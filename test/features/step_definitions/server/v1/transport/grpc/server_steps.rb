# frozen_string_literal: true

When('I request to generate a password with length {int} for gRPC') do |length|
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Auth::V1::GeneratePasswordRequest.new(length:)
  @response = Auth::V1.server_grpc.generate_password(request, { metadata: })
rescue StandardError => e
  @response = e
end

When('I request to generate a key with kind {string} with gRPC') do |kind|
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Auth::V1::GenerateKeyRequest.new(kind:)
  @response = Auth::V1.server_grpc.generate_key(request, { metadata: })
rescue StandardError => e
  @response = e
end

When('I request to get the public key with kind {string} with gRPC') do |kind|
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id }

  request = Auth::V1::GetPublicKeyRequest.new(kind:)
  @response = Auth::V1.server_grpc.get_public_key(request, { metadata: })
rescue StandardError => e
  @response = e
end

When('I request to generate an allowed access token with gRPC') do
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'authorization' => Auth::V1.basic_auth('valid_user')
  }

  request = Auth::V1::GenerateAccessTokenRequest.new
  @response = Auth::V1.server_grpc.generate_access_token(request, { metadata: })
rescue StandardError => e
  @response = e
end

When('I request to generate a disallowed access token with kind {string} with gRPC') do |kind|
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'authorization' => Auth::V1.basic_auth(kind)
  }

  request = Auth::V1::GenerateAccessTokenRequest.new
  @response = Auth::V1.server_grpc.generate_access_token(request, { metadata: })
rescue StandardError => e
  @response = e
end

When('I request to generate a allowed service token with kind {string} with gRPC') do |kind|
  @response = generate_service_token_with_grpc(kind, 'standort', Auth::V1.bearer_auth('valid_token'))
end

When('I request to generate a disallowed service token with kind {string} with gRPC') do |kind|
  @response = generate_service_token_with_grpc(kind, 'standort', Auth::V1.bearer_auth(kind))
end

When('I request to verify an allowed service token with kind {string} with gRPC') do |kind|
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'user-agent' => Auth.server_config['transport']['grpc']['user_agent'],
    'authorization' => Auth::V1.bearer_service_token('valid_token', @response.token.bearer)
  }

  request = Auth::V1::VerifyServiceTokenRequest.new(kind:, audience: 'standort', action: 'get-location')
  @response = Auth::V1.server_grpc.verify_service_token(request, { metadata: })
rescue StandardError => e
  @response = e
end

When('I request to verify a disallowed service token with gRPC:') do |table|
  rows = table.rows_hash
  response = generate_service_token_with_grpc(rows['token'], 'standort', Auth::V1.bearer_auth('valid_token'))
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'user-agent' => Auth.server_config['transport']['grpc']['user_agent'],
    'authorization' => Auth::V1.bearer_service_token(rows['issue'], response.token.bearer)
  }

  request = Auth::V1::VerifyServiceTokenRequest.new(kind: rows['token'], audience: 'standort', action: rows['issue'])
  @response = Auth::V1.server_grpc.verify_service_token(request, { metadata: })
rescue StandardError => e
  @response = e
end

Then('I should receive a valid password with length {int} for gRPC') do |length|
  length = 64 if length == 0

  expect(@response.meta.length).to be > 0
  expect(@response.password.plain).not_to include(':')
  expect(@response.password.plain.length).to eq(length)
  expect(@response.password['hash'].length).to be > 0
end

Then('I should receive an erroneous password with gRPC') do
  expect(@response).to be_a(GRPC::InvalidArgument)
end

Then('I should receive a valid key with kind {string} with gRPC') do |kind|
  expect(@response.meta.length).to be > 0

  pub = Base64.strict_decode64(@response.key['public'])
  pri = Base64.strict_decode64(@response.key['private'])

  expect(pub.length).to be > 0
  expect(pri.length).to be > 0

  kind = kind.strip

  if kind == 'rsa' || kind.empty?
    expect(OpenSSL::PKey::RSA.new(pub)).to be_public
    expect(OpenSSL::PKey::RSA.new(pri)).to be_private
  end

  expect(RbNaCl::Signatures::Ed25519::VerifyKey.new(pub).primitive).to eq(:ed25519) if kind == 'ed25519'
end

Then('I should receive a valid public key with kind {string} with gRPC') do |kind|
  expect(@response.meta.length).to be > 0
  expect(@response.key).to eq(Auth.server_config['key'][kind]['public'])
end

Then('I should receive a not found public key with gRPC') do
  expect(@response).to be_a(GRPC::NotFound)
end

Then('I should receive a valid access token with gRPC') do
  expect(@response.meta.length).to be > 0
  expect(@response.token.bearer.length).to be > 0
  expect(@response.token.password.plain).not_to include(':')
  expect(@response.token.password.plain.length).to eq(64)
  expect(@response.token.password['hash'].length).to be > 0
end

Then('I should receive a disallowed access token with gRPC') do
  expect(@response).to be_a(GRPC::Unauthenticated)
end

Then('I should receive a valid service token with kind {string} with gRPC') do |kind|
  expect(@response.meta.length).to be > 0
  expect(@response.token.bearer.length).to be > 0

  kind = kind.strip

  if kind == 'jwt' || kind.empty?
    decoded_token = Auth::V1.decode_jwt(@response.token.bearer)

    expect(decoded_token.length).to be > 0
    expect(decoded_token[0]['iss']).to eq(Auth.server_config['server']['v1']['issuer'])
    expect(decoded_token[0]['sub']).to eq('konfig')
    expect(decoded_token[0]['aud']).to eq(['standort'])
  end

  if kind == 'paseto'
    decoded_token = Auth::V1.decode_paseto(@response.token.bearer)

    expect(decoded_token.claims['iss']).to eq(Auth.server_config['server']['v1']['issuer'])
    expect(decoded_token.claims['sub']).to eq('konfig')
    expect(decoded_token.claims['aud']).to eq('standort')
  end
end

Then('I should receive a disallowed service token with gRPC') do
  expect(@response).to be_a(GRPC::Unauthenticated)
end

Then('I should have a valid service token with gRPC') do
  expect(@response).not_to be_a(GRPC::Unauthenticated)
end

Then('I should receive a disallowed verification of service token with gRPC') do
  expect(@response).to be_a(GRPC::Unauthenticated)
end

def generate_service_token_with_grpc(kind, audience, authorization)
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'user-agent' => Auth.server_config['transport']['grpc']['user_agent'],
    'authorization' => authorization
  }

  request = Auth::V1::GenerateServiceTokenRequest.new(kind:, audience:)
  Auth::V1.server_grpc.generate_service_token(request, { metadata: })
rescue StandardError => e
  @response = e
end
