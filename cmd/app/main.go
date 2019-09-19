package main

import (
	"github.com/fwidjaya20/demo-go-grpc-client/cmd/containers"
	"github.com/fwidjaya20/demo-go-grpc-client/cmd/grpcclient"
	"github.com/fwidjaya20/demo-go-grpc-client/cmd/http"
)

func main() {
	grpcClients := grpcclient.InitClients()

	svcContainers := containers.NewServiceContainer(grpcClients)

	http.HTTPServer(svcContainers)
}
