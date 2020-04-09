package main

import (
	"fmt"
)

func main() {
	// 配列確認
	a := [2][2]int{{1, 2}, {3, 4}}
	b := 5
	fmt.Println(a)
	// if確認
	if a[0][1] > b {
		fmt.Println("大きい", a[0][1])
	} else {
		fmt.Println("小さい", a[0][1])
	}
}
