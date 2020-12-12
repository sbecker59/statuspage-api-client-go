# \PageAccessUsersApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdPageAccessUsersPageAccessUserId**](PageAccessUsersApi.md#DeletePagesPageIdPageAccessUsersPageAccessUserId) | **Delete** /pages/{page_id}/page_access_users/{page_access_user_id} | Delete page access user
[**GetPagesPageIdPageAccessUsers**](PageAccessUsersApi.md#GetPagesPageIdPageAccessUsers) | **Get** /pages/{page_id}/page_access_users | Get a list of page access users
[**GetPagesPageIdPageAccessUsersPageAccessUserId**](PageAccessUsersApi.md#GetPagesPageIdPageAccessUsersPageAccessUserId) | **Get** /pages/{page_id}/page_access_users/{page_access_user_id} | Get page access user
[**PatchPagesPageIdPageAccessUsersPageAccessUserId**](PageAccessUsersApi.md#PatchPagesPageIdPageAccessUsersPageAccessUserId) | **Patch** /pages/{page_id}/page_access_users/{page_access_user_id} | Update page access user
[**PostPagesPageIdPageAccessUsers**](PageAccessUsersApi.md#PostPagesPageIdPageAccessUsers) | **Post** /pages/{page_id}/page_access_users | Add a page access user
[**PutPagesPageIdPageAccessUsersPageAccessUserId**](PageAccessUsersApi.md#PutPagesPageIdPageAccessUsersPageAccessUserId) | **Put** /pages/{page_id}/page_access_users/{page_access_user_id} | Update page access user


# **DeletePagesPageIdPageAccessUsersPageAccessUserId**
> DeletePagesPageIdPageAccessUsersPageAccessUserId(ctx, pageId, pageAccessUserId)
Delete page access user

Delete page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdPageAccessUsers**
> []PageAccessUser GetPagesPageIdPageAccessUsers(ctx, pageId, optional)
Get a list of page access users

Get a list of page access users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdPageAccessUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdPageAccessUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **optional.String**| Email address to search for | 

### Return type

[**[]PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdPageAccessUsersPageAccessUserId**
> PageAccessUser GetPagesPageIdPageAccessUsersPageAccessUserId(ctx, pageId, pageAccessUserId)
Get page access user

Get page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdPageAccessUsersPageAccessUserId**
> PageAccessUser PatchPagesPageIdPageAccessUsersPageAccessUserId(ctx, pageId, pageAccessUserId)
Update page access user

Update page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdPageAccessUsers**
> PageAccessUser PostPagesPageIdPageAccessUsers(ctx, pageId, postPagesPageIdPageAccessUsers)
Add a page access user

Add a page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdPageAccessUsers** | [**PostPagesPageIdPageAccessUsers**](PostPagesPageIdPageAccessUsers.md)|  | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdPageAccessUsersPageAccessUserId**
> PageAccessUser PutPagesPageIdPageAccessUsersPageAccessUserId(ctx, pageId, pageAccessUserId)
Update page access user

Update page access user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **pageAccessUserId** | **string**| Page Access User Identifier | 

### Return type

[**PageAccessUser**](PageAccessUser.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

