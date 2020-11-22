package main

import (
	"github.com/jroimartin/gocui"
	"regexp"
)

func replyKey(g *gocui.Gui, key rune) {
	W,_ := g.Size()
	highlight(key)
	if 2+len(dS)+2 < W && isValid(key) {
		dS += string(key)
		g.Update(updateVD)
	} else {
		highlight_error()
	}
}

// проверка корректного ввода
func isValid(key rune) bool {
	s := dS + string(key) // строка на выходе

	// первый символ не должен быть /*-+
	re0 := regexp.MustCompile(`^\D$`)
	if re0.MatchString(s) {
		return false
	}

	// отклонить при попытке после /*-+ ввести /*-+
	re1 := regexp.MustCompile(`\D\D`)
	if re1.MatchString(s) {
		return false
	}

	// отклонить при попытке после "/0","*0","-0","+0" ввести цифру
	re2 := regexp.MustCompile(`\D0\d`)
	if re2.MatchString(s) {
		return false
	}

	return true
}