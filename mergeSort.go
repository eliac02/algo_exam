package tiles

func merge(left, right []rule) []rule {
	result := make([]rule, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i].usage <= right[j].usage {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[i])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func mergeSort(arr []rule) []rule {
	if len(arr) <= 1 {
		return arr
	}
	half := len(arr) / 2
	left := mergeSort(arr[:half])
	right := mergeSort(arr[half:])

	return merge(left, right)
}
