package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestVersionAPI(t *testing.T) {
	InitializeTester[*onshape.VersionApiService](t)

	OpenAPITest{
		Call:   onshape.ApiGetAllVersionsRequest{},
		Expect: Todo(),
	}.Execute()

}
