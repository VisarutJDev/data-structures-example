package mysort

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Swap
	return i + 1
}

// QuickSort is a sorting algorithm based on the Divide and Conquer algorithm
// that picks an element as a pivot and partitions the given array around the picked pivot
// by placing the pivot in its correct position in the sorted array.
// Step
//  1. Start with set pivot, Normally it will be last element
//  2. find partition by first index (i) + 1 after sort
//     2.1 iterate to array by low -1 (which is i that means start from -1)
//     2.2 if element in position j is less than pivot, increase i by 1 and swap element i and j (which can be same place i==j)
//     2.3 lastly you will get new pivot position which is i+1 swap with last element
//  3. send new pivot position to other recursive quick sort
func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := Partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}
