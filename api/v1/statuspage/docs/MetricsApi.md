# \MetricsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdMetricsMetricId**](MetricsApi.md#DeletePagesPageIdMetricsMetricId) | **Delete** /pages/{page_id}/metrics/{metric_id} | Delete a metric
[**DeletePagesPageIdMetricsMetricIdData**](MetricsApi.md#DeletePagesPageIdMetricsMetricIdData) | **Delete** /pages/{page_id}/metrics/{metric_id}/data | Reset data for a metric
[**GetPagesPageIdMetrics**](MetricsApi.md#GetPagesPageIdMetrics) | **Get** /pages/{page_id}/metrics | Get a list of metrics
[**GetPagesPageIdMetricsMetricId**](MetricsApi.md#GetPagesPageIdMetricsMetricId) | **Get** /pages/{page_id}/metrics/{metric_id} | Get a metric
[**GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](MetricsApi.md#GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics) | **Get** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | List metrics for a metric provider
[**PatchPagesPageIdMetricsMetricId**](MetricsApi.md#PatchPagesPageIdMetricsMetricId) | **Patch** /pages/{page_id}/metrics/{metric_id} | Update a metric
[**PostPagesPageIdMetricsData**](MetricsApi.md#PostPagesPageIdMetricsData) | **Post** /pages/{page_id}/metrics/data | Add data points to metrics
[**PostPagesPageIdMetricsMetricIdData**](MetricsApi.md#PostPagesPageIdMetricsMetricIdData) | **Post** /pages/{page_id}/metrics/{metric_id}/data | Add data to a metric
[**PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](MetricsApi.md#PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics) | **Post** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | Create a metric for a metric provider
[**PutPagesPageIdMetricsMetricId**](MetricsApi.md#PutPagesPageIdMetricsMetricId) | **Put** /pages/{page_id}/metrics/{metric_id} | Update a metric


# **DeletePagesPageIdMetricsMetricId**
> Metric DeletePagesPageIdMetricsMetricId(ctx, pageId, metricId)
Delete a metric

Delete a metric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricId** | **string**| Metric Identifier | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePagesPageIdMetricsMetricIdData**
> Metric DeletePagesPageIdMetricsMetricIdData(ctx, pageId, metricId)
Reset data for a metric

Reset data for a metric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricId** | **string**| Metric Identifier | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdMetrics**
> Metric GetPagesPageIdMetrics(ctx, pageId)
Get a list of metrics

Get a list of metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdMetricsMetricId**
> Metric GetPagesPageIdMetricsMetricId(ctx, pageId, metricId)
Get a metric

Get a metric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricId** | **string**| Metric Identifier | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics**
> Metric GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics(ctx, pageId, metricsProviderId)
List metrics for a metric provider

List metrics for a metric provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricsProviderId** | **string**| Metric Provider Identifier | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdMetricsMetricId**
> Metric PatchPagesPageIdMetricsMetricId(ctx, pageId, metricId, patchPagesPageIdMetrics)
Update a metric

Update a metric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricId** | **string**| Metric Identifier | 
  **patchPagesPageIdMetrics** | [**PatchPagesPageIdMetrics**](PatchPagesPageIdMetrics.md)|  | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdMetricsData**
> MetricAddResponse PostPagesPageIdMetricsData(ctx, pageId, postPagesPageIdMetricsData)
Add data points to metrics

Add data points to metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdMetricsData** | [**PostPagesPageIdMetricsData**](PostPagesPageIdMetricsData.md)|  | 

### Return type

[**MetricAddResponse**](MetricAddResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdMetricsMetricIdData**
> SingleMetricAddResponse PostPagesPageIdMetricsMetricIdData(ctx, pageId, metricId, postPagesPageIdMetricsMetricIdData)
Add data to a metric

Add data to a metric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricId** | **string**| Metric Identifier | 
  **postPagesPageIdMetricsMetricIdData** | [**PostPagesPageIdMetricsMetricIdData**](PostPagesPageIdMetricsMetricIdData.md)|  | 

### Return type

[**SingleMetricAddResponse**](SingleMetricAddResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics**
> Metric PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics(ctx, pageId, metricsProviderId, postPagesPageIdMetricsProvidersMetricsProviderIdMetrics)
Create a metric for a metric provider

Create a metric for a metric provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricsProviderId** | **string**| Metric Provider Identifier | 
  **postPagesPageIdMetricsProvidersMetricsProviderIdMetrics** | [**PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics.md)|  | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdMetricsMetricId**
> Metric PutPagesPageIdMetricsMetricId(ctx, pageId, metricId, putPagesPageIdMetrics)
Update a metric

Update a metric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricId** | **string**| Metric Identifier | 
  **putPagesPageIdMetrics** | [**PutPagesPageIdMetrics**](PutPagesPageIdMetrics.md)|  | 

### Return type

[**Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

