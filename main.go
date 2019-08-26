package main

import (
	"context"
	"fmt"
	"log"
	"net"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	"google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	"google.golang.org/genproto/googleapis/ads/googleads/v1/services"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Search(ctx context.Context, in *services.SearchGoogleAdsRequest) (*services.SearchGoogleAdsResponse, error) {
	log.Printf("Received(Service): %v", in)
	return &services.SearchGoogleAdsResponse{
		Results: []*services.GoogleAdsRow{
			&services.GoogleAdsRow{
				SearchTermView: &resources.SearchTermView{
					SearchTerm: &wrappers.StringValue{Value: "abc"},
				},
				Metrics: &common.Metrics{
					Impressions:                     &wrappers.Int64Value{Value: 0},
					Clicks:                          &wrappers.Int64Value{Value: 0},
					Ctr:                             &wrappers.DoubleValue{Value: 0.0},
					Conversions:                     &wrappers.DoubleValue{Value: 0.0},
					ConversionsFromInteractionsRate: &wrappers.DoubleValue{Value: 0.0},
					ConversionsValue:                &wrappers.DoubleValue{Value: 0.0},
					CostMicros:                      &wrappers.Int64Value{Value: 0},
				},
			},
		},
		TotalResultsCount: 999,
	}, nil
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
