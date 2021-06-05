# PatchPagesPageIdIncidentsIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Incident Name | [optional] 
**Status** | Pointer to **string** | The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed. | [optional] 
**ImpactOverride** | Pointer to **string** | value to override calculated impact value | [optional] 
**ScheduledFor** | Pointer to **time.Time** | The timestamp the incident is scheduled for. | [optional] 
**ScheduledUntil** | Pointer to **time.Time** | The timestamp the incident is scheduled until. | [optional] 
**ScheduledRemindPrior** | Pointer to **bool** | Controls whether to remind subscribers prior to scheduled incidents. | [optional] 
**ScheduledAutoInProgress** | Pointer to **bool** | Controls whether the incident is scheduled to automatically change to in progress. | [optional] 
**ScheduledAutoCompleted** | Pointer to **bool** | Controls whether the incident is scheduled to automatically change to complete. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Attach a json object to the incident. All top-level values in the object must also be objects. | [optional] 
**DeliverNotifications** | Pointer to **bool** | Deliver notifications to subscribers if this is true. If this is false, create an incident without notifying customers. | [optional] [default to true]
**AutoTransitionDeliverNotificationsAtEnd** | Pointer to **bool** | Controls whether send notification when scheduled maintenances auto transition to completed. | [optional] 
**AutoTransitionDeliverNotificationsAtStart** | Pointer to **bool** | Controls whether send notification when scheduled maintenances auto transition to started. | [optional] 
**AutoTransitionToMaintenanceState** | Pointer to **bool** | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. | [optional] 
**AutoTransitionToOperationalState** | Pointer to **bool** | Controls whether change components status to operational once scheduled maintenance completes. | [optional] 
**AutoTweetAtBeginning** | Pointer to **bool** | Controls whether tweet automatically when scheduled maintenance starts. | [optional] 
**AutoTweetOnCompletion** | Pointer to **bool** | Controls whether tweet automatically when scheduled maintenance completes. | [optional] 
**AutoTweetOnCreation** | Pointer to **bool** | Controls whether tweet automatically when scheduled maintenance is created. | [optional] 
**AutoTweetOneHourBefore** | Pointer to **bool** | Controls whether tweet automatically one hour before scheduled maintenance starts. | [optional] 
**BackfillDate** | Pointer to **string** | TimeStamp when incident was backfilled. | [optional] 
**Backfilled** | Pointer to **bool** | Controls whether incident is backfilled. If true, components cannot be specified. | [optional] 
**Body** | Pointer to **string** | The initial message, created as the first incident update. | [optional] 
**Components** | Pointer to **map[string]interface{}** | Map of status changes to apply to affected components | [optional] 
**ComponentIds** | Pointer to **[]string** | List of component_ids affected by this incident | [optional] 
**ScheduledAutoTransition** | Pointer to **bool** | Same as :scheduled_auto_transition_in_progress. Controls whether the incident is scheduled to automatically change to in progress. | [optional] 

## Methods

### NewPatchPagesPageIdIncidentsIncident

`func NewPatchPagesPageIdIncidentsIncident() *PatchPagesPageIdIncidentsIncident`

NewPatchPagesPageIdIncidentsIncident instantiates a new PatchPagesPageIdIncidentsIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPagesPageIdIncidentsIncidentWithDefaults

`func NewPatchPagesPageIdIncidentsIncidentWithDefaults() *PatchPagesPageIdIncidentsIncident`

NewPatchPagesPageIdIncidentsIncidentWithDefaults instantiates a new PatchPagesPageIdIncidentsIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchPagesPageIdIncidentsIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchPagesPageIdIncidentsIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchPagesPageIdIncidentsIncident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchPagesPageIdIncidentsIncident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PatchPagesPageIdIncidentsIncident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchPagesPageIdIncidentsIncident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchPagesPageIdIncidentsIncident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchPagesPageIdIncidentsIncident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImpactOverride

