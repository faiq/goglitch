package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/faiq/goglitch/glitcher"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "goglitch"
	app.Usage = "Create and back some awesome projects"
	err := glitcher.HorizontalSort(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("%v \n", err)
	}
}
