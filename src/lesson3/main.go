package main

import "fmt"

// i5 := 100
// これはエラーになる
// 暗黙的定義は関数の中でしかできない

func main() {
	var i int = 100
	fmt.Println(i)

	var (
		i2 = 200
		i3 = "golang"
	)
	fmt.Println(i2, i3)

	// 暗黙的定義
	i4 := 400
	fmt.Println(i4)

	//  i4 := 500
	// これはエラーになる
	// 暗黙的定義は一度しかできない

}
