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
