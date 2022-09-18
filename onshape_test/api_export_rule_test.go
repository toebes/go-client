package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestExportRuleAPI(t *testing.T) {
	InitializeTester[*onshape.ExportRuleApiService](t)

	OpenAPITest{
		Call:   onshape.ApiGetValidRuleOptionsRequest{},
		Expect: Todo(),
	}.Execute()

}
