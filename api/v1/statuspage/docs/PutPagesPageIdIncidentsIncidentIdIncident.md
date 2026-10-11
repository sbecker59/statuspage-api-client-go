# PutPagesPageIdIncidentsIncidentIdIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Incident Name. There is a maximum limit of 255 characters. | [optional] 
**Status** | Pointer to **string** | The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed. | [optional] 
**ImpactOverride** | Pointer to **string** | value to override calculated impact value | [optional] 
**ScheduledFor** | Pointer to **time.Time** | The timestamp the incident is scheduled for. | [optional] 
**ScheduledUntil** | Pointer to **time.Time** | The timestamp the incident is scheduled until. | [optional] 
**ScheduledRemindPrior** | Pointer to **bool** | Controls whether to remind subscribers prior to scheduled incidents. | [optional] 
**AutoTransitionToMaintenanceState** | Pointer to **bool** | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. | [optional] 
**AutoTransitionToOperationalState** | Pointer to **bool** | Controls whether change components status to operational once scheduled maintenance completes. | [optional] 
**ScheduledAutoInProgress** | Pointer to **bool** | Controls whether the incident is scheduled to automatically change to in progress. | [optional] 
**ScheduledAutoCompleted** | Pointer to **bool** | Controls whether the incident is scheduled to automatically change to complete. | [optional] 
**AutoTransitionDeliverNotificationsAtStart** | Pointer to **bool** | Controls whether send notification when scheduled maintenances auto transition to started. | [optional] 
**AutoTransitionDeliverNotificationsAtEnd** | Pointer to **bool** | Controls whether send notification when scheduled maintenances auto transition to completed. | [optional] 
**ReminderIntervals** | Pointer to **string** | Custom reminder intervals for unresolved/open incidents. Not applicable for &lt;strong&gt;Scheduled maintenance&lt;/strong&gt;&lt;br&gt;There are 4 possible states for reminder_intervals:&lt;br&gt;&lt;strong&gt;DEFAULT:&lt;/strong&gt; NULL, representing a default behavior with intervals [3, 6, 12, 24].&lt;br&gt;&lt;strong&gt;AFTER:&lt;/strong&gt; A serialized array of strictly increasing intervals, each integer ranges from [1-24] (inclusive). Ex \&quot;[1, 5, 7, 10]\&quot;&lt;br&gt;&lt;strong&gt;EVERY:&lt;/strong&gt; An integer in the range [1-24] as a string, representing equal intervals. Ex \&quot;4\&quot; for [4, 8, 12, 16, 20, 24]&lt;br&gt;&lt;strong&gt;OFF:&lt;/strong&gt; A serialized empty array, for example, \&quot;[]\&quot;, meaning no reminder notifications will be sent. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Attach a json object to the incident. All top-level values in the object must also be objects. | [optional] [default to {}]
**DeliverNotifications** | Pointer to **bool** | Deliver notifications to subscribers if this is true. If this is false, create an incident without notifying customers. | [optional] [default to true]
**AutoTweetAtBeginning** | Pointer to **bool** | Controls whether tweet automatically when scheduled maintenance starts. | [optional] 
**AutoTweetOnCompletion** | Pointer to **bool** | Controls whether tweet automatically when scheduled maintenance completes. | [optional] 
**AutoTweetOnCreation** | Pointer to **bool** | Controls whether tweet automatically when scheduled maintenance is created. | [optional] 
**AutoTweetOneHourBefore** | Pointer to **bool** | Controls whether tweet automatically one hour before scheduled maintenance starts. | [optional] 
**BackfillDate** | Pointer to **string** | TimeStamp when incident was backfilled. | [optional] 
**Backfilled** | Pointer to **bool** | Controls whether incident is backfilled. If true, components cannot be specified. | [optional] 
**Body** | Pointer to **string** | The initial message, created as the first incident update. There is a maximum limit of 25000 characters | [optional] 
**Components** | Pointer to **map[string]string** | Map of status changes to apply to affected components | [optional] 
**ComponentIds** | Pointer to **[]string** | List of component_ids affected by this incident | [optional] 
**ScheduledAutoTransition** | Pointer to **bool** | Same as :scheduled_auto_transition_in_progress. Controls whether the incident is scheduled to automatically change to in progress. | [optional] 

## Methods

### NewPutPagesPageIdIncidentsIncidentIdIncident

`func NewPutPagesPageIdIncidentsIncidentIdIncident() *PutPagesPageIdIncidentsIncidentIdIncident`

