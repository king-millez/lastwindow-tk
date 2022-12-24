package main

import (
	"errors"
	"fmt"
	"lastwindow-tk/src/lastwindow-tk/file"
	"os"
)

func main() {
	cmdArgs := os.Args[1:]
	if len(cmdArgs) > 2 {
		if _, err := os.Stat(cmdArgs[1]); errors.Is(err, os.ErrNotExist) {
			panic(fmt.Sprintf("%v does not exist.\n", cmdArgs[1]))
		}
		if _, err := os.Stat(cmdArgs[2]); os.IsNotExist(err) {
			os.MkdirAll(cmdArgs[2], os.ModePerm) // Create nested directories to the output dir
		}
		switch cmdArgs[0] {
		case "packfile":
			file.UnpackPackfile(cmdArgs[1], cmdArgs[2])
		}
	} else {
		panic("Not enough arguments were provided.\nUsage: lastwindow-tk <mode> [file] [output directory]")
	}
}
