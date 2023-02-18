# \TemplatesApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPagesPageIdIncidentTemplates**](TemplatesApi.md#GetPagesPageIdIncidentTemplates) | **Get** /pages/{page_id}/incident_templates | Get a list of templates
[**PostPagesPageIdIncidentTemplates**](TemplatesApi.md#PostPagesPageIdIncidentTemplates) | **Post** /pages/{page_id}/incident_templates | Create a template



## GetPagesPageIdIncidentTemplates

> []IncidentTemplate GetPagesPageIdIncidentTemplates(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of templates



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
    page := int32(56) // int32 | Page offset to fetch. (optional) (default to 1)
    perPage := int32(56) // int32 | Number of results to return per page. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.GetPagesPageIdIncidentTemplates(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.GetPagesPageIdIncidentTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdIncidentTemplates`: []IncidentTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.GetPagesPageIdIncidentTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. | [default to 1]
 **perPage** | **int32** | Number of results to return per page. | [default to 100]

### Return type

[**[]IncidentTemplate**](IncidentTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdIncidentTemplates

> IncidentTemplate PostPagesPageIdIncidentTemplates(ctx, pageId).PostPagesPageIdIncidentTemplates(postPagesPageIdIncidentTemplates).Execute()

Create a template



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
    postPagesPageIdIncidentTemplates := *openapiclient.NewPostPagesPageIdIncidentTemplates() // PostPagesPageIdIncidentTemplates | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.PostPagesPageIdIncidentTemplates(context.Background(), pageId).PostPagesPageIdIncidentTemplates(postPagesPageIdIncidentTemplates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.PostPagesPageIdIncidentTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdIncidentTemplates`: IncidentTemplate
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.PostPagesPageIdIncidentTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdIncidentTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdIncidentTemplates** | [**PostPagesPageIdIncidentTemplates**](PostPagesPageIdIncidentTemplates.md) |  | 

### Return type

[**IncidentTemplate**](IncidentTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

