package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/genproto/googleapis/ads/googleads/v1/services"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Search(ctx context.Context, in *services.SearchGoogleAdsRequest) (*services.SearchGoogleAdsResponse, error) {
	log.Printf("Received(Service): %v", in)

	file, err := os.Open("./search_term_view.json")
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var response services.SearchGoogleAdsResponse
	err = jsonpb.Unmarshal(file, &response)
	if err != nil {
		log.Fatalf("%s", err)
	}

	return &response, nil
}

func (s *server) Mutate(ctx context.Context, in *services.MutateGoogleAdsRequest) (*services.MutateGoogleAdsResponse, error) {
	panic(fmt.Sprintf("Mutate of SearchGoogleAdsRequest has'nt been implemented yet."))
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	services.RegisterGoogleAdsServiceServer(s, &server{})

	log.Printf("start to serve: %v", port)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve on port %s", err)
	}
}
