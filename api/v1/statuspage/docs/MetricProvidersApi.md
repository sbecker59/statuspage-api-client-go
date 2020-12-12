# \MetricProvidersApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdMetricsProvidersMetricsProviderId**](MetricProvidersApi.md#DeletePagesPageIdMetricsProvidersMetricsProviderId) | **Delete** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Delete a metric provider
[**GetPagesPageIdMetricsProviders**](MetricProvidersApi.md#GetPagesPageIdMetricsProviders) | **Get** /pages/{page_id}/metrics_providers | Get a list of metric providers
[**GetPagesPageIdMetricsProvidersMetricsProviderId**](MetricProvidersApi.md#GetPagesPageIdMetricsProvidersMetricsProviderId) | **Get** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Get a metric provider
[**GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](MetricProvidersApi.md#GetPagesPageIdMetricsProvidersMetricsProviderIdMetrics) | **Get** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | List metrics for a metric provider
[**PatchPagesPageIdMetricsProvidersMetricsProviderId**](MetricProvidersApi.md#PatchPagesPageIdMetricsProvidersMetricsProviderId) | **Patch** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Update a metric provider
[**PostPagesPageIdMetricsProviders**](MetricProvidersApi.md#PostPagesPageIdMetricsProviders) | **Post** /pages/{page_id}/metrics_providers | Create a metric provider
[**PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics**](MetricProvidersApi.md#PostPagesPageIdMetricsProvidersMetricsProviderIdMetrics) | **Post** /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics | Create a metric for a metric provider
[**PutPagesPageIdMetricsProvidersMetricsProviderId**](MetricProvidersApi.md#PutPagesPageIdMetricsProvidersMetricsProviderId) | **Put** /pages/{page_id}/metrics_providers/{metrics_provider_id} | Update a metric provider


# **DeletePagesPageIdMetricsProvidersMetricsProviderId**
> MetricsProvider DeletePagesPageIdMetricsProvidersMetricsProviderId(ctx, pageId, metricsProviderId)
Delete a metric provider

Delete a metric provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricsProviderId** | **string**| Metric Provider Identifier | 

### Return type

[**MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdMetricsProviders**
> []MetricsProvider GetPagesPageIdMetricsProviders(ctx, pageId)
Get a list of metric providers

Get a list of metric providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 

### Return type

[**[]MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdMetricsProvidersMetricsProviderId**
> MetricsProvider GetPagesPageIdMetricsProvidersMetricsProviderId(ctx, pageId, metricsProviderId)
Get a metric provider

Get a metric provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricsProviderId** | **string**| Metric Provider Identifier | 

### Return type

[**MetricsProvider**](MetricsProvider.md)

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

# **PatchPagesPageIdMetricsProvidersMetricsProviderId**
> MetricsProvider PatchPagesPageIdMetricsProvidersMetricsProviderId(ctx, pageId, metricsProviderId, patchPagesPageIdMetricsProviders)
Update a metric provider

Update a metric provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricsProviderId** | **string**| Metric Provider Identifier | 
  **patchPagesPageIdMetricsProviders** | [**PatchPagesPageIdMetricsProviders**](PatchPagesPageIdMetricsProviders.md)|  | 

### Return type

[**MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdMetricsProviders**
> MetricsProvider PostPagesPageIdMetricsProviders(ctx, pageId, postPagesPageIdMetricsProviders)
Create a metric provider

Create a metric provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdMetricsProviders** | [**PostPagesPageIdMetricsProviders**](PostPagesPageIdMetricsProviders.md)|  | 

### Return type

[**MetricsProvider**](MetricsProvider.md)

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

# **PutPagesPageIdMetricsProvidersMetricsProviderId**
> MetricsProvider PutPagesPageIdMetricsProvidersMetricsProviderId(ctx, pageId, metricsProviderId, putPagesPageIdMetricsProviders)
Update a metric provider

Update a metric provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **metricsProviderId** | **string**| Metric Provider Identifier | 
  **putPagesPageIdMetricsProviders** | [**PutPagesPageIdMetricsProviders**](PutPagesPageIdMetricsProviders.md)|  | 

### Return type

[**MetricsProvider**](MetricsProvider.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

