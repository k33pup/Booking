package deliverySystemClient

import (
	"context"
	"flag"
	"log"
	"time"

	pb "booking/notification_svc/internal/pkg/api/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type MockDeliverySystemClient struct {
	client pb.SenderClient
}

func NewMockDeliverySystemClient() (*MockDeliverySystemClient, error) {
	client := MockDeliverySystemClient{}
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client.client = pb.NewSenderClient(conn)
	return &client, nil
}

func (ds *MockDeliverySystemClient) SendMail(str string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := ds.client.SendMail(ctx, &pb.SendMailRequest{Mail: str})
	if err != nil {
		log.Fatalf("could not send mail: %v", err)
		return err
	}
	return nil
}
