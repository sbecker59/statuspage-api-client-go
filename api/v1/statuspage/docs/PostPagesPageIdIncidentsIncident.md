# PostPagesPageIdIncidentsIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Incident Name. There is a maximum limit of 255 characters. | 
**Status** | Pointer to **string** | The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed. | [optional] 
**ImpactOverride** | Pointer to **string** | value to override calculated impact value | [optional] 
**ScheduledFor** | Pointer to **time.Time** | The timestamp the incident is scheduled for. | [optional] [default to "2013-05-07T03:00:00.007Z"]
**ScheduledUntil** | Pointer to **time.Time** | The timestamp the incident is scheduled until. | [optional] [default to "2013-05-07T06:00:00.007Z"]
**ScheduledRemindPrior** | Pointer to **bool** | Controls whether to remind subscribers prior to scheduled incidents. | [optional] 
**AutoTransitionToMaintenanceState** | Pointer to **bool** | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. | [optional] 
**AutoTransitionToOperationalState** | Pointer to **bool** | Controls whether change components status to operational once scheduled maintenance completes. | [optional] 
**ScheduledAutoInProgress** | Pointer to **bool** | Controls whether the incident is scheduled to automatically change to in progress. | [optional] 
**ScheduledAutoCompleted** | Pointer to **bool** | Controls whether the incident is scheduled to automatically change to complete. | [optional] 
**AutoTransitionDeliverNotificationsAtStart** | Pointer to **bool** | Controls whether send notification when scheduled maintenances auto transition to started. | [optional] 
**AutoTransitionDeliverNotificationsAtEnd** | Pointer to **bool** | Controls whether send notification when scheduled maintenances auto transition to completed. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Attach a json object to the incident. All top-level values in the object must also be objects. | [optional] 
**DeliverNotifications** | Pointer to **bool** | Deliver notifications to subscribers if this is true. If this is false, create an incident without notifying customers. | [optional] [default to true]
**AutoTweetAtBeginning** | Pointer to **bool** | Controls whether tweet automatically when scheduled maintenance starts. | [optional] 
**AutoTweetOnCompletion** | Pointer to **bool** | Controls whether tweet automatically when scheduled maintenance completes. | [optional] 
**AutoTweetOnCreation** | Pointer to **bool** | Controls whether tweet automatically when scheduled maintenance is created. | [optional] 
**AutoTweetOneHourBefore** | Pointer to **bool** | Controls whether tweet automatically one hour before scheduled maintenance starts. | [optional] 
**BackfillDate** | Pointer to **string** | TimeStamp when incident was backfilled. | [optional] 
**Backfilled** | Pointer to **bool** | Controls whether incident is backfilled. If true, components cannot be specified. | [optional] 
**Body** | Pointer to **string** | The initial message, created as the first incident update. There is a maximum limit of 25000 characters | [optional] 
**Components** | Pointer to [**PostPagesPageIdIncidentsIncidentComponents**](PostPagesPageIdIncidentsIncidentComponents.md) |  | [optional] 
**ComponentIds** | Pointer to **[]string** | List of component_ids affected by this incident | [optional] 
**ScheduledAutoTransition** | Pointer to **bool** | Same as :scheduled_auto_transition_in_progress. Controls whether the incident is scheduled to automatically change to in progress. | [optional] 

## Methods

### NewPostPagesPageIdIncidentsIncident

`func NewPostPagesPageIdIncidentsIncident(name string, ) *PostPagesPageIdIncidentsIncident`

NewPostPagesPageIdIncidentsIncident instantiates a new PostPagesPageIdIncidentsIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPagesPageIdIncidentsIncidentWithDefaults

`func NewPostPagesPageIdIncidentsIncidentWithDefaults() *PostPagesPageIdIncidentsIncident`

