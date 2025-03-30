package main

import "bytes"

func getFileSize(content []byte) int {
	return len(content)
}

func getLineCount(content []byte) int {
	return bytes.Count(content, []byte("\n"))
}

func getWordCount(content []byte) int {
	return len(bytes.Fields(content))
}

func getCharCount(content []byte) int {
	return len(bytes.Runes(content))
}
