package main

import (
	"image/color"
	"os"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui/widgets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	{ // set up logging
		consolelog := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05.000"}
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
		log.Logger = log.Output(consolelog)
	}
}

func init() {
	ebiten.SetWindowResizable(true)
}

func main() {
	err := ebiten.RunGame(&ui.GridUI{
		Columns: 2,
		Lines:   2,
		Widgets: []ui.Widget{
			&widgets.Text{Msg: `bull`, Bckgrd: color.RGBA{0x7f, 0xff, 0x7f, 0xae}},
			&widgets.Text{Msg: `shit`, Bckgrd: color.RGBA{0xff, 0x7f, 0x7f, 0xae}},
			&widgets.Text{Msg: `bin`, Bckgrd: color.RGBA{0xff, 0x7f, 0x3f, 0xae}},
			&widgets.Text{Msg: `Go`, Bckgrd: color.RGBA{0x7f, 0x7f, 0x7f, 0xae}},
		},
	})
	if err != nil {
		log.Fatal().Err(err).Msg(`exiting`)
	}
}
