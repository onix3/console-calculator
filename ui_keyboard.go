package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func keyboardBinding(key rune) {
	if err := g.SetKeybinding("", key, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			W,_ := g.Size()
			if 2+len(displayS)+2 < W {
				displayS += string(key)
				g.Update(updateVD)
			} else {
				fmt.Printf("\a")
			}
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
			if len(displayS) > 0 {
				displayS = displayS[:len(displayS)-1]
			}
			//fmt.Printf("%30s\n",displayS)
			g.Update(updateVD)
			return nil
		});
		err != nil {
		log_err(err)
	}
}