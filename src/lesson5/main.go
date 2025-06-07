package main

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Hello, World!")
	if i := 0; i > 0 {
		fmt.Println("i is greater than 0")
	} else if i < 0 {
		fmt.Println("i is less than 0")
	} else {
		fmt.Println("i is 0")
	}

	s := "a"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	for i := range 10 {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			// スキップ
			continue
		}
		fmt.Println(i)
	}

	// 配列から要素を取り出す
	// 配列の長さを取得するにはlen関数を使う
	// 配列の要素を取得するにはrangeを使う
	arr := [3]int{123, 456, 789}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// スライスから要素を取り出す
	// スライスの長さを取得するにはlen関数を使う
	// スライスの要素を取得するにはrangeを使う
	//　配列とスライスの違いは長さを変更できるかどうか
	// 配列は長さを変更できない
	// スライスは長さを変更できる
	sl := []int{123, 456, 789}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	//　スライスは長さを変更できる
	sl1 := []int{123, 456, 789}
	sl2 := append(sl1, 100)
	fmt.Println(sl2)

	// マップから要素を取り出す
	// マップの長さを取得するにはlen関数を使う
	// マップの要素を取得するにはrangeを使う
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// interface型
	// 型アサーション
	var x interface{} = 3
	fmt.Println(x)

	// 型アサーション
	i, isInt := x.(int)
	fmt.Println(i, isInt)

	// 型アサーション
	// 変換に失敗するとpanicが発生する
	// falseになる
	f, isFloat := x.(float64)
	fmt.Println(f, isFloat)

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}

}
