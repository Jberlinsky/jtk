package main

import (
	"os"
	"path"
)

var Dispatch map[string]func([]string)

func main() {
	binary := path.Base(os.Args[0])
	args := os.Args

	Dispatch = map[string]func([]string){
		"ec2-resource-for-ip": Ec2ResourceForIp,

		"explode": Explode,
		"help":    Help,
	}

	if binary == "leatherman" && len(args) > 1 {
		args = args[1:]
		binary = args[0]
	}

	fn, ok := Dispatch[binary]
	if !ok {
		Help(os.Args)
		os.Exit(1)
	}
	fn(args)
}
