package weasel

import "testing"

func TestSearch(test *testing.T) {
	actualResult := Search([]string{"cat", "dog", "pig"}, "dig")
	if actualResult != "dog" {
		test.Fail()
	}
}
