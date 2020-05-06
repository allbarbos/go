package closure

import "fmt"

func closure() {
	a, b := 0, 1
	fib := func() int {
		a, b = b, a+b
		return a
	}
	for i := 0; i < 8; i++ {
		fmt.Println("a:", a)
		fmt.Println("b:", b)
		fmt.Printf("%d ", fib())
		fmt.Println("\n")
	}
}
