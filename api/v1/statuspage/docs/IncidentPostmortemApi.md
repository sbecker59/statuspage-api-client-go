# \IncidentPostmortemApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdIncidentsIncidentIdPostmortem**](IncidentPostmortemApi.md#DeletePagesPageIdIncidentsIncidentIdPostmortem) | **Delete** /pages/{page_id}/incidents/{incident_id}/postmortem | Delete Postmortem
[**GetPagesPageIdIncidentsIncidentIdPostmortem**](IncidentPostmortemApi.md#GetPagesPageIdIncidentsIncidentIdPostmortem) | **Get** /pages/{page_id}/incidents/{incident_id}/postmortem | Get Postmortem
[**PutPagesPageIdIncidentsIncidentIdPostmortem**](IncidentPostmortemApi.md#PutPagesPageIdIncidentsIncidentIdPostmortem) | **Put** /pages/{page_id}/incidents/{incident_id}/postmortem | Create Postmortem
[**PutPagesPageIdIncidentsIncidentIdPostmortemPublish**](IncidentPostmortemApi.md#PutPagesPageIdIncidentsIncidentIdPostmortemPublish) | **Put** /pages/{page_id}/incidents/{incident_id}/postmortem/publish | Publish Postmortem
[**PutPagesPageIdIncidentsIncidentIdPostmortemRevert**](IncidentPostmortemApi.md#PutPagesPageIdIncidentsIncidentIdPostmortemRevert) | **Put** /pages/{page_id}/incidents/{incident_id}/postmortem/revert | Revert Postmortem


# **DeletePagesPageIdIncidentsIncidentIdPostmortem**
> DeletePagesPageIdIncidentsIncidentIdPostmortem(ctx, pageId, incidentId)
Delete Postmortem

Delete Postmortem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdIncidentsIncidentIdPostmortem**
> Postmortem GetPagesPageIdIncidentsIncidentIdPostmortem(ctx, pageId, incidentId)
Get Postmortem

Get Postmortem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 

### Return type

[**Postmortem**](Postmortem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdIncidentsIncidentIdPostmortem**
> Postmortem PutPagesPageIdIncidentsIncidentIdPostmortem(ctx, pageId, incidentId, putPagesPageIdIncidentsIncidentIdPostmortem)
Create Postmortem

Create Postmortem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **putPagesPageIdIncidentsIncidentIdPostmortem** | [**PutPagesPageIdIncidentsIncidentIdPostmortem**](PutPagesPageIdIncidentsIncidentIdPostmortem.md)|  | 

### Return type

[**Postmortem**](Postmortem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdIncidentsIncidentIdPostmortemPublish**
> Postmortem PutPagesPageIdIncidentsIncidentIdPostmortemPublish(ctx, pageId, incidentId, putPagesPageIdIncidentsIncidentIdPostmortemPublish)
Publish Postmortem

Publish Postmortem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **putPagesPageIdIncidentsIncidentIdPostmortemPublish** | [**PutPagesPageIdIncidentsIncidentIdPostmortemPublish**](PutPagesPageIdIncidentsIncidentIdPostmortemPublish.md)|  | 

### Return type

[**Postmortem**](Postmortem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdIncidentsIncidentIdPostmortemRevert**
> Postmortem PutPagesPageIdIncidentsIncidentIdPostmortemRevert(ctx, pageId, incidentId)
Revert Postmortem

Revert Postmortem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 

### Return type

[**Postmortem**](Postmortem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

