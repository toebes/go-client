package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestMetadataCategoryAPI(t *testing.T) {
	InitializeTester[*onshape.MetadataCategoryApiService](t)

	OpenAPITest{
		Call:   onshape.ApiGetCategoryPropertiesRequest{},
		Expect: Todo(),
	}.Execute()

}
