# \ThumbnailsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplicationThumbnails**](ThumbnailsApi.md#DeleteApplicationThumbnails) | **Delete** /api/thumbnails/d/{did}/{wv}/{wvid}/e/{eid} | 
[**GetConfiguredElementThumbnailWithSize**](ThumbnailsApi.md#GetConfiguredElementThumbnailWithSize) | **Get** /api/thumbnails/d/{did}/w/{wid}/e/{eid}/c/{cid}/s/{sz} | 
[**GetDocumentThumbnail**](ThumbnailsApi.md#GetDocumentThumbnail) | **Get** /api/thumbnails/d/{did}/w/{wid} | 
[**GetDocumentThumbnailWithSize**](ThumbnailsApi.md#GetDocumentThumbnailWithSize) | **Get** /api/thumbnails/d/{did}/w/{wid}/s/{sz} | 
[**GetElementThumbnail**](ThumbnailsApi.md#GetElementThumbnail) | **Get** /api/thumbnails/d/{did}/{wv}/{wvid}/e/{eid} | 
[**GetElementThumbnailWithApiConfiguration**](ThumbnailsApi.md#GetElementThumbnailWithApiConfiguration) | **Get** /api/thumbnails/d/{did}/w/{wid}/e/{eid}/ac/{cid}/s/{sz} | 
[**GetElementThumbnailWithSize**](ThumbnailsApi.md#GetElementThumbnailWithSize) | **Get** /api/thumbnails/d/{did}/w/{wid}/e/{eid}/s/{sz} | 
[**GetThumbnailForDocument**](ThumbnailsApi.md#GetThumbnailForDocument) | **Get** /api/thumbnails/d/{did} | 
[**GetThumbnailForDocumentAndVersion**](ThumbnailsApi.md#GetThumbnailForDocumentAndVersion) | **Get** /api/thumbnails/d/{did}/v/{vid} | 
[**GetThumbnailForDocumentAndVersionOld**](ThumbnailsApi.md#GetThumbnailForDocumentAndVersionOld) | **Get** /api/thumbnails/document/{did}/version/{vid} | 
[**GetThumbnailForDocumentOld**](ThumbnailsApi.md#GetThumbnailForDocumentOld) | **Get** /api/thumbnails/document/{did} | 
[**SetApplicationElementThumbnail**](ThumbnailsApi.md#SetApplicationElementThumbnail) | **Post** /api/thumbnails/d/{did}/{wv}/{wvid}/e/{eid} | 



## DeleteApplicationThumbnails

> DeleteApplicationThumbnails(ctx, did, wv, wvid, eid).Execute()



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
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    eid := "eid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.DeleteApplicationThumbnails(context.Background(), did, wv, wvid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.DeleteApplicationThumbnails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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


## GetConfiguredElementThumbnailWithSize

> GetConfiguredElementThumbnailWithSize(ctx, did, wid, eid, cid, sz).T(t).RejectEmpty(rejectEmpty).Execute()



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
    eid := "eid_example" // string | 
    cid := "cid_example" // string | 
    sz := "sz_example" // string | 
    t := "t_example" // string |  (optional)
    rejectEmpty := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetConfiguredElementThumbnailWithSize(context.Background(), did, wid, eid, cid, sz).T(t).RejectEmpty(rejectEmpty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetConfiguredElementThumbnailWithSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**cid** | **string** |  | 
**sz** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguredElementThumbnailWithSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **t** | **string** |  | 
 **rejectEmpty** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentThumbnail

> BTThumbnailInfo GetDocumentThumbnail(ctx, did, wid).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetDocumentThumbnail(context.Background(), did, wid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetDocumentThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentThumbnail`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailsApi.GetDocumentThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentThumbnailWithSize

> GetDocumentThumbnailWithSize(ctx, did, wid, sz).T(t).Execute()



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
    sz := "sz_example" // string | 
    t := "t_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetDocumentThumbnailWithSize(context.Background(), did, wid, sz).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetDocumentThumbnailWithSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**sz** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentThumbnailWithSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **t** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementThumbnail

> BTThumbnailInfo GetElementThumbnail(ctx, did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()



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
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    eid := "eid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetElementThumbnail(context.Background(), did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetElementThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementThumbnail`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailsApi.GetElementThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** |  | 

### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementThumbnailWithApiConfiguration

> GetElementThumbnailWithApiConfiguration(ctx, did, wid, eid, cid, sz).T(t).RejectEmpty(rejectEmpty).Execute()



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
    eid := "eid_example" // string | 
    cid := "cid_example" // string | 
    sz := "sz_example" // string | 
    t := "t_example" // string |  (optional)
    rejectEmpty := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetElementThumbnailWithApiConfiguration(context.Background(), did, wid, eid, cid, sz).T(t).RejectEmpty(rejectEmpty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetElementThumbnailWithApiConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**cid** | **string** |  | 
**sz** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementThumbnailWithApiConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **t** | **string** |  | 
 **rejectEmpty** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementThumbnailWithSize

> GetElementThumbnailWithSize(ctx, did, wid, eid, sz).T(t).RejectEmpty(rejectEmpty).Execute()



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
    eid := "eid_example" // string | 
    sz := "sz_example" // string | 
    t := "t_example" // string |  (optional)
    rejectEmpty := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetElementThumbnailWithSize(context.Background(), did, wid, eid, sz).T(t).RejectEmpty(rejectEmpty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetElementThumbnailWithSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**sz** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementThumbnailWithSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **t** | **string** |  | 
 **rejectEmpty** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnailForDocument

> BTThumbnailInfo GetThumbnailForDocument(ctx, did).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetThumbnailForDocument(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetThumbnailForDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThumbnailForDocument`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailsApi.GetThumbnailForDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnailForDocumentAndVersion

> GetThumbnailForDocumentAndVersion(ctx, did, vid).LinkDocumentId(linkDocumentId).Execute()



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
    vid := "vid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetThumbnailForDocumentAndVersion(context.Background(), did, vid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetThumbnailForDocumentAndVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentAndVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkDocumentId** | **string** |  | 

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


## GetThumbnailForDocumentAndVersionOld

> GetThumbnailForDocumentAndVersionOld(ctx, did, vid).Execute()



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
    vid := "vid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetThumbnailForDocumentAndVersionOld(context.Background(), did, vid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetThumbnailForDocumentAndVersionOld``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentAndVersionOldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetThumbnailForDocumentOld

> BTThumbnailInfo GetThumbnailForDocumentOld(ctx, did).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.GetThumbnailForDocumentOld(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.GetThumbnailForDocumentOld``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThumbnailForDocumentOld`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailsApi.GetThumbnailForDocumentOld`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentOldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetApplicationElementThumbnail

> SetApplicationElementThumbnail(ctx, did, wv, wvid, eid).BTApplicationElementThumbnailParamsArray(bTApplicationElementThumbnailParamsArray).Overwrite(overwrite).Execute()



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
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    eid := "eid_example" // string | 
    bTApplicationElementThumbnailParamsArray := openapiclient.BTApplicationElementThumbnailParamsArray{Thumbnails: []BTApplicationElementThumbnailParams{openapiclient.BTApplicationElementThumbnailParams{Base64EncodedImage: "Base64EncodedImage_example", Description: "Description_example", ImageHeight: 123, ImageWidth: 123, IsPrimary: false, UniqueId: "UniqueId_example"})} // BTApplicationElementThumbnailParamsArray | 
    overwrite := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ThumbnailsApi.SetApplicationElementThumbnail(context.Background(), did, wv, wvid, eid, bTApplicationElementThumbnailParamsArray).Overwrite(overwrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailsApi.SetApplicationElementThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetApplicationElementThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTApplicationElementThumbnailParamsArray** | [**BTApplicationElementThumbnailParamsArray**](BTApplicationElementThumbnailParamsArray.md) |  | 
 **overwrite** | **bool** |  | [default to false]

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

