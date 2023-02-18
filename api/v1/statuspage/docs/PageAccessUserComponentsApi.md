# \PageAccessUserComponentsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Remove components for page access user
[**DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId**](PageAccessUserComponentsApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id} | Remove component for page access user
[**GetPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#GetPagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Get** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Get components for page access user
[**PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Patch** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Add components for page access user
[**PostPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#PostPagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Post** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Replace components for page access user
[**PutPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#PutPagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Put** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Add components for page access user



## DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents

> PageAccessUser DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId).DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents(deletePagesPageIdPageAccessUsersPageAccessUserIdComponents).Execute()

Remove components for page access user



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
    deletePagesPageIdPageAccessUsersPageAccessUserIdComponents := *openapiclient.NewDeletePagesPageIdPageAccessUsersPageAccessUserIdComponents() // DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserComponentsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents(context.Background(), pageId, pageAccessUserId).DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents(deletePagesPageIdPageAccessUsersPageAccessUserIdComponents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserComponentsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserComponentsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deletePagesPageIdPageAccessUsersPageAccessUserIdComponents** | [**DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents**](DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents.md) |  | 

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


## DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId

> PageAccessUser DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId(ctx, pageId, pageAccessUserId, componentId).Execute()

Remove component for page access user



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
    componentId := "componentId_example" // string | Component identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserComponentsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId(context.Background(), pageId, pageAccessUserId, componentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserComponentsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserComponentsApi.DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentIdRequest struct via the builder pattern


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


## GetPagesPageIdPageAccessUsersPageAccessUserIdComponents

> []Component GetPagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId).Page(page).PerPage(perPage).Execute()

Get components for page access user



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
    resp, r, err := api_client.PageAccessUserComponentsApi.GetPagesPageIdPageAccessUsersPageAccessUserIdComponents(context.Background(), pageId, pageAccessUserId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserComponentsApi.GetPagesPageIdPageAccessUsersPageAccessUserIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdPageAccessUsersPageAccessUserIdComponents`: []Component
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserComponentsApi.GetPagesPageIdPageAccessUsersPageAccessUserIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdPageAccessUsersPageAccessUserIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

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


## PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents

> PageAccessUser PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId).PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents(patchPagesPageIdPageAccessUsersPageAccessUserIdComponents).Execute()

Add components for page access user



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
    patchPagesPageIdPageAccessUsersPageAccessUserIdComponents := *openapiclient.NewPatchPagesPageIdPageAccessUsersPageAccessUserIdComponents([]string{"ComponentIds_example"}) // PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserComponentsApi.PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents(context.Background(), pageId, pageAccessUserId).PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents(patchPagesPageIdPageAccessUsersPageAccessUserIdComponents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserComponentsApi.PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserComponentsApi.PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdPageAccessUsersPageAccessUserIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdPageAccessUsersPageAccessUserIdComponents** | [**PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents.md) |  | 

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


## PostPagesPageIdPageAccessUsersPageAccessUserIdComponents

> PageAccessUser PostPagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId).PostPagesPageIdPageAccessUsersPageAccessUserIdComponents(postPagesPageIdPageAccessUsersPageAccessUserIdComponents).Execute()

Replace components for page access user



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
    postPagesPageIdPageAccessUsersPageAccessUserIdComponents := *openapiclient.NewPostPagesPageIdPageAccessUsersPageAccessUserIdComponents([]string{"ComponentIds_example"}) // PostPagesPageIdPageAccessUsersPageAccessUserIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserComponentsApi.PostPagesPageIdPageAccessUsersPageAccessUserIdComponents(context.Background(), pageId, pageAccessUserId).PostPagesPageIdPageAccessUsersPageAccessUserIdComponents(postPagesPageIdPageAccessUsersPageAccessUserIdComponents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserComponentsApi.PostPagesPageIdPageAccessUsersPageAccessUserIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdPageAccessUsersPageAccessUserIdComponents`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserComponentsApi.PostPagesPageIdPageAccessUsersPageAccessUserIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdPageAccessUsersPageAccessUserIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postPagesPageIdPageAccessUsersPageAccessUserIdComponents** | [**PostPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PostPagesPageIdPageAccessUsersPageAccessUserIdComponents.md) |  | 

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


## PutPagesPageIdPageAccessUsersPageAccessUserIdComponents

> PageAccessUser PutPagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId).PutPagesPageIdPageAccessUsersPageAccessUserIdComponents(putPagesPageIdPageAccessUsersPageAccessUserIdComponents).Execute()

Add components for page access user



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
    putPagesPageIdPageAccessUsersPageAccessUserIdComponents := *openapiclient.NewPutPagesPageIdPageAccessUsersPageAccessUserIdComponents([]string{"ComponentIds_example"}) // PutPagesPageIdPageAccessUsersPageAccessUserIdComponents | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUserComponentsApi.PutPagesPageIdPageAccessUsersPageAccessUserIdComponents(context.Background(), pageId, pageAccessUserId).PutPagesPageIdPageAccessUsersPageAccessUserIdComponents(putPagesPageIdPageAccessUsersPageAccessUserIdComponents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUserComponentsApi.PutPagesPageIdPageAccessUsersPageAccessUserIdComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdPageAccessUsersPageAccessUserIdComponents`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUserComponentsApi.PutPagesPageIdPageAccessUsersPageAccessUserIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdPageAccessUsersPageAccessUserIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdPageAccessUsersPageAccessUserIdComponents** | [**PutPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PutPagesPageIdPageAccessUsersPageAccessUserIdComponents.md) |  | 

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

