package publicapi

import (
	"context"
	"log"

	public "github.com/ntaylor-barnett/BatchvStreamTest/gen/public"
)

// public service example implementation.
// The example methods log the requests and return zero values.
type publicsrvc struct {
	logger *log.Logger
}

// NewPublic returns the public service implementation.
func NewPublic(logger *log.Logger) public.Service {
	return &publicsrvc{logger}
}

// Receives an array of payloads
func (s *publicsrvc) BatchGRPC(ctx context.Context, p []*public.TestPayload) (res []*public.TestPayload, err error) {
	s.logger.Print("public.batchGRPC")
	return
}

// Receives an array of payloads
func (s *publicsrvc) StreamedBatchGRPC(ctx context.Context, stream public.StreamedBatchGRPCServerStream) (err error) {
	s.logger.Print("public.streamedBatchGRPC")
	return
}
