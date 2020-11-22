package main

import (
	"github.com/jroimartin/gocui"
)

func mouseBinding(key gocui.Key) {
	if err := g.SetKeybinding("", key, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			replyKey(g,rune(v.Name()[0]))
			return nil
		});
		err != nil {
		log_err(err)
	}
}

func initMouseBindings() {
	for _,key := range []gocui.Key{gocui.MouseLeft, gocui.MouseMiddle, gocui.MouseRight} {
		mouseBinding(key)
	}
}
