package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(100)
	fmt.Println("Random number: ", rand.Intn(786))
	fmt.Println("Random number: ", rand.Intn(786))

	rand.Seed(100)
	fmt.Println("Random number: ", rand.Intn(786))
	fmt.Println("Random number: ", rand.Intn(786))

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number: ", rand.Intn(786))
	fmt.Println("Random number: ", rand.Intn(786))

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number: ", rand.Intn(786))
	fmt.Println("Random number: ", rand.Intn(786))
}
