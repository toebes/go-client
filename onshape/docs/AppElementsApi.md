# \AppElementsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommitTransaction**](AppElementsApi.md#CommitTransaction) | **Post** /api/appelements/d/{did}/w/{wid}/e/{eid}/transactions/{tid} | Commit Transaction
[**CreateElement**](AppElementsApi.md#CreateElement) | **Post** /api/appelements/d/{did}/w/{wid} | Create Element.
[**CreateReference**](AppElementsApi.md#CreateReference) | **Post** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references | Create Reference
[**DeleteAppElementContent**](AppElementsApi.md#DeleteAppElementContent) | **Delete** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/subelements/{sid} | Delete a Sub-element
[**DeleteReference**](AppElementsApi.md#DeleteReference) | **Delete** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Delete Reference
[**GetAppElementHistory**](AppElementsApi.md#GetAppElementHistory) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/history | Get History
[**GetSubElementContent**](AppElementsApi.md#GetSubElementContent) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content | Get Content
[**GetSubelementIds**](AppElementsApi.md#GetSubelementIds) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/ids | Get Sub-element IDs
[**ResolveReference**](AppElementsApi.md#ResolveReference) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Resolve Reference
[**ResolveReferences**](AppElementsApi.md#ResolveReferences) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/resolvereferences | Resolve references.
[**StartTransaction**](AppElementsApi.md#StartTransaction) | **Post** /api/appelements/d/{did}/w/{wid}/e/{eid}/transactions | Start Transaction
[**UpdateAppElement**](AppElementsApi.md#UpdateAppElement) | **Post** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content | Update Element
[**UpdateReference**](AppElementsApi.md#UpdateReference) | **Post** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Update Reference



## CommitTransaction

> BTAppElementModifyInfo CommitTransaction(ctx, did, eid, wid, tid).BTAppElementCommitTransactionParams(bTAppElementCommitTransactionParams).Execute()

Commit Transaction

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
    wid := "wid_example" // string | 
    tid := "tid_example" // string | 
    bTAppElementCommitTransactionParams := openapiclient.BTAppElementCommitTransactionParams{Description: "Description_example", ReturnError: false, TransactionId: "TransactionId_example"} // BTAppElementCommitTransactionParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.CommitTransaction(context.Background(), did, eid, wid, tid, bTAppElementCommitTransactionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.CommitTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommitTransaction`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.CommitTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTAppElementCommitTransactionParams** | [**BTAppElementCommitTransactionParams**](BTAppElementCommitTransactionParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateElement

> BTAppElementModifyInfo CreateElement(ctx, did, wid).BTAppElementParams(bTAppElementParams).Execute()

Create Element.

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
    bTAppElementParams := openapiclient.BTAppElementParams{Description: "Description_example", FormatId: "FormatId_example", Location: openapiclient.BTElementLocationParams{ElementId: "ElementId_example", GroupId: "GroupId_example", Position: 123}, Name: "Name_example", Subelements: []BTAppElementChangeParams{openapiclient.BTAppElementChangeParams{BaseContent: "BaseContent_example", Delta: "Delta_example", SubelementId: "SubelementId_example"})} // BTAppElementParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.CreateElement(context.Background(), did, wid, bTAppElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.CreateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateElement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.CreateElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTAppElementParams** | [**BTAppElementParams**](BTAppElementParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReference

> BTAppElementReferenceInfo CreateReference(ctx, did, eid, wvm, wvmid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()

Create Reference

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    bTAppElementReferenceParams := openapiclient.BTAppElementReferenceParams{HasDocumentMicroversions: false, IdTag: "IdTag_example", IdTagMicroversionId: "IdTagMicroversionId_example", IsSketchOnly: false, ParentChangeId: "ParentChangeId_example", PartNumber: "PartNumber_example", PureSketch: false, ReferenceType: 123, ReturnError: false, Revision: "Revision_example", SketchIds: []string{"SketchIds_example"), TargetConfiguration: "TargetConfiguration_example", TargetDocumentId: "TargetDocumentId_example", TargetElementId: "TargetElementId_example", TargetMicroversionId: "TargetMicroversionId_example", TargetVersionId: "TargetVersionId_example", TrackNewVersions: false, TransactionId: "TransactionId_example", UpdateSketchInfo: false} // BTAppElementReferenceParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.CreateReference(context.Background(), did, eid, wvm, wvmid, bTAppElementReferenceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.CreateReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReference`: BTAppElementReferenceInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.CreateReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTAppElementReferenceParams** | [**BTAppElementReferenceParams**](BTAppElementReferenceParams.md) |  | 

### Return type

[**BTAppElementReferenceInfo**](BTAppElementReferenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppElementContent

> BTAppElementModifyInfo DeleteAppElementContent(ctx, did, eid, wvm, wvmid, sid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()

Delete a Sub-element

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    sid := "sid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    description := "description_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.DeleteAppElementContent(context.Background(), did, eid, wvm, wvmid, sid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.DeleteAppElementContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppElementContent`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.DeleteAppElementContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppElementContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **description** | **string** |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReference

> BTAppElementReferenceInfo DeleteReference(ctx, did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()

Delete Reference

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    rid := "rid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    description := "description_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.DeleteReference(context.Background(), did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.DeleteReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteReference`: BTAppElementReferenceInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.DeleteReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **description** | **string** |  | 

### Return type

[**BTAppElementReferenceInfo**](BTAppElementReferenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppElementHistory

> BTAppElementHistoryInfo GetAppElementHistory(ctx, did, eid, wvm, wvmid).Execute()

Get History

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.GetAppElementHistory(context.Background(), did, eid, wvm, wvmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.GetAppElementHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppElementHistory`: BTAppElementHistoryInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.GetAppElementHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppElementHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTAppElementHistoryInfo**](BTAppElementHistoryInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubElementContent

> BTAppElementContentInfo GetSubElementContent(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).BaseChangeId(baseChangeId).SubelementId(subelementId).LinkDocumentId(linkDocumentId).Execute()

Get Content

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    changeId := "changeId_example" // string |  (optional)
    baseChangeId := "baseChangeId_example" // string |  (optional)
    subelementId := "subelementId_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.GetSubElementContent(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).BaseChangeId(baseChangeId).SubelementId(subelementId).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.GetSubElementContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubElementContent`: BTAppElementContentInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.GetSubElementContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubElementContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 
 **baseChangeId** | **string** |  | 
 **subelementId** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTAppElementContentInfo**](BTAppElementContentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubelementIds

> BTAppElementModifyInfo GetSubelementIds(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()

Get Sub-element IDs

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    changeId := "changeId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.GetSubelementIds(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.GetSubelementIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubelementIds`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.GetSubelementIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubelementIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveReference

> BTAppElementReferenceResolveInfo ResolveReference(ctx, did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).Execute()

Resolve Reference

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    rid := "rid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    includeInternal := true // bool |  (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.ResolveReference(context.Background(), did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.ResolveReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveReference`: BTAppElementReferenceResolveInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.ResolveReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **includeInternal** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 

### Return type

[**BTAppElementReferenceResolveInfo**](BTAppElementReferenceResolveInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveReferences

> BTAppElementReferencesResolveInfo ResolveReferences(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).ReferenceIds(referenceIds).Execute()

Resolve references.

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    includeInternal := true // bool |  (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    referenceIds := "referenceIds_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.ResolveReferences(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).ReferenceIds(referenceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.ResolveReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveReferences`: BTAppElementReferencesResolveInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.ResolveReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **includeInternal** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 
 **referenceIds** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTAppElementReferencesResolveInfo**](BTAppElementReferencesResolveInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTransaction

> BTAppElementModifyInfo StartTransaction(ctx, did, eid, wid).BTAppElementStartTransactionParams(bTAppElementStartTransactionParams).Execute()

Start Transaction

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
    wid := "wid_example" // string | 
    bTAppElementStartTransactionParams := openapiclient.BTAppElementStartTransactionParams{ParentChangeId: "ParentChangeId_example", ReturnError: false} // BTAppElementStartTransactionParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.StartTransaction(context.Background(), did, eid, wid, bTAppElementStartTransactionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.StartTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartTransaction`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.StartTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAppElementStartTransactionParams** | [**BTAppElementStartTransactionParams**](BTAppElementStartTransactionParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppElement

> BTAppElementModifyInfo UpdateAppElement(ctx, did, eid, wvm, wvmid).BTAppElementUpdateParams(bTAppElementUpdateParams).Execute()

Update Element

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    bTAppElementUpdateParams := openapiclient.BTAppElementUpdateParams{Changes: []BTAppElementChangeParams{openapiclient.BTAppElementChangeParams{BaseContent: "BaseContent_example", Delta: "Delta_example", SubelementId: "SubelementId_example"}), Description: "Description_example", ParentChangeId: "ParentChangeId_example", ReturnError: false, TransactionId: "TransactionId_example"} // BTAppElementUpdateParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.UpdateAppElement(context.Background(), did, eid, wvm, wvmid, bTAppElementUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.UpdateAppElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppElement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.UpdateAppElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTAppElementUpdateParams** | [**BTAppElementUpdateParams**](BTAppElementUpdateParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReference

> BTAppElementReferenceInfo UpdateReference(ctx, did, eid, wvm, wvmid, rid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()

Update Reference

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    rid := "rid_example" // string | 
    bTAppElementReferenceParams := openapiclient.BTAppElementReferenceParams{HasDocumentMicroversions: false, IdTag: "IdTag_example", IdTagMicroversionId: "IdTagMicroversionId_example", IsSketchOnly: false, ParentChangeId: "ParentChangeId_example", PartNumber: "PartNumber_example", PureSketch: false, ReferenceType: 123, ReturnError: false, Revision: "Revision_example", SketchIds: []string{"SketchIds_example"), TargetConfiguration: "TargetConfiguration_example", TargetDocumentId: "TargetDocumentId_example", TargetElementId: "TargetElementId_example", TargetMicroversionId: "TargetMicroversionId_example", TargetVersionId: "TargetVersionId_example", TrackNewVersions: false, TransactionId: "TransactionId_example", UpdateSketchInfo: false} // BTAppElementReferenceParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppElementsApi.UpdateReference(context.Background(), did, eid, wvm, wvmid, rid, bTAppElementReferenceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementsApi.UpdateReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReference`: BTAppElementReferenceInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementsApi.UpdateReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **bTAppElementReferenceParams** | [**BTAppElementReferenceParams**](BTAppElementReferenceParams.md) |  | 

### Return type

[**BTAppElementReferenceInfo**](BTAppElementReferenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

