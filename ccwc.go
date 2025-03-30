package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	cFlag := flag.Bool("c", false, "Display the byte count of the file")
	lFlag := flag.Bool("l", false, "Display the number of lines in the file")
	wFlag := flag.Bool("w", false, "Display the number of words in the file")
	mFlag := flag.Bool("m", false, "Display the number of characters in the file")

	flag.Parse()
	fileName := flag.Arg(0)

	var fileContent []byte

	if len(fileName) > 0 {
		file, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		fileContent = file
	} else {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		fileContent = stdin
	}

	if !*cFlag && !*lFlag && !*wFlag && !*mFlag {
		fmt.Printf("%v %v %v %v\n", getLineCount(fileContent), getWordCount(fileContent), getFileSize(fileContent), fileName)
		return
	}

	var output string

	if *cFlag {
		byteSize := getFileSize(fileContent)
		output += fmt.Sprintf("%v ", byteSize)
	}
	if *lFlag {
		lines := getLineCount(fileContent)
		output += fmt.Sprintf("%v ", lines)
	}
	if *wFlag {
		words := getWordCount(fileContent)
		output += fmt.Sprintf("%v ", words)
	}
	if *mFlag {
		chars := getCharCount(fileContent)
		output += fmt.Sprintf("%v ", chars)
	}

	output += fileName

	fmt.Println(output)
}
