package main

import (
	context "context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strings"
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
	UnimplementedChatStreamServiceServer
}

// FetchResponse implements serverstr.StreamService
func (server) Chat(srv ChatStreamService_ChatServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF || strings.Contains(req.Msg, "bye") {
			log.Printf("fetch client msg: %v", req)
			return srv.Send(&Response{Nickname: "Server", Msg: "Bye bye " + req.Nickname})
		} else if err != nil {
			return err
		}
		log.Printf("fetch client msg: %v", req)
		srv.Send(&Response{Nickname: "Server", Msg: getMsg()})
	}
}

func getMsg() string {
	r := []string{"Really?", "Great!", "Wtf!", "I can't believe it"}
	s1 := rand.NewSource(time.Now().Local().UnixNano())
	r1 := rand.New(s1)
	min := 0
	index := r1.Intn(len(r)-min) + min
	return r[index] + ", tell me more"
}

func runServer() {
	// create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	RegisterChatStreamServiceServer(s, server{})

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
	client := NewChatStreamServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Start chatting
	stream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("%v.Chat(_) = _, %v", client, err)
	}
	defer stream.CloseSend()

	for _, s := range []string{"Hi!", "My name is Client", "I'm from Spain", "I looove pizza", "And I play basket", "I have to go, bye!"} {
		req := Request{Nickname: "Client", Msg: s}
		if err := stream.Send(&req); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, req, err)
		}

		reply, err := stream.Recv()
		if err != nil {
			log.Fatalf("%v.Recv() got error %v, want %v", stream, err, nil)
		}
		log.Printf("fetch server msg: %v", reply)
	}
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
