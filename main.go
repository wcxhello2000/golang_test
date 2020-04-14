package main

import (
	"fmt"
	"test/add"
	"test/div"
)

func main() {
	var a, b = 10, 20
	fmt.Printf("%d + %d = %d\n", a, b, add.Add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, div.DiV(a, b))
}
