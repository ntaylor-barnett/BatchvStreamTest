// Code generated by goa v3.1.1, DO NOT EDIT.
//
// public service
//
// Command:
// $ goa gen github.com/ntaylor-barnett/BatchvStreamTest/design

package public

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// A mock service to test different service communication models
type Service interface {
	// Receives an array of payloads
	BatchGRPC(context.Context, *TestPayloadBatch) (res []*ResponsePayload, err error)
	// Receives an array of payloads
	StreamedBatchGRPC(context.Context, StreamedBatchGRPCServerStream) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "public"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"batchGRPC", "streamedBatchGRPC"}

// StreamedBatchGRPCServerStream is the interface a "streamedBatchGRPC"
// endpoint server stream must satisfy.
type StreamedBatchGRPCServerStream interface {
	// Send streams instances of "ResponsePayload".
	Send(*ResponsePayload) error
	// Recv reads instances of "TestPayload" from the stream.
	Recv() (*TestPayload, error)
	// Close closes the stream.
	Close() error
}

// StreamedBatchGRPCClientStream is the interface a "streamedBatchGRPC"
// endpoint client stream must satisfy.
type StreamedBatchGRPCClientStream interface {
	// Send streams instances of "TestPayload".
	Send(*TestPayload) error
	// Recv reads instances of "ResponsePayload" from the stream.
	Recv() (*ResponsePayload, error)
	// Close closes the stream.
	Close() error
}

// TestPayloadBatch is the payload type of the public service batchGRPC method.
type TestPayloadBatch struct {
	Records []*TestPayload
}

// TestPayload is the streaming payload type of the public service
// streamedBatchGRPC method.
type TestPayload struct {
	FirstField     string
	SecondField    string
	ThirdField     string
	OrganizationID uint32
}

// an example response
type ResponsePayload struct {
	FirstField  string
	FourthField string
}

// MakeUnauthenticated builds a goa.ServiceError from an error.
func MakeUnauthenticated(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthenticated",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
