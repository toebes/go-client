package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestFolderAPI(t *testing.T) {
	InitializeTester[*onshape.FolderApiService](t)

	OpenAPITest{
		Call:   onshape.ApiGetFolderAclRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiShareRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUnShareRequest{},
		Expect: Todo(),
	}.Execute()

}
