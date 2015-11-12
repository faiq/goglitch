package main

import (
	"fmt"
	"github.com/faiq/goglitch/glitcher"
	"os"
)

func main() {
	err := glitcher.DripImage(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("%v \n", err)
	}
}
