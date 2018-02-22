package main

import (
	"fmt"
)

func Help(args []string) {
	str := "Tools:\n"
	for k := range Dispatch {
		str += " * " + k + "\n"
	}
	fmt.Println(str)
}
