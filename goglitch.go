package main

import (
	"github.com/faiq/goglitch/glitcher"
	"os"
)

func main() {
	glitcher.InvertImage(os.Args[1], os.Args[2])
}
