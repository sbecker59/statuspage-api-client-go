# \UsersApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganizationsOrganizationIdUsersUserId**](UsersApi.md#DeleteOrganizationsOrganizationIdUsersUserId) | **Delete** /organizations/{organization_id}/users/{user_id} | Delete a user
[**GetOrganizationsOrganizationIdUsers**](UsersApi.md#GetOrganizationsOrganizationIdUsers) | **Get** /organizations/{organization_id}/users | Get a list of users
[**PostOrganizationsOrganizationIdUsers**](UsersApi.md#PostOrganizationsOrganizationIdUsers) | **Post** /organizations/{organization_id}/users | Create a user



## DeleteOrganizationsOrganizationIdUsersUserId

> User DeleteOrganizationsOrganizationIdUsersUserId(ctx, organizationId, userId).Execute()

Delete a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization Identifier
    userId := "userId_example" // string | User Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.DeleteOrganizationsOrganizationIdUsersUserId(context.Background(), organizationId, userId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteOrganizationsOrganizationIdUsersUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrganizationsOrganizationIdUsersUserId`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.DeleteOrganizationsOrganizationIdUsersUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization Identifier | 
**userId** | **string** | User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationsOrganizationIdUsersUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](User.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationsOrganizationIdUsers

> []User GetOrganizationsOrganizationIdUsers(ctx, organizationId).Execute()

Get a list of users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetOrganizationsOrganizationIdUsers(context.Background(), organizationId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetOrganizationsOrganizationIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationsOrganizationIdUsers`: []User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetOrganizationsOrganizationIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsOrganizationIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]User**](User.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrganizationsOrganizationIdUsers

> User PostOrganizationsOrganizationIdUsers(ctx, organizationId).PostOrganizationsOrganizationIdUsers(postOrganizationsOrganizationIdUsers).Execute()

Create a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization Identifier
    postOrganizationsOrganizationIdUsers := *openapiclient.NewPostOrganizationsOrganizationIdUsers(*openapiclient.NewPostOrganizationsOrganizationIdUsersUser()) // PostOrganizationsOrganizationIdUsers | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.PostOrganizationsOrganizationIdUsers(context.Background(), organizationId).PostOrganizationsOrganizationIdUsers(postOrganizationsOrganizationIdUsers).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.PostOrganizationsOrganizationIdUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOrganizationsOrganizationIdUsers`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.PostOrganizationsOrganizationIdUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrganizationsOrganizationIdUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postOrganizationsOrganizationIdUsers** | [**PostOrganizationsOrganizationIdUsers**](PostOrganizationsOrganizationIdUsers.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

