package containers

import (
	"github.com/fwidjaya20/demo-go-grpc-client/cmd/grpcclient"
	"github.com/fwidjaya20/demo-go-grpc-client/internal"
)

type ServiceContainer struct {
	SendMoneyService internal.ServiceInterface
}

func NewServiceContainer(clients *grpcclient.Clients) ServiceContainer {
	return ServiceContainer{
		SendMoneyService: internal.NewTransferService(clients.TransferClient),
	}
}
