package main

import "fmt"

func main() {

	variablesFunction()
	pointers()
	scope()
	compareTypes()
	stringTests()
}

//this is the code how to define vars in golang
func variablesFunction() {
	var1 := "test1"

	var var2 = "test2"

	var var3 string = "test3"

	var var4 string
	var4 = "test4"

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
}
