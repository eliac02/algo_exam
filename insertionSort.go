package main

func insertionSort(rules *[]rule) {
	arr := *rules
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Sposta gli elementi dell'array arr[0..i-1], che sono maggiori del key,
		// di una posizione avanti rispetto alla loro posizione attuale
		for j >= 0 && arr[j].usage > key.usage {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	*rules = arr
}
