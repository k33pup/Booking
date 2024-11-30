package deliverySystemClient

import (
	"flag"
	"log"

	pb "booking/notification_svc/pkg/deliverySystem/api/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type SendlerClient struct {
	conn   *grpc.ClientConn
	client pb.SenderClient
}

func NewSenderClient() (*SendlerClient, error) {
	client := SendlerClient{}
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client.conn = conn
	client.client = pb.NewSenderClient(conn)
	return &client, nil
}

func NewClient() *pb.SenderClient {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSenderClient(conn)
	return &c
}
