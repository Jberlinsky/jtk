package main

import (
	"fmt"
	"os"
	"path"
)

func Explode(args []string) {
	exe, err := os.Executable()
	if err != nil {
		fmt.Println("Couldn't get Executable instnace:", err)
		os.Exit(1)
	}

	dir := path.Dir(exe)
	for k := range Dispatch {
		if k == "help" {
			continue
		}
		if k == "explode" {
			continue
		}

		err := os.Symlink(exe, dir+"/"+k)
		if err != nil {
			panic(err)
		}
	}
}
