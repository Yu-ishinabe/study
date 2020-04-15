package main

import (
	"fmt"
	"log"
	"strconv"
)

const count int = 10

var (
	c         int
	e         string
	plusAlias = plus
)

// type で Servant型を定義 何回も同じ内容を利用するなら type記載
type Servant = struct {
	name  string
	class string
}

func main() {
	// 配列確認
	a := [2][2]int{{1, 2}, {3, 4}} // 文字列の場合 [2][2]string{"a", "b"}, {"c", "d"}
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
	returnAsk()() // 戻り値を変数に入れなくても直接呼び出せる。最後の()で出力

	//  型変換とエラー処理
	str := "10"                 // "a"にするとintにはならないのでエラーになる
	i, err := strconv.Atoi(str) //stringからintへ (strconvが変換するパッケージ)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i) // float <=> int、string <=> []byte は変換可能

	// スライス 配列の一部を切り出す構造体
	m := [...]int{1, 2, 3, 4, 5}
	m2 := m[1:4] //mの配列内の1~3を参照している (最後の4は4番目までではなく「その前まで」になる)
	fmt.Println(m)
	fmt.Println(m2)

	m2[0] = 10     // スライスは元の配列を参照しているだけで、データを持っていない。
	fmt.Println(m) // m2[0]の処理のスライスを通して、元の配列も変わっている
	fmt.Println(m2)

	// マップ キーと値をマッピングさせている構造体
	n := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fmt.Println(n)
	delete(n, "b") // キーn bの要素の削除
	fmt.Println(n)
	res, has := n["c"] // 要素にアクセスした場合２つ戻り値を受けてれて、1つ目は値が入り、2つ目は要素が存在しているかのbool型による判定
	if has {
		fmt.Printf("%#+v\n", res)
	} else {
		fmt.Println("has not")
	}

	// 構造体
	servant := Servant{
		// Servant型を フィールド名:値 で初期化をする
		name:  "Altria",
		class: "Saber",
	}
	/* 一時的な場合はtypeではなく、直接以下のようにも書ける
	var servant struct {
		name  string
		class string
	}{
		name:  "Altria",
		class: "Saber",
	}
	*/

	// 「.」でフィールドにアクセス可能
	fmt.Printf("name: %s, class: %s\n", servant.name, servant.class)

	// 値の代入も「.」で可能
	servant.name = "Emiya"
	servant.class = "Archer"
	fmt.Printf("name: %s, class: %s\n", servant.name, servant.class)
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
