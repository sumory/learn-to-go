package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

var logFileName = flag.String("log", "logFile.log", "the log file")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if logErr != nil {
		fmt.Println("Fatal error")
		os.Exit(1)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("%v %v", "log something", "end")
}
