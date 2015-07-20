package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pocke/slot"
)

func main() {
	var syms []string
	if len(os.Args) == 2 {
		str := os.Args[1]
		syms = strings.Split(str, "")
	} else if len(os.Args) > 2 {
		syms = os.Args[1:]
	} else {
		fmt.Fprintln(os.Stderr, "arguments required")
		os.Exit(1)
	}
	s := slot.New(syms, os.Stdout)
	i := s.Start()
	fmt.Println()
	fmt.Printf("%d回で %s が揃いました!\n", i, strings.Join(syms, ""))
}
