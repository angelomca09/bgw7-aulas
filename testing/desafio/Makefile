test:
	go test ./... -coverprofile=coverage.out
	grep -v "_mock\.go" coverage.out > coverage_filtered.out && mv coverage_filtered.out coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out