NewPutPagesPageIdIncidentsIncidentIdIncident instantiates a new PutPagesPageIdIncidentsIncidentIdIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutPagesPageIdIncidentsIncidentIdIncidentWithDefaults

`func NewPutPagesPageIdIncidentsIncidentIdIncidentWithDefaults() *PutPagesPageIdIncidentsIncidentIdIncident`

NewPutPagesPageIdIncidentsIncidentIdIncidentWithDefaults instantiates a new PutPagesPageIdIncidentsIncidentIdIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImpactOverride

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetImpactOverride() string`

GetImpactOverride returns the ImpactOverride field if non-nil, zero value otherwise.

### GetImpactOverrideOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetImpactOverrideOk() (*string, bool)`

GetImpactOverrideOk returns a tuple with the ImpactOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactOverride

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetImpactOverride(v string)`

SetImpactOverride sets ImpactOverride field to given value.

### HasImpactOverride

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasImpactOverride() bool`

HasImpactOverride returns a boolean if a field has been set.

### GetScheduledFor

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetScheduledUntil

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledUntil() time.Time`

GetScheduledUntil returns the ScheduledUntil field if non-nil, zero value otherwise.

### GetScheduledUntilOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledUntilOk() (*time.Time, bool)`

GetScheduledUntilOk returns a tuple with the ScheduledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledUntil

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetScheduledUntil(v time.Time)`

SetScheduledUntil sets ScheduledUntil field to given value.

### HasScheduledUntil

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasScheduledUntil() bool`

HasScheduledUntil returns a boolean if a field has been set.

### GetScheduledRemindPrior

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledRemindPrior() bool`

GetScheduledRemindPrior returns the ScheduledRemindPrior field if non-nil, zero value otherwise.

### GetScheduledRemindPriorOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledRemindPriorOk() (*bool, bool)`

GetScheduledRemindPriorOk returns a tuple with the ScheduledRemindPrior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRemindPrior

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetScheduledRemindPrior(v bool)`

SetScheduledRemindPrior sets ScheduledRemindPrior field to given value.

### HasScheduledRemindPrior

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasScheduledRemindPrior() bool`

HasScheduledRemindPrior returns a boolean if a field has been set.

### GetAutoTransitionToMaintenanceState

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTransitionToMaintenanceState() bool`

GetAutoTransitionToMaintenanceState returns the AutoTransitionToMaintenanceState field if non-nil, zero value otherwise.

### GetAutoTransitionToMaintenanceStateOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTransitionToMaintenanceStateOk() (*bool, bool)`

GetAutoTransitionToMaintenanceStateOk returns a tuple with the AutoTransitionToMaintenanceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionToMaintenanceState

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetAutoTransitionToMaintenanceState(v bool)`

SetAutoTransitionToMaintenanceState sets AutoTransitionToMaintenanceState field to given value.

### HasAutoTransitionToMaintenanceState

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasAutoTransitionToMaintenanceState() bool`

HasAutoTransitionToMaintenanceState returns a boolean if a field has been set.

### GetAutoTransitionToOperationalState

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTransitionToOperationalState() bool`

GetAutoTransitionToOperationalState returns the AutoTransitionToOperationalState field if non-nil, zero value otherwise.

### GetAutoTransitionToOperationalStateOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTransitionToOperationalStateOk() (*bool, bool)`

GetAutoTransitionToOperationalStateOk returns a tuple with the AutoTransitionToOperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionToOperationalState

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetAutoTransitionToOperationalState(v bool)`

SetAutoTransitionToOperationalState sets AutoTransitionToOperationalState field to given value.

### HasAutoTransitionToOperationalState

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasAutoTransitionToOperationalState() bool`

HasAutoTransitionToOperationalState returns a boolean if a field has been set.

### GetScheduledAutoInProgress

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledAutoInProgress() bool`

GetScheduledAutoInProgress returns the ScheduledAutoInProgress field if non-nil, zero value otherwise.

### GetScheduledAutoInProgressOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledAutoInProgressOk() (*bool, bool)`

GetScheduledAutoInProgressOk returns a tuple with the ScheduledAutoInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoInProgress

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetScheduledAutoInProgress(v bool)`

SetScheduledAutoInProgress sets ScheduledAutoInProgress field to given value.

### HasScheduledAutoInProgress

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasScheduledAutoInProgress() bool`

HasScheduledAutoInProgress returns a boolean if a field has been set.

