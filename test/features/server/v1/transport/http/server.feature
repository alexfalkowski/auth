Feature: Server

  Server allows users to manage all your authn and authz needs.

  Scenario Outline: Generate password with HTTP
    When I request to generate a password with length <length> for HTTP
    Then I should receive a valid password with length <length> for HTTP

    Examples:
      | length |
      | 0      |
      | 32     |
      | 64     |

  Scenario Outline: Unsuccesfully generate password with HTTP
    When I request to generate a password with length <length> for HTTP
    Then I should receive an erroneous password with HTTP

    Examples:
      | length |
      | 1      |
      | 31     |

  Scenario Outline: Generate key with HTTP
    When I request to generate a key with kind "<kind>" with HTTP
    Then I should receive a valid key with kind "<kind>" with HTTP

    Examples:
      | kind    |
      |         |
      | rsa     |
      | ed25519 |

  Scenario Outline: Succesfully get public key with HTTP
    When I request to get the public key with kind "<kind>" with HTTP
    Then I should receive a valid public key with kind "<kind>" with HTTP

    Examples:
      | kind    |
      | rsa     |
      | ed25519 |

  Scenario Outline: Unsuccesfully get public key with HTTP
    When I request to get the public key with kind "<kind>" with HTTP
    Then I should receive a not found public key with HTTP

    Examples:
      | kind         |
      | non_existent |

  Scenario: Succesfully get jwks with HTTP
    When I request to get the jwks with HTTP
    Then I should receive a valid jwks with HTTP

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

  Scenario: Succesfully generate oauth token with HTTP
    When I request to generate an allowed oauth token with HTTP
    Then I should receive a valid oauth token with HTTP

  Scenario Outline: Unsuccesfully generate oauth token with HTTP
    When I request to generate a disallowed oauth token of kind "<kind>" with HTTP
    Then I should receive a disallowed oauth token with HTTP

    Examples:
      | kind                  |
      | missing_client_id     |
      | missing_client_secret |

  Scenario Outline: Succesfully generate service token with HTTP
    When I request to generate an allowed service token with kind "<kind>" with HTTP
    Then I should receive a valid service token with kind "<kind>" with HTTP

    Examples:
      | kind   |
      |        |
      | jwt    |
      | paseto |

  Scenario Outline: Unsuccesfully generate service token with HTTP
    When I request to generate a disallowed service token with kind "<kind>" with HTTP
    Then I should receive a disallowed service token with HTTP

    Examples:
      | kind            |
      | empty           |
      | not_supported   |
      | not_credentials |
      | invalid_token   |

  Scenario Outline: Succesfully verify service token with HTTP
    Given I request to generate an allowed service token with kind "<kind>" with HTTP
    When I request to verify an allowed service token with kind "<kind>" with HTTP
    Then I should have a valid service token with HTTP

    Examples:
      | kind   |
      |        |
      | jwt    |
      | paseto |

  Scenario Outline: Unsuccesfully verify service token with HTTP
    When I request to verify a disallowed service token with HTTP:
      | token | <token> |
      | issue | <issue> |
    Then I should receive a disallowed verification of service token with HTTP

    Examples: JWT token
      | token | issue           |
      | jwt   | empty           |
      | jwt   | not_supported   |
      | jwt   | not_credentials |
      | jwt   | invalid_token   |
      | jwt   | valid_token     |

    Examples: Paseto token
      | token  | issue           |
      | paseto | empty           |
      | paseto | not_supported   |
      | paseto | not_credentials |
      | paseto | invalid_token   |
      | paseto | valid_token     |
