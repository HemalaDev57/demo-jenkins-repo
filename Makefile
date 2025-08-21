    .PHONY: build test scan

    build:
		go mod tidy
		go build -o dist/app ./cmd/app

    test:
		gotestsum --junitfile test-report.xml -- -v ./...

    scan:
		gosec -fmt junit-xml -out gosec-report.xml ./...
