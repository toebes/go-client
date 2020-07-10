# \AccountsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPurchaseNew**](AccountsApi.md#CancelPurchaseNew) | **Delete** /api/accounts/{aid}/purchases/{pid} | Cancel Recurring Subscription
[**ConsumePurchase**](AccountsApi.md#ConsumePurchase) | **Post** /api/accounts/purchases/{pid}/consume | Mark Purchase Consumed For User
[**GetPlanPurchases**](AccountsApi.md#GetPlanPurchases) | **Get** /api/accounts/plans/{planId}/purchases | Get Plan Purchases
[**GetPurchases**](AccountsApi.md#GetPurchases) | **Get** /api/accounts/purchases | Get User&#39;s Appstore Purchases.



## CancelPurchaseNew

> CancelPurchaseNew(ctx, aid, pid).CancelImmediately(cancelImmediately).Execute()

Cancel Recurring Subscription

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
    aid := "aid_example" // string | 
    pid := "pid_example" // string | 
    cancelImmediately := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.CancelPurchaseNew(context.Background(), aid, pid).CancelImmediately(cancelImmediately).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CancelPurchaseNew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPurchaseNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cancelImmediately** | **bool** |  | [default to false]

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


## ConsumePurchase

> BTPurchaseInfo ConsumePurchase(ctx, pid).BTPurchaseUserParams(bTPurchaseUserParams).Execute()

Mark Purchase Consumed For User

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
    bTPurchaseUserParams := openapiclient.BTPurchaseUserParams{ConsumedQuantity: 123, PurchaseId: "PurchaseId_example", UserId: "UserId_example"} // BTPurchaseUserParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ConsumePurchase(context.Background(), pid).BTPurchaseUserParams(bTPurchaseUserParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ConsumePurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConsumePurchase`: BTPurchaseInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ConsumePurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumePurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTPurchaseUserParams** | [**BTPurchaseUserParams**](BTPurchaseUserParams.md) |  | 

### Return type

[**BTPurchaseInfo**](BTPurchaseInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanPurchases

> BTListResponseBTPurchaseInfo GetPlanPurchases(ctx, planId).Offset(offset).Limit(limit).Execute()

Get Plan Purchases

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
    planId := "planId_example" // string | 
    offset := 987 // int32 |  (optional) (default to 0)
    limit := 987 // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetPlanPurchases(context.Background(), planId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetPlanPurchases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlanPurchases`: BTListResponseBTPurchaseInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetPlanPurchases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTPurchaseInfo**](BTListResponseBTPurchaseInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchases

> []BTPurchaseInfo GetPurchases(ctx).All(all).OwnPurchaseOnly(ownPurchaseOnly).Execute()

Get User's Appstore Purchases.

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
    all := true // bool |  (optional) (default to false)
    ownPurchaseOnly := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetPurchases(context.Background(), ).All(all).OwnPurchaseOnly(ownPurchaseOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetPurchases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchases`: []BTPurchaseInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetPurchases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** |  | [default to false]
 **ownPurchaseOnly** | **bool** |  | [default to false]

### Return type

[**[]BTPurchaseInfo**](BTPurchaseInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

