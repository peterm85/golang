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