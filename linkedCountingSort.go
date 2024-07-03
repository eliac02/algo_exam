// Elia Cortesi 01911A

package main

func linkedCountingSort(arr *[]rule) {
    if len(*arr) == 0 {
        return
    }

    maxUsage := (*arr)[0].usage
    for _, r := range (*arr) {
        if r.usage > maxUsage {
            maxUsage = r.usage
        }
    }

    counts := make([]*linkedList, maxUsage+1)
    for index := range counts {
        counts[index] = newList()
    }

    for _, r := range (*arr) {
        node := newNode(r)
        counts[r.usage].addNode(node)
    }

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
