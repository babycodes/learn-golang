package helper

func BacklistWord(word string, keyword []string) bool {
	for _, result := range keyword {
		if result == word {
			return true
		}
	}
	return false
}
