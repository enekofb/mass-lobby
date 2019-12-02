package acceptance

import (
	"flag"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/enekofb/mass-lobby/acceptance/steps"
	"os"
	"testing"
)

var opt = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I want to search policies$`, steps.IWantToSearchPolicies)
	s.Step(`^I dont specify any search criteria$`, steps.IDontSpecifyAnySearchCriteria)
	s.Step(`^I search policies$`, steps.ISearchPolicies)
	s.Step(`^I receive a list o policies matching my search criteria$`, steps.IReceiveAListOPoliciesMatchingMySearchCriteria)
}
