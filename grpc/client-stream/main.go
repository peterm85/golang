package main

import (
	context "context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"golang.org/x/sync/errgroup"
	grpc "google.golang.org/grpc"
)

const (
	port    = ":50051"
	address = "localhost:50051"
)

// server is used to implement serverstr.StreamService.
type server struct {
	UnimplementedClientStreamServiceServer
}

// FetchResponse implements serverstr.StreamService
func (server) FetchResponse(srv ClientStreamService_FetchResponseServer) error {
	count := 0
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			resp := Response{Result: fmt.Sprintf("Total requests fetched: %d", count)}
			return srv.SendAndClose(&resp)
		}
		if err != nil {
			return err
		}
		count++
		log.Printf("fetch request for id : %d", req.Id)
	}
}

func runServer() {
	// create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()

	RegisterClientStreamServiceServer(s, server{})

	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("shutting-down server")
}

func runClient() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := NewClientStreamServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.FetchResponse(ctx)
	if err != nil {
		log.Fatalf("%v.FetchResponse(_) = _, %v", client, err)
	}
	for i := 0; i < 5; i++ {
		req := Request{Id: int32(i)}
		if err := stream.Send(&req); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, req, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Server response: %v", reply)
	log.Println("shutting-down client")
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
