package main

import "fmt"

// 関数の宣言
// 関数名(引数) 返り値
// 引数で同じ型の場合は一度だけ宣言すれば良い
func Plus(a, b int) int {
	return a + b
}

func Div(a, b int) (string, int) {
	return "result", a / b
}

func Later() func(string) string {
	i := ""
	return func(x string) string {
		i = x
		return i
	}
}
func main() {
	fmt.Println("Hello, World!")
	i := Plus(1, 2)
	fmt.Println(i)

	// 返り値が2つの場合は変数を宣言して受け取る
	// 2つ目の返り値は_で受け取ると捨てることができる
	s, _ := Div(100, 2)
	fmt.Println(s)

	f := func(x, y int) int {
		return x + y
	}
	i2 := f(1, 2)
	fmt.Println(i2)

	// 無名関数を即時実行
	//　クロージャー
	i3 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i3)

}
