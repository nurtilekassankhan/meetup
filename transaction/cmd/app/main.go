package main

import (
	"context"
	"flag"
	"github.com/nurtilekassankhan/meetup/transaction/grpc_server"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/config"
	"log"
)

func main() {
	confPath := flag.String("config", "configs/config.yaml", "config file path")
	flag.Parse()

	conf, err := config.New(*confPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("create server...")
	serv, err := grpc_server.New(context.Background(), conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server launch")
	if err = serv.Run(); err != nil {
		log.Fatal(err)
	}
}
