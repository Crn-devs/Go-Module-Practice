package main

import (
	"fmt"

	"github.com/Crn-devs/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(5)
	fmt.Println(s)
}
