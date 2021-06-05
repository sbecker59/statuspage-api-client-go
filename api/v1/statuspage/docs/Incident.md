# Incident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Incident Identifier | [optional] 
**Components** | Pointer to [**[]Component**](Component.md) | Incident components | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the incident was created at. | [optional] 
**Impact** | Pointer to **string** | The impact of the incident. | [optional] 
**ImpactOverride** | Pointer to **string** | value to override calculated impact value | [optional] 
**IncidentUpdates** | Pointer to [**[]IncidentUpdate**](IncidentUpdate.md) | The incident updates for incident. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata attached to the incident. Top level values must be objects. | [optional] 
**MonitoringAt** | Pointer to **time.Time** | The timestamp when incident entered monitoring state. | [optional] 
**Name** | Pointer to **string** | Incident Name | [optional] 
**PageId** | Pointer to **string** | Incident Page Identifier | [optional] 
**PostmortemBody** | Pointer to **string** | Body of the Postmortem. | [optional] 
**PostmortemBodyLastUpdatedAt** | Pointer to **time.Time** | The timestamp when the incident postmortem body was last updated at. | [optional] 
**PostmortemIgnored** | Pointer to **bool** | Controls whether the incident will have postmortem. | [optional] 
**PostmortemNotifiedSubscribers** | Pointer to **bool** | Indicates whether subscribers are already notificed about postmortem. | [optional] 
**PostmortemNotifiedTwitter** | Pointer to **bool** | Controls whether to decide if notify postmortem on twitter. | [optional] 
**PostmortemPublishedAt** | Pointer to **bool** | The timestamp when the postmortem was published. | [optional] 
**ResolvedAt** | Pointer to **time.Time** | The timestamp when incident was resolved. | [optional] 
**ScheduledAutoCompleted** | Pointer to **bool** | Controls whether the incident is scheduled to automatically change to complete. | [optional] 
**ScheduledAutoInProgress** | Pointer to **bool** | Controls whether the incident is scheduled to automatically change to in progress. | [optional] 
**ScheduledFor** | Pointer to **time.Time** | The timestamp the incident is scheduled for. | [optional] 
**ScheduledRemindPrior** | Pointer to **bool** | Controls whether to remind subscribers prior to scheduled incidents. | [optional] 
**ScheduledRemindedAt** | Pointer to **time.Time** | The timestamp when the scheduled incident reminder was sent at. | [optional] 
**ScheduledUntil** | Pointer to **time.Time** | The timestamp the incident is scheduled until. | [optional] 
**Shortlink** | Pointer to **string** | Incident Shortlink. | [optional] 
**Status** | Pointer to **string** | The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp when the incident was updated at. | [optional] 

## Methods

### NewIncident

`func NewIncident() *Incident`

NewIncident instantiates a new Incident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentWithDefaults

`func NewIncidentWithDefaults() *Incident`

NewIncidentWithDefaults instantiates a new Incident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Incident) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Incident) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Incident) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Incident) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComponents

`func (o *Incident) GetComponents() []Component`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Incident) GetComponentsOk() (*[]Component, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *Incident) SetComponents(v []Component)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *Incident) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Incident) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Incident) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Incident) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Incident) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetImpact

`func (o *Incident) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *Incident) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *Incident) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *Incident) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetImpactOverride

`func (o *Incident) GetImpactOverride() string`

GetImpactOverride returns the ImpactOverride field if non-nil, zero value otherwise.

### GetImpactOverrideOk

`func (o *Incident) GetImpactOverrideOk() (*string, bool)`

GetImpactOverrideOk returns a tuple with the ImpactOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactOverride

`func (o *Incident) SetImpactOverride(v string)`

SetImpactOverride sets ImpactOverride field to given value.

### HasImpactOverride

`func (o *Incident) HasImpactOverride() bool`

HasImpactOverride returns a boolean if a field has been set.

### GetIncidentUpdates

`func (o *Incident) GetIncidentUpdates() []IncidentUpdate`

GetIncidentUpdates returns the IncidentUpdates field if non-nil, zero value otherwise.

### GetIncidentUpdatesOk

`func (o *Incident) GetIncidentUpdatesOk() (*[]IncidentUpdate, bool)`

GetIncidentUpdatesOk returns a tuple with the IncidentUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentUpdates

`func (o *Incident) SetIncidentUpdates(v []IncidentUpdate)`

SetIncidentUpdates sets IncidentUpdates field to given value.

### HasIncidentUpdates

`func (o *Incident) HasIncidentUpdates() bool`

HasIncidentUpdates returns a boolean if a field has been set.

### GetMetadata

`func (o *Incident) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Incident) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Incident) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Incident) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMonitoringAt

`func (o *Incident) GetMonitoringAt() time.Time`

GetMonitoringAt returns the MonitoringAt field if non-nil, zero value otherwise.

### GetMonitoringAtOk

`func (o *Incident) GetMonitoringAtOk() (*time.Time, bool)`

GetMonitoringAtOk returns a tuple with the MonitoringAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringAt

