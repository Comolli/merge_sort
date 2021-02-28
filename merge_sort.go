package merge_sort

func merge(l []int, r []int) []int {
	result := []int{}

	for len(l) > 0 || len(r) > 0 {
		if len(l) > 0 && len(r) > 0 {
			if l[0] < r[0] {
				result = append(result, l[0])
				l = l[1:]
			} else {
				result = append(result, r[0])
				r = r[1:]
			}
		} else if len(l) > 0 {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}

	return result
}

// MergeSort algorithm applied to the given unsorted list
func MergeSort(list []int) []int {
	switch len(list) {
	case 0:
		return list
	case 1:
		return list
	default:
		m := len(list) / 2
		l := list[:m]
		r := list[m:]
		left := MergeSort(l)
		right := MergeSort(r)
		return merge(left, right)
	}
}
