package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	pb "github.com/akhilmk/gosamples/grpc/pushnotification/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	go handleRequests()
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
	log.Printf("grpc server started at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// server to implement push notification service
type server struct {
	pb.UnimplementedNotifSubscriberServer
}

func (s *server) SubscribeMessage(in *pb.SubscribeMsg, stream pb.NotifSubscriber_SubscribeMessageServer) error {
	log.Printf("subscriber called..")
	i := 0
	for t := range time.Tick(2 * time.Second) {
		_ = t
		log.Printf("publishing notification %v", i)
		if err := stream.Send(&pb.NotifReply{Replymessage: strconv.Itoa(i)}); err != nil {
			log.Printf("publishing err %v", err)
			return err
		}
		i++
	}
	log.Printf("subscriber end..")
	return nil
}

/***** REST API ****/
// REST notify endpoint
func notify(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "notify page!")
	fmt.Println("Endpoint Hit: notify")
}

func handleRequests() {
	http.HandleFunc("/notify", notify)
	fmt.Println("notify api started at.. 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
