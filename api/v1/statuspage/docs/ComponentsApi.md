# \ComponentsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdComponentsComponentId**](ComponentsApi.md#DeletePagesPageIdComponentsComponentId) | **Delete** /pages/{page_id}/components/{component_id} | Delete a component
[**DeletePagesPageIdComponentsComponentIdPageAccessGroups**](ComponentsApi.md#DeletePagesPageIdComponentsComponentIdPageAccessGroups) | **Delete** /pages/{page_id}/components/{component_id}/page_access_groups | Remove page access groups from a component
[**DeletePagesPageIdComponentsComponentIdPageAccessUsers**](ComponentsApi.md#DeletePagesPageIdComponentsComponentIdPageAccessUsers) | **Delete** /pages/{page_id}/components/{component_id}/page_access_users | Remove page access users from component
[**GetPagesPageIdComponents**](ComponentsApi.md#GetPagesPageIdComponents) | **Get** /pages/{page_id}/components | Get a list of components
[**GetPagesPageIdComponentsComponentId**](ComponentsApi.md#GetPagesPageIdComponentsComponentId) | **Get** /pages/{page_id}/components/{component_id} | Get a component
[**GetPagesPageIdComponentsComponentIdUptime**](ComponentsApi.md#GetPagesPageIdComponentsComponentIdUptime) | **Get** /pages/{page_id}/components/{component_id}/uptime | Get uptime data for a component
[**PatchPagesPageIdComponentsComponentId**](ComponentsApi.md#PatchPagesPageIdComponentsComponentId) | **Patch** /pages/{page_id}/components/{component_id} | Update a component
[**PostPagesPageIdComponents**](ComponentsApi.md#PostPagesPageIdComponents) | **Post** /pages/{page_id}/components | Create a component
[**PostPagesPageIdComponentsComponentIdPageAccessGroups**](ComponentsApi.md#PostPagesPageIdComponentsComponentIdPageAccessGroups) | **Post** /pages/{page_id}/components/{component_id}/page_access_groups | Add page access groups to a component
[**PostPagesPageIdComponentsComponentIdPageAccessUsers**](ComponentsApi.md#PostPagesPageIdComponentsComponentIdPageAccessUsers) | **Post** /pages/{page_id}/components/{component_id}/page_access_users | Add page access users to a component
[**PutPagesPageIdComponentsComponentId**](ComponentsApi.md#PutPagesPageIdComponentsComponentId) | **Put** /pages/{page_id}/components/{component_id} | Update a component


# **DeletePagesPageIdComponentsComponentId**
> DeletePagesPageIdComponentsComponentId(ctx, pageId, componentId)
Delete a component

Delete a component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **componentId** | **string**| Component identifier | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePagesPageIdComponentsComponentIdPageAccessGroups**
> Component DeletePagesPageIdComponentsComponentIdPageAccessGroups(ctx, pageId, componentId)
Remove page access groups from a component

Remove page access groups from a component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **componentId** | **string**| Component identifier | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePagesPageIdComponentsComponentIdPageAccessUsers**
> Component DeletePagesPageIdComponentsComponentIdPageAccessUsers(ctx, pageId, componentId)
Remove page access users from component

Remove page access users from component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **componentId** | **string**| Component identifier | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdComponents**
> []Component GetPagesPageIdComponents(ctx, pageId, optional)
Get a list of components

Get a list of components

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdComponentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page offset to fetch. | 
 **perPage** | **optional.Int32**| Number of results to return per page. | 

### Return type

[**[]Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdComponentsComponentId**
> Component GetPagesPageIdComponentsComponentId(ctx, pageId, componentId)
Get a component

Get a component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **componentId** | **string**| Component identifier | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdComponentsComponentIdUptime**
> ComponentUptime GetPagesPageIdComponentsComponentIdUptime(ctx, pageId, componentId, optional)
Get uptime data for a component

Get uptime data for a component that has uptime showcase enabled

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **componentId** | **string**| Component identifier | 
 **optional** | ***GetPagesPageIdComponentsComponentIdUptimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdComponentsComponentIdUptimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | [**optional.Interface of PartialStartDate**](.md)| The start date for uptime calculation (defaults to the component&#39;s start_date field or 90 days ago, whichever is more recent). The maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year. If the year and month are given, the start date defaults to the first day of that month. The earliest supported date is January 1, 1970.  | 
 **end** | [**optional.Interface of PartialEndDate**](.md)| The end date for uptime calculation (defaults to today in the page&#39;s time zone). The maximum supported date range is six calendar months. If the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month. The earliest supported date is January 1, 1970.  | 

### Return type

[**ComponentUptime**](ComponentUptime.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdComponentsComponentId**
> Component PatchPagesPageIdComponentsComponentId(ctx, pageId, componentId, patchPagesPageIdComponents)
Update a component

Update a component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **componentId** | **string**| Component identifier | 
  **patchPagesPageIdComponents** | [**PatchPagesPageIdComponents**](PatchPagesPageIdComponents.md)|  | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdComponents**
> Component PostPagesPageIdComponents(ctx, pageId, postPagesPageIdComponents)
Create a component

Create a component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdComponents** | [**PostPagesPageIdComponents**](PostPagesPageIdComponents.md)|  | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdComponentsComponentIdPageAccessGroups**
> Component PostPagesPageIdComponentsComponentIdPageAccessGroups(ctx, pageId, componentId)
Add page access groups to a component

Add page access groups to a component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **componentId** | **string**| Component identifier | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdComponentsComponentIdPageAccessUsers**
> Component PostPagesPageIdComponentsComponentIdPageAccessUsers(ctx, pageId, componentId, pageAccessUserIds)
Add page access users to a component

Add page access users to a component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **componentId** | **string**| Component identifier | 
  **pageAccessUserIds** | [**[]string**](string.md)| List of page access users to add to component | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdComponentsComponentId**
> Component PutPagesPageIdComponentsComponentId(ctx, pageId, componentId, putPagesPageIdComponents)
Update a component

Update a component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **componentId** | **string**| Component identifier | 
  **putPagesPageIdComponents** | [**PutPagesPageIdComponents**](PutPagesPageIdComponents.md)|  | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

