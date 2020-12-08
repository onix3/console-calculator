package src

import (
	"github.com/jroimartin/gocui"
	"time"
)

// Подсветка кнопки при нажатии
func highlight(key rune) {
	for _,v := range g.Views() {
		if v.Name() == string(key) || v.Name() == "ac" {
			go func() {
				v.BgColor = gocui.ColorMagenta
				time.Sleep(100*time.Millisecond)
				v.BgColor = gocui.ColorDefault
				g.Update(updateVD)
			}()
			break
		}
	}
}

// Подсветка дисплея при недопустимом вводе
func highlight_error() {
	go func() {
		vd.BgColor = gocui.ColorRed
		time.Sleep(100*time.Millisecond)
		vd.BgColor = gocui.ColorBlack
		g.Update(updateVD)
	}()
}