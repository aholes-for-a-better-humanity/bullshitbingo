package ui

import (
	"image"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/rs/zerolog/log"
)

var _ ebiten.Game = &GridUI{}

// GridUI a versatile structure for a grid-layout-based game screen.
type GridUI struct {
	Columns, Lines int         // dimensions of the grid
	Widgets        []Widget    // widgets in Grid
	ImW, ImH       int         // width and height of the layout
	Orig           image.Point // origin of the grid (relative to screen)
	sync.Mutex
}

// https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2@v2.2.2?utm_source=gopls#Game

// Update updates a game by one tick.
func (ui *GridUI) Update() error { return nil }

// Draw draws the game screen by one frame.
//
// The give argument represents a screen image. The updated content is adopted as the game screen.
func (ui *GridUI) Draw(screen *ebiten.Image) {
	w, h := screen.Size()
	ui.Layout(w, h)
	switch len(ui.Widgets) {
	case 0:
		ebitenutil.DebugPrint(screen, `no widget in grid`)
	case 1:
		ui.Widgets[0].Draw(screen)
	default:
		ui.Lock()
		defer ui.Unlock()
		for i := range ui.Widgets {
			ui.Widgets[i].Draw(ui.CellAt(screen, i))
		}
	}
}

// Layout accepts a native outside size in device-independent pixels and returns the game's logical screen
// size.
func (ui *GridUI) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	ui.ImW, ui.ImH = outsideWidth, outsideHeight
	return outsideWidth, outsideHeight
}

// CellAt returns the portion of the source image that is in the cell numbered i.
//
// numbers are laid out line by line, from left to right, lines laid out top to bottom.
// A 4*3 grid would be numbered :
//
//	// 0  1  2  3
//	// 4  5  6  7
//	// 8  9 10 11
//
// Warning, if you increase size of a cell, the next cell will have an offset, too
func (ui *GridUI) CellAt(screen *ebiten.Image, i int) *ebiten.Image {
	lin, col := gridPos(i, ui.Lines, ui.Columns)

	remainderX := ui.ImW % ui.Columns
	remainderY := ui.ImH % ui.Lines
	var adjX int
	var adjY int
	var adjW int
	var adjH int
	if col < remainderX {
		adjW = 1
		adjX = col
	} else {
		adjX = remainderX
	}
	if lin < remainderY {
		adjH = 1
		adjY = lin
	} else {
		adjY = remainderY
	}

	cellWidth := ui.ImW / ui.Columns
	cellHeight := ui.ImH / ui.Lines
	crop := image.Rectangle{
		Min: image.Point{
			X: adjX + col*cellWidth + screen.Bounds().Min.X,
			Y: adjY + lin*cellHeight + screen.Bounds().Min.Y,
		},
		Max: image.Point{
			X: adjX + adjW + (col+1)*cellWidth + screen.Bounds().Min.X,
			Y: adjY + adjH + (lin+1)*cellHeight + screen.Bounds().Min.Y,
		},
	}
	// log.Debug().
	// 	Int(`colWidth`, colWidth).
	// 	Int(`linWidth`, linWidth).
	// 	Interface(`Min`, crop.Min).
	// 	Interface(`Max`, crop.Max).Msg(``)
	return screen.SubImage(crop).(*ebiten.Image)
}

// gridPos returns the position of a cell by its number.
func gridPos(index, lines, cols int) (lin, col int) {
	if index >= lines*cols || cols == 0 {
		log.Warn().Msg(`out of bonds`)
		return -1, -1
	}
	return index / cols, index % cols
}
