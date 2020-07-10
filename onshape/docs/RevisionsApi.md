# \RevisionsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnumerateRevisions**](RevisionsApi.md#EnumerateRevisions) | **Get** /api/revisions/companies/{cid} | 
[**GetLatestInDocumentOrCompany**](RevisionsApi.md#GetLatestInDocumentOrCompany) | **Get** /api/revisions/{cd}/{cdid}/p/{pnum}/latest | 
[**GetRevisionHistoryInCompany**](RevisionsApi.md#GetRevisionHistoryInCompany) | **Get** /api/revisions/companies/{cid}/partnumber/{pnum} | 



## EnumerateRevisions

> BTListResponseBTRevisionInfo EnumerateRevisions(ctx, cid).ElementType(elementType).Limit(limit).Offset(offset).LatestOnly(latestOnly).After(after).Execute()



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
    elementType := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional) (default to 20)
    offset := 987 // int32 |  (optional) (default to 0)
    latestOnly := true // bool |  (optional) (default to false)
    after := 987 // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RevisionsApi.EnumerateRevisions(context.Background(), cid).ElementType(elementType).Limit(limit).Offset(offset).LatestOnly(latestOnly).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionsApi.EnumerateRevisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnumerateRevisions`: BTListResponseBTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionsApi.EnumerateRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnumerateRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **elementType** | **int32** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **latestOnly** | **bool** |  | [default to false]
 **after** | **int64** |  | 

### Return type

[**BTListResponseBTRevisionInfo**](BTListResponseBTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestInDocumentOrCompany

> BTRevisionInfo GetLatestInDocumentOrCompany(ctx, cd, cdid, pnum).Et(et).Execute()



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
    cd := "cd_example" // string | 
    cdid := "cdid_example" // string | 
    pnum := "pnum_example" // string | 
    et := "et_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RevisionsApi.GetLatestInDocumentOrCompany(context.Background(), cd, cdid, pnum).Et(et).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionsApi.GetLatestInDocumentOrCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestInDocumentOrCompany`: BTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionsApi.GetLatestInDocumentOrCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cd** | **string** |  | 
**cdid** | **string** |  | 
**pnum** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestInDocumentOrCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **et** | **string** |  | 

### Return type

[**BTRevisionInfo**](BTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionHistoryInCompany

> BTListResponseBTRevisionInfo GetRevisionHistoryInCompany(ctx, cid, pnum).ElementType(elementType).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()



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
    pnum := "pnum_example" // string | 
    elementType := "elementType_example" // string |  (optional)
    fillApprovers := true // bool |  (optional) (default to false)
    fillExportPermission := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RevisionsApi.GetRevisionHistoryInCompany(context.Background(), cid, pnum).ElementType(elementType).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionsApi.GetRevisionHistoryInCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionHistoryInCompany`: BTListResponseBTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionsApi.GetRevisionHistoryInCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 
**pnum** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionHistoryInCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **elementType** | **string** |  | 
 **fillApprovers** | **bool** |  | [default to false]
 **fillExportPermission** | **bool** |  | [default to false]

### Return type

[**BTListResponseBTRevisionInfo**](BTListResponseBTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

