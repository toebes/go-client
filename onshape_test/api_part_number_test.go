package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestPartNumberAPI(t *testing.T) {
	InitializeTester[*onshape.PartNumberApiService](t)

	OpenAPITest{
		Call:   onshape.ApiNextNumbersRequest{},
		Expect: Todo(),
	}.Execute()

}
