package ui

import "github.com/hajimehoshi/ebiten/v2"

var _ ebiten.Game = &GridUI{}

// GridUI a versatile structure for a grid-layout-based game screen.
type GridUI struct{}

// https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2@v2.2.2?utm_source=gopls#Game

// Update updates a game by one tick.
func (ui *GridUI) Update() error { return nil }

// Draw draws the game screen by one frame.
//
// The give argument represents a screen image. The updated content is adopted as the game screen.
func (ui *GridUI) Draw(screen *ebiten.Image) {}

// Layout accepts a native outside size in device-independent pixels and returns the game's logical screen
// size.
func (ui *GridUI) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
