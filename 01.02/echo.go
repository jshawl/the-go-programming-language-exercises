package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + fmt.Sprintf("%d", i) + ": " + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}
