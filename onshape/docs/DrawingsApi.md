# \DrawingsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrawingAppElement**](DrawingsApi.md#CreateDrawingAppElement) | **Post** /api/drawings/create | 
[**CreateDrawingTranslation**](DrawingsApi.md#CreateDrawingTranslation) | **Post** /api/drawings/d/{did}/{wv}/{wvid}/e/{eid}/translations | Create Drawing translation
[**GetDrawingTranslatorFormats**](DrawingsApi.md#GetDrawingTranslatorFormats) | **Get** /api/drawings/d/{did}/w/{wid}/e/{eid}/translationformats | 



## CreateDrawingAppElement

> BTDocumentElementInfo CreateDrawingAppElement(ctx).BTDrawingParams(bTDrawingParams).Execute()



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
    bTDrawingParams := openapiclient.BTDrawingParams{Border: false, ComputeIntersection: false, DecimalSeparator: "DecimalSeparator_example", DocumentId: "DocumentId_example", DocumentMicroversionId: "DocumentMicroversionId_example", DrawingName: "DrawingName_example", ElementConfiguration: "ElementConfiguration_example", ElementId: "ElementId_example", ElementMicroversionId: "ElementMicroversionId_example", ExternalDocumentId: "ExternalDocumentId_example", ExternalDocumentVersionId: "ExternalDocumentVersionId_example", HiddenLines: "HiddenLines_example", IncludeSurfaces: false, IsFlattenedPart: false, IsSketchOnly: false, IsSurface: false, Language: "Language_example", Location: openapiclient.BTElementLocationParams{ElementId: "ElementId_example", GroupId: "GroupId_example", Position: 123}, ModelType: "ModelType_example", NumberHorizontalZones: 123, NumberVerticalZones: 123, PartId: "PartId_example", PartNumber: "PartNumber_example", PartQuery: "PartQuery_example", Projection: "Projection_example", PureSketch: false, QualityOption: "QualityOption_example", ReferenceType: 123, ReferenceTypeEnum: "ReferenceTypeEnum_example", Revision: "Revision_example", ShowCutGeomOnly: false, SimplificationOption: "SimplificationOption_example", SimplificationThreshold: 123, Size: "Size_example", SketchIds: []string{"SketchIds_example"), Standard: "Standard_example", StartZones: "StartZones_example", TemplateArgs: []string{"TemplateArgs_example"), TemplateDocumentId: "TemplateDocumentId_example", TemplateElementId: "TemplateElementId_example", TemplateName: "TemplateName_example", TemplateVersionId: "TemplateVersionId_example", TemplateWorkspaceId: "TemplateWorkspaceId_example", Titleblock: false, Units: "Units_example", Views: "Views_example", WorkspaceId: "WorkspaceId_example"} // BTDrawingParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DrawingsApi.CreateDrawingAppElement(context.Background(), bTDrawingParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingsApi.CreateDrawingAppElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDrawingAppElement`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `DrawingsApi.CreateDrawingAppElement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDrawingAppElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTDrawingParams** | [**BTDrawingParams**](BTDrawingParams.md) |  | 

### Return type

[**BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDrawingTranslation

> BTTranslationRequestInfo CreateDrawingTranslation(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Create Drawing translation

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
    resp, r, err := api_client.DrawingsApi.CreateDrawingTranslation(context.Background(), did, wv, wvid, eid, bTTranslateFormatParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingsApi.CreateDrawingTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDrawingTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `DrawingsApi.CreateDrawingTranslation`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateDrawingTranslationRequest struct via the builder pattern


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


## GetDrawingTranslatorFormats

> []BTModelFormatInfo GetDrawingTranslatorFormats(ctx, did, wid, eid).Execute()



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
    resp, r, err := api_client.DrawingsApi.GetDrawingTranslatorFormats(context.Background(), did, wid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingsApi.GetDrawingTranslatorFormats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDrawingTranslatorFormats`: []BTModelFormatInfo
    fmt.Fprintf(os.Stdout, "Response from `DrawingsApi.GetDrawingTranslatorFormats`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDrawingTranslatorFormatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

