# My sort explaination

### Insertion Sort
Insertion Sort is a simple sorting algorithm that builds the final sorted array one item at a time. It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.

**Explanation:**
The algorithm iterates through the array, and for each element, it compares it with the elements before it and inserts it in the correct position relative to the already sorted part of the array.

**Time Complexity:**
Best/Average Case: 𝑂(𝑛) 
    Occurs when the array is already sorted.
    The algorithm only needs to iterate through the array once, making n−1 comparisons with no swaps.
Worst Case: O(n^2) —
    Occurs when the array is sorted in reverse order.
    The algorithm needs to shift every element for every insertion, resulting in n(n−1)/2 comparisons and shifts.

### Merge Sort
Merge Sort is an efficient, stable, comparison-based, divide-and-conquer sorting algorithm. Most implementations produce a stable sort, meaning that the order of equal elements is preserved.

**Explanation:**
Divide: The array is recursively divided into two halves until each subarray contains a single element.
Conquer: The sorted halves are merged back together to form the final sorted array.
Merge: The merging process ensures that the final array is sorted by comparing the elements of the two halves and merging them in the correct order.

**Time Complexity:**
Best/Average Case: 𝑂(𝑛log𝑛) 
    The best case is the same as the average and worst case since Merge Sort always divides the array into two halves and merges them, regardless of the initial order of elements.
Worst Case: 𝑂(𝑛log𝑛) 
    The worst-case time complexity also remains 𝑂(𝑛log𝑛) because the number of operations is determined by the recursive splitting and merging steps, which are consistent for any input.

### Quick sort
Quick Sort is a popular and efficient sorting algorithm that uses a divide-and-conquer approach to sort elements in an array or list. It’s well-known for its average-case time complexity of 𝑂(𝑛log𝑛) although in the worst case, it can degrade to O(n^2). However, with good pivot selection, the worst-case scenario is rare.

**Explanation:**
1. Base Case: If the array has fewer than 2 elements, it is already sorted, so the function returns immediately.
2. Partitioning:
    2.1 The pivot is selected as the last element of the array.
    2.2 The array is then partitioned so that all elements smaller than the pivot are on the left and all elements larger are on the right.
3. Recursive Calls: The QuickSort function is called recursively on the left and right subarrays until the entire array is sorted.

**Time Complexity:**
Best/Average Case: 𝑂(𝑛log𝑛) 
    Occurs when the pivot selection always divides the array into two nearly equal parts.
    This allows for an efficient division of the problem, reducing the depth of the recursion tree.
Worst Case: 𝑂(𝑛^2) 
    Occurs when the pivot selection consistently results in the most unbalanced partitions (e.g., choosing the smallest or largest element as the pivot in an already sorted array).
    The depth of the recursion tree becomes 𝑛 leading to n(n−1)/2 comparisons.

| Algorithm | Best Case | Worst Case | Stability | In-Place | Use Cases | 
| --- | --- | --- | --- |--- |--- |--- |--- |--- |--- |--- |--- | 
| Insertion Sort | O(n) | O(n^2) | Stable | Yes | Small datasets, nearly sorted data  | 
| Merge Sort | O(nlogn) | O(nlogn) | Stable | No | Large datasets, requires stable sorting  | 
| Quick Sort | O(nlogn) | O(n^2) | Unstable | Yes | General-purpose, often fastest in practice |  