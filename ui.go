package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

var (
	g  *gocui.Gui
	vd *gocui.View // windows of "calculator's display"
	displayS string
	displayW int // width of display
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

	// создание кнопки
	butt := func(s string, x0,y0 int) {
		x1 := x0+4
		if s == "0" {
			x1 = 14
		}
		_,err := g.SetView(s, x0, y0, x1, y0+2)
		if err != nil && err != gocui.ErrUnknownView{
			log_err(err)
		}
	}

	y := 4 // номер ряда по вертикали, с которого начинается панель кнопок
	butt("7",0, y);   butt("8",5, y);   butt("9",10, y);   butt("/",15, y)
	butt("4",0, y+3); butt("5",5, y+3); butt("6",10, y+3); butt("*",15, y+3)
	butt("1",0, y+6); butt("2",5, y+6); butt("3",10, y+6); butt("-",15, y+6)
	butt("0",0, y+9); butt("+",15, y+9)

	for _,v := range g.Views() {
		s := " " + v.Name()
		if v.Name() == "0" {
			s = "      " + v.Name()
		}
		_,err = fmt.Fprint(v,s)
		if err != nil {
			log_err(err)
		}
	}

	initKeyboardBindings()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	//W,H := g.Size()

	displayW = 17
	if len(displayS) > 16 {
		displayW = len(displayS)+1
	}

	var err error
	vd,err = g.SetView("display", 0, 0, 1+displayW+1, 3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log_err(err)
		}
	}

	if _,err := g.SetCurrentView("display"); err != nil {
		log_err(err)
	}

	return nil
}