package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

var (
	g  *gocui.Gui
	vd *gocui.View // windows of "calculator's display"
)

func go_ui() {
	var err error
	if g,err = gocui.NewGui(gocui.OutputNormal); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	g.SetManagerFunc(layout)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	//W,H := g.Size()

	var err error
	vd,err = g.SetView("display", 0, 0, 19, 3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log_err(err)
		}
	}

	if _,err := g.SetCurrentView("display"); err != nil {
		log_err(err)
	}

	y := 4 // номер ряда по вертикали, с которого начинается панель кнопок
	butt(g," 7",0, y); butt(g," 8",5, y); butt(g," 9",10, y); butt(g," /",15, y)
	butt(g," 4",0, y+3); butt(g," 5",5, y+3); butt(g," 6",10, y+3); butt(g," *",15, y+3)
	butt(g," 1",0, y+6); butt(g," 2",5, y+6); butt(g," 3",10, y+6); butt(g," -",15, y+6)
	butt(g,"   0 ",0, y+9); butt(g," .",10, y+9); butt(g," +",15, y+9)

	for _,v := range g.Views()[1:] {
		_,err = fmt.Fprint(v,v.Name())
		if err != nil {
			log_err(err)
		}
	}

	return nil
}

// кнопка с символом
func butt(g *gocui.Gui, s string, x0,y0 int) {
	x1 := x0+4
	if s == "   0 " {
		x1 = 9
	}
	_,err := g.SetView(s, x0, y0, x1, y0+2)
	if err != nil && err != gocui.ErrUnknownView{
		log_err(err)
	}
}