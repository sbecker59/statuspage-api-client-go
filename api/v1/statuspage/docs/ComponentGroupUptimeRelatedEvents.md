# ComponentGroupUptimeRelatedEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | Pointer to **string** | Component identifier | [optional] 
**Incidents** | Pointer to [**ComponentUptimeRelatedEvents**](ComponentUptime_related_events.md) |  | [optional] 

## Methods

### NewComponentGroupUptimeRelatedEvents

`func NewComponentGroupUptimeRelatedEvents() *ComponentGroupUptimeRelatedEvents`

NewComponentGroupUptimeRelatedEvents instantiates a new ComponentGroupUptimeRelatedEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentGroupUptimeRelatedEventsWithDefaults

`func NewComponentGroupUptimeRelatedEventsWithDefaults() *ComponentGroupUptimeRelatedEvents`

NewComponentGroupUptimeRelatedEventsWithDefaults instantiates a new ComponentGroupUptimeRelatedEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentId

`func (o *ComponentGroupUptimeRelatedEvents) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *ComponentGroupUptimeRelatedEvents) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *ComponentGroupUptimeRelatedEvents) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *ComponentGroupUptimeRelatedEvents) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetIncidents

`func (o *ComponentGroupUptimeRelatedEvents) GetIncidents() ComponentUptimeRelatedEvents`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ComponentGroupUptimeRelatedEvents) GetIncidentsOk() (*ComponentUptimeRelatedEvents, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ComponentGroupUptimeRelatedEvents) SetIncidents(v ComponentUptimeRelatedEvents)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ComponentGroupUptimeRelatedEvents) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


