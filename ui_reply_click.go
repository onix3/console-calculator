package main

import (
	"github.com/jroimartin/gocui"
	"regexp"
)

func replyKey(g *gocui.Gui, key rune) {
	W,_ := g.Size()
	highlight(key)
	if 2+len(dE)+2 < W && isValid(key) {
		dE += string(key)
		g.Update(updateVD)
	} else {
		highlight_error()
	}
}

// проверка корректного ввода
func isValid(key rune) bool {
	s := dE + string(key) // строка на выходе

	// первый символ не должен быть /*+. (- можно)
	re1 := regexp.MustCompile(`^[/*+.]$`)

	// отклонить при "--"
	re2 := regexp.MustCompile(`--`)

	// отклонить при попытке после /*+ ввести /*+- (- можно)
	re3 := regexp.MustCompile(`[/*+-.][/*+.]`)

	// отклонить деление на 0 (именно на 0, на 0.2 можно)
	re4 := regexp.MustCompile(`/0[*/+-]`)

	// отклонить при попытке после "*0","/0","+0","-0", ввести цифру (после ".0" можно)
	re5 := regexp.MustCompile(`[*/+-]0\d`)

	// не допускать "0.0.", введение второй точки в дробное число
	re6 := regexp.MustCompile(`\.\d+\.`)

	if re1.MatchString(s) { return false }
	if re2.MatchString(s) { return false }
	if re3.MatchString(s) { return false }
	if re4.MatchString(s) { return false }
	if re5.MatchString(s) { return false }
	if re6.MatchString(s) { return false }

	return true
}