package main

import (
	"math/rand"
	"time"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui/widgets"
)

func runBBGame(UI *ui.GridUI) {
	Hydrate(UI)
}

func Hydrate(UI *ui.GridUI) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	UI.Columns = 5
	UI.Lines = 5

	UI.Widgets = make([]ui.Widget, UI.Columns*UI.Lines)
	for i := 0; i < len(UI.Widgets); i++ {
		UI.Widgets[i] = &widgets.Text{Msg: words[i], Pad: 8, Bckgrd: ui.Colors[i%len(ui.Colors)]}
		UI.Widgets[i].PreloadBbox()
	}
	UI.Widgets[12].(*widgets.Text).Msg = "."
	UI.Widgets[12].PreloadBbox()
}
