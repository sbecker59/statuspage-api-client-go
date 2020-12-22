# ComponentGroupUptime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RangeStart** | Pointer to **time.Time** | Start date used for uptime calculation (see the warnings field in the response if this value does not match the start parameter you provided). | [optional] 
**RangeEnd** | Pointer to **time.Time** | End date used for uptime calculation (see the warnings field in the response if this value does not match the end parameter you provided). | [optional] 
**UptimePercentage** | Pointer to **float32** | Uptime percentage for a component | [optional] 
**MajorOutage** | Pointer to **int32** | Seconds of major outage | [optional] 
**PartialOutage** | Pointer to **int32** | Seconds of partial outage | [optional] 
**Warnings** | Pointer to **string** | Warning messages related to the uptime query that may occur | [optional] 
**Id** | Pointer to **string** | Component group identifier | [optional] 
**Name** | Pointer to **string** | Component group display name | [optional] 
**RelatedEvents** | Pointer to [**ComponentGroupUptimeRelatedEvents**](ComponentGroupUptime_related_events.md) |  | [optional] 

## Methods

### NewComponentGroupUptime

`func NewComponentGroupUptime() *ComponentGroupUptime`

NewComponentGroupUptime instantiates a new ComponentGroupUptime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentGroupUptimeWithDefaults

`func NewComponentGroupUptimeWithDefaults() *ComponentGroupUptime`

NewComponentGroupUptimeWithDefaults instantiates a new ComponentGroupUptime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRangeStart

`func (o *ComponentGroupUptime) GetRangeStart() time.Time`

GetRangeStart returns the RangeStart field if non-nil, zero value otherwise.

### GetRangeStartOk

`func (o *ComponentGroupUptime) GetRangeStartOk() (*time.Time, bool)`

GetRangeStartOk returns a tuple with the RangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStart

`func (o *ComponentGroupUptime) SetRangeStart(v time.Time)`

SetRangeStart sets RangeStart field to given value.

### HasRangeStart

`func (o *ComponentGroupUptime) HasRangeStart() bool`

HasRangeStart returns a boolean if a field has been set.

### GetRangeEnd

`func (o *ComponentGroupUptime) GetRangeEnd() time.Time`

GetRangeEnd returns the RangeEnd field if non-nil, zero value otherwise.

### GetRangeEndOk

`func (o *ComponentGroupUptime) GetRangeEndOk() (*time.Time, bool)`

GetRangeEndOk returns a tuple with the RangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEnd

`func (o *ComponentGroupUptime) SetRangeEnd(v time.Time)`

SetRangeEnd sets RangeEnd field to given value.

### HasRangeEnd

`func (o *ComponentGroupUptime) HasRangeEnd() bool`

HasRangeEnd returns a boolean if a field has been set.

### GetUptimePercentage

`func (o *ComponentGroupUptime) GetUptimePercentage() float32`

GetUptimePercentage returns the UptimePercentage field if non-nil, zero value otherwise.

### GetUptimePercentageOk

`func (o *ComponentGroupUptime) GetUptimePercentageOk() (*float32, bool)`

GetUptimePercentageOk returns a tuple with the UptimePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimePercentage

`func (o *ComponentGroupUptime) SetUptimePercentage(v float32)`

SetUptimePercentage sets UptimePercentage field to given value.

### HasUptimePercentage

`func (o *ComponentGroupUptime) HasUptimePercentage() bool`

HasUptimePercentage returns a boolean if a field has been set.

### GetMajorOutage

`func (o *ComponentGroupUptime) GetMajorOutage() int32`

GetMajorOutage returns the MajorOutage field if non-nil, zero value otherwise.

### GetMajorOutageOk

`func (o *ComponentGroupUptime) GetMajorOutageOk() (*int32, bool)`

GetMajorOutageOk returns a tuple with the MajorOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorOutage

`func (o *ComponentGroupUptime) SetMajorOutage(v int32)`

SetMajorOutage sets MajorOutage field to given value.

### HasMajorOutage

`func (o *ComponentGroupUptime) HasMajorOutage() bool`

HasMajorOutage returns a boolean if a field has been set.

### GetPartialOutage

`func (o *ComponentGroupUptime) GetPartialOutage() int32`

GetPartialOutage returns the PartialOutage field if non-nil, zero value otherwise.

### GetPartialOutageOk

`func (o *ComponentGroupUptime) GetPartialOutageOk() (*int32, bool)`

GetPartialOutageOk returns a tuple with the PartialOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialOutage

`func (o *ComponentGroupUptime) SetPartialOutage(v int32)`

SetPartialOutage sets PartialOutage field to given value.

### HasPartialOutage

`func (o *ComponentGroupUptime) HasPartialOutage() bool`

HasPartialOutage returns a boolean if a field has been set.

### GetWarnings

`func (o *ComponentGroupUptime) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ComponentGroupUptime) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ComponentGroupUptime) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ComponentGroupUptime) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetId

`func (o *ComponentGroupUptime) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentGroupUptime) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentGroupUptime) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentGroupUptime) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComponentGroupUptime) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentGroupUptime) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentGroupUptime) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentGroupUptime) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelatedEvents

`func (o *ComponentGroupUptime) GetRelatedEvents() ComponentGroupUptimeRelatedEvents`

GetRelatedEvents returns the RelatedEvents field if non-nil, zero value otherwise.

### GetRelatedEventsOk

`func (o *ComponentGroupUptime) GetRelatedEventsOk() (*ComponentGroupUptimeRelatedEvents, bool)`

GetRelatedEventsOk returns a tuple with the RelatedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEvents

`func (o *ComponentGroupUptime) SetRelatedEvents(v ComponentGroupUptimeRelatedEvents)`

SetRelatedEvents sets RelatedEvents field to given value.

### HasRelatedEvents

`func (o *ComponentGroupUptime) HasRelatedEvents() bool`

HasRelatedEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


