# \PartsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportPS**](PartsApi.md#ExportPS) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/parasolid | Export Part to Parasolid.
[**GetBendTable**](PartsApi.md#GetBendTable) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/sheetmetal/bendtable | Get Sheet Metal Bend Table.
[**GetBodyDetails**](PartsApi.md#GetBodyDetails) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/bodydetails | 
[**GetBoundingBoxes**](PartsApi.md#GetBoundingBoxes) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/boundingboxes | 
[**GetEdges**](PartsApi.md#GetEdges) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/tessellatededges | Tessellated Edges
[**GetFaces1**](PartsApi.md#GetFaces1) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/tessellatedfaces | Get Tessellated Faces
[**GetMassProperties**](PartsApi.md#GetMassProperties) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/massproperties | 
[**GetPartMetadata**](PartsApi.md#GetPartMetadata) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/metadata | 
[**GetPartsWMV**](PartsApi.md#GetPartsWMV) | **Get** /api/parts/d/{did}/{wvm}/{wvmid} | Get list of parts
[**GetPartsWMVE**](PartsApi.md#GetPartsWMVE) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid} | Get parts from an element.
[**GetShadedViews1**](PartsApi.md#GetShadedViews1) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/shadedviews | 
[**GetStandardContentPartMetadata**](PartsApi.md#GetStandardContentPartMetadata) | **Get** /api/parts/standardcontent/d/{did}/v/{vid}/e/{eid}/{otype}/{oid}/partid/{partid}/metadata | 
[**UpdatePartMetadata**](PartsApi.md#UpdatePartMetadata) | **Post** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/metadata | 
[**UpdateStandardContentPartMetadata**](PartsApi.md#UpdateStandardContentPartMetadata) | **Post** /api/parts/standardcontent/d/{did}/v/{vid}/e/{eid}/{otype}/{oid}/partid/{partid}/metadata | 



## ExportPS

> *os.File ExportPS(ctx, did, wvm, wvmid, eid, partid).Version(version).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Export Part to Parasolid.

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
    partid := "partid_example" // string | 
    version := "version_example" // string |  (optional) (default to "0")
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.ExportPS(context.Background(), did, wvm, wvmid, eid, partid).Version(version).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.ExportPS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPS`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.ExportPS`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **version** | **string** |  | [default to &quot;0&quot;]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

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


## GetBendTable

> BTTableResponse1546 GetBendTable(ctx, did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Execute()

Get Sheet Metal Bend Table.

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
    partid := "partid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetBendTable(context.Background(), did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetBendTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBendTable`: BTTableResponse1546
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetBendTable`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBendTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** |  | 

### Return type

[**BTTableResponse1546**](BTTableResponse-1546.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBodyDetails

> BTExportModelBodiesResponse734 GetBodyDetails(ctx, did, wvm, wvmid, eid, partid).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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
    partid := "partid_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetBodyDetails(context.Background(), did, wvm, wvmid, eid, partid).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetBodyDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBodyDetails`: BTExportModelBodiesResponse734
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetBodyDetails`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBodyDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTExportModelBodiesResponse734**](BTExportModelBodiesResponse-734.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoundingBoxes

> BTBoundingBoxInfo GetBoundingBoxes(ctx, did, wvm, wvmid, eid, partid).IncludeHidden(includeHidden).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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
    partid := "partid_example" // string | 
    includeHidden := true // bool |  (optional) (default to false)
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetBoundingBoxes(context.Background(), did, wvm, wvmid, eid, partid).IncludeHidden(includeHidden).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetBoundingBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoundingBoxes`: BTBoundingBoxInfo
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetBoundingBoxes`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoundingBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **includeHidden** | **bool** |  | [default to false]
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


## GetEdges

> BTExportTessellatedEdgesResponse327 GetEdges(ctx, did, wvm, wvmid, eid, partid).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).EdgeId(edgeId).Configuration(configuration).LinkDocumentId(linkDocumentId).Body(body).Execute()

Tessellated Edges

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
    partid := "partid_example" // string | 
    angleTolerance := 987 // float64 |  (optional)
    chordTolerance := 987 // float64 |  (optional)
    edgeId := []string{"Inner_example"} // []string |  (optional)
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetEdges(context.Background(), did, wvm, wvmid, eid, partid).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).EdgeId(edgeId).Configuration(configuration).LinkDocumentId(linkDocumentId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetEdges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdges`: BTExportTessellatedEdgesResponse327
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetEdges`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **edgeId** | [**[]string**](string.md) |  | 
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **body** | **string** |  | 

### Return type

[**BTExportTessellatedEdgesResponse327**](BTExportTessellatedEdgesResponse-327.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFaces1

> BTExportTessellatedFacesResponse898 GetFaces1(ctx, did, wvm, wvmid, eid, partid).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).OutputVertexNormals(outputVertexNormals).OutputFacetNormals(outputFacetNormals).OutputTextureCoordinates(outputTextureCoordinates).OutputFaceAppearances(outputFaceAppearances).OutputIndexTable(outputIndexTable).FaceId(faceId).Configuration(configuration).OutputErrorFaces(outputErrorFaces).LinkDocumentId(linkDocumentId).Body(body).Execute()

Get Tessellated Faces

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
    partid := "partid_example" // string | 
    angleTolerance := 987 // float64 |  (optional)
    chordTolerance := 987 // float64 |  (optional)
    maxFacetWidth := 987 // float64 |  (optional)
    outputVertexNormals := true // bool |  (optional) (default to false)
    outputFacetNormals := true // bool |  (optional) (default to true)
    outputTextureCoordinates := true // bool |  (optional) (default to false)
    outputFaceAppearances := true // bool |  (optional) (default to false)
    outputIndexTable := true // bool |  (optional) (default to false)
    faceId := []string{"Inner_example"} // []string |  (optional)
    configuration := "configuration_example" // string |  (optional)
    outputErrorFaces := true // bool |  (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetFaces1(context.Background(), did, wvm, wvmid, eid, partid).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).OutputVertexNormals(outputVertexNormals).OutputFacetNormals(outputFacetNormals).OutputTextureCoordinates(outputTextureCoordinates).OutputFaceAppearances(outputFaceAppearances).OutputIndexTable(outputIndexTable).FaceId(faceId).Configuration(configuration).OutputErrorFaces(outputErrorFaces).LinkDocumentId(linkDocumentId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetFaces1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFaces1`: BTExportTessellatedFacesResponse898
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetFaces1`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFaces1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **maxFacetWidth** | **float64** |  | 
 **outputVertexNormals** | **bool** |  | [default to false]
 **outputFacetNormals** | **bool** |  | [default to true]
 **outputTextureCoordinates** | **bool** |  | [default to false]
 **outputFaceAppearances** | **bool** |  | [default to false]
 **outputIndexTable** | **bool** |  | [default to false]
 **faceId** | [**[]string**](string.md) |  | 
 **configuration** | **string** |  | 
 **outputErrorFaces** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 
 **body** | **string** |  | 

### Return type

[**BTExportTessellatedFacesResponse898**](BTExportTessellatedFacesResponse-898.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMassProperties

> BTMassPropertiesBulkInfo GetMassProperties(ctx, did, wvm, wvmid, eid, partid).InferMetadataOwner(inferMetadataOwner).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()



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
    partid := "partid_example" // string | 
    inferMetadataOwner := true // bool |  (optional) (default to true)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    configuration := "configuration_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetMassProperties(context.Background(), did, wvm, wvmid, eid, partid).InferMetadataOwner(inferMetadataOwner).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetMassProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMassProperties`: BTMassPropertiesBulkInfo
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetMassProperties`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMassPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **inferMetadataOwner** | **bool** |  | [default to true]
 **linkDocumentId** | **string** |  | 
 **configuration** | **string** |  | 

### Return type

[**BTMassPropertiesBulkInfo**](BTMassPropertiesBulkInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartMetadata

> BTPartMetadataInfo GetPartMetadata(ctx, did, wvm, wvmid, eid, partid).InferMetadataOwner(inferMetadataOwner).IncludePropertyDefaults(includePropertyDefaults).FriendlyUserIds(friendlyUserIds).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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
    partid := "partid_example" // string | 
    inferMetadataOwner := true // bool |  (optional) (default to false)
    includePropertyDefaults := true // bool |  (optional) (default to false)
    friendlyUserIds := true // bool |  (optional) (default to false)
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetPartMetadata(context.Background(), did, wvm, wvmid, eid, partid).InferMetadataOwner(inferMetadataOwner).IncludePropertyDefaults(includePropertyDefaults).FriendlyUserIds(friendlyUserIds).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetPartMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartMetadata`: BTPartMetadataInfo
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetPartMetadata`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **inferMetadataOwner** | **bool** |  | [default to false]
 **includePropertyDefaults** | **bool** |  | [default to false]
 **friendlyUserIds** | **bool** |  | [default to false]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartsWMV

> []BTPartMetadataInfo GetPartsWMV(ctx, did, wvm, wvmid).ElementId(elementId).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()

Get list of parts

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
    did := "did_example" // string | Document ID.
    wvm := "wvm_example" // string | One of w or v or m corresponding to whether a workspace or version or microversion was entered.
    wvmid := "wvmid_example" // string | Workspace (w), Version (v) or Microversion (m) ID.
    elementId := "elementId_example" // string | Element ID (optional)
    withThumbnails := true // bool | Whether or not to include thumbnails (not supported for microversion) (optional) (default to false)
    includePropertyDefaults := true // bool | If true, include metadata schema property defaults in response (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)
    configuration := "configuration_example" // string | Configuration string. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetPartsWMV(context.Background(), did, wvm, wvmid).ElementId(elementId).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetPartsWMV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartsWMV`: []BTPartMetadataInfo
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetPartsWMV`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wvm** | **string** | One of w or v or m corresponding to whether a workspace or version or microversion was entered. | 
**wvmid** | **string** | Workspace (w), Version (v) or Microversion (m) ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartsWMVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **elementId** | **string** | Element ID | 
 **withThumbnails** | **bool** | Whether or not to include thumbnails (not supported for microversion) | [default to false]
 **includePropertyDefaults** | **bool** | If true, include metadata schema property defaults in response | [default to false]
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 
 **configuration** | **string** | Configuration string. | 

### Return type

[**[]BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartsWMVE

> []BTPartMetadataInfo GetPartsWMVE(ctx, did, wvm, wvmid, eid).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Get parts from an element.

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
    did := "did_example" // string | Document ID.
    wvm := "wvm_example" // string | One of w or v or m corresponding to whether a workspace or version or microversion was entered.
    wvmid := "wvmid_example" // string | Workspace (w), Version (v) or Microversion (m) ID.
    eid := "eid_example" // string | Element ID.
    withThumbnails := true // bool | Whether or not to include thumbnails (not supported for microversion) (optional) (default to false)
    includePropertyDefaults := true // bool | If true, include metadata schema property defaults in response (optional) (default to false)
    configuration := "configuration_example" // string | Configuration string. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetPartsWMVE(context.Background(), did, wvm, wvmid, eid).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetPartsWMVE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartsWMVE`: []BTPartMetadataInfo
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetPartsWMVE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wvm** | **string** | One of w or v or m corresponding to whether a workspace or version or microversion was entered. | 
**wvmid** | **string** | Workspace (w), Version (v) or Microversion (m) ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartsWMVERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **withThumbnails** | **bool** | Whether or not to include thumbnails (not supported for microversion) | [default to false]
 **includePropertyDefaults** | **bool** | If true, include metadata schema property defaults in response | [default to false]
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**[]BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShadedViews1

> BTShadedViewsInfo GetShadedViews1(ctx, did, wvm, wvmid, eid, partid).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).UseAntiAliasing(useAntiAliasing).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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
    partid := "partid_example" // string | 
    viewMatrix := "viewMatrix_example" // string |  (optional) (default to "front")
    outputHeight := 987 // int32 |  (optional) (default to 500)
    outputWidth := 987 // int32 |  (optional) (default to 500)
    pixelSize := 987 // float64 |  (optional) (default to 0.003)
    edges := "edges_example" // string |  (optional) (default to "show")
    useAntiAliasing := true // bool |  (optional) (default to false)
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetShadedViews1(context.Background(), did, wvm, wvmid, eid, partid).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).UseAntiAliasing(useAntiAliasing).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetShadedViews1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShadedViews1`: BTShadedViewsInfo
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetShadedViews1`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShadedViews1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **viewMatrix** | **string** |  | [default to &quot;front&quot;]
 **outputHeight** | **int32** |  | [default to 500]
 **outputWidth** | **int32** |  | [default to 500]
 **pixelSize** | **float64** |  | [default to 0.003]
 **edges** | **string** |  | [default to &quot;show&quot;]
 **useAntiAliasing** | **bool** |  | [default to false]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

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


## GetStandardContentPartMetadata

> BTPartMetadataInfo GetStandardContentPartMetadata(ctx, did, vid, eid, otype, oid, partid).IncludePropertyDefaults(includePropertyDefaults).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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
    partid := "partid_example" // string | 
    includePropertyDefaults := true // bool |  (optional) (default to false)
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.GetStandardContentPartMetadata(context.Background(), did, vid, eid, otype, oid, partid).IncludePropertyDefaults(includePropertyDefaults).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.GetStandardContentPartMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStandardContentPartMetadata`: BTPartMetadataInfo
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.GetStandardContentPartMetadata`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStandardContentPartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **includePropertyDefaults** | **bool** |  | [default to false]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartMetadata

> BTPartMetadataInfo UpdatePartMetadata(ctx, did, wvm, wvmid, eid, partid).BTWorkspacePartParams(bTWorkspacePartParams).Execute()



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
    partid := "partid_example" // string | 
    bTWorkspacePartParams := openapiclient.BTWorkspacePartParams{Appearance: openapiclient.BTPartAppearanceParams{Color: openapiclient.BTColorParams{Blue: 123, Green: 123, Red: 123}, Opacity: 123}, ApplyUpdateToAllConfigurations: false, Configuration: "Configuration_example", ConnectionId: "ConnectionId_example", CustomProperties: []BTNameValuePair{openapiclient.BTNameValuePair{Name: "Name_example", Value: "Value_example"}), CustomPropertyDefinitions: []BTCustomPropertyDefinitionParams{openapiclient.BTCustomPropertyDefinitionParams{Description: "Description_example", EnumDefinition: []string{"EnumDefinition_example"), Name: "Name_example", Template: "Template_example", Type: "Type_example"}), Description: "Description_example", ElementId: "ElementId_example", Material: openapiclient.BTMaterialParams{DisplayName: "DisplayName_example", Id: "Id_example", LibraryName: "LibraryName_example", LibraryReference: openapiclient.BTExternalElementReferenceInfo{DocumentId: "DocumentId_example", ElementId: "ElementId_example", ElementMicroversionId: "ElementMicroversionId_example", VersionId: "VersionId_example"}, Properties: []BTMaterialPropertyParams{openapiclient.BTMaterialPropertyParams{Category: "Category_example", Description: "Description_example", DisplayName: "DisplayName_example", Name: "Name_example", Type: "Type_example", Units: "Units_example", Value: "Value_example"})}, Name: "Name_example", PartId: "PartId_example", PartNumber: "PartNumber_example", ProductLine: "ProductLine_example", Project: "Project_example", Revision: "Revision_example", Title1: "Title1_example", Title2: "Title2_example", Title3: "Title3_example", Vendor: "Vendor_example"} // BTWorkspacePartParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.UpdatePartMetadata(context.Background(), did, wvm, wvmid, eid, partid).BTWorkspacePartParams(bTWorkspacePartParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.UpdatePartMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePartMetadata`: BTPartMetadataInfo
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.UpdatePartMetadata`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **bTWorkspacePartParams** | [**BTWorkspacePartParams**](BTWorkspacePartParams.md) |  | 

### Return type

[**BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStandardContentPartMetadata

> BTPartMetadataInfo UpdateStandardContentPartMetadata(ctx, did, vid, eid, otype, oid, partid).LinkDocumentId(linkDocumentId).IncludePropertyDefaults(includePropertyDefaults).BTWorkspacePartParams(bTWorkspacePartParams).Execute()



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
    partid := "partid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    includePropertyDefaults := true // bool |  (optional) (default to false)
    bTWorkspacePartParams := openapiclient.BTWorkspacePartParams{Appearance: openapiclient.BTPartAppearanceParams{Color: openapiclient.BTColorParams{Blue: 123, Green: 123, Red: 123}, Opacity: 123}, ApplyUpdateToAllConfigurations: false, Configuration: "Configuration_example", ConnectionId: "ConnectionId_example", CustomProperties: []BTNameValuePair{openapiclient.BTNameValuePair{Name: "Name_example", Value: "Value_example"}), CustomPropertyDefinitions: []BTCustomPropertyDefinitionParams{openapiclient.BTCustomPropertyDefinitionParams{Description: "Description_example", EnumDefinition: []string{"EnumDefinition_example"), Name: "Name_example", Template: "Template_example", Type: "Type_example"}), Description: "Description_example", ElementId: "ElementId_example", Material: openapiclient.BTMaterialParams{DisplayName: "DisplayName_example", Id: "Id_example", LibraryName: "LibraryName_example", LibraryReference: openapiclient.BTExternalElementReferenceInfo{DocumentId: "DocumentId_example", ElementId: "ElementId_example", ElementMicroversionId: "ElementMicroversionId_example", VersionId: "VersionId_example"}, Properties: []BTMaterialPropertyParams{openapiclient.BTMaterialPropertyParams{Category: "Category_example", Description: "Description_example", DisplayName: "DisplayName_example", Name: "Name_example", Type: "Type_example", Units: "Units_example", Value: "Value_example"})}, Name: "Name_example", PartId: "PartId_example", PartNumber: "PartNumber_example", ProductLine: "ProductLine_example", Project: "Project_example", Revision: "Revision_example", Title1: "Title1_example", Title2: "Title2_example", Title3: "Title3_example", Vendor: "Vendor_example"} // BTWorkspacePartParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartsApi.UpdateStandardContentPartMetadata(context.Background(), did, vid, eid, otype, oid, partid).LinkDocumentId(linkDocumentId).IncludePropertyDefaults(includePropertyDefaults).BTWorkspacePartParams(bTWorkspacePartParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartsApi.UpdateStandardContentPartMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStandardContentPartMetadata`: BTPartMetadataInfo
    fmt.Fprintf(os.Stdout, "Response from `PartsApi.UpdateStandardContentPartMetadata`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStandardContentPartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **linkDocumentId** | **string** |  | 
 **includePropertyDefaults** | **bool** |  | [default to false]
 **bTWorkspacePartParams** | [**BTWorkspacePartParams**](BTWorkspacePartParams.md) |  | 

### Return type

[**BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

