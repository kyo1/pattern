package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/kyo1/pattern"
	"log"
	"strconv"
	"strings"
)

var cli struct {
	Create struct {
		Length int `arg name:"length" help:"length of the pattern to be generated" default:"0"`
	} `cmd help:"generate a cyclic pattern"`

	Offset struct {
		BigEndian bool `default:"false"`

		Needle string `arg name:"needle" default:""`
	} `cmd help:"find the offset"`
}

func main() {
	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "create <length>":
		fmt.Println(pattern.Create(cli.Create.Length))
	case "offset <needle>":
		if strings.HasPrefix(cli.Offset.Needle, "0x") {
			x, err := strconv.ParseUint(cli.Offset.Needle, 0, 64)
			if err != nil {
				log.Fatal(err)
			}
			cli.Offset.Needle = pattern.Hex2str(x, cli.Offset.BigEndian)

		}
		offset := pattern.Offset(cli.Offset.Needle)
		if offset != -1 {
			fmt.Println(cli.Offset.Needle, "found at offset", offset)
		} else {
			fmt.Println("not found")
		}
	default:
		panic(ctx.Command)
	}
}
