# \DocumentsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyWorkspace**](DocumentsApi.md#CopyWorkspace) | **Post** /api/documents/{did}/workspaces/{wid}/copy | Copy Workspace
[**CreateDocument**](DocumentsApi.md#CreateDocument) | **Post** /api/documents | Create document.
[**CreateVersion**](DocumentsApi.md#CreateVersion) | **Post** /api/documents/d/{did}/versions | Create Version.
[**CreateWorkspace**](DocumentsApi.md#CreateWorkspace) | **Post** /api/documents/d/{did}/workspaces | Create Workspace
[**DeleteDocument**](DocumentsApi.md#DeleteDocument) | **Delete** /api/documents/{did} | Delete Document
[**DeleteWorkspace**](DocumentsApi.md#DeleteWorkspace) | **Delete** /api/documents/d/{did}/workspaces/{wid} | Delete Workspace
[**DownloadExternalData**](DocumentsApi.md#DownloadExternalData) | **Get** /api/documents/d/{did}/externaldata/{fid} | Download External Data
[**Export2Json**](DocumentsApi.md#Export2Json) | **Post** /api/documents/d/{did}/{wv}/{wvid}/e/{eid}/export | 
[**GetCurrentMicroversion**](DocumentsApi.md#GetCurrentMicroversion) | **Get** /api/documents/d/{did}/{wv}/{wvid}/currentmicroversion | Get Current Document Microversion
[**GetDocument**](DocumentsApi.md#GetDocument) | **Get** /api/documents/{did} | Get Document
[**GetDocumentAcl**](DocumentsApi.md#GetDocumentAcl) | **Get** /api/documents/{did}/acl | Get Access Control List
[**GetDocumentPermissionSet**](DocumentsApi.md#GetDocumentPermissionSet) | **Get** /api/documents/{did}/permissionset | Get Document Permissions
[**GetDocumentVersions**](DocumentsApi.md#GetDocumentVersions) | **Get** /api/documents/d/{did}/versions | Get Versions
[**GetDocumentWorkspaces**](DocumentsApi.md#GetDocumentWorkspaces) | **Get** /api/documents/d/{did}/workspaces | Get Workspaces
[**GetDocuments**](DocumentsApi.md#GetDocuments) | **Get** /api/documents | Get Documents
[**GetElementsInDocument**](DocumentsApi.md#GetElementsInDocument) | **Get** /api/documents/d/{did}/{wvm}/{wvmid}/elements | Get a list of elements in the workspace, version, or microversion of the document.
[**GetInsertables**](DocumentsApi.md#GetInsertables) | **Get** /api/documents/d/{did}/{wvm}/{wvmid}/insertables | Insertable List for Document Version.
[**GetVersion**](DocumentsApi.md#GetVersion) | **Get** /api/documents/d/{did}/versions/{vid} | Get Version
[**MergeIntoWorkspace**](DocumentsApi.md#MergeIntoWorkspace) | **Post** /api/documents/{did}/workspaces/{wid}/merge | Merge into workspace
[**MoveElementsToDocument**](DocumentsApi.md#MoveElementsToDocument) | **Post** /api/documents/d/{did}/w/{wid}/moveelement | Move Elements
[**RestoreFromHistory**](DocumentsApi.md#RestoreFromHistory) | **Post** /api/documents/{did}/w/{wid}/restore/{vm}/{vmid} | Restore version or microversion to workspace.
[**Search**](DocumentsApi.md#Search) | **Post** /api/documents/search | 
[**ShareDocument**](DocumentsApi.md#ShareDocument) | **Post** /api/documents/{did}/share | Share Document
[**SyncApplicationElements**](DocumentsApi.md#SyncApplicationElements) | **Post** /api/documents/d/{did}/w/{wid}/syncApplicationElements | Sync Application Elements
[**UnShareDocument**](DocumentsApi.md#UnShareDocument) | **Delete** /api/documents/{did}/share/{eid} | Unshare Document
[**UpdateDocumentAttributes**](DocumentsApi.md#UpdateDocumentAttributes) | **Post** /api/documents/{did} | Update Document Attributes.
[**UpdateExternalReferencesToLatestDocuments**](DocumentsApi.md#UpdateExternalReferencesToLatestDocuments) | **Post** /api/documents/d/{did}/w/{wid}/e/{eid}/latestdocumentreferences | Update External References to Latest



## CopyWorkspace

> BTCopyDocumentInfo CopyWorkspace(ctx, did, wid).BTCopyDocumentParams(bTCopyDocumentParams).Execute()

Copy Workspace

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
    bTCopyDocumentParams := openapiclient.BTCopyDocumentParams{BetaCapabilityIds: []string{"BetaCapabilityIds_example"), IsPublic: false, NewName: "NewName_example", OwnerId: "OwnerId_example", OwnerTypeIndex: 123, ParentId: "ParentId_example", ProjectId: "ProjectId_example"} // BTCopyDocumentParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.CopyWorkspace(context.Background(), did, wid).BTCopyDocumentParams(bTCopyDocumentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.CopyWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyWorkspace`: BTCopyDocumentInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.CopyWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTCopyDocumentParams** | [**BTCopyDocumentParams**](BTCopyDocumentParams.md) |  | 

### Return type

[**BTCopyDocumentInfo**](BTCopyDocumentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDocument

> BTDocumentInfo CreateDocument(ctx).BTDocumentParams(bTDocumentParams).Execute()

Create document.

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
    bTDocumentParams := openapiclient.BTDocumentParams{BetaCapabilityIds: []string{"BetaCapabilityIds_example"), Description: "Description_example", GenerateUnknownMessages: false, IsEmptyContent: false, IsPublic: false, Name: "Name_example", NotRevisionManaged: false, OwnerEmail: "OwnerEmail_example", OwnerId: "OwnerId_example", OwnerType: 123, ParentId: "ParentId_example", ProjectId: "ProjectId_example", Tags: []string{"Tags_example")} // BTDocumentParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.CreateDocument(context.Background(), bTDocumentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.CreateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDocument`: BTDocumentInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.CreateDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTDocumentParams** | [**BTDocumentParams**](BTDocumentParams.md) |  | 

### Return type

[**BTDocumentInfo**](BTDocumentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVersion

> BTVersionInfo CreateVersion(ctx, did).BTVersionOrWorkspaceParams(bTVersionOrWorkspaceParams).Execute()

Create Version.

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
    bTVersionOrWorkspaceParams := openapiclient.BTVersionOrWorkspaceParams{ClientInteractionMode: "ClientInteractionMode_example", Description: "Description_example", DocumentId: "DocumentId_example", FromHistory: false, IsRelease: false, MicroversionId: "MicroversionId_example", Name: "Name_example", Purpose: 123, ReadOnly: false, VersionId: "VersionId_example", WorkspaceId: "WorkspaceId_example"} // BTVersionOrWorkspaceParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.CreateVersion(context.Background(), did, bTVersionOrWorkspaceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.CreateVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVersion`: BTVersionInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.CreateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTVersionOrWorkspaceParams** | [**BTVersionOrWorkspaceParams**](BTVersionOrWorkspaceParams.md) |  | 

### Return type

[**BTVersionInfo**](BTVersionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspace

> BTWorkspaceInfo CreateWorkspace(ctx, did).BTVersionOrWorkspaceParams(bTVersionOrWorkspaceParams).Execute()

Create Workspace

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
    bTVersionOrWorkspaceParams := openapiclient.BTVersionOrWorkspaceParams{ClientInteractionMode: "ClientInteractionMode_example", Description: "Description_example", DocumentId: "DocumentId_example", FromHistory: false, IsRelease: false, MicroversionId: "MicroversionId_example", Name: "Name_example", Purpose: 123, ReadOnly: false, VersionId: "VersionId_example", WorkspaceId: "WorkspaceId_example"} // BTVersionOrWorkspaceParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.CreateWorkspace(context.Background(), did).BTVersionOrWorkspaceParams(bTVersionOrWorkspaceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.CreateWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkspace`: BTWorkspaceInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.CreateWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTVersionOrWorkspaceParams** | [**BTVersionOrWorkspaceParams**](BTVersionOrWorkspaceParams.md) |  | 

### Return type

[**BTWorkspaceInfo**](BTWorkspaceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> DeleteDocument(ctx, did).Forever(forever).Execute()

Delete Document

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
    forever := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.DeleteDocument(context.Background(), did).Forever(forever).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.DeleteDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forever** | **bool** |  | [default to false]

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


## DeleteWorkspace

> DeleteWorkspace(ctx, did, wid).Execute()

Delete Workspace

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
    resp, r, err := api_client.DocumentsApi.DeleteWorkspace(context.Background(), did, wid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.DeleteWorkspace``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRequest struct via the builder pattern


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


## DownloadExternalData

> *os.File DownloadExternalData(ctx, did, fid).IfNoneMatch(ifNoneMatch).Execute()

Download External Data

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
    fid := "fid_example" // string | 
    ifNoneMatch := "ifNoneMatch_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.DownloadExternalData(context.Background(), did, fid).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.DownloadExternalData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadExternalData`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.DownloadExternalData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadExternalDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+octet-stream;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Export2Json

> Export2Json(ctx, did, wv, wvid, eid).BTExportModelParams(bTExportModelParams).Execute()



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
    bTExportModelParams := openapiclient.BTExportModelParams{AngleTolerance: 123, BatchFlatPatterns: false, ChordTolerance: 123, CloudObjectId: "CloudObjectId_example", CloudStorageAccountId: "CloudStorageAccountId_example", Configuration: "Configuration_example", DeepSearchForForeignData: false, DestinationName: "DestinationName_example", DocumentId: "DocumentId_example", DocumentVersionId: "DocumentVersionId_example", ElementId: "ElementId_example", ElementIds: "ElementIds_example", EmailLink: false, EmailMessage: "EmailMessage_example", EmailSubject: "EmailSubject_example", EmailTo: "EmailTo_example", ExtractToS3: false, FeatureIds: "FeatureIds_example", Flatten: false, Format: "Format_example", FromUserId: "FromUserId_example", Grouping: "Grouping_example", IncludeBendCenterlines: false, IncludeBendLines: false, IncludeCustomProperties: false, IncludeCustomPropertiesData: false, IncludeExportIds: false, IncludeForeignData: false, IncludeItemsData: false, IncludeLinkedDocuments: false, IncludeReleaseManagementData: false, IncludeSketches: false, IncludeStd: false, IsPartingOut: false, LinkDocumentId: "LinkDocumentId_example", LinkDocumentWorkspaceId: "LinkDocumentWorkspaceId_example", MaxFacetWidth: 123, Microversion: "Microversion_example", MinFacetWidth: 123, Mode: "Mode_example", PartIds: "PartIds_example", PartQuery: "PartQuery_example", Password: "Password_example", PasswordRequired: false, Resolution: "Resolution_example", Scale: 123, SendCopyToMe: false, SheetMetalFlat: false, SplinesAsPolylines: false, StoreInDocument: false, TriggerAutoDownload: false, Units: "Units_example", UserId: "UserId_example", ValidForDays: 123, Version: "Version_example", View: "View_example", WorkspaceId: "WorkspaceId_example", ZipSingleFileOutput: false} // BTExportModelParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.Export2Json(context.Background(), did, wv, wvid, eid).BTExportModelParams(bTExportModelParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.Export2Json``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExport2JsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTExportModelParams** | [**BTExportModelParams**](BTExportModelParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+octet-stream;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentMicroversion

> BTMicroversionInfo GetCurrentMicroversion(ctx, did, wv, wvid).Execute()

Get Current Document Microversion

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.GetCurrentMicroversion(context.Background(), did, wv, wvid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetCurrentMicroversion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentMicroversion`: BTMicroversionInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetCurrentMicroversion`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCurrentMicroversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BTMicroversionInfo**](BTMicroversionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocument

> BTDocumentInfo GetDocument(ctx, did).Execute()

Get Document

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
    resp, r, err := api_client.DocumentsApi.GetDocument(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocument`: BTDocumentInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTDocumentInfo**](BTDocumentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentAcl

> BTAclInfo GetDocumentAcl(ctx, did).Execute()

Get Access Control List

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
    resp, r, err := api_client.DocumentsApi.GetDocumentAcl(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetDocumentAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentAcl`: BTAclInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetDocumentAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTAclInfo**](BTAclInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentPermissionSet

> []string GetDocumentPermissionSet(ctx, did).Execute()

Get Document Permissions

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
    resp, r, err := api_client.DocumentsApi.GetDocumentPermissionSet(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetDocumentPermissionSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentPermissionSet`: []string
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetDocumentPermissionSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentPermissionSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentVersions

> []BTVersionInfo GetDocumentVersions(ctx, did).Offset(offset).Limit(limit).Execute()

Get Versions

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
    offset := 987 // int32 |  (optional) (default to 0)
    limit := 987 // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.GetDocumentVersions(context.Background(), did).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetDocumentVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentVersions`: []BTVersionInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetDocumentVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 0]

### Return type

[**[]BTVersionInfo**](BTVersionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentWorkspaces

> []BTWorkspaceInfo GetDocumentWorkspaces(ctx, did).Execute()

Get Workspaces

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
    resp, r, err := api_client.DocumentsApi.GetDocumentWorkspaces(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetDocumentWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentWorkspaces`: []BTWorkspaceInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetDocumentWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BTWorkspaceInfo**](BTWorkspaceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocuments

> BTGlobalTreeNodeListResponse GetDocuments(ctx).Q(q).Filter(filter).Owner(owner).OwnerType(ownerType).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Label(label).Project(project).ParentId(parentId).Execute()

Get Documents

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
    q := "q_example" // string |  (optional) (default to "")
    filter := 987 // int32 |  (optional)
    owner := "owner_example" // string |  (optional) (default to "")
    ownerType := 987 // int32 |  (optional) (default to 1)
    sortColumn := "sortColumn_example" // string |  (optional) (default to "createdAt")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "desc")
    offset := 987 // int32 |  (optional) (default to 0)
    limit := 987 // int32 |  (optional) (default to 20)
    label := "label_example" // string |  (optional)
    project := "project_example" // string |  (optional)
    parentId := "parentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.GetDocuments(context.Background(), ).Q(q).Filter(filter).Owner(owner).OwnerType(ownerType).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Label(label).Project(project).ParentId(parentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocuments`: BTGlobalTreeNodeListResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | [default to &quot;&quot;]
 **filter** | **int32** |  | 
 **owner** | **string** |  | [default to &quot;&quot;]
 **ownerType** | **int32** |  | [default to 1]
 **sortColumn** | **string** |  | [default to &quot;createdAt&quot;]
 **sortOrder** | **string** |  | [default to &quot;desc&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **label** | **string** |  | 
 **project** | **string** |  | 
 **parentId** | **string** |  | 

### Return type

[**BTGlobalTreeNodeListResponse**](BTGlobalTreeNodeListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementsInDocument

> []BTDocumentElementInfo GetElementsInDocument(ctx, did, wvm, wvmid).ElementType(elementType).ElementId(elementId).WithThumbnails(withThumbnails).LinkDocumentId(linkDocumentId).Execute()

Get a list of elements in the workspace, version, or microversion of the document.

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
    elementType := "elementType_example" // string |  (optional) (default to "")
    elementId := "elementId_example" // string |  (optional) (default to "")
    withThumbnails := true // bool |  (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.GetElementsInDocument(context.Background(), did, wvm, wvmid).ElementType(elementType).ElementId(elementId).WithThumbnails(withThumbnails).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetElementsInDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementsInDocument`: []BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetElementsInDocument`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetElementsInDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **elementType** | **string** |  | [default to &quot;&quot;]
 **elementId** | **string** |  | [default to &quot;&quot;]
 **withThumbnails** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 

### Return type

[**[]BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsertables

> BTInsertablesListResponse GetInsertables(ctx, did, wvm, wvmid).BetaCapabilityIds(betaCapabilityIds).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeWires(includeWires).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatures(includeFeatures).IncludeFeatureStudios(includeFeatureStudios).IncludePartStudios(includePartStudios).IncludeBlobs(includeBlobs).IncludeMeshes(includeMeshes).IncludeFlattenedBodies(includeFlattenedBodies).AllowedBlobMimeTypes(allowedBlobMimeTypes).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).Execute()

Insertable List for Document Version.

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
    betaCapabilityIds := []string{"Inner_example"} // []string |  (optional)
    includeParts := true // bool |  (optional) (default to false)
    includeSurfaces := true // bool |  (optional) (default to false)
    includeWires := true // bool |  (optional) (default to false)
    includeSketches := true // bool |  (optional) (default to false)
    includeReferenceFeatures := true // bool |  (optional) (default to false)
    includeAssemblies := true // bool |  (optional) (default to false)
    includeFeatures := true // bool |  (optional) (default to false)
    includeFeatureStudios := true // bool |  (optional) (default to false)
    includePartStudios := true // bool |  (optional) (default to false)
    includeBlobs := true // bool |  (optional) (default to false)
    includeMeshes := true // bool |  (optional) (default to false)
    includeFlattenedBodies := true // bool |  (optional) (default to false)
    allowedBlobMimeTypes := "allowedBlobMimeTypes_example" // string |  (optional) (default to "")
    maxFeatureScriptVersion := 987 // int32 |  (optional) (default to 0)
    includeApplications := true // bool |  (optional) (default to false)
    allowedApplicationMimeTypes := "allowedApplicationMimeTypes_example" // string |  (optional) (default to "")
    includeCompositeParts := true // bool |  (optional) (default to false)
    includeFSTables := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.GetInsertables(context.Background(), did, wvm, wvmid).BetaCapabilityIds(betaCapabilityIds).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeWires(includeWires).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatures(includeFeatures).IncludeFeatureStudios(includeFeatureStudios).IncludePartStudios(includePartStudios).IncludeBlobs(includeBlobs).IncludeMeshes(includeMeshes).IncludeFlattenedBodies(includeFlattenedBodies).AllowedBlobMimeTypes(allowedBlobMimeTypes).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetInsertables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsertables`: BTInsertablesListResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetInsertables`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetInsertablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **betaCapabilityIds** | [**[]string**](string.md) |  | 
 **includeParts** | **bool** |  | [default to false]
 **includeSurfaces** | **bool** |  | [default to false]
 **includeWires** | **bool** |  | [default to false]
 **includeSketches** | **bool** |  | [default to false]
 **includeReferenceFeatures** | **bool** |  | [default to false]
 **includeAssemblies** | **bool** |  | [default to false]
 **includeFeatures** | **bool** |  | [default to false]
 **includeFeatureStudios** | **bool** |  | [default to false]
 **includePartStudios** | **bool** |  | [default to false]
 **includeBlobs** | **bool** |  | [default to false]
 **includeMeshes** | **bool** |  | [default to false]
 **includeFlattenedBodies** | **bool** |  | [default to false]
 **allowedBlobMimeTypes** | **string** |  | [default to &quot;&quot;]
 **maxFeatureScriptVersion** | **int32** |  | [default to 0]
 **includeApplications** | **bool** |  | [default to false]
 **allowedApplicationMimeTypes** | **string** |  | [default to &quot;&quot;]
 **includeCompositeParts** | **bool** |  | [default to false]
 **includeFSTables** | **bool** |  | [default to false]

### Return type

[**BTInsertablesListResponse**](BTInsertablesListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> BTVersionInfo GetVersion(ctx, did, vid).Parents(parents).LinkDocumentId(linkDocumentId).Execute()

Get Version

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
    parents := true // bool |  (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.GetVersion(context.Background(), did, vid).Parents(parents).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.GetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersion`: BTVersionInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.GetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parents** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 

### Return type

[**BTVersionInfo**](BTVersionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeIntoWorkspace

> BTDocumentMergeInfo MergeIntoWorkspace(ctx, did, wid).BTVersionOrWorkspaceInfo(bTVersionOrWorkspaceInfo).Execute()

Merge into workspace

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
    bTVersionOrWorkspaceInfo := openapiclient.BTVersionOrWorkspaceInfo{Id: "Id_example", Type: "Type_example"} // BTVersionOrWorkspaceInfo | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.MergeIntoWorkspace(context.Background(), did, wid, bTVersionOrWorkspaceInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.MergeIntoWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MergeIntoWorkspace`: BTDocumentMergeInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.MergeIntoWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeIntoWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTVersionOrWorkspaceInfo** | [**BTVersionOrWorkspaceInfo**](BTVersionOrWorkspaceInfo.md) |  | 

### Return type

[**BTDocumentMergeInfo**](BTDocumentMergeInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveElementsToDocument

> BTMoveElementInfo MoveElementsToDocument(ctx, did, wid).BTMoveElementParams(bTMoveElementParams).Execute()

Move Elements

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
    bTMoveElementParams := openapiclient.BTMoveElementParams{AnchorElementId: "AnchorElementId_example", Description: "Description_example", ElementOriginalToNewMap: map[string]string{ "Key" = "Value" }, Elements: []string{"Elements_example"), GenerateUnknownMessages: false, ImportData: []string{123), IsCopy: false, IsDeepCopy: false, IsGroupAnchor: false, IsNewDocument: false, IsPublic: false, IsSelectivePartOut: false, Name: "Name_example", NeedNewVersion: false, OwnerEmail: "OwnerEmail_example", OwnerId: "OwnerId_example", OwnerType: 123, ParentId: "ParentId_example", ProjectId: "ProjectId_example", SourceDocumentId: "SourceDocumentId_example", SourceWorkspaceId: "SourceWorkspaceId_example", Tags: []string{"Tags_example"), TargetDocumentId: "TargetDocumentId_example", TargetWorkspaceId: "TargetWorkspaceId_example", VersionName: "VersionName_example"} // BTMoveElementParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.MoveElementsToDocument(context.Background(), did, wid, bTMoveElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.MoveElementsToDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveElementsToDocument`: BTMoveElementInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.MoveElementsToDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveElementsToDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTMoveElementParams** | [**BTMoveElementParams**](BTMoveElementParams.md) |  | 

### Return type

[**BTMoveElementInfo**](BTMoveElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFromHistory

> RestoreFromHistory(ctx, did, wid, vm, vmid).Execute()

Restore version or microversion to workspace.

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
    vm := "vm_example" // string | 
    vmid := "vmid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.RestoreFromHistory(context.Background(), did, wid, vm, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.RestoreFromHistory``: %v\n", err)
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
**vm** | **string** |  | 
**vmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFromHistoryRequest struct via the builder pattern


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


## Search

> Search(ctx).BTDocumentSearchParams(bTDocumentSearchParams).Execute()



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
    bTDocumentSearchParams := openapiclient.BTDocumentSearchParams{DocumentFilter: 123, FoundIn: "FoundIn_example", Limit: 123, Offset: 123, OwnerId: "OwnerId_example", ParentId: "ParentId_example", RawQuery: "RawQuery_example", SortColumn: "SortColumn_example", SortOrder: "SortOrder_example", Type: "Type_example", When: "When_example"} // BTDocumentSearchParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.Search(context.Background(), bTDocumentSearchParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTDocumentSearchParams** | [**BTDocumentSearchParams**](BTDocumentSearchParams.md) |  | 

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


## ShareDocument

> BTAclInfo ShareDocument(ctx, did).BTShareParams(bTShareParams).Execute()

Share Document

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
    bTShareParams := openapiclient.BTShareParams{DocumentId: "DocumentId_example", ElementId: "ElementId_example", EncodedConfiguration: "EncodedConfiguration_example", Entries: []BTShareEntryParams{openapiclient.BTShareEntryParams{ApplicationId: "ApplicationId_example", CompanyId: "CompanyId_example", Email: "Email_example", EntryType: 123, TeamId: "TeamId_example", UserId: "UserId_example"}), FolderId: "FolderId_example", Message: "Message_example", Permission: int64(123), PermissionSet: []string{"PermissionSet_example"), Update: false, WorkspaceId: "WorkspaceId_example"} // BTShareParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.ShareDocument(context.Background(), did, bTShareParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.ShareDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShareDocument`: BTAclInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.ShareDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTShareParams** | [**BTShareParams**](BTShareParams.md) |  | 

### Return type

[**BTAclInfo**](BTAclInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncApplicationElements

> SyncApplicationElements(ctx, did, wid).ApplicationElementIds(applicationElementIds).Description(description).Execute()

Sync Application Elements

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
    applicationElementIds := []string{"Inner_example"} // []string | 
    description := "description_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.SyncApplicationElements(context.Background(), did, wid, applicationElementIds).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.SyncApplicationElements``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiSyncApplicationElementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationElementIds** | [**[]string**](string.md) |  | 
 **description** | **string** |  | 

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


## UnShareDocument

> UnShareDocument(ctx, did, eid).EntryType(entryType).Execute()

Unshare Document

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
    entryType := 987 // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.UnShareDocument(context.Background(), did, eid).EntryType(entryType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.UnShareDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnShareDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entryType** | **int32** |  | [default to 0]

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


## UpdateDocumentAttributes

> UpdateDocumentAttributes(ctx, did).BTDocumentParams(bTDocumentParams).Execute()

Update Document Attributes.

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
    bTDocumentParams := openapiclient.BTDocumentParams{BetaCapabilityIds: []string{"BetaCapabilityIds_example"), Description: "Description_example", GenerateUnknownMessages: false, IsEmptyContent: false, IsPublic: false, Name: "Name_example", NotRevisionManaged: false, OwnerEmail: "OwnerEmail_example", OwnerId: "OwnerId_example", OwnerType: 123, ParentId: "ParentId_example", ProjectId: "ProjectId_example", Tags: []string{"Tags_example")} // BTDocumentParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.UpdateDocumentAttributes(context.Background(), did, bTDocumentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.UpdateDocumentAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTDocumentParams** | [**BTDocumentParams**](BTDocumentParams.md) |  | 

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


## UpdateExternalReferencesToLatestDocuments

> BTLinkToLatestDocumentInfo UpdateExternalReferencesToLatestDocuments(ctx, did, wid, eid).BTLinkToLatestDocumentParams(bTLinkToLatestDocumentParams).Execute()

Update External References to Latest

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
    bTLinkToLatestDocumentParams := openapiclient.BTLinkToLatestDocumentParams{Elements: []string{"Elements_example")} // BTLinkToLatestDocumentParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DocumentsApi.UpdateExternalReferencesToLatestDocuments(context.Background(), did, wid, eid).BTLinkToLatestDocumentParams(bTLinkToLatestDocumentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsApi.UpdateExternalReferencesToLatestDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExternalReferencesToLatestDocuments`: BTLinkToLatestDocumentInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentsApi.UpdateExternalReferencesToLatestDocuments`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateExternalReferencesToLatestDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTLinkToLatestDocumentParams** | [**BTLinkToLatestDocumentParams**](BTLinkToLatestDocumentParams.md) |  | 

### Return type

[**BTLinkToLatestDocumentInfo**](BTLinkToLatestDocumentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

