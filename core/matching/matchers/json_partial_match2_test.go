package matchers_test

import (
	"testing"

	"github.com/SpectoLabs/hoverfly/core/matching/matchers"
	. "github.com/onsi/gomega"
)

func Test0True(t *testing.T) {
	RegisterTestingT(t)

	Expect(matchers.JsonPartialMatch2(
		`{"fieldA": "valueA"}`,
		`{
	"test": {
		"json": true,
		"minified": true,
		"someObject": {
			"fieldA": "valueA"
		}
	}}`)).To(BeTrue())
}

func Test1True(t *testing.T) {
	RegisterTestingT(t)

	Expect(matchers.JsonPartialMatch2(
		`{"NAME": "79684881033", "REDIRECT_NUMBER": "79684881033"}`,
		`{
      "jsonrpc": "2.0",
      "id": "1",
      "result": {
            "redirect_type": 1,
            "followme_struct": [
                  3, [{
                        "I_FOLLOW_ORDER": "1",
                        "ACTIVE": "Y",
                        "NAME": "79684881033",
                        "REDIRECT_NUMBER": "79684881033",
                        "PERIOD": "always",
                        "PERIOD_DESCRIPTION": "always",
                        "TIMEOUT": "15"
                  }, {
                        "I_FOLLOW_ORDER": "2"
                        "ACTIVE": "Y",
                        "NAME": "79684881034",
                        "REDIRECT_NUMBER": "79684881034",
                        "PERIOD": "always",
                        "PERIOD_DESCRIPTION": "always",
                        "TIMEOUT": "15"

                  }]
            ]
	}`)).To(BeTrue())
}

func Test2True(t *testing.T) {
	RegisterTestingT(t)

	Expect(matchers.JsonPartialMatch2(
		`
{
  "I_FOLLOW_ORDER": "2",
  "ACTIVE": "Y",
}`,
		`
{
  "followme_struct": [
    3,
    [
      {
        "I_FOLLOW_ORDER": "1",
        "ACTIVE": "Y",
        "NAME": "79684881033",
        "REDIRECT_NUMBER": "79684881033",
        "PERIOD": "always",
        "PERIOD_DESCRIPTION": "always",
        "TIMEOUT": "15"
      },
      {
        "I_FOLLOW_ORDER": "2",
        "ACTIVE": "Y",
        "NAME": "79684881034",
        "REDIRECT_NUMBER": "79684881034",
        "PERIOD": "always",
        "PERIOD_DESCRIPTION": "always",
        "TIMEOUT": "15"
      }
    ]
  ]
}`)).To(BeTrue())
}
