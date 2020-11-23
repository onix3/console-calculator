package main

import (
	"fmt"
	"os"
	"time"
)

var (
	logFile   *os.File
)

// I like short code
func log_err(err error) {
	_,err2 := fmt.Fprintf(logFile, "[%s]   %s\n",
		time.Now().Format("2006.01.02 15:04:05"), err.Error())
	if err2 != nil {
		fmt.Println(err2)
	}
	//log.New(logFile, err.Error(), log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// init logging
	var err error
	logFile,err = os.OpenFile("errors.txt", os.O_APPEND, os.ModeAppend);
	if err != nil {
		logFile,_ = os.Create("errors.txt")
	}
	defer logFile.Close()

	for i:=0 ; i<17; i++ {
		fmt.Println()
	}

	// go_ui() launches infinite cycle that's
	// it↓ should launch in goroutine before
	go func() {

	}()

	// launch console user interface
	go_ui()
}
