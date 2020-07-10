# \TeamsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Find**](TeamsApi.md#Find) | **Get** /api/teams | 
[**GetTeam**](TeamsApi.md#GetTeam) | **Get** /api/teams/{tid} | 



## Find

> BTGlobalTreeNodeListResponseBTTeamInfo Find(ctx).Prefix(prefix).Uid(uid).CompanyId(companyId).IncludeCompanyOwnedTeams(includeCompanyOwnedTeams).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    prefix := "prefix_example" // string |  (optional) (default to "")
    uid := "uid_example" // string |  (optional)
    companyId := "companyId_example" // string |  (optional)
    includeCompanyOwnedTeams := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.Find(context.Background(), ).Prefix(prefix).Uid(uid).CompanyId(companyId).IncludeCompanyOwnedTeams(includeCompanyOwnedTeams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.Find``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Find`: BTGlobalTreeNodeListResponseBTTeamInfo
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.Find`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** |  | [default to &quot;&quot;]
 **uid** | **string** |  | 
 **companyId** | **string** |  | 
 **includeCompanyOwnedTeams** | **bool** |  | [default to true]

### Return type

[**BTGlobalTreeNodeListResponseBTTeamInfo**](BTGlobalTreeNodeListResponseBTTeamInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeam

> BTTeamInfo GetTeam(ctx, tid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tid := "tid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsApi.GetTeam(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.GetTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeam`: BTTeamInfo
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.GetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTTeamInfo**](BTTeamInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

