package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)


// Get colored string
func colorStr(s string, c1, c2 int) string {
	return fmt.Sprintf("\033[3%d;%dm%s\033[0m", c1, c2, s)
}

// update window with string
func updateVS(g *gocui.Gui) error {
	vd.Clear()

	return nil
}