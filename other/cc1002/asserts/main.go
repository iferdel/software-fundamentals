package main

import (
	"fmt"
	"math/rand"
)

func randomNumber(min, max int) int {
	num := rand.Intn(max-min+1) + min
	fmt.Println("picked random number", num)
	return num
}
