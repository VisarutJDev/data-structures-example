package arraystring

// BinarySearch is a search that that finds the position of a target value within a sorted array
// start from mid value and check that value is less or more than target value or equal
// if less, change high position to -1 of mid
// if more, change low position to +1 of mid
// this will search by continuing to cut array in half to search
func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Element not found
}
