protoc -I ./ proto/*.proto -I/usr/local/include \
	-I$GOPATH/src \
	-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:./gen/proto
protoc -I ./ proto/*.proto -I/usr/local/include \
	-I$GOPATH/src \
	-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--grpc-gateway_out=logtostderr=true,paths=source_relative:./gen \
