# \TemplatesApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPagesPageIdIncidentTemplates**](TemplatesApi.md#GetPagesPageIdIncidentTemplates) | **Get** /pages/{page_id}/incident_templates | Get a list of templates
[**PostPagesPageIdIncidentTemplates**](TemplatesApi.md#PostPagesPageIdIncidentTemplates) | **Post** /pages/{page_id}/incident_templates | Create a template


# **GetPagesPageIdIncidentTemplates**
> []IncidentTemplate GetPagesPageIdIncidentTemplates(ctx, pageId, optional)
Get a list of templates

Get a list of templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdIncidentTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdIncidentTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page offset to fetch. | [default to 1]
 **perPage** | **optional.Int32**| Number of results to return per page. | [default to 100]

### Return type

[**[]IncidentTemplate**](IncidentTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdIncidentTemplates**
> IncidentTemplate PostPagesPageIdIncidentTemplates(ctx, pageId, postPagesPageIdIncidentTemplates)
Create a template

Create a template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdIncidentTemplates** | [**PostPagesPageIdIncidentTemplates**](PostPagesPageIdIncidentTemplates.md)|  | 

### Return type

[**IncidentTemplate**](IncidentTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

