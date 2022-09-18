package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestPublicationAPI(t *testing.T) {
	InitializeTester[*onshape.PublicationApiService](t)

	OpenAPITest{
		Call:   onshape.ApiGetPublicationItemsRequest{},
		Expect: Todo(),
	}.Execute()

}
