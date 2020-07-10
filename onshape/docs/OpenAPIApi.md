# \OpenAPIApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOpenApi**](OpenAPIApi.md#GetOpenApi) | **Get** /api/openapi | OpenAPI spec documentation for the Onshape REST API.



## GetOpenApi

> GetOpenApi(ctx).ExcludedTags(excludedTags).IncludedTags(includedTags).IncludeDeprecated(includeDeprecated).DocumentationStatus(documentationStatus).Execute()

OpenAPI spec documentation for the Onshape REST API.

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
    excludedTags := "excludedTags_example" // string | If an operation contains an excluded tag, it is not returned from this endpoint. (optional)
    includedTags := "includedTags_example" // string | Return only operations with tags included in includedTags. (optional)
    includeDeprecated := true // bool | Include deprecated endpoints. (optional)
    documentationStatus := []string{"DocumentationStatus_example"} // []string | Only return endpoints that have the specified document status. Default is to return all the endpoints the user should have access to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OpenAPIApi.GetOpenApi(context.Background(), ).ExcludedTags(excludedTags).IncludedTags(includedTags).IncludeDeprecated(includeDeprecated).DocumentationStatus(documentationStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIApi.GetOpenApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludedTags** | **string** | If an operation contains an excluded tag, it is not returned from this endpoint. | 
 **includedTags** | **string** | Return only operations with tags included in includedTags. | 
 **includeDeprecated** | **bool** | Include deprecated endpoints. | 
 **documentationStatus** | [**[]string**](string.md) | Only return endpoints that have the specified document status. Default is to return all the endpoints the user should have access to. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

