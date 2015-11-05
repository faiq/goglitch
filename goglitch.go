package main

import (
	"github.com/faiq/goglitch/glitcher"
	"os"
)

func main() {
	glitcher.DripImage(os.Args[1], os.Args[2])
}
