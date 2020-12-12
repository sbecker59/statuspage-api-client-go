# \ComponentGroupsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdComponentGroupsId**](ComponentGroupsApi.md#DeletePagesPageIdComponentGroupsId) | **Delete** /pages/{page_id}/component-groups/{id} | Delete a component group
[**GetPagesPageIdComponentGroups**](ComponentGroupsApi.md#GetPagesPageIdComponentGroups) | **Get** /pages/{page_id}/component-groups | Get a list of component groups
[**GetPagesPageIdComponentGroupsId**](ComponentGroupsApi.md#GetPagesPageIdComponentGroupsId) | **Get** /pages/{page_id}/component-groups/{id} | Get a component group
[**GetPagesPageIdComponentGroupsIdUptime**](ComponentGroupsApi.md#GetPagesPageIdComponentGroupsIdUptime) | **Get** /pages/{page_id}/component-groups/{id}/uptime | Get uptime data for a component group
[**PatchPagesPageIdComponentGroupsId**](ComponentGroupsApi.md#PatchPagesPageIdComponentGroupsId) | **Patch** /pages/{page_id}/component-groups/{id} | Update a component group
[**PostPagesPageIdComponentGroups**](ComponentGroupsApi.md#PostPagesPageIdComponentGroups) | **Post** /pages/{page_id}/component-groups | Create a component group
[**PutPagesPageIdComponentGroupsId**](ComponentGroupsApi.md#PutPagesPageIdComponentGroupsId) | **Put** /pages/{page_id}/component-groups/{id} | Update a component group


# **DeletePagesPageIdComponentGroupsId**
> GroupComponent DeletePagesPageIdComponentGroupsId(ctx, pageId, id)
Delete a component group

Delete a component group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **id** | **string**| Component group identifier | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdComponentGroups**
> []GroupComponent GetPagesPageIdComponentGroups(ctx, pageId)
Get a list of component groups

Get a list of component groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 

### Return type

[**[]GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdComponentGroupsId**
> GroupComponent GetPagesPageIdComponentGroupsId(ctx, pageId, id)
Get a component group

Get a component group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **id** | **string**| Component group identifier | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdComponentGroupsIdUptime**
> ComponentGroupUptime GetPagesPageIdComponentGroupsIdUptime(ctx, pageId, id, optional)
Get uptime data for a component group

Get uptime data for a component group that has uptime showcase enabled for at least one component.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **id** | **string**| Component group identifier | 
 **optional** | ***GetPagesPageIdComponentGroupsIdUptimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdComponentGroupsIdUptimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | [**optional.Interface of PartialStartDate**](.md)| The start date for uptime calculation (defaults to the date of the component in the group with the earliest start_date, or 90 days ago, whichever is more recent). The maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year. If the year and month are given, the start date defaults to the first day of that month. The earliest supported date is January 1, 1970.  | 
 **end** | [**optional.Interface of PartialEndDate**](.md)| The end date for uptime calculation (defaults to today in the page&#39;s time zone). The maximum supported date range is six calendar months. If the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month. The earliest supported date is January 1, 1970.  | 

### Return type

[**ComponentGroupUptime**](ComponentGroupUptime.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdComponentGroupsId**
> GroupComponent PatchPagesPageIdComponentGroupsId(ctx, pageId, id, patchPagesPageIdComponentGroups)
Update a component group

Update a component group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **id** | **string**| Component group identifier | 
  **patchPagesPageIdComponentGroups** | [**PatchPagesPageIdComponentGroups**](PatchPagesPageIdComponentGroups.md)|  | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdComponentGroups**
> GroupComponent PostPagesPageIdComponentGroups(ctx, pageId, postPagesPageIdComponentGroups)
Create a component group

Create a component group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdComponentGroups** | [**PostPagesPageIdComponentGroups**](PostPagesPageIdComponentGroups.md)|  | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdComponentGroupsId**
> GroupComponent PutPagesPageIdComponentGroupsId(ctx, pageId, id, putPagesPageIdComponentGroups)
Update a component group

Update a component group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **id** | **string**| Component group identifier | 
  **putPagesPageIdComponentGroups** | [**PutPagesPageIdComponentGroups**](PutPagesPageIdComponentGroups.md)|  | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

