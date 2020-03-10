// Code generated by goa v3.1.1, DO NOT EDIT.
//
// public gRPC client
//
// Command:
// $ goa gen github.com/ntaylor-barnett/BatchvStreamTest/design

package client

import (
	"context"

	publicpb "github.com/ntaylor-barnett/BatchvStreamTest/gen/grpc/public/pb"
	public "github.com/ntaylor-barnett/BatchvStreamTest/gen/public"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli publicpb.PublicClient
	opts    []grpc.CallOption
}

// StreamedBatchGRPCClientStream implements the
// public.StreamedBatchGRPCClientStream interface.
type StreamedBatchGRPCClientStream struct {
	stream publicpb.Public_StreamedBatchGRPCClient
}

// NewClient instantiates gRPC client for all the public service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: publicpb.NewPublicClient(cc),
		opts:    opts,
	}
}

// BatchGRPC calls the "BatchGRPC" function in publicpb.PublicClient interface.
func (c *Client) BatchGRPC() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildBatchGRPCFunc(c.grpccli, c.opts...),
			EncodeBatchGRPCRequest,
			DecodeBatchGRPCResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// StreamedBatchGRPC calls the "StreamedBatchGRPC" function in
// publicpb.PublicClient interface.
func (c *Client) StreamedBatchGRPC() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildStreamedBatchGRPCFunc(c.grpccli, c.opts...),
			nil,
			DecodeStreamedBatchGRPCResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Recv reads instances of "publicpb.StreamedBatchGRPCResponse" from the
// "streamedBatchGRPC" endpoint gRPC stream.
func (s *StreamedBatchGRPCClientStream) Recv() (*public.ResponsePayload, error) {
	var res *public.ResponsePayload
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewResponsePayload(v), nil
}

// Send streams instances of "publicpb.TestPayload" to the "streamedBatchGRPC"
// endpoint gRPC stream.
func (s *StreamedBatchGRPCClientStream) Send(res *public.TestPayload) error {
	v := NewTestPayload(res)
	return s.stream.Send(v)
}

func (s *StreamedBatchGRPCClientStream) Close() error {
	// Close the send direction of the stream
	return s.stream.CloseSend()
}
