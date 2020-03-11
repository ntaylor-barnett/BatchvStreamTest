#test

depend:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go	
	go get -u github.com/myitcv/gobin
	

gen:
	gobin -run goa.design/goa/v3/cmd/goa gen github.com/ntaylor-barnett/BatchvStreamTest/design

example:
	gobin -run goa.design/goa/v3/cmd/goa example github.com/ntaylor-barnett/BatchvStreamTest/design

build:
	go mod vendor
	go mod download
	go build ./cmd/public 
	go build ./cmd/sender

.PHONY: gen build example