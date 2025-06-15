build:
  GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bin/gpbus-linux-arm64 -trimpath
  GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o bin/gpbus-linux-arm -trimpath
