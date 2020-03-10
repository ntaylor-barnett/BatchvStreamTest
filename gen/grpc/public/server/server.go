// Code generated by goa v3.1.1, DO NOT EDIT.
//
// public gRPC server
//
// Command:
// $ goa gen github.com/ntaylor-barnett/BatchvStreamTest/design

package server

import (
	"context"

	publicpb "github.com/ntaylor-barnett/BatchvStreamTest/gen/grpc/public/pb"
	public "github.com/ntaylor-barnett/BatchvStreamTest/gen/public"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the publicpb.PublicServer interface.
type Server struct {
	BatchGRPCH         goagrpc.UnaryHandler
	StreamedBatchGRPCH goagrpc.StreamHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// StreamedBatchGRPCServerStream implements the
// public.StreamedBatchGRPCServerStream interface.
type StreamedBatchGRPCServerStream struct {
	stream publicpb.Public_StreamedBatchGRPCServer
}

// New instantiates the server struct with the public service endpoints.
func New(e *public.Endpoints, uh goagrpc.UnaryHandler, sh goagrpc.StreamHandler) *Server {
	return &Server{
		BatchGRPCH:         NewBatchGRPCHandler(e.BatchGRPC, uh),
		StreamedBatchGRPCH: NewStreamedBatchGRPCHandler(e.StreamedBatchGRPC, sh),
	}
}

// NewBatchGRPCHandler creates a gRPC handler which serves the "public" service
// "batchGRPC" endpoint.
func NewBatchGRPCHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeBatchGRPCRequest, EncodeBatchGRPCResponse)
	}
	return h
}

// BatchGRPC implements the "BatchGRPC" method in publicpb.PublicServer
// interface.
func (s *Server) BatchGRPC(ctx context.Context, message *publicpb.BatchGRPCRequest) (*publicpb.BatchGRPCResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "batchGRPC")
	ctx = context.WithValue(ctx, goa.ServiceKey, "public")
	resp, err := s.BatchGRPCH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "unauthenticated":
				return nil, goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*publicpb.BatchGRPCResponse), nil
}

// NewStreamedBatchGRPCHandler creates a gRPC handler which serves the "public"
// service "streamedBatchGRPC" endpoint.
func NewStreamedBatchGRPCHandler(endpoint goa.Endpoint, h goagrpc.StreamHandler) goagrpc.StreamHandler {
	if h == nil {
		h = goagrpc.NewStreamHandler(endpoint, nil)
	}
	return h
}

// StreamedBatchGRPC implements the "StreamedBatchGRPC" method in
// publicpb.PublicServer interface.
func (s *Server) StreamedBatchGRPC(stream publicpb.Public_StreamedBatchGRPCServer) error {
	ctx := stream.Context()
	ctx = context.WithValue(ctx, goa.MethodKey, "streamedBatchGRPC")
	ctx = context.WithValue(ctx, goa.ServiceKey, "public")
	_, err := s.StreamedBatchGRPCH.Decode(ctx, nil)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "unauthenticated":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	ep := &public.StreamedBatchGRPCEndpointInput{
		Stream: &StreamedBatchGRPCServerStream{stream: stream},
	}
	err = s.StreamedBatchGRPCH.Handle(ctx, ep)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "unauthenticated":
				return goagrpc.NewStatusError(codes.Unauthenticated, err, goagrpc.NewErrorResponse(err))
			}
		}
		return goagrpc.EncodeError(err)
	}
	return nil
}

// Send streams instances of "publicpb.StreamedBatchGRPCResponse" to the
// "streamedBatchGRPC" endpoint gRPC stream.
func (s *StreamedBatchGRPCServerStream) Send(res []*public.TestPayload) error {
	v := NewStreamedBatchGRPCResponse(res)
	return s.stream.Send(v)
}

// Recv reads instances of "publicpb.StreamedBatchGRPCStreamingRequest" from
// the "streamedBatchGRPC" endpoint gRPC stream.
func (s *StreamedBatchGRPCServerStream) Recv() ([]*public.TestPayload, error) {
	var res []*public.TestPayload
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewStreamedBatchGRPCStreamingRequest(v), nil
}

func (s *StreamedBatchGRPCServerStream) Close() error {
	// nothing to do here
	return nil
}
