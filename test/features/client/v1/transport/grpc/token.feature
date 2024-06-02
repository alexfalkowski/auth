Feature: Token

  Token allows users to manage all generation needs.

  Scenario: Succesfully generate a service token
    When I generate a service token
    Then I should have a generated service token
    And I should see a log entry of "generated service token" in the file "reports/client.log"

  Scenario: Succesfully verify a service token
    Given I request to generate an allowed service token with kind "jwt" with gRPC
    When I verify a service token
    Then I should have a verified service token
    And I should see a log entry of "verified service token" in the file "reports/client.log"
