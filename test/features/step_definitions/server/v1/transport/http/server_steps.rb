# frozen_string_literal: true

When('I request to generate a password with HTTP') do
  headers = { request_id: SecureRandom.uuid, user_agent: Auth.server_config['transport']['grpc']['user_agent'] }

  @response = Auth::V1.server_http.generate_password(headers)
end

When('I request to generate a key with HTTP') do
  headers = { request_id: SecureRandom.uuid, user_agent: Auth.server_config['transport']['grpc']['user_agent'] }

  @response = Auth::V1.server_http.generate_key(headers)
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
