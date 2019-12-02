package steps

import "github.com/DATA-DOG/godog"

func IWantToSearchPolicies() error {
	NewPolicies()
	return godog.ErrPending
}

func NewPolicies() {

}

func IDontSpecifyAnySearchCriteria() error {
	return godog.ErrPending
}

func ISearchPolicies() error {
	return godog.ErrPending
}

func IReceiveAListOPoliciesMatchingMySearchCriteria() error {
	return godog.ErrPending
}