### GetScheduledAutoCompleted

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledAutoCompleted() bool`

GetScheduledAutoCompleted returns the ScheduledAutoCompleted field if non-nil, zero value otherwise.

### GetScheduledAutoCompletedOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledAutoCompletedOk() (*bool, bool)`

GetScheduledAutoCompletedOk returns a tuple with the ScheduledAutoCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoCompleted

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetScheduledAutoCompleted(v bool)`

SetScheduledAutoCompleted sets ScheduledAutoCompleted field to given value.

### HasScheduledAutoCompleted

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasScheduledAutoCompleted() bool`

HasScheduledAutoCompleted returns a boolean if a field has been set.

### GetAutoTransitionDeliverNotificationsAtStart

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTransitionDeliverNotificationsAtStart() bool`

GetAutoTransitionDeliverNotificationsAtStart returns the AutoTransitionDeliverNotificationsAtStart field if non-nil, zero value otherwise.

### GetAutoTransitionDeliverNotificationsAtStartOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTransitionDeliverNotificationsAtStartOk() (*bool, bool)`

GetAutoTransitionDeliverNotificationsAtStartOk returns a tuple with the AutoTransitionDeliverNotificationsAtStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionDeliverNotificationsAtStart

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetAutoTransitionDeliverNotificationsAtStart(v bool)`

SetAutoTransitionDeliverNotificationsAtStart sets AutoTransitionDeliverNotificationsAtStart field to given value.

### HasAutoTransitionDeliverNotificationsAtStart

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasAutoTransitionDeliverNotificationsAtStart() bool`

HasAutoTransitionDeliverNotificationsAtStart returns a boolean if a field has been set.

### GetAutoTransitionDeliverNotificationsAtEnd

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTransitionDeliverNotificationsAtEnd() bool`

GetAutoTransitionDeliverNotificationsAtEnd returns the AutoTransitionDeliverNotificationsAtEnd field if non-nil, zero value otherwise.

### GetAutoTransitionDeliverNotificationsAtEndOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTransitionDeliverNotificationsAtEndOk() (*bool, bool)`

GetAutoTransitionDeliverNotificationsAtEndOk returns a tuple with the AutoTransitionDeliverNotificationsAtEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTransitionDeliverNotificationsAtEnd

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetAutoTransitionDeliverNotificationsAtEnd(v bool)`

SetAutoTransitionDeliverNotificationsAtEnd sets AutoTransitionDeliverNotificationsAtEnd field to given value.

### HasAutoTransitionDeliverNotificationsAtEnd

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasAutoTransitionDeliverNotificationsAtEnd() bool`

HasAutoTransitionDeliverNotificationsAtEnd returns a boolean if a field has been set.

### GetReminderIntervals

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetReminderIntervals() string`

GetReminderIntervals returns the ReminderIntervals field if non-nil, zero value otherwise.

### GetReminderIntervalsOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetReminderIntervalsOk() (*string, bool)`

GetReminderIntervalsOk returns a tuple with the ReminderIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderIntervals

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetReminderIntervals(v string)`

SetReminderIntervals sets ReminderIntervals field to given value.

### HasReminderIntervals

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasReminderIntervals() bool`

HasReminderIntervals returns a boolean if a field has been set.

### GetMetadata

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDeliverNotifications

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetDeliverNotifications() bool`

GetDeliverNotifications returns the DeliverNotifications field if non-nil, zero value otherwise.

### GetDeliverNotificationsOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetDeliverNotificationsOk() (*bool, bool)`

GetDeliverNotificationsOk returns a tuple with the DeliverNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverNotifications

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetDeliverNotifications(v bool)`

SetDeliverNotifications sets DeliverNotifications field to given value.

### HasDeliverNotifications

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasDeliverNotifications() bool`

HasDeliverNotifications returns a boolean if a field has been set.

### GetAutoTweetAtBeginning

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTweetAtBeginning() bool`

GetAutoTweetAtBeginning returns the AutoTweetAtBeginning field if non-nil, zero value otherwise.

### GetAutoTweetAtBeginningOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTweetAtBeginningOk() (*bool, bool)`

GetAutoTweetAtBeginningOk returns a tuple with the AutoTweetAtBeginning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetAtBeginning

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetAutoTweetAtBeginning(v bool)`

SetAutoTweetAtBeginning sets AutoTweetAtBeginning field to given value.

### HasAutoTweetAtBeginning

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasAutoTweetAtBeginning() bool`

