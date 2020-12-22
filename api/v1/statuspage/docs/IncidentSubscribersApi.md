# \IncidentSubscribersApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId**](IncidentSubscribersApi.md#DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId) | **Delete** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id} | Unsubscribe an incident subscriber
[**GetPagesPageIdIncidentsIncidentIdSubscribers**](IncidentSubscribersApi.md#GetPagesPageIdIncidentsIncidentIdSubscribers) | **Get** /pages/{page_id}/incidents/{incident_id}/subscribers | Get a list of incident subscribers
[**GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId**](IncidentSubscribersApi.md#GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId) | **Get** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id} | Get an incident subscriber
[**PostPagesPageIdIncidentsIncidentIdSubscribers**](IncidentSubscribersApi.md#PostPagesPageIdIncidentsIncidentIdSubscribers) | **Post** /pages/{page_id}/incidents/{incident_id}/subscribers | Create an incident subscriber
[**PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation**](IncidentSubscribersApi.md#PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation) | **Post** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation | Resend confirmation to an incident subscriber



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
    openapiclient "./openapi"
)

func main() {
    pageId := "pageId_example" // string | Page identifier
    incidentId := "incidentId_example" // string | Incident Identifier
    subscriberId := "subscriberId_example" // string | Subscriber Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentSubscribersApi.DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId(context.Background(), pageId, incidentId, subscriberId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersApi.DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `IncidentSubscribersApi.DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId`: %v\n", resp)
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

> []Subscriber GetPagesPageIdIncidentsIncidentIdSubscribers(ctx, pageId, incidentId).Execute()

Get a list of incident subscribers



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
    pageId := "pageId_example" // string | Page identifier
    incidentId := "incidentId_example" // string | Incident Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentSubscribersApi.GetPagesPageIdIncidentsIncidentIdSubscribers(context.Background(), pageId, incidentId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersApi.GetPagesPageIdIncidentsIncidentIdSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdIncidentsIncidentIdSubscribers`: []Subscriber
    fmt.Fprintf(os.Stdout, "Response from `IncidentSubscribersApi.GetPagesPageIdIncidentsIncidentIdSubscribers`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    pageId := "pageId_example" // string | Page identifier
    incidentId := "incidentId_example" // string | Incident Identifier
    subscriberId := "subscriberId_example" // string | Subscriber Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentSubscribersApi.GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId(context.Background(), pageId, incidentId, subscriberId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersApi.GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `IncidentSubscribersApi.GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    pageId := "pageId_example" // string | Page identifier
    incidentId := "incidentId_example" // string | Incident Identifier
    postPagesPageIdIncidentsIncidentIdSubscribers := *openapiclient.NewPostPagesPageIdIncidentsIncidentIdSubscribers() // PostPagesPageIdIncidentsIncidentIdSubscribers | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentSubscribersApi.PostPagesPageIdIncidentsIncidentIdSubscribers(context.Background(), pageId, incidentId).PostPagesPageIdIncidentsIncidentIdSubscribers(postPagesPageIdIncidentsIncidentIdSubscribers).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersApi.PostPagesPageIdIncidentsIncidentIdSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdIncidentsIncidentIdSubscribers`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `IncidentSubscribersApi.PostPagesPageIdIncidentsIncidentIdSubscribers`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    pageId := "pageId_example" // string | Page identifier
    incidentId := "incidentId_example" // string | Incident Identifier
    subscriberId := "subscriberId_example" // string | Subscriber Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentSubscribersApi.PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation(context.Background(), pageId, incidentId, subscriberId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentSubscribersApi.PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation``: %v\n", err)
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

