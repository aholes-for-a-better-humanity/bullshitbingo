package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/internal/bbingo"
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
	rand.Seed(time.Now().UnixNano())
}

func init() {
	ebiten.SetWindowResizable(true)
}

func main() {
	log.Error().Err(ebiten.RunGame(&bbingo.Game{})).Msg(`exiting`)
}
