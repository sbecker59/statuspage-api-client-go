# \MetricsAPI

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdMetricsMetricId**](MetricsAPI.md#DeletePagesPageIdMetricsMetricId) | **Delete** /pages/{page_id}/metrics/{metric_id} | Delete a metric
[**DeletePagesPageIdMetricsMetricIdData**](MetricsAPI.md#DeletePagesPageIdMetricsMetricIdData) | **Delete** /pages/{page_id}/metrics/{metric_id}/data | Reset data for a metric
[**GetPagesPageIdMetrics**](MetricsAPI.md#GetPagesPageIdMetrics) | **Get** /pages/{page_id}/metrics | Get a list of metrics
[**GetPagesPageIdMetricsMetricId**](MetricsAPI.md#GetPagesPageIdMetricsMetricId) | **Get** /pages/{page_id}/metrics/{metric_id} | Get a metric
[**GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](MetricsAPI.md#GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics) | **Get** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | List metrics for a metric provider
[**PatchPagesPageIdMetricsMetricId**](MetricsAPI.md#PatchPagesPageIdMetricsMetricId) | **Patch** /pages/{page_id}/metrics/{metric_id} | Update a metric
[**PostPagesPageIdMetricsData**](MetricsAPI.md#PostPagesPageIdMetricsData) | **Post** /pages/{page_id}/metrics/data | Add data points to metrics
[**PostPagesPageIdMetricsMetricIdData**](MetricsAPI.md#PostPagesPageIdMetricsMetricIdData) | **Post** /pages/{page_id}/metrics/{metric_id}/data | Add data to a metric
[**PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](MetricsAPI.md#PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics) | **Post** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | Create a metric for a metric provider
[**PutPagesPageIdMetricsMetricId**](MetricsAPI.md#PutPagesPageIdMetricsMetricId) | **Put** /pages/{page_id}/metrics/{metric_id} | Update a metric



## DeletePagesPageIdMetricsMetricId

> Metric DeletePagesPageIdMetricsMetricId(ctx, pageId, metricId).Execute()

Delete a metric



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricId := "metricId_example" // string | Metric Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.DeletePagesPageIdMetricsMetricId(context.Background(), pageId, metricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.DeletePagesPageIdMetricsMetricId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePagesPageIdMetricsMetricId`: Metric
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.DeletePagesPageIdMetricsMetricId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricId** | **string** | Metric Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdMetricsMetricIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePagesPageIdMetricsMetricIdData

> Metric DeletePagesPageIdMetricsMetricIdData(ctx, pageId, metricId).Execute()

Reset data for a metric



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricId := "metricId_example" // string | Metric Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.DeletePagesPageIdMetricsMetricIdData(context.Background(), pageId, metricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.DeletePagesPageIdMetricsMetricIdData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePagesPageIdMetricsMetricIdData`: Metric
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.DeletePagesPageIdMetricsMetricIdData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricId** | **string** | Metric Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdMetricsMetricIdDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdMetrics

> Metric GetPagesPageIdMetrics(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	page := int32(56) // int32 | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. (optional)
	perPage := int32(56) // int32 | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetPagesPageIdMetrics(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetPagesPageIdMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdMetrics`: Metric
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetPagesPageIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdMetricsMetricId

> Metric GetPagesPageIdMetricsMetricId(ctx, pageId, metricId).Execute()

Get a metric



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricId := "metricId_example" // string | Metric Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetPagesPageIdMetricsMetricId(context.Background(), pageId, metricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetPagesPageIdMetricsMetricId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdMetricsMetricId`: Metric
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetPagesPageIdMetricsMetricId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricId** | **string** | Metric Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdMetricsMetricIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics

> Metric GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics(ctx, pageId, metricsProviderId).Page(page).PerPage(perPage).Execute()

List metrics for a metric provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricsProviderId := "metricsProviderId_example" // string | Metric Provider Identifier
	page := int32(56) // int32 | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. (optional)
	perPage := int32(56) // int32 | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics(context.Background(), pageId, metricsProviderId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics`: Metric
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricsProviderId** | **string** | Metric Provider Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdMetricsProvidersMetricsProviderIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdMetricsMetricId

> Metric PatchPagesPageIdMetricsMetricId(ctx, pageId, metricId).PatchPagesPageIdMetrics(patchPagesPageIdMetrics).Execute()

Update a metric



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricId := "metricId_example" // string | Metric Identifier
	patchPagesPageIdMetrics := *openapiclient.NewPatchPagesPageIdMetrics() // PatchPagesPageIdMetrics | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PatchPagesPageIdMetricsMetricId(context.Background(), pageId, metricId).PatchPagesPageIdMetrics(patchPagesPageIdMetrics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PatchPagesPageIdMetricsMetricId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPagesPageIdMetricsMetricId`: Metric
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PatchPagesPageIdMetricsMetricId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricId** | **string** | Metric Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdMetricsMetricIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdMetrics** | [**PatchPagesPageIdMetrics**](PatchPagesPageIdMetrics.md) |  | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdMetricsData

> MetricAddResponse PostPagesPageIdMetricsData(ctx, pageId).PostPagesPageIdMetricsData(postPagesPageIdMetricsData).Execute()

Add data points to metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	postPagesPageIdMetricsData := *openapiclient.NewPostPagesPageIdMetricsData(*openapiclient.NewMetricAddResponse()) // PostPagesPageIdMetricsData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PostPagesPageIdMetricsData(context.Background(), pageId).PostPagesPageIdMetricsData(postPagesPageIdMetricsData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PostPagesPageIdMetricsData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdMetricsData`: MetricAddResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PostPagesPageIdMetricsData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdMetricsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdMetricsData** | [**PostPagesPageIdMetricsData**](PostPagesPageIdMetricsData.md) |  | 

### Return type

[**MetricAddResponse**](MetricAddResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdMetricsMetricIdData

> SingleMetricAddResponse PostPagesPageIdMetricsMetricIdData(ctx, pageId, metricId).PostPagesPageIdMetricsMetricIdData(postPagesPageIdMetricsMetricIdData).Execute()

Add data to a metric



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricId := "metricId_example" // string | Metric Identifier
	postPagesPageIdMetricsMetricIdData := *openapiclient.NewPostPagesPageIdMetricsMetricIdData(*openapiclient.NewPostPagesPageIdMetricsMetricIdDataData()) // PostPagesPageIdMetricsMetricIdData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PostPagesPageIdMetricsMetricIdData(context.Background(), pageId, metricId).PostPagesPageIdMetricsMetricIdData(postPagesPageIdMetricsMetricIdData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PostPagesPageIdMetricsMetricIdData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdMetricsMetricIdData`: SingleMetricAddResponse
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PostPagesPageIdMetricsMetricIdData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricId** | **string** | Metric Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdMetricsMetricIdDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postPagesPageIdMetricsMetricIdData** | [**PostPagesPageIdMetricsMetricIdData**](PostPagesPageIdMetricsMetricIdData.md) |  | 

### Return type

[**SingleMetricAddResponse**](SingleMetricAddResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics

> Metric PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics(ctx, pageId, metricsProviderId).PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics(postPagesPageIdMetricsProvidersMetricsProviderIdMetrics).Execute()

Create a metric for a metric provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricsProviderId := "metricsProviderId_example" // string | Metric Provider Identifier
	postPagesPageIdMetricsProvidersMetricsProviderIdMetrics := *openapiclient.NewPostPagesPageIdMetricsProvidersMetricsProviderIdMetrics() // PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics(context.Background(), pageId, metricsProviderId).PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics(postPagesPageIdMetricsProvidersMetricsProviderIdMetrics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics`: Metric
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricsProviderId** | **string** | Metric Provider Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdMetricsProvidersMetricsProviderIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postPagesPageIdMetricsProvidersMetricsProviderIdMetrics** | [**PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics.md) |  | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdMetricsMetricId

> Metric PutPagesPageIdMetricsMetricId(ctx, pageId, metricId).PutPagesPageIdMetrics(putPagesPageIdMetrics).Execute()

Update a metric



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	metricId := "metricId_example" // string | Metric Identifier
	putPagesPageIdMetrics := *openapiclient.NewPutPagesPageIdMetrics() // PutPagesPageIdMetrics | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricsAPI.PutPagesPageIdMetricsMetricId(context.Background(), pageId, metricId).PutPagesPageIdMetrics(putPagesPageIdMetrics).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.PutPagesPageIdMetricsMetricId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPagesPageIdMetricsMetricId`: Metric
	fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.PutPagesPageIdMetricsMetricId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**metricId** | **string** | Metric Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdMetricsMetricIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdMetrics** | [**PutPagesPageIdMetrics**](PutPagesPageIdMetrics.md) |  | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

