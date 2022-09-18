package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestInsertableAPI(t *testing.T) {
	InitializeTester[*onshape.InsertableApiService](t)

	OpenAPITest{
		Call:   onshape.ApiGetLatestInDocumentRequest{},
		Expect: Todo(),
	}.Execute()

}
