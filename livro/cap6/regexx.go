package regexx

import (
	"fmt"
	"regexp"
)

func regexx() {
	texto := "Anderson tem 21 anos"
	expr := regexp.MustCompile("\\d")
	fmt.Println(expr.ReplaceAllString(texto, "3"))
}
