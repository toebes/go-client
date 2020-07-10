# \TranslationsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranslation**](TranslationsApi.md#CreateTranslation) | **Post** /api/translations/d/{did}/w/{wid} | Create translation from upload.
[**DeleteTranslation**](TranslationsApi.md#DeleteTranslation) | **Delete** /api/translations/{tid} | 
[**GetAllTranslatorFormats**](TranslationsApi.md#GetAllTranslatorFormats) | **Get** /api/translations/translationformats | 
[**GetDocumentTranslations**](TranslationsApi.md#GetDocumentTranslations) | **Get** /api/translations/d/{did} | 
[**GetTranslation**](TranslationsApi.md#GetTranslation) | **Get** /api/translations/{tid} | 



## CreateTranslation

> BTTranslationRequestInfo CreateTranslation(ctx, did, wid).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).File(file).FileBodyWithDetails(fileBodyWithDetails).FileContentLength(fileContentLength).FileDetail(fileDetail).FlattenAssemblies(flattenAssemblies).FormatName(formatName).IsyAxisIsUp(isyAxisIsUp).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).OwnerType(ownerType).ParentId(parentId).ProjectId(projectId).Public(public).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).Execute()

Create translation from upload.

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
    resp, r, err := api_client.TranslationsApi.CreateTranslation(context.Background(), did, wid).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).File(file).FileBodyWithDetails(fileBodyWithDetails).FileContentLength(fileContentLength).FileDetail(fileDetail).FlattenAssemblies(flattenAssemblies).FormatName(formatName).IsyAxisIsUp(isyAxisIsUp).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).OwnerType(ownerType).ParentId(parentId).ProjectId(projectId).Public(public).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationsApi.CreateTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationsApi.CreateTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranslationRequest struct via the builder pattern


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

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTranslation

> DeleteTranslation(ctx, tid).Execute()



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
    tid := "tid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TranslationsApi.DeleteTranslation(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationsApi.DeleteTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTranslationRequest struct via the builder pattern


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


## GetAllTranslatorFormats

> []BTModelFormatFullInfo GetAllTranslatorFormats(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TranslationsApi.GetAllTranslatorFormats(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationsApi.GetAllTranslatorFormats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTranslatorFormats`: []BTModelFormatFullInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationsApi.GetAllTranslatorFormats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTranslatorFormatsRequest struct via the builder pattern


### Return type

[**[]BTModelFormatFullInfo**](BTModelFormatFullInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentTranslations

> BTListResponseBTTranslationRequestInfo GetDocumentTranslations(ctx, did).Offset(offset).Limit(limit).Execute()



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
    limit := 987 // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TranslationsApi.GetDocumentTranslations(context.Background(), did).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationsApi.GetDocumentTranslations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentTranslations`: BTListResponseBTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationsApi.GetDocumentTranslations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTTranslationRequestInfo**](BTListResponseBTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTranslation

> BTTranslationRequestInfo GetTranslation(ctx, tid).Execute()



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
    tid := "tid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TranslationsApi.GetTranslation(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationsApi.GetTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationsApi.GetTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

