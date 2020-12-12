# \PageAccessUserMetricsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Delete metrics for page access user
[**DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId**](PageAccessUserMetricsApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id} | Delete metric for page access user
[**GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Get** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Get metrics for page access user
[**PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Patch** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Add metrics for page access user
[**PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Post** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Replace metrics for page access user
[**PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PageAccessUserMetricsApi.md#PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics) | **Put** /pages/{page_id}/page_access_users/{page_access_user_id}/metrics | Add metrics for page access user


# **DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics**
> PageAccessUser DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId, deletePagesPageIdPageAccessUsersPageAccessUserIdMetrics)
Delete metrics for page access user

Delete metrics for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
  **deletePagesPageIdPageAccessUsersPageAccessUserIdMetrics** | [**DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics**](DeletePagesPageIdPageAccessUsersPageAccessUserIdMetrics.md)|  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId**
> PageAccessUser DeletePagesPageIdPageAccessUsersPageAccessUserIdMetricsMetricId(ctx, pageId, pageAccessUserId, metricId)
Delete metric for page access user

Delete metric for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
  **metricId** | **string**| Identifier of metric requested | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics**
> []Metric GetPagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId)
Get metrics for page access user

Get metrics for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 

### Return type

[**[]Metric**](Metric.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics**
> PageAccessUser PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId, patchPagesPageIdPageAccessUsersPageAccessUserIdMetrics)
Add metrics for page access user

Add metrics for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
  **patchPagesPageIdPageAccessUsersPageAccessUserIdMetrics** | [**PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PatchPagesPageIdPageAccessUsersPageAccessUserIdMetrics.md)|  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics**
> PageAccessUser PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId, postPagesPageIdPageAccessUsersPageAccessUserIdMetrics)
Replace metrics for page access user

Replace metrics for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
  **postPagesPageIdPageAccessUsersPageAccessUserIdMetrics** | [**PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PostPagesPageIdPageAccessUsersPageAccessUserIdMetrics.md)|  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics**
> PageAccessUser PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics(ctx, pageId, pageAccessUserId, putPagesPageIdPageAccessUsersPageAccessUserIdMetrics)
Add metrics for page access user

Add metrics for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
  **putPagesPageIdPageAccessUsersPageAccessUserIdMetrics** | [**PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics**](PutPagesPageIdPageAccessUsersPageAccessUserIdMetrics.md)|  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

