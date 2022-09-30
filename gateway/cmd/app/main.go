package main

import (
	"flag"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/config"
	"github.com/nurtilekassankhan/meetup/gateway/server"
	"log"
)

func main() {
	confPath := flag.String("config", "configs/config.yaml", "config file path")
	flag.Parse()

	// create config
	conf, err := config.New(*confPath)
	if err != nil {
		log.Fatal(err)
	}

	serv, err := server.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	if err = serv.RunGracefulShutdown(); err != nil {
		log.Println(err)
	}
}
