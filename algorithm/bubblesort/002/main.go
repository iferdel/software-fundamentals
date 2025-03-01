package main

import (
	"log"

	"github.com/iferdel/software-fundamentals/algorithm/bubblesort/002/bubblesort"
)

func main() {
	slice := []int{2, 5, 1, 7, 8, 10, 4}
	sortedSlice, err := bubblesort.BubbleSort(slice)
	if err != nil {
		log.Fatalln("bubblesort error:", err)
	}
	log.Println("slice sorted successfully:", sortedSlice)

}
