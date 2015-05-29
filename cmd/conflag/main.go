package main

import (
	"fmt"

	"github.com/monochromegane/conflag"
)

func main() {
	args, err := conflag.ArgsFrom("test.toml", "args")
	fmt.Printf("%v, %v\n", args, err)
}