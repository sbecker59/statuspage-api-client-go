# \IncidentSubscribersApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId**](IncidentSubscribersApi.md#DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId) | **Delete** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id} | Unsubscribe an incident subscriber
[**GetPagesPageIdIncidentsIncidentIdSubscribers**](IncidentSubscribersApi.md#GetPagesPageIdIncidentsIncidentIdSubscribers) | **Get** /pages/{page_id}/incidents/{incident_id}/subscribers | Get a list of incident subscribers
[**GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId**](IncidentSubscribersApi.md#GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId) | **Get** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id} | Get an incident subscriber
[**PostPagesPageIdIncidentsIncidentIdSubscribers**](IncidentSubscribersApi.md#PostPagesPageIdIncidentsIncidentIdSubscribers) | **Post** /pages/{page_id}/incidents/{incident_id}/subscribers | Create an incident subscriber
[**PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation**](IncidentSubscribersApi.md#PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation) | **Post** /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation | Resend confirmation to an incident subscriber


# **DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId**
> Subscriber DeletePagesPageIdIncidentsIncidentIdSubscribersSubscriberId(ctx, pageId, incidentId, subscriberId)
Unsubscribe an incident subscriber

Unsubscribe an incident subscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **subscriberId** | **string**| Subscriber Identifier | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdIncidentsIncidentIdSubscribers**
> []Subscriber GetPagesPageIdIncidentsIncidentIdSubscribers(ctx, pageId, incidentId)
Get a list of incident subscribers

Get a list of incident subscribers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId**
> Subscriber GetPagesPageIdIncidentsIncidentIdSubscribersSubscriberId(ctx, pageId, incidentId, subscriberId)
Get an incident subscriber

Get an incident subscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **subscriberId** | **string**| Subscriber Identifier | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdIncidentsIncidentIdSubscribers**
> Subscriber PostPagesPageIdIncidentsIncidentIdSubscribers(ctx, pageId, incidentId, postPagesPageIdIncidentsIncidentIdSubscribers)
Create an incident subscriber

Create an incident subscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **postPagesPageIdIncidentsIncidentIdSubscribers** | [**PostPagesPageIdIncidentsIncidentIdSubscribers**](PostPagesPageIdIncidentsIncidentIdSubscribers.md)|  | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation**
> PostPagesPageIdIncidentsIncidentIdSubscribersSubscriberIdResendConfirmation(ctx, pageId, incidentId, subscriberId)
Resend confirmation to an incident subscriber

Resend confirmation to an incident subscriber

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **subscriberId** | **string**| Subscriber Identifier | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

