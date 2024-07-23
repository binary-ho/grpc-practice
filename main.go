package main

import (
	"flag"
	"grpc-practice/cmd"
	"grpc-practice/config"
	"grpc-practice/grpc/server"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {

	flag.Parse()
	config := config.NewConfig(*configFlag)

	server.NewServer(config)

	cmd.NewApp(config)
}
