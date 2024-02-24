package weasel

func Populate(text string, rate float64, count int) []string {
	var result []string
	for i := 0; i < count; i++ {
		mutatedText := Mutate(text, rate)
		result = append(result, mutatedText)
	}

	return result
}