NewPostPagesPageIdIncidentsIncidentWithDefaults instantiates a new PostPagesPageIdIncidentsIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostPagesPageIdIncidentsIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostPagesPageIdIncidentsIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostPagesPageIdIncidentsIncident) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *PostPagesPageIdIncidentsIncident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostPagesPageIdIncidentsIncident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostPagesPageIdIncidentsIncident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostPagesPageIdIncidentsIncident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImpactOverride

`func (o *PostPagesPageIdIncidentsIncident) GetImpactOverride() string`

GetImpactOverride returns the ImpactOverride field if non-nil, zero value otherwise.

### GetImpactOverrideOk

`func (o *PostPagesPageIdIncidentsIncident) GetImpactOverrideOk() (*string, bool)`

GetImpactOverrideOk returns a tuple with the ImpactOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactOverride

`func (o *PostPagesPageIdIncidentsIncident) SetImpactOverride(v string)`

SetImpactOverride sets ImpactOverride field to given value.

### HasImpactOverride

`func (o *PostPagesPageIdIncidentsIncident) HasImpactOverride() bool`

HasImpactOverride returns a boolean if a field has been set.

### GetScheduledFor

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *PostPagesPageIdIncidentsIncident) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *PostPagesPageIdIncidentsIncident) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetScheduledUntil

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledUntil() time.Time`

GetScheduledUntil returns the ScheduledUntil field if non-nil, zero value otherwise.

### GetScheduledUntilOk

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledUntilOk() (*time.Time, bool)`

GetScheduledUntilOk returns a tuple with the ScheduledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledUntil

`func (o *PostPagesPageIdIncidentsIncident) SetScheduledUntil(v time.Time)`

SetScheduledUntil sets ScheduledUntil field to given value.

### HasScheduledUntil

`func (o *PostPagesPageIdIncidentsIncident) HasScheduledUntil() bool`

HasScheduledUntil returns a boolean if a field has been set.

### GetScheduledRemindPrior

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledRemindPrior() bool`

GetScheduledRemindPrior returns the ScheduledRemindPrior field if non-nil, zero value otherwise.

### GetScheduledRemindPriorOk

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledRemindPriorOk() (*bool, bool)`

GetScheduledRemindPriorOk returns a tuple with the ScheduledRemindPrior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRemindPrior

`func (o *PostPagesPageIdIncidentsIncident) SetScheduledRemindPrior(v bool)`

SetScheduledRemindPrior sets ScheduledRemindPrior field to given value.

### HasScheduledRemindPrior

`func (o *PostPagesPageIdIncidentsIncident) HasScheduledRemindPrior() bool`

HasScheduledRemindPrior returns a boolean if a field has been set.

### GetAutoTransitionToMaintenanceState

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTransitionToMaintenanceState() bool`

GetAutoTransitionToMaintenanceState returns the AutoTransitionToMaintenanceState field if non-nil, zero value otherwise.

### GetAutoTransitionToMaintenanceStateOk

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTransitionToMaintenanceStateOk() (*bool, bool)`

GetAutoTransitionToMaintenanceStateOk returns a tuple with the AutoTransitionToMaintenanceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionToMaintenanceState

`func (o *PostPagesPageIdIncidentsIncident) SetAutoTransitionToMaintenanceState(v bool)`

SetAutoTransitionToMaintenanceState sets AutoTransitionToMaintenanceState field to given value.

### HasAutoTransitionToMaintenanceState

`func (o *PostPagesPageIdIncidentsIncident) HasAutoTransitionToMaintenanceState() bool`

HasAutoTransitionToMaintenanceState returns a boolean if a field has been set.

### GetAutoTransitionToOperationalState

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTransitionToOperationalState() bool`

GetAutoTransitionToOperationalState returns the AutoTransitionToOperationalState field if non-nil, zero value otherwise.

### GetAutoTransitionToOperationalStateOk

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTransitionToOperationalStateOk() (*bool, bool)`

GetAutoTransitionToOperationalStateOk returns a tuple with the AutoTransitionToOperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionToOperationalState

`func (o *PostPagesPageIdIncidentsIncident) SetAutoTransitionToOperationalState(v bool)`

