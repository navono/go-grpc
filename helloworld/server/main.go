package main

import (
	"log"
	"net"
	"strconv"

	pb "github.com/navono/go-grpc/helloworld/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":20188"
)

// Config for rpc server
type Config struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	log.Println("get request")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHello2(in *pb.HelloRequest, gs pb.Greeter_SayHello2Server) error {
	name := in.Name
	for i := 0; i < 100000000; i++ {
		gs.Send(&pb.HelloReply{Message: "Hello " + name + strconv.Itoa(i)})

		log.Println("Hello " + name + strconv.Itoa(i))

		//time.Sleep(time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	//reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
