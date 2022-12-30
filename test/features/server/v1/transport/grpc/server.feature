Feature: Server

  Server allows users to manage all your authn and authz needs.

  Scenario Outline: Succesfully generate password with gRPC
    When I request to generate a password with length <length> for gRPC
    Then I should receive a valid password with length <length> for gRPC

    Examples:
      | length |
      | 0      |
      | 32     |
      | 64     |

  Scenario Outline: Unsuccesfully generate password with gRPC
    When I request to generate a password with length <length> for gRPC
    Then I should receive an erroneous password with gRPC

    Examples:
      | length |
      | 1      |
      | 31     |

  Scenario Outline: Generate key with gRPC
    When I request to generate a key with kind "<kind>" with gRPC
    Then I should receive a valid key with kind "<kind>" with gRPC

    Examples:
      | kind    |
      |         |
      | rsa     |
      | ed25519 |

  Scenario Outline: Succesfully get public key with gRPC
    When I request to get the public key with kind "<kind>" with gRPC
    Then I should receive a valid public key with kind "<kind>" with gRPC

    Examples:
      | kind    |
      | rsa     |
      | ed25519 |

  Scenario Outline: Unsuccesfully get public key with gRPC
    When I request to get the public key with kind "<kind>" with gRPC
    Then I should receive a not found public key with gRPC

    Examples:
      | kind         |
      | non_existent |

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

  Scenario Outline: Succesfully generate service token with gRPC
    When I request to generate a allowed service token with kind "<kind>" with gRPC
    Then I should receive a valid service token with kind "<kind>" with gRPC

    Examples:
      | kind   |
      |        |
      | jwt    |
      | branca |
      | paseto |

  Scenario Outline: Unsuccesfully generate service token with gRPC
    When I request to generate a disallowed service token with kind "<kind>" with gRPC
    Then I should receive a disallowed service token with gRPC

    Examples:
      | kind            |
      | empty           |
      | not_supported   |
      | not_credentials |
      | invalid_token   |

  Scenario Outline: Succesfully verify service token with gRPC
    Given I request to generate a allowed service token with kind "<kind>" with gRPC
    When I request to verify an allowed service token with kind "<kind>" with gRPC
    Then I should have a valid service token with gRPC

    Examples:
      | kind   |
      |        |
      | jwt    |
      | branca |
      | paseto |

  Scenario Outline: Unsuccesfully verify service token with gRPC
    When I request to verify a disallowed service token with gRPC:
      | token | <token> |
      | issue | <issue> |
    Then I should receive a disallowed verification of service token with gRPC

    Examples: JWT token
      | token | issue           |
      | jwt   | empty           |
      | jwt   | not_supported   |
      | jwt   | not_credentials |
      | jwt   | invalid_token   |
      | jwt   | valid_token     |

    Examples: Branca token
      | token  | issue           |
      | branca | empty           |
      | branca | not_supported   |
      | branca | not_credentials |
      | branca | invalid_token   |
      | branca | valid_token     |

    Examples: Paseto token
      | token  | issue           |
      | paseto | empty           |
      | paseto | not_supported   |
      | paseto | not_credentials |
      | paseto | invalid_token   |
      | paseto | valid_token     |
