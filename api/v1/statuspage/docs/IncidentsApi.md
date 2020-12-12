# \IncidentsApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdIncidentsIncidentId**](IncidentsApi.md#DeletePagesPageIdIncidentsIncidentId) | **Delete** /pages/{page_id}/incidents/{incident_id} | Delete an incident
[**GetPagesPageIdIncidents**](IncidentsApi.md#GetPagesPageIdIncidents) | **Get** /pages/{page_id}/incidents | Get a list of incidents
[**GetPagesPageIdIncidentsActiveMaintenance**](IncidentsApi.md#GetPagesPageIdIncidentsActiveMaintenance) | **Get** /pages/{page_id}/incidents/active_maintenance | Get a list of active maintenances
[**GetPagesPageIdIncidentsIncidentId**](IncidentsApi.md#GetPagesPageIdIncidentsIncidentId) | **Get** /pages/{page_id}/incidents/{incident_id} | Get an incident
[**GetPagesPageIdIncidentsScheduled**](IncidentsApi.md#GetPagesPageIdIncidentsScheduled) | **Get** /pages/{page_id}/incidents/scheduled | Get a list of scheduled incidents
[**GetPagesPageIdIncidentsUnresolved**](IncidentsApi.md#GetPagesPageIdIncidentsUnresolved) | **Get** /pages/{page_id}/incidents/unresolved | Get a list of unresolved incidents
[**GetPagesPageIdIncidentsUpcoming**](IncidentsApi.md#GetPagesPageIdIncidentsUpcoming) | **Get** /pages/{page_id}/incidents/upcoming | Get a list of upcoming incidents
[**PatchPagesPageIdIncidentsIncidentId**](IncidentsApi.md#PatchPagesPageIdIncidentsIncidentId) | **Patch** /pages/{page_id}/incidents/{incident_id} | Update an incident
[**PostPagesPageIdIncidents**](IncidentsApi.md#PostPagesPageIdIncidents) | **Post** /pages/{page_id}/incidents | Create an incident
[**PutPagesPageIdIncidentsIncidentId**](IncidentsApi.md#PutPagesPageIdIncidentsIncidentId) | **Put** /pages/{page_id}/incidents/{incident_id} | Update an incident


# **DeletePagesPageIdIncidentsIncidentId**
> Incident DeletePagesPageIdIncidentsIncidentId(ctx, pageId, incidentId)
Delete an incident

Delete an incident

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 

### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdIncidents**
> []Incident GetPagesPageIdIncidents(ctx, pageId, optional)
Get a list of incidents

Get a list of incidents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdIncidentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdIncidentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**| If this is specified, search for the text query string in the incidents&#39; name, status, postmortem_body, and incident_updates fields. | 
 **limit** | **optional.Int32**| The maximum number of rows to return per page. The default and maximum limit is 100. | 
 **page** | **optional.Int32**| Page offset to fetch. | 

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdIncidentsActiveMaintenance**
> []Incident GetPagesPageIdIncidentsActiveMaintenance(ctx, pageId, optional)
Get a list of active maintenances

Get a list of active maintenances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdIncidentsActiveMaintenanceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdIncidentsActiveMaintenanceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page offset to fetch. | [default to 1]
 **perPage** | **optional.Int32**| Number of results to return per page. | [default to 100]

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdIncidentsIncidentId**
> Incident GetPagesPageIdIncidentsIncidentId(ctx, pageId, incidentId)
Get an incident

Get an incident

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 

### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdIncidentsScheduled**
> []Incident GetPagesPageIdIncidentsScheduled(ctx, pageId, optional)
Get a list of scheduled incidents

Get a list of scheduled incidents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdIncidentsScheduledOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdIncidentsScheduledOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page offset to fetch. | [default to 1]
 **perPage** | **optional.Int32**| Number of results to return per page. | [default to 100]

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdIncidentsUnresolved**
> []Incident GetPagesPageIdIncidentsUnresolved(ctx, pageId, optional)
Get a list of unresolved incidents

Get a list of unresolved incidents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdIncidentsUnresolvedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdIncidentsUnresolvedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page offset to fetch. | [default to 1]
 **perPage** | **optional.Int32**| Number of results to return per page. | [default to 100]

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPagesPageIdIncidentsUpcoming**
> []Incident GetPagesPageIdIncidentsUpcoming(ctx, pageId, optional)
Get a list of upcoming incidents

Get a list of upcoming incidents

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
 **optional** | ***GetPagesPageIdIncidentsUpcomingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPagesPageIdIncidentsUpcomingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Page offset to fetch. | [default to 1]
 **perPage** | **optional.Int32**| Number of results to return per page. | [default to 100]

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPagesPageIdIncidentsIncidentId**
> Incident PatchPagesPageIdIncidentsIncidentId(ctx, pageId, incidentId, patchPagesPageIdIncidents)
Update an incident

Update an incident

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **patchPagesPageIdIncidents** | [**PatchPagesPageIdIncidents**](PatchPagesPageIdIncidents.md)|  | 

### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostPagesPageIdIncidents**
> Incident PostPagesPageIdIncidents(ctx, pageId, postPagesPageIdIncidents)
Create an incident

Create an incident

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **postPagesPageIdIncidents** | [**PostPagesPageIdIncidents**](PostPagesPageIdIncidents.md)|  | 

### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPagesPageIdIncidentsIncidentId**
> Incident PutPagesPageIdIncidentsIncidentId(ctx, pageId, incidentId, putPagesPageIdIncidents)
Update an incident

Update an incident

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pageId** | **string**| Page identifier | 
  **incidentId** | **string**| Incident Identifier | 
  **putPagesPageIdIncidents** | [**PutPagesPageIdIncidents**](PutPagesPageIdIncidents.md)|  | 

### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

