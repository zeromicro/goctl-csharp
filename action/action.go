package action

import (
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/goctl-csharp/generate"
)

func CSharp(ctx *cli.Context) error {
	ns := ctx.String("namespace")

	plugin, err := plugin.NewPlugin()
	if err != nil {
		return err
	}

	return generate.CSharpCommand(plugin, ns)
}
