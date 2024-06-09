//Elia Cortesi 01911A

package main

// insertionSort sorts an array in place with O(n) time complexity in the best case and O(n^2) time complexity in the worst case
//
// @param rules The pointer to the array containing all the rules of the system
func insertionSort(rules *[]rule) {
	arr := *rules
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j].usage > key.usage {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	*rules = arr
}
