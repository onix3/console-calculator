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
	g.Mouse = true

	g.SetManagerFunc(layout)

	// создание кнопки
	butt := func(s string, x0,y0 int) {
		x1 := x0+4
		if s == "0" {
			x1 = 14
		}
		// стандартно нажатия мыши работают только внутри Frame
		// Расширяю границы на 1, чтобы срабатывали и на рамке кнопок
		v,err := g.SetView(s, x0-1, y0-1, x1+1, y0+2+1)
		if err != nil && err != gocui.ErrUnknownView{
			log_err(err)
		}
		v.Frame = false
	}

	y := 4 // номер ряда по вертикали, с которого начинается панель кнопок
	butt("7",0, y);   butt("8",5, y);   butt("9",10, y);   butt("/",15, y)
	butt("4",0, y+3); butt("5",5, y+3); butt("6",10, y+3); butt("*",15, y+3)
	butt("1",0, y+6); butt("2",5, y+6); butt("3",10, y+6); butt("-",15, y+6)
	butt("0",0, y+9); butt("+",15, y+9)

	//butt("7",0, y);   butt("8",5, y);   butt("9",10, y);   butt("/",15, y)
	//butt("4",0, y+3); butt("5",5, y+3); butt("6",10, y+3); butt("*",15, y+3)
	//butt("1",0, y+6); butt("2",5, y+6); butt("3",10, y+6); butt("-",15, y+6)
	//butt("0",0, y+9); butt("+",15, y+9)

	for _,v := range g.Views() {
		s := "┌───┐\n│ " + v.Name() + " │\n└───┘"
		if v.Name() == "0" {
			s = "┌─────────────┐\n│      " + v.Name() + "      │\n└─────────────┘"
		}
		_,err = fmt.Fprint(v,s)
		if err != nil {
			log_err(err)
		}
	}

	initKeyboardBindings()
	initMouseBindings()

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
	// границы дисплея расширены, чтобы я мог сам вывести не одинарные, а двойные линии
	vd,err = g.SetView("display", 0, 0, 1+displayW+1, 3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log_err(err)
		}
	}
	//vd.Frame = false

	if _,err := g.SetCurrentView("display"); err != nil {
		log_err(err)
	}

	return nil
}