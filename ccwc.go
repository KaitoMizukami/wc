package main

import (
	"flag"
	"fmt"
	"os"
)

func run(args []string) {
	cFlag := flag.Bool("c", false, "Display the byte count of the file")

	flag.Parse()

	FileNameIdx := len(args) - 1
	fileName := args[FileNameIdx]
	if len(fileName) == 0 {
		fmt.Println("file name is necessary")
		os.Exit(1)
	}

	if *cFlag {
		fileInfo, err := os.Stat(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}

		size := fileInfo.Size()
		fmt.Println(size, fileName)
	}
}

func main() {
	run(os.Args[1:])
}
