Feature: Server

  Server allows users to manage all your authn and authz needs.

  Scenario: Generate password with HTTP
    When I request to generate a password with HTTP
    Then I should receive a valid password with HTTP
