package stringmanipulate

// CountFrequency count charactor that duplicate in string
func CountFrequency(str string) map[rune]int {
	freq := make(map[rune]int)
	for _, ch := range str {
		freq[ch]++
	}
	return freq
}
