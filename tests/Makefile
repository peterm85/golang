import ../Makefile

deps: 
	echo "[dep] fetching package dependencies"; \
	go get github.com/golang/mock/gomock; \
	go get github.com/golang/mock/mockgen

mockgen: 
	mockgen -source='iclient.go' IClient -destination='mocks/mock_iclient.go'

test-run:
	go test -v -run=TestServer -timeout 10s

test-subrun:
	go test -v -run=Sum/All

test-bench:
	go test -benchmem -run=^$ -bench .
