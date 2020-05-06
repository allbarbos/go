package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	expr := regexp.MustCompile("\\b\\w")
	transformadora := func(s string) string {
		return strings.ToUpper(s)
	}
	texto := "antonio carlos jobim"

	result := expr.ReplaceAllStringFunc(texto, transformadora)

	fmt.Println(transformadora(texto))
	fmt.Println(result)
}
