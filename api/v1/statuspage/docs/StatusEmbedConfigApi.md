# \StatusEmbedConfigApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPagesPageIdStatusEmbedConfig**](StatusEmbedConfigApi.md#GetPagesPageIdStatusEmbedConfig) | **Get** /pages/{page_id}/status_embed_config | Get status embed config settings
[**PatchPagesPageIdStatusEmbedConfig**](StatusEmbedConfigApi.md#PatchPagesPageIdStatusEmbedConfig) | **Patch** /pages/{page_id}/status_embed_config | Update status embed config settings
[**PutPagesPageIdStatusEmbedConfig**](StatusEmbedConfigApi.md#PutPagesPageIdStatusEmbedConfig) | **Put** /pages/{page_id}/status_embed_config | Update status embed config settings



## GetPagesPageIdStatusEmbedConfig

> StatusEmbedConfig GetPagesPageIdStatusEmbedConfig(ctx, pageId).Execute()

Get status embed config settings



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
    resp, r, err := api_client.StatusEmbedConfigApi.GetPagesPageIdStatusEmbedConfig(context.Background(), pageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusEmbedConfigApi.GetPagesPageIdStatusEmbedConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdStatusEmbedConfig`: StatusEmbedConfig
    fmt.Fprintf(os.Stdout, "Response from `StatusEmbedConfigApi.GetPagesPageIdStatusEmbedConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdStatusEmbedConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusEmbedConfig**](StatusEmbedConfig.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdStatusEmbedConfig

> StatusEmbedConfig PatchPagesPageIdStatusEmbedConfig(ctx, pageId).PatchPagesPageIdStatusEmbedConfig(patchPagesPageIdStatusEmbedConfig).Execute()

Update status embed config settings



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
    patchPagesPageIdStatusEmbedConfig := *openapiclient.NewPatchPagesPageIdStatusEmbedConfig() // PatchPagesPageIdStatusEmbedConfig | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatusEmbedConfigApi.PatchPagesPageIdStatusEmbedConfig(context.Background(), pageId).PatchPagesPageIdStatusEmbedConfig(patchPagesPageIdStatusEmbedConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusEmbedConfigApi.PatchPagesPageIdStatusEmbedConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageIdStatusEmbedConfig`: StatusEmbedConfig
    fmt.Fprintf(os.Stdout, "Response from `StatusEmbedConfigApi.PatchPagesPageIdStatusEmbedConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdStatusEmbedConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchPagesPageIdStatusEmbedConfig** | [**PatchPagesPageIdStatusEmbedConfig**](PatchPagesPageIdStatusEmbedConfig.md) |  | 

### Return type

[**StatusEmbedConfig**](StatusEmbedConfig.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdStatusEmbedConfig

> StatusEmbedConfig PutPagesPageIdStatusEmbedConfig(ctx, pageId).PutPagesPageIdStatusEmbedConfig(putPagesPageIdStatusEmbedConfig).Execute()

Update status embed config settings



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
    putPagesPageIdStatusEmbedConfig := *openapiclient.NewPutPagesPageIdStatusEmbedConfig() // PutPagesPageIdStatusEmbedConfig | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatusEmbedConfigApi.PutPagesPageIdStatusEmbedConfig(context.Background(), pageId).PutPagesPageIdStatusEmbedConfig(putPagesPageIdStatusEmbedConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatusEmbedConfigApi.PutPagesPageIdStatusEmbedConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdStatusEmbedConfig`: StatusEmbedConfig
    fmt.Fprintf(os.Stdout, "Response from `StatusEmbedConfigApi.PutPagesPageIdStatusEmbedConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdStatusEmbedConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putPagesPageIdStatusEmbedConfig** | [**PutPagesPageIdStatusEmbedConfig**](PutPagesPageIdStatusEmbedConfig.md) |  | 

### Return type

[**StatusEmbedConfig**](StatusEmbedConfig.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

