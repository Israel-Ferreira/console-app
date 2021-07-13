package app

import "github.com/urfave/cli"


func Gerar() *cli.App{
	app := cli.NewApp()
	app.Name = ""

	return app
}