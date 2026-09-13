# \IncidentSubscribersAPI

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId**](IncidentSubscribersAPI.md#DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId) | **Delete** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id} | Unsubscribe an incident subscriber
[**GetPagesPageIdIncidentsIncidentIdSubscribers**](IncidentSubscribersAPI.md#GetPagesPageIdIncidentsIncidentIdSubscribers) | **Get** /pages/{page_id}/incidents/{incident_id}/subscribers | Get a list of incident subscribers
[**GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId**](IncidentSubscribersAPI.md#GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId) | **Get** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id} | Get an incident subscriber
[**PostPagesPageIdIncidentsIncidentIdSubscribers**](IncidentSubscribersAPI.md#PostPagesPageIdIncidentsIncidentIdSubscribers) | **Post** /pages/{page_id}/incidents/{incident_id}/subscribers | Create an incident subscriber
[**PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation**](IncidentSubscribersAPI.md#PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation) | **Post** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation | Resend confirmation to an incident subscriber



## DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId

> Subscriber DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId(ctx, pageId, incidentId, subscriberId).Execute()

Unsubscribe an incident subscriber



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
	incidentId := "incidentId_example" // string | Incident Identifier
	subscriberId := "subscriberId_example" // string | Subscriber Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentSubscribersAPI.DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId(context.Background(), pageId, incidentId, subscriberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersAPI.DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId`: Subscriber
	fmt.Fprintf(os.Stdout, "Response from `IncidentSubscribersAPI.DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 
**subscriberId** | **string** | Subscriber Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdIncidentsIncidentIdSubscribers

> []Subscriber GetPagesPageIdIncidentsIncidentIdSubscribers(ctx, pageId, incidentId).Page(page).PerPage(perPage).Execute()

Get a list of incident subscribers



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
	incidentId := "incidentId_example" // string | Incident Identifier
	page := int32(56) // int32 | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. (optional)
	perPage := int32(56) // int32 | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentSubscribersAPI.GetPagesPageIdIncidentsIncidentIdSubscribers(context.Background(), pageId, incidentId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersAPI.GetPagesPageIdIncidentsIncidentIdSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdIncidentsIncidentIdSubscribers`: []Subscriber
	fmt.Fprintf(os.Stdout, "Response from `IncidentSubscribersAPI.GetPagesPageIdIncidentsIncidentIdSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentsIncidentIdSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | 
 **perPage** | **int32** | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. | 

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId

> Subscriber GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId(ctx, pageId, incidentId, subscriberId).Execute()

Get an incident subscriber



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
	incidentId := "incidentId_example" // string | Incident Identifier
	subscriberId := "subscriberId_example" // string | Subscriber Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentSubscribersAPI.GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId(context.Background(), pageId, incidentId, subscriberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersAPI.GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId`: Subscriber
	fmt.Fprintf(os.Stdout, "Response from `IncidentSubscribersAPI.GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 
**subscriberId** | **string** | Subscriber Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdIncidentsIncidentIdSubscribers

> Subscriber PostPagesPageIdIncidentsIncidentIdSubscribers(ctx, pageId, incidentId).PostPagesPageIdIncidentsIncidentIdSubscribers(postPagesPageIdIncidentsIncidentIdSubscribers).Execute()

Create an incident subscriber



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
	incidentId := "incidentId_example" // string | Incident Identifier
	postPagesPageIdIncidentsIncidentIdSubscribers := *openapiclient.NewPostPagesPageIdIncidentsIncidentIdSubscribers() // PostPagesPageIdIncidentsIncidentIdSubscribers | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentSubscribersAPI.PostPagesPageIdIncidentsIncidentIdSubscribers(context.Background(), pageId, incidentId).PostPagesPageIdIncidentsIncidentIdSubscribers(postPagesPageIdIncidentsIncidentIdSubscribers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersAPI.PostPagesPageIdIncidentsIncidentIdSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdIncidentsIncidentIdSubscribers`: Subscriber
	fmt.Fprintf(os.Stdout, "Response from `IncidentSubscribersAPI.PostPagesPageIdIncidentsIncidentIdSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdIncidentsIncidentIdSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postPagesPageIdIncidentsIncidentIdSubscribers** | [**PostPagesPageIdIncidentsIncidentIdSubscribers**](PostPagesPageIdIncidentsIncidentIdSubscribers.md) |  | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation

> PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation(ctx, pageId, incidentId, subscriberId).Execute()

Resend confirmation to an incident subscriber



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
	incidentId := "incidentId_example" // string | Incident Identifier
	subscriberId := "subscriberId_example" // string | Subscriber Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IncidentSubscribersAPI.PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation(context.Background(), pageId, incidentId, subscriberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersAPI.PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 
**subscriberId** | **string** | Subscriber Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

