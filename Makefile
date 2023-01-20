.PHONY: vendor

include bin/build/make/service.mak

# Build release binary.
build:
	go build -race -ldflags="-X 'github.com/alexfalkowski/auth/cmd.Version=latest'" -mod vendor -o auth main.go

# Build test binary.
build-test:
	go test -race -ldflags="-X 'github.com/alexfalkowski/auth/cmd.Version=latest'" -mod vendor -c -tags features -covermode=atomic -o auth -coverpkg=./... github.com/alexfalkowski/auth

sanitize-coverage:
	bin/quality/go/cov

# Get the HTML coverage for go.
html-coverage: sanitize-coverage
	go tool cover -html test/reports/final.cov

# Get the func coverage for go.
func-coverage: sanitize-coverage
	go tool cover -func test/reports/final.cov

# Send coveralls data.
goveralls: sanitize-coverage
	goveralls -coverprofile=test/reports/final.cov -service=circle-ci -repotoken=TmQmSnygHAOrsUNorqpZE62DH0WYJy453

# Run go security checks.
go-sec:
	gosec -quiet -exclude-dir=test -exclude=G104 ./...

# Run security checks.
sec: go-sec

# Release to docker hub.
docker:
	bin/build/docker/push auth

# Start the environment.
start:
	bin/build/docker/env start

# Stop the environment.
stop:
	bin/build/docker/env stop
