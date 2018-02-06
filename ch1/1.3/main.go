package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()
	fmt.Println("+", end.Sub(start).Nanoseconds(), "ns")

	start = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	end = time.Now()
	fmt.Println("Join", end.Sub(start).Nanoseconds(), "ns")
}
