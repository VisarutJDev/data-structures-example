package mysort

// MergeSort recursively sorts an array using the merge sort algorithm
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle point and divide the array into two halves
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

// merge merges two sorted arrays into one sorted array
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare each element and merge the two arrays
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
