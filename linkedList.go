// Elia Cortesi 01911A

package main

type linkedList struct {
	head *listNode
	tail *listNode
}

type listNode struct {
	next *listNode
	data rule
}

func newList() *linkedList {
	return &linkedList{nil, nil}
}

func newNode(r rule) *listNode {
	return &listNode{nil, r}
}

func (l *linkedList) addNode(node *listNode) {
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
}
