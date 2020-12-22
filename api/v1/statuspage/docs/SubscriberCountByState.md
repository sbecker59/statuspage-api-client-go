# SubscriberCountByState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | The number of active subscribers found by the query. | [optional] 
**Unconfirmed** | Pointer to **int32** | The number of unconfirmed subscribers found by the query. | [optional] 
**Quarantined** | Pointer to **int32** | The number of quarantined subscribers found by the query. | [optional] 
**Total** | Pointer to **int32** | The total number of subscribers found by the query. | [optional] 

## Methods

### NewSubscriberCountByState

`func NewSubscriberCountByState() *SubscriberCountByState`

NewSubscriberCountByState instantiates a new SubscriberCountByState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberCountByStateWithDefaults

`func NewSubscriberCountByStateWithDefaults() *SubscriberCountByState`

NewSubscriberCountByStateWithDefaults instantiates a new SubscriberCountByState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SubscriberCountByState) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SubscriberCountByState) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SubscriberCountByState) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *SubscriberCountByState) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUnconfirmed

`func (o *SubscriberCountByState) GetUnconfirmed() int32`

GetUnconfirmed returns the Unconfirmed field if non-nil, zero value otherwise.

### GetUnconfirmedOk

`func (o *SubscriberCountByState) GetUnconfirmedOk() (*int32, bool)`

GetUnconfirmedOk returns a tuple with the Unconfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmed

`func (o *SubscriberCountByState) SetUnconfirmed(v int32)`

SetUnconfirmed sets Unconfirmed field to given value.

### HasUnconfirmed

`func (o *SubscriberCountByState) HasUnconfirmed() bool`

HasUnconfirmed returns a boolean if a field has been set.

### GetQuarantined

`func (o *SubscriberCountByState) GetQuarantined() int32`

GetQuarantined returns the Quarantined field if non-nil, zero value otherwise.

### GetQuarantinedOk

`func (o *SubscriberCountByState) GetQuarantinedOk() (*int32, bool)`

GetQuarantinedOk returns a tuple with the Quarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantined

`func (o *SubscriberCountByState) SetQuarantined(v int32)`

SetQuarantined sets Quarantined field to given value.

### HasQuarantined

`func (o *SubscriberCountByState) HasQuarantined() bool`

HasQuarantined returns a boolean if a field has been set.

### GetTotal

`func (o *SubscriberCountByState) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SubscriberCountByState) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SubscriberCountByState) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SubscriberCountByState) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


