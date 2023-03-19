package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "github.com/akhilmk/gosamples/grpc/pushnotification/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	startGrpcClient()
}

func startGrpcClient() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewNotifSubscriberClient(conn)

	// stream message
	log.Printf("stream started")
	// ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	stream, er := c.SubscribeMessage(context.Background(), &pb.SubscribeMsg{})
	if er != nil {
		log.Printf("stream error")
	}
	for {
		hellloreply, serr := stream.Recv()
		if serr == io.EOF {
			log.Printf("stream break")
			break
		}
		if serr != nil {
			log.Fatalf("stream failed: %v", serr)
		}
		log.Printf("stream notification: %v", hellloreply.GetReplymessage())
	}
	log.Printf("stream exit")
}
