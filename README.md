[![CircleCI](https://circleci.com/gh/alexfalkowski/auth.svg?style=svg)](https://circleci.com/gh/alexfalkowski/auth)
[![Coverage Status](https://coveralls.io/repos/github/alexfalkowski/auth/badge.svg?branch=master)](https://coveralls.io/github/alexfalkowski/auth?branch=master)

# Auth

A service for all your authn and authz needs.

## Background

A list of articles we have found useful to make this service:
- https://cheatsheetseries.owasp.org/cheatsheets/Microservices_security.html
- https://auth0.com/docs/secure/tokens/token-best-practices

## Generate Password

The service allows you to generate a secure password. The system uses:
- https://github.com/sethvargo/go-password
- https://pkg.go.dev/golang.org/x/crypto/bcrypt

The user can specify a length. The default is 32 and above.

We take recommendations from [NIST Password Guidelines](https://blog.netwrix.com/2022/11/14/nist-password-guidelines).

## Generate Key

The service allows you to generate a secure public and private keys.

### RSA

The service will create a key pair using 4096 bits. This is the default or when you pass the kind `rsa`.

### Ed25519

The service will generate a key pair of [Ed25519](https://ed25519.cr.yp.to/). This is achieved by passing the `ed25519` kind.

## Generate Access Token

Access tokens are generated with RSA key pair as created by the key service. The user can specify a length. The default is 32 and above.

Once this is generated it is configured as follows:

```yaml
server:
  v1:
    key:
      rsa:
        public: base64-public-key
        private: base64-private-key
```

```toml
[server.v1.key.rsa]
public = "base64-public-key"
private = "base64-private-key"
```

These keys should be stored and retrieved from an [application configuration system](https://github.com/alexfalkowski/konfig).

The service needs administrators to create these access tokens. This is configured as follows:

```yaml
server:
  v1:
    admins:
      - id: su-1234
        hash: bcrypt-password
```

```toml
[[server.v1.admins]]
id = "su-1234"
hash = "bcrypt-password"
```

Each admin has an id and a hash. The password and hash are generated by the password service. The user then sends `id:password` as [Basic Authentication](https://swagger.io/docs/specification/authentication/basic-authentication/). This will give you an encrypted access token. This token is a password that is encrypted with the public key. So you could always generate your own token if needed.

The password and token should be stored and retrieved from an [application configuration system](https://github.com/alexfalkowski/konfig). The hash is safe to just leave as is, not need to securely store it.

## Issuer

This is used to add the issuer to service tokens. This is configured as follows:

```yaml
server:
  v1:
    issuer: https://auth.falkowski.io
```

```toml
[server.v1]
issuer = "https://auth.falkowski.io"
```

## Get Public Key

The service allows you to get the public key by kind. The supported kinds are the same as generating a key.

## Generate Service Tokens

Service tokens are generated using Ed25519 key pair. Once this is generated it is configured as follows:

```yaml
server:
  v1:
    key:
      ed25519:
        public: base64-public-key
        private: base64-private-key
```

```toml
[server.v1.key.ed25519]
public = "base64-public-key"
private = "base64-private-key"
```

The system generates service tokens from the access tokens. This is configured as follows:

```yaml
server:
  v1:
    services:
      - id: service-name
        hash: bcrypt-password
        duration: 24h
```

```toml
[[server.v1.services]]
id = "service-name"
hash = "bcrypt-password"
duration = "24h"
```

Each service has an id, hash and duration. The access token is generated by the access token service. The user then sends the access token as [Bearer Authentication](https://swagger.io/docs/specification/authentication/bearer-authentication/). This will give you an encrypted service token with subject, audience and issuer that is valid for the duration.

### JWT

The system will by default generate [JWT tokens](https://jwt.io/), or if we specify kind to be `jwt`.

### Branca

The system will generate a [branca token](https://branca.io/) if the kind that is passed in is `branca`. This is configured as follows:

```yaml
server:
  v1:
    secret:
      branca: random-32-character-string
```

```toml
[server.v1.secret]
branca = "random-32-character-string"
```

The token is JSON encoded as follows:

```json
{"sub":"sub","aud":"aud","iss":"iss"}
```
### Paseto

The system will generate a [paseto token](https://github.com/paseto-standard/paseto-spec) if the kind that is passed in is `paseto`.

## Validate Service Tokens

The systems allows the validation of service tokens. The token is passed as [Bearer Authentication](https://swagger.io/docs/specification/authentication/bearer-authentication/) along with a kind and action.

### Casbin

[Casbin](https://github.com/casbin/casbin) is used to authorize the token. Currently it is configured using ACL, using the following:

```yaml
server:
  v1:
    casbin:
      model: |
        [request_definition]
        r = sub, obj, act

        [policy_definition]
        p = sub, obj, act

        [policy_effect]
        e = some(where (p.eft == allow))

        [matchers]
        m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
      policy: |
        p, subject, audience, action
```

```toml
[server.v1.casbin]
model = """
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
"""
policy = """
p, subject, audience, action
"""
```

#### Policy

Here is an explanation of the different terms:
- `Subject`: This is the current service that gets the token to send to the service.
- `Audience`: Is the destination service you will be sending the token to.
- `Action`: Is the action/method of the service you are trying to authorize against.

Check out [How it works?](https://github.com/casbin/casbin#how-it-works)

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
