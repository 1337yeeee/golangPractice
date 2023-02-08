package main

/*
************

Двусвязный список: предыдущий элемент, следующий элемент, значение

************
*/

import (
	"fmt"
	"os"
	"os/exec"
)


type Element struct {
	prev *Element
	next *Element
	value string  // may be anything
}

func clear() { 
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
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
	length := root.lenElements()

	for i:=0; i<length; i++ {
		fmt.Printf("+---+\t")
	}
	fmt.Print("\n")

	for pointer != nil {
		fmt.Printf("| %v |\t", pointer.value)
		pointer = pointer.next
	}
	fmt.Print("\n")

	for i:=0; i<length; i++ {
		fmt.Printf("+---+\t")
	}
	fmt.Print("\n")
}

func (root *Element) printElementsReverse() {
	pointer := root.getLastElemnt()
	length := root.lenElements()

	for i:=0; i<length; i++ {
		fmt.Printf("+---+\t")
	}
	fmt.Print("\n")

	for pointer != nil {
		fmt.Printf("| %v |\t", pointer.value)
		pointer = pointer.prev
	}
	fmt.Print("\n")

	for i:=0; i<length; i++ {
		fmt.Printf("+---+\t")
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

	if pointer.next == nil {
		return 0
	}

	for pointer != nil {
		length ++
		pointer = pointer.next
	}

	return length
}

func menuFuncOption1(root *Element) {
	clear()
	fmt.Print("Enter a string: ")
	var s string
	fmt.Scan(&s)
	root.createFromString(s)
}

func menuFuncOption2(root *Element) {
	clear()
	fmt.Print("You've chosen `Print list reverse`\n\n")
	root.printElementsReverse()
	fmt.Print("\nPress Enter to continue\n\n")
	fmt.Scanln()
}

func menuFuncOption3(root *Element) {
	newRoot := Element{}
	*root = newRoot
}

func main() {
	var root Element
	// cursor := &root

	run:=true
	for run {
		clear()
		root.printElements()
		fmt.Print(
"Shose what to do:\n",
"1. Create Elements from string\n",
"2. Print list reverse\n",
"3. Delete list\n",
"0. Quite.\n")
		var choice rune
		fmt.Scanf("%c", &choice)
		fmt.Scanln()


		switch choice {
			case '1':
				menuFuncOption1(&root)
			case '2':
				menuFuncOption2(&root)
			case '3':
				menuFuncOption3(&root)
			case '0':
				run = false 
		}
	}
}

// TODO укзатель на рабочий элемент
