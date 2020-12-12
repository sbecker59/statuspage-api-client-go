# PatchPagesPageIdIncidentsIncident

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Incident Name | [optional] 
**Status** | **string** | The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed. | [optional] 
**ImpactOverride** | **string** | value to override calculated impact value | [optional] 
**ScheduledFor** | [**time.Time**](time.Time.md) | The timestamp the incident is scheduled for. | [optional] 
**ScheduledUntil** | [**time.Time**](time.Time.md) | The timestamp the incident is scheduled until. | [optional] 
**ScheduledRemindPrior** | **bool** | Controls whether to remind subscribers prior to scheduled incidents. | [optional] 
**ScheduledAutoInProgress** | **bool** | Controls whether the incident is scheduled to automatically change to in progress. | [optional] 
**ScheduledAutoCompleted** | **bool** | Controls whether the incident is scheduled to automatically change to complete. | [optional] 
**Metadata** | [***interface{}**](interface{}.md) | Attach a json object to the incident. All top-level values in the object must also be objects. | [optional] 
**DeliverNotifications** | **bool** | Deliver notifications to subscribers if this is true. If this is false, create an incident without notifying customers. | [optional] [default to true]
**AutoTransitionDeliverNotificationsAtEnd** | **bool** | Controls whether send notification when scheduled maintenances auto transition to completed. | [optional] 
**AutoTransitionDeliverNotificationsAtStart** | **bool** | Controls whether send notification when scheduled maintenances auto transition to started. | [optional] 
**AutoTransitionToMaintenanceState** | **bool** | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. | [optional] 
**AutoTransitionToOperationalState** | **bool** | Controls whether change components status to operational once scheduled maintenance completes. | [optional] 
**AutoTweetAtBeginning** | **bool** | Controls whether tweet automatically when scheduled maintenance starts. | [optional] 
**AutoTweetOnCompletion** | **bool** | Controls whether tweet automatically when scheduled maintenance completes. | [optional] 
**AutoTweetOnCreation** | **bool** | Controls whether tweet automatically when scheduled maintenance is created. | [optional] 
**AutoTweetOneHourBefore** | **bool** | Controls whether tweet automatically one hour before scheduled maintenance starts. | [optional] 
**BackfillDate** | **string** | TimeStamp when incident was backfilled. | [optional] 
**Backfilled** | **bool** | Controls whether incident is backfilled. If true, components cannot be specified. | [optional] 
**Body** | **string** | The initial message, created as the first incident update. | [optional] 
**Components** | [***PostPagesPageIdIncidentsIncidentComponents**](postPagesPageIdIncidents_incident_components.md) |  | [optional] 
**ComponentIds** | **[]string** | List of component_ids affected by this incident | [optional] 
**ScheduledAutoTransition** | **bool** | Same as :scheduled_auto_transition_in_progress. Controls whether the incident is scheduled to automatically change to in progress. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


