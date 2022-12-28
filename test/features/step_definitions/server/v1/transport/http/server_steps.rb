# frozen_string_literal: true

When('I request to generate a password with HTTP') do
  headers = { request_id: SecureRandom.uuid, user_agent: Auth.server_config['transport']['grpc']['user_agent'] }

  @response = Auth::V1.server_http.generate_password(headers)
end

When('I request to generate a key with kind {string} with HTTP') do |kind|
  headers = { request_id: SecureRandom.uuid, user_agent: Auth.server_config['transport']['grpc']['user_agent'] }

  @response = Auth::V1.server_http.generate_key(kind, headers)
end

When('I request to get the public key with kind {string} with HTTP') do |kind|
  headers = { request_id: SecureRandom.uuid, user_agent: Auth.server_config['transport']['grpc']['user_agent'] }

  @response = Auth::V1.server_http.get_public_key(kind, headers)
end

When('I request to generate an allowed access token with HTTP') do
  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Auth.server_config['transport']['grpc']['user_agent'],
    authorization: Auth::V1.basic_auth('valid_user')
  }

  @response = Auth::V1.server_http.generate_access_token(headers)
end

When('I request to generate a disallowed access token with kind {string} with HTTP') do |kind|
  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Auth.server_config['transport']['grpc']['user_agent'],
    authorization: Auth::V1.basic_auth(kind)
  }

  @response = Auth::V1.server_http.generate_access_token(headers)
end

When('I request to generate a allowed service token with kind {string} with HTTP') do |kind|
  @response = generate_service_token_with_http(kind, 'standort', Auth::V1.bearer_auth('valid_token'))
end

When('I request to generate a disallowed service token with kind {string} with HTTP') do |kind|
  @response = generate_service_token_with_http(kind, kind,  Auth::V1.bearer_auth(kind))
end

When('I request to verify an allowed service token with kind {string} with HTTP') do |kind|
  resp = JSON.parse(@response.body)
  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Auth.server_config['transport']['grpc']['user_agent'],
    authorization: Auth::V1.bearer_service_token('valid_token', resp['token']['bearer'])
  }

  @response = Auth::V1.server_http.verify_service_token(kind, 'get-location', headers)
end

When('I request to verify a disallowed service token with HTTP:') do |table|
  rows = table.rows_hash
  resp = JSON.parse(generate_service_token_with_http(rows['token'], 'standort', Auth::V1.bearer_auth('valid_token')).body)

  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Auth.server_config['transport']['grpc']['user_agent'],
    authorization: Auth::V1.bearer_service_token(rows['issue'], resp['token']['bearer'])
  }

  @response = Auth::V1.server_http.verify_service_token(rows['token'], rows['issue'], headers)
end

Then('I should receive a valid password with HTTP') do
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['password']['plain'].length).to eq(64)
  expect(resp['password']['hash'].length).to be > 0
end

Then('I should receive a valid key with kind {string} with HTTP') do |kind|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
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

  expect(resp['key']).to eq(Auth.server_config['server']['v1']['key'][kind]['public'])
end

Then('I should receive a not found public key with HTTP') do
  expect(@response.code).to eq(404)
end

Then('I should receive a valid access token with HTTP') do
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['token']['bearer'].length).to be > 0
  expect(resp['token']['password']['plain'].length).to eq(64)
  expect(resp['token']['password']['hash'].length).to be > 0
end

Then('I should receive a disallowed access token with HTTP') do
  expect(@response.code).to eq(401)
end

Then('I should receive a valid service token with kind {string} with HTTP') do |kind|
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)
  kind = kind.strip

  if kind == 'jwt' || kind.empty?
    decoded_token = Auth::V1.decode_jwt(resp['token']['bearer'])

    expect(decoded_token.length).to be > 0
    expect(decoded_token[0]['iss']).to eq(Auth.server_config['server']['v1']['issuer'])
    expect(decoded_token[0]['sub']).to eq('konfig')
    expect(decoded_token[0]['aud']).to eq(['standort'])
  end

  if kind == 'branca'
    decoded_token = Auth::V1.decode_branca(resp['token']['bearer'])
    message = JSON.parse(decoded_token.message)

    expect(message).to eq({ 'aud' => 'standort', 'iss' => Auth.server_config['server']['v1']['issuer'], 'sub' => 'konfig' })
  end

  if kind == 'paseto'
    decoded_token = Auth::V1.decode_paseto(resp['token']['bearer'])

    expect(decoded_token.claims['iss']).to eq(Auth.server_config['server']['v1']['issuer'])
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

def generate_service_token_with_http(kind, audience, authorization)
  headers = {
    request_id: SecureRandom.uuid,
    user_agent: Auth.server_config['transport']['grpc']['user_agent'],
    authorization: authorization
  }

  Auth::V1.server_http.generate_service_token(kind, audience, headers)
end
