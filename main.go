package main

import (
	"fmt"
	"github.com/kyo1/pattern/cmd"
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
			Usage:   "generate a pattern",
			Action: func(c *cli.Context) error {
				length, err := strconv.ParseInt(c.Args().First(), 0, 0)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(cmd.Create(int(length)))
				return nil
			},
		},
		{
			Name:    "offset",
			Aliases: []string{"o", "search", "s"},
			Usage:   "search a offset",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "big",
					Aliases: []string{"b"},
					Usage:   "Using big endian",
				},
			},
			Action: func(c *cli.Context) error {
				pattern := c.Args().First()
				if strings.HasPrefix(pattern, "0x") {
					x, err := strconv.ParseUint(pattern, 0, 64)
					if err != nil {
						log.Fatal(err)
					}
					pattern = cmd.Hex2str(x, c.Bool("big"))
				}
				offset := cmd.Offset(pattern)
				if offset != -1 {
					fmt.Println(pattern, "found at offset", offset)
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
