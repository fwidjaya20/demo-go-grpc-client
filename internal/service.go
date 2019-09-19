package internal

import (
	"context"
	pb "github.com/fwidjaya20/demo-go-grpc-server/pkg/atm"
)

type ServiceInterface interface {
	SendMoney(transfer *pb.Transfer) *pb.TransferResponse
}

type service struct {
	Transfer pb.TransfersClient
}

func NewTransferService(transferClient pb.TransfersClient) ServiceInterface {
	return &service{
		Transfer: transferClient,
	}
}

func (s *service) SendMoney(transfer *pb.Transfer) *pb.TransferResponse {
	response, _ := s.Transfer.Send(context.Background(), transfer)

	return response
}