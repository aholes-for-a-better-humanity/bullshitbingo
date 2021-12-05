package main

import (
	"math/rand"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui/widgets"
)

func runBBGame(UI *ui.GridUI) {
	Hydrate(UI)
}

func Hydrate(UI *ui.GridUI) {
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	UI.Columns = 5
	UI.Lines = 5
	UI.Widgets = make([]ui.Widget, 25)
	for i := 0; i < 25; i++ {
		UI.Widgets[i] = &widgets.Text{Msg: words[i], Pad: 4, Bckgrd: ui.Colors[i%len(ui.Colors)]}
	}
}
