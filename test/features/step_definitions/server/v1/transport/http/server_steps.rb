# frozen_string_literal: true

When('I request to generate a password with length {int} for HTTP') do |length|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Auth::V1.server_http.generate_password(length, opts.merge(Auth.creds_http))
end

When('I request to generate a key with kind {string} with HTTP') do |kind|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 30, open_timeout: 30
  }

  @response = Auth::V1.server_http.generate_key(kind, opts.merge(Auth.creds_http))
end

When('I request to get the public key with kind {string} with HTTP') do |kind|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Auth::V1.server_http.get_public_key(kind, opts.merge(Auth.creds_http))
end

When('I request to generate an allowed access token with HTTP') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json,
      authorization: Auth::V1.basic_auth('valid_user')
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Auth::V1.server_http.generate_access_token(0, opts.merge(Auth.creds_http))
end

When('I request to generate a disallowed access token with kind {string} with HTTP') do |kind|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json,
      authorization: Auth::V1.basic_auth(kind)
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Auth::V1.server_http.generate_access_token(0, opts.merge(Auth.creds_http))
end

When('I request to generate an allowed service token with kind {string} with HTTP') do |kind|
  @response = generate_service_token_with_http(kind, 'standort', Auth::V1.bearer_auth('valid_token'))
end

When('I request to generate a disallowed service token with kind {string} with HTTP') do |kind|
  @response = generate_service_token_with_http(kind, kind,  Auth::V1.bearer_auth(kind))
end

When('I request to verify an allowed service token with kind {string} with HTTP') do |kind|
  resp = JSON.parse(@response.body)
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json,
      authorization: Auth::V1.bearer_service_token('valid_token', resp['token']['bearer'])
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Auth::V1.server_http.verify_service_token(kind, 'standort', 'get-location', opts.merge(Auth.creds_http))
end

When('I request to verify a disallowed service token with HTTP:') do |table|
  rows = table.rows_hash
  resp = JSON.parse(generate_service_token_with_http(rows['token'], 'standort', Auth::V1.bearer_auth('valid_token')).body)
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json,
      authorization: Auth::V1.bearer_service_token(rows['issue'], resp['token']['bearer'])
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Auth::V1.server_http.verify_service_token(rows['token'], 'standort', rows['issue'], opts.merge(Auth.creds_http))
end

When('I request to generate an allowed oauth token with HTTP') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }
  client = Auth::V1.client('valid')
  audience = 'standort'
  grant_type = 'client_credentials'

  @response = Auth::V1.server_http.generate_oauth_token(client[:id], client[:secret], audience, grant_type, opts.merge(Auth.creds_http))
end

When('I request to get the jwks with HTTP') do
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }

  @response = Auth::V1.server_http.get_jwks(opts.merge(Auth.creds_http))
end

When('I request to generate a disallowed oauth token of kind {string} with HTTP') do |kind|
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json
    },
    read_timeout: 10, open_timeout: 10
  }
  client = Auth::V1.client(kind)
  audience = 'standort'
  grant_type = 'client_credentials'

  @response = Auth::V1.server_http.generate_oauth_token(client[:id], client[:secret], audience, grant_type, opts.merge(Auth.creds_http))
end

Then('I should receive a valid password with length {int} for HTTP') do |length|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  pass = resp['password']['plain']
  length = 64 if length == 0

  expect(resp['meta'].length).to be > 0
  expect(pass).not_to include(':')
  expect(pass.length).to eq(length)
  expect(resp['password']['hash'].length).to be > 0
end

Then('I should receive an erroneous password with HTTP') do
  expect(@response.code).to eq(400)
end

Then('I should receive a valid key with kind {string} with HTTP') do |kind|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['meta'].length).to be > 0

  pub = Base64.strict_decode64(resp['key']['public'])
  pri = Base64.strict_decode64(resp['key']['private'])

  expect(pub.length).to be > 0
  expect(pri.length).to be > 0

  kind = kind.strip

  if kind == 'rsa' || kind.empty?
    expect(OpenSSL::PKey::RSA.new(pub)).to be_public
    expect(OpenSSL::PKey::RSA.new(pri)).to be_private
  end

  expect(RbNaCl::Signatures::Ed25519::VerifyKey.new(pub).primitive).to eq(:ed25519) if kind == 'ed25519'
end

Then('I should receive a valid public key with kind {string} with HTTP') do |kind|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['meta'].length).to be > 0
  expect(resp['key']).to eq(Auth.server_config.crypto.send(kind).public)
end

Then('I should receive a not found public key with HTTP') do
  expect(@response.code).to eq(404)
end

Then('I should receive a valid access token with HTTP') do
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  pass = resp['token']['password']['plain']

  expect(resp['meta'].length).to be > 0
  expect(resp['token']['bearer'].length).to be > 0
  expect(pass).not_to include(':')
  expect(pass.length).to eq(64)
  expect(resp['token']['password']['hash'].length).to be > 0
end

Then('I should receive a disallowed access token with HTTP') do
  expect(@response.code).to eq(401)
end

Then('I should receive a valid service token with kind {string} with HTTP') do |kind|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  kind = kind.strip

  expect(resp['meta'].length).to be > 0

  if kind == 'jwt' || kind.empty?
    decoded_token = Auth::V1.decode_jwt(resp['token']['bearer'])

    expect(decoded_token.length).to be > 0
    expect(decoded_token[0]['iss']).to eq(Auth.server_config.server.v1.issuer)
    expect(decoded_token[0]['sub']).to eq('konfig')
    expect(decoded_token[0]['aud']).to eq(['standort'])
  end

  if kind == 'paseto'
    decoded_token = Auth::V1.decode_paseto(resp['token']['bearer'])

    expect(decoded_token.claims['iss']).to eq(Auth.server_config.server.v1.issuer)
    expect(decoded_token.claims['sub']).to eq('konfig')
    expect(decoded_token.claims['aud']).to eq('standort')
  end
end

Then('I should receive a disallowed service token with HTTP') do
  expect(@response.code).to eq(401)
end

Then('I should have a valid service token with HTTP') do
  expect(@response.code).to eq(200)
end

Then('I should receive a disallowed verification of service token with HTTP') do
  expect(@response.code).to eq(401)
end

Then('I should receive a valid oauth token with HTTP') do
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['meta'].length).to be > 0
  expect(resp['access_token'].length).to be > 0
  expect(resp['token_type']).to eq('Bearer')
end

Then('I should receive a valid jwks with HTTP') do
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['meta'].length).to be > 0
  expect(resp['keys'].length).to be > 0

  key = resp['keys'][0]
  expect(key['kid'].length).to be > 0
  expect(key['kty']).to eq('EC')
  expect(key['use']).to eq('sig')
  expect(key['x5c'].length).to be > 0
end

Then('I should receive a disallowed oauth token with HTTP') do
  expect(@response.code).to eq(401)
end

def generate_service_token_with_http(kind, audience, authorization)
  opts = {
    headers: {
      request_id: SecureRandom.uuid, user_agent: 'Auth-ruby-client/1.0 HTTP/1.0',
      content_type: :json, accept: :json,
      authorization:
    },
    read_timeout: 10, open_timeout: 10
  }

  Auth::V1.server_http.generate_service_token(kind, audience, opts.merge(Auth.creds_http))
end
