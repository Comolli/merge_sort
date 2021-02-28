package merge_sort

func MergeSortInts(arr []int) int {
	return _mergeSortInts(arr)
}

func _mergeSortInts(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	mid := (len(arr) / 2)
	//Gather inversions from the two sub arrays
	inv_low := _mergeSortInts(arr[:mid])
	inv_high := _mergeSortInts(arr[mid:])
	//Gather the split inversions
	split_inv := _merge(arr, mid)
	return inv_low + inv_high + split_inv
}

func _merge(arr []int, mid int) int {
	arr_lo := make([]int, len(arr[:mid]))
	arr_hi := make([]int, len(arr[mid:]))
	copy(arr_lo, arr[:mid])
	copy(arr_hi, arr[mid:])
	var i, j, k, split_inv int

	//Copy till either of the sub arrays gets completely copied
	for i < len(arr_lo) && j < len(arr_hi) {
		if arr_lo[i] <= arr_hi[j] {
			arr[k] = arr_lo[i]
			i++
		} else {
			arr[k] = arr_hi[j]
			j++
			split_inv += mid - i
		}
		k++
	}
	//Copy the remaining elements of the sub-array if any
	for i < len(arr_lo) {
		arr[k] = arr_lo[i]
		i++
		k++
	}
	//Copy the remaining elements of the sub-array if any
	for j < len(arr_hi) {
		arr[k] = arr_hi[j]
		j++
		k++
	}
	return split_inv
}
