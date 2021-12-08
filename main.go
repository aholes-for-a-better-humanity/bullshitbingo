package main

import (
	"image/color"
	"os"
	"time"

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
	UI := &ui.GridUI{
		Columns: 2,
		Lines:   2,
		Widgets: []ui.Widget{
			&widgets.Text{Msg: `bull`, Bckgrd: color.RGBA{0x7f, 0xff, 0x7f, 0xae}},
			&widgets.Text{Msg: `shit`, Bckgrd: color.RGBA{0xff, 0x7f, 0x7f, 0xae}},
			&widgets.Text{Msg: `bin`, Bckgrd: color.RGBA{0xff, 0x7f, 0x3f, 0xae}},
			&widgets.Text{Msg: `Go`, Bckgrd: color.RGBA{0x7f, 0x7f, 0x7f, 0xae}},
		},
	}
	for i := range UI.Widgets {
		go UI.Widgets[i].Preload()
	}

	go func() { time.Sleep(5 * time.Second); runBBGame(UI) }()
	err := ebiten.RunGame(UI)
	if err != nil {
		log.Fatal().Err(err).Msg(`exiting`)
	}
}
