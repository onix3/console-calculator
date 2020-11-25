package main

import (
	"fmt"
	"os"
	"time"
)

var (
	logFile   *os.File
)

// Обработка ошибки
func log_err(err error) {
	_,_ = fmt.Fprintf(logFile, "[%s]   %s\n",
		time.Now().Format("2006.01.02 15:04:05"), err.Error())
	// звуковой сигнал
	fmt.Printf("\a")
}

func main() {
	// логирование
	var err error
	logFile,err = os.OpenFile("console-calculator-errors.txt", os.O_APPEND, os.ModeAppend);
	if err != nil {
		logFile,_ = os.Create("console-calculator-errors.txt")
	}
	defer logFile.Close()

	for i:=0; i<20; i++ {
		fmt.Println()
	}

	go_ui()
}
