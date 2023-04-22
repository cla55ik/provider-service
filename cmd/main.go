package main

import (
	"flag"
	"github.com/cla55ik/provider-service/include/server"
	"github.com/cla55ik/provider-service/pkg/simulator"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Println("Configuration file error")
	}

	mode := flag.String("mode", "server", "a string")
	addr := flag.String("addr", viper.GetString("server.addr"), "a string")
	port := flag.String("port", viper.GetString("server.port"), "a string")
	flag.Parse()

	simUrl := viper.GetString("simulator.addr") + ":" + viper.GetString("simulator.port")
	url := *addr + ":" + *port

	switch {
	case *mode == "server":
		server.StartServer(url)
	case *mode == "simulator":
		simulator.Start(simUrl)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
