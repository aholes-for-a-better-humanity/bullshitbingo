package widgets

import (
	"image/color"

	"github.com/rs/zerolog/log"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
)

var (
	mplusNormalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal().Err(err).Msg(``)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal().Err(err).Msg(``)
	}
}

var _ ui.Widget = &Text{}

type Text struct {
	Msg    string
	Bckgrd color.RGBA
}

func (t *Text) Update() error { return nil }
func (t *Text) Draw(screen *ebiten.Image) {
	screen.Fill(t.Bckgrd)
	textDims := text.BoundString(mplusNormalFont, t.Msg)
	text.Draw(screen, t.Msg, mplusNormalFont,
		screen.Bounds().Min.X+screen.Bounds().Dx()/2-textDims.Dx()/2,
		screen.Bounds().Min.Y+screen.Bounds().Dy()/2+textDims.Dy()/2, // origin is on the text baseline (bottom of letters)
		color.White)
	//log.Debug().Str(`txt`, t.Msg).Msg(`drawn`)
}
