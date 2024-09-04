package stringmanipulate

// IsPalindrome by iterate through each string with 2 index 1 iterate
// start from first and last index then countiue increase frist index and decrease last index
// then compare is it equal charactor
func IsPalindrome(str string) bool {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
