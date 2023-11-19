# frozen_string_literal: true

lib = File.expand_path('lib', __dir__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)

require 'auth/version'

Gem::Specification.new do |spec|
  spec.name          = 'auth'
  spec.version       = Auth::VERSION
  spec.authors       = ['Alejandro Falkowski']
  spec.email         = ['alexrfalkowski@gmail.com']

  spec.summary       = 'A service for all your authn and authz needs.'
  spec.description   = spec.summary
  spec.homepage      = 'https://github.com/alexfalkowski/auth'
  spec.license       = 'Unlicense'
  spec.files         = Dir.chdir(File.expand_path(__dir__)) do
    `git ls-files -z`.split("\x0").reject { |f| f.match(%r{^(test|spec|features)/}) }
  end
  spec.bindir        = 'exe'
  spec.executables   = spec.files.grep(%r{^exe/}) { |f| File.basename(f) }
  spec.require_paths = ['lib']
  spec.required_ruby_version = ['>= 3.2.0', '< 4.0.0']
  spec.metadata['rubygems_mfa_required'] = 'true'
end
