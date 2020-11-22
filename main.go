package main

import (
	"log"
	"os"
)

var (
	logFile   *os.File
)

// I like short code
func log_err(err error) {
	log.New(logFile, err.Error(), log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	//fmt.Printf("\a")
	// init logging
	var err error
	if logFile,err = os.Open("errors.txt"); err != nil {
		logFile,_ = os.Create("errors.txt")
	}
	defer logFile.Close()

	// go_ui() launches infinite cycle that's
	// it↓ should launch in goroutine before
	go func() {

	}()

	// launch console user interface
	go_ui()
}
