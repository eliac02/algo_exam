//Elia Cortesi 01911A

package main

// merge merges to sorted arrays in one single sorted array
// 
// @param left right Two arrays of rule
// @return []rule The sorted array of rules
func merge(left, right []rule) []rule {
	sorted := make([]rule, 0, len(left)+len(right))
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

// mergeSort sorts an array with O(n*(log(n))) time complexity and O(n) space complexity
//
// @param arr The pointer to the array containing the rules of the system
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
    for i:=0; i< len(*arr); i++ {
        (*arr)[i] = merged[i]
    }
}
