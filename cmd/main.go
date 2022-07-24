package main

import (
	"log"

	"github.com/asialeaf/automation/pkg/core/connector"
)

func main() {

	var hosts []connector.Host
	Hosts := []connector.BaseHost{
		{
			Name:     "node1",
			Address:  "10.1.2.104",
			User:     "root",
			Password: "abc@123",
			Port:     22,
			Timeout:  10,
		},
	}
	for _, v := range Hosts {
		hosts = append(hosts, &v)
		for _, host := range hosts {

			dialer := connector.NewDialer()
			conn, err := dialer.Connect(host)
			if err != nil {
				log.Fatal(err)
			}
			stdout, _, err := conn.Exec("uname -sr", host)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(stdout)
			dialer.Close(host)

		}
	}
}
