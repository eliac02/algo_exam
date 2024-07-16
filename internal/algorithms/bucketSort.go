package algorithms

import (
	models "tiles/internal/models"
)

// bucketSort sorts the list of rules with O(n+k) time complexity, where n is the number of rules and k is the usage of the most-used rule
//
// @param The pointer to the array containing all the rules of the system
func bucketSort(arr *[]models.Rule) {
	if len(*arr) == 0 {
		return
	}

	// retrieve the most used rule
	maxUsage := (*arr)[0].Usage
	for _, r := range *arr {
		if r.Usage > maxUsage {
			maxUsage = r.Usage
		}
	}

	// count the rules
	counts := make([]*linkedList, maxUsage+1)
	for index := range counts {
		counts[index] = newList()
	}

	for _, r := range *arr {
		node := newNode(r)
		counts[r.Usage].addNode(node)
	}

	// sort the rules
	index := 0
	for _, list := range counts {
		current := list.head
		for current != nil {
			(*arr)[index] = current.data
			index++
			current = current.next
		}
	}
}
