Feature: Server

  Server allows users to manage all your authn and authz needs.

  Scenario: Generate password with gRPC
    When I request to generate a password with gRPC
    Then I should receive a valid password with gRPC

  Scenario: Generate key with gRPC
    When I request to generate a key with gRPC
    Then I should receive a valid key with gRPC
