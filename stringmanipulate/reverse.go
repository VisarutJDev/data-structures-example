package stringmanipulate

//ReverseString by iterate through each string with 2 index 1 iterate
// start from first and last index then countiue increase frist index and decrease last index
// then switch value in each loop
func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
