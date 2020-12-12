# \PageAccessUserComponentsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Remove components for page access user
[**DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId**](PageAccessUserComponentsApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id} | Remove component for page access user
[**GetPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#GetPagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Get** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Get components for page access user
[**PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Patch** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Add components for page access user
[**PostPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#PostPagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Post** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Replace components for page access user
[**PutPagesPageIdPageAccessUsersPageAccessUserIdComponents**](PageAccessUserComponentsApi.md#PutPagesPageIdPageAccessUsersPageAccessUserIdComponents) | **Put** /pages/{page_id}/page_access_users/{page_access_user_id}/components | Add components for page access user


# **DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents**
> PageAccessUser DeletePagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId, optional)
Remove components for page access user

Remove components for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
 **optional** | ***DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentIds** | [**optional.Interface of []string**](string.md)| List of components codes to remove.  If omitted, all components will be removed. | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId**
> PageAccessUser DeletePagesPageIdPageAccessUsersPageAccessUserIdComponentsComponentId(ctx, pageId, pageAccessUserId, componentId)
Remove component for page access user

Remove component for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
  **componentId** | **string**| Component identifier | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdPageAccessUsersPageAccessUserIdComponents**
> []Component GetPagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId)
Get components for page access user

Get components for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 

### Return type

[**[]Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents**
> PageAccessUser PatchPagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId, optional)
Add components for page access user

Add components for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
 **optional** | ***PatchPagesPageIdPageAccessUsersPageAccessUserIdComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PatchPagesPageIdPageAccessUsersPageAccessUserIdComponentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**optional.Interface of interface{}**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdPageAccessUsersPageAccessUserIdComponents**
> PageAccessUser PostPagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId, optional)
Replace components for page access user

Replace components for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
 **optional** | ***PostPagesPageIdPageAccessUsersPageAccessUserIdComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostPagesPageIdPageAccessUsersPageAccessUserIdComponentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**optional.Interface of interface{}**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdPageAccessUsersPageAccessUserIdComponents**
> PageAccessUser PutPagesPageIdPageAccessUsersPageAccessUserIdComponents(ctx, pageId, pageAccessUserId, optional)
Add components for page access user

Add components for page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 
 **optional** | ***PutPagesPageIdPageAccessUsersPageAccessUserIdComponentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PutPagesPageIdPageAccessUsersPageAccessUserIdComponentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uNKNOWNBASETYPE** | [**optional.Interface of interface{}**](UNKNOWN_BASE_TYPE.md)|  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

