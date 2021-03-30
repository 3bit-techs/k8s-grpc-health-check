# download dependecies and configure go path
configure-go:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	@echo "##### Execute the export below before compile-pb #####"
	export PATH=$(PATH):$(HOME)/go/bin

# compile hello world protobuffer
compile-pb:
	protoc --go_out=plugins=grpc:. proto/helloworld.proto

# generate docker image for server
gen-server-image:
	docker build -t 3bittechs/k8s-grpc-server:1.0 .

# run gRPC server local
run-server:
	docker run --network=host 3bittechs/k8s-grpc-server:1.0