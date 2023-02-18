# \PageAccessUserMetricsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Delete metrics for page access user
[**DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId**](PageAccessUserMetricsApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id} | Delete metric for page access user
[**GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Get** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Get metrics for page access user
[**PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Patch** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Add metrics for page access user
[**PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Post** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Replace metrics for page access user
[**PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Put** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Add metrics for page access user



## DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics

> PageAccessUser DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId).DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics(deletePagesPageIdPageAccessUsersPageAccessUserIdMetrics).Execute()

Delete metrics for page access user



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
    pageId := "pageId_example" // string | Page identifier
    pageAccessUserId := "pageAccessUserId_example" // string | Page Access User Identifier
    deletePagesPageIdPageAccessUsersPageAccessUserIdMetrics := *openapiclient.NewDeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics() // DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserMetricsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics(context.Background(), pageId, pageAccessUserId).DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics(deletePagesPageIdPageAccessUsersPageAccessUserIdMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserMetricsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserMetricsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deletePagesPageIdPageAccessUsersPageAccessUserIdMetrics** | [**DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics**](DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics.md) |  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId

> PageAccessUser DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId(ctx, pageId, pageAccessUserId, metricId).Execute()

Delete metric for page access user



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
    pageId := "pageId_example" // string | Page identifier
    pageAccessUserId := "pageAccessUserId_example" // string | Page Access User Identifier
    metricId := "metricId_example" // string | Identifier of metric requested

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserMetricsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId(context.Background(), pageId, pageAccessUserId, metricId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserMetricsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserMetricsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 
**metricId** | **string** | Identifier of metric requested | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics

> []Metric GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId).Page(page).PerPage(perPage).Execute()

Get metrics for page access user



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
    pageId := "pageId_example" // string | Page identifier
    pageAccessUserId := "pageAccessUserId_example" // string | Page Access User Identifier
    page := int32(56) // int32 | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. (optional)
    perPage := int32(56) // int32 | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserMetricsApi.GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics(context.Background(), pageId, pageAccessUserId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserMetricsApi.GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserMetricsApi.GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdPageAccessUsersPageAccessUserIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

### Return type

[**[]Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics

> PageAccessUser PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId).PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics(patchPagesPageIdPageAccessUsersPageAccessUserIdMetrics).Execute()

Add metrics for page access user



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
    pageId := "pageId_example" // string | Page identifier
    pageAccessUserId := "pageAccessUserId_example" // string | Page Access User Identifier
    patchPagesPageIdPageAccessUsersPageAccessUserIdMetrics := *openapiclient.NewPatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics([]string{"MetricIds_example"}) // PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserMetricsApi.PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics(context.Background(), pageId, pageAccessUserId).PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics(patchPagesPageIdPageAccessUsersPageAccessUserIdMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserMetricsApi.PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserMetricsApi.PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdPageAccessUsersPageAccessUserIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdPageAccessUsersPageAccessUserIdMetrics** | [**PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics.md) |  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics

> PageAccessUser PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId).PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics(postPagesPageIdPageAccessUsersPageAccessUserIdMetrics).Execute()

Replace metrics for page access user



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
    pageId := "pageId_example" // string | Page identifier
    pageAccessUserId := "pageAccessUserId_example" // string | Page Access User Identifier
    postPagesPageIdPageAccessUsersPageAccessUserIdMetrics := *openapiclient.NewPostPagesPageIdPageAccessUsersPageAccessUserIdMetrics([]string{"MetricIds_example"}) // PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserMetricsApi.PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics(context.Background(), pageId, pageAccessUserId).PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics(postPagesPageIdPageAccessUsersPageAccessUserIdMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserMetricsApi.PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserMetricsApi.PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdPageAccessUsersPageAccessUserIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postPagesPageIdPageAccessUsersPageAccessUserIdMetrics** | [**PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics.md) |  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics

> PageAccessUser PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId).PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics(putPagesPageIdPageAccessUsersPageAccessUserIdMetrics).Execute()

Add metrics for page access user



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
    pageId := "pageId_example" // string | Page identifier
    pageAccessUserId := "pageAccessUserId_example" // string | Page Access User Identifier
    putPagesPageIdPageAccessUsersPageAccessUserIdMetrics := *openapiclient.NewPutPagesPageIdPageAccessUsersPageAccessUserIdMetrics([]string{"MetricIds_example"}) // PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserMetricsApi.PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics(context.Background(), pageId, pageAccessUserId).PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics(putPagesPageIdPageAccessUsersPageAccessUserIdMetrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserMetricsApi.PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserMetricsApi.PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdPageAccessUsersPageAccessUserIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdPageAccessUsersPageAccessUserIdMetrics** | [**PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics.md) |  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

