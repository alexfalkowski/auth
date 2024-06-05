FROM golang:1.22.4-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-s -w -X 'github.com/alexfalkowski/auth/cmd.Version=${version}'" -a -o auth main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /app/auth /auth
ENTRYPOINT ["/auth"]