SetAutoTransitionToOperationalState sets AutoTransitionToOperationalState field to given value.

### HasAutoTransitionToOperationalState

`func (o *PostPagesPageIdIncidentsIncident) HasAutoTransitionToOperationalState() bool`

HasAutoTransitionToOperationalState returns a boolean if a field has been set.

### GetScheduledAutoInProgress

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledAutoInProgress() bool`

GetScheduledAutoInProgress returns the ScheduledAutoInProgress field if non-nil, zero value otherwise.

### GetScheduledAutoInProgressOk

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledAutoInProgressOk() (*bool, bool)`

GetScheduledAutoInProgressOk returns a tuple with the ScheduledAutoInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoInProgress

`func (o *PostPagesPageIdIncidentsIncident) SetScheduledAutoInProgress(v bool)`

SetScheduledAutoInProgress sets ScheduledAutoInProgress field to given value.

### HasScheduledAutoInProgress

`func (o *PostPagesPageIdIncidentsIncident) HasScheduledAutoInProgress() bool`

HasScheduledAutoInProgress returns a boolean if a field has been set.

### GetScheduledAutoCompleted

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledAutoCompleted() bool`

GetScheduledAutoCompleted returns the ScheduledAutoCompleted field if non-nil, zero value otherwise.

### GetScheduledAutoCompletedOk

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledAutoCompletedOk() (*bool, bool)`

GetScheduledAutoCompletedOk returns a tuple with the ScheduledAutoCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoCompleted

`func (o *PostPagesPageIdIncidentsIncident) SetScheduledAutoCompleted(v bool)`

SetScheduledAutoCompleted sets ScheduledAutoCompleted field to given value.

### HasScheduledAutoCompleted

`func (o *PostPagesPageIdIncidentsIncident) HasScheduledAutoCompleted() bool`

HasScheduledAutoCompleted returns a boolean if a field has been set.

### GetAutoTransitionDeliverNotificationsAtStart

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtStart() bool`

GetAutoTransitionDeliverNotificationsAtStart returns the AutoTransitionDeliverNotificationsAtStart field if non-nil, zero value otherwise.

### GetAutoTransitionDeliverNotificationsAtStartOk

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtStartOk() (*bool, bool)`

GetAutoTransitionDeliverNotificationsAtStartOk returns a tuple with the AutoTransitionDeliverNotificationsAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionDeliverNotificationsAtStart

`func (o *PostPagesPageIdIncidentsIncident) SetAutoTransitionDeliverNotificationsAtStart(v bool)`

SetAutoTransitionDeliverNotificationsAtStart sets AutoTransitionDeliverNotificationsAtStart field to given value.

### HasAutoTransitionDeliverNotificationsAtStart

`func (o *PostPagesPageIdIncidentsIncident) HasAutoTransitionDeliverNotificationsAtStart() bool`

HasAutoTransitionDeliverNotificationsAtStart returns a boolean if a field has been set.

### GetAutoTransitionDeliverNotificationsAtEnd

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtEnd() bool`

GetAutoTransitionDeliverNotificationsAtEnd returns the AutoTransitionDeliverNotificationsAtEnd field if non-nil, zero value otherwise.

### GetAutoTransitionDeliverNotificationsAtEndOk

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTransitionDeliverNotificationsAtEndOk() (*bool, bool)`

GetAutoTransitionDeliverNotificationsAtEndOk returns a tuple with the AutoTransitionDeliverNotificationsAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionDeliverNotificationsAtEnd

`func (o *PostPagesPageIdIncidentsIncident) SetAutoTransitionDeliverNotificationsAtEnd(v bool)`

SetAutoTransitionDeliverNotificationsAtEnd sets AutoTransitionDeliverNotificationsAtEnd field to given value.

### HasAutoTransitionDeliverNotificationsAtEnd

`func (o *PostPagesPageIdIncidentsIncident) HasAutoTransitionDeliverNotificationsAtEnd() bool`

HasAutoTransitionDeliverNotificationsAtEnd returns a boolean if a field has been set.

