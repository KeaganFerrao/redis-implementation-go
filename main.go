package main

import (
	"flag"
	"log"

	"github.com/KeaganFerrao/redis-implementation-go/config"
	"github.com/KeaganFerrao/redis-implementation-go/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the server")
	flag.IntVar(&config.Port, "port", 7379, "port for the server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("Starting the server...")
	server.RunAsyncTCPServer()
}
