[![CircleCI](https://circleci.com/gh/alexfalkowski/auth.svg?style=svg)](https://circleci.com/gh/alexfalkowski/auth)
[![Coverage Status](https://coveralls.io/repos/github/alexfalkowski/auth/badge.svg?branch=master)](https://coveralls.io/github/alexfalkowski/auth?branch=master)

# Auth

A service for all your authn and authz needs.

## Password

The service allows you to generate a secure password. The system uses:
- https://github.com/sethvargo/go-password
- https://pkg.go.dev/golang.org/x/crypto/bcrypt

We take recommendations from [NIST Password Guidelines](https://blog.netwrix.com/2022/11/14/nist-password-guidelines).

## Key

The service allows you to generate a secure RSA public and private keys.

We take recommendations from [A Guide to RSA Encryption in Go](https://levelup.gitconnected.com/a-guide-to-rsa-encryption-in-go-1a18d827f35d).

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- Golang

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

The get yourself setup, please run the following:

```sh
make setup
```

### Binaries

To make sure everything compiles for the app, please run the following:

```sh
make build-test
```

### Features

To run all the features, please run the following:

```sh
make features
```

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
