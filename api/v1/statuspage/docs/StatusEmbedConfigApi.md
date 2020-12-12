# \StatusEmbedConfigApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPagesPageIdStatusEmbedConfig**](StatusEmbedConfigApi.md#GetPagesPageIdStatusEmbedConfig) | **Get** /pages/{page_id}/status_embed_config | Get status embed config settings
[**PatchPagesPageIdStatusEmbedConfig**](StatusEmbedConfigApi.md#PatchPagesPageIdStatusEmbedConfig) | **Patch** /pages/{page_id}/status_embed_config | Update status embed config settings
[**PutPagesPageIdStatusEmbedConfig**](StatusEmbedConfigApi.md#PutPagesPageIdStatusEmbedConfig) | **Put** /pages/{page_id}/status_embed_config | Update status embed config settings


# **GetPagesPageIdStatusEmbedConfig**
> StatusEmbedConfig GetPagesPageIdStatusEmbedConfig(ctx, pageId)
Get status embed config settings

Get status embed config settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 

### Return type

[**StatusEmbedConfig**](StatusEmbedConfig.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdStatusEmbedConfig**
> StatusEmbedConfig PatchPagesPageIdStatusEmbedConfig(ctx, pageId, patchPagesPageIdStatusEmbedConfig)
Update status embed config settings

Update status embed config settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **patchPagesPageIdStatusEmbedConfig** | [**PatchPagesPageIdStatusEmbedConfig**](PatchPagesPageIdStatusEmbedConfig.md)|  | 

### Return type

[**StatusEmbedConfig**](StatusEmbedConfig.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdStatusEmbedConfig**
> StatusEmbedConfig PutPagesPageIdStatusEmbedConfig(ctx, pageId, putPagesPageIdStatusEmbedConfig)
Update status embed config settings

Update status embed config settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **putPagesPageIdStatusEmbedConfig** | [**PutPagesPageIdStatusEmbedConfig**](PutPagesPageIdStatusEmbedConfig.md)|  | 

### Return type

[**StatusEmbedConfig**](StatusEmbedConfig.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

