package arraystring

//LinearSearch is a search that search from starting point to the end
func LinearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1 // Element not found
}
