# \FeatureStudiosApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeatureStudio**](FeatureStudiosApi.md#CreateFeatureStudio) | **Post** /api/featurestudios/d/{did}/w/{wid} | Create Feature Studio
[**GetFeatureStudioContents**](FeatureStudiosApi.md#GetFeatureStudioContents) | **Get** /api/featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid} | Get Feature Studio Contents.
[**GetFeatureStudioSpecs**](FeatureStudiosApi.md#GetFeatureStudioSpecs) | **Get** /api/featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs | Get Feature Studio Specs
[**UpdateFeatureStudioContents**](FeatureStudiosApi.md#UpdateFeatureStudioContents) | **Post** /api/featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid} | Update Feature Studio contents



## CreateFeatureStudio

> BTDocumentElementInfo CreateFeatureStudio(ctx, did, wid).BTModelElementParams(bTModelElementParams).Execute()

Create Feature Studio

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
    did := "did_example" // string | 
    wid := "wid_example" // string | 
    bTModelElementParams := openapiclient.BTModelElementParams{Name: "Name_example"} // BTModelElementParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureStudiosApi.CreateFeatureStudio(context.Background(), did, wid, bTModelElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureStudiosApi.CreateFeatureStudio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeatureStudio`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `FeatureStudiosApi.CreateFeatureStudio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureStudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTModelElementParams** | [**BTModelElementParams**](BTModelElementParams.md) |  | 

### Return type

[**BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureStudioContents

> BTFeatureStudioContents2239 GetFeatureStudioContents(ctx, did, wvm, wvmid, eid).Execute()

Get Feature Studio Contents.

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
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureStudiosApi.GetFeatureStudioContents(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureStudiosApi.GetFeatureStudioContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureStudioContents`: BTFeatureStudioContents2239
    fmt.Fprintf(os.Stdout, "Response from `FeatureStudiosApi.GetFeatureStudioContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureStudioContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTFeatureStudioContents2239**](BTFeatureStudioContents-2239.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureStudioSpecs

> BTFeatureSpecsResponse664 GetFeatureStudioSpecs(ctx, did, wvm, wvmid, eid).Execute()

Get Feature Studio Specs

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
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureStudiosApi.GetFeatureStudioSpecs(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureStudiosApi.GetFeatureStudioSpecs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureStudioSpecs`: BTFeatureSpecsResponse664
    fmt.Fprintf(os.Stdout, "Response from `FeatureStudiosApi.GetFeatureStudioSpecs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureStudioSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTFeatureSpecsResponse664**](BTFeatureSpecsResponse-664.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureStudioContents

> BTFeatureStudioContents2239 UpdateFeatureStudioContents(ctx, did, wvm, wvmid, eid).Body(body).Execute()

Update Feature Studio contents

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
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FeatureStudiosApi.UpdateFeatureStudioContents(context.Background(), did, wvm, wvmid, eid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureStudiosApi.UpdateFeatureStudioContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureStudioContents`: BTFeatureStudioContents2239
    fmt.Fprintf(os.Stdout, "Response from `FeatureStudiosApi.UpdateFeatureStudioContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureStudioContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**BTFeatureStudioContents2239**](BTFeatureStudioContents-2239.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

