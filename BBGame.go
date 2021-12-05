package main

import (
	"math/rand"

	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui"
	"github.com/aholes-for-a-better-humanity/bullshitbingo/ui/widgets"
)

func runBBGame(UI *ui.GridUI) {
	Hydrate(UI)
}

var words = []string{
	"aggregate",
	"architect",
	"benchmark",
	"brand",
	"cultivate",
	"deliver",
	"deploy",
	"disintermediate",
	"drive",
	"e-enable",
	"embrace",
	"empower",
	"enable",
	"engage",
	"engineer",
	"enhance",
	"envisioneer",
	"evolve",
	"expedite",
	"exploit",
	"extend",
	"facilitate",
	"generate",
	"grow",
	"harness",
	"implement",
	"incentivize",
	"incubate",
	"innovate",
	"integrate",
	"iterate",
	"leverage",
	"matrix",
	"maximize",
	"mesh",
	"monetize",
	"morph",
	"optimize",
	"orchestrate",
	"productize",
	"recontextualize",
	"redefine",
	"reintermediate",
	"reinvent",
	"repurpose",
	"revolutionize",
	"scale",
	"seize",
	"strategize",
	"streamline",
	"syndicate",
	"synergize",
	"synthesize",
	"target",
	"transform",
	"transition",
	"unleash",
	"utilize",
	"visualize",
	"whiteboard",
	"24/365",
	"24/7",
	"B2B",
	"B2C",
	"back-end",
	"best-of-breed",
	"bleeding-edge",
	"bricks-and-clicks",
	"clicks-and-mortar",
	"collaborative",
	"compelling",
	"cross-platform",
	"cross-media",
	"customized",
	"cutting-edge",
	"distributed",
	"dot-com",
	"dynamic",
	"e-business",
	"efficient",
	"end-to-end",
	"enterprise",
	"extensible",
	"frictionless",
	"front-end",
	"global",
	"granular",
	"holistic",
	"impactful",
	"innovative",
	"integrated",
	"interactive",
	"intuitive",
	"killer",
	"leading-edge",
	"magnetic",
	"mission-critical",
	"next-generation",
	"one-to-one",
	"open-source",
	"out-of-the-box",
	"plug-and-play",
	"proactive",
	"real-time",
	"revolutionary",
	"rich",
	"robust",
	"scalable",
	"seamless",
	"sexy",
	"sticky",
	"strategic",
	"synergistic",
	"transparent",
	"turn-key",
	"ubiquitous",
	"user-centric",
	"value-added",
	"vertical",
	"viral",
	"virtual",
	"visionary",
	"web-enabled",
	"wireless",
	"world-class",
	"action-items",
	"applications",
	"architectures",
	"bandwidth",
	"channels",
	"communities",
	"content",
	"convergence",
	"deliverables",
	"e-business",
	"e-commerce",
	"e-markets",
	"e-services",
	"e-tailers",
	"experiences",
	"eyeballs",
	"functionalities",
	"infomediaries",
	"infrastructures",
	"initiatives",
	"interfaces",
	"markets",
	"methodologies",
	"metrics",
	"mindshare",
	"models",
	"networks",
	"niches",
	"paradigms",
	"partnerships",
	"platforms",
	"portals",
	"relationships",
	"ROI",
	"synergies",
	"web-readiness",
	"schemas",
	"solutions",
	"supply-chains",
	"systems",
	"technologies",
	"users",
	"vortals",
	"web services",
}

func Hydrate(UI *ui.GridUI) {
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	UI.Columns = 5
	UI.Lines = 5
	UI.Widgets = make([]ui.Widget, 25)
	for i := 0; i < 25; i++ {
		UI.Widgets[i] = &widgets.Text{Msg: words[i]}
	}
}
