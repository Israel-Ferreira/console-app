package services

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)


func BuscarIp(c *cli.Context){
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err !=  nil {
		log.Fatalln("Erro: {}", err)
	}


	for _, ip := range ips {
		fmt.Println("Ip: ", ip)
	}
}