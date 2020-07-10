# \CompaniesApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindCompany**](CompaniesApi.md#FindCompany) | **Get** /api/companies | Get User companies.
[**GetCompany**](CompaniesApi.md#GetCompany) | **Get** /api/companies/{cid} | Get company.



## FindCompany

> BTListResponseBTCompanyInfo FindCompany(ctx).Uid(uid).ActiveOnly(activeOnly).IncludeAll(includeAll).Execute()

Get User companies.

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
    uid := "uid_example" // string |  (optional)
    activeOnly := true // bool |  (optional) (default to true)
    includeAll := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.FindCompany(context.Background(), ).Uid(uid).ActiveOnly(activeOnly).IncludeAll(includeAll).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.FindCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCompany`: BTListResponseBTCompanyInfo
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.FindCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** |  | 
 **activeOnly** | **bool** |  | [default to true]
 **includeAll** | **bool** |  | [default to false]

### Return type

[**BTListResponseBTCompanyInfo**](BTListResponseBTCompanyInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> BTCompanyInfo GetCompany(ctx, cid).Execute()

Get company.

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
    cid := "cid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CompaniesApi.GetCompany(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.GetCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompany`: BTCompanyInfo
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.GetCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTCompanyInfo**](BTCompanyInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

