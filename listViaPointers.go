package main

/*
************

Двусвязный список: предыдущий элемент, следующий элемент, значение

************
*/

import "fmt"


type Element struct {
	prev *Element
	next *Element
	value string  // may be anything
}

func (root *Element) createFromString(s string) {
	cur := root

	for _, v := range s {		
		cur.value = string(v)
		var next Element
		cur.next = &next
		next.prev = cur
		cur = &next
	}
	cur = cur.prev
	cur.next = nil
}

func (root *Element) printElements() {
	pointer := root

	for pointer != nil {
		fmt.Printf("| %v |\t", pointer.value)
		pointer = pointer.next
	}
	fmt.Print("\n")
}

func (root *Element) printReverseElements() {
	pointer := root.getLastElemnt()

	for pointer != nil {
		fmt.Printf("| %v |\t", pointer.value)
		pointer = pointer.prev
	}
	fmt.Print("\n")
}

func (root *Element) getLastElemnt() *Element {
	pointer := root

	for pointer.next != nil {
		pointer = pointer.next
	}

	return pointer
}

func (root *Element) lenElements() int {
	var length int
	pointer := root

	for pointer != nil {
		length ++
		pointer = pointer.next
	}

	return length
}

func main() {
	fmt.Println("Enter a string")
	var s string
	fmt.Scan(&s)

	var root Element
	root.createFromString(s)

	root.printElements()
	fmt.Print("\n\n")
	root.printReverseElements()
	fmt.Print("\n\n")
	root.printElements()
}
