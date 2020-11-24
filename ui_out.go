package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strconv"
)

// Сделать цветную строку
func colorStr(s string, c1, c2 int) string {
	return fmt.Sprintf("\033[3%d;%dm%s\033[0m", c1, c2, s)
}

// Обновить инфо на дисплее
func updateVD(g *gocui.Gui) error {
	vd.Clear()

	outE := ""
	count := 0
	for _,c := range dE {
		if c == '/' || c == '*' || c == '+' || c == '-' {
			outE += colorStr(string(c),3,1)
			count++
		} else {
			outE += string(c)
		}
	}

	// длина цветной строки "1" составляет 12
	// поэтому с каждым символом к ширине для выравнивания добавляю 11
	fmt.Fprintf(vd, "%" + strconv.Itoa(dW+11*count) + "s\n", outE)
	dR := compute()       // результат вычисления
	dR = colorStr(dR,2,1) // стал зелёным
	fmt.Fprintf(vd, "%" + strconv.Itoa(dW+11) + "s", dR)

	return nil
}