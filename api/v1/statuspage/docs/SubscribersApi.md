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


# **DeletePagesPageIdSubscribersSubscriberId**
> Subscriber DeletePagesPageIdSubscribersSubscriberId(ctx, pageId, subscriberId, optional)
Unsubscribe a subscriber

Unsubscribe a subscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **subscriberId** | **string**| Subscriber Identifier | 
 **optional** | ***DeletePagesPageIdSubscribersSubscriberIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeletePagesPageIdSubscribersSubscriberIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipUnsubscriptionNotification** | **optional.Bool**| If skip_unsubscription_notification is true, the subscriber does not receive any notifications when they are unsubscribed. | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdSubscribers**
> []Subscriber GetPagesPageIdSubscribers(ctx, pageId, optional)
Get a list of subscribers

Get a list of subscribers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdSubscribersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdSubscribersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**| If this is specified, search the contact information (email, endpoint, or phone number) for the provided value. This parameter doesn’t support searching for Slack subscribers. | 
 **type_** | **optional.String**| If specified, only return subscribers of the indicated type. | 
 **state** | **optional.String**| If this is present, only return subscribers in this state. Specify state \&quot;all\&quot; to find subscribers in any states. | [default to active]
 **limit** | **optional.Int32**| The maximum number of rows to return. If a text query string is specified (q&#x3D;), the default and maximum limit is 100. If the text query string is not specified, the default and maximum limit are not set, and not providing a limit will return all the subscribers. | 
 **page** | **optional.Int32**| The page offset of subscribers. The first page is page 0, the second page 1, etc. This skips page * limit subscribers. | [default to 0]
 **sortField** | **optional.String**| The field on which to sort: &#39;primary&#39; to indicate sorting by the identifying field, &#39;created_at&#39; for sorting by creation timestamp, &#39;quarantined_at&#39; for sorting by quarantine timestamp, and &#39;relevance&#39; which sorts by the relevancy of the search text. &#39;relevance&#39; is not a valid parameter if no search text is supplied. | [default to primary]
 **sortDirection** | **optional.String**| The sort direction of the listing. | [default to asc]

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdSubscribersCount**
> SubscriberCountByType GetPagesPageIdSubscribersCount(ctx, pageId, optional)
Get a count of subscribers by type

Get a count of subscribers by type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdSubscribersCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdSubscribersCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| If this is present, only count subscribers of this type. | 
 **state** | **optional.String**| If this is present, only count subscribers in this state. Specify state \&quot;all\&quot; to count subscribers in any states. | [default to active]

### Return type

[**SubscriberCountByType**](SubscriberCountByType.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdSubscribersHistogramByState**
> SubscriberCountByTypeAndState GetPagesPageIdSubscribersHistogramByState(ctx, pageId)
Get a histogram of subscribers by type and then state

Get a histogram of subscribers by type and then state

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 

### Return type

[**SubscriberCountByTypeAndState**](SubscriberCountByTypeAndState.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdSubscribersSubscriberId**
> Subscriber GetPagesPageIdSubscribersSubscriberId(ctx, pageId, subscriberId)
Get a subscriber

Get a subscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **subscriberId** | **string**| Subscriber Identifier | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdSubscribersUnsubscribed**
> []Subscriber GetPagesPageIdSubscribersUnsubscribed(ctx, pageId)
Get a list of unsubscribed subscribers

Get a list of unsubscribed subscribers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdSubscribersSubscriberId**
> Subscriber PatchPagesPageIdSubscribersSubscriberId(ctx, pageId, subscriberId, patchPagesPageIdSubscribers)
Update a subscriber

Update a subscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **subscriberId** | **string**| Subscriber Identifier | 
  **patchPagesPageIdSubscribers** | [**PatchPagesPageIdSubscribers**](PatchPagesPageIdSubscribers.md)|  | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdSubscribers**
> Subscriber PostPagesPageIdSubscribers(ctx, pageId, postPagesPageIdSubscribers)
Create a subscriber

Create a subscriber. Not applicable for Slack subscribers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdSubscribers** | [**PostPagesPageIdSubscribers**](PostPagesPageIdSubscribers.md)|  | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdSubscribersReactivate**
> PostPagesPageIdSubscribersReactivate(ctx, pageId, postPagesPageIdSubscribersReactivate)
Reactivate a list of subscribers

Reactivate a list of quarantined subscribers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdSubscribersReactivate** | [**PostPagesPageIdSubscribersReactivate**](PostPagesPageIdSubscribersReactivate.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdSubscribersResendConfirmation**
> PostPagesPageIdSubscribersResendConfirmation(ctx, pageId, postPagesPageIdSubscribersResendConfirmation)
Resend confirmations to a list of subscribers

Resend confirmations to a list of subscribers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdSubscribersResendConfirmation** | [**PostPagesPageIdSubscribersResendConfirmation**](PostPagesPageIdSubscribersResendConfirmation.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdSubscribersSubscriberIdResendConfirmation**
> PostPagesPageIdSubscribersSubscriberIdResendConfirmation(ctx, pageId, subscriberId)
Resend confirmation to a subscriber

Resend confirmation to a subscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **subscriberId** | **string**| Subscriber Identifier | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdSubscribersUnsubscribe**
> PostPagesPageIdSubscribersUnsubscribe(ctx, pageId, postPagesPageIdSubscribersUnsubscribe)
Unsubscribe a list of subscribers

Unsubscribe a list of subscribers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdSubscribersUnsubscribe** | [**PostPagesPageIdSubscribersUnsubscribe**](PostPagesPageIdSubscribersUnsubscribe.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

