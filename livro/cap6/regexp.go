package regexp

import (
	"fmt"
)

func regexp() {
	texto := "Anderson tem 21 anos"
	expr := regexp.MustCompile("\\d")
	fmt.Println(expr.ReplaceAllString(texto, "3"))
}