`func (o *PatchPagesPageIdIncidentsIncident) GetImpactOverride() string`

GetImpactOverride returns the ImpactOverride field if non-nil, zero value otherwise.

### GetImpactOverrideOk

`func (o *PatchPagesPageIdIncidentsIncident) GetImpactOverrideOk() (*string, bool)`

GetImpactOverrideOk returns a tuple with the ImpactOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactOverride

`func (o *PatchPagesPageIdIncidentsIncident) SetImpactOverride(v string)`

SetImpactOverride sets ImpactOverride field to given value.

### HasImpactOverride

`func (o *PatchPagesPageIdIncidentsIncident) HasImpactOverride() bool`

HasImpactOverride returns a boolean if a field has been set.

### GetScheduledFor

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *PatchPagesPageIdIncidentsIncident) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *PatchPagesPageIdIncidentsIncident) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetScheduledUntil

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledUntil() time.Time`

GetScheduledUntil returns the ScheduledUntil field if non-nil, zero value otherwise.

### GetScheduledUntilOk

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledUntilOk() (*time.Time, bool)`

GetScheduledUntilOk returns a tuple with the ScheduledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledUntil

`func (o *PatchPagesPageIdIncidentsIncident) SetScheduledUntil(v time.Time)`

SetScheduledUntil sets ScheduledUntil field to given value.

### HasScheduledUntil

`func (o *PatchPagesPageIdIncidentsIncident) HasScheduledUntil() bool`

HasScheduledUntil returns a boolean if a field has been set.

### GetScheduledRemindPrior

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledRemindPrior() bool`

GetScheduledRemindPrior returns the ScheduledRemindPrior field if non-nil, zero value otherwise.

### GetScheduledRemindPriorOk

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledRemindPriorOk() (*bool, bool)`

GetScheduledRemindPriorOk returns a tuple with the ScheduledRemindPrior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRemindPrior

`func (o *PatchPagesPageIdIncidentsIncident) SetScheduledRemindPrior(v bool)`

SetScheduledRemindPrior sets ScheduledRemindPrior field to given value.

### HasScheduledRemindPrior

`func (o *PatchPagesPageIdIncidentsIncident) HasScheduledRemindPrior() bool`

HasScheduledRemindPrior returns a boolean if a field has been set.

### GetScheduledAutoInProgress

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoInProgress() bool`

GetScheduledAutoInProgress returns the ScheduledAutoInProgress field if non-nil, zero value otherwise.

### GetScheduledAutoInProgressOk

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoInProgressOk() (*bool, bool)`

GetScheduledAutoInProgressOk returns a tuple with the ScheduledAutoInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoInProgress

`func (o *PatchPagesPageIdIncidentsIncident) SetScheduledAutoInProgress(v bool)`

SetScheduledAutoInProgress sets ScheduledAutoInProgress field to given value.

### HasScheduledAutoInProgress

`func (o *PatchPagesPageIdIncidentsIncident) HasScheduledAutoInProgress() bool`

HasScheduledAutoInProgress returns a boolean if a field has been set.

### GetScheduledAutoCompleted

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoCompleted() bool`

GetScheduledAutoCompleted returns the ScheduledAutoCompleted field if non-nil, zero value otherwise.

### GetScheduledAutoCompletedOk

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoCompletedOk() (*bool, bool)`

GetScheduledAutoCompletedOk returns a tuple with the ScheduledAutoCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoCompleted

`func (o *PatchPagesPageIdIncidentsIncident) SetScheduledAutoCompleted(v bool)`

SetScheduledAutoCompleted sets ScheduledAutoCompleted field to given value.

### HasScheduledAutoCompleted

`func (o *PatchPagesPageIdIncidentsIncident) HasScheduledAutoCompleted() bool`

HasScheduledAutoCompleted returns a boolean if a field has been set.

### GetMetadata

