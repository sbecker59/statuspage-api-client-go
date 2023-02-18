# \ComponentGroupsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdComponentGroupsId**](ComponentGroupsApi.md#DeletePagesPageIdComponentGroupsId) | **Delete** /pages/{page_id}/component-groups/{id} | Delete a component group
[**GetPagesPageIdComponentGroups**](ComponentGroupsApi.md#GetPagesPageIdComponentGroups) | **Get** /pages/{page_id}/component-groups | Get a list of component groups
[**GetPagesPageIdComponentGroupsId**](ComponentGroupsApi.md#GetPagesPageIdComponentGroupsId) | **Get** /pages/{page_id}/component-groups/{id} | Get a component group
[**GetPagesPageIdComponentGroupsIdUptime**](ComponentGroupsApi.md#GetPagesPageIdComponentGroupsIdUptime) | **Get** /pages/{page_id}/component-groups/{id}/uptime | Get uptime data for a component group
[**PatchPagesPageIdComponentGroupsId**](ComponentGroupsApi.md#PatchPagesPageIdComponentGroupsId) | **Patch** /pages/{page_id}/component-groups/{id} | Update a component group
[**PostPagesPageIdComponentGroups**](ComponentGroupsApi.md#PostPagesPageIdComponentGroups) | **Post** /pages/{page_id}/component-groups | Create a component group
[**PutPagesPageIdComponentGroupsId**](ComponentGroupsApi.md#PutPagesPageIdComponentGroupsId) | **Put** /pages/{page_id}/component-groups/{id} | Update a component group



## DeletePagesPageIdComponentGroupsId

> GroupComponent DeletePagesPageIdComponentGroupsId(ctx, pageId, id).Execute()

Delete a component group



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
    id := "id_example" // string | Component group identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentGroupsApi.DeletePagesPageIdComponentGroupsId(context.Background(), pageId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsApi.DeletePagesPageIdComponentGroupsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdComponentGroupsId`: GroupComponent
    fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsApi.DeletePagesPageIdComponentGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdComponentGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentGroups

> []GroupComponent GetPagesPageIdComponentGroups(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of component groups



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
    page := int32(56) // int32 | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. (optional)
    perPage := int32(56) // int32 | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentGroupsApi.GetPagesPageIdComponentGroups(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsApi.GetPagesPageIdComponentGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdComponentGroups`: []GroupComponent
    fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsApi.GetPagesPageIdComponentGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

### Return type

[**[]GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentGroupsId

> GroupComponent GetPagesPageIdComponentGroupsId(ctx, pageId, id).Execute()

Get a component group



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
    id := "id_example" // string | Component group identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentGroupsApi.GetPagesPageIdComponentGroupsId(context.Background(), pageId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsApi.GetPagesPageIdComponentGroupsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdComponentGroupsId`: GroupComponent
    fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsApi.GetPagesPageIdComponentGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentGroupsIdUptime

> ComponentGroupUptime GetPagesPageIdComponentGroupsIdUptime(ctx, pageId, id).Execute()

Get uptime data for a component group



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
    id := "id_example" // string | Component group identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentGroupsApi.GetPagesPageIdComponentGroupsIdUptime(context.Background(), pageId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsApi.GetPagesPageIdComponentGroupsIdUptime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdComponentGroupsIdUptime`: ComponentGroupUptime
    fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsApi.GetPagesPageIdComponentGroupsIdUptime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentGroupsIdUptimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ComponentGroupUptime**](ComponentGroupUptime.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdComponentGroupsId

> GroupComponent PatchPagesPageIdComponentGroupsId(ctx, pageId, id).PatchPagesPageIdComponentGroups(patchPagesPageIdComponentGroups).Execute()

Update a component group



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
    id := "id_example" // string | Component group identifier
    patchPagesPageIdComponentGroups := *openapiclient.NewPatchPagesPageIdComponentGroups() // PatchPagesPageIdComponentGroups | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentGroupsApi.PatchPagesPageIdComponentGroupsId(context.Background(), pageId, id).PatchPagesPageIdComponentGroups(patchPagesPageIdComponentGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsApi.PatchPagesPageIdComponentGroupsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageIdComponentGroupsId`: GroupComponent
    fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsApi.PatchPagesPageIdComponentGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdComponentGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdComponentGroups** | [**PatchPagesPageIdComponentGroups**](PatchPagesPageIdComponentGroups.md) |  | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdComponentGroups

> GroupComponent PostPagesPageIdComponentGroups(ctx, pageId).PostPagesPageIdComponentGroups(postPagesPageIdComponentGroups).Execute()

Create a component group



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
    postPagesPageIdComponentGroups := *openapiclient.NewPostPagesPageIdComponentGroups() // PostPagesPageIdComponentGroups | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentGroupsApi.PostPagesPageIdComponentGroups(context.Background(), pageId).PostPagesPageIdComponentGroups(postPagesPageIdComponentGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsApi.PostPagesPageIdComponentGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdComponentGroups`: GroupComponent
    fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsApi.PostPagesPageIdComponentGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdComponentGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdComponentGroups** | [**PostPagesPageIdComponentGroups**](PostPagesPageIdComponentGroups.md) |  | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdComponentGroupsId

> GroupComponent PutPagesPageIdComponentGroupsId(ctx, pageId, id).PutPagesPageIdComponentGroups(putPagesPageIdComponentGroups).Execute()

Update a component group



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
    id := "id_example" // string | Component group identifier
    putPagesPageIdComponentGroups := *openapiclient.NewPutPagesPageIdComponentGroups() // PutPagesPageIdComponentGroups | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentGroupsApi.PutPagesPageIdComponentGroupsId(context.Background(), pageId, id).PutPagesPageIdComponentGroups(putPagesPageIdComponentGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsApi.PutPagesPageIdComponentGroupsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdComponentGroupsId`: GroupComponent
    fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsApi.PutPagesPageIdComponentGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdComponentGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdComponentGroups** | [**PutPagesPageIdComponentGroups**](PutPagesPageIdComponentGroups.md) |  | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

