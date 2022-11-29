package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	print(programName)
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in GO!")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some progarm!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_MAIL: Logging in go!")
	fmt.Println("will you see this?")
}
