# \PageAccessGroupsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessGroupsPageAccessGroupId**](PageAccessGroupsApi.md#DeletePagesPageIdPageAccessGroupsPageAccessGroupId) | **Delete** /pages/{page_id}/page_access_groups/{page_access_group_id} | Remove a page access group
[**GetPagesPageIdPageAccessGroups**](PageAccessGroupsApi.md#GetPagesPageIdPageAccessGroups) | **Get** /pages/{page_id}/page_access_groups | Get a list of page access groups
[**GetPagesPageIdPageAccessGroupsPageAccessGroupId**](PageAccessGroupsApi.md#GetPagesPageIdPageAccessGroupsPageAccessGroupId) | **Get** /pages/{page_id}/page_access_groups/{page_access_group_id} | Get a page access group
[**PatchPagesPageIdPageAccessGroupsPageAccessGroupId**](PageAccessGroupsApi.md#PatchPagesPageIdPageAccessGroupsPageAccessGroupId) | **Patch** /pages/{page_id}/page_access_groups/{page_access_group_id} | Update a page access group
[**PostPagesPageIdPageAccessGroups**](PageAccessGroupsApi.md#PostPagesPageIdPageAccessGroups) | **Post** /pages/{page_id}/page_access_groups | Create a page access group
[**PutPagesPageIdPageAccessGroupsPageAccessGroupId**](PageAccessGroupsApi.md#PutPagesPageIdPageAccessGroupsPageAccessGroupId) | **Put** /pages/{page_id}/page_access_groups/{page_access_group_id} | Update a page access group



## DeletePagesPageIdPageAccessGroupsPageAccessGroupId

> PageAccessGroup DeletePagesPageIdPageAccessGroupsPageAccessGroupId(ctx, pageId, pageAccessGroupId).Execute()

Remove a page access group



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
    pageAccessGroupId := "pageAccessGroupId_example" // string | Page Access Group Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupsApi.DeletePagesPageIdPageAccessGroupsPageAccessGroupId(context.Background(), pageId, pageAccessGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupsApi.DeletePagesPageIdPageAccessGroupsPageAccessGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdPageAccessGroupsPageAccessGroupId`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupsApi.DeletePagesPageIdPageAccessGroupsPageAccessGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdPageAccessGroupsPageAccessGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdPageAccessGroups

> []PageAccessGroup GetPagesPageIdPageAccessGroups(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of page access groups



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
    resp, r, err := api_client.PageAccessGroupsApi.GetPagesPageIdPageAccessGroups(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupsApi.GetPagesPageIdPageAccessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdPageAccessGroups`: []PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupsApi.GetPagesPageIdPageAccessGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdPageAccessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

### Return type

[**[]PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdPageAccessGroupsPageAccessGroupId

> PageAccessGroup GetPagesPageIdPageAccessGroupsPageAccessGroupId(ctx, pageId, pageAccessGroupId).Execute()

Get a page access group



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
    pageAccessGroupId := "pageAccessGroupId_example" // string | Page Access Group Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupsApi.GetPagesPageIdPageAccessGroupsPageAccessGroupId(context.Background(), pageId, pageAccessGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupsApi.GetPagesPageIdPageAccessGroupsPageAccessGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdPageAccessGroupsPageAccessGroupId`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupsApi.GetPagesPageIdPageAccessGroupsPageAccessGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdPageAccessGroupsPageAccessGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdPageAccessGroupsPageAccessGroupId

> PageAccessGroup PatchPagesPageIdPageAccessGroupsPageAccessGroupId(ctx, pageId, pageAccessGroupId).PatchPagesPageIdPageAccessGroups(patchPagesPageIdPageAccessGroups).Execute()

Update a page access group



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
    pageAccessGroupId := "pageAccessGroupId_example" // string | Page Access Group Identifier
    patchPagesPageIdPageAccessGroups := *openapiclient.NewPatchPagesPageIdPageAccessGroups() // PatchPagesPageIdPageAccessGroups | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupsApi.PatchPagesPageIdPageAccessGroupsPageAccessGroupId(context.Background(), pageId, pageAccessGroupId).PatchPagesPageIdPageAccessGroups(patchPagesPageIdPageAccessGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupsApi.PatchPagesPageIdPageAccessGroupsPageAccessGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageIdPageAccessGroupsPageAccessGroupId`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupsApi.PatchPagesPageIdPageAccessGroupsPageAccessGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdPageAccessGroups** | [**PatchPagesPageIdPageAccessGroups**](PatchPagesPageIdPageAccessGroups.md) |  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdPageAccessGroups

> PageAccessGroup PostPagesPageIdPageAccessGroups(ctx, pageId).PostPagesPageIdPageAccessGroups(postPagesPageIdPageAccessGroups).Execute()

Create a page access group



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
    postPagesPageIdPageAccessGroups := *openapiclient.NewPostPagesPageIdPageAccessGroups() // PostPagesPageIdPageAccessGroups | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupsApi.PostPagesPageIdPageAccessGroups(context.Background(), pageId).PostPagesPageIdPageAccessGroups(postPagesPageIdPageAccessGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupsApi.PostPagesPageIdPageAccessGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdPageAccessGroups`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupsApi.PostPagesPageIdPageAccessGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdPageAccessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdPageAccessGroups** | [**PostPagesPageIdPageAccessGroups**](PostPagesPageIdPageAccessGroups.md) |  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdPageAccessGroupsPageAccessGroupId

> PageAccessGroup PutPagesPageIdPageAccessGroupsPageAccessGroupId(ctx, pageId, pageAccessGroupId).PutPagesPageIdPageAccessGroups(putPagesPageIdPageAccessGroups).Execute()

Update a page access group



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
    pageAccessGroupId := "pageAccessGroupId_example" // string | Page Access Group Identifier
    putPagesPageIdPageAccessGroups := *openapiclient.NewPutPagesPageIdPageAccessGroups() // PutPagesPageIdPageAccessGroups | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupsApi.PutPagesPageIdPageAccessGroupsPageAccessGroupId(context.Background(), pageId, pageAccessGroupId).PutPagesPageIdPageAccessGroups(putPagesPageIdPageAccessGroups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupsApi.PutPagesPageIdPageAccessGroupsPageAccessGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdPageAccessGroupsPageAccessGroupId`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupsApi.PutPagesPageIdPageAccessGroupsPageAccessGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdPageAccessGroupsPageAccessGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdPageAccessGroups** | [**PutPagesPageIdPageAccessGroups**](PutPagesPageIdPageAccessGroups.md) |  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

