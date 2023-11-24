package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Print("%v\n", args)
	fmt.Print("%v\n", args[1])

}
