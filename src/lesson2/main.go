package main // goはパッケージ宣言から始まる mainからスタートする

import (
	"fmt"  // フォーマットパッケージ
	"time" // 時間パッケージ
)

// 定数で一度宣言した値をもう一度追加回せる
// ABCは1,1,1という値が入る
// DEは"A"という値が入る
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// iotaは0から始まる
// 連番を返す
const (
	i0 = iota
	i1
	i2
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("The time is", time.Now())

	fmt.Println(A, B, C, D, E, F)
	fmt.Println(i0, i1, i2)
}

// goだとbuildして出来上がったバイナリをそのまま配布するだけで動作する
// pyだと実行環境と実装環境で差分がある場合は動作しないk
