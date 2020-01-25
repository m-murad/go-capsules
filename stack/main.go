package main

import (
	"errors"
	"fmt"
)

type stack []string

func (s *stack) Push(str string) {
	*s = append(*s, str)
}

func (s *stack) Pop() (str string, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err =  errors.New("stack under flow")
		}
	}()

	str = (*s)[len(*s)-1]
	(*s)[len(*s)-1] = ""
	*s = (*s)[:len(*s)-1]

	return str, nil
}

func main() {
	var myStk = new(stack) // Create a new stack instance. Stack becomes []

	myStk.Push("a") // Will add "a" to the stack. Stack becomes ["a"]
	myStk.Push("b") // Will add "b" to the stack. Stack becomes ["a", "b"]
	myStk.Push("c") // Will add "c" to the stack. Stack becomes ["a", "b", "c"]

	fmt.Printf("Stack is %v\n", myStk)


	str, err := myStk.Pop() // Will return string "c" and nil error. Stack becomes ["a", "b"]
	fmt.Printf("Poped item is %s\n", str)
	str, err = myStk.Pop() // Will return string "b" and nil error. Stack becomes ["a"]
	str, err = myStk.Pop() // Will return string "a" and nil error. Stack becomes []
	str, err = myStk.Pop() // Will return empty string "" and non-nil error.
	fmt.Printf("Failed to pop stack: %v\n", err)
}
