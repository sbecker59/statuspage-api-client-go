# \PagesApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPages**](PagesApi.md#GetPages) | **Get** /pages | Get a list of pages
[**GetPagesPageId**](PagesApi.md#GetPagesPageId) | **Get** /pages/{page_id} | Get a page
[**PatchPagesPageId**](PagesApi.md#PatchPagesPageId) | **Patch** /pages/{page_id} | Update a page
[**PutPagesPageId**](PagesApi.md#PutPagesPageId) | **Put** /pages/{page_id} | Update a page


# **GetPages**
> []Page GetPages(ctx, )
Get a list of pages

Get a list of pages

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Page**](Page.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageId**
> Page GetPagesPageId(ctx, pageId)
Get a page

Get a page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 

### Return type

[**Page**](Page.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageId**
> Page PatchPagesPageId(ctx, pageId, patchPages)
Update a page

Update a page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **patchPages** | [**PatchPages**](PatchPages.md)|  | 

### Return type

[**Page**](Page.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageId**
> Page PutPagesPageId(ctx, pageId, putPages)
Update a page

Update a page

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **putPages** | [**PutPages**](PutPages.md)|  | 

### Return type

[**Page**](Page.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

