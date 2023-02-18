# \IncidentPostmortemApi

All URIs are relative to *https://api.statuspage.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePagesPageIdIncidentsIncidentIdPostmortem**](IncidentPostmortemApi.md#DeletePagesPageIdIncidentsIncidentIdPostmortem) | **Delete** /pages/{page_id}/incidents/{incident_id}/postmortem | Delete Postmortem
[**GetPagesPageIdIncidentsIncidentIdPostmortem**](IncidentPostmortemApi.md#GetPagesPageIdIncidentsIncidentIdPostmortem) | **Get** /pages/{page_id}/incidents/{incident_id}/postmortem | Get Postmortem
[**PutPagesPageIdIncidentsIncidentIdPostmortem**](IncidentPostmortemApi.md#PutPagesPageIdIncidentsIncidentIdPostmortem) | **Put** /pages/{page_id}/incidents/{incident_id}/postmortem | Create Postmortem
[**PutPagesPageIdIncidentsIncidentIdPostmortemPublish**](IncidentPostmortemApi.md#PutPagesPageIdIncidentsIncidentIdPostmortemPublish) | **Put** /pages/{page_id}/incidents/{incident_id}/postmortem/publish | Publish Postmortem
[**PutPagesPageIdIncidentsIncidentIdPostmortemRevert**](IncidentPostmortemApi.md#PutPagesPageIdIncidentsIncidentIdPostmortemRevert) | **Put** /pages/{page_id}/incidents/{incident_id}/postmortem/revert | Revert Postmortem



## DeletePagesPageIdIncidentsIncidentIdPostmortem

> DeletePagesPageIdIncidentsIncidentIdPostmortem(ctx, pageId, incidentId).Execute()

Delete Postmortem



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
    resp, r, err := api_client.IncidentPostmortemApi.DeletePagesPageIdIncidentsIncidentIdPostmortem(context.Background(), pageId, incidentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentPostmortemApi.DeletePagesPageIdIncidentsIncidentIdPostmortem``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagesPageIdIncidentsIncidentIdPostmortemRequest struct via the builder pattern


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


## GetPagesPageIdIncidentsIncidentIdPostmortem

> Postmortem GetPagesPageIdIncidentsIncidentIdPostmortem(ctx, pageId, incidentId).Execute()

Get Postmortem



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
    resp, r, err := api_client.IncidentPostmortemApi.GetPagesPageIdIncidentsIncidentIdPostmortem(context.Background(), pageId, incidentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentPostmortemApi.GetPagesPageIdIncidentsIncidentIdPostmortem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPagesPageIdIncidentsIncidentIdPostmortem`: Postmortem
    fmt.Fprintf(os.Stdout, "Response from `IncidentPostmortemApi.GetPagesPageIdIncidentsIncidentIdPostmortem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagesPageIdIncidentsIncidentIdPostmortemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Postmortem**](Postmortem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdIncidentsIncidentIdPostmortem

> Postmortem PutPagesPageIdIncidentsIncidentIdPostmortem(ctx, pageId, incidentId).PutPagesPageIdIncidentsIncidentIdPostmortem(putPagesPageIdIncidentsIncidentIdPostmortem).Execute()

Create Postmortem



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
    putPagesPageIdIncidentsIncidentIdPostmortem := *openapiclient.NewPutPagesPageIdIncidentsIncidentIdPostmortem() // PutPagesPageIdIncidentsIncidentIdPostmortem | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentPostmortemApi.PutPagesPageIdIncidentsIncidentIdPostmortem(context.Background(), pageId, incidentId).PutPagesPageIdIncidentsIncidentIdPostmortem(putPagesPageIdIncidentsIncidentIdPostmortem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentPostmortemApi.PutPagesPageIdIncidentsIncidentIdPostmortem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdIncidentsIncidentIdPostmortem`: Postmortem
    fmt.Fprintf(os.Stdout, "Response from `IncidentPostmortemApi.PutPagesPageIdIncidentsIncidentIdPostmortem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdIncidentsIncidentIdPostmortemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdIncidentsIncidentIdPostmortem** | [**PutPagesPageIdIncidentsIncidentIdPostmortem**](PutPagesPageIdIncidentsIncidentIdPostmortem.md) |  | 

### Return type

[**Postmortem**](Postmortem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdIncidentsIncidentIdPostmortemPublish

> Postmortem PutPagesPageIdIncidentsIncidentIdPostmortemPublish(ctx, pageId, incidentId).PutPagesPageIdIncidentsIncidentIdPostmortemPublish(putPagesPageIdIncidentsIncidentIdPostmortemPublish).Execute()

Publish Postmortem



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
    putPagesPageIdIncidentsIncidentIdPostmortemPublish := *openapiclient.NewPutPagesPageIdIncidentsIncidentIdPostmortemPublish() // PutPagesPageIdIncidentsIncidentIdPostmortemPublish | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IncidentPostmortemApi.PutPagesPageIdIncidentsIncidentIdPostmortemPublish(context.Background(), pageId, incidentId).PutPagesPageIdIncidentsIncidentIdPostmortemPublish(putPagesPageIdIncidentsIncidentIdPostmortemPublish).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentPostmortemApi.PutPagesPageIdIncidentsIncidentIdPostmortemPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdIncidentsIncidentIdPostmortemPublish`: Postmortem
    fmt.Fprintf(os.Stdout, "Response from `IncidentPostmortemApi.PutPagesPageIdIncidentsIncidentIdPostmortemPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdIncidentsIncidentIdPostmortemPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPagesPageIdIncidentsIncidentIdPostmortemPublish** | [**PutPagesPageIdIncidentsIncidentIdPostmortemPublish**](PutPagesPageIdIncidentsIncidentIdPostmortemPublish.md) |  | 

### Return type

[**Postmortem**](Postmortem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPagesPageIdIncidentsIncidentIdPostmortemRevert

> Postmortem PutPagesPageIdIncidentsIncidentIdPostmortemRevert(ctx, pageId, incidentId).Execute()

Revert Postmortem



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
    resp, r, err := api_client.IncidentPostmortemApi.PutPagesPageIdIncidentsIncidentIdPostmortemRevert(context.Background(), pageId, incidentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IncidentPostmortemApi.PutPagesPageIdIncidentsIncidentIdPostmortemRevert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPagesPageIdIncidentsIncidentIdPostmortemRevert`: Postmortem
    fmt.Fprintf(os.Stdout, "Response from `IncidentPostmortemApi.PutPagesPageIdIncidentsIncidentIdPostmortemRevert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | Page identifier | 
**incidentId** | **string** | Incident Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPagesPageIdIncidentsIncidentIdPostmortemRevertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Postmortem**](Postmortem.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

