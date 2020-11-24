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
	_,err2 := fmt.Fprintf(logFile, "[%s]   %s\n",
		time.Now().Format("2006.01.02 15:04:05"), err.Error())
	if err2 != nil {
		fmt.Println(err2)
	}
	// звуковой сигнал
	fmt.Printf("\a")
}

func main() {
	// логирование
	var err error
	logFile,err = os.OpenFile("errors.txt", os.O_APPEND, os.ModeAppend);
	if err != nil {
		logFile,_ = os.Create("errors.txt")
	}
	defer logFile.Close()

	go_ui()
}
