// Elia Cortesi 01911A

package main

// linkedCountingSort sorts the list of rules with O(n+k) time complexity, where n is the number of rules and k is the usage of the most-used rule
//
// @param The pointer to the array containing all the rules of the system
func linkedCountingSort(arr *[]rule) {
	if len(*arr) == 0 {
		return
	}

	// retrieve the most used rule
	maxUsage := (*arr)[0].usage
	for _, r := range *arr {
		if r.usage > maxUsage {
			maxUsage = r.usage
		}
	}

	// count the rules
	counts := make([]*linkedList, maxUsage+1)
	for index := range counts {
		counts[index] = newList()
	}

	for _, r := range *arr {
		node := newNode(r)
		counts[r.usage].addNode(node)
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
