package ui

import "github.com/hajimehoshi/ebiten/v2"

// Widget is a elemnt of the UI.
//
// Combined together, they allow building more complex UIs,
// using simple building blocks.
type Widget interface {
	Update() error
	Draw(screen *ebiten.Image)
	PreloadBbox()
}
