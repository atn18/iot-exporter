.PHONY: run
run:
	(source .env; go run cmd/main.go)

.PHONY: build
build:
	GOARM=7 GOOS=linux GOARCH=arm go build -o bin/iot_exporter_arm7 cmd/main.go
