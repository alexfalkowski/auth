FROM golang:1.20.0-bullseye AS build

ARG version=latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -ldflags="-X 'github.com/alexfalkowski/auth/cmd.Version=${version}'" -a -o auth main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/auth /auth
ENTRYPOINT ["/auth"]
