# IncidentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Incident Update Identifier. | [optional] 
**IncidentId** | Pointer to **string** | Incident Identifier. | [optional] 
**AffectedComponents** | Pointer to **[]map[string]interface{}** | Affected components associated with the incident update. | [optional] 
**Body** | Pointer to **string** | Incident update body. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the incident update was created at. | [optional] 
**CustomTweet** | Pointer to **string** | An optional customized tweet message for incident postmortem. | [optional] 
**DeliverNotifications** | Pointer to **bool** | Controls whether to delivery notifications. | [optional] 
**DisplayAt** | Pointer to **time.Time** | Timestamp when incident update is happened. | [optional] 
**Status** | Pointer to **string** | The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed. | [optional] 
**TweetId** | Pointer to **string** | Tweet identifier associated to this incident update. | [optional] 
**TwitterUpdatedAt** | Pointer to **time.Time** | The timestamp when twitter updated at. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp when the incident update is updated. | [optional] 
**WantsTwitterUpdate** | Pointer to **bool** | Controls whether to create twitter update. | [optional] 

## Methods

### NewIncidentUpdate

`func NewIncidentUpdate() *IncidentUpdate`

NewIncidentUpdate instantiates a new IncidentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentUpdateWithDefaults

`func NewIncidentUpdateWithDefaults() *IncidentUpdate`

NewIncidentUpdateWithDefaults instantiates a new IncidentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IncidentUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncidentId

`func (o *IncidentUpdate) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *IncidentUpdate) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *IncidentUpdate) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *IncidentUpdate) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### GetAffectedComponents

`func (o *IncidentUpdate) GetAffectedComponents() []map[string]interface{}`

GetAffectedComponents returns the AffectedComponents field if non-nil, zero value otherwise.

### GetAffectedComponentsOk

`func (o *IncidentUpdate) GetAffectedComponentsOk() (*[]map[string]interface{}, bool)`

GetAffectedComponentsOk returns a tuple with the AffectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedComponents

`func (o *IncidentUpdate) SetAffectedComponents(v []map[string]interface{})`

SetAffectedComponents sets AffectedComponents field to given value.

### HasAffectedComponents

`func (o *IncidentUpdate) HasAffectedComponents() bool`

HasAffectedComponents returns a boolean if a field has been set.

### GetBody

`func (o *IncidentUpdate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *IncidentUpdate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *IncidentUpdate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *IncidentUpdate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IncidentUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IncidentUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IncidentUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IncidentUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomTweet

`func (o *IncidentUpdate) GetCustomTweet() string`

GetCustomTweet returns the CustomTweet field if non-nil, zero value otherwise.

### GetCustomTweetOk

`func (o *IncidentUpdate) GetCustomTweetOk() (*string, bool)`

GetCustomTweetOk returns a tuple with the CustomTweet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTweet

`func (o *IncidentUpdate) SetCustomTweet(v string)`

SetCustomTweet sets CustomTweet field to given value.

### HasCustomTweet

`func (o *IncidentUpdate) HasCustomTweet() bool`

HasCustomTweet returns a boolean if a field has been set.

### GetDeliverNotifications

`func (o *IncidentUpdate) GetDeliverNotifications() bool`

GetDeliverNotifications returns the DeliverNotifications field if non-nil, zero value otherwise.

### GetDeliverNotificationsOk

`func (o *IncidentUpdate) GetDeliverNotificationsOk() (*bool, bool)`

GetDeliverNotificationsOk returns a tuple with the DeliverNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverNotifications

`func (o *IncidentUpdate) SetDeliverNotifications(v bool)`

SetDeliverNotifications sets DeliverNotifications field to given value.

### HasDeliverNotifications

`func (o *IncidentUpdate) HasDeliverNotifications() bool`

HasDeliverNotifications returns a boolean if a field has been set.

### GetDisplayAt

`func (o *IncidentUpdate) GetDisplayAt() time.Time`

GetDisplayAt returns the DisplayAt field if non-nil, zero value otherwise.

### GetDisplayAtOk

`func (o *IncidentUpdate) GetDisplayAtOk() (*time.Time, bool)`

GetDisplayAtOk returns a tuple with the DisplayAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAt

`func (o *IncidentUpdate) SetDisplayAt(v time.Time)`

SetDisplayAt sets DisplayAt field to given value.

### HasDisplayAt

`func (o *IncidentUpdate) HasDisplayAt() bool`

HasDisplayAt returns a boolean if a field has been set.

### GetStatus

`func (o *IncidentUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncidentUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncidentUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IncidentUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTweetId

`func (o *IncidentUpdate) GetTweetId() string`

GetTweetId returns the TweetId field if non-nil, zero value otherwise.

### GetTweetIdOk

`func (o *IncidentUpdate) GetTweetIdOk() (*string, bool)`

GetTweetIdOk returns a tuple with the TweetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetId

`func (o *IncidentUpdate) SetTweetId(v string)`

SetTweetId sets TweetId field to given value.

### HasTweetId

`func (o *IncidentUpdate) HasTweetId() bool`

HasTweetId returns a boolean if a field has been set.

### GetTwitterUpdatedAt

`func (o *IncidentUpdate) GetTwitterUpdatedAt() time.Time`

GetTwitterUpdatedAt returns the TwitterUpdatedAt field if non-nil, zero value otherwise.

### GetTwitterUpdatedAtOk

`func (o *IncidentUpdate) GetTwitterUpdatedAtOk() (*time.Time, bool)`

GetTwitterUpdatedAtOk returns a tuple with the TwitterUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUpdatedAt

`func (o *IncidentUpdate) SetTwitterUpdatedAt(v time.Time)`

SetTwitterUpdatedAt sets TwitterUpdatedAt field to given value.

### HasTwitterUpdatedAt

`func (o *IncidentUpdate) HasTwitterUpdatedAt() bool`

HasTwitterUpdatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IncidentUpdate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IncidentUpdate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IncidentUpdate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IncidentUpdate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWantsTwitterUpdate

`func (o *IncidentUpdate) GetWantsTwitterUpdate() bool`

GetWantsTwitterUpdate returns the WantsTwitterUpdate field if non-nil, zero value otherwise.

### GetWantsTwitterUpdateOk

`func (o *IncidentUpdate) GetWantsTwitterUpdateOk() (*bool, bool)`

GetWantsTwitterUpdateOk returns a tuple with the WantsTwitterUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantsTwitterUpdate

`func (o *IncidentUpdate) SetWantsTwitterUpdate(v bool)`

SetWantsTwitterUpdate sets WantsTwitterUpdate field to given value.

### HasWantsTwitterUpdate

`func (o *IncidentUpdate) HasWantsTwitterUpdate() bool`

HasWantsTwitterUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


