package main

import (
	"SimpleInferencer/internal/server"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/logger"
)

var (
	verbose = flag.Bool("verbose", true, "print info level logs to stdout")
	port    = flag.Int("port", 5001, "port for pool manager to listen")
)

func main() {
	logger.SetFlags(log.Lmicroseconds | log.Lshortfile)
	flag.Parse()
	logger.Init("SimpleInferencer", *verbose, false, ioutil.Discard)
	ctx := context.Background()
	s, _ := server.NewServer(ctx)

	address := fmt.Sprintf(":%v", *port)
	hs := &http.Server{
		Addr:    address,
		Handler: s.Routes(),
	}
	logger.Infof("SimpleInferencer started and listening in %v", *port)
	hs.ListenAndServe()
}
