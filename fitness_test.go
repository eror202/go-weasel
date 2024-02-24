package weasel

import "testing"

func TestFitness_withDifference(test *testing.T) {
	text := "the quick brown fox jumps over the lazy dog"
	sample := "the quick brown cat jumps over the lazy pig"

	result := Fitness(text, sample)

	if result != 5 {
		test.Fail()
	}
}

func TestFitness_withoutDifference(test *testing.T) {
	text := "the quick brown fox jumps over the lazy dog"

	result := Fitness(text, text)

	if result != 0 {
		test.Fail()
	}
}
