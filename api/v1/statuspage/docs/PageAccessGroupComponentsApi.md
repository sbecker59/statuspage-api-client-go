# \PageAccessGroupComponentsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Delete** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Delete components for a page access group
[**DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId**](PageAccessGroupComponentsApi.md#DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId) | **Delete** /pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id} | Remove a component from a page access group
[**GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Get** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | List components for a page access group
[**PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Patch** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Add components to page access group
[**PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Post** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Replace components for a page access group
[**PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Put** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Add components to page access group



## DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents

> PageAccessGroup DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId).DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents(deletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents).Execute()

Delete components for a page access group



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
    deletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents := *openapiclient.NewDeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents() // DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupComponentsApi.DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents(context.Background(), pageId, pageAccessGroupId).DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents(deletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupComponentsApi.DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupComponentsApi.DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents** | [**DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md) |  | 

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


## DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId

> PageAccessGroup DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId(ctx, pageId, pageAccessGroupId, componentId).Execute()

Remove a component from a page access group



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
    componentId := "componentId_example" // string | Component identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupComponentsApi.DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId(context.Background(), pageId, pageAccessGroupId, componentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupComponentsApi.DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupComponentsApi.DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentIdRequest struct via the builder pattern


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


## GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents

> []Component GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId).Execute()

List components for a page access group



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
    resp, r, err := api_client.PageAccessGroupComponentsApi.GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(context.Background(), pageId, pageAccessGroupId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupComponentsApi.GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: []Component
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupComponentsApi.GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdPageAccessGroupsPageAccessGroupIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents

> PageAccessGroup PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId).PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(patchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents).Execute()

Add components to page access group



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
    patchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents := *openapiclient.NewPatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents() // PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupComponentsApi.PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(context.Background(), pageId, pageAccessGroupId).PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(patchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupComponentsApi.PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupComponentsApi.PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents** | [**PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md) |  | 

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


## PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents

> PageAccessGroup PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId).PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(postPagesPageIdPageAccessGroupsPageAccessGroupIdComponents).Execute()

Replace components for a page access group



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
    postPagesPageIdPageAccessGroupsPageAccessGroupIdComponents := *openapiclient.NewPostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents([]string{"ComponentIds_example"}) // PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupComponentsApi.PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(context.Background(), pageId, pageAccessGroupId).PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(postPagesPageIdPageAccessGroupsPageAccessGroupIdComponents).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupComponentsApi.PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupComponentsApi.PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdPageAccessGroupsPageAccessGroupIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postPagesPageIdPageAccessGroupsPageAccessGroupIdComponents** | [**PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md) |  | 

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


## PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents

> PageAccessGroup PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId).PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(putPagesPageIdPageAccessGroupsPageAccessGroupIdComponents).Execute()

Add components to page access group



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
    putPagesPageIdPageAccessGroupsPageAccessGroupIdComponents := *openapiclient.NewPutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents() // PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessGroupComponentsApi.PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(context.Background(), pageId, pageAccessGroupId).PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(putPagesPageIdPageAccessGroupsPageAccessGroupIdComponents).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessGroupComponentsApi.PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: PageAccessGroup
    fmt.Fprintf(os.Stdout, "Response from `PageAccessGroupComponentsApi.PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessGroupId** | **string** | Page Access Group Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdPageAccessGroupsPageAccessGroupIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdPageAccessGroupsPageAccessGroupIdComponents** | [**PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md) |  | 

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