`func (o *PatchPagesPageIdIncidentsIncident) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchPagesPageIdIncidentsIncident) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchPagesPageIdIncidentsIncident) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchPagesPageIdIncidentsIncident) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDeliverNotifications

`func (o *PatchPagesPageIdIncidentsIncident) GetDeliverNotifications() bool`

GetDeliverNotifications returns the DeliverNotifications field if non-nil, zero value otherwise.

### GetDeliverNotificationsOk

`func (o *PatchPagesPageIdIncidentsIncident) GetDeliverNotificationsOk() (*bool, bool)`

GetDeliverNotificationsOk returns a tuple with the DeliverNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverNotifications

`func (o *PatchPagesPageIdIncidentsIncident) SetDeliverNotifications(v bool)`

SetDeliverNotifications sets DeliverNotifications field to given value.

### HasDeliverNotifications

`func (o *PatchPagesPageIdIncidentsIncident) HasDeliverNotifications() bool`

HasDeliverNotifications returns a boolean if a field has been set.

### GetAutoTransitionDeliverNotificationsAtEnd

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtEnd() bool`

GetAutoTransitionDeliverNotificationsAtEnd returns the AutoTransitionDeliverNotificationsAtEnd field if non-nil, zero value otherwise.

### GetAutoTransitionDeliverNotificationsAtEndOk

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtEndOk() (*bool, bool)`

GetAutoTransitionDeliverNotificationsAtEndOk returns a tuple with the AutoTransitionDeliverNotificationsAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionDeliverNotificationsAtEnd

`func (o *PatchPagesPageIdIncidentsIncident) SetAutoTransitionDeliverNotificationsAtEnd(v bool)`

SetAutoTransitionDeliverNotificationsAtEnd sets AutoTransitionDeliverNotificationsAtEnd field to given value.

### HasAutoTransitionDeliverNotificationsAtEnd

`func (o *PatchPagesPageIdIncidentsIncident) HasAutoTransitionDeliverNotificationsAtEnd() bool`

HasAutoTransitionDeliverNotificationsAtEnd returns a boolean if a field has been set.

### GetAutoTransitionDeliverNotificationsAtStart

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtStart() bool`

GetAutoTransitionDeliverNotificationsAtStart returns the AutoTransitionDeliverNotificationsAtStart field if non-nil, zero value otherwise.

### GetAutoTransitionDeliverNotificationsAtStartOk

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtStartOk() (*bool, bool)`

GetAutoTransitionDeliverNotificationsAtStartOk returns a tuple with the AutoTransitionDeliverNotificationsAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionDeliverNotificationsAtStart

`func (o *PatchPagesPageIdIncidentsIncident) SetAutoTransitionDeliverNotificationsAtStart(v bool)`

SetAutoTransitionDeliverNotificationsAtStart sets AutoTransitionDeliverNotificationsAtStart field to given value.

### HasAutoTransitionDeliverNotificationsAtStart

`func (o *PatchPagesPageIdIncidentsIncident) HasAutoTransitionDeliverNotificationsAtStart() bool`

HasAutoTransitionDeliverNotificationsAtStart returns a boolean if a field has been set.

### GetAutoTransitionToMaintenanceState

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionToMaintenanceState() bool`

GetAutoTransitionToMaintenanceState returns the AutoTransitionToMaintenanceState field if non-nil, zero value otherwise.

### GetAutoTransitionToMaintenanceStateOk

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionToMaintenanceStateOk() (*bool, bool)`

GetAutoTransitionToMaintenanceStateOk returns a tuple with the AutoTransitionToMaintenanceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionToMaintenanceState

`func (o *PatchPagesPageIdIncidentsIncident) SetAutoTransitionToMaintenanceState(v bool)`

SetAutoTransitionToMaintenanceState sets AutoTransitionToMaintenanceState field to given value.

### HasAutoTransitionToMaintenanceState

`func (o *PatchPagesPageIdIncidentsIncident) HasAutoTransitionToMaintenanceState() bool`

HasAutoTransitionToMaintenanceState returns a boolean if a field has been set.

### GetAutoTransitionToOperationalState

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionToOperationalState() bool`

