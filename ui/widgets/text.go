package widgets

import (
	_ "embed"
	"image"
	"image/color"
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
	"github.com/hajimehoshi/ebiten/v2"

	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed Lato-Regular.ttf
var myfont []byte

var (
	fontsL []font.Face
)

func init() {
	tt, err := opentype.Parse(myfont)
	if err != nil {
		log.Fatal().Err(err).Msg(``)
	}
	const dpi = 72
	var fsizes []float64
	for i := 100.00; i > 0; i -= 2.0 {
		fsizes = append(fsizes, i)
	}
	fontsL = make([]font.Face, len(fsizes))
	for i, fsize := range fsizes {
		var err error
		fontsL[i], err = opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    fsize,
			DPI:     dpi,
			Hinting: font.HintingNone,
		})
		if err != nil {
			log.Fatal().Err(err).Msg(``)
		}
	}

}

var _ ui.Widget = &Text{}

type Text struct {
	Msg     string
	Bckgrd  color.RGBA
	Padding int // pixel value f the padding in p
	lines   int // number of lines of text (if wrapped)
	Fsize   float64
	bboxL   []image.Rectangle
	o       sync.Once
}

func (t *Text) Preload() {
	//log.Debug().Str("text", t.Msg).Msg("preloading")
	t.lines = 1
	for _, run := range t.Msg {
		if run == '\n' {
			t.lines++
		}
	}
	t.bboxL = make([]image.Rectangle, len(fontsL))
	for i := 0; i < len(fontsL); i++ {
		t.bboxL[i] = text.BoundString(fontsL[i], t.Msg)
	}
}

func (t *Text) Update() error { t.o.Do(t.Preload); return nil }
func (t *Text) Draw(screen *ebiten.Image) {
	t.o.Do(t.Preload)
	// log.Debug().Str("text", t.Msg).Msg("drawing")
	screen.Fill(t.Bckgrd)
	var textDims image.Rectangle
	var fontFace font.Face
	for i := 0; i < len(fontsL); i++ {
		fontFace = fontsL[i]
		textDims = t.bboxL[i]
		// textDims = text.BoundString(fontFace, t.Msg)
		if textDims.Dx() < screen.Bounds().Dx()-t.Padding*2 && textDims.Dy() < screen.Bounds().Dy()-t.Padding*2 {
			break
		}
	}

	// draws a bouding box of the text
	// ebitenutil.DrawRect(screen,
	// 	float64(screen.Bounds().Min.X+(screen.Bounds().Dx()-textDims.Dx())/2),
	// 	float64(screen.Bounds().Min.Y+(screen.Bounds().Dy()-textDims.Dy())/2),
	// 	float64(textDims.Dx()), float64(textDims.Dy()), color.RGBA{0, 0, 0, 0xFF})
	vcorrect := 0
	if t.lines > 1 {
		vcorrect = textDims.Min.Y * t.lines
	}
	decal := struct{ X, Y int }{ // origin vector to place text in center
		X: screen.Bounds().Min.X - textDims.Min.X,
		Y: screen.Bounds().Min.Y + vcorrect, // Align baselines (do not correct vertically with textDims.Min.Y)
	}
	text.Draw(screen, t.Msg, fontFace,
		decal.X+(screen.Bounds().Dx()-textDims.Dx())/2,
		decal.Y+(screen.Bounds().Dy()+textDims.Dy())/2,
		color.White) // textDims is at first character position, so pixels start at Min.X,Min.Y
}
