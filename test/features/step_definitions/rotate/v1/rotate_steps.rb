# frozen_string_literal: true

Given('I have an existing configuration') do
  FileUtils.rm_f('reports/server.yml')
end

When('I rotate an existing configuration') do
  env = {
    'CONFIG_FILE' => '.config/server.yml',
    'ROTATE_CONFIG_FILE' => 'reports/server.yml'
  }
  cmd = Nonnative.go_executable(%w[cover], 'reports', '../auth', 'rotate')
  pid = spawn(env, cmd, %i[out err] => ['reports/rotate.log', 'a'])

  _, @status = Process.waitpid2(pid)
end

Then('I should have a working rotated configuration') do
  expect(@status.exitstatus).to eq(0)
  expect(File.exist?('reports/server.yml')).to be true

  src = Nonnative.configurations('.config/server.yml')
  dest = Nonnative.configurations('reports/server.yml')

  expect(src.key.rsa.public).to_not eq(dest.key.rsa.public)
  expect(src.key.rsa.private).to_not eq(dest.key.rsa.private)
  expect(src.key.ed25519.public).to_not eq(dest.key.ed25519.public)
  expect(src.key.ed25519.private).to_not eq(dest.key.ed25519.private)
end