### GetMetadata

`func (o *PostPagesPageIdIncidentsIncident) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostPagesPageIdIncidentsIncident) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostPagesPageIdIncidentsIncident) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PostPagesPageIdIncidentsIncident) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDeliverNotifications

`func (o *PostPagesPageIdIncidentsIncident) GetDeliverNotifications() bool`

GetDeliverNotifications returns the DeliverNotifications field if non-nil, zero value otherwise.

### GetDeliverNotificationsOk

`func (o *PostPagesPageIdIncidentsIncident) GetDeliverNotificationsOk() (*bool, bool)`

GetDeliverNotificationsOk returns a tuple with the DeliverNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverNotifications

`func (o *PostPagesPageIdIncidentsIncident) SetDeliverNotifications(v bool)`

SetDeliverNotifications sets DeliverNotifications field to given value.

### HasDeliverNotifications

`func (o *PostPagesPageIdIncidentsIncident) HasDeliverNotifications() bool`

HasDeliverNotifications returns a boolean if a field has been set.

### GetAutoTweetAtBeginning

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTweetAtBeginning() bool`

GetAutoTweetAtBeginning returns the AutoTweetAtBeginning field if non-nil, zero value otherwise.

### GetAutoTweetAtBeginningOk

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTweetAtBeginningOk() (*bool, bool)`

GetAutoTweetAtBeginningOk returns a tuple with the AutoTweetAtBeginning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetAtBeginning

`func (o *PostPagesPageIdIncidentsIncident) SetAutoTweetAtBeginning(v bool)`

SetAutoTweetAtBeginning sets AutoTweetAtBeginning field to given value.

### HasAutoTweetAtBeginning

`func (o *PostPagesPageIdIncidentsIncident) HasAutoTweetAtBeginning() bool`

HasAutoTweetAtBeginning returns a boolean if a field has been set.

### GetAutoTweetOnCompletion

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTweetOnCompletion() bool`

GetAutoTweetOnCompletion returns the AutoTweetOnCompletion field if non-nil, zero value otherwise.

### GetAutoTweetOnCompletionOk

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTweetOnCompletionOk() (*bool, bool)`

GetAutoTweetOnCompletionOk returns a tuple with the AutoTweetOnCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetOnCompletion

`func (o *PostPagesPageIdIncidentsIncident) SetAutoTweetOnCompletion(v bool)`

SetAutoTweetOnCompletion sets AutoTweetOnCompletion field to given value.

### HasAutoTweetOnCompletion

`func (o *PostPagesPageIdIncidentsIncident) HasAutoTweetOnCompletion() bool`

HasAutoTweetOnCompletion returns a boolean if a field has been set.

### GetAutoTweetOnCreation

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTweetOnCreation() bool`

GetAutoTweetOnCreation returns the AutoTweetOnCreation field if non-nil, zero value otherwise.

### GetAutoTweetOnCreationOk

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTweetOnCreationOk() (*bool, bool)`

GetAutoTweetOnCreationOk returns a tuple with the AutoTweetOnCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetOnCreation

`func (o *PostPagesPageIdIncidentsIncident) SetAutoTweetOnCreation(v bool)`

SetAutoTweetOnCreation sets AutoTweetOnCreation field to given value.

### HasAutoTweetOnCreation

`func (o *PostPagesPageIdIncidentsIncident) HasAutoTweetOnCreation() bool`

HasAutoTweetOnCreation returns a boolean if a field has been set.

### GetAutoTweetOneHourBefore

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTweetOneHourBefore() bool`

GetAutoTweetOneHourBefore returns the AutoTweetOneHourBefore field if non-nil, zero value otherwise.

### GetAutoTweetOneHourBeforeOk

`func (o *PostPagesPageIdIncidentsIncident) GetAutoTweetOneHourBeforeOk() (*bool, bool)`

GetAutoTweetOneHourBeforeOk returns a tuple with the AutoTweetOneHourBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetOneHourBefore

