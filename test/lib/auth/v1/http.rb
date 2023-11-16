# frozen_string_literal: true

module Auth
  module V1
    class HTTP < Nonnative::HTTPClient
      def generate_password(length, opts = {})
        post('/v1/password/generate', { 'length' => length }.to_json, opts)
      end

      def generate_key(kind, opts = {})
        post('/v1/key/generate', { 'kind' => kind }.to_json, opts)
      end

      def get_public_key(kind, opts = {})
        get("/v1/key/public/#{kind}", opts)
      end

      def get_jwks(opts = {})
        get('/v1/.well-known/jwks.json', opts)
      end

      def generate_access_token(length, opts = {})
        post('/v1/access-token/generate', { 'length' => length }.to_json, opts)
      end

      def generate_oauth_token(client_id, client_secret, audience, grant_type, opts = {})
        payload = { 'client_id' => client_id, 'client_secret' => client_secret, 'audience' => audience, 'grant_type' => grant_type }.to_json
        post('/v1/oauth/token', payload, opts)
      end

      def generate_service_token(kind, audience, opts = {})
        post('/v1/service-token/generate', { 'kind' => kind, 'audience' => audience }.to_json, opts)
      end

      def verify_service_token(kind, audience, action, opts = {})
        get("/v1/service-token/verify/#{kind}/#{audience}/#{action}", opts)
      end
    end
  end
end
