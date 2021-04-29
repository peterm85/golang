test: 
	echo "[go test] running tests and collecting coverage metrics"; \
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

test-units: 
	go test -v -tags unit_tests ./...

test-integration: 
	go test -v -tags integration_tests ./...

build-prod:
	go build -tags prod ./...

doc:
	go doc
env:
	go env

deps:
	go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/unary/unary.proto