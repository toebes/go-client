package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestBillingAPI(t *testing.T) {
	InitializeTester[*onshape.BillingApiService](t)

	OpenAPITest{
		Call:   onshape.ApiGetClientPlansRequest{},
		Expect: Todo(),
	}.Execute()

}
