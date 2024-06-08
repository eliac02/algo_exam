package main

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
