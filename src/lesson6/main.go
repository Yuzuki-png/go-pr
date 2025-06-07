package main

import "fmt"

func main() {
	sl := []int{100, 200}
	sl = append(sl, 300)
	fmt.Println(sl)

	// make関数
	// スライスを作成する
	// スライスの長さを指定する
	// スライスの容量を指定する
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)

	// マップの要素を取得する
	fmt.Println(m["apple"])

	s, ok := m["apple"]
	if !ok {
		fmt.Println("apple is not exist")
	} else {
		fmt.Println(s)
	}

}
