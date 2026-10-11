# \IncidentUpdatesAPI

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**](IncidentUpdatesAPI.md#PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId) | **Patch** /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id} | Update a previous incident update
[**PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**](IncidentUpdatesAPI.md#PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId) | **Put** /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id} | Update a previous incident update



## PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId

> IncidentUpdate PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(ctx, pageId, incidentId, incidentUpdateId).PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(patchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId).Execute()

Update a previous incident update



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	incidentId := "incidentId_example" // string | Incident Identifier
	incidentUpdateId := "incidentUpdateId_example" // string | Incident Update Identifier
	patchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId := *openapiclient.NewPatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(*openapiclient.NewPutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateIdIncidentUpdate()) // PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentUpdatesAPI.PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(context.Background(), pageId, incidentId, incidentUpdateId).PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(patchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentUpdatesAPI.PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId`: IncidentUpdate
	fmt.Fprintf(os.Stdout, "Response from `IncidentUpdatesAPI.PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 
**incidentUpdateId** | **string** | Incident Update Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId** | [**PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**](PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId.md) |  | 

### Return type

[**IncidentUpdate**](IncidentUpdate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId

> IncidentUpdate PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(ctx, pageId, incidentId, incidentUpdateId).PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(putPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId).Execute()

Update a previous incident update



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sbecker59/statuspage-api-client-go"
)

func main() {
	pageId := "pageId_example" // string | Page identifier
	incidentId := "incidentId_example" // string | Incident Identifier
	incidentUpdateId := "incidentUpdateId_example" // string | Incident Update Identifier
	putPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId := *openapiclient.NewPutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(*openapiclient.NewPutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateIdIncidentUpdate()) // PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentUpdatesAPI.PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(context.Background(), pageId, incidentId, incidentUpdateId).PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId(putPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentUpdatesAPI.PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId`: IncidentUpdate
	fmt.Fprintf(os.Stdout, "Response from `IncidentUpdatesAPI.PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 
**incidentUpdateId** | **string** | Incident Update Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId** | [**PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId**](PutPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId.md) |  | 

### Return type

[**IncidentUpdate**](IncidentUpdate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