HasAutoTweetAtBeginning returns a boolean if a field has been set.

### GetAutoTweetOnCompletion

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTweetOnCompletion() bool`

GetAutoTweetOnCompletion returns the AutoTweetOnCompletion field if non-nil, zero value otherwise.

### GetAutoTweetOnCompletionOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTweetOnCompletionOk() (*bool, bool)`

GetAutoTweetOnCompletionOk returns a tuple with the AutoTweetOnCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetOnCompletion

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetAutoTweetOnCompletion(v bool)`

SetAutoTweetOnCompletion sets AutoTweetOnCompletion field to given value.

### HasAutoTweetOnCompletion

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasAutoTweetOnCompletion() bool`

HasAutoTweetOnCompletion returns a boolean if a field has been set.

### GetAutoTweetOnCreation

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTweetOnCreation() bool`

GetAutoTweetOnCreation returns the AutoTweetOnCreation field if non-nil, zero value otherwise.

### GetAutoTweetOnCreationOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTweetOnCreationOk() (*bool, bool)`

GetAutoTweetOnCreationOk returns a tuple with the AutoTweetOnCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetOnCreation

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetAutoTweetOnCreation(v bool)`

SetAutoTweetOnCreation sets AutoTweetOnCreation field to given value.

### HasAutoTweetOnCreation

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasAutoTweetOnCreation() bool`

HasAutoTweetOnCreation returns a boolean if a field has been set.

### GetAutoTweetOneHourBefore

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTweetOneHourBefore() bool`

GetAutoTweetOneHourBefore returns the AutoTweetOneHourBefore field if non-nil, zero value otherwise.

### GetAutoTweetOneHourBeforeOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetAutoTweetOneHourBeforeOk() (*bool, bool)`

GetAutoTweetOneHourBeforeOk returns a tuple with the AutoTweetOneHourBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTweetOneHourBefore

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetAutoTweetOneHourBefore(v bool)`

SetAutoTweetOneHourBefore sets AutoTweetOneHourBefore field to given value.

### HasAutoTweetOneHourBefore

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasAutoTweetOneHourBefore() bool`

HasAutoTweetOneHourBefore returns a boolean if a field has been set.

### GetBackfillDate

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetBackfillDate() string`

GetBackfillDate returns the BackfillDate field if non-nil, zero value otherwise.

### GetBackfillDateOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetBackfillDateOk() (*string, bool)`

GetBackfillDateOk returns a tuple with the BackfillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillDate

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetBackfillDate(v string)`

SetBackfillDate sets BackfillDate field to given value.

### HasBackfillDate

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasBackfillDate() bool`

HasBackfillDate returns a boolean if a field has been set.

### GetBackfilled

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetBackfilled() bool`

GetBackfilled returns the Backfilled field if non-nil, zero value otherwise.

### GetBackfilledOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetBackfilledOk() (*bool, bool)`

GetBackfilledOk returns a tuple with the Backfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilled

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetBackfilled(v bool)`

SetBackfilled sets Backfilled field to given value.

### HasBackfilled

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasBackfilled() bool`

HasBackfilled returns a boolean if a field has been set.

### GetBody

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetComponents

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetComponents() map[string]string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetComponentsOk() (*map[string]string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetComponents(v map[string]string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetComponentIds

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetComponentIds() []string`

GetComponentIds returns the ComponentIds field if non-nil, zero value otherwise.

### GetComponentIdsOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetComponentIdsOk() (*[]string, bool)`

GetComponentIdsOk returns a tuple with the ComponentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIds

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetComponentIds(v []string)`

SetComponentIds sets ComponentIds field to given value.

### HasComponentIds

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasComponentIds() bool`

HasComponentIds returns a boolean if a field has been set.

### GetScheduledAutoTransition

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledAutoTransition() bool`

GetScheduledAutoTransition returns the ScheduledAutoTransition field if non-nil, zero value otherwise.

### GetScheduledAutoTransitionOk

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) GetScheduledAutoTransitionOk() (*bool, bool)`

GetScheduledAutoTransitionOk returns a tuple with the ScheduledAutoTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoTransition

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) SetScheduledAutoTransition(v bool)`

SetScheduledAutoTransition sets ScheduledAutoTransition field to given value.

### HasScheduledAutoTransition

`func (o *PutPagesPageIdIncidentsIncidentIdIncident) HasScheduledAutoTransition() bool`

HasScheduledAutoTransition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


