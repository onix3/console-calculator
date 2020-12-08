package src

import (
	"github.com/jroimartin/gocui"
	"strings"
)

// Привязка нажатия кнопки мыши
func mouseBinding(key gocui.Key) {
	err := g.SetKeybinding("", key, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			if v.Name() != "display" && !strings.HasPrefix(v.Name(), "info") {
				replyKey(g,rune(v.Name()[0]))
			}
			if v.Name() == "ac" {
				dE = ""
				g.Update(updateVD)
			}
			return nil
		});
	IsErr(err)
}

// Привязка нажатий кнопок мыши
func initMouseBindings() {
	for _,key := range []gocui.Key{gocui.MouseLeft, gocui.MouseMiddle, gocui.MouseRight} {
		mouseBinding(key)
	}
}