package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)

	ch1 <- 3
	fmt.Println(<-ch1)

	ch1 <- 100
	fmt.Println(<-ch1)

	ch1 <- 200
	fmt.Println(<-ch1)

	ch1 <- 300
	fmt.Println(<-ch1)

	ch1 <- 400
	fmt.Println(<-ch1)

	fmt.Println(ch1)

	ch1 <- 500
	i := <-ch1
	fmt.Println(i)

	ch1 <- 600
	fmt.Println(<-ch1)
}
