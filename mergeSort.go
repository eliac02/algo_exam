package main

func merge(left, right []rule) []rule {
	var sorted []rule
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i].usage <= right[j].usage {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[i])
			j++
		}
	}

	sorted = append(sorted, left[i:]...)
	sorted = append(sorted, right[j:]...)

	return sorted
}

func mergeSort(arr *[]rule) {
	if len(*arr) <= 1 {
		return 
	}
	half := len(*arr) / 2
    left := make([]rule, half)
    right := make([]rule, len(*arr)-half)

    copy(left, (*arr)[:half])
    copy(right, (*arr)[half:])

    mergeSort(&left)
    mergeSort(&right)

    merged := merge(left, right)
	copy(*arr, merged)
}
