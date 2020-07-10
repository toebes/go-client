# \WebhooksApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksApi.md#CreateWebhook) | **Post** /api/webhooks | 
[**GetWebhook**](WebhooksApi.md#GetWebhook) | **Get** /api/webhooks/{webhookid} | 
[**GetWebhooks**](WebhooksApi.md#GetWebhooks) | **Get** /api/webhooks | 
[**PingWebhook**](WebhooksApi.md#PingWebhook) | **Post** /api/webhooks/{webhookid}/ping | 
[**UnregisterWebhook**](WebhooksApi.md#UnregisterWebhook) | **Delete** /api/webhooks/{webhookid} | 
[**UpdateWebhook**](WebhooksApi.md#UpdateWebhook) | **Post** /api/webhooks/{webhookid} | 



## CreateWebhook

> BTWebhookInfo CreateWebhook(ctx).BTWebhookParams(bTWebhookParams).Execute()



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
    bTWebhookParams := openapiclient.BTWebhookParams{ClientId: "ClientId_example", CompanyId: "CompanyId_example", Data: "Data_example", DocumentId: "DocumentId_example", ElementId: "ElementId_example", Events: []string{"Events_example"), Filter: "Filter_example", FolderId: "FolderId_example", Id: "Id_example", Options: openapiclient.BTWebhookOptions{CollapseEvents: false}, PartId: "PartId_example", ProjectId: "ProjectId_example", Url: "Url_example", UserId: "UserId_example", VersionId: "VersionId_example", WorkspaceId: "WorkspaceId_example"} // BTWebhookParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.CreateWebhook(context.Background(), ).BTWebhookParams(bTWebhookParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: BTWebhookInfo
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTWebhookParams** | [**BTWebhookParams**](BTWebhookParams.md) |  | 

### Return type

[**BTWebhookInfo**](BTWebhookInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> BTWebhookInfo GetWebhook(ctx, webhookid).Execute()



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
    webhookid := "webhookid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.GetWebhook(context.Background(), webhookid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhook`: BTWebhookInfo
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWebhookInfo**](BTWebhookInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> BTListResponseBTWebhookInfo GetWebhooks(ctx).Company(company).User(user).Offset(offset).Limit(limit).Execute()



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
    company := "company_example" // string |  (optional) (default to "")
    user := "user_example" // string |  (optional)
    offset := 987 // int32 |  (optional) (default to 0)
    limit := 987 // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.GetWebhooks(context.Background(), ).Company(company).User(user).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooks`: BTListResponseBTWebhookInfo
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **company** | **string** |  | [default to &quot;&quot;]
 **user** | **string** |  | 
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTWebhookInfo**](BTListResponseBTWebhookInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingWebhook

> PingWebhook(ctx, webhookid).Execute()



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
    webhookid := "webhookid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.PingWebhook(context.Background(), webhookid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PingWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingWebhookRequest struct via the builder pattern


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


## UnregisterWebhook

> UnregisterWebhook(ctx, webhookid).Execute()



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
    webhookid := "webhookid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.UnregisterWebhook(context.Background(), webhookid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UnregisterWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterWebhookRequest struct via the builder pattern


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


## UpdateWebhook

> BTWebhookInfo UpdateWebhook(ctx, webhookid).BTWebhookParams(bTWebhookParams).Execute()



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
    webhookid := "webhookid_example" // string | 
    bTWebhookParams := openapiclient.BTWebhookParams{ClientId: "ClientId_example", CompanyId: "CompanyId_example", Data: "Data_example", DocumentId: "DocumentId_example", ElementId: "ElementId_example", Events: []string{"Events_example"), Filter: "Filter_example", FolderId: "FolderId_example", Id: "Id_example", Options: openapiclient.BTWebhookOptions{CollapseEvents: false}, PartId: "PartId_example", ProjectId: "ProjectId_example", Url: "Url_example", UserId: "UserId_example", VersionId: "VersionId_example", WorkspaceId: "WorkspaceId_example"} // BTWebhookParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WebhooksApi.UpdateWebhook(context.Background(), webhookid).BTWebhookParams(bTWebhookParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhook`: BTWebhookInfo
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTWebhookParams** | [**BTWebhookParams**](BTWebhookParams.md) |  | 

### Return type

[**BTWebhookInfo**](BTWebhookInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

