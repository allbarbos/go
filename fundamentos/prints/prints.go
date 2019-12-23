package main

import "fmt"

func main() {
	fmt.Print("linha")
	fmt.Print(" um.")

	x := 3.141516

	fmt.Println()
	fmt.Println("valor de x: " + fmt.Sprint(x))
	fmt.Println("valor de x:", x)

	fmt.Printf("valor de x: %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "bile"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
}
