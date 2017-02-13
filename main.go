package main

import (
	"fmt"

	"github.com/marcusljx/eevee/entities"
)

func main() {
	var a entities.SimpleTextEntity = "HELLO WORLD"
	x := fmt.Sprintf("%s", a)
	fmt.Println(x)
}
