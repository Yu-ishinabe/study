package main

import (
	"fmt"

	"./animals"
)

func main() {
	// 構造確認
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
