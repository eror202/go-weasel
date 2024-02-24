package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eror202/go-weasel"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sample := "METHINKS IT IS LIKE A WEASEL"
	current := weasel.Initialize(len(sample))
	var generationCouner int
	for current != sample {
		const outputRate = 10
		if generationCouner%outputRate == 0 {
			fmt.Println(generationCouner, current)
		}

		variants := weasel.Populate(current, 0.05, 100)
		current = weasel.Search(variants, sample)

		generationCouner++
	}

	fmt.Println(generationCouner, current)
}
