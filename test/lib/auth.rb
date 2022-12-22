# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'grpc/health/v1/health_services_pb'

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
    end
  end
end
