package main

import (
	"fmt"
)

func compareTypes() {
	var x int32 = 1
	var y int64 = 2

	y = int64(x)
	fmt.Println(y)

}
