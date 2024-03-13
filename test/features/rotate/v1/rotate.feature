Feature: Rotate

  Rotate allows users to regenerate an existing configuration

  Scenario: Succesfully rotate an existing configuration
    Given I have an existing configuration
    When I rotate an existing configuration
    Then I should have a working rotated configuration
    And I should see a log entry of "generated admin password" in the file "reports/rotate.log"
    And I should see a log entry of "generated service password/token" in the file "reports/rotate.log"
