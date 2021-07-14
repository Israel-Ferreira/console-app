package app

import (
	"fmt"

	"github.com/urfave/cli"
)


func Gerar() *cli.App{
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"
	

	app.Action = func (c *cli.Context) error {
		fmt.Println("Bem vindo !!!")

		return nil
	}


	return app
}