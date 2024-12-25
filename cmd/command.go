package main

import (
	"github.com/urfave/cli/v2"

	"github.com/tanqiangyes/rouguelike/internal/config"
	"github.com/tanqiangyes/rouguelike/internal/window"
)

// RunCommand 启动子命令
func RunCommand() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "run game",
		Action: func(ctx *cli.Context) error {
			window.NewWindow(config.GetConfig()).Run()
			return nil
		},
	}
}
