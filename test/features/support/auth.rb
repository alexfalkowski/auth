# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'grpc/health/v1/health_services_pb'
require 'rbnacl'
require 'jwt'
require 'paseto'

require 'auth/v1/service_services_pb'
require 'auth/v1/http'

module Auth
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('https://localhost:8080')
    end

    def server_config
      @server_config ||= Nonnative.configurations('.config/server.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:9090', Auth.creds_grpc, channel_args: Auth.user_agent)
    end

    def user_agent
      @user_agent ||= Nonnative::Header.grpc_user_agent(server_config.transport.grpc.user_agent)
    end

    def creds_grpc
      @creds_grpc ||= GRPC::Core::ChannelCredentials.new(File.read(ENV.fetch('AUTH_ROOT_CA', nil)), File.read('certs/client-key.pem'),
                                                         File.read('certs/client-cert.pem'))
    end

    def creds_http
      @creds_http ||= { ssl_client_cert: OpenSSL::X509::Certificate.new(File.read('certs/client-cert.pem')),
                        ssl_client_key: OpenSSL::PKey::RSA.new(File.read('certs/client-key.pem')),
                        ssl_ca_file: ENV.fetch('AUTH_ROOT_CA', nil), verify_ssl: OpenSSL::SSL::VERIFY_PEER }
    end
  end

  module V1
    class << self
      def server_http
        @server_http ||= Auth::V1::HTTP.new('https://localhost:8080')
      end

      def server_grpc
        @server_grpc ||= Auth::V1::Service::Stub.new('localhost:9090', Auth.creds_grpc, channel_args: Auth.user_agent)
      end

      def basic_auth(kind)
        lookup = {
          'empty' => '', 'not_supported' => 'Bob test', 'not_credentials' => 'Basic', 'invalid_encoding' => 'Basic test',
          'no_user' => Nonnative::Header.auth_basic(':o{6wJTESvAy/Z8alkC0]bjsQd2*~zf}DGp=)eLtXg#nV51IRrP974Y3mKBxq%OWi')[:authorization],
          'missing_separator' => Nonnative::Header.auth_basic('su-1234')[:authorization],
          'no_password' => Nonnative::Header.auth_basic('su-1234:')[:authorization],
          'invalid_user' => Nonnative::Header.auth_basic('no:o{6wJTESvAy/Z8alkC0]bjsQd2*~zf}DGp=)eLtXg#nV51IRrP974Y3mKBxq%OWi')[:authorization],
          'invalid_password' => Nonnative::Header.auth_basic('su-1234:nooo')[:authorization],
          'valid_user' => Nonnative::Header.auth_basic('su-1234:o{6wJTESvAy/Z8alkC0]bjsQd2*~zf}DGp=)eLtXg#nV51IRrP974Y3mKBxq%OWi')[:authorization]
        }

        lookup[kind]
      end

      def bearer_auth(kind)
        lookup = {
          'empty' => '', 'not_supported' => 'Bob test', 'not_credentials' => 'Bearer', 'invalid_token' => 'Bearer test',
          # rubocop:disable Layout/LineLength
          'valid_token' => Nonnative::Header.auth_bearer('jiptGfppR5U9uVHSDPDMlnoOfTysUfruuEngvJDno/LgDt3YiYT+/WMBRr/V91dsXCcSYfut6pgUWh5evBrD4mXGoCc8zrlIM+fllgcGPNuGSYk21Q9C3JKaBm617uOV3nIt1kJT/VKrJGTfqOeiRcPg3+11Urj4/R8NfywZChAQtekoZKjGeB4S9g+oTNDxOkSWfYyY3fiGhxQf4R4KpXaJFP4jEWNoTNslwTMSbFAdF+j2+Ne3rA1jSUJvkUZxjP0snUkbSDLB/BirkZBUDHOnVPOpdLZckGfDpU3Ne/+ZJZAOAECkDBuzMv+oOptRgK/ASxUEs5RWC9AzOjFJ2VdoUJ3yqpqu0+3rn3qAr8YCUwUDYLQS57TH9ESaw/NLW9Qjqn+ku8y9zqmfqRRx4/lxkmV/gTxJuZHE2AJdirkzjp6QKTu3DTCn3qSqlWVnmE6Zo8sH3YCBPzojf57DeZLjP2jgVzsIqTtRT7d/qXH7+7B1f4bB9MhUhbIMzHDq97Zomx+JDrobK5bAtNHYZV4cJukWwlzZDEYLNGUbuuKXiEr8AuXilwNhi4vUDYzq/r/URoj23jMRJH/cb2Kq2KCHPs1u362TmAGRN1/ybQZduwSQHDmxaWTIc0lA0ZBLWvbFqLnNPF5ly3swn5EDBdX+Ze3sMthzWSPcgTs0tog=')[:authorization]
          # rubocop:enable Layout/LineLength
        }

        lookup[kind]
      end

      def bearer_service_token(kind, token = '')
        lookup = {
          'empty' => '', 'not_supported' => 'Bob test', 'not_credentials' => 'Bearer', 'invalid_token' => 'Bearer test',
          'valid_token' => Nonnative::Header.auth_bearer(token)[:authorization]
        }

        lookup[kind]
      end

      def client(kind)
        lookup = {
          'missing_client_id' => { id: 'missing', secret: '$VzL_-HdP3Y7oE(64?jf@Irau|BJ!<ei0)51WcDhnQkZA2NtXMT8yObGUsgvKRl9' },
          'missing_client_secret' => { id: 'e1602e185cba2a90d8bbcfc3f3c5530c', secret: 'missing' },
          'valid' => { id: 'e1602e185cba2a90d8bbcfc3f3c5530c', secret: '$VzL_-HdP3Y7oE(64?jf@Irau|BJ!<ei0)51WcDhnQkZA2NtXMT8yObGUsgvKRl9' }
        }

        lookup[kind]
      end

      def decode_jwt(token)
        k = Base64.strict_decode64(Auth.server_config.key.ed25519.public)
        key = RbNaCl::Signatures::Ed25519::VerifyKey.new(k)

        JWT.decode(token, key, true, { algorithm: 'EdDSA' })
      end

      def decode_paseto(token)
        k = Base64.strict_decode64(Auth.server_config.key.ed25519.public)
        key = RbNaCl::Signatures::Ed25519::VerifyKey.new(k)
        verifier = Paseto::V4::Public.new(key)

        verifier.decode(token)
      end
    end
  end
end
