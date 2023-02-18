# \SubscribersApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdSubscribersSubscriberId**](SubscribersApi.md#DeletePagesPageIdSubscribersSubscriberId) | **Delete** /pages/{page_id}/subscribers/{subscriber_id} | Unsubscribe a subscriber
[**GetPagesPageIdSubscribers**](SubscribersApi.md#GetPagesPageIdSubscribers) | **Get** /pages/{page_id}/subscribers | Get a list of subscribers
[**GetPagesPageIdSubscribersCount**](SubscribersApi.md#GetPagesPageIdSubscribersCount) | **Get** /pages/{page_id}/subscribers/count | Get a count of subscribers by type
[**GetPagesPageIdSubscribersHistogramByState**](SubscribersApi.md#GetPagesPageIdSubscribersHistogramByState) | **Get** /pages/{page_id}/subscribers/histogram_by_state | Get a histogram of subscribers by type and then state
[**GetPagesPageIdSubscribersSubscriberId**](SubscribersApi.md#GetPagesPageIdSubscribersSubscriberId) | **Get** /pages/{page_id}/subscribers/{subscriber_id} | Get a subscriber
[**GetPagesPageIdSubscribersUnsubscribed**](SubscribersApi.md#GetPagesPageIdSubscribersUnsubscribed) | **Get** /pages/{page_id}/subscribers/unsubscribed | Get a list of unsubscribed subscribers
[**PatchPagesPageIdSubscribersSubscriberId**](SubscribersApi.md#PatchPagesPageIdSubscribersSubscriberId) | **Patch** /pages/{page_id}/subscribers/{subscriber_id} | Update a subscriber
[**PostPagesPageIdSubscribers**](SubscribersApi.md#PostPagesPageIdSubscribers) | **Post** /pages/{page_id}/subscribers | Create a subscriber
[**PostPagesPageIdSubscribersReactivate**](SubscribersApi.md#PostPagesPageIdSubscribersReactivate) | **Post** /pages/{page_id}/subscribers/reactivate | Reactivate a list of subscribers
[**PostPagesPageIdSubscribersResendConfirmation**](SubscribersApi.md#PostPagesPageIdSubscribersResendConfirmation) | **Post** /pages/{page_id}/subscribers/resend_confirmation | Resend confirmations to a list of subscribers
[**PostPagesPageIdSubscribersSubscriberIdResendConfirmation**](SubscribersApi.md#PostPagesPageIdSubscribersSubscriberIdResendConfirmation) | **Post** /pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation | Resend confirmation to a subscriber
[**PostPagesPageIdSubscribersUnsubscribe**](SubscribersApi.md#PostPagesPageIdSubscribersUnsubscribe) | **Post** /pages/{page_id}/subscribers/unsubscribe | Unsubscribe a list of subscribers



## DeletePagesPageIdSubscribersSubscriberId

> Subscriber DeletePagesPageIdSubscribersSubscriberId(ctx, pageId, subscriberId).SkipUnsubscriptionNotification(skipUnsubscriptionNotification).Execute()

Unsubscribe a subscriber



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
    subscriberId := "subscriberId_example" // string | Subscriber Identifier
    skipUnsubscriptionNotification := true // bool | If skip_unsubscription_notification is true, the subscriber does not receive any notifications when they are unsubscribed. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.DeletePagesPageIdSubscribersSubscriberId(context.Background(), pageId, subscriberId).SkipUnsubscriptionNotification(skipUnsubscriptionNotification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.DeletePagesPageIdSubscribersSubscriberId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePagesPageIdSubscribersSubscriberId`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.DeletePagesPageIdSubscribersSubscriberId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**subscriberId** | **string** | Subscriber Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdSubscribersSubscriberIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipUnsubscriptionNotification** | **bool** | If skip_unsubscription_notification is true, the subscriber does not receive any notifications when they are unsubscribed. | 

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


## GetPagesPageIdSubscribers

> []Subscriber GetPagesPageIdSubscribers(ctx, pageId).Q(q).Type_(type_).State(state).Limit(limit).Page(page).SortField(sortField).SortDirection(sortDirection).Execute()

Get a list of subscribers



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
    q := "q_example" // string | If this is specified, search the contact information (email, endpoint, or phone number) for the provided value. This parameter doesn’t support searching for Slack subscribers. (optional)
    type_ := "type__example" // string | If specified, only return subscribers of the indicated type. (optional)
    state := "state_example" // string | If this is present, only return subscribers in this state. Specify state \"all\" to find subscribers in any states. (optional) (default to "active")
    limit := int32(56) // int32 | The maximum number of rows to return. If a text query string is specified (q=), the default and maximum limit is 100. If the text query string is not specified, the default and maximum limit are not set, and not providing a limit will return all the subscribers. Beginning February 28, 2023, a default limit of 100 will be imposed and this endpoint will return paginated data (i.e. will no longer return all subscribers) even if this query parameter is not provided. (optional)
    page := int32(56) // int32 | The page offset of subscribers. The first page is page 0, the second page 1, etc. This skips page * limit subscribers. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. (optional) (default to 0)
    sortField := "sortField_example" // string | The field on which to sort: 'primary' to indicate sorting by the identifying field, 'created_at' for sorting by creation timestamp, 'quarantined_at' for sorting by quarantine timestamp, and 'relevance' which sorts by the relevancy of the search text. 'relevance' is not a valid parameter if no search text is supplied. (optional) (default to "primary")
    sortDirection := "sortDirection_example" // string | The sort direction of the listing. (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.GetPagesPageIdSubscribers(context.Background(), pageId).Q(q).Type_(type_).State(state).Limit(limit).Page(page).SortField(sortField).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.GetPagesPageIdSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdSubscribers`: []Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.GetPagesPageIdSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | If this is specified, search the contact information (email, endpoint, or phone number) for the provided value. This parameter doesn’t support searching for Slack subscribers. | 
 **type_** | **string** | If specified, only return subscribers of the indicated type. | 
 **state** | **string** | If this is present, only return subscribers in this state. Specify state \&quot;all\&quot; to find subscribers in any states. | [default to &quot;active&quot;]
 **limit** | **int32** | The maximum number of rows to return. If a text query string is specified (q&#x3D;), the default and maximum limit is 100. If the text query string is not specified, the default and maximum limit are not set, and not providing a limit will return all the subscribers. Beginning February 28, 2023, a default limit of 100 will be imposed and this endpoint will return paginated data (i.e. will no longer return all subscribers) even if this query parameter is not provided. | 
 **page** | **int32** | The page offset of subscribers. The first page is page 0, the second page 1, etc. This skips page * limit subscribers. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. | [default to 0]
 **sortField** | **string** | The field on which to sort: &#39;primary&#39; to indicate sorting by the identifying field, &#39;created_at&#39; for sorting by creation timestamp, &#39;quarantined_at&#39; for sorting by quarantine timestamp, and &#39;relevance&#39; which sorts by the relevancy of the search text. &#39;relevance&#39; is not a valid parameter if no search text is supplied. | [default to &quot;primary&quot;]
 **sortDirection** | **string** | The sort direction of the listing. | [default to &quot;asc&quot;]

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


## GetPagesPageIdSubscribersCount

> SubscriberCountByType GetPagesPageIdSubscribersCount(ctx, pageId).Type_(type_).State(state).Execute()

Get a count of subscribers by type



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
    type_ := "type__example" // string | If this is present, only count subscribers of this type. (optional)
    state := "state_example" // string | If this is present, only count subscribers in this state. Specify state \"all\" to count subscribers in any states. (optional) (default to "active")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.GetPagesPageIdSubscribersCount(context.Background(), pageId).Type_(type_).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.GetPagesPageIdSubscribersCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdSubscribersCount`: SubscriberCountByType
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.GetPagesPageIdSubscribersCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdSubscribersCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | If this is present, only count subscribers of this type. | 
 **state** | **string** | If this is present, only count subscribers in this state. Specify state \&quot;all\&quot; to count subscribers in any states. | [default to &quot;active&quot;]

### Return type

[**SubscriberCountByType**](SubscriberCountByType.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdSubscribersHistogramByState

> SubscriberCountByTypeAndState GetPagesPageIdSubscribersHistogramByState(ctx, pageId).Execute()

Get a histogram of subscribers by type and then state



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.GetPagesPageIdSubscribersHistogramByState(context.Background(), pageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.GetPagesPageIdSubscribersHistogramByState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdSubscribersHistogramByState`: SubscriberCountByTypeAndState
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.GetPagesPageIdSubscribersHistogramByState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdSubscribersHistogramByStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriberCountByTypeAndState**](SubscriberCountByTypeAndState.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdSubscribersSubscriberId

> Subscriber GetPagesPageIdSubscribersSubscriberId(ctx, pageId, subscriberId).Execute()

Get a subscriber



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
    subscriberId := "subscriberId_example" // string | Subscriber Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.GetPagesPageIdSubscribersSubscriberId(context.Background(), pageId, subscriberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.GetPagesPageIdSubscribersSubscriberId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdSubscribersSubscriberId`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.GetPagesPageIdSubscribersSubscriberId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**subscriberId** | **string** | Subscriber Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdSubscribersSubscriberIdRequest struct via the builder pattern


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


## GetPagesPageIdSubscribersUnsubscribed

> []Subscriber GetPagesPageIdSubscribersUnsubscribed(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of unsubscribed subscribers



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
    page := int32(56) // int32 | Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided. (optional)
    perPage := int32(56) // int32 | Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.GetPagesPageIdSubscribersUnsubscribed(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.GetPagesPageIdSubscribersUnsubscribed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdSubscribersUnsubscribed`: []Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.GetPagesPageIdSubscribersUnsubscribed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdSubscribersUnsubscribedRequest struct via the builder pattern


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


## PatchPagesPageIdSubscribersSubscriberId

> Subscriber PatchPagesPageIdSubscribersSubscriberId(ctx, pageId, subscriberId).PatchPagesPageIdSubscribers(patchPagesPageIdSubscribers).Execute()

Update a subscriber



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
    subscriberId := "subscriberId_example" // string | Subscriber Identifier
    patchPagesPageIdSubscribers := *openapiclient.NewPatchPagesPageIdSubscribers() // PatchPagesPageIdSubscribers | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.PatchPagesPageIdSubscribersSubscriberId(context.Background(), pageId, subscriberId).PatchPagesPageIdSubscribers(patchPagesPageIdSubscribers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.PatchPagesPageIdSubscribersSubscriberId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPagesPageIdSubscribersSubscriberId`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.PatchPagesPageIdSubscribersSubscriberId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**subscriberId** | **string** | Subscriber Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdSubscribersSubscriberIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdSubscribers** | [**PatchPagesPageIdSubscribers**](PatchPagesPageIdSubscribers.md) |  | 

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


## PostPagesPageIdSubscribers

> Subscriber PostPagesPageIdSubscribers(ctx, pageId).PostPagesPageIdSubscribers(postPagesPageIdSubscribers).Execute()

Create a subscriber



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
    postPagesPageIdSubscribers := *openapiclient.NewPostPagesPageIdSubscribers() // PostPagesPageIdSubscribers | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.PostPagesPageIdSubscribers(context.Background(), pageId).PostPagesPageIdSubscribers(postPagesPageIdSubscribers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.PostPagesPageIdSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPagesPageIdSubscribers`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscribersApi.PostPagesPageIdSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdSubscribers** | [**PostPagesPageIdSubscribers**](PostPagesPageIdSubscribers.md) |  | 

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


## PostPagesPageIdSubscribersReactivate

> PostPagesPageIdSubscribersReactivate(ctx, pageId).PostPagesPageIdSubscribersReactivate(postPagesPageIdSubscribersReactivate).Execute()

Reactivate a list of subscribers



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
    postPagesPageIdSubscribersReactivate := *openapiclient.NewPostPagesPageIdSubscribersReactivate("Subscribers_example") // PostPagesPageIdSubscribersReactivate | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.PostPagesPageIdSubscribersReactivate(context.Background(), pageId).PostPagesPageIdSubscribersReactivate(postPagesPageIdSubscribersReactivate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.PostPagesPageIdSubscribersReactivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdSubscribersReactivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdSubscribersReactivate** | [**PostPagesPageIdSubscribersReactivate**](PostPagesPageIdSubscribersReactivate.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdSubscribersResendConfirmation

> PostPagesPageIdSubscribersResendConfirmation(ctx, pageId).PostPagesPageIdSubscribersResendConfirmation(postPagesPageIdSubscribersResendConfirmation).Execute()

Resend confirmations to a list of subscribers



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
    postPagesPageIdSubscribersResendConfirmation := *openapiclient.NewPostPagesPageIdSubscribersResendConfirmation("Subscribers_example") // PostPagesPageIdSubscribersResendConfirmation | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.PostPagesPageIdSubscribersResendConfirmation(context.Background(), pageId).PostPagesPageIdSubscribersResendConfirmation(postPagesPageIdSubscribersResendConfirmation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.PostPagesPageIdSubscribersResendConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdSubscribersResendConfirmationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdSubscribersResendConfirmation** | [**PostPagesPageIdSubscribersResendConfirmation**](PostPagesPageIdSubscribersResendConfirmation.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdSubscribersSubscriberIdResendConfirmation

> PostPagesPageIdSubscribersSubscriberIdResendConfirmation(ctx, pageId, subscriberId).Execute()

Resend confirmation to a subscriber



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
    subscriberId := "subscriberId_example" // string | Subscriber Identifier

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.PostPagesPageIdSubscribersSubscriberIdResendConfirmation(context.Background(), pageId, subscriberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.PostPagesPageIdSubscribersSubscriberIdResendConfirmation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**subscriberId** | **string** | Subscriber Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdSubscribersSubscriberIdResendConfirmationRequest struct via the builder pattern


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


## PostPagesPageIdSubscribersUnsubscribe

> PostPagesPageIdSubscribersUnsubscribe(ctx, pageId).PostPagesPageIdSubscribersUnsubscribe(postPagesPageIdSubscribersUnsubscribe).Execute()

Unsubscribe a list of subscribers



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
    postPagesPageIdSubscribersUnsubscribe := *openapiclient.NewPostPagesPageIdSubscribersUnsubscribe("Subscribers_example") // PostPagesPageIdSubscribersUnsubscribe | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscribersApi.PostPagesPageIdSubscribersUnsubscribe(context.Background(), pageId).PostPagesPageIdSubscribersUnsubscribe(postPagesPageIdSubscribersUnsubscribe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscribersApi.PostPagesPageIdSubscribersUnsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdSubscribersUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdSubscribersUnsubscribe** | [**PostPagesPageIdSubscribersUnsubscribe**](PostPagesPageIdSubscribersUnsubscribe.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