`func (o *PostPagesPageIdIncidentsIncident) SetAutoTweetOneHourBefore(v bool)`

SetAutoTweetOneHourBefore sets AutoTweetOneHourBefore field to given value.

### HasAutoTweetOneHourBefore

`func (o *PostPagesPageIdIncidentsIncident) HasAutoTweetOneHourBefore() bool`

HasAutoTweetOneHourBefore returns a boolean if a field has been set.

### GetBackfillDate

`func (o *PostPagesPageIdIncidentsIncident) GetBackfillDate() string`

GetBackfillDate returns the BackfillDate field if non-nil, zero value otherwise.

### GetBackfillDateOk

`func (o *PostPagesPageIdIncidentsIncident) GetBackfillDateOk() (*string, bool)`

GetBackfillDateOk returns a tuple with the BackfillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillDate

`func (o *PostPagesPageIdIncidentsIncident) SetBackfillDate(v string)`

SetBackfillDate sets BackfillDate field to given value.

### HasBackfillDate

`func (o *PostPagesPageIdIncidentsIncident) HasBackfillDate() bool`

HasBackfillDate returns a boolean if a field has been set.

### GetBackfilled

`func (o *PostPagesPageIdIncidentsIncident) GetBackfilled() bool`

GetBackfilled returns the Backfilled field if non-nil, zero value otherwise.

### GetBackfilledOk

`func (o *PostPagesPageIdIncidentsIncident) GetBackfilledOk() (*bool, bool)`

GetBackfilledOk returns a tuple with the Backfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilled

`func (o *PostPagesPageIdIncidentsIncident) SetBackfilled(v bool)`

SetBackfilled sets Backfilled field to given value.

### HasBackfilled

`func (o *PostPagesPageIdIncidentsIncident) HasBackfilled() bool`

HasBackfilled returns a boolean if a field has been set.

### GetBody

`func (o *PostPagesPageIdIncidentsIncident) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PostPagesPageIdIncidentsIncident) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PostPagesPageIdIncidentsIncident) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PostPagesPageIdIncidentsIncident) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetComponents

`func (o *PostPagesPageIdIncidentsIncident) GetComponents() PostPagesPageIdIncidentsIncidentComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PostPagesPageIdIncidentsIncident) GetComponentsOk() (*PostPagesPageIdIncidentsIncidentComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PostPagesPageIdIncidentsIncident) SetComponents(v PostPagesPageIdIncidentsIncidentComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *PostPagesPageIdIncidentsIncident) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetComponentIds

`func (o *PostPagesPageIdIncidentsIncident) GetComponentIds() []string`

GetComponentIds returns the ComponentIds field if non-nil, zero value otherwise.

### GetComponentIdsOk

`func (o *PostPagesPageIdIncidentsIncident) GetComponentIdsOk() (*[]string, bool)`

GetComponentIdsOk returns a tuple with the ComponentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIds

`func (o *PostPagesPageIdIncidentsIncident) SetComponentIds(v []string)`

SetComponentIds sets ComponentIds field to given value.

### HasComponentIds

`func (o *PostPagesPageIdIncidentsIncident) HasComponentIds() bool`

HasComponentIds returns a boolean if a field has been set.

### GetScheduledAutoTransition

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledAutoTransition() bool`

GetScheduledAutoTransition returns the ScheduledAutoTransition field if non-nil, zero value otherwise.

### GetScheduledAutoTransitionOk

`func (o *PostPagesPageIdIncidentsIncident) GetScheduledAutoTransitionOk() (*bool, bool)`

GetScheduledAutoTransitionOk returns a tuple with the ScheduledAutoTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoTransition

`func (o *PostPagesPageIdIncidentsIncident) SetScheduledAutoTransition(v bool)`

SetScheduledAutoTransition sets ScheduledAutoTransition field to given value.

### HasScheduledAutoTransition

`func (o *PostPagesPageIdIncidentsIncident) HasScheduledAutoTransition() bool`

HasScheduledAutoTransition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


