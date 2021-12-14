package main

import (
	"fmt"
)

func stringTests() {

	var tmp string = "grzegorz"
	var ptr *string

	ptr = &tmp
	fmt.Println("addr of tmp string: ", ptr)

	fmt.Println("val of tmp string: ", tmp)
	fmt.Println("addr of tmp string: ", ptr)

	//var t int = 5
	for t := 5; t > 0; t-- {

		fmt.Println("done", t)
		t--

	}

	//var m string
	//fmt.Scan(&m)
	//fmt.Println(m)
	////fmt.Println(num)
	//fmt.Println(err)

}
