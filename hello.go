package main

import (
	"fmt"
	"github.com/user/stringutil"
	"github.com/user/fp"
)

func main() {
	fmt.Println(stringutil.Reverse("Hdddello, world."))

	var ff = fp.Compose(f, g)
	var res = ff("crap")
	fmt.Println(res)
}

func g(s string) string {
	fmt.Println("in g")
	return s + "fred"
}

func f(s string) string {
	fmt.Println("in f")
	return s + "simon"
}
