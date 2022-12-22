# frozen_string_literal: true

module Auth
  module V1
    class HTTP < Nonnative::HTTPClient
      def generate_password(headers = {})
        headers.merge!(default_headers)

        post('/v1/password/generate', {}, headers, 10)
      end

      def generate_key(headers = {})
        headers.merge!(default_headers)

        post('/v1/key/generate', {}, headers, 10)
      end

      def generate_access_token(headers = {})
        headers.merge!(default_headers)

        post('/v1/access-token/generate', {}, headers, 10)
      end

      def generate_service_token(headers = {})
        headers.merge!(default_headers)

        post('/v1/service-token/generate', {}, headers, 10)
      end

      private

      def default_headers
        { content_type: :json, accept: :json }
      end
    end
  end
end