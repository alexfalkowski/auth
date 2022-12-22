# frozen_string_literal: true

When('I request to generate a password with gRPC') do
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id, 'ua' => Auth.server_config['transport']['grpc']['user_agent'] }

  request = Auth::V1::GeneratePasswordRequest.new
  @response = Auth::V1.server_grpc.generate_password(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

When('I request to generate a key with gRPC') do
  @request_id = SecureRandom.uuid
  metadata = { 'request-id' => @request_id, 'ua' => Auth.server_config['transport']['grpc']['user_agent'] }

  request = Auth::V1::GenerateKeyRequest.new
  @response = Auth::V1.server_grpc.generate_key(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

When('I request to generate an allowed access token with gRPC') do
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'ua' => Auth.server_config['transport']['grpc']['user_agent'],
    'authorization' => Auth::V1.basic_auth('valid_user')
  }

  request = Auth::V1::GenerateAccessTokenRequest.new
  @response = Auth::V1.server_grpc.generate_access_token(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

When('I request to generate a disallowed access token with kind {string} with gRPC') do |kind|
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'ua' => Auth.server_config['transport']['grpc']['user_agent'],
    'authorization' => Auth::V1.basic_auth(kind)
  }

  request = Auth::V1::GenerateAccessTokenRequest.new
  @response = Auth::V1.server_grpc.generate_access_token(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

When('I request to generate an allowed service token with gRPC') do
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'ua' => Auth.server_config['transport']['grpc']['user_agent'],
    'authorization' => Auth::V1.bearer_auth('valid_token')
  }

  request = Auth::V1::GenerateServiceTokenRequest.new
  @response = Auth::V1.server_grpc.generate_service_token(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

When('I request to generate a disallowed service token with kind {string} with gRPC') do |kind|
  @request_id = SecureRandom.uuid
  metadata = {
    'request-id' => @request_id,
    'ua' => Auth.server_config['transport']['grpc']['user_agent'],
    'authorization' => Auth::V1.bearer_auth(kind)
  }

  request = Auth::V1::GenerateServiceTokenRequest.new
  @response = Auth::V1.server_grpc.generate_service_token(request, { metadata: metadata })
rescue StandardError => e
  @response = e
end

Then('I should receive a valid password with gRPC') do
  expect(@response.password.plain.length).to eq(64)
  expect(@response.password['hash'].length).to be > 0
end

Then('I should receive a valid key with gRPC') do
  expect(@response.key['public'].length).to be > 0
  expect(@response.key['private'].length).to be > 0
  expect(OpenSSL::PKey::RSA.new(@response.key['public'])).to be_public
  expect(OpenSSL::PKey::RSA.new(@response.key['private'])).to be_private
end

Then('I should receive a valid access token with gRPC') do
  expect(@response.token.bearer.length).to be > 0
  expect(@response.token.password.plain.length).to eq(64)
  expect(@response.token.password['hash'].length).to be > 0
end

Then('I should receive a disallowed access token with gRPC') do
  expect(@response).to be_a(GRPC::Unauthenticated)
end

Then('I should receive a valid service token with gRPC') do
  expect(@response.token.bearer.length).to be > 0

  decoded_token = Auth::V1.decode_token(@response.token.bearer)

  expect(decoded_token.length).to be > 0
  expect(decoded_token[0]['iss']).to eq(Auth.server_config['server']['v1']['issuer'])
end

Then('I should receive a disallowed service token with gRPC') do
  expect(@response).to be_a(GRPC::Unauthenticated)
end