`func (o *Incident) SetMonitoringAt(v time.Time)`

SetMonitoringAt sets MonitoringAt field to given value.

### HasMonitoringAt

`func (o *Incident) HasMonitoringAt() bool`

HasMonitoringAt returns a boolean if a field has been set.

### GetName

`func (o *Incident) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Incident) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Incident) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Incident) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPageId

`func (o *Incident) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *Incident) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *Incident) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *Incident) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetPostmortemBody

`func (o *Incident) GetPostmortemBody() string`

GetPostmortemBody returns the PostmortemBody field if non-nil, zero value otherwise.

### GetPostmortemBodyOk

`func (o *Incident) GetPostmortemBodyOk() (*string, bool)`

GetPostmortemBodyOk returns a tuple with the PostmortemBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostmortemBody

`func (o *Incident) SetPostmortemBody(v string)`

SetPostmortemBody sets PostmortemBody field to given value.

### HasPostmortemBody

`func (o *Incident) HasPostmortemBody() bool`

HasPostmortemBody returns a boolean if a field has been set.

### GetPostmortemBodyLastUpdatedAt

`func (o *Incident) GetPostmortemBodyLastUpdatedAt() time.Time`

GetPostmortemBodyLastUpdatedAt returns the PostmortemBodyLastUpdatedAt field if non-nil, zero value otherwise.

### GetPostmortemBodyLastUpdatedAtOk

`func (o *Incident) GetPostmortemBodyLastUpdatedAtOk() (*time.Time, bool)`

GetPostmortemBodyLastUpdatedAtOk returns a tuple with the PostmortemBodyLastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostmortemBodyLastUpdatedAt

`func (o *Incident) SetPostmortemBodyLastUpdatedAt(v time.Time)`

SetPostmortemBodyLastUpdatedAt sets PostmortemBodyLastUpdatedAt field to given value.

### HasPostmortemBodyLastUpdatedAt

`func (o *Incident) HasPostmortemBodyLastUpdatedAt() bool`

HasPostmortemBodyLastUpdatedAt returns a boolean if a field has been set.

### GetPostmortemIgnored

`func (o *Incident) GetPostmortemIgnored() bool`

GetPostmortemIgnored returns the PostmortemIgnored field if non-nil, zero value otherwise.

### GetPostmortemIgnoredOk

`func (o *Incident) GetPostmortemIgnoredOk() (*bool, bool)`

GetPostmortemIgnoredOk returns a tuple with the PostmortemIgnored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostmortemIgnored

`func (o *Incident) SetPostmortemIgnored(v bool)`

SetPostmortemIgnored sets PostmortemIgnored field to given value.

### HasPostmortemIgnored

`func (o *Incident) HasPostmortemIgnored() bool`

HasPostmortemIgnored returns a boolean if a field has been set.

### GetPostmortemNotifiedSubscribers

`func (o *Incident) GetPostmortemNotifiedSubscribers() bool`

GetPostmortemNotifiedSubscribers returns the PostmortemNotifiedSubscribers field if non-nil, zero value otherwise.

### GetPostmortemNotifiedSubscribersOk

`func (o *Incident) GetPostmortemNotifiedSubscribersOk() (*bool, bool)`

GetPostmortemNotifiedSubscribersOk returns a tuple with the PostmortemNotifiedSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostmortemNotifiedSubscribers

`func (o *Incident) SetPostmortemNotifiedSubscribers(v bool)`

SetPostmortemNotifiedSubscribers sets PostmortemNotifiedSubscribers field to given value.

### HasPostmortemNotifiedSubscribers

`func (o *Incident) HasPostmortemNotifiedSubscribers() bool`

HasPostmortemNotifiedSubscribers returns a boolean if a field has been set.

### GetPostmortemNotifiedTwitter

`func (o *Incident) GetPostmortemNotifiedTwitter() bool`

GetPostmortemNotifiedTwitter returns the PostmortemNotifiedTwitter field if non-nil, zero value otherwise.

### GetPostmortemNotifiedTwitterOk

`func (o *Incident) GetPostmortemNotifiedTwitterOk() (*bool, bool)`

GetPostmortemNotifiedTwitterOk returns a tuple with the PostmortemNotifiedTwitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostmortemNotifiedTwitter

`func (o *Incident) SetPostmortemNotifiedTwitter(v bool)`

SetPostmortemNotifiedTwitter sets PostmortemNotifiedTwitter field to given value.

### HasPostmortemNotifiedTwitter

`func (o *Incident) HasPostmortemNotifiedTwitter() bool`

HasPostmortemNotifiedTwitter returns a boolean if a field has been set.

### GetPostmortemPublishedAt

`func (o *Incident) GetPostmortemPublishedAt() bool`

GetPostmortemPublishedAt returns the PostmortemPublishedAt field if non-nil, zero value otherwise.

### GetPostmortemPublishedAtOk

`func (o *Incident) GetPostmortemPublishedAtOk() (*bool, bool)`

