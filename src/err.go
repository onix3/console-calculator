package src

import (
	"fmt"
	"os"
	"time"
)

// Открытие лог-файла, запись, закрытие лог-файла
func LogErr(err error) {
	logFile, fileErr := os.OpenFile("console-calculator-ERRORS.txt", os.O_APPEND, os.ModeAppend)
	if fileErr != nil {
		logFile,_ = os.Create("console-calculator-ERRORS.txt")
	}
	defer logFile.Close()

	_,_ = fmt.Fprintf(logFile, "[%s]   %s\n",
		time.Now().Format("2006.01.02 15:04:05"), err.Error())
	fmt.Printf("\a") // звуковой сигнал
}

// Обработка ошибки, включающая стандартаное условие
func IsErr(err error) {
	if err != nil {
		LogErr(err)
	}
}