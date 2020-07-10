# \DefaultApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowableTestObject**](DefaultApi.md#CreateWorkflowableTestObject) | **Post** /api/workflowabletestobject/testobject/{wfid} | 
[**DeleteAssociativeData**](DefaultApi.md#DeleteAssociativeData) | **Delete** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata | 
[**GetAssociativeData**](DefaultApi.md#GetAssociativeData) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata | 
[**GetLatestInDocument**](DefaultApi.md#GetLatestInDocument) | **Get** /api/insertables/d/{did}/latest | insertables for a document
[**GetMetadataSchema**](DefaultApi.md#GetMetadataSchema) | **Get** /api/metadataschema | 
[**GetProperties**](DefaultApi.md#GetProperties) | **Get** /api/metadataschema/properties | 
[**GetPropertyInfo**](DefaultApi.md#GetPropertyInfo) | **Get** /api/metadataschema/propertyinfo/{pid} | 
[**GetSchema**](DefaultApi.md#GetSchema) | **Get** /api/metadataschema/{sid} | 
[**GetSketchBoundingBoxes**](DefaultApi.md#GetSketchBoundingBoxes) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches/{sid}/boundingboxes | 
[**GetSketchInfo**](DefaultApi.md#GetSketchInfo) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches | 
[**GetTessellatedEntities**](DefaultApi.md#GetTessellatedEntities) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches/{sid}/tessellatedentities | 
[**GetWorkflowableTestObject**](DefaultApi.md#GetWorkflowableTestObject) | **Get** /api/workflowabletestobject/{oid} | 
[**PostAssociativeData**](DefaultApi.md#PostAssociativeData) | **Post** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata | 
[**TransitionWorkflowableTestObject**](DefaultApi.md#TransitionWorkflowableTestObject) | **Post** /api/workflowabletestobject/{oid}/{transition} | 
[**UpdateWorkflowableTestObject**](DefaultApi.md#UpdateWorkflowableTestObject) | **Post** /api/workflowabletestobject/{oid} | 



## CreateWorkflowableTestObject

> BTWorkflowableTestObjectInfo CreateWorkflowableTestObject(ctx, wfid).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWorkflowableTestObject(context.Background(), wfid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWorkflowableTestObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowableTestObject`: BTWorkflowableTestObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWorkflowableTestObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssociativeData

> BTAppElementBasicInfo DeleteAssociativeData(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ParentChangeId(parentChangeId).AssociativeDataId(associativeDataId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).Execute()



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
    transactionId := "transactionId_example" // string |  (optional) (default to "")
    parentChangeId := "parentChangeId_example" // string |  (optional) (default to "")
    associativeDataId := []string{"Inner_example"} // []string |  (optional)
    elementId := "elementId_example" // string |  (optional) (default to "")
    viewId := "viewId_example" // string |  (optional) (default to "")
    microversionId := "microversionId_example" // string |  (optional) (default to "")
    documentMicroversion := "documentMicroversion_example" // string |  (optional) (default to "")
    deterministicId := "deterministicId_example" // string |  (optional) (default to "")
    featureId := "featureId_example" // string |  (optional) (default to "")
    entityId := "entityId_example" // string |  (optional) (default to "")
    occurrenceId := "occurrenceId_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAssociativeData(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ParentChangeId(parentChangeId).AssociativeDataId(associativeDataId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAssociativeData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAssociativeData`: BTAppElementBasicInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteAssociativeData`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDeleteAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | [default to &quot;&quot;]
 **parentChangeId** | **string** |  | [default to &quot;&quot;]
 **associativeDataId** | [**[]string**](string.md) |  | 
 **elementId** | **string** |  | [default to &quot;&quot;]
 **viewId** | **string** |  | [default to &quot;&quot;]
 **microversionId** | **string** |  | [default to &quot;&quot;]
 **documentMicroversion** | **string** |  | [default to &quot;&quot;]
 **deterministicId** | **string** |  | [default to &quot;&quot;]
 **featureId** | **string** |  | [default to &quot;&quot;]
 **entityId** | **string** |  | [default to &quot;&quot;]
 **occurrenceId** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTAppElementBasicInfo**](BTAppElementBasicInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssociativeData

> BTAppAssociativeDataInfoArray GetAssociativeData(ctx, did, wvm, wvmid, eid).TransactionId(transactionId).ChangeId(changeId).AssociativeDataId(associativeDataId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).ReturnIdTags(returnIdTags).Execute()



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
    transactionId := "transactionId_example" // string |  (optional) (default to "")
    changeId := "changeId_example" // string |  (optional) (default to "")
    associativeDataId := []string{"Inner_example"} // []string |  (optional)
    elementId := "elementId_example" // string |  (optional) (default to "")
    viewId := "viewId_example" // string |  (optional) (default to "")
    microversionId := "microversionId_example" // string |  (optional) (default to "")
    documentMicroversion := "documentMicroversion_example" // string |  (optional) (default to "")
    deterministicId := "deterministicId_example" // string |  (optional) (default to "")
    featureId := "featureId_example" // string |  (optional) (default to "")
    entityId := "entityId_example" // string |  (optional) (default to "")
    occurrenceId := "occurrenceId_example" // string |  (optional) (default to "")
    returnIdTags := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAssociativeData(context.Background(), did, wvm, wvmid, eid).TransactionId(transactionId).ChangeId(changeId).AssociativeDataId(associativeDataId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).ReturnIdTags(returnIdTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAssociativeData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssociativeData`: BTAppAssociativeDataInfoArray
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAssociativeData`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | [default to &quot;&quot;]
 **changeId** | **string** |  | [default to &quot;&quot;]
 **associativeDataId** | [**[]string**](string.md) |  | 
 **elementId** | **string** |  | [default to &quot;&quot;]
 **viewId** | **string** |  | [default to &quot;&quot;]
 **microversionId** | **string** |  | [default to &quot;&quot;]
 **documentMicroversion** | **string** |  | [default to &quot;&quot;]
 **deterministicId** | **string** |  | [default to &quot;&quot;]
 **featureId** | **string** |  | [default to &quot;&quot;]
 **entityId** | **string** |  | [default to &quot;&quot;]
 **occurrenceId** | **string** |  | [default to &quot;&quot;]
 **returnIdTags** | **bool** |  | [default to false]

### Return type

[**BTAppAssociativeDataInfoArray**](BTAppAssociativeDataInfoArray.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestInDocument

> BTListResponseBTInsertableInfo GetLatestInDocument(ctx, did).BetaCapabilityIds(betaCapabilityIds).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeWires(includeWires).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatures(includeFeatures).IncludeFeatureStudios(includeFeatureStudios).IncludePartStudios(includePartStudios).IncludeBlobs(includeBlobs).IncludeMeshes(includeMeshes).IncludeFlattenedBodies(includeFlattenedBodies).AllowedBlobMimeTypes(allowedBlobMimeTypes).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).Execute()

insertables for a document

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
    resp, r, err := api_client.DefaultApi.GetLatestInDocument(context.Background(), did).BetaCapabilityIds(betaCapabilityIds).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeWires(includeWires).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatures(includeFeatures).IncludeFeatureStudios(includeFeatureStudios).IncludePartStudios(includePartStudios).IncludeBlobs(includeBlobs).IncludeMeshes(includeMeshes).IncludeFlattenedBodies(includeFlattenedBodies).AllowedBlobMimeTypes(allowedBlobMimeTypes).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLatestInDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestInDocument`: BTListResponseBTInsertableInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLatestInDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestInDocumentRequest struct via the builder pattern


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

[**BTListResponseBTInsertableInfo**](BTListResponseBTInsertableInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataSchema

> BTMetadataSchemaInfo GetMetadataSchema(ctx).ObjectType(objectType).OwnerId(ownerId).DocumentId(documentId).OwnerType(ownerType).Execute()



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
    objectType := 987 // int32 | 
    ownerId := "ownerId_example" // string |  (optional)
    documentId := "documentId_example" // string |  (optional)
    ownerType := 987 // int32 |  (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetMetadataSchema(context.Background(), objectType).OwnerId(ownerId).DocumentId(documentId).OwnerType(ownerType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMetadataSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataSchema`: BTMetadataSchemaInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMetadataSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectType** | **int32** |  | 
 **ownerId** | **string** |  | 
 **documentId** | **string** |  | 
 **ownerType** | **int32** |  | [default to 1]

### Return type

[**BTMetadataSchemaInfo**](BTMetadataSchemaInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProperties

> BTListResponseBTMetadataPropertySummaryInfo GetProperties(ctx).SchemaId(schemaId).OwnerId(ownerId).DocumentId(documentId).OwnerType(ownerType).ObjectType(objectType).Strict(strict).ActiveOnly(activeOnly).Offset(offset).Limit(limit).Execute()



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
    schemaId := "schemaId_example" // string |  (optional)
    ownerId := "ownerId_example" // string |  (optional)
    documentId := "documentId_example" // string |  (optional)
    ownerType := 987 // int32 |  (optional) (default to 1)
    objectType := 987 // int32 |  (optional)
    strict := true // bool |  (optional) (default to false)
    activeOnly := true // bool |  (optional) (default to false)
    offset := 987 // int32 |  (optional) (default to 0)
    limit := 987 // int32 |  (optional) (default to 200)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetProperties(context.Background(), ).SchemaId(schemaId).OwnerId(ownerId).DocumentId(documentId).OwnerType(ownerType).ObjectType(objectType).Strict(strict).ActiveOnly(activeOnly).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProperties`: BTListResponseBTMetadataPropertySummaryInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaId** | **string** |  | 
 **ownerId** | **string** |  | 
 **documentId** | **string** |  | 
 **ownerType** | **int32** |  | [default to 1]
 **objectType** | **int32** |  | 
 **strict** | **bool** |  | [default to false]
 **activeOnly** | **bool** |  | [default to false]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 200]

### Return type

[**BTListResponseBTMetadataPropertySummaryInfo**](BTListResponseBTMetadataPropertySummaryInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPropertyInfo

> BTMetadataPropertyInfo GetPropertyInfo(ctx, pid).DocumentId(documentId).SchemaId(schemaId).OwnerId(ownerId).OwnerType(ownerType).ObjectType(objectType).Execute()



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
    pid := "pid_example" // string | 
    documentId := "documentId_example" // string |  (optional)
    schemaId := "schemaId_example" // string |  (optional)
    ownerId := "ownerId_example" // string |  (optional)
    ownerType := 987 // int32 |  (optional) (default to 1)
    objectType := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetPropertyInfo(context.Background(), pid).DocumentId(documentId).SchemaId(schemaId).OwnerId(ownerId).OwnerType(ownerType).ObjectType(objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPropertyInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPropertyInfo`: BTMetadataPropertyInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPropertyInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documentId** | **string** |  | 
 **schemaId** | **string** |  | 
 **ownerId** | **string** |  | 
 **ownerType** | **int32** |  | [default to 1]
 **objectType** | **int32** |  | 

### Return type

[**BTMetadataPropertyInfo**](BTMetadataPropertyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchema

> BTMetadataSchemaInfo GetSchema(ctx, sid).DocumentId(documentId).Execute()



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
    sid := "sid_example" // string | 
    documentId := "documentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSchema(context.Background(), sid).DocumentId(documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchema`: BTMetadataSchemaInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documentId** | **string** |  | 

### Return type

[**BTMetadataSchemaInfo**](BTMetadataSchemaInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSketchBoundingBoxes

> BTBoundingBoxInfo GetSketchBoundingBoxes(ctx, did, wvm, wvmid, eid, sid).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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
    sid := "sid_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSketchBoundingBoxes(context.Background(), did, wvm, wvmid, eid, sid).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSketchBoundingBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSketchBoundingBoxes`: BTBoundingBoxInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSketchBoundingBoxes`: %v\n", resp)
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
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSketchBoundingBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTBoundingBoxInfo**](BTBoundingBoxInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSketchInfo

> GetSketchInfo(ctx, did, wvm, wvmid, eid).Configuration(configuration).SketchId(sketchId).Output3D(output3D).CurvePoints(curvePoints).IncludeGeometry(includeGeometry).LinkDocumentId(linkDocumentId).Execute()



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
    sketchId := []string{"Inner_example"} // []string |  (optional)
    output3D := true // bool |  (optional) (default to false)
    curvePoints := true // bool |  (optional) (default to false)
    includeGeometry := true // bool |  (optional) (default to true)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetSketchInfo(context.Background(), did, wvm, wvmid, eid).Configuration(configuration).SketchId(sketchId).Output3D(output3D).CurvePoints(curvePoints).IncludeGeometry(includeGeometry).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSketchInfo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetSketchInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configuration** | **string** |  | 
 **sketchId** | [**[]string**](string.md) |  | 
 **output3D** | **bool** |  | [default to false]
 **curvePoints** | **bool** |  | [default to false]
 **includeGeometry** | **bool** |  | [default to true]
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


## GetTessellatedEntities

> GetTessellatedEntities(ctx, did, wvm, wvmid, eid, sid).Configuration(configuration).EntityId(entityId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).LinkDocumentId(linkDocumentId).Execute()



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
    sid := "sid_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    entityId := []string{"Inner_example"} // []string |  (optional)
    angleTolerance := 987 // float64 |  (optional)
    chordTolerance := 987 // float64 |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetTessellatedEntities(context.Background(), did, wvm, wvmid, eid, sid).Configuration(configuration).EntityId(entityId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTessellatedEntities``: %v\n", err)
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
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTessellatedEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **configuration** | **string** |  | 
 **entityId** | [**[]string**](string.md) |  | 
 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
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


## GetWorkflowableTestObject

> BTWorkflowableTestObjectInfo GetWorkflowableTestObject(ctx, oid).Execute()



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
    oid := "oid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetWorkflowableTestObject(context.Background(), oid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWorkflowableTestObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowableTestObject`: BTWorkflowableTestObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWorkflowableTestObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAssociativeData

> BTAppAssociativeDataInfoArray PostAssociativeData(ctx, did, eid, wvm, wvmid).Body(body).Execute()



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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostAssociativeData(context.Background(), did, eid, wvm, wvmid, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostAssociativeData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAssociativeData`: BTAppAssociativeDataInfoArray
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostAssociativeData`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiPostAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**BTAppAssociativeDataInfoArray**](BTAppAssociativeDataInfoArray.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransitionWorkflowableTestObject

> BTWorkflowableTestObjectInfo TransitionWorkflowableTestObject(ctx, oid, transition).Execute()



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
    oid := "oid_example" // string | 
    transition := "transition_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TransitionWorkflowableTestObject(context.Background(), oid, transition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TransitionWorkflowableTestObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransitionWorkflowableTestObject`: BTWorkflowableTestObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TransitionWorkflowableTestObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string** |  | 
**transition** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransitionWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowableTestObject

> BTWorkflowableTestObjectInfo UpdateWorkflowableTestObject(ctx, oid).BTUpdateWorkflowableTestObjectParams(bTUpdateWorkflowableTestObjectParams).Execute()



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
    oid := "oid_example" // string | 
    bTUpdateWorkflowableTestObjectParams := openapiclient.BTUpdateWorkflowableTestObjectParams{Properties: []BTPropertyValueParam{openapiclient.BTPropertyValueParam{PropertyId: "PropertyId_example", Value: "TODO"})} // BTUpdateWorkflowableTestObjectParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWorkflowableTestObject(context.Background(), oid, bTUpdateWorkflowableTestObjectParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWorkflowableTestObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowableTestObject`: BTWorkflowableTestObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWorkflowableTestObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTUpdateWorkflowableTestObjectParams** | [**BTUpdateWorkflowableTestObjectParams**](BTUpdateWorkflowableTestObjectParams.md) |  | 

### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

