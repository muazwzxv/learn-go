package main

import (
	"errors"
	"fmt"
)

type node struct {
	data *int
	next *node
}

type linkedlist struct {
	head *node
}

func newNode(data int) *node {
	return &node{&data, nil}
}

func (list *linkedlist) addAtFront(data int) {
	node := newNode(data)
	node.next = list.head
	list.head = node
}

func (list *linkedlist) addAtBack(data int) {
	node := newNode(data)

	if list.head == nil {
		list.head = node
	}
	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = node
}

func (list *linkedlist) deleteAtFront() (*int, error) {
	if list.head == nil {
		code := -1
		return &code, errors.New("Houstan we got a problem: the list is empty")
	}

	current := list.head
	list.head = current.next

	return current.data, nil
}

func (list *linkedlist) deleteAtBack() (*int, error) {
	if list.head == nil {
		code := -1
		return &code, errors.New("Houstan we got a problem: the list is empty")
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	toReturn := current.data
	current.next = nil
	return toReturn, nil
}

func (list *linkedlist) count() (*int, error) {
	if list.head == nil {
		code := -1
		return &code, errors.New("Houstan we got a problem: the list is empty")
	}

	count := 0
	for current := list.head; current.next != nil; current = current.next {
		count++
	}

	return &count, nil
}

func (list *linkedlist) displayList() error {
	current := list.head

	if current.data == nil {
		return errors.New("Houstan we got a problem: the list is empty")
	}

	for current.next != nil {
		fmt.Println(current.data)
		current = current.next
	}

	return nil
}

func main() {
	fmt.Println("LinkedList")
}
