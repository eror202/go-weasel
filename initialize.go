package weasel

import (
	"math/rand"
	"strings"
)

func Initialize(length int) string {
	var result strings.Builder
	for i := 0; i < length; i++ {
		randomCharacter := makeRandomCharacter()
		result.WriteRune(randomCharacter)
	}

	return result.String()
}

func makeRandomCharacter() rune {
	const alpabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ "
	randomIndex := rand.Intn(len(alpabet))
	return rune(alpabet[randomIndex])
}
