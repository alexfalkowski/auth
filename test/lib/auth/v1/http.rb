# frozen_string_literal: true

module Auth
  module V1
    class HTTP < Nonnative::HTTPClient
      def generate_password(headers = {})
        default_headers = { content_type: :json, accept: :json }
        default_headers.merge!(headers)

        get('/v1/password/generate', headers, 10)
      end

      def generate_key(headers = {})
        default_headers = { content_type: :json, accept: :json }
        default_headers.merge!(headers)

        get('/v1/key/generate', headers, 10)
      end
    end
  end
end
