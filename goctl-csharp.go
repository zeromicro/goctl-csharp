package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
	"github.com/zeromicro/goctl-csharp/action"
)

var (
	version  = "20250207"
	commands = []*cli.Command{
		{
			Name:   "csharp",
			Usage:  "generates http client for csharp",
			Action: action.CSharp,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "namespace",
					Usage:    "the namespace of csharp",
					Required: true,
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Usage = "a plugin of goctl to generate http client code for csharp."
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-csharp: %+v\n", err)
	}
}
