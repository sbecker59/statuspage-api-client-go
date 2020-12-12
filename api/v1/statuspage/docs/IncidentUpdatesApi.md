# \IncidentUpdatesApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**](IncidentUpdatesApi.md#PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId) | **Patch** /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id} | Update a previous incident update
[**PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**](IncidentUpdatesApi.md#PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId) | **Put** /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id} | Update a previous incident update


# **PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**
> IncidentUpdate PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(ctx, pageId, incidentId, incidentUpdateId, patchPagesPageIdIncidentsIncidentIdIncidentUpdates)
Update a previous incident update

Update a previous incident update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **incidentUpdateId** | **string**| Incident Update Identifier | 
  **patchPagesPageIdIncidentsIncidentIdIncidentUpdates** | [**PatchPagesPageIdIncidentsIncidentIdIncidentUpdates**](PatchPagesPageIdIncidentsIncidentIdIncidentUpdates.md)|  | 

### Return type

[**IncidentUpdate**](IncidentUpdate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**
> IncidentUpdate PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(ctx, pageId, incidentId, incidentUpdateId, putPagesPageIdIncidentsIncidentIdIncidentUpdates)
Update a previous incident update

Update a previous incident update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **incidentUpdateId** | **string**| Incident Update Identifier | 
  **putPagesPageIdIncidentsIncidentIdIncidentUpdates** | [**PutPagesPageIdIncidentsIncidentIdIncidentUpdates**](PutPagesPageIdIncidentsIncidentIdIncidentUpdates.md)|  | 

### Return type

[**IncidentUpdate**](IncidentUpdate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

