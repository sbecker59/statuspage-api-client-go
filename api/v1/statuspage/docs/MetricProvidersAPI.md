# \MetricProvidersAPI

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdMetricsProvidersMetricsProviderId**](MetricProvidersAPI.md#DeletePagesPageIdMetricsProvidersMetricsProviderId) | **Delete** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Delete a metric provider
[**GetPagesPageIdMetricsProviders**](MetricProvidersAPI.md#GetPagesPageIdMetricsProviders) | **Get** /pages/{page_id}/metrics_providers | Get a list of metric providers
[**GetPagesPageIdMetricsProvidersMetricsProviderId**](MetricProvidersAPI.md#GetPagesPageIdMetricsProvidersMetricsProviderId) | **Get** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Get a metric provider
[**PatchPagesPageIdMetricsProvidersMetricsProviderId**](MetricProvidersAPI.md#PatchPagesPageIdMetricsProvidersMetricsProviderId) | **Patch** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Update a metric provider
[**PostPagesPageIdMetricsProviders**](MetricProvidersAPI.md#PostPagesPageIdMetricsProviders) | **Post** /pages/{page_id}/metrics_providers | Create a metric provider
[**PutPagesPageIdMetricsProvidersMetricsProviderId**](MetricProvidersAPI.md#PutPagesPageIdMetricsProvidersMetricsProviderId) | **Put** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Update a metric provider



## DeletePagesPageIdMetricsProvidersMetricsProviderId

> MetricsProvider DeletePagesPageIdMetricsProvidersMetricsProviderId(ctx, pageId, metricsProviderId).Execute()

Delete a metric provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricsProviderId := "metricsProviderId_example" // string | Metric Provider Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricProvidersAPI.DeletePagesPageIdMetricsProvidersMetricsProviderId(context.Background(), pageId, metricsProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricProvidersAPI.DeletePagesPageIdMetricsProvidersMetricsProviderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePagesPageIdMetricsProvidersMetricsProviderId`: MetricsProvider
	fmt.Fprintf(os.Stdout, "Response from `MetricProvidersAPI.DeletePagesPageIdMetricsProvidersMetricsProviderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricsProviderId** | **string** | Metric Provider Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdMetricsProvidersMetricsProviderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdMetricsProviders

> []MetricsProvider GetPagesPageIdMetricsProviders(ctx, pageId).Execute()

Get a list of metric providers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageId := "pageId_example" // string | Page identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricProvidersAPI.GetPagesPageIdMetricsProviders(context.Background(), pageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricProvidersAPI.GetPagesPageIdMetricsProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdMetricsProviders`: []MetricsProvider
	fmt.Fprintf(os.Stdout, "Response from `MetricProvidersAPI.GetPagesPageIdMetricsProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdMetricsProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdMetricsProvidersMetricsProviderId

> MetricsProvider GetPagesPageIdMetricsProvidersMetricsProviderId(ctx, pageId, metricsProviderId).Execute()

Get a metric provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricsProviderId := "metricsProviderId_example" // string | Metric Provider Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricProvidersAPI.GetPagesPageIdMetricsProvidersMetricsProviderId(context.Background(), pageId, metricsProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricProvidersAPI.GetPagesPageIdMetricsProvidersMetricsProviderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdMetricsProvidersMetricsProviderId`: MetricsProvider
	fmt.Fprintf(os.Stdout, "Response from `MetricProvidersAPI.GetPagesPageIdMetricsProvidersMetricsProviderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricsProviderId** | **string** | Metric Provider Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdMetricsProvidersMetricsProviderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdMetricsProvidersMetricsProviderId

> MetricsProvider PatchPagesPageIdMetricsProvidersMetricsProviderId(ctx, pageId, metricsProviderId).PatchPagesPageIdMetricsProviders(patchPagesPageIdMetricsProviders).Execute()

Update a metric provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricsProviderId := "metricsProviderId_example" // string | Metric Provider Identifier
	patchPagesPageIdMetricsProviders := *openapiclient.NewPatchPagesPageIdMetricsProviders() // PatchPagesPageIdMetricsProviders | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricProvidersAPI.PatchPagesPageIdMetricsProvidersMetricsProviderId(context.Background(), pageId, metricsProviderId).PatchPagesPageIdMetricsProviders(patchPagesPageIdMetricsProviders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricProvidersAPI.PatchPagesPageIdMetricsProvidersMetricsProviderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPagesPageIdMetricsProvidersMetricsProviderId`: MetricsProvider
	fmt.Fprintf(os.Stdout, "Response from `MetricProvidersAPI.PatchPagesPageIdMetricsProvidersMetricsProviderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricsProviderId** | **string** | Metric Provider Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdMetricsProvidersMetricsProviderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdMetricsProviders** | [**PatchPagesPageIdMetricsProviders**](PatchPagesPageIdMetricsProviders.md) |  | 

### Return type

[**MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdMetricsProviders

> MetricsProvider PostPagesPageIdMetricsProviders(ctx, pageId).PostPagesPageIdMetricsProviders(postPagesPageIdMetricsProviders).Execute()

Create a metric provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	postPagesPageIdMetricsProviders := *openapiclient.NewPostPagesPageIdMetricsProviders() // PostPagesPageIdMetricsProviders | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricProvidersAPI.PostPagesPageIdMetricsProviders(context.Background(), pageId).PostPagesPageIdMetricsProviders(postPagesPageIdMetricsProviders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricProvidersAPI.PostPagesPageIdMetricsProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdMetricsProviders`: MetricsProvider
	fmt.Fprintf(os.Stdout, "Response from `MetricProvidersAPI.PostPagesPageIdMetricsProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdMetricsProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdMetricsProviders** | [**PostPagesPageIdMetricsProviders**](PostPagesPageIdMetricsProviders.md) |  | 

### Return type

[**MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdMetricsProvidersMetricsProviderId

> MetricsProvider PutPagesPageIdMetricsProvidersMetricsProviderId(ctx, pageId, metricsProviderId).PutPagesPageIdMetricsProviders(putPagesPageIdMetricsProviders).Execute()

Update a metric provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricsProviderId := "metricsProviderId_example" // string | Metric Provider Identifier
	putPagesPageIdMetricsProviders := *openapiclient.NewPutPagesPageIdMetricsProviders() // PutPagesPageIdMetricsProviders | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricProvidersAPI.PutPagesPageIdMetricsProvidersMetricsProviderId(context.Background(), pageId, metricsProviderId).PutPagesPageIdMetricsProviders(putPagesPageIdMetricsProviders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricProvidersAPI.PutPagesPageIdMetricsProvidersMetricsProviderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPagesPageIdMetricsProvidersMetricsProviderId`: MetricsProvider
	fmt.Fprintf(os.Stdout, "Response from `MetricProvidersAPI.PutPagesPageIdMetricsProvidersMetricsProviderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricsProviderId** | **string** | Metric Provider Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdMetricsProvidersMetricsProviderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdMetricsProviders** | [**PutPagesPageIdMetricsProviders**](PutPagesPageIdMetricsProviders.md) |  | 

### Return type

[**MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

