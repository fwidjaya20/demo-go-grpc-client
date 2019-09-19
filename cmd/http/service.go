package http

import (
	"fmt"
	"github.com/fwidjaya20/demo-go-grpc-client/cmd/containers"
	"github.com/fwidjaya20/demo-go-grpc-client/config"
	"github.com/fwidjaya20/demo-go-grpc-client/internal/transports/http"
	"github.com/go-chi/chi"
	"log"
	netHttp "net/http"
)

// HTTPServer .
func HTTPServer(containers containers.ServiceContainer) {
	var router *chi.Mux

	router = chi.NewRouter()

	router.Post("/transfer", http.Transfer(containers.SendMoneyService))

	log.Println(fmt.Sprintf("HTTP Transport running on %s", config.HTTP_ADDR))
	_ = netHttp.ListenAndServe(config.HTTP_ADDR, router)
}
