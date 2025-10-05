package main

import (
	"log"

	"github.com/jinrai-js/go/pkg/jinrai"
)

func main() {
	port := 8422

	ssr, err := jinrai.New("./test/export", "https://fld.ru/Api", "https://fld.ru/Api/Meta/GetMetaTags")
	if err != nil {
		log.Fatal(err)
	}

	if err = ssr.Serve(port); err != nil {
		log.Fatal(err)
	}

	log.Println("Server run on port:", port)
}
