# \MetadataApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVEOPStandardContentMetadata**](MetadataApi.md#GetVEOPStandardContentMetadata) | **Get** /api/metadata/standardcontent/d/{did}/v/{vid}/e/{eid}/{otype}/{oid}/p/{pid} | 
[**GetWMVEMetadata**](MetadataApi.md#GetWMVEMetadata) | **Get** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid} | 
[**GetWMVEPMetadata**](MetadataApi.md#GetWMVEPMetadata) | **Get** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/p/{pid} | 
[**GetWMVEPsMetadata**](MetadataApi.md#GetWMVEPsMetadata) | **Get** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/p | 
[**GetWMVEsMetadata**](MetadataApi.md#GetWMVEsMetadata) | **Get** /api/metadata/d/{did}/{wvm}/{wvmid}/e | 
[**GetWVMetadata**](MetadataApi.md#GetWVMetadata) | **Get** /api/metadata/d/{did}/{wv}/{wvid} | 
[**UpdateVEOPStandardContentPartMetadata**](MetadataApi.md#UpdateVEOPStandardContentPartMetadata) | **Post** /api/metadata/standardcontent/d/{did}/v/{vid}/e/{eid}/{otype}/{oid}/p/{pid} | 
[**UpdateWVEMetadata**](MetadataApi.md#UpdateWVEMetadata) | **Post** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid} | 
[**UpdateWVEPMetadata**](MetadataApi.md#UpdateWVEPMetadata) | **Post** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/p/{pid} | 
[**UpdateWVMetadata**](MetadataApi.md#UpdateWVMetadata) | **Post** /api/metadata/d/{did}/{wv}/{wvid} | 



## GetVEOPStandardContentMetadata

> GetVEOPStandardContentMetadata(ctx, did, vid, eid, otype, oid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).Thumbnail(thumbnail).Execute()



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
    eid := "eid_example" // string | 
    otype := "otype_example" // string | 
    oid := "oid_example" // string | 
    pid := "pid_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.GetVEOPStandardContentMetadata(context.Background(), did, vid, eid, otype, oid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetVEOPStandardContentMetadata``: %v\n", err)
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
**eid** | **string** |  | 
**otype** | **string** |  | 
**oid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVEOPStandardContentMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **thumbnail** | **bool** |  | [default to false]

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


## GetWMVEMetadata

> GetWMVEMetadata(ctx, did, wvm, wvmid, eid).Configuration(configuration).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).Thumbnail(thumbnail).Execute()



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
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    depth := "depth_example" // string |  (optional) (default to "1")
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.GetWMVEMetadata(context.Background(), did, wvm, wvmid, eid).Configuration(configuration).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWMVEMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiGetWMVEMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **depth** | **string** |  | [default to &quot;1&quot;]
 **thumbnail** | **bool** |  | [default to false]

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


## GetWMVEPMetadata

> GetWMVEPMetadata(ctx, did, wvm, wvmid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Thumbnail(thumbnail).Execute()



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
    pid := "pid_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.GetWMVEPMetadata(context.Background(), did, wvm, wvmid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWMVEPMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWMVEPMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

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


## GetWMVEPsMetadata

> GetWMVEPsMetadata(ctx, did, wvm, wvmid, eid).Configuration(configuration).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Thumbnail(thumbnail).Execute()



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
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.GetWMVEPsMetadata(context.Background(), did, wvm, wvmid, eid).Configuration(configuration).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWMVEPsMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiGetWMVEPsMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

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


## GetWMVEsMetadata

> GetWMVEsMetadata(ctx, did, wvm, wvmid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).Thumbnail(thumbnail).Execute()



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
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    depth := "depth_example" // string |  (optional) (default to "1")
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.GetWMVEsMetadata(context.Background(), did, wvm, wvmid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWMVEsMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWMVEsMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **depth** | **string** |  | [default to &quot;1&quot;]
 **thumbnail** | **bool** |  | [default to false]

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


## GetWVMetadata

> GetWVMetadata(ctx, did, wv, wvid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).Thumbnail(thumbnail).Execute()



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
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    depth := "depth_example" // string |  (optional) (default to "1")
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.GetWVMetadata(context.Background(), did, wv, wvid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWVMetadata``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiGetWVMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **depth** | **string** |  | [default to &quot;1&quot;]
 **thumbnail** | **bool** |  | [default to false]

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


## UpdateVEOPStandardContentPartMetadata

> UpdateVEOPStandardContentPartMetadata(ctx, did, vid, eid, otype, oid, pid).Body(body).LinkDocumentId(linkDocumentId).Execute()



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
    eid := "eid_example" // string | 
    otype := "otype_example" // string | 
    oid := "oid_example" // string | 
    pid := "pid_example" // string | 
    body := "body_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.UpdateVEOPStandardContentPartMetadata(context.Background(), did, vid, eid, otype, oid, pid, body).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateVEOPStandardContentPartMetadata``: %v\n", err)
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
**eid** | **string** |  | 
**otype** | **string** |  | 
**oid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVEOPStandardContentPartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | **string** |  | 
 **linkDocumentId** | **string** |  | 

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


## UpdateWVEMetadata

> UpdateWVEMetadata(ctx, did, wvm, wvmid, eid).Body(body).Configuration(configuration).Execute()



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
    body := "body_example" // string | 
    configuration := "configuration_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.UpdateWVEMetadata(context.Background(), did, wvm, wvmid, eid, body).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateWVEMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiUpdateWVEMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 
 **configuration** | **string** |  | 

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


## UpdateWVEPMetadata

> UpdateWVEPMetadata(ctx, did, wvm, wvmid, eid, pid, subResource).Body(body).Configuration(configuration).Execute()



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
    pid := "pid_example" // string | 
    subResource := "subResource_example" // string | 
    body := "body_example" // string | 
    configuration := "configuration_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.UpdateWVEPMetadata(context.Background(), did, wvm, wvmid, eid, pid, subResource, body).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateWVEPMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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
**pid** | **string** |  | 
**subResource** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWVEPMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | **string** |  | 
 **configuration** | **string** |  | 

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


## UpdateWVMetadata

> UpdateWVMetadata(ctx, did, wv, wvid).Body(body).Execute()



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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetadataApi.UpdateWVMetadata(context.Background(), did, wv, wvid, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateWVMetadata``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWVMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

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

