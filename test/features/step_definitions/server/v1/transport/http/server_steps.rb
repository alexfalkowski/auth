# frozen_string_literal: true

When('I request to generate a password with HTTP') do
  headers = { request_id: SecureRandom.uuid, user_agent: Auth.server_config['transport']['grpc']['user_agent'] }

  @response = Auth::V1.server_http.generate_password(headers)
end

When('I request to generate a key with HTTP') do
  headers = { request_id: SecureRandom.uuid, user_agent: Auth.server_config['transport']['grpc']['user_agent'] }

  @response = Auth::V1.server_http.generate_key(headers)
end

Then('I should receive a valid password with HTTP') do
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['password']['plain'].length).to eq(64)
  expect(resp['password']['hash'].length).to be > 0
end

Then('I should receive a valid key with HTTP') do
  expect(@response.code).to eq(200)

  resp = JSON.parse(@response.body)

  expect(resp['key']['public'].length).to be > 0
  expect(resp['key']['private'].length).to be > 0

  expect(OpenSSL::PKey::RSA.new(resp['key']['public'])).to be_public
  expect(OpenSSL::PKey::RSA.new(resp['key']['private'])).to be_private
end
