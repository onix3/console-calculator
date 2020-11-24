package main

import (
	"github.com/jroimartin/gocui"
	"regexp"
)

// Собственно реакция на нажатие
func replyKey(g *gocui.Gui, key rune) {
	W,_ := g.Size()
	highlight(key)
	// если ввод корректен и дисплей меньше ширины консоли
	if isValid(key) && sx+2+len(dE)+2 < W {
		dE += string(key)
		g.Update(updateVD)
	} else {
		highlight_error()
	}
}

// Проверка корректного ввода
func isValid(key rune) bool {
	s := dE + string(key) // строка на выходе

	// первый символ не должен быть /*+. (- можно)
	re1 := regexp.MustCompile(`^[/*+.]$`)

	// отклонить при "--"
	re2 := regexp.MustCompile(`--`)

	// отклонить при попытке после /*+ ввести /*+- (- можно)
	re3 := regexp.MustCompile(`[/*+-.][/*+.]`)

	// отклонить деление на 0 или -0 (именно на 0, на 0.2 можно)
	re4 := regexp.MustCompile(`/-?0[*/+-]`)

	// отклонить при попытке после "*0","/0","+0","-0", ввести цифру (после ".0" можно)
	re5 := regexp.MustCompile(`[*/+-]0\d`)

	// не допускать "0.0.", введение второй точки в дробное число
	re6 := regexp.MustCompile(`\.\d+\.`)

	// отклонить при попытке после "." ввести /*-+
	re7 := regexp.MustCompile(`\.\D`)

	// не допускать NaN
	re8 := regexp.MustCompile(`-?0/-?0[^.]`)

	if re1.MatchString(s) { return false }
	if re2.MatchString(s) { return false }
	if re3.MatchString(s) { return false }
	if re4.MatchString(s) { return false }
	if re5.MatchString(s) { return false }
	if re6.MatchString(s) { return false }
	if re7.MatchString(s) { return false }
	if re8.MatchString(s) { return false }

	return true
}