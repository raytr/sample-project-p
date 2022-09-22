build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o app/backend-challenge ./cmd/*.go
	#cp conf.yml /app
test:
	go test ./...