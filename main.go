package main

import (
	"cgoTest/cfun"
	"fmt"
)

func main() {
	fmt.Println(cfun.Number_add_mod(10, 5, 12))
}
