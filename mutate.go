package weasel

import (
	"math/rand"
	"strings"
)

func Mutate(text string, rate float64) string {
	var result strings.Builder
	for _, character := range text {
		if rand.Float64() < rate {
			randomCharacter := makeRandomCharacter()
			result.WriteRune(randomCharacter)
		} else {
			result.WriteRune(character)
		}
	}

	return result.String()
}
