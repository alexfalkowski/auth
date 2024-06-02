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
      @observability ||= Nonnative::Observability.new('https://localhost:11000')
    end

    def server_config
      @server_config ||= Nonnative.configurations('.config/server.yml')
    end

    def health_grpc
      @health_grpc ||= Grpc::Health::V1::Health::Stub.new('localhost:12000', Auth.creds_grpc, channel_args: Auth.user_agent)
    end

    def user_agent
      @user_agent ||= Nonnative::Header.grpc_user_agent('Auth-ruby-client/1.0 gRPC/1.0')
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
        @server_http ||= Auth::V1::HTTP.new('https://localhost:11000')
      end

      def server_grpc
        @server_grpc ||= Auth::V1::Service::Stub.new('localhost:12000', Auth.creds_grpc, channel_args: Auth.user_agent)
      end

      def basic_auth(kind)
        lookup = {
          'empty' => '', 'not_supported' => 'Bob test', 'not_credentials' => 'Basic', 'invalid_encoding' => 'Basic test',
          'no_user' => Nonnative::Header.auth_basic(':9ZMFeknVFo|1S-js5)r)HmYLvHTpq>wbo-=jNute@==q&%<Ms]Ff4vYWE[,7B3_#')[:authorization],
          'missing_separator' => Nonnative::Header.auth_basic('su-1234')[:authorization],
          'no_password' => Nonnative::Header.auth_basic('su-1234:')[:authorization],
          'invalid_user' => Nonnative::Header.auth_basic('no:9ZMFeknVFo|1S-js5)r)HmYLvHTpq>wbo-=jNute@==q&%<Ms]Ff4vYWE[,7B3_#')[:authorization],
          'invalid_password' => Nonnative::Header.auth_basic('su-1234:nooo')[:authorization],
          'valid_user' => Nonnative::Header.auth_basic('su-1234:9ZMFeknVFo|1S-js5)r)HmYLvHTpq>wbo-=jNute@==q&%<Ms]Ff4vYWE[,7B3_#')[:authorization]
        }

        lookup[kind]
      end

      def bearer_auth(kind)
        lookup = {
          'empty' => '', 'not_supported' => 'Bob test', 'not_credentials' => 'Bearer', 'invalid_token' => 'Bearer test',
          # rubocop:disable Layout/LineLength
          'valid_token' => Nonnative::Header.auth_bearer('bpmgvMpr1paMRxwGlmlEHjLAoPZdGKeyDAEVQODZ4tdnFe4T/VdAqDn+SyszZiyfEkSlUjSKjxjaPFFn6VTrHgD7UkeQFyXABt20UWSFJn4ktZVH+gDz5o7peLzsJOxT8toz+fNLMmApawEsd7ij9fBkNKArf6NJOTeyj6qdBKCh1bEdfED5egR9j31uoR4dVOo+20XtPFEhh5B6jLzjYAj8iV0S+9tCvCsm1hX+zDggW8Kr5DhjH+pUp6gVB5KY9zk9dixvI+xKb7jA+Pj4mkxATrguj3oROFTiXO9M4WJd7TOdMURpNqEYUyOQIXFicxf/JvTGhxWsocuk0xxvCgfaVw5pZENxXJUdiM/fCAQe1Tsm21DsoMXkXE2ccPjmJoRAqQKDyxys1nOspGhd9XNPpcpMytHWpUPPn0JL2H37UYimbGLC7VsxU6w6npZW7r9nG2PQPTEcxmUklDweB9b+ygP8ise/hhUCx9IEIDlkdgFHyNN+5Ii9XdoxAueeT4oWhBMRBK9NNlZ1B8xIXLP6uqBwn1v1iKtkr+OAYAAuE2++F96oJqOJuqpEDzrGSyLChdMfBXcdngMAbgoivA6U7vHFr7+Xe3w4LZyB2BkQCq6czx6TUqQC8M92lGfFa8kC0mRdNdHNW7Wa+/CQBDFVp29+S0QtxtRLhdK+iWY=')[:authorization]
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
          'missing_client_id' => { id: 'missing', secret: '/5?O?.%1xU[&@ba8ov=<Kzq~J=}YfpfAyf0=bV1MaGCPD!P&I(6@cBHl}wIM)W3<' },
          'missing_client_secret' => { id: 'e1602e185cba2a90d8bbcfc3f3c5530c', secret: 'missing' },
          'valid' => { id: 'e1602e185cba2a90d8bbcfc3f3c5530c', secret: '/5?O?.%1xU[&@ba8ov=<Kzq~J=}YfpfAyf0=bV1MaGCPD!P&I(6@cBHl}wIM)W3<' }
        }

        lookup[kind]
      end

      def decode_jwt(token)
        k = RbNaCl::Signatures::Ed25519::VerifyKey.new(key)

        JWT.decode(token, k, true, { algorithm: 'EdDSA' })
      end

      def decode_paseto(token)
        k = RbNaCl::Signatures::Ed25519::VerifyKey.new(key)
        verifier = Paseto::V4::Public.new(k)

        verifier.decode(token)
      end

      def key
        OpenSSL::PKey.read(File.read(Auth.server_config.crypto.ed25519.public)).raw_public_key
      end
    end
  end
end
