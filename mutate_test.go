package weasel

import (
	"math/rand"
	"testing"
)

func TestMutate(test *testing.T) {
	rand.Seed(1)

	text := "the quick brown fox jumps over the lazy dog"
	actualResult := Mutate(text, 0.2)

	if actualResult != "the qu Pk brown fox jumps oveF tGD Nazy dog" {
		test.Fail()
	}
}
