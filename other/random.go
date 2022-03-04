/*
generates a random number
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	// rand.Seed(86)
	// seeding with the same value returns same random sequence each run
	// seed with different value like time.Now().UnixNano()
	fmt.Println("my favourite number is:", rand.Intn(10))

}
