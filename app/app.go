package app

import (
	"console-app/services"
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


	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca Ips de Endereços na Internet",
			Flags: flags,
			Action: services.BuscarIp,
		},

		{
			Name: "server",
			Usage: "Busca Nomes de Servidores na Internet",
			Flags: flags,
			Action: services.BuscarNomeDoServidor,
		},
	}


	return app
}



