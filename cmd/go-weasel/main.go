package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/eror202/go-weasel"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sample := flag.String("sample", "METHINKS IT IS LIKE A WEASEL", "target string")
	rate := flag.Float64("rate", 0.05, "character mutate rate")
	populationCount := flag.Int("count", 100, "population size")
	flag.Parse()

	current := weasel.Initialize(len(*sample))
	var generationCouner int
	for current != *sample {
		const outputRate = 10
		if generationCouner%outputRate == 0 {
			fmt.Println(generationCouner, current)
		}

		variants := weasel.Populate(current, *rate, *populationCount)
		current = weasel.Search(variants, *sample)

		generationCouner++
	}

	fmt.Println(generationCouner, current)
}
