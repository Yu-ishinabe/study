package main

import (
	"fmt"
)

func main() {
	a := [2][2]int{{1, 2}, {3, 4}}
	b := 5
	fmt.Println(a)
	if a[0][1] > b {
		fmt.Println("大きい", a[0][1])
	} else {
		fmt.Println("小さい", a[0][1])
	}
}
