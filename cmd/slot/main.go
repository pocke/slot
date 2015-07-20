package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pocke/slot"
)

func main() {
	str := os.Args[1]
	syms := strings.Split(str, "")
	s := slot.New(syms, os.Stdout)
	i := s.Start()
	fmt.Println()
	fmt.Printf("%d回で %s が揃いました!\n", i, str)
}
