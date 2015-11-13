package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/faiq/goglitch/glitcher"
)

func main() {
	app := cli.NewApp()
	app.Name = "goglitch"
	app.Usage = "A CLI to glitch your images. Most commands are "
	app.Commands = []cli.Command{
		{
			Name:    "horrizontal-sort",
			Aliases: []string{"h-s"},
			Usage:   "Horrizontal sort the values of all the pixels in every row.",
			Action: func(c *cli.Context) {
				args := c.Args()
				err := glitcher.HorizontalSort(args[0], args[1])
				if err != nil {
					fmt.Printf("%v \n", err)
				}
			},
		},
		{
			Name:    "vertical-sort",
			Aliases: []string{"v-s"},
			Usage:   "Vertically sort the values of all the pixels in every row.",
			Action: func(c *cli.Context) {
				args := c.Args()
				err := glitcher.VerticalSort(args[0], args[1])
				if err != nil {
					fmt.Printf("%v \n", err)
				}
			},
		},
		{
			Name:    "invert",
			Aliases: []string{"inv"},
			Usage:   "Invert your picture",
			Action: func(c *cli.Context) {
				args := c.Args()
				glitcher.InvertImage(args[0], args[1])
			},
		},
	}
	app.RunAndExitOnError()
}
