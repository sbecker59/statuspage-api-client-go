# PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address for creating Email subscribers. | [optional] 
**PhoneCountry** | Pointer to **string** | The two-character country where the phone number is located to use for the new SMS subscriber. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number (as you would dial from the phone_country) to use for the new SMS subscriber. | [optional] 
**SkipConfirmationNotification** | Pointer to **bool** | If skip_confirmation_notification is true, the subscriber does not receive any notifications when their subscription changes. Email subscribers will be automatically opted in. This option is only available for paid pages. This option has no effect for trial customers. | [optional] 

## Methods

### NewPostPagesPageIdIncidentsIncidentIdSubscribersSubscriber

`func NewPostPagesPageIdIncidentsIncidentIdSubscribersSubscriber() *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber`

NewPostPagesPageIdIncidentsIncidentIdSubscribersSubscriber instantiates a new PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPagesPageIdIncidentsIncidentIdSubscribersSubscriberWithDefaults

`func NewPostPagesPageIdIncidentsIncidentIdSubscribersSubscriberWithDefaults() *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber`

NewPostPagesPageIdIncidentsIncidentIdSubscribersSubscriberWithDefaults instantiates a new PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhoneCountry

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) GetPhoneCountry() string`

GetPhoneCountry returns the PhoneCountry field if non-nil, zero value otherwise.

### GetPhoneCountryOk

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) GetPhoneCountryOk() (*string, bool)`

GetPhoneCountryOk returns a tuple with the PhoneCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountry

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) SetPhoneCountry(v string)`

SetPhoneCountry sets PhoneCountry field to given value.

### HasPhoneCountry

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) HasPhoneCountry() bool`

HasPhoneCountry returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSkipConfirmationNotification

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) GetSkipConfirmationNotification() bool`

GetSkipConfirmationNotification returns the SkipConfirmationNotification field if non-nil, zero value otherwise.

### GetSkipConfirmationNotificationOk

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) GetSkipConfirmationNotificationOk() (*bool, bool)`

GetSkipConfirmationNotificationOk returns a tuple with the SkipConfirmationNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirmationNotification

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) SetSkipConfirmationNotification(v bool)`

SetSkipConfirmationNotification sets SkipConfirmationNotification field to given value.

### HasSkipConfirmationNotification

`func (o *PostPagesPageIdIncidentsIncidentIdSubscribersSubscriber) HasSkipConfirmationNotification() bool`

HasSkipConfirmationNotification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


