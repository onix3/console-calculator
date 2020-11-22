package main

import "github.com/jroimartin/gocui"

func replyKey(g *gocui.Gui, key rune) {
	W,_ := g.Size()
	highlight(key)
	if 2+len(displayS)+2 < W {
		displayS += string(key)
		g.Update(updateVD)
	} else {
		highlight_error()
	}
}
