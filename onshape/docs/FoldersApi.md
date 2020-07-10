# \FoldersApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFolderAcl**](FoldersApi.md#GetFolderAcl) | **Get** /api/folders/{fid}/acl | Get Access Control List
[**Share**](FoldersApi.md#Share) | **Post** /api/folders/{fid}/share | Share Folder
[**UnShare**](FoldersApi.md#UnShare) | **Delete** /api/folders/{fid}/share/{eid} | Unshare Folder



## GetFolderAcl

> BTAclInfo GetFolderAcl(ctx, fid).Execute()

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
    fid := "fid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FoldersApi.GetFolderAcl(context.Background(), fid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.GetFolderAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolderAcl`: BTAclInfo
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.GetFolderAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderAclRequest struct via the builder pattern


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


## Share

> BTAclInfo Share(ctx, fid).BTShareParams(bTShareParams).Execute()

Share Folder

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
    fid := "fid_example" // string | 
    bTShareParams := openapiclient.BTShareParams{DocumentId: "DocumentId_example", ElementId: "ElementId_example", EncodedConfiguration: "EncodedConfiguration_example", Entries: []BTShareEntryParams{openapiclient.BTShareEntryParams{ApplicationId: "ApplicationId_example", CompanyId: "CompanyId_example", Email: "Email_example", EntryType: 123, TeamId: "TeamId_example", UserId: "UserId_example"}), FolderId: "FolderId_example", Message: "Message_example", Permission: int64(123), PermissionSet: []string{"PermissionSet_example"), Update: false, WorkspaceId: "WorkspaceId_example"} // BTShareParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FoldersApi.Share(context.Background(), fid, bTShareParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.Share``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Share`: BTAclInfo
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.Share`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareRequest struct via the builder pattern


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


## UnShare

> UnShare(ctx, fid, eid).EntryType(entryType).Execute()

Unshare Folder

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
    fid := "fid_example" // string | 
    eid := "eid_example" // string | 
    entryType := 987 // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FoldersApi.UnShare(context.Background(), fid, eid).EntryType(entryType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.UnShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnShareRequest struct via the builder pattern


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

