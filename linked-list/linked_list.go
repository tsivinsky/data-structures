package main

import (
	"errors"
	"strconv"
	"strings"
)

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) String() string {
	var s []string

	current := list.Head
	for current != nil {
		s = append(s, strconv.Itoa(current.Data))
		current = current.Next
	}

	return strings.Join(s, ", ")
}

func (list *LinkedList) Len() int {
	l := 0

	current := list.Head
	for current != nil {
		l++
		current = current.Next
	}

	return l
}

func (list *LinkedList) Insert(data int) {
	node := Node{
		Data: data,
	}

	node.Next = list.Head
	if list.Head != nil {
		list.Head.Prev = &node
	}

	list.Head = &node

	node.Prev = nil
}

func (list *LinkedList) Find(data int) *Node {
	current := list.Head
	for current != nil {
		if current.Data == data {
			return current
		}

		current = current.Next
	}

	return nil
}

func (list *LinkedList) Delete(node *Node) error {
	if node == nil {
		return errors.New("No node")
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		list.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	return nil
}
