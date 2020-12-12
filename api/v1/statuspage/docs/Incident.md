# Incident

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Incident Identifier | [optional] 
**Components** | [**[]Component**](Component.md) | Incident components | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The timestamp when the incident was created at. | [optional] 
**Impact** | **string** | The impact of the incident. | [optional] 
**ImpactOverride** | **string** | value to override calculated impact value | [optional] 
**IncidentUpdates** | [**[]IncidentUpdate**](IncidentUpdate.md) | The incident updates for incident. | [optional] 
**Metadata** | [***Object**](Object.md) | Metadata attached to the incident. Top level values must be objects. | [optional] 
**MonitoringAt** | [**time.Time**](time.Time.md) | The timestamp when incident entered monitoring state. | [optional] 
**Name** | **string** | Incident Name | [optional] 
**PageId** | **string** | Incident Page Identifier | [optional] 
**PostmortemBody** | **string** | Body of the Postmortem. | [optional] 
**PostmortemBodyLastUpdatedAt** | [**time.Time**](time.Time.md) | The timestamp when the incident postmortem body was last updated at. | [optional] 
**PostmortemIgnored** | **bool** | Controls whether the incident will have postmortem. | [optional] 
**PostmortemNotifiedSubscribers** | **bool** | Indicates whether subscribers are already notificed about postmortem. | [optional] 
**PostmortemNotifiedTwitter** | **bool** | Controls whether to decide if notify postmortem on twitter. | [optional] 
**PostmortemPublishedAt** | **bool** | The timestamp when the postmortem was published. | [optional] 
**ResolvedAt** | [**time.Time**](time.Time.md) | The timestamp when incident was resolved. | [optional] 
**ScheduledAutoCompleted** | **bool** | Controls whether the incident is scheduled to automatically change to complete. | [optional] 
**ScheduledAutoInProgress** | **bool** | Controls whether the incident is scheduled to automatically change to in progress. | [optional] 
**ScheduledFor** | [**time.Time**](time.Time.md) | The timestamp the incident is scheduled for. | [optional] 
**ScheduledRemindPrior** | **bool** | Controls whether to remind subscribers prior to scheduled incidents. | [optional] 
**ScheduledRemindedAt** | [**time.Time**](time.Time.md) | The timestamp when the scheduled incident reminder was sent at. | [optional] 
**ScheduledUntil** | [**time.Time**](time.Time.md) | The timestamp the incident is scheduled until. | [optional] 
**Shortlink** | **string** | Incident Shortlink. | [optional] 
**Status** | **string** | The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The timestamp when the incident was updated at. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


