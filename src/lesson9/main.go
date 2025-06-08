package main

import "fmt"

func Double(i *int) {
	*i = *i * 2
}

func main() {
	i := 100

	fmt.Println(i)
	fmt.Println(&i)

	var p *int = &i
	fmt.Println(p)
	fmt.Println(*p)

	Double(&i)
	fmt.Println(i)

}
