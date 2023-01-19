# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'grpc/health/v1/health_services_pb'
require 'jwt'
require 'paseto'

require 'auth/v1/service_services_pb'
require 'auth/v1/http'

module Auth
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:8080')
    end

    def server_config
      @server_config ||= YAML.load_file('.config/server.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:8080', :this_channel_is_insecure)
    end
  end

  module V1
    class << self
      def server_http
        @server_http ||= Auth::V1::HTTP.new('http://localhost:8080')
      end

      def server_grpc
        @server_grpc ||= Auth::V1::Service::Stub.new('localhost:8080', :this_channel_is_insecure)
      end

      def basic_auth(kind)
        lookup = {
          'empty' => '', 'not_supported' => 'Bob test', 'not_credentials' => 'Basic', 'invalid_encoding' => 'Basic test',
          'no_user' => "Basic #{Base64.strict_encode64(':MCZxL$Y5beypAWj<JQENft@P_DXVuh#,]02rq1Hwd69mFg(R|7ci&TlaoBU8k3s4')}",
          'missing_separator' => "Basic #{Base64.strict_encode64('su-1234')}", 'no_password' => "Basic #{Base64.strict_encode64('su-1234:')}",
          'invalid_user' => "Basic #{Base64.strict_encode64('no:MCZxL$Y5beypAWj<JQENft@P_DXVuh#,]02rq1Hwd69mFg(R|7ci&TlaoBU8k3s4')}",
          'invalid_password' => "Basic #{Base64.strict_encode64('su-1234:nooo')}",
          'valid_user' => "Basic #{Base64.strict_encode64('su-1234:MCZxL$Y5beypAWj<JQENft@P_DXVuh#,]02rq1Hwd69mFg(R|7ci&TlaoBU8k3s4')}"
        }

        lookup[kind]
      end

      def bearer_auth(kind)
        lookup = {
          'empty' => '', 'not_supported' => 'Bob test', 'not_credentials' => 'Bearer', 'invalid_token' => 'Bearer test',
          # rubocop:disable Layout/LineLength
          'valid_token' => 'Bearer PE7+1MdFLPkwb3BCTEYxesCsd96bb+3cfZbObbijWOsQ39HtvlQE9TuptDrurWOeD4gjhxzP1eGEF7A8CE9ddL1gUeulvWYJ16MKog+Rosbsxk3dze5j1yxhRiMzhH4bIe4MvCHdi2NPIRbT7qQcjBHr04KuHjO2qYya398kMuFq5Xezpl1uvv7idcpPQmgT6vKzFH14hVFR1R1S1ABCe7x2Fwxl6xDjaetoJ7vrpVwwbqOLl79L5U98QGyAHVE1kxEkursPQBGa7rb3s0LHAQCJOxS6daeV6Xkbd/y4rC2L+65xfB2FAMZtvg+bX+Tr6S4EyQyQt0GfJvX/8cNy6TN21UG73FNEk3TYIz46JAhRaGk2Atn4+AO35Ypz75ovZghn8snAkSSNjeZMnTFKf0uNIS6W/xULZcaZnxYKYwwjxsTwGsM/H97n2YqbaEOzG4fU69mJsz0KwA/2TE1aqhb5Hpf2GbJVxoN7AkNp6kRmmufMssZluXC5Xd+bKSrNsK2yzxAgPrV7X/xIfqeJZ653Vp6HevN7G41jQuXBXcgn4nnubnf+f/3CyvBpIcwm9OoVFFgb0C60nZFMmSex/yc6EwaRP+9EEESztCzX+W9kN0fNHCnq2+rYsVfc4wZVlQq0aRgtl1k9umGI1ikjGwWaJ8uwmxSOJHPyg6xFmDo='
          # rubocop:enable Layout/LineLength
        }

        lookup[kind]
      end

      def bearer_service_token(kind, token = '')
        lookup = {
          'empty' => '', 'not_supported' => 'Bob test', 'not_credentials' => 'Bearer', 'invalid_token' => 'Bearer test',
          'valid_token' => "Bearer #{token}"
        }

        lookup[kind]
      end

      def decode_jwt(token)
        k = Base64.strict_decode64(Auth.server_config['server']['v1']['key']['ed25519']['public'])
        key = RbNaCl::Signatures::Ed25519::VerifyKey.new(k)

        JWT.decode(token, key, true, { algorithm: 'EdDSA' })
      end

      def decode_paseto(token)
        k = Base64.strict_decode64(Auth.server_config['server']['v1']['key']['ed25519']['public'])
        key = RbNaCl::Signatures::Ed25519::VerifyKey.new(k)
        verifier = Paseto::V4::Public.new(key)

        verifier.decode(token)
      end
    end
  end
end