GetPostmortemPublishedAtOk returns a tuple with the PostmortemPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostmortemPublishedAt

`func (o *Incident) SetPostmortemPublishedAt(v bool)`

SetPostmortemPublishedAt sets PostmortemPublishedAt field to given value.

### HasPostmortemPublishedAt

`func (o *Incident) HasPostmortemPublishedAt() bool`

HasPostmortemPublishedAt returns a boolean if a field has been set.

### GetResolvedAt

`func (o *Incident) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *Incident) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *Incident) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *Incident) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetScheduledAutoCompleted

`func (o *Incident) GetScheduledAutoCompleted() bool`

GetScheduledAutoCompleted returns the ScheduledAutoCompleted field if non-nil, zero value otherwise.

### GetScheduledAutoCompletedOk

`func (o *Incident) GetScheduledAutoCompletedOk() (*bool, bool)`

GetScheduledAutoCompletedOk returns a tuple with the ScheduledAutoCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoCompleted

`func (o *Incident) SetScheduledAutoCompleted(v bool)`

SetScheduledAutoCompleted sets ScheduledAutoCompleted field to given value.

### HasScheduledAutoCompleted

`func (o *Incident) HasScheduledAutoCompleted() bool`

HasScheduledAutoCompleted returns a boolean if a field has been set.

### GetScheduledAutoInProgress

`func (o *Incident) GetScheduledAutoInProgress() bool`

GetScheduledAutoInProgress returns the ScheduledAutoInProgress field if non-nil, zero value otherwise.

### GetScheduledAutoInProgressOk

`func (o *Incident) GetScheduledAutoInProgressOk() (*bool, bool)`

GetScheduledAutoInProgressOk returns a tuple with the ScheduledAutoInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAutoInProgress

`func (o *Incident) SetScheduledAutoInProgress(v bool)`

SetScheduledAutoInProgress sets ScheduledAutoInProgress field to given value.

### HasScheduledAutoInProgress

`func (o *Incident) HasScheduledAutoInProgress() bool`

HasScheduledAutoInProgress returns a boolean if a field has been set.

### GetScheduledFor

`func (o *Incident) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *Incident) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *Incident) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *Incident) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetScheduledRemindPrior

`func (o *Incident) GetScheduledRemindPrior() bool`

GetScheduledRemindPrior returns the ScheduledRemindPrior field if non-nil, zero value otherwise.

### GetScheduledRemindPriorOk

`func (o *Incident) GetScheduledRemindPriorOk() (*bool, bool)`

GetScheduledRemindPriorOk returns a tuple with the ScheduledRemindPrior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRemindPrior

`func (o *Incident) SetScheduledRemindPrior(v bool)`

SetScheduledRemindPrior sets ScheduledRemindPrior field to given value.

### HasScheduledRemindPrior

`func (o *Incident) HasScheduledRemindPrior() bool`

HasScheduledRemindPrior returns a boolean if a field has been set.

### GetScheduledRemindedAt

`func (o *Incident) GetScheduledRemindedAt() time.Time`

GetScheduledRemindedAt returns the ScheduledRemindedAt field if non-nil, zero value otherwise.

### GetScheduledRemindedAtOk

`func (o *Incident) GetScheduledRemindedAtOk() (*time.Time, bool)`

GetScheduledRemindedAtOk returns a tuple with the ScheduledRemindedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRemindedAt

`func (o *Incident) SetScheduledRemindedAt(v time.Time)`

SetScheduledRemindedAt sets ScheduledRemindedAt field to given value.

### HasScheduledRemindedAt

`func (o *Incident) HasScheduledRemindedAt() bool`

HasScheduledRemindedAt returns a boolean if a field has been set.

### GetScheduledUntil

`func (o *Incident) GetScheduledUntil() time.Time`

GetScheduledUntil returns the ScheduledUntil field if non-nil, zero value otherwise.

### GetScheduledUntilOk

`func (o *Incident) GetScheduledUntilOk() (*time.Time, bool)`

GetScheduledUntilOk returns a tuple with the ScheduledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledUntil

`func (o *Incident) SetScheduledUntil(v time.Time)`

SetScheduledUntil sets ScheduledUntil field to given value.

### HasScheduledUntil

`func (o *Incident) HasScheduledUntil() bool`

HasScheduledUntil returns a boolean if a field has been set.

### GetShortlink

`func (o *Incident) GetShortlink() string`

GetShortlink returns the Shortlink field if non-nil, zero value otherwise.

### GetShortlinkOk

`func (o *Incident) GetShortlinkOk() (*string, bool)`

GetShortlinkOk returns a tuple with the Shortlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortlink

`func (o *Incident) SetShortlink(v string)`

SetShortlink sets Shortlink field to given value.

### HasShortlink

`func (o *Incident) HasShortlink() bool`

HasShortlink returns a boolean if a field has been set.

### GetStatus

`func (o *Incident) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Incident) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Incident) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Incident) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Incident) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Incident) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Incident) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Incident) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


