package widgets

import (
	"image"
	"image/color"

	"github.com/rs/zerolog/log"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	fontsL []font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal().Err(err).Msg(``)
	}
	const dpi = 72
	var fsizes []float64
	for i := 100.00; i > 0; i -= 0.25 {
		fsizes = append(fsizes, i)
	}
	fontsL = make([]font.Face, len(fsizes))
	for i, fsize := range fsizes {
		var err error
		fontsL[i], err = opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    fsize,
			DPI:     dpi,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Fatal().Err(err).Msg(``)
		}
	}

}

var _ ui.Widget = &Text{}

type Text struct {
	Msg    string
	Bckgrd color.RGBA
	Pad    int // pixel value f the padding in p
	Fsize  float64
}

func (t *Text) Update() error { return nil }
func (t *Text) Draw(screen *ebiten.Image) {
	screen.Fill(t.Bckgrd)
	var textDims image.Rectangle
	var fontFace font.Face
	for i := 0; i < len(fontsL); i++ {
		fontFace = fontsL[i]
		textDims = text.BoundString(fontFace, t.Msg)
		if textDims.Dx() < screen.Bounds().Dx()-t.Pad*2 && textDims.Dy() < screen.Bounds().Dy()-t.Pad*2 {
			break
		}
	}
	text.Draw(screen, t.Msg, fontFace,
		screen.Bounds().Min.X+screen.Bounds().Dx()/2-textDims.Dx()/2,
		screen.Bounds().Min.Y+screen.Bounds().Dy()/2+textDims.Dy()/2, // origin is on the text baseline (bottom of letters)
		color.White)
	//log.Debug().Str(`txt`, t.Msg).Msg(`drawn`)
}