GetAutoTransitionToOperationalState returns the AutoTransitionToOperationalState field if non-nil, zero value otherwise.

### GetAutoTransitionToOperationalStateOk

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTransitionToOperationalStateOk() (*bool, bool)`

GetAutoTransitionToOperationalStateOk returns a tuple with the AutoTransitionToOperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionToOperationalState

`func (o *PatchPagesPageIdIncidentsIncident) SetAutoTransitionToOperationalState(v bool)`

SetAutoTransitionToOperationalState sets AutoTransitionToOperationalState field to given value.

### HasAutoTransitionToOperationalState

`func (o *PatchPagesPageIdIncidentsIncident) HasAutoTransitionToOperationalState() bool`

HasAutoTransitionToOperationalState returns a boolean if a field has been set.

### GetAutoTweetAtBeginning

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetAtBeginning() bool`

GetAutoTweetAtBeginning returns the AutoTweetAtBeginning field if non-nil, zero value otherwise.

### GetAutoTweetAtBeginningOk

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetAtBeginningOk() (*bool, bool)`

GetAutoTweetAtBeginningOk returns a tuple with the AutoTweetAtBeginning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetAtBeginning

`func (o *PatchPagesPageIdIncidentsIncident) SetAutoTweetAtBeginning(v bool)`

SetAutoTweetAtBeginning sets AutoTweetAtBeginning field to given value.

### HasAutoTweetAtBeginning

`func (o *PatchPagesPageIdIncidentsIncident) HasAutoTweetAtBeginning() bool`

HasAutoTweetAtBeginning returns a boolean if a field has been set.

### GetAutoTweetOnCompletion

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOnCompletion() bool`

GetAutoTweetOnCompletion returns the AutoTweetOnCompletion field if non-nil, zero value otherwise.

### GetAutoTweetOnCompletionOk

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOnCompletionOk() (*bool, bool)`

GetAutoTweetOnCompletionOk returns a tuple with the AutoTweetOnCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetOnCompletion

`func (o *PatchPagesPageIdIncidentsIncident) SetAutoTweetOnCompletion(v bool)`

SetAutoTweetOnCompletion sets AutoTweetOnCompletion field to given value.

### HasAutoTweetOnCompletion

`func (o *PatchPagesPageIdIncidentsIncident) HasAutoTweetOnCompletion() bool`

HasAutoTweetOnCompletion returns a boolean if a field has been set.

### GetAutoTweetOnCreation

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOnCreation() bool`

GetAutoTweetOnCreation returns the AutoTweetOnCreation field if non-nil, zero value otherwise.

### GetAutoTweetOnCreationOk

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOnCreationOk() (*bool, bool)`

GetAutoTweetOnCreationOk returns a tuple with the AutoTweetOnCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetOnCreation

`func (o *PatchPagesPageIdIncidentsIncident) SetAutoTweetOnCreation(v bool)`

SetAutoTweetOnCreation sets AutoTweetOnCreation field to given value.

### HasAutoTweetOnCreation

`func (o *PatchPagesPageIdIncidentsIncident) HasAutoTweetOnCreation() bool`

HasAutoTweetOnCreation returns a boolean if a field has been set.

### GetAutoTweetOneHourBefore

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOneHourBefore() bool`

GetAutoTweetOneHourBefore returns the AutoTweetOneHourBefore field if non-nil, zero value otherwise.

### GetAutoTweetOneHourBeforeOk

`func (o *PatchPagesPageIdIncidentsIncident) GetAutoTweetOneHourBeforeOk() (*bool, bool)`

GetAutoTweetOneHourBeforeOk returns a tuple with the AutoTweetOneHourBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetOneHourBefore

`func (o *PatchPagesPageIdIncidentsIncident) SetAutoTweetOneHourBefore(v bool)`

