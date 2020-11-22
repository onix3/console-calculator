package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strconv"
)


// Get colored string
func colorStr(s string, c1, c2 int) string {
	return fmt.Sprintf("\033[3%d;%dm%s\033[0m", c1, c2, s)
}

// update window with string
func updateVD(g *gocui.Gui) error {
	vd.Clear()

	fmt.Fprintf(vd, "%" + strconv.Itoa(displayW) + "s", displayS)

	//fmt.Fprintln(vd,"╔" + strings.Repeat("═",displayW-2) + "╗")
	//fmt.Fprintf(vd, "║%" + strconv.Itoa(displayW) + "s║\n", displayS)
	//fmt.Fprintln(vd,"╚" + strings.Repeat("═",displayW-2) + "╝")

	return nil
}