package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Omikuji() int {
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(100)
	return res
}

func main() {
	res := Omikuji()
	switch {
	case res < 5:
		fmt.Println("大凶・・・だと！？")
	case res < 15:
		fmt.Println("まさか、凶を引くとは・・・。")
	case res < 45:
		fmt.Println("普通の吉かぁ。")
	case res < 75:
		fmt.Println("少しいい小吉か。")
	case res < 85:
		fmt.Println("お、中吉だ。中々だな。")
	case res > 95:
		fmt.Println("めっちゃ運いいじゃん！大吉や！！")
	}
}
