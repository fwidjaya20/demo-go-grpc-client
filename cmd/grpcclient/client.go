package grpcclient

import (
	"github.com/fwidjaya20/demo-go-grpc-client/config"
	pb "github.com/fwidjaya20/demo-go-grpc-server/pkg/atm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
)

type Clients struct {
	TransferClient pb.TransfersClient
	Connections []*grpc.ClientConn
}

func InitClients() *Clients {
	transferClientConn, err := grpc.Dial(config.TRANSFER_GRPC, grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		PermitWithoutStream: true,
	}))

	if nil != err {
		log.Fatal("could not connect to", config.TRANSFER_GRPC, err)
	}

	log.Println("connected to client ", config.TRANSFER_GRPC)

	return &Clients{
		TransferClient: pb.NewTransfersClient(transferClientConn),
		Connections: []*grpc.ClientConn{
			transferClientConn,
		},
	}
}

func (c *Clients) Close() {
	for _, v := range c.Connections {
		_  = v.Close()
	}
}