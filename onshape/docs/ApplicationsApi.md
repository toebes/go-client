# \ApplicationsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppSettings**](ApplicationsApi.md#DeleteAppSettings) | **Delete** /api/applications/clients/{cid}/settings/users/{uid} | Delete Application Settings
[**GetUserAppSettings**](ApplicationsApi.md#GetUserAppSettings) | **Get** /api/applications/clients/{cid}/settings/users/{uid} | Get User Application Settings
[**UpdateAppSettings**](ApplicationsApi.md#UpdateAppSettings) | **Post** /api/applications/clients/{cid}/settings/users/{uid} | Update Application Settings



## DeleteAppSettings

> DeleteAppSettings(ctx, uid, cid).Key(key).Execute()

Delete Application Settings

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
    uid := "uid_example" // string | 
    cid := "cid_example" // string | 
    key := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsApi.DeleteAppSettings(context.Background(), uid, cid).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.DeleteAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**[]string**](string.md) |  | 

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


## GetUserAppSettings

> BTUserAppSettingsInfo GetUserAppSettings(ctx, uid, cid).Key(key).Execute()

Get User Application Settings

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
    uid := "uid_example" // string | 
    cid := "cid_example" // string | 
    key := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsApi.GetUserAppSettings(context.Background(), uid, cid).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.GetUserAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserAppSettings`: BTUserAppSettingsInfo
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.GetUserAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | [**[]string**](string.md) |  | 

### Return type

[**BTUserAppSettingsInfo**](BTUserAppSettingsInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppSettings

> UpdateAppSettings(ctx, uid, cid).BTUserAppSettingsParams(bTUserAppSettingsParams).Execute()

Update Application Settings

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
    uid := "uid_example" // string | 
    cid := "cid_example" // string | 
    bTUserAppSettingsParams := openapiclient.BTUserAppSettingsParams{Settings: []BTSettingParam{openapiclient.BTSettingParam{Key: "Key_example", Value: "TODO"})} // BTUserAppSettingsParams | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationsApi.UpdateAppSettings(context.Background(), uid, cid, bTUserAppSettingsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.UpdateAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTUserAppSettingsParams** | [**BTUserAppSettingsParams**](BTUserAppSettingsParams.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

