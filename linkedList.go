// Elia Cortesi 01911A
package main

// linkedList represents a linked list
//
// head The first element of the list
// tail the last element of the list
type linkedList struct {
	head *listNode
	tail *listNode
}

// listNode represents an element of the linked list
//
// next The next element of the list
// data The info inside the node
type listNode struct {
	next *listNode
	data rule
}

// newList creates a ne empty linked newList
//
// @return *linkedList The newly created newList
func newList() *linkedList {
	return &linkedList{nil, nil}
}

// newNode creates a new element of the list
//
// @param r The rule to be stored
// @return *listNode The newly created node
func newNode(r rule) *listNode {
	return &listNode{nil, r}
}

// addNode adds an element to the end of the list
//
// @param node The node to be added
func (l *linkedList) addNode(node *listNode) {
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
}
