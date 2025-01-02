package main

import "fmt"

func main() {
	a := 10
	fmt.Println("a value: ", a)
	fmt.Println("a address: ", &a)

	pointer := &a
	var pointer2 *int = pointer

	fmt.Println("pointer value: ", *pointer)
	fmt.Println("pointer address: ", pointer)
	fmt.Println("pointer2 value : ", *pointer2)
	fmt.Println("pointer2 address: ", pointer2)

	*pointer = 50

	fmt.Println("a value: ", a)
	fmt.Println("a address: ", &a)

	fmt.Println("pointer value: ", *pointer)
	fmt.Println("pointer address: ", pointer)
	fmt.Println("pointer2 value : ", *pointer2)
	fmt.Println("pointer2 address: ", pointer2)

}
