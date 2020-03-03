package main

import (
	"log"

	"github.com/hanjunlee/gocui-architecture/internal/ui"
	"github.com/hanjunlee/gocui-architecture/pkg/service"

	"github.com/jroimartin/gocui"
)

func main() {
	g := newGui()
	defer g.Close()

	m := &ui.Manager{
		Svc: &service.MartService{},
		SvcEntries: map[string]service.UseCase{
			"mart":   &service.MartService{},
			"zoo": &service.ZooService{},
		},
	}
	g.SetManager(m)
	m.Keybinding(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panic(err)
	}
}

func newGui() *gocui.Gui {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panic(err)
	}

	g.Highlight = false
	g.InputEsc = true
	return g
}
