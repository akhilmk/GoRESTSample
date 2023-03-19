package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	pb "github.com/akhilmk/gosamples/grpc/pushnotification/proto"
	"google.golang.org/grpc"
)

var (
	port      = flag.Int("port", 50051, "The server port")
	notifChan = make(chan string)
	noCount   = 0
)

func main() {
	go startNotifApi()
	startGrpcServer()
}

func startGrpcServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotifSubscriberServer(s, &server{})
	log.Printf("grpc server started at : %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// server to implement push notification service
type server struct {
	pb.UnimplementedNotifSubscriberServer
}

func (s *server) SubscribeMessage(in *pb.SubscribeMsg, stream pb.NotifSubscriber_SubscribeMessageServer) error {
	log.Printf("new client subscribed ..")
	for {
		select {
		case m := <-notifChan:
			if err := stream.Send(&pb.NotifReply{Replymessage: m}); err != nil {
				log.Printf("publishing err %v", err)
				return err
			}
		}
	}
}

/***** REST API ****/
func startNotifApi() {
	http.HandleFunc("/notify", notify)
	log.Println("notify api started at : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func notify(w http.ResponseWriter, r *http.Request) {
	msg := "notification " + strconv.Itoa(noCount)
	notifChan <- msg
	fmt.Println(msg)
	fmt.Fprintf(w, "notify count:"+strconv.Itoa(noCount))
	noCount++
}
