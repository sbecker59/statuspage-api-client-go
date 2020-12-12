# \UsersApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganizationsOrganizationIdUsersUserId**](UsersApi.md#DeleteOrganizationsOrganizationIdUsersUserId) | **Delete** /organizations/{organization_id}/users/{user_id} | Delete a user
[**GetOrganizationsOrganizationIdUsers**](UsersApi.md#GetOrganizationsOrganizationIdUsers) | **Get** /organizations/{organization_id}/users | Get a list of users
[**PostOrganizationsOrganizationIdUsers**](UsersApi.md#PostOrganizationsOrganizationIdUsers) | **Post** /organizations/{organization_id}/users | Create a user


# **DeleteOrganizationsOrganizationIdUsersUserId**
> User DeleteOrganizationsOrganizationIdUsersUserId(ctx, organizationId, userId)
Delete a user

Delete a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **string**| Organization Identifier | 
  **userId** | **string**| User Identifier | 

### Return type

[**User**](User.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizationsOrganizationIdUsers**
> []User GetOrganizationsOrganizationIdUsers(ctx, organizationId)
Get a list of users

Get a list of users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **string**| Organization Identifier | 

### Return type

[**[]User**](User.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostOrganizationsOrganizationIdUsers**
> User PostOrganizationsOrganizationIdUsers(ctx, organizationId, postOrganizationsOrganizationIdUsers)
Create a user

Create a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | **string**| Organization Identifier | 
  **postOrganizationsOrganizationIdUsers** | [**PostOrganizationsOrganizationIdUsers**](PostOrganizationsOrganizationIdUsers.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

