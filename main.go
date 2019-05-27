package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/pmmp/pmmp-stats/db"
	"github.com/pmmp/pmmp-stats/process"
	"github.com/pmmp/pmmp-stats/web"
)

var (
	webAddr = flag.String("web.addr", "localhost:8080", "The address to listen for http connections on")
)

func main() {
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		for range c {
			os.Exit(0)
		}
	}()

	dataBase := db.Setup()
	go web.Serve(c, *webAddr, dataBase)
	process.Process(c, dataBase)
}
