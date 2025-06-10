package main

import (
	"github.com/ataberkcanitez/araqr/cmd"
)

func main() {
	if err := cmd.RunRootCmd(); err != nil {
		panic(err)
	}
}
