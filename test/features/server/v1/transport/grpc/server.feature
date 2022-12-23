Feature: Server

  Server allows users to manage all your authn and authz needs.

  Scenario: Generate password with gRPC
    When I request to generate a password with gRPC
    Then I should receive a valid password with gRPC

  Scenario: Generate key with gRPC
    When I request to generate a key with gRPC
    Then I should receive a valid key with gRPC

  Scenario: Succesfully generate access token with gRPC
    When I request to generate an allowed access token with gRPC
    Then I should receive a valid access token with gRPC

  Scenario Outline: Unsuccesfully generate access token with gRPC
    When I request to generate a disallowed access token with kind "<kind>" with gRPC
    Then I should receive a disallowed access token with gRPC

    Examples:
      | kind              |
      | empty             |
      | not_supported     |
      | not_credentials   |
      | invalid_encoding  |
      | missing_separator |
      | no_user           |
      | no_password       |
      | invalid_user      |
      | invalid_password  |

  Scenario: Succesfully generate service token with gRPC
    When I request to generate an allowed service token with gRPC
    Then I should receive a valid service token with gRPC

  Scenario Outline: Unsuccesfully generate service token with gRPC
    When I request to generate a disallowed service token with kind "<kind>" with gRPC
    Then I should receive a disallowed service token with gRPC

    Examples:
      | kind            |
      | empty           |
      | not_supported   |
      | not_credentials |
      | invalid_token   |
