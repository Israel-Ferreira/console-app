package services

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func BuscarIp(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatalln("Erro: {}", err)
	}

	for _, ip := range ips {
		fmt.Println("Ip: ", ip)
	}
}

func BuscarNomeDoServidor(c *cli.Context) {
	host := c.String("host")

	serverNames, err :=  net.LookupNS(host)

	if err != nil {
		log.Fatalln("Erro: {}", err)
	}


	for _, serverName := range serverNames {
		fmt.Println("Server Name: ",serverName.Host )
	}

}
