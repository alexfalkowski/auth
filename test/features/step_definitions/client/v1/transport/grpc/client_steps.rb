# frozen_string_literal: true

When('I generate an access token') do
  env = {
    'CONFIG_FILE' => '.config/client.yaml'
  }
  cmd = Nonnative.go_executable('reports', '../auth', 'client', '--generate-access-token 0')
  pid = spawn(env, cmd, %i[out err] => ['reports/client.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('I generate a service token') do
  env = {
    'CONFIG_FILE' => '.config/client.yaml'
  }
  cmd = Nonnative.go_executable('reports', '../auth', 'client', '--generate-service-token jwt:standort')
  pid = spawn(env, cmd, %i[out err] => ['reports/client.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('I verify a service token') do
  env = {
    'CONFIG_FILE' => '.config/client.yaml'
  }
  cmd = Nonnative.go_executable('reports', '../auth', 'client', "--verify-service-token jwt:standort:get-location:#{@response.token.bearer}")
  pid = spawn(env, cmd, %i[out err] => ['reports/client.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have a generated access token') do
  expect(@status.exitstatus).to eq(0)
end

Then('I should have a generated service token') do
  expect(@status.exitstatus).to eq(0)
end

Then('I should have a verified service token') do
  expect(@status.exitstatus).to eq(0)
end
