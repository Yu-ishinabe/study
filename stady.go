package main

import (
	"fmt"
)

const count int = 10

var (
	c         int
	e         string
	plusAlias = plus
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

	// for確認
	for i := 0; i < count; i++ {
		// count = 5 これを実施した場合,書き換えられない
		// cannot assign to count というエラーになる
		if i == a[1][0] {
			break
		}
		fmt.Println(i, a[1][0])
	}

	// 関数の確認
	c := plus(a[1][0], b)
	// 変数への暗示的な代入は1回のみ
	// c := 1 をするとコンパイルエラーとなる
	// 「演算子 =」の場合は再代入は制限なし
	d := c + b
	e := "dog"
	// Printfは書式指定式を含んだフォーマット文字列と、続く可変長を記載して出力する
	// 送るパラメータが足らない場合、MISSINGとなり、多い場合はEXTRAとなる
	fmt.Printf("合計は%d この動物は%v\n", d, e)

	// 複数の戻り値
	f, g := division(d, a[1][1]) // 戻り値の破棄をする場合は「_」を使う
	fmt.Printf("商=%d 余=%d\n", f, g)

	h := doSomething()
	fmt.Println(h)

	// 無名関数
	j := func(x, y int) int { return x + y } // 無名関数を定義
	k := j(d, h)                             // 変数jにd,hの引数を渡して関数を実行
	fmt.Println(k)

	//無名関数と名前関数
	l := plusAlias(k, d) // plusAlias(12行目)はplus関数を直接代入している
	fmt.Println(l)

	// 関数を返す関数
	returnAsk()() // 戻り値を変数に入れなくても直接呼び出せる
}

func plus(a, b int) int {
	return a + b
}

func division(a, b int) (int, int) {
	f := a / b
	g := a % b
	return f, g
}
func doSomething() (a int) {
	return
	// エラーになりそうだが、実際は下記のようになっている
	/*
		func doSomething() int {
			var a int
			return a
		}
	*/
	// var a int によって変数aが整数0に初期化されているため
}

func returnAsk() func() {
	return func() {
		fmt.Println("Are you my master?")
	}
}
