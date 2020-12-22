# \ComponentsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdComponentsComponentId**](ComponentsApi.md#DeletePagesPageIdComponentsComponentId) | **Delete** /pages/{page_id}/components/{component_id} | Delete a component
[**DeletePagesPageIdComponentsComponentIdPageAccessGroups**](ComponentsApi.md#DeletePagesPageIdComponentsComponentIdPageAccessGroups) | **Delete** /pages/{page_id}/components/{component_id}/page_access_groups | Remove page access groups from a component
[**DeletePagesPageIdComponentsComponentIdPageAccessUsers**](ComponentsApi.md#DeletePagesPageIdComponentsComponentIdPageAccessUsers) | **Delete** /pages/{page_id}/components/{component_id}/page_access_users | Remove page access users from component
[**GetPagesPageIdComponents**](ComponentsApi.md#GetPagesPageIdComponents) | **Get** /pages/{page_id}/components | Get a list of components
[**GetPagesPageIdComponentsComponentId**](ComponentsApi.md#GetPagesPageIdComponentsComponentId) | **Get** /pages/{page_id}/components/{component_id} | Get a component
[**GetPagesPageIdComponentsComponentIdUptime**](ComponentsApi.md#GetPagesPageIdComponentsComponentIdUptime) | **Get** /pages/{page_id}/components/{component_id}/uptime | Get uptime data for a component
[**PatchPagesPageIdComponentsComponentId**](ComponentsApi.md#PatchPagesPageIdComponentsComponentId) | **Patch** /pages/{page_id}/components/{component_id} | Update a component
[**PostPagesPageIdComponents**](ComponentsApi.md#PostPagesPageIdComponents) | **Post** /pages/{page_id}/components | Create a component
[**PostPagesPageIdComponentsComponentIdPageAccessGroups**](ComponentsApi.md#PostPagesPageIdComponentsComponentIdPageAccessGroups) | **Post** /pages/{page_id}/components/{component_id}/page_access_groups | Add page access groups to a component
[**PostPagesPageIdComponentsComponentIdPageAccessUsers**](ComponentsApi.md#PostPagesPageIdComponentsComponentIdPageAccessUsers) | **Post** /pages/{page_id}/components/{component_id}/page_access_users | Add page access users to a component
[**PutPagesPageIdComponentsComponentId**](ComponentsApi.md#PutPagesPageIdComponentsComponentId) | **Put** /pages/{page_id}/components/{component_id} | Update a component



## DeletePagesPageIdComponentsComponentId

> DeletePagesPageIdComponentsComponentId(ctx, pageId, componentId).Execute()

Delete a component



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
    componentId := "componentId_example" // string | Component identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.DeletePagesPageIdComponentsComponentId(context.Background(), pageId, componentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.DeletePagesPageIdComponentsComponentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdComponentsComponentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePagesPageIdComponentsComponentIdPageAccessGroups

> Component DeletePagesPageIdComponentsComponentIdPageAccessGroups(ctx, pageId, componentId).Execute()

Remove page access groups from a component



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
    componentId := "componentId_example" // string | Component identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.DeletePagesPageIdComponentsComponentIdPageAccessGroups(context.Background(), pageId, componentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.DeletePagesPageIdComponentsComponentIdPageAccessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdComponentsComponentIdPageAccessGroups`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.DeletePagesPageIdComponentsComponentIdPageAccessGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdComponentsComponentIdPageAccessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePagesPageIdComponentsComponentIdPageAccessUsers

> Component DeletePagesPageIdComponentsComponentIdPageAccessUsers(ctx, pageId, componentId).Execute()

Remove page access users from component



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
    componentId := "componentId_example" // string | Component identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.DeletePagesPageIdComponentsComponentIdPageAccessUsers(context.Background(), pageId, componentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.DeletePagesPageIdComponentsComponentIdPageAccessUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdComponentsComponentIdPageAccessUsers`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.DeletePagesPageIdComponentsComponentIdPageAccessUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdComponentsComponentIdPageAccessUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponents

> []Component GetPagesPageIdComponents(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of components



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
    page := int32(56) // int32 | Page offset to fetch. (optional)
    perPage := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.GetPagesPageIdComponents(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.GetPagesPageIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdComponents`: []Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.GetPagesPageIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. | 
 **perPage** | **int32** | Number of results to return per page. | 

### Return type

[**[]Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentsComponentId

> Component GetPagesPageIdComponentsComponentId(ctx, pageId, componentId).Execute()

Get a component



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
    componentId := "componentId_example" // string | Component identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.GetPagesPageIdComponentsComponentId(context.Background(), pageId, componentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.GetPagesPageIdComponentsComponentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdComponentsComponentId`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.GetPagesPageIdComponentsComponentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentsComponentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentsComponentIdUptime

> ComponentUptime GetPagesPageIdComponentsComponentIdUptime(ctx, pageId, componentId).Start(start).End(end).Execute()

Get uptime data for a component



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
    componentId := "componentId_example" // string | Component identifier
    start := TODO // PartialStartDate | The start date for uptime calculation (defaults to the component's start_date field or 90 days ago, whichever is more recent). The maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year. If the year and month are given, the start date defaults to the first day of that month. The earliest supported date is January 1, 1970.  (optional)
    end := TODO // PartialEndDate | The end date for uptime calculation (defaults to today in the page's time zone). The maximum supported date range is six calendar months. If the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month. The earliest supported date is January 1, 1970.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.GetPagesPageIdComponentsComponentIdUptime(context.Background(), pageId, componentId).Start(start).End(end).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.GetPagesPageIdComponentsComponentIdUptime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdComponentsComponentIdUptime`: ComponentUptime
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.GetPagesPageIdComponentsComponentIdUptime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentsComponentIdUptimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | [**PartialStartDate**](PartialStartDate.md) | The start date for uptime calculation (defaults to the component&#39;s start_date field or 90 days ago, whichever is more recent). The maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year. If the year and month are given, the start date defaults to the first day of that month. The earliest supported date is January 1, 1970.  | 
 **end** | [**PartialEndDate**](PartialEndDate.md) | The end date for uptime calculation (defaults to today in the page&#39;s time zone). The maximum supported date range is six calendar months. If the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month. The earliest supported date is January 1, 1970.  | 

### Return type

[**ComponentUptime**](ComponentUptime.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdComponentsComponentId

> Component PatchPagesPageIdComponentsComponentId(ctx, pageId, componentId).PatchPagesPageIdComponents(patchPagesPageIdComponents).Execute()

Update a component



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
    componentId := "componentId_example" // string | Component identifier
    patchPagesPageIdComponents := *openapiclient.NewPatchPagesPageIdComponents() // PatchPagesPageIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.PatchPagesPageIdComponentsComponentId(context.Background(), pageId, componentId).PatchPagesPageIdComponents(patchPagesPageIdComponents).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.PatchPagesPageIdComponentsComponentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageIdComponentsComponentId`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.PatchPagesPageIdComponentsComponentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdComponentsComponentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdComponents** | [**PatchPagesPageIdComponents**](PatchPagesPageIdComponents.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdComponents

> Component PostPagesPageIdComponents(ctx, pageId).PostPagesPageIdComponents(postPagesPageIdComponents).Execute()

Create a component



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
    postPagesPageIdComponents := *openapiclient.NewPostPagesPageIdComponents() // PostPagesPageIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.PostPagesPageIdComponents(context.Background(), pageId).PostPagesPageIdComponents(postPagesPageIdComponents).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.PostPagesPageIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdComponents`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.PostPagesPageIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdComponents** | [**PostPagesPageIdComponents**](PostPagesPageIdComponents.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdComponentsComponentIdPageAccessGroups

> Component PostPagesPageIdComponentsComponentIdPageAccessGroups(ctx, pageId, componentId).Execute()

Add page access groups to a component



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
    componentId := "componentId_example" // string | Component identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.PostPagesPageIdComponentsComponentIdPageAccessGroups(context.Background(), pageId, componentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.PostPagesPageIdComponentsComponentIdPageAccessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdComponentsComponentIdPageAccessGroups`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.PostPagesPageIdComponentsComponentIdPageAccessGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdComponentsComponentIdPageAccessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdComponentsComponentIdPageAccessUsers

> Component PostPagesPageIdComponentsComponentIdPageAccessUsers(ctx, pageId, componentId).PageAccessUserIds(pageAccessUserIds).Execute()

Add page access users to a component



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
    componentId := "componentId_example" // string | Component identifier
    pageAccessUserIds := []string{"Inner_example"} // []string | List of page access users to add to component

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.PostPagesPageIdComponentsComponentIdPageAccessUsers(context.Background(), pageId, componentId).PageAccessUserIds(pageAccessUserIds).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.PostPagesPageIdComponentsComponentIdPageAccessUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdComponentsComponentIdPageAccessUsers`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.PostPagesPageIdComponentsComponentIdPageAccessUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdComponentsComponentIdPageAccessUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageAccessUserIds** | **[]string** | List of page access users to add to component | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdComponentsComponentId

> Component PutPagesPageIdComponentsComponentId(ctx, pageId, componentId).PutPagesPageIdComponents(putPagesPageIdComponents).Execute()

Update a component



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
    componentId := "componentId_example" // string | Component identifier
    putPagesPageIdComponents := *openapiclient.NewPutPagesPageIdComponents() // PutPagesPageIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ComponentsApi.PutPagesPageIdComponentsComponentId(context.Background(), pageId, componentId).PutPagesPageIdComponents(putPagesPageIdComponents).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsApi.PutPagesPageIdComponentsComponentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdComponentsComponentId`: Component
    fmt.Fprintf(os.Stdout, "Response from `ComponentsApi.PutPagesPageIdComponentsComponentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdComponentsComponentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdComponents** | [**PutPagesPageIdComponents**](PutPagesPageIdComponents.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

