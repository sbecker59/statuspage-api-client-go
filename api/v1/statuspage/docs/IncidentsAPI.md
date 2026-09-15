# \IncidentsAPI

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdIncidentsIncidentId**](IncidentsAPI.md#DeletePagesPageIdIncidentsIncidentId) | **Delete** /pages/{page_id}/incidents/{incident_id} | Delete an incident
[**GetPagesPageIdIncidents**](IncidentsAPI.md#GetPagesPageIdIncidents) | **Get** /pages/{page_id}/incidents | Get a list of incidents
[**GetPagesPageIdIncidentsActiveMaintenance**](IncidentsAPI.md#GetPagesPageIdIncidentsActiveMaintenance) | **Get** /pages/{page_id}/incidents/active_maintenance | Get a list of active maintenances
[**GetPagesPageIdIncidentsIncidentId**](IncidentsAPI.md#GetPagesPageIdIncidentsIncidentId) | **Get** /pages/{page_id}/incidents/{incident_id} | Get an incident
[**GetPagesPageIdIncidentsScheduled**](IncidentsAPI.md#GetPagesPageIdIncidentsScheduled) | **Get** /pages/{page_id}/incidents/scheduled | Get a list of scheduled incidents
[**GetPagesPageIdIncidentsUnresolved**](IncidentsAPI.md#GetPagesPageIdIncidentsUnresolved) | **Get** /pages/{page_id}/incidents/unresolved | Get a list of unresolved incidents
[**GetPagesPageIdIncidentsUpcoming**](IncidentsAPI.md#GetPagesPageIdIncidentsUpcoming) | **Get** /pages/{page_id}/incidents/upcoming | Get a list of upcoming incidents
[**PatchPagesPageIdIncidentsIncidentId**](IncidentsAPI.md#PatchPagesPageIdIncidentsIncidentId) | **Patch** /pages/{page_id}/incidents/{incident_id} | Update an incident
[**PostPagesPageIdIncidents**](IncidentsAPI.md#PostPagesPageIdIncidents) | **Post** /pages/{page_id}/incidents | Create an incident
[**PutPagesPageIdIncidentsIncidentId**](IncidentsAPI.md#PutPagesPageIdIncidentsIncidentId) | **Put** /pages/{page_id}/incidents/{incident_id} | Update an incident



## DeletePagesPageIdIncidentsIncidentId

> Incident DeletePagesPageIdIncidentsIncidentId(ctx, pageId, incidentId).Execute()

Delete an incident



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.DeletePagesPageIdIncidentsIncidentId(context.Background(), pageId, incidentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.DeletePagesPageIdIncidentsIncidentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePagesPageIdIncidentsIncidentId`: Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.DeletePagesPageIdIncidentsIncidentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdIncidentsIncidentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdIncidents

> []Incident GetPagesPageIdIncidents(ctx, pageId).Q(q).Limit(limit).Page(page).Execute()

Get a list of incidents



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
	q := "q_example" // string | If this is specified, search for the text query string in the incidents' name, status, postmortem_body, and incident_updates fields. (optional)
	limit := int32(56) // int32 | The maximum number of rows to return per page. The default and maximum limit is 100. (optional)
	page := int32(56) // int32 | Page offset to fetch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetPagesPageIdIncidents(context.Background(), pageId).Q(q).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetPagesPageIdIncidents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdIncidents`: []Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetPagesPageIdIncidents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | If this is specified, search for the text query string in the incidents&#39; name, status, postmortem_body, and incident_updates fields. | 
 **limit** | **int32** | The maximum number of rows to return per page. The default and maximum limit is 100. | 
 **page** | **int32** | Page offset to fetch. | 

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdIncidentsActiveMaintenance

> []Incident GetPagesPageIdIncidentsActiveMaintenance(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of active maintenances



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
	page := int32(56) // int32 | Page offset to fetch. (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetPagesPageIdIncidentsActiveMaintenance(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetPagesPageIdIncidentsActiveMaintenance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdIncidentsActiveMaintenance`: []Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetPagesPageIdIncidentsActiveMaintenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentsActiveMaintenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. | [default to 1]
 **perPage** | **int32** | Number of results to return per page. | [default to 100]

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdIncidentsIncidentId

> Incident GetPagesPageIdIncidentsIncidentId(ctx, pageId, incidentId).Execute()

Get an incident



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetPagesPageIdIncidentsIncidentId(context.Background(), pageId, incidentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetPagesPageIdIncidentsIncidentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdIncidentsIncidentId`: Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetPagesPageIdIncidentsIncidentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentsIncidentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdIncidentsScheduled

> []Incident GetPagesPageIdIncidentsScheduled(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of scheduled incidents



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
	page := int32(56) // int32 | Page offset to fetch. (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetPagesPageIdIncidentsScheduled(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetPagesPageIdIncidentsScheduled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdIncidentsScheduled`: []Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetPagesPageIdIncidentsScheduled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentsScheduledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. | [default to 1]
 **perPage** | **int32** | Number of results to return per page. | [default to 100]

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdIncidentsUnresolved

> []Incident GetPagesPageIdIncidentsUnresolved(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of unresolved incidents



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
	page := int32(56) // int32 | Page offset to fetch. (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetPagesPageIdIncidentsUnresolved(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetPagesPageIdIncidentsUnresolved``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdIncidentsUnresolved`: []Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetPagesPageIdIncidentsUnresolved`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentsUnresolvedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. | [default to 1]
 **perPage** | **int32** | Number of results to return per page. | [default to 100]

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPagesPageIdIncidentsUpcoming

> []Incident GetPagesPageIdIncidentsUpcoming(ctx, pageId).Page(page).PerPage(perPage).Execute()

Get a list of upcoming incidents



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
	page := int32(56) // int32 | Page offset to fetch. (optional) (default to 1)
	perPage := int32(56) // int32 | Number of results to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.GetPagesPageIdIncidentsUpcoming(context.Background(), pageId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.GetPagesPageIdIncidentsUpcoming``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPagesPageIdIncidentsUpcoming`: []Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.GetPagesPageIdIncidentsUpcoming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentsUpcomingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page offset to fetch. | [default to 1]
 **perPage** | **int32** | Number of results to return per page. | [default to 100]

### Return type

[**[]Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPagesPageIdIncidentsIncidentId

> Incident PatchPagesPageIdIncidentsIncidentId(ctx, pageId, incidentId).PatchPagesPageIdIncidents(patchPagesPageIdIncidents).Execute()

Update an incident



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
	patchPagesPageIdIncidents := *openapiclient.NewPatchPagesPageIdIncidents() // PatchPagesPageIdIncidents | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.PatchPagesPageIdIncidentsIncidentId(context.Background(), pageId, incidentId).PatchPagesPageIdIncidents(patchPagesPageIdIncidents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.PatchPagesPageIdIncidentsIncidentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPagesPageIdIncidentsIncidentId`: Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.PatchPagesPageIdIncidentsIncidentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPagesPageIdIncidentsIncidentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPagesPageIdIncidents** | [**PatchPagesPageIdIncidents**](PatchPagesPageIdIncidents.md) |  | 

### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPagesPageIdIncidents

> Incident PostPagesPageIdIncidents(ctx, pageId).PostPagesPageIdIncidents(postPagesPageIdIncidents).Execute()

Create an incident



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
	postPagesPageIdIncidents := *openapiclient.NewPostPagesPageIdIncidents() // PostPagesPageIdIncidents | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.PostPagesPageIdIncidents(context.Background(), pageId).PostPagesPageIdIncidents(postPagesPageIdIncidents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.PostPagesPageIdIncidents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPagesPageIdIncidents`: Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.PostPagesPageIdIncidents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPagesPageIdIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPagesPageIdIncidents** | [**PostPagesPageIdIncidents**](PostPagesPageIdIncidents.md) |  | 

### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdIncidentsIncidentId

> Incident PutPagesPageIdIncidentsIncidentId(ctx, pageId, incidentId).PutPagesPageIdIncidents(putPagesPageIdIncidents).Execute()

Update an incident



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
	putPagesPageIdIncidents := *openapiclient.NewPutPagesPageIdIncidents() // PutPagesPageIdIncidents | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IncidentsAPI.PutPagesPageIdIncidentsIncidentId(context.Background(), pageId, incidentId).PutPagesPageIdIncidents(putPagesPageIdIncidents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IncidentsAPI.PutPagesPageIdIncidentsIncidentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPagesPageIdIncidentsIncidentId`: Incident
	fmt.Fprintf(os.Stdout, "Response from `IncidentsAPI.PutPagesPageIdIncidentsIncidentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdIncidentsIncidentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdIncidents** | [**PutPagesPageIdIncidents**](PutPagesPageIdIncidents.md) |  | 

### Return type

[**Incident**](Incident.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

