Feature: search policies
  As a policy evaluator 
  I want to search all existing policies that matches certain criteria
  In order to evaluate
  
  Scenario: search all policies
    Given I want to search policies
    And I dont specify any search criteria
    When I search policies 
    Then I receive a list o policies matching my search criteria
