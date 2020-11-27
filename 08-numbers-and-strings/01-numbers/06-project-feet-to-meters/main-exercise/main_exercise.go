package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	c, _ := strconv.ParseFloat(os.Args[1], 64)
	f := (c * 1.8) + 32
	fmt.Printf("%g C is %g F.", c, f)
}
