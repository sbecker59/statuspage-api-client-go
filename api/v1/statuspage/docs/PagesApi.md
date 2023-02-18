# \PagesApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPages**](PagesApi.md#GetPages) | **Get** /pages | Get a list of pages
[**GetPagesPageId**](PagesApi.md#GetPagesPageId) | **Get** /pages/{page_id} | Get a page
[**PatchPagesPageId**](PagesApi.md#PatchPagesPageId) | **Patch** /pages/{page_id} | Update a page
[**PutPagesPageId**](PagesApi.md#PutPagesPageId) | **Put** /pages/{page_id} | Update a page



## GetPages

> []Page GetPages(ctx).Execute()

Get a list of pages



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PagesApi.GetPages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagesApi.GetPages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPages`: []Page
    fmt.Fprintf(os.Stdout, "Response from `PagesApi.GetPages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesRequest struct via the builder pattern


### Return type

[**[]Page**](Page.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageId

> Page GetPagesPageId(ctx, pageId).Execute()

Get a page



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PagesApi.GetPagesPageId(context.Background(), pageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagesApi.GetPagesPageId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageId`: Page
    fmt.Fprintf(os.Stdout, "Response from `PagesApi.GetPagesPageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Page**](Page.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageId

> Page PatchPagesPageId(ctx, pageId).PatchPages(patchPages).Execute()

Update a page



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
    patchPages := *openapiclient.NewPatchPages() // PatchPages | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PagesApi.PatchPagesPageId(context.Background(), pageId).PatchPages(patchPages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagesApi.PatchPagesPageId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageId`: Page
    fmt.Fprintf(os.Stdout, "Response from `PagesApi.PatchPagesPageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchPages** | [**PatchPages**](PatchPages.md) |  | 

### Return type

[**Page**](Page.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageId

> Page PutPagesPageId(ctx, pageId).PutPages(putPages).Execute()

Update a page



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
    putPages := *openapiclient.NewPutPages() // PutPages | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PagesApi.PutPagesPageId(context.Background(), pageId).PutPages(putPages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PagesApi.PutPagesPageId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageId`: Page
    fmt.Fprintf(os.Stdout, "Response from `PagesApi.PutPagesPageId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putPages** | [**PutPages**](PutPages.md) |  | 

### Return type

[**Page**](Page.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

