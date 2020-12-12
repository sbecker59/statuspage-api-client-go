# \PageAccessGroupsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessGroupsPageAccessGroupId**](PageAccessGroupsApi.md#DeletePagesPageIdPageAccessGroupsPageAccessGroupId) | **Delete** /pages/{page_id}/page_access_groups/{page_access_group_id} | Remove a page access group
[**GetPagesPageIdPageAccessGroups**](PageAccessGroupsApi.md#GetPagesPageIdPageAccessGroups) | **Get** /pages/{page_id}/page_access_groups | Get a list of page access groups
[**GetPagesPageIdPageAccessGroupsPageAccessGroupId**](PageAccessGroupsApi.md#GetPagesPageIdPageAccessGroupsPageAccessGroupId) | **Get** /pages/{page_id}/page_access_groups/{page_access_group_id} | Get a page access group
[**PatchPagesPageIdPageAccessGroupsPageAccessGroupId**](PageAccessGroupsApi.md#PatchPagesPageIdPageAccessGroupsPageAccessGroupId) | **Patch** /pages/{page_id}/page_access_groups/{page_access_group_id} | Update a page access group
[**PostPagesPageIdPageAccessGroups**](PageAccessGroupsApi.md#PostPagesPageIdPageAccessGroups) | **Post** /pages/{page_id}/page_access_groups | Create a page access group
[**PutPagesPageIdPageAccessGroupsPageAccessGroupId**](PageAccessGroupsApi.md#PutPagesPageIdPageAccessGroupsPageAccessGroupId) | **Put** /pages/{page_id}/page_access_groups/{page_access_group_id} | Update a page access group


# **DeletePagesPageIdPageAccessGroupsPageAccessGroupId**
> PageAccessGroup DeletePagesPageIdPageAccessGroupsPageAccessGroupId(ctx, pageId, pageAccessGroupId)
Remove a page access group

Remove a page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdPageAccessGroups**
> []PageAccessGroup GetPagesPageIdPageAccessGroups(ctx, pageId)
Get a list of page access groups

Get a list of page access groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 

### Return type

[**[]PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdPageAccessGroupsPageAccessGroupId**
> PageAccessGroup GetPagesPageIdPageAccessGroupsPageAccessGroupId(ctx, pageId, pageAccessGroupId)
Get a page access group

Get a page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdPageAccessGroupsPageAccessGroupId**
> PageAccessGroup PatchPagesPageIdPageAccessGroupsPageAccessGroupId(ctx, pageId, pageAccessGroupId, patchPagesPageIdPageAccessGroups)
Update a page access group

Update a page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 
  **patchPagesPageIdPageAccessGroups** | [**PatchPagesPageIdPageAccessGroups**](PatchPagesPageIdPageAccessGroups.md)|  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdPageAccessGroups**
> PageAccessGroup PostPagesPageIdPageAccessGroups(ctx, pageId, postPagesPageIdPageAccessGroups)
Create a page access group

Create a page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdPageAccessGroups** | [**PostPagesPageIdPageAccessGroups**](PostPagesPageIdPageAccessGroups.md)|  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdPageAccessGroupsPageAccessGroupId**
> PageAccessGroup PutPagesPageIdPageAccessGroupsPageAccessGroupId(ctx, pageId, pageAccessGroupId, putPagesPageIdPageAccessGroups)
Update a page access group

Update a page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 
  **putPagesPageIdPageAccessGroups** | [**PutPagesPageIdPageAccessGroups**](PutPagesPageIdPageAccessGroups.md)|  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

