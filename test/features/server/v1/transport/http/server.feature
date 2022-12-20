Feature: Server

  Server allows users to manage all your authn and authz needs.

  Scenario: Generate password with HTTP
    When I request to generate a password with HTTP
    Then I should receive a valid password with HTTP

  Scenario: Generate key with HTTP
    When I request to generate a key with HTTP
    Then I should receive a valid key with HTTP
