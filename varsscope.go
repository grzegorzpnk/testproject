package main

import "fmt"

var ptr *int

func scope() {

	fmt.Println("Score function:")
	printVarX()
	printVarX()
	printVarX2()

}

func printVarX() {

	var x = 4
	ptr = &x
	fmt.Println(x)
	fmt.Println(" x addr", ptr)

}
func printVarX2() {

	fmt.Println(" x value", *ptr)
}
