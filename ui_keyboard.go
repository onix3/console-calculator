package main

import (
	"github.com/jroimartin/gocui"
)

// Привязка нажатия на клавиатуре
func keyboardBinding(key rune) {
	if err := g.SetKeybinding("", key, gocui.ModNone,
		func(g *gocui.Gui, _ *gocui.View) error {
			// при нажатии ',' вводить '.'
			if key == ',' { key = '.' }
			replyKey(g,key)
			return nil
		});
	err != nil {
		log_err(err)
	}
}

// Привязка нажатий на клавиатуре
func initKeyboardBindings() {
	for _,c := range "789/456*123-0.+," {
		keyboardBinding(c)
	}

	if err := g.SetKeybinding("", gocui.KeyBackspace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if len(dE) > 0 {
				dE = dE[:len(dE)-1]
				g.Update(updateVD)
			}
			return nil
		});
		err != nil {
		log_err(err)
	}
}