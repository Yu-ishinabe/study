package main

import (
	"fmt"
)

const count int = 10

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

	// for確認
	for i := 0; i < count; i++ {
		// count = 5 これを実施した場合,書き換えられない
		// cannot assign to count というエラーになる
		if i == a[1][0] {
			break
		}
		fmt.Println(i, a[1][0])
	}
}
