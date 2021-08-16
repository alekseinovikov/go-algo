package lists

import "errors"

type singleLinkedListNode struct {
	head  bool
	value interface{}
	next  *singleLinkedListNode
}

func NewSingleLinkedList() List {
	return &singleLinkedListNode{
		head:  true,
		value: nil,
		next:  nil,
	}
}

func (s *singleLinkedListNode) Add(elem interface{}) {
	curr := s
	for curr.next != nil {
		curr = curr.next
	}
	addCreateAndAdd(curr, elem)
}

func addCreateAndAdd(node *singleLinkedListNode, elem interface{}) {
	newNode := new(singleLinkedListNode)
	newNode.head = false
	newNode.next = nil
	newNode.value = elem
	node.next = newNode
}

func (s *singleLinkedListNode) GetAt(index int) (elem interface{}, ok bool) {
	pointer := 0
	node := s.next
	for pointer < index && node != nil {
		node = node.next
		pointer++
	}

	if pointer != index || node == nil {
		return nil, false
	}

	return node.value, true
}

func (s *singleLinkedListNode) AddAt(index int, elem interface{}) error {
	if index == 0 {
		curr := s.next
		addCreateAndAdd(s, elem)
		s.next.next = curr

		return nil
	}

	pointer := 1
	curr := s.next
	prev := s
	for pointer < index && curr != nil {
		pointer++
		prev = curr
		curr = curr.next
	}

	if pointer != index {
		return errors.New("list is too short")
	}

	addCreateAndAdd(prev, elem)
	prev.next.next = curr
	return nil
}

func (s *singleLinkedListNode) RemoveLast() bool {
	curr, last := findLast(s)
	if last == nil {
		return false
	}

	curr.next = nil
	return true
}

func findLast(root *singleLinkedListNode) (prev *singleLinkedListNode, last *singleLinkedListNode) {
	if root.next == nil && root.head {
		return nil, nil
	}

	curr := root
	next := root.next
	for next.next != nil {
		curr = next
		next = next.next
	}

	return curr, next
}

func (s *singleLinkedListNode) RemoveFirst() bool {
	if s.next == nil {
		return false
	}

	s.next = s.next.next
	return true
}
