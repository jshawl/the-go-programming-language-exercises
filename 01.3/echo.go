package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start)
	fmt.Println("\nfor loop took", elapsed)
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("\n.Join took", time.Since(start))
}
