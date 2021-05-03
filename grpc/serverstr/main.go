package main

import (
	context "context"
	"fmt"
	"io"
	"log"
	"net"
	sync "sync"
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
	UnimplementedStreamServiceServer
}

// FetchResponse implements serverstr.StreamService
func (server) FetchResponse(in *Request, srv StreamService_FetchResponseServer) error {

	log.Printf("fetch response for id : %d", in.Id)

	//use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()

			//time sleep to simulate server process time
			time.Sleep(time.Duration(count) * time.Second)
			resp := Response{Result: fmt.Sprintf("Request #%d For Id:%d", count, in.Id)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}

	wg.Wait()
	return nil
}

func runServer() {
	// create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	RegisterStreamServiceServer(s, server{})

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
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := NewStreamServiceClient(conn)
	in := &Request{Id: 1}
	stream, err := client.FetchResponse(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.Result)
		}
	}()

	<-done //we will wait until all response is received
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
