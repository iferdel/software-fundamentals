package main

import (
	"bytes"
	"log"
)

func main() {

	var (
		buf        bytes.Buffer
		loggerInfo = log.New(&buf, "INFO: ", log.Ldate|log.Llongfile|log.Lmsgprefix|log.LstdFlags|log.Ltime)
	)

	slice := [][]int{
		{1, 2, 3, 4, 5, 6},
		{4, 2, 3, 1, 0, 2},
		{9, 2, 4, 5, 6, 5},
	}
	target := 1

	_, err := binarySearch(slice, target)
	if err != nil {
		loggerInfo.Fatal("binarySearch error:", err)
	}

	// loggerInfo.Println("target found in idx:", s.idx...)

}
