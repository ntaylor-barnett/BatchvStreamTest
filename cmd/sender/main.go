package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"

	public "github.com/ntaylor-barnett/BatchvStreamTest/gen/public"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	//public "../../gen/public"
	grpcmap "github.com/ntaylor-barnett/BatchvStreamTest/gen/grpc/public/client"
	grpc "google.golang.org/grpc"
)

func main() {
	modeStrPtr := flag.String("mode", "", "either stream or batch")
	recordCount := flag.Int("records", 1000, "How many records to push")
	iterations := flag.Int("iter", 1, "how many times to execute the test")
	address := flag.String("addr", "127.0.0.1:8000", "address to connect to")
	flag.Parse()

	client := GetClient(*address)
	ctx := context.Background()
	switch strings.ToLower(*modeStrPtr) {
	case "stream":
		StreamData(ctx, client, *recordCount, *iterations)
	case "batch":
		BatchData(ctx, client, *recordCount, *iterations)
	default:
		fmt.Println("Invalid command")
		flag.PrintDefaults()
	}

}

func BatchData(ctx context.Context, client *public.Client, records, repeat int) {
	fmt.Println(fmt.Sprintf("Started Batch test. Records: %v, Repetitions: %v", records, repeat))
	ctx, canceller := context.WithCancel(ctx)
	defer canceller()
	var totaltime float64
	for iter := 1; iter <= repeat; iter++ {
		timestarted := time.Now()
		p := &public.TestPayloadBatch{
			Records: make([]*public.TestPayload, records),
		}
		datachan := generateRecords(ctx, records)
		i := 0
		for rec := range datachan {
			p.Records[i] = rec
			i++
		}
		res, err := client.BatchGRPC(ctx, p)
		if err != nil {
			panic(err)
		}
		_ = res
		elapsed := time.Since(timestarted)
		totaltime += elapsed.Seconds()
		fmt.Println(fmt.Sprintf("Iteration completed in %vms", elapsed.Seconds()*1000))
	}
	fmt.Println(fmt.Sprintf("Average time %vms", (totaltime/float64(repeat))*1000))
}

func StreamData(ctx context.Context, client *public.Client, records, repeat int) {
	fmt.Println(fmt.Sprintf("Started Bidrectional streaming test. Records: %v, Repetitions: %v", records, repeat))
	ctx, canceller := context.WithCancel(ctx)
	defer canceller()
	var totaltime float64
	for iter := 1; iter <= repeat; iter++ {
		timestarted := time.Now()
		// this will execute the loops
		stream, err := client.StreamedBatchGRPC(ctx)
		if err != nil {
			panic(err)
		}
		// we need to set up two main goroutines. One to send data, one to recieve
		eg, egctx := errgroup.WithContext(ctx)
		datachan := generateRecords(ctx, records)
		sendcomplete := false
		eg.Go(func() (rErr error) {
			// sender
			defer func() {
				err := stream.Close()
				if rErr == nil && err != nil {
					rErr = errors.Wrap(err, "failed to close output stream")
				}
			}()
			select {
			case <-egctx.Done():
				return
			case p, ok := <-datachan:
				if !ok {
					return
				}
				err := stream.Send(p)
				if err != nil {
					return errors.Wrap(err, "error returned from sender")
				}
			}
			sendcomplete = true
			return nil
		})
		eg.Go(func() error {
			// reciever
			for {
				p, err := stream.Recv()
				if err == io.EOF {
					if !sendcomplete {
						return errors.New("server closed the response stream before we had finished sending data")
					}
					return nil // graceful closure
				}
				_ = p // we dont actually care about the response, only that we got it
			}
		})
		resulterr := eg.Wait()
		if resulterr != nil {
			panic(resulterr)
		}
		elapsed := time.Since(timestarted)
		totaltime += elapsed.Seconds()
		fmt.Println(fmt.Sprintf("Iteration completed in %vms", elapsed.Seconds()*1000))
	}
	fmt.Println(fmt.Sprintf("Average time %vms", (totaltime/float64(repeat))*1000))

}

func generateRecords(ctx context.Context, recordCount int) chan *public.TestPayload {
	out := make(chan *public.TestPayload, 1000)
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	go func() {
		defer close(out)
		for i := 0; i < recordCount; i++ {

			p := &public.TestPayload{
				FirstField:     strconv.Itoa(i),
				SecondField:    String(seededRand, 20),
				ThirdField:     String(seededRand, 40),
				OrganizationID: 12,
			}
			select {
			case <-ctx.Done():
				return
			case out <- p:
			}
		}
	}()
	return out
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(r *rand.Rand, length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset)-1)]
	}
	return string(b)
}

func String(r *rand.Rand, length int) string {
	return StringWithCharset(r, length, charset)
}

func GetClient(address string) *public.Client {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	gclient := grpcmap.NewClient(conn)
	client := public.NewClient(gclient.BatchGRPC(), gclient.StreamedBatchGRPC())
	return client
}
