# \PageAccessGroupComponentsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Delete** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Delete components for a page access group
[**DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId**](PageAccessGroupComponentsApi.md#DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId) | **Delete** /pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id} | Remove a component from a page access group
[**GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Get** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | List components for a page access group
[**PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Patch** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Add components to page access group
[**PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Post** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Replace components for a page access group
[**PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PageAccessGroupComponentsApi.md#PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents) | **Put** /pages/{page_id}/page_access_groups/{page_access_group_id}/components | Add components to page access group


# **DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents**
> PageAccessGroup DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId, deletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents)
Delete components for a page access group

Delete components for a page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 
  **deletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents** | [**DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md)|  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId**
> PageAccessGroup DeletePagesPageIdPageAccessGroupsPageAccessGroupIdComponentsComponentId(ctx, pageId, pageAccessGroupId, componentId)
Remove a component from a page access group

Remove a component from a page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 
  **componentId** | **string**| Component identifier | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**
> []Component GetPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId)
List components for a page access group

List components for a page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 

### Return type

[**[]Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**
> PageAccessGroup PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId, patchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents)
Add components to page access group

Add components to page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 
  **patchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents** | [**PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PatchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md)|  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**
> PageAccessGroup PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId, postPagesPageIdPageAccessGroupsPageAccessGroupIdComponents)
Replace components for a page access group

Replace components for a page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 
  **postPagesPageIdPageAccessGroupsPageAccessGroupIdComponents** | [**PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PostPagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md)|  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**
> PageAccessGroup PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents(ctx, pageId, pageAccessGroupId, putPagesPageIdPageAccessGroupsPageAccessGroupIdComponents)
Add components to page access group

Add components to page access group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessGroupId** | **string**| Page Access Group Identifier | 
  **putPagesPageIdPageAccessGroupsPageAccessGroupIdComponents** | [**PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents**](PutPagesPageIdPageAccessGroupsPageAccessGroupIdComponents.md)|  | 

### Return type

[**PageAccessGroup**](PageAccessGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

