package publicapi

import (
	"context"
	"log"

	public "github.com/ntaylor-barnett/BatchvStreamTest/gen/public"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
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
func (s *publicsrvc) BatchGRPC(ctx context.Context, p *public.TestPayloadBatch) (res []*public.ResponsePayload, err error) {
	s.logger.Print("public.batchGRPC")
	return
}

// Receives an array of payloads
func (s *publicsrvc) StreamedBatchGRPC(ctx context.Context, stream public.StreamedBatchGRPCServerStream) (err error) {
	s.logger.Print("public.streamedBatchGRPC")
	eg, egctx := errgroup.WithContext(ctx)
	datachan := make(chan *public.TestPayload, 1000)
	eg.Go(func() error {
		// sender
		select {
		case <-egctx.Done():
			return nil
		case v, ok := <-datachan:
			if !ok {
				return nil
			}
			// we will test with immediate send at the moment. This is best case scenario
			resp := &public.ResponsePayload{
				FirstField:  v.FirstField,
				FourthField: "yeah, no probs. All good mate",
			}
			err := stream.Send(resp)
			if err != nil {
				return errors.Wrap(err, "error when sending reply")
			}
		}
		return nil
	})

	eg.Go(func() error {
		//reciever
		defer close(datachan)
		payload, err := stream.Recv()
		if err != nil {
			return err
		}
		select {
		case <-egctx.Done():
			return nil
		case datachan <- payload:
			// nothing to do here. data was recieved
		}
		return nil
	})
	return eg.Wait()

}
