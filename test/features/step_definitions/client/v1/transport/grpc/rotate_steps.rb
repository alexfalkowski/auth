# frozen_string_literal: true

Before('@rotate') do
  FileUtils.rm_f('reports/server.yml')
end

When('I rotate an all of the configuration') do
  env = {
    'AUTH_CONFIG_FILE' => '.config/server.yml',
    'AUTH_APP_CONFIG_FILE' => 'reports/server.yml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../auth', 'rotate', '--admins', '--services')
  pid = spawn(env, cmd, %i[out err] => ['reports/all_rotate.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('I rotate an admins of the configuration') do
  env = {
    'AUTH_CONFIG_FILE' => '.config/server.yml',
    'AUTH_APP_CONFIG_FILE' => 'reports/server.yml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../auth', 'rotate', '--admins')
  pid = spawn(env, cmd, %i[out err] => ['reports/admins_rotate.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

When('I rotate an services of the configuration') do
  env = {
    'AUTH_CONFIG_FILE' => '.config/server.yml',
    'AUTH_APP_CONFIG_FILE' => 'reports/server.yml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../auth', 'rotate', '--services')
  pid = spawn(env, cmd, %i[out err] => ['reports/services_rotate.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have a complete rotated configuration') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/server.yml')).to be true

  src = Nonnative.configurations('.config/server.yml')
  dest = Nonnative.configurations('reports/server.yml')

  expect(src.crypto.rsa.public).to eq(dest.crypto.rsa.public)
  expect(src.crypto.rsa.private).to eq(dest.crypto.rsa.private)
  expect(src.crypto.ed25519.public).to eq(dest.crypto.ed25519.public)
  expect(src.crypto.ed25519.private).to eq(dest.crypto.ed25519.private)
end

Then('I should have the admins rotated in the configuration') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/server.yml')).to be true

  src = Nonnative.configurations('.config/server.yml')
  dest = Nonnative.configurations('reports/server.yml')

  expect(src.crypto.rsa.public).to eq(dest.crypto.rsa.public)
  expect(src.crypto.rsa.private).to eq(dest.crypto.rsa.private)
  expect(src.crypto.ed25519.public).to eq(dest.crypto.ed25519.public)
  expect(src.crypto.ed25519.private).to eq(dest.crypto.ed25519.private)
end

Then('I should have the services rotated in the configuration') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/server.yml')).to be true

  src = Nonnative.configurations('.config/server.yml')
  dest = Nonnative.configurations('reports/server.yml')

  expect(src.crypto.rsa.public).to eq(dest.crypto.rsa.public)
  expect(src.crypto.rsa.private).to eq(dest.crypto.rsa.private)
  expect(src.crypto.ed25519.public).to eq(dest.crypto.ed25519.public)
  expect(src.crypto.ed25519.private).to eq(dest.crypto.ed25519.private)
end
