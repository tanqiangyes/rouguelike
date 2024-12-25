package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	configs "github.com/tanqiangyes/rouguelike/config"
	"github.com/tanqiangyes/rouguelike/internal/config"
	"github.com/tanqiangyes/rouguelike/pkg/i18n"
	"github.com/tanqiangyes/rouguelike/pkg/logger"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = "Rouguelike"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Usage: "config file",
		},
		&cli.StringFlag{
			Name:  "lang",
			Usage: "lang file",
		},
		&cli.StringFlag{
			Name:  "env",
			Usage: "environment",
		},
	}
	app.Before = func(ctx *cli.Context) error {
		return mustLoadConfig(ctx)
	}
}

func main() {
	app.Commands = []*cli.Command{
		RunCommand(),
	}
	if err := app.Run(os.Args); err != nil {
		logger.NewEntry().WithError(err).Fatal()
	}
}

func mustLoadConfig(ctx *cli.Context) error {
	if confFile := ctx.String("config"); confFile != "" {
		if err := config.LoadFile(confFile); err != nil {
			return errors.Wrap(err, "load config file")
		}
	} else {
		if err := configs.InitConfData(ctx.String("env")); err != nil {
			return errors.Wrap(err, "load env config")
		}
	}

	cfg := config.GetConfig()
	cfg.Icon = configs.GetIcon()
	if err := config.CheckConfig(cfg); err != nil {
		return errors.Wrap(err, "check config")
	}

	if langFile := ctx.String("lang"); langFile != "" {
		if err := i18n.Load(langFile); err != nil {
			return errors.Wrap(err, "load language translation file")
		}
	} else {
		if err := i18n.LoadData(configs.GetLangData()); err != nil {
			return errors.Wrap(err, "load language translation file")
		}
	}

	return nil
}
