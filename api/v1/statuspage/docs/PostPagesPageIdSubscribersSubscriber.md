# PostPagesPageIdSubscribersSubscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address for creating Email and Webhook subscribers. | [optional] 
**Endpoint** | Pointer to **string** | The endpoint URI for creating Webhook subscribers. | [optional] 
**PhoneCountry** | Pointer to **string** | The two-character country where the phone number is located to use for the new SMS subscriber. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number (as you would dial from the phone_country) to use for the new SMS subscriber. | [optional] 
**SkipConfirmationNotification** | Pointer to **bool** | If skip_confirmation_notification is true, the subscriber does not receive any notifications when their subscription changes.  Email subscribers will be automatically opted in. This option is only available for paid pages. This option has no effect for trial customers. | [optional] 
**PageAccessUser** | Pointer to **string** | The code of the page access user to which the subscriber belongs. | [optional] 
**ComponentIds** | Pointer to **[]string** | A list of component ids for which the subscriber should recieve updates for. Components must be an array with at least one element if it is passed at all. Each component must belong to the page indicated in the path. | [optional] 

## Methods

### NewPostPagesPageIdSubscribersSubscriber

`func NewPostPagesPageIdSubscribersSubscriber() *PostPagesPageIdSubscribersSubscriber`

NewPostPagesPageIdSubscribersSubscriber instantiates a new PostPagesPageIdSubscribersSubscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPagesPageIdSubscribersSubscriberWithDefaults

`func NewPostPagesPageIdSubscribersSubscriberWithDefaults() *PostPagesPageIdSubscribersSubscriber`

NewPostPagesPageIdSubscribersSubscriberWithDefaults instantiates a new PostPagesPageIdSubscribersSubscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostPagesPageIdSubscribersSubscriber) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostPagesPageIdSubscribersSubscriber) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostPagesPageIdSubscribersSubscriber) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PostPagesPageIdSubscribersSubscriber) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndpoint

`func (o *PostPagesPageIdSubscribersSubscriber) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PostPagesPageIdSubscribersSubscriber) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PostPagesPageIdSubscribersSubscriber) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PostPagesPageIdSubscribersSubscriber) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetPhoneCountry

`func (o *PostPagesPageIdSubscribersSubscriber) GetPhoneCountry() string`

GetPhoneCountry returns the PhoneCountry field if non-nil, zero value otherwise.

### GetPhoneCountryOk

`func (o *PostPagesPageIdSubscribersSubscriber) GetPhoneCountryOk() (*string, bool)`

GetPhoneCountryOk returns a tuple with the PhoneCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountry

`func (o *PostPagesPageIdSubscribersSubscriber) SetPhoneCountry(v string)`

SetPhoneCountry sets PhoneCountry field to given value.

### HasPhoneCountry

`func (o *PostPagesPageIdSubscribersSubscriber) HasPhoneCountry() bool`

HasPhoneCountry returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PostPagesPageIdSubscribersSubscriber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PostPagesPageIdSubscribersSubscriber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PostPagesPageIdSubscribersSubscriber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PostPagesPageIdSubscribersSubscriber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSkipConfirmationNotification

`func (o *PostPagesPageIdSubscribersSubscriber) GetSkipConfirmationNotification() bool`

GetSkipConfirmationNotification returns the SkipConfirmationNotification field if non-nil, zero value otherwise.

### GetSkipConfirmationNotificationOk

`func (o *PostPagesPageIdSubscribersSubscriber) GetSkipConfirmationNotificationOk() (*bool, bool)`

GetSkipConfirmationNotificationOk returns a tuple with the SkipConfirmationNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirmationNotification

`func (o *PostPagesPageIdSubscribersSubscriber) SetSkipConfirmationNotification(v bool)`

SetSkipConfirmationNotification sets SkipConfirmationNotification field to given value.

### HasSkipConfirmationNotification

`func (o *PostPagesPageIdSubscribersSubscriber) HasSkipConfirmationNotification() bool`

HasSkipConfirmationNotification returns a boolean if a field has been set.

### GetPageAccessUser

`func (o *PostPagesPageIdSubscribersSubscriber) GetPageAccessUser() string`

GetPageAccessUser returns the PageAccessUser field if non-nil, zero value otherwise.

### GetPageAccessUserOk

`func (o *PostPagesPageIdSubscribersSubscriber) GetPageAccessUserOk() (*string, bool)`

GetPageAccessUserOk returns a tuple with the PageAccessUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageAccessUser

`func (o *PostPagesPageIdSubscribersSubscriber) SetPageAccessUser(v string)`

SetPageAccessUser sets PageAccessUser field to given value.

### HasPageAccessUser

`func (o *PostPagesPageIdSubscribersSubscriber) HasPageAccessUser() bool`

HasPageAccessUser returns a boolean if a field has been set.

### GetComponentIds

`func (o *PostPagesPageIdSubscribersSubscriber) GetComponentIds() []string`

GetComponentIds returns the ComponentIds field if non-nil, zero value otherwise.

### GetComponentIdsOk

`func (o *PostPagesPageIdSubscribersSubscriber) GetComponentIdsOk() (*[]string, bool)`

GetComponentIdsOk returns a tuple with the ComponentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIds

`func (o *PostPagesPageIdSubscribersSubscriber) SetComponentIds(v []string)`

SetComponentIds sets ComponentIds field to given value.

### HasComponentIds

`func (o *PostPagesPageIdSubscribersSubscriber) HasComponentIds() bool`

HasComponentIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


