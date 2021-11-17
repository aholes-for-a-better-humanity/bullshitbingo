package main

import (
	"os"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
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

func main() {
	err := ebiten.RunGame(&ui.GridUI{})
	if err != nil {
		log.Fatal().Err(err).Msg(`exiting`)
	}
}
