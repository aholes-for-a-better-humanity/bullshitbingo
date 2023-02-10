package bbingo

import (
	"context"
	"errors"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui/widgets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/pioz/faker"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

var Finished = errors.New("ended")
var invalidationDelay = 10 * time.Second // time we wait before invalidating a word checked, if no one confirms
var quorum = .5

type gameEventSignal int

const (
	gameEventStartPlay        gameEventSignal = iota // no extra data
	gameEventRandomizeSplash                         // no extra data
	gameEventEndRun                                  // no extra data
	gameEventMsg                                     // ticker message, has `payload`,`color`,`duration`
	gameWordPressed                                  // word pressed, has `word`
	gameWordPressInvalidation                        // word previously pressed by this user is invalidated, has `word`
	gameWordPressedByOther                           // word pressed by someone else, has `word`
)

type gameEvent struct {
	sig     gameEventSignal // what is it
	word    string          // what word does it refer to
	sender  string          // who sent it
	payload string          // does it have extra info
	color   color.RGBA      // does it convey color
	dur     time.Duration   // does it have a duration ?
}

var _ ebiten.Game = &Game{}

// Game should hold a single Widget
// this widget is updated and drawn by Draw and Update
type Game struct {
	initOnce         sync.Once
	eg               *errgroup.Group
	ctx              context.Context
	widget           ui.Widget
	events           chan gameEvent
	toNetwork        chan netMsg // for things to publish towards the others
	touchIDs         []ebiten.TouchID
	cursorX, cursorY int // position of last touch or click
	//
	nickname    string
	gameWeAreIn string                      // the game we are in
	ourWords    map[string]*validationLevel // list of reference expressions (unsplitted) in the GUI
	whoHas      map[string][]string         // key:(real)word, value:nicknames of posessors
}
type presenceMap = map[string]bool
type netMsg struct{ topic, content string }
type validationLevel struct {
	Validated     bool
	Self          int    // have we pressed this word ourselves ?
	total         int    // how many players (self included) have this word ?
	OthersPressed int    // how many others have pressed it ?
	Word          string // FIXME remove (test)
}

// Init prepares the game, and sets up the lifecycle
func (g *Game) Init() { g.initOnce.Do(g.init) }

// init is called once. It initializes g, and makes it
// "live" async in between calls to Draw and Update
func (g *Game) init() {
	g.events = make(chan gameEvent, 4)  // buffer to avoid blocking
	g.toNetwork = make(chan netMsg, 64) // same here
	g.ourWords = make(map[string]*validationLevel, 25)
	g.whoHas = make(map[string][]string)
	// ctx, _ := context.WithCancel(context.Background())
	g.eg, g.ctx = errgroup.WithContext(context.Background())
	g.nickname = strings.Join([]string{faker.Username(), faker.Username()}, " ")
	g.eg.Go(g.lifecycle)
	// initial widget is the splash/welcome screen
	bsbg := strings.Split("bull shit bin Go", " ")
	g.widget = &ui.GridUI{
		Columns: 2,
		Lines:   2,
		Widgets: []ui.Widget{
			&widgets.Text{Msg: bsbg[0], Bckgrd: color.RGBA{0x7f, 0xff, 0x7f, 0xae}},
			&widgets.Text{Msg: bsbg[1], Bckgrd: color.RGBA{0xff, 0x7f, 0x7f, 0xae}},
			&widgets.Text{Msg: bsbg[2], Bckgrd: color.RGBA{0xff, 0x7f, 0x3f, 0xae}},
			&widgets.Text{Msg: bsbg[3], Bckgrd: color.RGBA{0x7f, 0x7f, 0x7f, 0xae}},
		},
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.Init()
	if g.isKeyJustPressed() { // process touch/click
		log.Debug().Int("X", g.cursorX).Int("Y", g.cursorY).Send()
		if gui, ok := g.widget.(*GUI); ok {
			word := gui.WordAt(g.cursorX, g.cursorY)
			g.events <- gameEvent{
				sig:  gameWordPressed,
				word: word,
			}
		}
	}
	select {
	case <-g.ctx.Done():
		log.Fatal().Caller().Err(g.ctx.Err()).Send()
		return g.ctx.Err()
	default:
		return nil
	}
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.Init()
	// Write your game's rendering.
	if g.widget != nil {
		g.widget.Draw(screen)
	}
}

func (g *Game) isKeyJustPressed() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.cursorX, g.cursorY = ebiten.CursorPosition()
		return true
	}
	g.touchIDs = inpututil.AppendJustPressedTouchIDs(g.touchIDs)
	if len(g.touchIDs) > 0 { // use position of last touch
		g.cursorX, g.cursorY = ebiten.TouchPosition(g.touchIDs[len(g.touchIDs)-1])
		return true
	}
	return false
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// lifecycle runs as long as the game is on screen
func (g *Game) lifecycle() error {
	gui := NewGUI(g.nickname, "")
	for _, w := range gui.Words() {
		g.ourWords[w] = &validationLevel{Word: w}
	}
	g.eg.Go(g.timeline)
	g.eg.Go(g.sigCatcher)
	g.eg.Go(g.network)
	for {
		select {
		case <-g.ctx.Done():
			log.Error().Err(g.ctx.Err()).Msg("")
			return g.ctx.Err()
		case ev := <-g.events:
			switch ev.sig {
			case gameEventRandomizeSplash:
				g.randomizeSplash()
			case gameEventStartPlay:
				// go to "game mode"
				g.widget = gui
				log.Info().Int("words", len(g.ourWords)).Send()
				for w := range g.ourWords {
					log.Info().Str("w", w).Msg("coloring")
					g.ColorWord(w, g.ourWords[w])
				}
				log.Info().Msg("done colouring")
			case gameEventEndRun:
				log.Debug().Msg("byebye")
				return Finished // an error here closes the context and ends the program
			case gameEventMsg:
				if gui, ok := g.widget.(*GUI); ok {
					gui.setFooter(ev.payload, ev.color)
					g.eg.Go(func() error { time.Sleep(ev.dur); gui.setFooter("", ev.color); return nil })
				}
			case gameWordPressed:
				// a word has just been touched/pressed
				g.gameWordPressed(ev.word)
			case gameWordPressedByOther:
				realWord := strings.ReplaceAll(ev.word, "\n", " ")
				// register external press
				g.ourWords[realWord].OthersPressed++ // TODO if pressed twice by same player, don't increment again
				if float64(g.ourWords[realWord].Self+g.ourWords[realWord].OthersPressed) > float64(g.ourWords[realWord].total)*quorum {
					// check validation if more than quorum
					g.ourWords[realWord].Validated = true
				}
				log.Debug().Msg("press by other")
				g.ColorWord(ev.word, g.ourWords[realWord]) // TODO validate also for sender when quorum is reached
			case gameWordPressInvalidation:
				g.invalidate(ev.word)
			default:
			}
		}
	}
}

func (g *Game) invalidate(word string) {
	realWord := strings.ReplaceAll(word, "\n", " ")
	r := g.ourWords[realWord]
	r.Self = 0
	g.ColorWord(word, r)
}

// ColorWord fetches necessary info for a word and calls eponymous method from the gui
// word argument is RAW, with line feeds.
func (g *Game) ColorWord(word string, vl *validationLevel) {
	log.Debug().Caller().
		Interface("vl", vl).
		Str("word", word).
		Msg("(g *Game) ColorWord(word string)")
	if gui, ok := g.widget.(*GUI); ok {
		gui.ColorWord(word, vl)
	}
}

// gameWordPressed is inviked when *we* press a word onscreen
func (g *Game) gameWordPressed(word string) error {
	realWord := strings.ReplaceAll(word, "\n", " ")
	// register state
	v, ok := g.ourWords[realWord]
	if !ok {
		log.Warn().Str("rw", realWord).Msg("NOT OK")
		return nil
	}
	g.ourWords[realWord] = v
	// FIXME pointers or smth makes sharing the value NOK
	log := log.With().Str(".", "(g *Game) gameWordPressed").Logger()
	log.Debug().Interface("vl", g.ourWords[realWord]).Interface("v", v).Send()
	if (g.ourWords[realWord].Self == 1) && (!g.ourWords[realWord].Validated) {
		log.Warn().Msg("unchecking")
		v = g.ourWords[realWord]
		v.Self = 0
	} else {
		v.Self = 1
		// communicate (beware:split lines in result of WordAt)
		g.eg.Go(func() error { g.toNetwork <- netMsg{topic: pressedWordsTopic, content: realWord}; return nil })
		g.eg.Go(func() error { return g.invalidateLater(realWord) })
	}
	// show
	g.ColorWord(word, v)
	return nil
}

//invalidateLater unchecks a word after a while
func (g *Game) invalidateLater(word string) error {
	select {
	case <-g.ctx.Done():
		return g.ctx.Err()
	case <-time.After(invalidationDelay):
		log.Debug().Str("word_to_invalidate", word).Send()
		if !g.ourWords[word].Validated {
			g.events <- gameEvent{sig: gameWordPressInvalidation, word: word}
		}
		return nil
	}
}

// timeline sends events to g.event when it's time to change life phase
func (g *Game) timeline() error {
	splash := time.NewTimer(1500 * time.Millisecond)
	randsplashDelay := 800 * time.Millisecond
splashScreen:
	for {
		select {
		case <-splash.C:
			g.events <- gameEvent{sig: gameEventStartPlay}
			break splashScreen
		case <-time.After(randsplashDelay):
			randsplashDelay = randsplashDelay / time.Duration(2)
			g.events <- gameEvent{sig: gameEventRandomizeSplash}
		}
	}
	g.eg.Go(g.footerTimeline)

	// wait for end of run.
	select {
	case <-g.ctx.Done():
		log.Debug().Err(g.ctx.Err()).Send()
		return g.ctx.Err()
	}
}

//footerTimeline ticks messages in the first seconds
func (g *Game) footerTimeline() error {
	type tickMsg struct {
		durMillis int64
		m         string
		color     color.RGBA
	}
	ticks := []tickMsg{
		{500, "Hello", ui.Green},
		{1500, "Enjoy your game", ui.Red},
		{2000, "Your nickname is on top", ui.Green},
		{2000, "your nickname is random", ui.Green},
		{1, "", ui.Grey},
	}
	for _, tick := range ticks {
		g.events <- gameEvent{
			sig:     gameEventMsg,
			payload: tick.m,
			color:   tick.color,
			dur:     time.Duration(tick.durMillis) * time.Millisecond,
		}
		select {
		case <-time.After(time.Duration(tick.durMillis) * time.Millisecond):
			continue
		case <-g.ctx.Done():
			return g.ctx.Err()
		}
	}
	return nil
}

// randomizeSplash shuffles the widgets inside the top-level widget.
// It is used to make the splash screen look alive at startup
func (g *Game) randomizeSplash() {
	if gr, ok := g.widget.(*ui.GridUI); ok {
		rand.Shuffle(
			len(gr.Widgets),
			func(i, j int) {
				gr.Widgets[i], gr.Widgets[j] = gr.Widgets[j], gr.Widgets[i]
			},
		)
	}
}

func (g *Game) sigCatcher() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-g.ctx.Done():
		log.Debug().Err(g.ctx.Err()).Send()
		return g.ctx.Err()
	case s := <-c:
		return fmt.Errorf(`signal: %s`, s)
	}
}
