.PHONY: vendor

include bin/build/make/service.mak

# Build release binary.
build:
	go build -race -ldflags="-X 'github.com/alexfalkowski/auth/cmd.Version=latest'" -mod vendor -o auth main.go

# Build test binary.
build-test:
	go test -race -ldflags="-X 'github.com/alexfalkowski/auth/cmd.Version=latest'" -mod vendor -c -tags features -covermode=atomic -o auth -coverpkg=./... github.com/alexfalkowski/auth

# Release to docker hub.
docker:
	bin/build/docker/push auth
