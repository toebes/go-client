# \GlobalTreeNodesApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalTreeNodes**](GlobalTreeNodesApi.md#GlobalTreeNodes) | **Get** /api/globaltreenodes/ | Get Searchable Trees
[**GlobalTreeNodesFolder**](GlobalTreeNodesApi.md#GlobalTreeNodesFolder) | **Get** /api/globaltreenodes/folder/{fid} | Get Tree Node List
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


## GlobalTreeNodesFolder

> BTGlobalTreeNodesInfo GlobalTreeNodesFolder(ctx, fid).GetPathToRoot(getPathToRoot).Offset(offset).Limit(limit).SortColumn(sortColumn).SortOrder(sortOrder).Execute()

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
    fid := "fid_example" // string | ID of the folder to traverse.
    getPathToRoot := true // bool | When true, requests that the path of all nodes up to the root are to be returned. (optional) (default to false)
    offset := 987 // int32 | Starting offset in the sorted list of items to return. (optional) (default to 0)
    limit := 987 // int32 | Maximum number of items in the list to return. (optional) (default to 20)
    sortColumn := "sortColumn_example" // string | Field to sort the items in the list by. (optional) (default to "modifiedAt")
    sortOrder := "sortOrder_example" // string | Which way to sort the list (asc/desc).  Defaults to asc. (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GlobalTreeNodesApi.GlobalTreeNodesFolder(context.Background(), fid).GetPathToRoot(getPathToRoot).Offset(offset).Limit(limit).SortColumn(sortColumn).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalTreeNodesApi.GlobalTreeNodesFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GlobalTreeNodesFolder`: BTGlobalTreeNodesInfo
    fmt.Fprintf(os.Stdout, "Response from `GlobalTreeNodesApi.GlobalTreeNodesFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fid** | **string** | ID of the folder to traverse. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalTreeNodesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getPathToRoot** | **bool** | When true, requests that the path of all nodes up to the root are to be returned. | [default to false]
 **offset** | **int32** | Starting offset in the sorted list of items to return. | [default to 0]
 **limit** | **int32** | Maximum number of items in the list to return. | [default to 20]
 **sortColumn** | **string** | Field to sort the items in the list by. | [default to &quot;modifiedAt&quot;]
 **sortOrder** | **string** | Which way to sort the list (asc/desc).  Defaults to asc. | [default to &quot;asc&quot;]

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
    mid := "mid_example" // string | Magic ID of list to return. Known values include: | MAGIC ID | TITLE | USAGE | |-|-|-| | 0 | Recently Opened | Most recently opened documents | | 1 | My Onshape | Root folder and contents | | 2 | Created by Me | Documents created by the logged in user | | 3 | Public | All public documents (a very long list) | | 4 | Trash | Trashcan for the logged in user | | 5 | Tutorials & Samples | Desktop Tutorials | | 6 | FeatureScript samples |  FeatureScript samples (found when you select Other documents while adding a custom feature) | | 7 | Community spotlight | Community spotlight (found when you select Other documents while adding a custom feature) | | 8 | Tutorials | IOS Tutorials | | 9 | Tutorials | Android Tutorials | | 10 | Labels | Labels created by the user  | | 11 | Teams | Teams that the user is connected to | | 12 | Shared with me | Documents shared with the user | | 13 | Cloud Storage | Visual list of cloud accounts associated with the logged in user | | 14 | Custom table samples | Custom table samples (found when you select Other documents while adding a custom table) | 
    getPathToRoot := true // bool | When true, requests that the path of all nodes up to the root are to be returned. (optional) (default to false)
    offset := 987 // int32 | Starting offset in the sorted list of items to return. (optional) (default to 0)
    limit := 987 // int32 | Maximum number of items in the list to return. (optional) (default to 20)
    sortColumn := "sortColumn_example" // string | Field to sort the items in the list by. (optional) (default to "modifiedAt")
    sortOrder := "sortOrder_example" // string | Which way to sort the list (asc/desc).  Defaults to asc. (optional) (default to "asc")

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
**mid** | **string** | Magic ID of list to return. Known values include: | MAGIC ID | TITLE | USAGE | |-|-|-| | 0 | Recently Opened | Most recently opened documents | | 1 | My Onshape | Root folder and contents | | 2 | Created by Me | Documents created by the logged in user | | 3 | Public | All public documents (a very long list) | | 4 | Trash | Trashcan for the logged in user | | 5 | Tutorials &amp; Samples | Desktop Tutorials | | 6 | FeatureScript samples |  FeatureScript samples (found when you select Other documents while adding a custom feature) | | 7 | Community spotlight | Community spotlight (found when you select Other documents while adding a custom feature) | | 8 | Tutorials | IOS Tutorials | | 9 | Tutorials | Android Tutorials | | 10 | Labels | Labels created by the user  | | 11 | Teams | Teams that the user is connected to | | 12 | Shared with me | Documents shared with the user | | 13 | Cloud Storage | Visual list of cloud accounts associated with the logged in user | | 14 | Custom table samples | Custom table samples (found when you select Other documents while adding a custom table) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGlobalTreeNodesMagicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getPathToRoot** | **bool** | When true, requests that the path of all nodes up to the root are to be returned. | [default to false]
 **offset** | **int32** | Starting offset in the sorted list of items to return. | [default to 0]
 **limit** | **int32** | Maximum number of items in the list to return. | [default to 20]
 **sortColumn** | **string** | Field to sort the items in the list by. | [default to &quot;modifiedAt&quot;]
 **sortOrder** | **string** | Which way to sort the list (asc/desc).  Defaults to asc. | [default to &quot;asc&quot;]

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

