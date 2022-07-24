package main

import (
	"github.com/asialeaf/automation/pkg/core/connector"
	"github.com/asialeaf/automation/pkg/core/logger"
)

func main() {
	logger.Log = logger.NewLogger("", true)
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
				logger.Log.Errorf("connector failed: %s", host.GetName())
			}
			stdout, _, err := conn.Exec("uname -sr", host)
			if err != nil {
				logger.Log.Errorf("exec failed: %s", err)
			}
			logger.Log.Infof("Output: %s", stdout)
			conn.Chmod("/root/anaconda-ks.cfg", 777)
			dialer.Close(host)

		}
	}
}
