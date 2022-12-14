package main

import (
	"log"
	"net/http"

	"github.com/galaxyed/prometheus-proxy/internal/conf"
	"github.com/galaxyed/prometheus-proxy/internal/server"
)

func main() {
	// Read Config
	cfgPath, err := server.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := conf.NewConfig(cfgPath)

	// initialize a reverse proxy and pass the actual backend server url here
	proxy, err := server.NewProxy("http://10.100.0.52:9090", cfg)
	if err != nil {
		panic(err)
	}
	log.Println(cfg.Policies[0].Name)

	// handle all requests to your server using the proxy
	http.HandleFunc("/", server.ProxyRequestHandler(proxy))
	log.Println("Server Started")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
