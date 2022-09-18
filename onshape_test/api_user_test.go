package onshape_test

import (
	"testing"

	"github.com/toebes/go-client/onshape"
)

func TestUserAPI(t *testing.T) {
	InitializeTester[*onshape.UserApiService](t)

	OpenAPITest{
		Call:   onshape.ApiSessionRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiSessionInfoRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetUserSettingsCurrentLoggedInUserRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetUserSettingsRequest{},
		Expect: Todo(),
	}.Execute()

}