SetAutoTweetOneHourBefore sets AutoTweetOneHourBefore field to given value.

### HasAutoTweetOneHourBefore

`func (o *PatchPagesPageIdIncidentsIncident) HasAutoTweetOneHourBefore() bool`

HasAutoTweetOneHourBefore returns a boolean if a field has been set.

### GetBackfillDate

`func (o *PatchPagesPageIdIncidentsIncident) GetBackfillDate() string`

GetBackfillDate returns the BackfillDate field if non-nil, zero value otherwise.

### GetBackfillDateOk

`func (o *PatchPagesPageIdIncidentsIncident) GetBackfillDateOk() (*string, bool)`

GetBackfillDateOk returns a tuple with the BackfillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillDate

`func (o *PatchPagesPageIdIncidentsIncident) SetBackfillDate(v string)`

SetBackfillDate sets BackfillDate field to given value.

### HasBackfillDate

`func (o *PatchPagesPageIdIncidentsIncident) HasBackfillDate() bool`

HasBackfillDate returns a boolean if a field has been set.

### GetBackfilled

`func (o *PatchPagesPageIdIncidentsIncident) GetBackfilled() bool`

GetBackfilled returns the Backfilled field if non-nil, zero value otherwise.

### GetBackfilledOk

`func (o *PatchPagesPageIdIncidentsIncident) GetBackfilledOk() (*bool, bool)`

GetBackfilledOk returns a tuple with the Backfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilled

`func (o *PatchPagesPageIdIncidentsIncident) SetBackfilled(v bool)`

SetBackfilled sets Backfilled field to given value.

### HasBackfilled

`func (o *PatchPagesPageIdIncidentsIncident) HasBackfilled() bool`

HasBackfilled returns a boolean if a field has been set.

### GetBody

`func (o *PatchPagesPageIdIncidentsIncident) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PatchPagesPageIdIncidentsIncident) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PatchPagesPageIdIncidentsIncident) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PatchPagesPageIdIncidentsIncident) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetComponents

`func (o *PatchPagesPageIdIncidentsIncident) GetComponents() map[string]interface{}`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PatchPagesPageIdIncidentsIncident) GetComponentsOk() (*map[string]interface{}, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PatchPagesPageIdIncidentsIncident) SetComponents(v map[string]interface{})`

SetComponents sets Components field to given value.

### HasComponents

`func (o *PatchPagesPageIdIncidentsIncident) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetComponentIds

`func (o *PatchPagesPageIdIncidentsIncident) GetComponentIds() []string`

GetComponentIds returns the ComponentIds field if non-nil, zero value otherwise.

### GetComponentIdsOk

`func (o *PatchPagesPageIdIncidentsIncident) GetComponentIdsOk() (*[]string, bool)`

GetComponentIdsOk returns a tuple with the ComponentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIds

`func (o *PatchPagesPageIdIncidentsIncident) SetComponentIds(v []string)`

SetComponentIds sets ComponentIds field to given value.

### HasComponentIds

`func (o *PatchPagesPageIdIncidentsIncident) HasComponentIds() bool`

HasComponentIds returns a boolean if a field has been set.

### GetScheduledAutoTransition

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoTransition() bool`

GetScheduledAutoTransition returns the ScheduledAutoTransition field if non-nil, zero value otherwise.

### GetScheduledAutoTransitionOk

`func (o *PatchPagesPageIdIncidentsIncident) GetScheduledAutoTransitionOk() (*bool, bool)`

GetScheduledAutoTransitionOk returns a tuple with the ScheduledAutoTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoTransition

`func (o *PatchPagesPageIdIncidentsIncident) SetScheduledAutoTransition(v bool)`

SetScheduledAutoTransition sets ScheduledAutoTransition field to given value.

### HasScheduledAutoTransition

`func (o *PatchPagesPageIdIncidentsIncident) HasScheduledAutoTransition() bool`

HasScheduledAutoTransition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


