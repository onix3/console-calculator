package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

var (
	g     *gocui.Gui
	vd    *gocui.View // окошко дисплея
	sx,sy int         // startX, startY
	dE    string      // выражение в дисплее
	dW    int         // актуальная ширина дисплея
)

// Запустить интерфейс в консоли
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
		// стандартно нажатия мыши работают только внутри рамки
		// Расширяю границы на 1, чтобы срабатывали и на рамке
		v,err := g.SetView(s, x0-1, y0-1, x1+1, y0+2+1)
		if err != nil && err != gocui.ErrUnknownView{
			log_err(err)
		}
		v.Frame = false
	}

	sx,sy = 3,2 // startX, startY
	by := sy+7 // номер ряда по вертикали, с которого начинается панель кнопок
	butt("7",sx+0, by);   butt("8",sx+5, by);    butt("9",sx+10, by);   butt("/",sx+15, by)
	butt("4",sx+0, by+3); butt("5",sx+5, by+3);  butt("6",sx+10, by+3); butt("*",sx+15, by+3)
	butt("1",sx+0, by+6); butt("2",sx+5, by+6);  butt("3",sx+10, by+6); butt("-",sx+15, by+6)
	butt("0",sx+0, by+9); butt(".",sx+10, by+9); butt("+",sx+15, by+9)

	for _,v := range g.Views() {
		s := "┌───┐\n│ " + v.Name() + " │\n└───┘"
		if v.Name() == "0" {
			s = "┌────────┐\n│   " + v.Name() + "    │\n└────────┘"
		}
		_,err = fmt.Fprint(v,s)
		if err != nil {
			log_err(err)
		}
	}

	// кнопка AC. Также статичная, поэтому не в layout
	ac,err := g.SetView("ac", sx-1, sy-1, sx+19+1, sy+3)
	if err != nil && err != gocui.ErrUnknownView{
		log_err(err)
	}
	ac.Frame = false
	fmt.Fprintln(ac,colorStr("┌──────────────────┐",0,1))
	fmt.Fprintln(ac,colorStr("│        AC        │",0,1))
	fmt.Fprintln(ac,colorStr("└──────────────────┘",0,1))

	initKeyboardBindings()
	initMouseBindings()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func max(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Функция, вызываемая для перерисовки
func layout(g *gocui.Gui) error {
	// ширина дисплея подстраивается под длины строк
	dW = 17
	dEb := addBrackets(dE)
	max_len := max(len(dEb),len(compute()))
	if max_len > 16 {
		dW = max_len+1
	}

	// display объявляется в layout, поскольку постоянно обновляется, в отличие от статичных кнопок
	var err error
	vd,err = g.SetView("display", sx, sy+3, sx+1+dW+1, sy+3+3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log_err(err)
		}
	}

	// рамка становится цветной только когда view в фокусе. Рамка дисплея в фокусе
	if _,err := g.SetCurrentView("display"); err != nil {
		log_err(err)
	}

	return nil
}