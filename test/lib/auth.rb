# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'grpc/health/v1/health_services_pb'
require 'jwt'

require 'auth/v1/service_services_pb'
require 'auth/v1/http'

module Auth
  class << self
    def observability
      @observability ||= Nonnative::Observability.new('http://localhost:8080')
    end

    def server_config
      @server_config ||= YAML.load_file('.config/server.config.yml')
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
          'valid_token' => 'Bearer LhO4FgGdqBpsXwnYGX565aLYLAT8SZn2mH/QlF0oOPvDKAognUtx4FyoQNx1UF/f3v3D7A17ZOlO4iVtRtlBblAh1Tys7TofVJ33B7ApzYPbKEh+bE7fx1+1HwlgZ61bkdHp1xKW8p14cydigVMtvXIwRl0qLJQSOdMLVy2wkq2Xu7C/GHGINl8G35v2eR7BVp2eArfj8AkTFNSI3eJ6mjMJpu73gKALSkcMb2V+9T0PwZvH981wTV7kPz2UWHLVtzmghh3n1nMNFPkCVWAbaISlsZ0NYYFj3gglDPPPIw398fDJN5Kf677H+nTBBIQDSojZQ7tcPXUThdgm6XwB/LLDbF/4zM7l+kkLy5qaYyA9eaOyDKSL+Xcid9y2GZjEdMqhMDRpANu6JjpWjPAWQ41LiWEbE04AsmsgXrfHouRRsXDR3O6YGb680OC2B2ZNM97atqtPvO4l54QDYVwWXR0qIt8pvw55bE96fNFgSlMZ9k/58xXnnrBHxs8/wEgVlC7NwfACemtVMlS2nBx4zaY3cgOMGEmC1ycbPgxo+IJ4WmzQEdthZ1Vk7WGkr+SnHHJcIINlsQTHkroGsoOEV/VxhvDYRVAretwGHwcSL9Af5L/Eoml734biGL4jzb6e85Znqp4/0XqgjMQSrYPlhOFigrfu1Xz6nn4SlIkM0lY='
          # rubocop:enable Layout/LineLength
        }

        lookup[kind]
      end

      def decode_token(token)
        key = OpenSSL::PKey::RSA.new(Auth.server_config['server']['v1']['key']['public'])

        JWT.decode(token, key, true, { algorithm: 'RS512' })
      end
    end
  end
end
