# \ComponentsAPI

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdComponentsComponentId**](ComponentsAPI.md#DeletePagesPageIdComponentsComponentId) | **Delete** /pages/{page_id}/components/{component_id} | Delete a component
[**DeletePagesPageIdComponentsComponentIdPageAccessGroups**](ComponentsAPI.md#DeletePagesPageIdComponentsComponentIdPageAccessGroups) | **Delete** /pages/{page_id}/components/{component_id}/page_access_groups | Remove page access groups from a component
[**DeletePagesPageIdComponentsComponentIdPageAccessUsers**](ComponentsAPI.md#DeletePagesPageIdComponentsComponentIdPageAccessUsers) | **Delete** /pages/{page_id}/components/{component_id}/page_access_users | Remove page access users from component
[**GetPagesPageIdComponents**](ComponentsAPI.md#GetPagesPageIdComponents) | **Get** /pages/{page_id}/components | Get a list of components
[**GetPagesPageIdComponentsComponentId**](ComponentsAPI.md#GetPagesPageIdComponentsComponentId) | **Get** /pages/{page_id}/components/{component_id} | Get a component
[**GetPagesPageIdComponentsComponentIdUptime**](ComponentsAPI.md#GetPagesPageIdComponentsComponentIdUptime) | **Get** /pages/{page_id}/components/{component_id}/uptime | Get uptime data for a component
[**PatchPagesPageIdComponentsComponentId**](ComponentsAPI.md#PatchPagesPageIdComponentsComponentId) | **Patch** /pages/{page_id}/components/{component_id} | Update a component
[**PostPagesPageIdComponents**](ComponentsAPI.md#PostPagesPageIdComponents) | **Post** /pages/{page_id}/components | Create a component
[**PostPagesPageIdComponentsComponentIdPageAccessGroups**](ComponentsAPI.md#PostPagesPageIdComponentsComponentIdPageAccessGroups) | **Post** /pages/{page_id}/components/{component_id}/page_access_groups | Add page access groups to a component
[**PostPagesPageIdComponentsComponentIdPageAccessUsers**](ComponentsAPI.md#PostPagesPageIdComponentsComponentIdPageAccessUsers) | **Post** /pages/{page_id}/components/{component_id}/page_access_users | Add page access users to a component
[**PutPagesPageIdComponentsComponentId**](ComponentsAPI.md#PutPagesPageIdComponentsComponentId) | **Put** /pages/{page_id}/components/{component_id} | Update a component



## DeletePagesPageIdComponentsComponentId

> DeletePagesPageIdComponentsComponentId(ctx, pageId, componentId).Execute()

Delete a component



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
	pageId := "pageId_example" // string | Page identifier
	componentId := "componentId_example" // string | Component identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentsAPI.DeletePagesPageIdComponentsComponentId(context.Background(), pageId, componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.DeletePagesPageIdComponentsComponentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdComponentsComponentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePagesPageIdComponentsComponentIdPageAccessGroups

> Component DeletePagesPageIdComponentsComponentIdPageAccessGroups(ctx, pageId, componentId).Execute()

Remove page access groups from a component



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
	pageId := "pageId_example" // string | Page identifier
	componentId := "componentId_example" // string | Component identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.DeletePagesPageIdComponentsComponentIdPageAccessGroups(context.Background(), pageId, componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.DeletePagesPageIdComponentsComponentIdPageAccessGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePagesPageIdComponentsComponentIdPageAccessGroups`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.DeletePagesPageIdComponentsComponentIdPageAccessGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdComponentsComponentIdPageAccessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePagesPageIdComponentsComponentIdPageAccessUsers

> Component DeletePagesPageIdComponentsComponentIdPageAccessUsers(ctx, pageId, componentId).Execute()

Remove page access users from component



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
	pageId := "pageId_example" // string | Page identifier
	componentId := "componentId_example" // string | Component identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.DeletePagesPageIdComponentsComponentIdPageAccessUsers(context.Background(), pageId, componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.DeletePagesPageIdComponentsComponentIdPageAccessUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePagesPageIdComponentsComponentIdPageAccessUsers`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.DeletePagesPageIdComponentsComponentIdPageAccessUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdComponentsComponentIdPageAccessUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponents

> []Component GetPagesPageIdComponents(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of components



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
	pageId := "pageId_example" // string | Page identifier
	page := int32(56) // int32 | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. (optional)
	perPage := int32(56) // int32 | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetPagesPageIdComponents(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetPagesPageIdComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdComponents`: []Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetPagesPageIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

### Return type

[**[]Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentsComponentId

> Component GetPagesPageIdComponentsComponentId(ctx, pageId, componentId).Execute()

Get a component



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
	pageId := "pageId_example" // string | Page identifier
	componentId := "componentId_example" // string | Component identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetPagesPageIdComponentsComponentId(context.Background(), pageId, componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetPagesPageIdComponentsComponentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdComponentsComponentId`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetPagesPageIdComponentsComponentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentsComponentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentsComponentIdUptime

> ComponentUptime GetPagesPageIdComponentsComponentIdUptime(ctx, pageId, componentId).SkipRelatedEvents(skipRelatedEvents).Start(start).End(end).Execute()

Get uptime data for a component



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	componentId := "componentId_example" // string | Component identifier
	skipRelatedEvents := true // bool | Skips supplying the related events data along with the component uptime data. (optional)
	start := time.Now() // string | The start date for uptime calculation (defaults to the component's start_date field or 90 days ago, whichever is more recent). The maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year. If the year and month are given, the start date defaults to the first day of that month. The earliest supported date is January 1, 1970.  (optional)
	end := time.Now() // string | The end date for uptime calculation (defaults to today in the page's time zone). The maximum supported date range is six calendar months. If the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month. The earliest supported date is January 1, 1970.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetPagesPageIdComponentsComponentIdUptime(context.Background(), pageId, componentId).SkipRelatedEvents(skipRelatedEvents).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetPagesPageIdComponentsComponentIdUptime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdComponentsComponentIdUptime`: ComponentUptime
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetPagesPageIdComponentsComponentIdUptime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentsComponentIdUptimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipRelatedEvents** | **bool** | Skips supplying the related events data along with the component uptime data. | 
 **start** | **string** | The start date for uptime calculation (defaults to the component&#39;s start_date field or 90 days ago, whichever is more recent). The maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year. If the year and month are given, the start date defaults to the first day of that month. The earliest supported date is January 1, 1970.  | 
 **end** | **string** | The end date for uptime calculation (defaults to today in the page&#39;s time zone). The maximum supported date range is six calendar months. If the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month. The earliest supported date is January 1, 1970.  | 

### Return type

[**ComponentUptime**](ComponentUptime.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdComponentsComponentId

> Component PatchPagesPageIdComponentsComponentId(ctx, pageId, componentId).PatchPagesPageIdComponents(patchPagesPageIdComponents).Execute()

Update a component



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
	pageId := "pageId_example" // string | Page identifier
	componentId := "componentId_example" // string | Component identifier
	patchPagesPageIdComponents := *openapiclient.NewPatchPagesPageIdComponents() // PatchPagesPageIdComponents | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.PatchPagesPageIdComponentsComponentId(context.Background(), pageId, componentId).PatchPagesPageIdComponents(patchPagesPageIdComponents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.PatchPagesPageIdComponentsComponentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPagesPageIdComponentsComponentId`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.PatchPagesPageIdComponentsComponentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdComponentsComponentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdComponents** | [**PatchPagesPageIdComponents**](PatchPagesPageIdComponents.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdComponents

> Component PostPagesPageIdComponents(ctx, pageId).PostPagesPageIdComponents(postPagesPageIdComponents).Execute()

Create a component



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
	pageId := "pageId_example" // string | Page identifier
	postPagesPageIdComponents := *openapiclient.NewPostPagesPageIdComponents() // PostPagesPageIdComponents | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.PostPagesPageIdComponents(context.Background(), pageId).PostPagesPageIdComponents(postPagesPageIdComponents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.PostPagesPageIdComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdComponents`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.PostPagesPageIdComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdComponents** | [**PostPagesPageIdComponents**](PostPagesPageIdComponents.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdComponentsComponentIdPageAccessGroups

> Component PostPagesPageIdComponentsComponentIdPageAccessGroups(ctx, pageId, componentId).Execute()

Add page access groups to a component



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
	pageId := "pageId_example" // string | Page identifier
	componentId := "componentId_example" // string | Component identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.PostPagesPageIdComponentsComponentIdPageAccessGroups(context.Background(), pageId, componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.PostPagesPageIdComponentsComponentIdPageAccessGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdComponentsComponentIdPageAccessGroups`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.PostPagesPageIdComponentsComponentIdPageAccessGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdComponentsComponentIdPageAccessGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdComponentsComponentIdPageAccessUsers

> Component PostPagesPageIdComponentsComponentIdPageAccessUsers(ctx, pageId, componentId).PageAccessUserIds(pageAccessUserIds).Execute()

Add page access users to a component



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
	pageId := "pageId_example" // string | Page identifier
	componentId := "componentId_example" // string | Component identifier
	pageAccessUserIds := []string{"Inner_example"} // []string | List of page access users to add to component

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.PostPagesPageIdComponentsComponentIdPageAccessUsers(context.Background(), pageId, componentId).PageAccessUserIds(pageAccessUserIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.PostPagesPageIdComponentsComponentIdPageAccessUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdComponentsComponentIdPageAccessUsers`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.PostPagesPageIdComponentsComponentIdPageAccessUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdComponentsComponentIdPageAccessUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageAccessUserIds** | **[]string** | List of page access users to add to component | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdComponentsComponentId

> Component PutPagesPageIdComponentsComponentId(ctx, pageId, componentId).PutPagesPageIdComponents(putPagesPageIdComponents).Execute()

Update a component



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
	pageId := "pageId_example" // string | Page identifier
	componentId := "componentId_example" // string | Component identifier
	putPagesPageIdComponents := *openapiclient.NewPutPagesPageIdComponents() // PutPagesPageIdComponents | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.PutPagesPageIdComponentsComponentId(context.Background(), pageId, componentId).PutPagesPageIdComponents(putPagesPageIdComponents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.PutPagesPageIdComponentsComponentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPagesPageIdComponentsComponentId`: Component
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.PutPagesPageIdComponentsComponentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**componentId** | **string** | Component identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdComponentsComponentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdComponents** | [**PutPagesPageIdComponents**](PutPagesPageIdComponents.md) |  | 

### Return type

[**Component**](Component.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

