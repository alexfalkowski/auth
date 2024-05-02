# frozen_string_literal: true

require 'base64'

ENV['SERVER_CERT_PEM'] = Base64.encode64(File.read('certs/cert.pem'))
ENV['SERVER_KEY_PEM'] = Base64.encode64(File.read('certs/key.pem'))

ENV['CLIENT_CERT_PEM'] = Base64.encode64(File.read('certs/client-cert.pem'))
ENV['CLIENT_KEY_PEM'] = Base64.encode64(File.read('certs/client-key.pem'))

require 'nonnative'
require 'auth'
