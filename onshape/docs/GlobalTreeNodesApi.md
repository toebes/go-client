# \GlobalTreeNodesApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalTreeNodes**](GlobalTreeNodesApi.md#GlobalTreeNodes) | **Get** /api/globaltreenodes/ | Get Searchable Trees
[**GlobalTreeNodesMagic**](GlobalTreeNodesApi.md#GlobalTreeNodesMagic) | **Get** /api/globaltreenodes/magic/{mid} | Get Tree Node List



## GlobalTreeNodes

> BTGlobalTreeNodesInfo GlobalTreeNodes(ctx).Execute()

Get Searchable Trees

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
    resp, r, err := api_client.GlobalTreeNodesApi.GlobalTreeNodes(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalTreeNodesApi.GlobalTreeNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GlobalTreeNodes`: BTGlobalTreeNodesInfo
    fmt.Fprintf(os.Stdout, "Response from `GlobalTreeNodesApi.GlobalTreeNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalTreeNodesRequest struct via the builder pattern


### Return type

[**BTGlobalTreeNodesInfo**](BTGlobalTreeNodesInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GlobalTreeNodesMagic

> BTGlobalTreeNodesInfo GlobalTreeNodesMagic(ctx, mid).GetPathToRoot(getPathToRoot).Offset(offset).Limit(limit).SortColumn(sortColumn).SortOrder(sortOrder).Execute()

Get Tree Node List

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
    mid := "mid_example" // string | 
    getPathToRoot := true // bool |  (optional) (default to false)
    offset := 987 // int32 |  (optional) (default to 0)
    limit := 987 // int32 |  (optional) (default to 20)
    sortColumn := "sortColumn_example" // string |  (optional) (default to "modifiedAt")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalTreeNodesApi.GlobalTreeNodesMagic(context.Background(), mid).GetPathToRoot(getPathToRoot).Offset(offset).Limit(limit).SortColumn(sortColumn).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalTreeNodesApi.GlobalTreeNodesMagic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GlobalTreeNodesMagic`: BTGlobalTreeNodesInfo
    fmt.Fprintf(os.Stdout, "Response from `GlobalTreeNodesApi.GlobalTreeNodesMagic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalTreeNodesMagicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getPathToRoot** | **bool** |  | [default to false]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **sortColumn** | **string** |  | [default to &quot;modifiedAt&quot;]
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]

### Return type

[**BTGlobalTreeNodesInfo**](BTGlobalTreeNodesInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

