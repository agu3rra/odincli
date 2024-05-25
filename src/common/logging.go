package common

import (
	"log"
	"os"
)

var (
	LogInfo = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogWarn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogErr  = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
