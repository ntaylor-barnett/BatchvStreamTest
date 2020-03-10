#test

depend:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go	

gen: depend
	gobin -m -run goa.design/goa/v3/cmd/goa gen github.com/ntaylor-barnett/BatchvStreamTest/design

example:
	gobin -m -run goa.design/goa/v3/cmd/goa example github.com/ntaylor-barnett/BatchvStreamTest/design

build:
	go mod vendor
	go mod download
	go build ./cmd/public 