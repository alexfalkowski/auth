Feature: Server

  Server allows users to manage all your authn and authz needs.

  Scenario: Generate password with HTTP
    When I request to generate a password with HTTP
    Then I should receive a valid password with HTTP

  Scenario: Generate key with HTTP
    When I request to generate a key with HTTP
    Then I should receive a valid key with HTTP

  Scenario: Succesfully generate access token with HTTP
    When I request to generate an allowed access token with HTTP
    Then I should receive a valid access token with HTTP

  Scenario Outline: Unsuccesfully generate access token with HTTP
    When I request to generate a disallowed access token with kind "<kind>" with HTTP
    Then I should receive a disallowed access token with HTTP

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

  Scenario: Succesfully generate service token with HTTP
    When I request to generate an allowed service token with HTTP
    Then I should receive a valid service token with HTTP

  Scenario Outline: Unsuccesfully generate service token with HTTP
    When I request to generate a disallowed service token with kind "<kind>" with HTTP
    Then I should receive a disallowed service token with HTTP

    Examples:
      | kind            |
      | empty           |
      | not_supported   |
      | not_credentials |
      | invalid_token   |
