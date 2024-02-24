package weasel

func Fitness(text string, sample string) int {
	var result int
	for index, character := range text {
		if character != rune(sample[index]) {
			result++
		}
	}

	return result
}
