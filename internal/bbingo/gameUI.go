package bbingo

import (
	_ "embed"
	"image"
	"image/color"
	"strings"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui/widgets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog/log"
)

func NewGUI(head, foot string) *GUI {
	g := &GUI{}
	g.setHeader(head, ui.DarkGreen)
	g.setFooter(foot, ui.Sienne)
	g.B = &ui.GridUI{
		Columns: 5,
		Lines:   5,
		Widgets: make([]ui.Widget, 25),
	}
	for i := 0; i < len(g.B.Widgets); i++ {
		// rand.Shuffle(len(ui.Greys),
		// 	func(i, j int) { ui.Greys[i], ui.Greys[j] = ui.Greys[j], ui.Greys[i] })
		g.B.Widgets[i] = &widgets.Text{
			Msg: terms[i], Padding: 4,
			Bckgrd: ui.Red,
		}
		g.words = append(g.words, strings.ReplaceAll(terms[i], "\n", " "))
	}
	return g
}

var _ ui.Widget = &GUI{}

// GUI is the holder for the screen layout when playing
type GUI struct {
	H, B, F *ui.GridUI // Header, Body, Footer
	// header and footers are "one line"
	// grid in the middle (Body) holds the terms
	head, foot string
	words      []string
}

func (gui *GUI) setHeader(s string, c color.RGBA) {
	gui.H = &ui.GridUI{
		Columns: 1,
		Lines:   1,
		Widgets: []ui.Widget{&widgets.Text{
			Msg: s, Padding: 2,
			Bckgrd: c,
		}},
	}
}
func (gui *GUI) setFooter(s string, c color.RGBA) {
	gui.F = &ui.GridUI{
		Columns: 1,
		Lines:   1,
		Widgets: []ui.Widget{&widgets.Text{
			Msg: s, Padding: 2,
			Bckgrd: c,
		}},
	}
}
func (gui *GUI) Words() []string {
	return gui.words
}
func (gui *GUI) Update() error {
	return nil
}

// WordAt returns the word pressed at the given coordinates
func (gui *GUI) WordAt(x, y int) string {
	decal := image.Point{x - gui.B.Orig.X, y - gui.B.Orig.Y}
	line := decal.Y / (gui.B.ImH / gui.B.Lines)
	column := decal.X / (gui.B.ImW / gui.B.Columns)
	if 5*line+column > 24 {
		return ""
	}
	if widg, ok := gui.B.Widgets[5*line+column].(*widgets.Text); ok {
		w := widg.Msg
		log.Debug().Int("x", decal.X).Int("y", decal.Y).Int("line", line).Int("column", column).Str("w", w).Msg("(*GUI).WordAt")
		return w
	}
	return ""
}

// ColorWord colors the corresponding widget(s) with this colour
func (gui *GUI) ColorWord(word string, vl *validationLevel) {
	color := ui.Grey
	// TODO VARY COLOR
	// if self-touched, increase green
	// if validated, set Blue
	// in any case, set a "nice" Grey
	for pos, widg := range gui.B.Widgets {
		if widg, ok := widg.(*widgets.Text); ok {
			if MkRealWord(widg.Msg) != word {
				// log.Trace().Str("word", word).Int("pos", pos).Str("msg", widg.Msg).Send()
				continue
			}
			log.Info().Str("word", word).Send()
			color = ui.Greys[pos%len(ui.Greys)]
			widg.Bckgrd = color
			return
		}
	}
}

func (gui *GUI) Draw(screen *ebiten.Image) {
	_, h := screen.Size()
	headerHeight := h / 16
	if gui.H != nil {
		headerZone := image.Rectangle{
			Min: screen.Bounds().Min,
			Max: image.Point{X: screen.Bounds().Max.X, Y: headerHeight},
		}
		gui.H.Orig = headerZone.Min
		gui.H.Draw(screen.SubImage(headerZone).(*ebiten.Image))
	}
	if gui.F != nil {
		footerZone := image.Rectangle{
			Min: image.Point{X: screen.Bounds().Min.X, Y: screen.Bounds().Max.Y - headerHeight},
			Max: screen.Bounds().Max,
		}
		gui.F.Orig = footerZone.Min
		gui.F.Draw(screen.SubImage(footerZone).(*ebiten.Image))
	}
	if gui.B != nil {
		bodyZone := image.Rectangle{
			Min: image.Point{X: screen.Bounds().Min.X, Y: headerHeight},
			Max: image.Point{X: screen.Bounds().Max.X, Y: screen.Bounds().Max.Y - headerHeight},
		}
		gui.B.Orig = bodyZone.Min
		gui.B.Draw(screen.SubImage(bodyZone).(*ebiten.Image))
	}
}
