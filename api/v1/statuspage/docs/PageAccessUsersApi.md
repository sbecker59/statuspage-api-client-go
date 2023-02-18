# \PageAccessUsersApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessUsersPageAccessUserId**](PageAccessUsersApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserId) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id} | Delete page access user
[**GetPagesPageIdPageAccessUsers**](PageAccessUsersApi.md#GetPagesPageIdPageAccessUsers) | **Get** /pages/{page_id}/page_access_users | Get a list of page access users
[**GetPagesPageIdPageAccessUsersPageAccessUserId**](PageAccessUsersApi.md#GetPagesPageIdPageAccessUsersPageAccessUserId) | **Get** /pages/{page_id}/page_access_users/{page_access_user_id} | Get page access user
[**PatchPagesPageIdPageAccessUsersPageAccessUserId**](PageAccessUsersApi.md#PatchPagesPageIdPageAccessUsersPageAccessUserId) | **Patch** /pages/{page_id}/page_access_users/{page_access_user_id} | Update page access user
[**PostPagesPageIdPageAccessUsers**](PageAccessUsersApi.md#PostPagesPageIdPageAccessUsers) | **Post** /pages/{page_id}/page_access_users | Add a page access user
[**PutPagesPageIdPageAccessUsersPageAccessUserId**](PageAccessUsersApi.md#PutPagesPageIdPageAccessUsersPageAccessUserId) | **Put** /pages/{page_id}/page_access_users/{page_access_user_id} | Update page access user



## DeletePagesPageIdPageAccessUsersPageAccessUserId

> DeletePagesPageIdPageAccessUsersPageAccessUserId(ctx, pageId, pageAccessUserId).Execute()

Delete page access user



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUsersApi.DeletePagesPageIdPageAccessUsersPageAccessUserId(context.Background(), pageId, pageAccessUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUsersApi.DeletePagesPageIdPageAccessUsersPageAccessUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdPageAccessUsersPageAccessUserIdRequest struct via the builder pattern


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


## GetPagesPageIdPageAccessUsers

> []PageAccessUser GetPagesPageIdPageAccessUsers(ctx, pageId).Email(email).Page(page).PerPage(perPage).Execute()

Get a list of page access users



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
    email := "email_example" // string | Email address to search for (optional)
    page := int32(56) // int32 | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. (optional)
    perPage := int32(56) // int32 | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUsersApi.GetPagesPageIdPageAccessUsers(context.Background(), pageId).Email(email).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUsersApi.GetPagesPageIdPageAccessUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdPageAccessUsers`: []PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUsersApi.GetPagesPageIdPageAccessUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdPageAccessUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **string** | Email address to search for | 
 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

### Return type

[**[]PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdPageAccessUsersPageAccessUserId

> PageAccessUser GetPagesPageIdPageAccessUsersPageAccessUserId(ctx, pageId, pageAccessUserId).Execute()

Get page access user



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUsersApi.GetPagesPageIdPageAccessUsersPageAccessUserId(context.Background(), pageId, pageAccessUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUsersApi.GetPagesPageIdPageAccessUsersPageAccessUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdPageAccessUsersPageAccessUserId`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUsersApi.GetPagesPageIdPageAccessUsersPageAccessUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdPageAccessUsersPageAccessUserIdRequest struct via the builder pattern


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


## PatchPagesPageIdPageAccessUsersPageAccessUserId

> PageAccessUser PatchPagesPageIdPageAccessUsersPageAccessUserId(ctx, pageId, pageAccessUserId).Execute()

Update page access user



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUsersApi.PatchPagesPageIdPageAccessUsersPageAccessUserId(context.Background(), pageId, pageAccessUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUsersApi.PatchPagesPageIdPageAccessUsersPageAccessUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageIdPageAccessUsersPageAccessUserId`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUsersApi.PatchPagesPageIdPageAccessUsersPageAccessUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdPageAccessUsersPageAccessUserIdRequest struct via the builder pattern


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


## PostPagesPageIdPageAccessUsers

> PageAccessUser PostPagesPageIdPageAccessUsers(ctx, pageId).PostPagesPageIdPageAccessUsers(postPagesPageIdPageAccessUsers).Execute()

Add a page access user



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
    postPagesPageIdPageAccessUsers := *openapiclient.NewPostPagesPageIdPageAccessUsers() // PostPagesPageIdPageAccessUsers | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUsersApi.PostPagesPageIdPageAccessUsers(context.Background(), pageId).PostPagesPageIdPageAccessUsers(postPagesPageIdPageAccessUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUsersApi.PostPagesPageIdPageAccessUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdPageAccessUsers`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUsersApi.PostPagesPageIdPageAccessUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdPageAccessUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdPageAccessUsers** | [**PostPagesPageIdPageAccessUsers**](PostPagesPageIdPageAccessUsers.md) |  | 

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


## PutPagesPageIdPageAccessUsersPageAccessUserId

> PageAccessUser PutPagesPageIdPageAccessUsersPageAccessUserId(ctx, pageId, pageAccessUserId).Execute()

Update page access user



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PageAccessUsersApi.PutPagesPageIdPageAccessUsersPageAccessUserId(context.Background(), pageId, pageAccessUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PageAccessUsersApi.PutPagesPageIdPageAccessUsersPageAccessUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdPageAccessUsersPageAccessUserId`: PageAccessUser
    fmt.Fprintf(os.Stdout, "Response from `PageAccessUsersApi.PutPagesPageIdPageAccessUsersPageAccessUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**pageAccessUserId** | **string** | Page Access User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdPageAccessUsersPageAccessUserIdRequest struct via the builder pattern


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

