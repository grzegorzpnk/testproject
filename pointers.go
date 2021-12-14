package main

import "fmt"

func pointers() {

	var x, y int
	var ptr *int
	x = 1
	ptr = &x
	y = *ptr
	fmt.Println("y: ", y)
	fmt.Println("ptr: ", ptr)
	fmt.Println("addr of ptr: ", &ptr)
	fmt.Println("value of ptr: ", *ptr)
	g := new(int)
	fmt.Println("g: ", g)
	fmt.Println("g val: ", *g)

}
