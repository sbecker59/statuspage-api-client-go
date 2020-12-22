# PostPagesPageIdSubscribersUnsubscribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscribers** | **string** | The array of subscriber codes to unsubscribe (limited to 100), or \&quot;all\&quot; to unsubscribe all subscribers if the number of subscribers is less than 100. | 
**Type** | Pointer to **string** | If this is present, only unsubscribe subscribers of this type. | [optional] 
**State** | Pointer to **string** | If this is present, only unsubscribe subscribers in this state. Specify state \&quot;all\&quot; to unsubscribe subscribers in any states. | [optional] [default to "active"]
**SkipUnsubscriptionNotification** | Pointer to **bool** | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. | [optional] 

## Methods

### NewPostPagesPageIdSubscribersUnsubscribe

`func NewPostPagesPageIdSubscribersUnsubscribe(subscribers string, ) *PostPagesPageIdSubscribersUnsubscribe`

NewPostPagesPageIdSubscribersUnsubscribe instantiates a new PostPagesPageIdSubscribersUnsubscribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPagesPageIdSubscribersUnsubscribeWithDefaults

`func NewPostPagesPageIdSubscribersUnsubscribeWithDefaults() *PostPagesPageIdSubscribersUnsubscribe`

NewPostPagesPageIdSubscribersUnsubscribeWithDefaults instantiates a new PostPagesPageIdSubscribersUnsubscribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribers

`func (o *PostPagesPageIdSubscribersUnsubscribe) GetSubscribers() string`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *PostPagesPageIdSubscribersUnsubscribe) GetSubscribersOk() (*string, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *PostPagesPageIdSubscribersUnsubscribe) SetSubscribers(v string)`

SetSubscribers sets Subscribers field to given value.


### GetType

`func (o *PostPagesPageIdSubscribersUnsubscribe) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostPagesPageIdSubscribersUnsubscribe) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostPagesPageIdSubscribersUnsubscribe) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostPagesPageIdSubscribersUnsubscribe) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *PostPagesPageIdSubscribersUnsubscribe) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostPagesPageIdSubscribersUnsubscribe) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostPagesPageIdSubscribersUnsubscribe) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PostPagesPageIdSubscribersUnsubscribe) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSkipUnsubscriptionNotification

`func (o *PostPagesPageIdSubscribersUnsubscribe) GetSkipUnsubscriptionNotification() bool`

GetSkipUnsubscriptionNotification returns the SkipUnsubscriptionNotification field if non-nil, zero value otherwise.

### GetSkipUnsubscriptionNotificationOk

`func (o *PostPagesPageIdSubscribersUnsubscribe) GetSkipUnsubscriptionNotificationOk() (*bool, bool)`

GetSkipUnsubscriptionNotificationOk returns a tuple with the SkipUnsubscriptionNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipUnsubscriptionNotification

`func (o *PostPagesPageIdSubscribersUnsubscribe) SetSkipUnsubscriptionNotification(v bool)`

SetSkipUnsubscriptionNotification sets SkipUnsubscriptionNotification field to given value.

### HasSkipUnsubscriptionNotification

`func (o *PostPagesPageIdSubscribersUnsubscribe) HasSkipUnsubscriptionNotification() bool`

HasSkipUnsubscriptionNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


