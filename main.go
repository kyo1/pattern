package main

import (
	"fmt"
	"github.com/kyo1/pattern"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "generate a cyclic pattern",
			Action: func(c *cli.Context) error {
				length, err := strconv.ParseInt(c.Args().First(), 0, 0)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(pattern.Create(int(length)))
				return nil
			},
		},
		{
			Name:    "offset",
			Aliases: []string{"o", "search", "s"},
			Usage:   "find the offset",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "big",
					Aliases: []string{"b"},
					Usage:   "using big endian",
				},
			},
			Action: func(c *cli.Context) error {
				needle := c.Args().First()
				if strings.HasPrefix(needle, "0x") {
					x, err := strconv.ParseUint(needle, 0, 64)
					if err != nil {
						log.Fatal(err)
					}
					needle = pattern.Hex2str(x, c.Bool("big"))
				}
				offset := pattern.Offset(needle)
				if offset != -1 {
					fmt.Println(needle, "found at offset", offset)
				} else {
					fmt.Println("not found")
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
