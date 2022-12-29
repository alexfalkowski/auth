# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: auth/v1/service.proto for package 'Auth.V1'

require 'grpc'
require 'auth/v1/service_pb'

module Auth
  module V1
    module Service
      # Service for auth.
      class Service

        include ::GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'auth.v1.Service'

        # GeneratePassword that is secure.
        rpc :GeneratePassword, ::Auth::V1::GeneratePasswordRequest, ::Auth::V1::GeneratePasswordResponse
        # GenerateKey public and private key based on kind.
        rpc :GenerateKey, ::Auth::V1::GenerateKeyRequest, ::Auth::V1::GenerateKeyResponse
        # GetPublicKey from kind.
        rpc :GetPublicKey, ::Auth::V1::GetPublicKeyRequest, ::Auth::V1::GetPublicKeyResponse
        # GenerateAccessToken from RSA keys.
        rpc :GenerateAccessToken, ::Auth::V1::GenerateAccessTokenRequest, ::Auth::V1::GenerateAccessTokenResponse
        # GenerateServiceToken from Ed25519 keys.
        rpc :GenerateServiceToken, ::Auth::V1::GenerateServiceTokenRequest, ::Auth::V1::GenerateServiceTokenResponse
        # VerifyServiceToken based on kind.
        rpc :VerifyServiceToken, ::Auth::V1::VerifyServiceTokenRequest, ::Auth::V1::VerifyServiceTokenResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end
