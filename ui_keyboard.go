package main

import (
	"github.com/jroimartin/gocui"
)

func keyboardBinding(key rune) {
	if err := g.SetKeybinding("", key, gocui.ModNone,
		func(g *gocui.Gui, _ *gocui.View) error {
			replyKey(g,key)
			return nil
		});
	err != nil {
		log_err(err)
	}
}

func initKeyboardBindings() {
	for _,c := range "789/456*123-0+" {
		keyboardBinding(c)
	}

	if err := g.SetKeybinding("", gocui.KeyBackspace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if len(dS) > 0 {
				dS = dS[:len(dS)-1]
			}
			g.Update(updateVD)
			return nil
		});
		err != nil {
		log_err(err)
	}
}