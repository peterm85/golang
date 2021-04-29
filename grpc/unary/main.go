package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

const (
	port        = ":50051"
	address     = "localhost:50051"
	defaultName = "world"
)

// server is used to implement unary.GreeterServer.
type server struct {
	UnimplementedGreeterServer
}

// SayHello implements unary.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func runServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func runClient() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func main() {
	var eg errgroup.Group
	eg.Go(func() error {
		log.Printf("Running grpc server")
		runServer()
		return nil
	})

	eg.Go(func() error {
		log.Printf("Running grpc client")
		runClient()
		return nil
	})

	if err := eg.Wait(); err != nil {
		fmt.Println("Encountered error:", err)
	}
}
