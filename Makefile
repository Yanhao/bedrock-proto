proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		metaserver.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		dataserver.proto
	protoc --cpp_out=. dataserver.proto
	protoc --cpp_out=. metaserver.proto
	protoc --cpp_out=. proxy.proto


clean:
	rm -f *.go *.h *.cc

.PHONY: proto clean
