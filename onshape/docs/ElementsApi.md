# \ElementsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyElementFromSourceDocument**](ElementsApi.md#CopyElementFromSourceDocument) | **Post** /api/elements/copyelement/{did}/workspace/{wid} | Copy Element
[**DecodeConfiguration**](ElementsApi.md#DecodeConfiguration) | **Get** /api/elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configurationencodings/{cid} | Decode Configuration String
[**DeleteElement**](ElementsApi.md#DeleteElement) | **Delete** /api/elements/d/{did}/w/{wid}/e/{eid} | Delete Element
[**EncodeConfigurationMap**](ElementsApi.md#EncodeConfigurationMap) | **Post** /api/elements/d/{did}/e/{eid}/configurationencodings | Encode Configuration
[**GetConfiguration**](ElementsApi.md#GetConfiguration) | **Get** /api/elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration | Get Configuration
[**GetElementTranslatorFormatsByVersionOrWorkspace**](ElementsApi.md#GetElementTranslatorFormatsByVersionOrWorkspace) | **Get** /api/elements/translatorFormats/{did}/{wv}/{wvid}/{eid} | Get Element Translator Formats
[**UpdateConfiguration**](ElementsApi.md#UpdateConfiguration) | **Post** /api/elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration | Update Configuration
[**UpdateReferences**](ElementsApi.md#UpdateReferences) | **Post** /api/elements/d/{did}/w/{wid}/e/{eid}/updatereferences | Update or replace node references



## CopyElementFromSourceDocument

> BTDocumentElementInfo CopyElementFromSourceDocument(ctx, did, wid).BTCopyElementParams(bTCopyElementParams).Execute()

Copy Element

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
    bTCopyElementParams := openapiclient.BTCopyElementParams{AnchorElementId: "AnchorElementId_example", DocumentIdSource: "DocumentIdSource_example", ElementIdSource: "ElementIdSource_example", IsGroupAnchor: false, WorkspaceIdSource: "WorkspaceIdSource_example"} // BTCopyElementParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ElementsApi.CopyElementFromSourceDocument(context.Background(), did, wid, bTCopyElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.CopyElementFromSourceDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyElementFromSourceDocument`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.CopyElementFromSourceDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyElementFromSourceDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTCopyElementParams** | [**BTCopyElementParams**](BTCopyElementParams.md) |  | 

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


## DecodeConfiguration

> BTConfigurationInfo DecodeConfiguration(ctx, did, wvm, wvmid, eid, cid).LinkDocumentId(linkDocumentId).IncludeDisplay(includeDisplay).ConfigurationIsId(configurationIsId).Execute()

Decode Configuration String

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
    cid := "cid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    includeDisplay := true // bool |  (optional) (default to false)
    configurationIsId := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ElementsApi.DecodeConfiguration(context.Background(), did, wvm, wvmid, eid, cid).LinkDocumentId(linkDocumentId).IncludeDisplay(includeDisplay).ConfigurationIsId(configurationIsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.DecodeConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DecodeConfiguration`: BTConfigurationInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.DecodeConfiguration`: %v\n", resp)
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
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecodeConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** |  | 
 **includeDisplay** | **bool** |  | [default to false]
 **configurationIsId** | **bool** |  | [default to false]

### Return type

[**BTConfigurationInfo**](BTConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteElement

> DeleteElement(ctx, did, wid, eid).Execute()

Delete Element

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ElementsApi.DeleteElement(context.Background(), did, wid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.DeleteElement``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteElementRequest struct via the builder pattern


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


## EncodeConfigurationMap

> BTEncodedConfigurationInfo EncodeConfigurationMap(ctx, did, eid).BTConfigurationParams(bTConfigurationParams).VersionId(versionId).LinkDocumentId(linkDocumentId).Execute()

Encode Configuration

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
    eid := "eid_example" // string | 
    bTConfigurationParams := openapiclient.BTConfigurationParams{Parameters: []ConfigurationEntry{openapiclient.ConfigurationEntry{ParameterId: "ParameterId_example", ParameterValue: "ParameterValue_example"}), StandardContentParametersId: "StandardContentParametersId_example"} // BTConfigurationParams | 
    versionId := "versionId_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ElementsApi.EncodeConfigurationMap(context.Background(), did, eid, bTConfigurationParams).VersionId(versionId).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.EncodeConfigurationMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncodeConfigurationMap`: BTEncodedConfigurationInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.EncodeConfigurationMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEncodeConfigurationMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTConfigurationParams** | [**BTConfigurationParams**](BTConfigurationParams.md) |  | 
 **versionId** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTEncodedConfigurationInfo**](BTEncodedConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> BTConfigurationInfo GetConfiguration(ctx, did, wvm, wvmid, eid).Execute()

Get Configuration

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
    resp, r, err := api_client.ElementsApi.GetConfiguration(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration`: BTConfigurationInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetConfiguration`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTConfigurationInfo**](BTConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementTranslatorFormatsByVersionOrWorkspace

> []BTModelFormatInfo GetElementTranslatorFormatsByVersionOrWorkspace(ctx, did, wv, wvid, eid).CheckContent(checkContent).Configuration(configuration).Execute()

Get Element Translator Formats

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
    checkContent := true // bool |  (optional) (default to true)
    configuration := "configuration_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ElementsApi.GetElementTranslatorFormatsByVersionOrWorkspace(context.Background(), did, wv, wvid, eid).CheckContent(checkContent).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.GetElementTranslatorFormatsByVersionOrWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementTranslatorFormatsByVersionOrWorkspace`: []BTModelFormatInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.GetElementTranslatorFormatsByVersionOrWorkspace`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetElementTranslatorFormatsByVersionOrWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **checkContent** | **bool** |  | [default to true]
 **configuration** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]BTModelFormatInfo**](BTModelFormatInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> BTConfigurationInfo UpdateConfiguration(ctx, did, wvm, wvmid, eid).Body(body).Execute()

Update Configuration

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
    resp, r, err := api_client.ElementsApi.UpdateConfiguration(context.Background(), did, wvm, wvmid, eid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: BTConfigurationInfo
    fmt.Fprintf(os.Stdout, "Response from `ElementsApi.UpdateConfiguration`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**BTConfigurationInfo**](BTConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReferences

> UpdateReferences(ctx, did, wid, eid).BTUpdateReferenceParams(bTUpdateReferenceParams).Execute()

Update or replace node references

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
    bTUpdateReferenceParams := openapiclient.BTUpdateReferenceParams{ReferenceUpdates: []UpdateParams{openapiclient.UpdateParams{FromReference: openapiclient.BTUniqueDocumentItemParams{ApiConfiguration: "ApiConfiguration_example", DocumentId: "DocumentId_example", ElementId: "ElementId_example", ElementType: "ElementType_example", PartId: "PartId_example", PartNumber: "PartNumber_example", Revision: "Revision_example", VersionId: "VersionId_example", WorkspaceId: "WorkspaceId_example"}, IdsToUpdate: []string{"IdsToUpdate_example"), ToReference: openapiclient.BTUniqueDocumentItemParams{ApiConfiguration: "ApiConfiguration_example", DocumentId: "DocumentId_example", ElementId: "ElementId_example", ElementType: "ElementType_example", PartId: "PartId_example", PartNumber: "PartNumber_example", Revision: "Revision_example", VersionId: "VersionId_example", WorkspaceId: "WorkspaceId_example"}})} // BTUpdateReferenceParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ElementsApi.UpdateReferences(context.Background(), did, wid, eid, bTUpdateReferenceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ElementsApi.UpdateReferences``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTUpdateReferenceParams** | [**BTUpdateReferenceParams**](BTUpdateReferenceParams.md) |  | 

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

