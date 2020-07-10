# \BlobElementsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlobTranslation**](BlobElementsApi.md#CreateBlobTranslation) | **Post** /api/blobelements/d/{did}/{wv}/{wvid}/e/{eid}/translations | Create Translation.
[**DownloadFileWorkspace**](BlobElementsApi.md#DownloadFileWorkspace) | **Get** /api/blobelements/d/{did}/w/{wid}/e/{eid} | Download File From Blob Element.
[**UpdateUnits**](BlobElementsApi.md#UpdateUnits) | **Post** /api/blobelements/d/{did}/w/{wid}/e/{eid}/units | Update Mesh Units.
[**UploadFileCreateElement**](BlobElementsApi.md#UploadFileCreateElement) | **Post** /api/blobelements/d/{did}/w/{wid} | Upload file to new element.
[**UploadFileUpdateElement**](BlobElementsApi.md#UploadFileUpdateElement) | **Post** /api/blobelements/d/{did}/w/{wid}/e/{eid} | Update Blob Element.



## CreateBlobTranslation

> BTTranslationRequestInfo CreateBlobTranslation(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Create Translation.

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
    bTTranslateFormatParams := openapiclient.BTTranslateFormatParams{AllowFaultyParts: false, AngularTolerance: 123, BlobElementId: "BlobElementId_example", BlobMicroversionId: "BlobMicroversionId_example", CloudObjectId: "CloudObjectId_example", CloudStorageAccountId: "CloudStorageAccountId_example", ColorMethod: "ColorMethod_example", Configuration: "Configuration_example", ConnectionId: "ConnectionId_example", CreateComposite: false, CurrentSheetOnly: false, DestinationName: "DestinationName_example", DistanceTolerance: 123, ElementId: "ElementId_example", EmailLink: false, EmailMessage: "EmailMessage_example", EmailSubject: "EmailSubject_example", EmailTo: []string{"EmailTo_example"), ExtractAssemblyHierarchy: false, Flatten: false, FlattenAssemblies: false, ForeignId: "ForeignId_example", FormatName: "FormatName_example", FromUserId: "FromUserId_example", GetyAxisIsUp: false, Grouping: false, ImageHeight: 123, ImageWidth: 123, ImportInBackground: false, ImportWithinDocument: false, IncludeExportIds: false, JoinAdjacentSurfaces: false, LinkDocumentId: "LinkDocumentId_example", LinkDocumentWorkspaceId: "LinkDocumentWorkspaceId_example", MaximumChordLength: 123, NotifyUser: false, OriginalForeignId: "OriginalForeignId_example", ParentId: "ParentId_example", PartIds: "PartIds_example", Password: "Password_example", PasswordRequired: false, ProcessedForeignId: "ProcessedForeignId_example", ProjectId: "ProjectId_example", SelectablePdfText: false, SendCopyToMe: false, ShowOverriddenDimensions: false, SourceName: "SourceName_example", SpecifyUnits: false, SplinesAsPolylines: false, SplitAssembliesIntoMultipleDocuments: false, StoreInDocument: false, TextAsGeometry: false, TriggerAutoDownload: false, Unit: "Unit_example", UploadId: "UploadId_example", ValidForDays: 123, VersionString: "VersionString_example"} // BTTranslateFormatParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlobElementsApi.CreateBlobTranslation(context.Background(), did, wv, wvid, eid, bTTranslateFormatParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementsApi.CreateBlobTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBlobTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementsApi.CreateBlobTranslation`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateBlobTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTTranslateFormatParams** | [**BTTranslateFormatParams**](BTTranslateFormatParams.md) |  | 

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFileWorkspace

> *os.File DownloadFileWorkspace(ctx, did, wid, eid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).LinkDocumentId(linkDocumentId).Execute()

Download File From Blob Element.

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
    contentDisposition := "contentDisposition_example" // string |  (optional)
    ifNoneMatch := "ifNoneMatch_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlobElementsApi.DownloadFileWorkspace(context.Background(), did, wid, eid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementsApi.DownloadFileWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFileWorkspace`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `BlobElementsApi.DownloadFileWorkspace`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDownloadFileWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentDisposition** | **string** |  | 
 **ifNoneMatch** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUnits

> BTDocumentElementProcessingInfo UpdateUnits(ctx, did, eid, wid).BTUpdateMeshUnitsParams(bTUpdateMeshUnitsParams).Execute()

Update Mesh Units.

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
    bTUpdateMeshUnitsParams := openapiclient.BTUpdateMeshUnitsParams{Units: "Units_example"} // BTUpdateMeshUnitsParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlobElementsApi.UpdateUnits(context.Background(), did, eid, wid, bTUpdateMeshUnitsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementsApi.UpdateUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUnits`: BTDocumentElementProcessingInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementsApi.UpdateUnits`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTUpdateMeshUnitsParams** | [**BTUpdateMeshUnitsParams**](BTUpdateMeshUnitsParams.md) |  | 

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileCreateElement

> BTDocumentElementProcessingInfo UploadFileCreateElement(ctx, did, wid).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).File(file).FileBodyWithDetails(fileBodyWithDetails).FileContentLength(fileContentLength).FileDetail(fileDetail).FlattenAssemblies(flattenAssemblies).FormatName(formatName).IsyAxisIsUp(isyAxisIsUp).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).OwnerType(ownerType).ParentId(parentId).ProjectId(projectId).Public(public).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).Execute()

Upload file to new element.

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
    allowFaultyParts := true // bool |  (optional)
    createComposite := true // bool |  (optional)
    createDrawingIfPossible := true // bool |  (optional)
    encodedFilename := "encodedFilename_example" // string |  (optional)
    extractAssemblyHierarchy := true // bool |  (optional)
    file := 987 // *os.File |  (optional)
    fileBodyWithDetails := openapiclient.FormDataBodyPart{ContentDisposition: openapiclient.ContentDisposition{CreationDate: "TODO", FileName: "FileName_example", ModificationDate: "TODO", Parameters: map[string]string{ "Key" = "Value" }, ReadDate: "TODO", Size: int64(123), Type: "Type_example"}, Entity: "TODO", FormDataContentDisposition: openapiclient.FormDataContentDisposition{CreationDate: "TODO", FileName: "FileName_example", ModificationDate: "TODO", Name: "Name_example", Parameters: map[string]string{ "Key" = "Value" }, ReadDate: "TODO", Size: int64(123), Type: "Type_example"}, Headers: map[string]string{ "Key" = "Value" }, MediaType: openapiclient.BodyPart_mediaType{Type: "Type_example", Subtype: "Subtype_example", Parameters: map[string]string{ "Key" = "Value" }, WildcardType: false, WildcardSubtype: false}, MessageBodyWorkers: "TODO", Name: "Name_example", ParameterizedHeaders: map[string]string{ "Key" = "Value" }, Parent: openapiclient.MultiPart{BodyParts: []BodyPart{openapiclient.BodyPart{ContentDisposition: openapiclient.ContentDisposition{CreationDate: "TODO", FileName: "FileName_example", ModificationDate: "TODO", Parameters: map[string]string{ "Key" = "Value" }, ReadDate: "TODO", Size: int64(123), Type: "Type_example"}, Entity: "TODO", Headers: map[string]string{ "Key" = "Value" }, MediaType: openapiclient.BodyPart_mediaType{Type: "Type_example", Subtype: "Subtype_example", Parameters: map[string]string{ "Key" = "Value" }, WildcardType: false, WildcardSubtype: false}, MessageBodyWorkers: "TODO", ParameterizedHeaders: map[string]string{ "Key" = "Value" }, Parent: openapiclient.MultiPart{BodyParts: []BodyPart{openapiclient.BodyPart{ContentDisposition: , Entity: "TODO", Headers: map[string]string{ "Key" = "Value" }, MediaType: , MessageBodyWorkers: "TODO", ParameterizedHeaders: map[string]string{ "Key" = "Value" }, Parent: , Providers: "TODO"}), ContentDisposition: , Entity: "TODO", Headers: map[string]string{ "Key" = "Value" }, MediaType: , MessageBodyWorkers: "TODO", ParameterizedHeaders: map[string]string{ "Key" = "Value" }, Parent: , Providers: "TODO"}, Providers: "TODO"}), ContentDisposition: , Entity: "TODO", Headers: map[string]string{ "Key" = "Value" }, MediaType: , MessageBodyWorkers: "TODO", ParameterizedHeaders: map[string]string{ "Key" = "Value" }, Parent: , Providers: "TODO"}, Providers: "TODO", Simple: false, Value: "Value_example"} // FormDataBodyPart |  (optional)
    fileContentLength := 987 // int64 |  (optional)
    fileDetail := openapiclient.FormDataContentDisposition{CreationDate: "TODO", FileName: "FileName_example", ModificationDate: "TODO", Name: "Name_example", Parameters: map[string]string{ "Key" = "Value" }, ReadDate: "TODO", Size: int64(123), Type: "Type_example"} // FormDataContentDisposition |  (optional)
    flattenAssemblies := true // bool |  (optional)
    formatName := "formatName_example" // string |  (optional)
    isyAxisIsUp := true // bool |  (optional)
    joinAdjacentSurfaces := true // bool |  (optional)
    locationElementId := "locationElementId_example" // string |  (optional)
    locationGroupId := "locationGroupId_example" // string |  (optional)
    locationPosition := 987 // int32 |  (optional)
    notifyUser := true // bool |  (optional)
    ownerId := "ownerId_example" // string |  (optional)
    ownerType := "ownerType_example" // string |  (optional)
    parentId := "parentId_example" // string |  (optional)
    projectId := "projectId_example" // string |  (optional)
    public := true // bool |  (optional)
    splitAssembliesIntoMultipleDocuments := true // bool |  (optional)
    storeInDocument := true // bool |  (optional)
    translate := true // bool |  (optional)
    unit := "unit_example" // string |  (optional)
    uploadId := "uploadId_example" // string |  (optional)
    versionString := "versionString_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlobElementsApi.UploadFileCreateElement(context.Background(), did, wid).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).File(file).FileBodyWithDetails(fileBodyWithDetails).FileContentLength(fileContentLength).FileDetail(fileDetail).FlattenAssemblies(flattenAssemblies).FormatName(formatName).IsyAxisIsUp(isyAxisIsUp).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).OwnerType(ownerType).ParentId(parentId).ProjectId(projectId).Public(public).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementsApi.UploadFileCreateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFileCreateElement`: BTDocumentElementProcessingInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementsApi.UploadFileCreateElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileCreateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allowFaultyParts** | **bool** |  | 
 **createComposite** | **bool** |  | 
 **createDrawingIfPossible** | **bool** |  | 
 **encodedFilename** | **string** |  | 
 **extractAssemblyHierarchy** | **bool** |  | 
 **file** | ***os.File** |  | 
 **fileBodyWithDetails** | [**FormDataBodyPart**](FormDataBodyPart.md) |  | 
 **fileContentLength** | **int64** |  | 
 **fileDetail** | [**FormDataContentDisposition**](FormDataContentDisposition.md) |  | 
 **flattenAssemblies** | **bool** |  | 
 **formatName** | **string** |  | 
 **isyAxisIsUp** | **bool** |  | 
 **joinAdjacentSurfaces** | **bool** |  | 
 **locationElementId** | **string** |  | 
 **locationGroupId** | **string** |  | 
 **locationPosition** | **int32** |  | 
 **notifyUser** | **bool** |  | 
 **ownerId** | **string** |  | 
 **ownerType** | **string** |  | 
 **parentId** | **string** |  | 
 **projectId** | **string** |  | 
 **public** | **bool** |  | 
 **splitAssembliesIntoMultipleDocuments** | **bool** |  | 
 **storeInDocument** | **bool** |  | 
 **translate** | **bool** |  | 
 **unit** | **string** |  | 
 **uploadId** | **string** |  | 
 **versionString** | **string** |  | 

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileUpdateElement

> BTDocumentElementProcessingInfo UploadFileUpdateElement(ctx, did, eid, wid).ParentChangeId(parentChangeId).Execute()

Update Blob Element.

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
    parentChangeId := "parentChangeId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BlobElementsApi.UploadFileUpdateElement(context.Background(), did, eid, wid).ParentChangeId(parentChangeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementsApi.UploadFileUpdateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFileUpdateElement`: BTDocumentElementProcessingInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementsApi.UploadFileUpdateElement`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUploadFileUpdateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **parentChangeId** | **string** |  | 

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

