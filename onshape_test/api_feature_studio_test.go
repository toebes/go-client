package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestFeatureStudioAPI(t *testing.T) {
	InitializeTester[*onshape.FeatureStudioApiService](t)

	OpenAPITest{
		Call:   onshape.ApiCreateFeatureStudioRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetFeatureStudioContentsRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateFeatureStudioContentsRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetFeatureStudioSpecsRequest{},
		Expect: Todo(),
	}.Execute()

}
