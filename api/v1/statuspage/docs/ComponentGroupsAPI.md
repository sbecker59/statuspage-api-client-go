# \ComponentGroupsAPI

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdComponentGroupsId**](ComponentGroupsAPI.md#DeletePagesPageIdComponentGroupsId) | **Delete** /pages/{page_id}/component-groups/{id} | Delete a component group
[**GetPagesPageIdComponentGroups**](ComponentGroupsAPI.md#GetPagesPageIdComponentGroups) | **Get** /pages/{page_id}/component-groups | Get a list of component groups
[**GetPagesPageIdComponentGroupsId**](ComponentGroupsAPI.md#GetPagesPageIdComponentGroupsId) | **Get** /pages/{page_id}/component-groups/{id} | Get a component group
[**GetPagesPageIdComponentGroupsIdUptime**](ComponentGroupsAPI.md#GetPagesPageIdComponentGroupsIdUptime) | **Get** /pages/{page_id}/component-groups/{id}/uptime | Get uptime data for a component group
[**PatchPagesPageIdComponentGroupsId**](ComponentGroupsAPI.md#PatchPagesPageIdComponentGroupsId) | **Patch** /pages/{page_id}/component-groups/{id} | Update a component group
[**PostPagesPageIdComponentGroups**](ComponentGroupsAPI.md#PostPagesPageIdComponentGroups) | **Post** /pages/{page_id}/component-groups | Create a component group
[**PutPagesPageIdComponentGroupsId**](ComponentGroupsAPI.md#PutPagesPageIdComponentGroupsId) | **Put** /pages/{page_id}/component-groups/{id} | Update a component group



## DeletePagesPageIdComponentGroupsId

> GroupComponent DeletePagesPageIdComponentGroupsId(ctx, pageId, id).Execute()

Delete a component group



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
	id := "id_example" // string | Component group identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentGroupsAPI.DeletePagesPageIdComponentGroupsId(context.Background(), pageId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsAPI.DeletePagesPageIdComponentGroupsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePagesPageIdComponentGroupsId`: GroupComponent
	fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsAPI.DeletePagesPageIdComponentGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdComponentGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentGroups

> []GroupComponent GetPagesPageIdComponentGroups(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of component groups



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
	resp, r, err := apiClient.ComponentGroupsAPI.GetPagesPageIdComponentGroups(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsAPI.GetPagesPageIdComponentGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdComponentGroups`: []GroupComponent
	fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsAPI.GetPagesPageIdComponentGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

### Return type

[**[]GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentGroupsId

> GroupComponent GetPagesPageIdComponentGroupsId(ctx, pageId, id).Execute()

Get a component group



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
	id := "id_example" // string | Component group identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentGroupsAPI.GetPagesPageIdComponentGroupsId(context.Background(), pageId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsAPI.GetPagesPageIdComponentGroupsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdComponentGroupsId`: GroupComponent
	fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsAPI.GetPagesPageIdComponentGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdComponentGroupsIdUptime

> ComponentGroupUptime GetPagesPageIdComponentGroupsIdUptime(ctx, pageId, id).SkipRelatedEvents(skipRelatedEvents).Start(start).End(end).Execute()

Get uptime data for a component group



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
	id := "id_example" // string | Component group identifier
	skipRelatedEvents := true // bool | Skips supplying the related events data along with the component uptime data. (optional)
	start := TODO // PartialStartDate | The start date for uptime calculation (defaults to the date of the component in the group with the earliest start_date, or 90 days ago, whichever is more recent). The maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year. If the year and month are given, the start date defaults to the first day of that month. The earliest supported date is January 1, 1970.  (optional)
	end := TODO // PartialEndDate | The end date for uptime calculation (defaults to today in the page's time zone). The maximum supported date range is six calendar months. If the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month. The earliest supported date is January 1, 1970.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentGroupsAPI.GetPagesPageIdComponentGroupsIdUptime(context.Background(), pageId, id).SkipRelatedEvents(skipRelatedEvents).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsAPI.GetPagesPageIdComponentGroupsIdUptime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdComponentGroupsIdUptime`: ComponentGroupUptime
	fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsAPI.GetPagesPageIdComponentGroupsIdUptime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdComponentGroupsIdUptimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipRelatedEvents** | **bool** | Skips supplying the related events data along with the component uptime data. | 
 **start** | [**PartialStartDate**](PartialStartDate.md) | The start date for uptime calculation (defaults to the date of the component in the group with the earliest start_date, or 90 days ago, whichever is more recent). The maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year. If the year and month are given, the start date defaults to the first day of that month. The earliest supported date is January 1, 1970.  | 
 **end** | [**PartialEndDate**](PartialEndDate.md) | The end date for uptime calculation (defaults to today in the page&#39;s time zone). The maximum supported date range is six calendar months. If the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month. The earliest supported date is January 1, 1970.  | 

### Return type

[**ComponentGroupUptime**](ComponentGroupUptime.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdComponentGroupsId

> GroupComponent PatchPagesPageIdComponentGroupsId(ctx, pageId, id).PatchPagesPageIdComponentGroups(patchPagesPageIdComponentGroups).Execute()

Update a component group



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
	id := "id_example" // string | Component group identifier
	patchPagesPageIdComponentGroups := *openapiclient.NewPatchPagesPageIdComponentGroups() // PatchPagesPageIdComponentGroups | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentGroupsAPI.PatchPagesPageIdComponentGroupsId(context.Background(), pageId, id).PatchPagesPageIdComponentGroups(patchPagesPageIdComponentGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsAPI.PatchPagesPageIdComponentGroupsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPagesPageIdComponentGroupsId`: GroupComponent
	fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsAPI.PatchPagesPageIdComponentGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdComponentGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdComponentGroups** | [**PatchPagesPageIdComponentGroups**](PatchPagesPageIdComponentGroups.md) |  | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdComponentGroups

> GroupComponent PostPagesPageIdComponentGroups(ctx, pageId).PostPagesPageIdComponentGroups(postPagesPageIdComponentGroups).Execute()

Create a component group



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
	postPagesPageIdComponentGroups := *openapiclient.NewPostPagesPageIdComponentGroups() // PostPagesPageIdComponentGroups | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentGroupsAPI.PostPagesPageIdComponentGroups(context.Background(), pageId).PostPagesPageIdComponentGroups(postPagesPageIdComponentGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsAPI.PostPagesPageIdComponentGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdComponentGroups`: GroupComponent
	fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsAPI.PostPagesPageIdComponentGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdComponentGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdComponentGroups** | [**PostPagesPageIdComponentGroups**](PostPagesPageIdComponentGroups.md) |  | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdComponentGroupsId

> GroupComponent PutPagesPageIdComponentGroupsId(ctx, pageId, id).PutPagesPageIdComponentGroups(putPagesPageIdComponentGroups).Execute()

Update a component group



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
	id := "id_example" // string | Component group identifier
	putPagesPageIdComponentGroups := *openapiclient.NewPutPagesPageIdComponentGroups() // PutPagesPageIdComponentGroups | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentGroupsAPI.PutPagesPageIdComponentGroupsId(context.Background(), pageId, id).PutPagesPageIdComponentGroups(putPagesPageIdComponentGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentGroupsAPI.PutPagesPageIdComponentGroupsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPagesPageIdComponentGroupsId`: GroupComponent
	fmt.Fprintf(os.Stdout, "Response from `ComponentGroupsAPI.PutPagesPageIdComponentGroupsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**id** | **string** | Component group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdComponentGroupsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdComponentGroups** | [**PutPagesPageIdComponentGroups**](PutPagesPageIdComponentGroups.md) |  | 

### Return type

[**GroupComponent**](GroupComponent.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

