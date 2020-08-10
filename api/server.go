package api

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/leandroudala/golang_jwt/api/auto"
	"github.com/leandroudala/golang_jwt/api/config"
	"github.com/leandroudala/golang_jwt/api/router"
)

// Run start the api server
func Run() {
	config.Load()

	// auto.Load()

	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Printf("Starting app at localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
