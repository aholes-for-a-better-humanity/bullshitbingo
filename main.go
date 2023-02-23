package main

import (
	"os"
	"runtime"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/internal/bbingo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	{ // set up logging
		consolelog := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05.000", NoColor: (runtime.GOOS == "js")}
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
		log.Logger = log.Output(consolelog)
	}
}

func init() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowClosingHandled(true)
}

func main() {
	log.Error().Err(ebiten.RunGame(&bbingo.Game{})).Msg(`exiting`)
}
