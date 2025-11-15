package main

import (
	"flag"
	"log"
	"strings"

	"github.com/jinrai-js/go/pkg/jinrai"
)

func main() {
	dist, server, meta, port, assets, caching, verbose := initFlags()
	ssr := jinrai.NewX(*dist)
	if meta != nil {
		ssr.SetMeta(*meta)
	}

	log.Printf("jinrai: http://localhost:%d\n", *port)

	if *assets {
		log.Println("+ serve assets")
		ssr.ServeAssets(true)
	}

	if *verbose {
		ssr.Debug()
		log.Println("+ verbose")
	}

	ssr.SetStringProxy(*server)
	if caching != nil {
		ssr.SetChashing(strings.Split(*caching, ","))
	}
	ssr.ServeX(*port)
}

func initFlags() (dist, proxy, meta *string, port *int, assets *bool, caching *string, verbose *bool) {
	dist = flag.String("dist", "dist", "dist folder")
	proxy = flag.String("proxy", "/Api=http://localhost", `list of proxy servers
	example:
	/api=http://localhost,/profile=http//localhost:3000`)
	port = flag.Int("port", 80, "html server port")
	meta = flag.String("meta", "/Api/Meta/GetMetaDate", "meta date url")

	assets = flag.Bool("a", false, "serve assets")
	verbose = flag.Bool("v", true, "verbose")
	caching = flag.String("caching", "/Api", "caching proxy requests (example: \"/api/v1,/api/v2\")")

	flag.Parse()
	return
}
