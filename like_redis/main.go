package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/alexeyfrank/learn_go/like_redis/client"
	"github.com/alexeyfrank/learn_go/like_redis/server"
)

func main() {
	var (
		isServer bool
		isClient bool
		addr     string
	)
	flag.BoolVar(&isServer, "server", false, "")
	flag.BoolVar(&isClient, "client", false, "")
	flag.StringVar(&addr, "addr", "localhost:8888", "")
	flag.Parse()

	if isServer {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
		log.Println("Is server")

		storage := server.NewStorage()
		server := server.NewServer(addr, storage)

		server.Run()
	}

	if isClient {
		fmt.Println("Is client")

		commandReader := client.NewCommandReader(os.Stdin)
		responseWriter := client.NewResponseWriter(os.Stdout)
		client, err := client.NewClient(addr, commandReader, responseWriter)

		if err != nil {
			log.Fatal(err)
		}

		client.Run()
	}
}
