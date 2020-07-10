# \ReleaseManagementApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObsoletionPackage**](ReleaseManagementApi.md#CreateObsoletionPackage) | **Post** /api/releasepackages/obsoletion/{wfid} | 
[**CreateReleasePackage**](ReleaseManagementApi.md#CreateReleasePackage) | **Post** /api/releasepackages/release/{wfid} | 
[**GetCompanyReleaseWorkflow**](ReleaseManagementApi.md#GetCompanyReleaseWorkflow) | **Get** /api/releasepackages/companyreleaseworkflow | 
[**GetReleasePackage**](ReleaseManagementApi.md#GetReleasePackage) | **Get** /api/releasepackages/{rpid} | 
[**UpdateReleasePackage**](ReleaseManagementApi.md#UpdateReleasePackage) | **Post** /api/releasepackages/{rpid} | 



## CreateObsoletionPackage

> CreateObsoletionPackage(ctx, wfid).RevisionId(revisionId).Execute()



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
    wfid := "wfid_example" // string | 
    revisionId := "revisionId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.CreateObsoletionPackage(context.Background(), wfid).RevisionId(revisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.CreateObsoletionPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObsoletionPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revisionId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReleasePackage

> CreateReleasePackage(ctx, wfid).BTReleasePackageParams(bTReleasePackageParams).Execute()



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
    wfid := "wfid_example" // string | 
    bTReleasePackageParams := openapiclient.BTReleasePackageParams{Items: []BTReleasePackageItemParams{openapiclient.BTReleasePackageItemParams{Configuration: "Configuration_example", DocumentId: "DocumentId_example", ElementId: "ElementId_example", Href: "Href_example", Id: "Id_example", IsIncluded: false, PartId: "PartId_example", PartNumber: "PartNumber_example", Properties: []BTPropertyValueParam{openapiclient.BTPropertyValueParam{PropertyId: "PropertyId_example", Value: "TODO"}), VersionId: "VersionId_example", WorkspaceId: "WorkspaceId_example"})} // BTReleasePackageParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.CreateReleasePackage(context.Background(), wfid, bTReleasePackageParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.CreateReleasePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleasePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTReleasePackageParams** | [**BTReleasePackageParams**](BTReleasePackageParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyReleaseWorkflow

> BTActiveWorkflowInfo GetCompanyReleaseWorkflow(ctx).DocumentId(documentId).Execute()



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
    documentId := "documentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.GetCompanyReleaseWorkflow(context.Background(), ).DocumentId(documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.GetCompanyReleaseWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyReleaseWorkflow`: BTActiveWorkflowInfo
    fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementApi.GetCompanyReleaseWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyReleaseWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentId** | **string** |  | 

### Return type

[**BTActiveWorkflowInfo**](BTActiveWorkflowInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleasePackage

> BTReleasePackageInfo GetReleasePackage(ctx, rpid).Detailed(detailed).Execute()



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
    rpid := "rpid_example" // string | 
    detailed := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.GetReleasePackage(context.Background(), rpid).Detailed(detailed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.GetReleasePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleasePackage`: BTReleasePackageInfo
    fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementApi.GetReleasePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detailed** | **bool** |  | [default to false]

### Return type

[**BTReleasePackageInfo**](BTReleasePackageInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleasePackage

> BTReleasePackageInfo UpdateReleasePackage(ctx, rpid).BTUpdateReleasePackageParams(bTUpdateReleasePackageParams).Action(action).Wfaction(wfaction).Execute()



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
    rpid := "rpid_example" // string | 
    bTUpdateReleasePackageParams := openapiclient.BTUpdateReleasePackageParams{Empty: false, ItemIds: []string{"ItemIds_example"), Items: []BTReleasePackageItemParams{openapiclient.BTReleasePackageItemParams{Configuration: "Configuration_example", DocumentId: "DocumentId_example", ElementId: "ElementId_example", Href: "Href_example", Id: "Id_example", IsIncluded: false, PartId: "PartId_example", PartNumber: "PartNumber_example", Properties: []BTPropertyValueParam{openapiclient.BTPropertyValueParam{PropertyId: "PropertyId_example", Value: "TODO"}), VersionId: "VersionId_example", WorkspaceId: "WorkspaceId_example"}), Properties: []BTPropertyValueParam{)} // BTUpdateReleasePackageParams | 
    action := "action_example" // string |  (optional) (default to "UPDATE")
    wfaction := "wfaction_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReleaseManagementApi.UpdateReleasePackage(context.Background(), rpid, bTUpdateReleasePackageParams).Action(action).Wfaction(wfaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementApi.UpdateReleasePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReleasePackage`: BTReleasePackageInfo
    fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementApi.UpdateReleasePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleasePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTUpdateReleasePackageParams** | [**BTUpdateReleasePackageParams**](BTUpdateReleasePackageParams.md) |  | 
 **action** | **string** |  | [default to &quot;UPDATE&quot;]
 **wfaction** | **string** |  | 

### Return type

[**BTReleasePackageInfo**](BTReleasePackageInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

