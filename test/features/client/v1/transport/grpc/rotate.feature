@rotate
Feature: Rotate

  Rotate allows users to regenerate an existing configuration

  Scenario: Succesfully rotate all of the configuration
    When I rotate an all of the configuration
    Then I should have a complete rotated configuration
    And I should see a log entry of "generated admin password" in the file "reports/all_rotate.log"
    And I should see a log entry of "generated service password/token" in the file "reports/all_rotate.log"

  Scenario: Succesfully rotate admins of the configuration
    When I rotate an admins of the configuration
    Then I should have the admins rotated in the configuration
    And I should see a log entry of "generated admin password" in the file "reports/admins_rotate.log"

  Scenario: Succesfully rotate services of the configuration
    When I rotate an services of the configuration
    Then I should have the services rotated in the configuration
    And I should see a log entry of "generated service password/token" in the file "reports/services_rotate.log"
