# \PermissionsAPI

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationsOrganizationIdPermissionsUserId**](PermissionsAPI.md#GetOrganizationsOrganizationIdPermissionsUserId) | **Get** /organizations/{organization_id}/permissions/{user_id} | Get a user&#39;s permissions
[**PutOrganizationsOrganizationIdPermissionsUserId**](PermissionsAPI.md#PutOrganizationsOrganizationIdPermissionsUserId) | **Put** /organizations/{organization_id}/permissions/{user_id} | Update a user&#39;s role permissions



## GetOrganizationsOrganizationIdPermissionsUserId

> Permissions GetOrganizationsOrganizationIdPermissionsUserId(ctx, organizationId, userId).Execute()

Get a user's permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization Identifier
	userId := "userId_example" // string | User identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.GetOrganizationsOrganizationIdPermissionsUserId(context.Background(), organizationId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetOrganizationsOrganizationIdPermissionsUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationsOrganizationIdPermissionsUserId`: Permissions
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetOrganizationsOrganizationIdPermissionsUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization Identifier | 
**userId** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsOrganizationIdPermissionsUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Permissions**](Permissions.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrganizationsOrganizationIdPermissionsUserId

> Permissions PutOrganizationsOrganizationIdPermissionsUserId(ctx, organizationId, userId).PutOrganizationsOrganizationIdPermissions(putOrganizationsOrganizationIdPermissions).Execute()

Update a user's role permissions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization Identifier
	userId := "userId_example" // string | User identifier
	putOrganizationsOrganizationIdPermissions := *openapiclient.NewPutOrganizationsOrganizationIdPermissions() // PutOrganizationsOrganizationIdPermissions | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.PutOrganizationsOrganizationIdPermissionsUserId(context.Background(), organizationId, userId).PutOrganizationsOrganizationIdPermissions(putOrganizationsOrganizationIdPermissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.PutOrganizationsOrganizationIdPermissionsUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutOrganizationsOrganizationIdPermissionsUserId`: Permissions
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.PutOrganizationsOrganizationIdPermissionsUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization Identifier | 
**userId** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrganizationsOrganizationIdPermissionsUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putOrganizationsOrganizationIdPermissions** | [**PutOrganizationsOrganizationIdPermissions**](PutOrganizationsOrganizationIdPermissions.md) |  | 

### Return type

[**Permissions**](Permissions.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

