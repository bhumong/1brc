package main

import (
	"fmt"

	"github.com/bhumong/1brc/filepath"
)

type City struct {
	Name string
	Temp float64
}

func main() {
	fmt.Println(filepath.GetPath())
}
