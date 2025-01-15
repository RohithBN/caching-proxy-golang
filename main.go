package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/RohithBN/caching-proxy/proxy"
)

func main() {
	Port := flag.Int("port", 8080, "port to listen on")
	Origin := flag.String("origin", "", "origin to allow")
	clearCache := flag.Bool("clear-cache", false, "clear the cache")
	flag.Parse()
	if *Origin != "" || *Port != 0 {
		if *Origin == "" {
			log.Fatal("Origin server URL is required when starting the server")
		}
	}
	proxy := proxy.NewProxy(*Origin, *Port)
	if *clearCache {
		proxy.ClearCache()
	}
	fmt.Println("Listening on port", *Port)
	http.HandleFunc("/", proxy.ServeHTTP)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *Port), nil))
}
