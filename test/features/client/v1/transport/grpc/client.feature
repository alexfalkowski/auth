Feature: Client

  Client allows users to manage all your authn and authz needs.

  Scenario: Succesfully generate an access token
    When I generate an access token
    Then I should have a generated access token
    And I should see a log entry of "generated access token" in the file "reports/client.log"

  Scenario: Succesfully generate a service token
    When I generate a service token
    Then I should have a generated service token
    And I should see a log entry of "generated service token" in the file "reports/client.log"

  Scenario: Succesfully verify a service token
    Given I request to generate a allowed service token with kind "jwt" with gRPC
    When I verify a service token
    Then I should have a verified service token
    And I should see a log entry of "verified service token" in the file "reports/client.log"
