# \AssembliesApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFeature**](AssembliesApi.md#AddFeature) | **Post** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/features | 
[**CreateAssembly**](AssembliesApi.md#CreateAssembly) | **Post** /api/assemblies/d/{did}/w/{wid} | Create Assembly
[**CreateInstance**](AssembliesApi.md#CreateInstance) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/instances | Create assembly instance
[**DeleteFeature**](AssembliesApi.md#DeleteFeature) | **Delete** /api/assemblies/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Delete Feature
[**DeleteInstance**](AssembliesApi.md#DeleteInstance) | **Delete** /api/assemblies/d/{did}/w/{wid}/e/{eid}/instance/nodeid/{nid} | Delete assembly instance.
[**GetAssemblyBoundingBoxes**](AssembliesApi.md#GetAssemblyBoundingBoxes) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/boundingboxes | Bounding Boxes.
[**GetAssemblyDefinition**](AssembliesApi.md#GetAssemblyDefinition) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid} | Assembly Definition.
[**GetAssemblyShadedViews**](AssembliesApi.md#GetAssemblyShadedViews) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/shadedviews | 
[**GetBillOfMaterials**](AssembliesApi.md#GetBillOfMaterials) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/bom | Get Bill of Materials
[**GetFeatureSpecs**](AssembliesApi.md#GetFeatureSpecs) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs | 
[**GetFeatures**](AssembliesApi.md#GetFeatures) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Get Feature List
[**GetNamedViews**](AssembliesApi.md#GetNamedViews) | **Get** /api/assemblies/d/{did}/e/{eid}/namedViews | 
[**GetOrCreateBillOfMaterialsElement**](AssembliesApi.md#GetOrCreateBillOfMaterialsElement) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/bomelement | Get or Create Bill of Materials Element
[**GetTranslatorFormats**](AssembliesApi.md#GetTranslatorFormats) | **Get** /api/assemblies/d/{did}/w/{wid}/e/{eid}/translationformats | Get Translation Formats
[**InsertTransformedInstances**](AssembliesApi.md#InsertTransformedInstances) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/transformedinstances | Create and transform assembly instances
[**TransformOccurrences**](AssembliesApi.md#TransformOccurrences) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/occurrencetransforms | Transform assembly occurrences.
[**TranslateFormat**](AssembliesApi.md#TranslateFormat) | **Post** /api/assemblies/d/{did}/{wv}/{wvid}/e/{eid}/translations | Create Assembly translation.
[**UpdateFeature**](AssembliesApi.md#UpdateFeature) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | 



## AddFeature

> BTFeatureDefinitionResponse1617 AddFeature(ctx, did, wvm, wvmid, eid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()



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
    bTFeatureDefinitionCall1406 := openapiclient.BTFeatureDefinitionCall-1406{BtType: "BtType_example", LibraryVersion: 123, MicroversionSkew: false, RejectMicroversionSkew: false, SerializationVersion: "SerializationVersion_example", SourceMicroversion: "SourceMicroversion_example", Feature: openapiclient.BTMFeature-134{BtType: "BtType_example", FeatureId: "FeatureId_example", FeatureType: "FeatureType_example", ImportMicroversion: "ImportMicroversion_example", Name: "Name_example", Namespace: "Namespace_example", NodeId: "NodeId_example", Parameters: []BTMParameter1{openapiclient.BTMParameter-1{BtType: "BtType_example", ImportMicroversion: "ImportMicroversion_example", NodeId: "NodeId_example", ParameterId: "ParameterId_example"}), ReturnAfterSubfeatures: false, SubFeatures: []BTMFeature134{openapiclient.BTMFeature-134{BtType: "BtType_example", FeatureId: "FeatureId_example", FeatureType: "FeatureType_example", ImportMicroversion: "ImportMicroversion_example", Name: "Name_example", Namespace: "Namespace_example", NodeId: "NodeId_example", Parameters: []BTMParameter1{openapiclient.BTMParameter-1{BtType: "BtType_example", ImportMicroversion: "ImportMicroversion_example", NodeId: "NodeId_example", ParameterId: "ParameterId_example"}), ReturnAfterSubfeatures: false, SubFeatures: []BTMFeature134{), Suppressed: false}), Suppressed: false}} // BTFeatureDefinitionCall1406 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.AddFeature(context.Background(), did, wvm, wvmid, eid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.AddFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFeature`: BTFeatureDefinitionResponse1617
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.AddFeature`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTFeatureDefinitionCall1406** | [**BTFeatureDefinitionCall1406**](BTFeatureDefinitionCall1406.md) |  | 

### Return type

[**BTFeatureDefinitionResponse1617**](BTFeatureDefinitionResponse-1617.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssembly

> BTDocumentElementInfo CreateAssembly(ctx, did, wid).BTModelElementParams(bTModelElementParams).Execute()

Create Assembly

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
    resp, r, err := api_client.AssembliesApi.CreateAssembly(context.Background(), did, wid, bTModelElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.CreateAssembly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssembly`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.CreateAssembly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssemblyRequest struct via the builder pattern


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


## CreateInstance

> []BTOccurrence74 CreateInstance(ctx, did, wid, eid).BTAssemblyInstanceDefinitionParams(bTAssemblyInstanceDefinitionParams).Execute()

Create assembly instance

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
    bTAssemblyInstanceDefinitionParams := openapiclient.BTAssemblyInstanceDefinitionParams{Configuration: "Configuration_example", DocumentId: "DocumentId_example", ElementId: "ElementId_example", FeatureId: "FeatureId_example", IsAssembly: false, IsHidden: false, IsSuppressed: false, IsWholePartStudio: false, MicroversionId: "MicroversionId_example", PartId: "PartId_example", PartNumber: "PartNumber_example", Revision: "Revision_example", VersionId: "VersionId_example"} // BTAssemblyInstanceDefinitionParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.CreateInstance(context.Background(), did, wid, eid, bTAssemblyInstanceDefinitionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.CreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstance`: []BTOccurrence74
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.CreateInstance`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAssemblyInstanceDefinitionParams** | [**BTAssemblyInstanceDefinitionParams**](BTAssemblyInstanceDefinitionParams.md) |  | 

### Return type

[**[]BTOccurrence74**](BTOccurrence-74.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeature

> BTFeatureApiBase1430 DeleteFeature(ctx, did, wid, eid, fid).Execute()

Delete Feature

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
    fid := "fid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.DeleteFeature(context.Background(), did, wid, eid, fid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.DeleteFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFeature`: BTFeatureApiBase1430
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.DeleteFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTFeatureApiBase1430**](BTFeatureApiBase-1430.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance(ctx, did, eid, wid, nid).Execute()

Delete assembly instance.

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
    nid := "nid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.DeleteInstance(context.Background(), did, eid, wid, nid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.DeleteInstance``: %v\n", err)
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
**wid** | **string** |  | 
**nid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


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


## GetAssemblyBoundingBoxes

> BTBoundingBoxInfo GetAssemblyBoundingBoxes(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).IncludeHidden(includeHidden).DisplayStateId(displayStateId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()

Bounding Boxes.

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
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    includeHidden := true // bool |  (optional)
    displayStateId := "displayStateId_example" // string |  (optional)
    configuration := "configuration_example" // string |  (optional)
    explodedViewId := "explodedViewId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.GetAssemblyBoundingBoxes(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).IncludeHidden(includeHidden).DisplayStateId(displayStateId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.GetAssemblyBoundingBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssemblyBoundingBoxes`: BTBoundingBoxInfo
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.GetAssemblyBoundingBoxes`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAssemblyBoundingBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** |  | 
 **includeHidden** | **bool** |  | 
 **displayStateId** | **string** |  | 
 **configuration** | **string** |  | 
 **explodedViewId** | **string** |  | 

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


## GetAssemblyDefinition

> BTAssemblyDefinitionInfo GetAssemblyDefinition(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).IncludeMateFeatures(includeMateFeatures).IncludeNonSolids(includeNonSolids).IncludeMateConnectors(includeMateConnectors).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()

Assembly Definition.

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
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    includeMateFeatures := true // bool |  (optional)
    includeNonSolids := true // bool |  (optional)
    includeMateConnectors := true // bool |  (optional)
    configuration := "configuration_example" // string |  (optional)
    explodedViewId := "explodedViewId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.GetAssemblyDefinition(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).IncludeMateFeatures(includeMateFeatures).IncludeNonSolids(includeNonSolids).IncludeMateConnectors(includeMateConnectors).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.GetAssemblyDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssemblyDefinition`: BTAssemblyDefinitionInfo
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.GetAssemblyDefinition`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAssemblyDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** |  | 
 **includeMateFeatures** | **bool** |  | 
 **includeNonSolids** | **bool** |  | 
 **includeMateConnectors** | **bool** |  | 
 **configuration** | **string** |  | 
 **explodedViewId** | **string** |  | 

### Return type

[**BTAssemblyDefinitionInfo**](BTAssemblyDefinitionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssemblyShadedViews

> BTShadedViewsInfo GetAssemblyShadedViews(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).ShowAllParts(showAllParts).IncludeSurfaces(includeSurfaces).UseAntiAliasing(useAntiAliasing).DisplayStateId(displayStateId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()



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
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    viewMatrix := "viewMatrix_example" // string |  (optional) (default to "front")
    outputHeight := 987 // int32 |  (optional) (default to 500)
    outputWidth := 987 // int32 |  (optional) (default to 500)
    pixelSize := 987 // float64 |  (optional) (default to 0.003)
    edges := "edges_example" // string |  (optional) (default to "show")
    showAllParts := true // bool |  (optional) (default to false)
    includeSurfaces := true // bool |  (optional) (default to true)
    useAntiAliasing := true // bool |  (optional) (default to false)
    displayStateId := "displayStateId_example" // string |  (optional)
    configuration := "configuration_example" // string |  (optional)
    explodedViewId := "explodedViewId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.GetAssemblyShadedViews(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).ShowAllParts(showAllParts).IncludeSurfaces(includeSurfaces).UseAntiAliasing(useAntiAliasing).DisplayStateId(displayStateId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.GetAssemblyShadedViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssemblyShadedViews`: BTShadedViewsInfo
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.GetAssemblyShadedViews`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAssemblyShadedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** |  | 
 **viewMatrix** | **string** |  | [default to &quot;front&quot;]
 **outputHeight** | **int32** |  | [default to 500]
 **outputWidth** | **int32** |  | [default to 500]
 **pixelSize** | **float64** |  | [default to 0.003]
 **edges** | **string** |  | [default to &quot;show&quot;]
 **showAllParts** | **bool** |  | [default to false]
 **includeSurfaces** | **bool** |  | [default to true]
 **useAntiAliasing** | **bool** |  | [default to false]
 **displayStateId** | **string** |  | 
 **configuration** | **string** |  | 
 **explodedViewId** | **string** |  | 

### Return type

[**BTShadedViewsInfo**](BTShadedViewsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillOfMaterials

> JsonNode GetBillOfMaterials(ctx, did, wvm, wvmid, eid).MetadataWorkspaceId(metadataWorkspaceId).BomColumnIds(bomColumnIds).Indented(indented).MultiLevel(multiLevel).GenerateIfAbsent(generateIfAbsent).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()

Get Bill of Materials

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
    metadataWorkspaceId := "metadataWorkspaceId_example" // string |  (optional) (default to "")
    bomColumnIds := []string{"Inner_example"} // []string |  (optional)
    indented := true // bool |  (optional) (default to true)
    multiLevel := true // bool |  (optional) (default to false)
    generateIfAbsent := true // bool |  (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    configuration := "configuration_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.GetBillOfMaterials(context.Background(), did, wvm, wvmid, eid).MetadataWorkspaceId(metadataWorkspaceId).BomColumnIds(bomColumnIds).Indented(indented).MultiLevel(multiLevel).GenerateIfAbsent(generateIfAbsent).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.GetBillOfMaterials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillOfMaterials`: JsonNode
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.GetBillOfMaterials`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetBillOfMaterialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **metadataWorkspaceId** | **string** |  | [default to &quot;&quot;]
 **bomColumnIds** | [**[]string**](string.md) |  | 
 **indented** | **bool** |  | [default to true]
 **multiLevel** | **bool** |  | [default to false]
 **generateIfAbsent** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 
 **configuration** | **string** |  | 

### Return type

[**JsonNode**](JsonNode.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureSpecs

> BTFeatureSpecsResponse664 GetFeatureSpecs(ctx, did, wvm, wvmid, eid).Execute()



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
    resp, r, err := api_client.AssembliesApi.GetFeatureSpecs(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.GetFeatureSpecs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureSpecs`: BTFeatureSpecsResponse664
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.GetFeatureSpecs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFeatureSpecsRequest struct via the builder pattern


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


## GetFeatures

> BTAssemblyFeatureListResponse1174 GetFeatures(ctx, did, wvm, wvmid, eid).FeatureId(featureId).LinkDocumentId(linkDocumentId).Execute()

Get Feature List

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
    featureId := []string{"Inner_example"} // []string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.GetFeatures(context.Background(), did, wvm, wvmid, eid).FeatureId(featureId).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.GetFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatures`: BTAssemblyFeatureListResponse1174
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.GetFeatures`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **featureId** | [**[]string**](string.md) |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTAssemblyFeatureListResponse1174**](BTAssemblyFeatureListResponse-1174.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v2+json;charset=UTF-8;qs=0.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamedViews

> BTNamedViewsInfo GetNamedViews(ctx, did, eid).SkipPerspective(skipPerspective).LinkDocumentId(linkDocumentId).Execute()



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
    skipPerspective := true // bool |  (optional) (default to true)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.GetNamedViews(context.Background(), did, eid).SkipPerspective(skipPerspective).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.GetNamedViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamedViews`: BTNamedViewsInfo
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.GetNamedViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipPerspective** | **bool** |  | [default to true]
 **linkDocumentId** | **string** |  | 

### Return type

[**BTNamedViewsInfo**](BTNamedViewsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrCreateBillOfMaterialsElement

> BTDocumentElementInfo GetOrCreateBillOfMaterialsElement(ctx, did, wid, eid).Execute()

Get or Create Bill of Materials Element

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
    resp, r, err := api_client.AssembliesApi.GetOrCreateBillOfMaterialsElement(context.Background(), did, wid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.GetOrCreateBillOfMaterialsElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrCreateBillOfMaterialsElement`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.GetOrCreateBillOfMaterialsElement`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOrCreateBillOfMaterialsElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTranslatorFormats

> []BTModelFormatInfo GetTranslatorFormats(ctx, did, wid, eid).CheckContent(checkContent).Execute()

Get Translation Formats

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
    checkContent := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.GetTranslatorFormats(context.Background(), did, wid, eid).CheckContent(checkContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.GetTranslatorFormats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTranslatorFormats`: []BTModelFormatInfo
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.GetTranslatorFormats`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetTranslatorFormatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **checkContent** | **bool** |  | [default to true]

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


## InsertTransformedInstances

> BTAssemblyInsertTransformedInstancesResponse InsertTransformedInstances(ctx, did, eid, wid).BTAssemblyTransformedInstancesDefinitionParams(bTAssemblyTransformedInstancesDefinitionParams).Execute()

Create and transform assembly instances

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
    bTAssemblyTransformedInstancesDefinitionParams := openapiclient.BTAssemblyTransformedInstancesDefinitionParams{TransformGroups: []TransformGroup{openapiclient.TransformGroup{Instances: []BTAssemblyInstanceDefinitionParams{openapiclient.BTAssemblyInstanceDefinitionParams{Configuration: "Configuration_example", DocumentId: "DocumentId_example", ElementId: "ElementId_example", FeatureId: "FeatureId_example", IsAssembly: false, IsHidden: false, IsSuppressed: false, IsWholePartStudio: false, MicroversionId: "MicroversionId_example", PartId: "PartId_example", PartNumber: "PartNumber_example", Revision: "Revision_example", VersionId: "VersionId_example"}), Transform: []float64{123)})} // BTAssemblyTransformedInstancesDefinitionParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.InsertTransformedInstances(context.Background(), did, eid, wid, bTAssemblyTransformedInstancesDefinitionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.InsertTransformedInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsertTransformedInstances`: BTAssemblyInsertTransformedInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.InsertTransformedInstances`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiInsertTransformedInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAssemblyTransformedInstancesDefinitionParams** | [**BTAssemblyTransformedInstancesDefinitionParams**](BTAssemblyTransformedInstancesDefinitionParams.md) |  | 

### Return type

[**BTAssemblyInsertTransformedInstancesResponse**](BTAssemblyInsertTransformedInstancesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransformOccurrences

> TransformOccurrences(ctx, did, eid, wid).BTAssemblyTransformDefinitionParams(bTAssemblyTransformDefinitionParams).Execute()

Transform assembly occurrences.

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
    bTAssemblyTransformDefinitionParams := openapiclient.BTAssemblyTransformDefinitionParams{IsRelative: false, Occurrences: []BTOccurrence74{openapiclient.BTOccurrence-74{BtType: "BtType_example", FullPathAsString: "FullPathAsString_example", HeadInstanceId: "HeadInstanceId_example", OccurrenceWithoutHead: openapiclient.BTOccurrence-74{BtType: "BtType_example", FullPathAsString: "FullPathAsString_example", HeadInstanceId: "HeadInstanceId_example", OccurrenceWithoutHead: , OccurrenceWithoutTail: , Parent: , Path: []string{"Path_example"), PatternDescendant: false, RootOccurrence: false, TailInstanceId: "TailInstanceId_example"}, OccurrenceWithoutTail: , Parent: , Path: []string{"Path_example"), PatternDescendant: false, RootOccurrence: false, TailInstanceId: "TailInstanceId_example"}), Transform: []float64{123)} // BTAssemblyTransformDefinitionParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.TransformOccurrences(context.Background(), did, eid, wid, bTAssemblyTransformDefinitionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.TransformOccurrences``: %v\n", err)
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
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransformOccurrencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAssemblyTransformDefinitionParams** | [**BTAssemblyTransformDefinitionParams**](BTAssemblyTransformDefinitionParams.md) |  | 

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


## TranslateFormat

> BTTranslationRequestInfo TranslateFormat(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Create Assembly translation.

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
    resp, r, err := api_client.AssembliesApi.TranslateFormat(context.Background(), did, wv, wvid, eid, bTTranslateFormatParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.TranslateFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TranslateFormat`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.TranslateFormat`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiTranslateFormatRequest struct via the builder pattern


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


## UpdateFeature

> BTFeatureDefinitionResponse1617 UpdateFeature(ctx, did, wid, eid, fid).Body(body).Execute()



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
    fid := "fid_example" // string | 
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AssembliesApi.UpdateFeature(context.Background(), did, wid, eid, fid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssembliesApi.UpdateFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeature`: BTFeatureDefinitionResponse1617
    fmt.Fprintf(os.Stdout, "Response from `AssembliesApi.UpdateFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**BTFeatureDefinitionResponse1617**](BTFeatureDefinitionResponse-1617.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

