package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	num := rand.Intn(10)             // generate a random integer between 0 and 9
	fmt.Println(num)                 // print the random number
}
