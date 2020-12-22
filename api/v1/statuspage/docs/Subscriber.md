# Subscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Subscriber Identifier | [optional] 
**SkipConfirmationNotification** | Pointer to **bool** | If this is true, do not notify the user with changes to their subscription. | [optional] 
**Mode** | Pointer to **string** | The communication mode of the subscriber. | [optional] 
**Email** | Pointer to **string** | The email address to use to contact the subscriber. Used for Email and Webhook subscribers. | [optional] 
**Endpoint** | Pointer to **string** | The URL where a webhook subscriber elects to receive updates. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number used to contact an SMS subscriber | [optional] 
**PhoneCountry** | Pointer to **string** | The two-character country code representing the country of which the phone_number is a part. | [optional] 
**DisplayPhoneNumber** | Pointer to **string** | A formatted version of the phone_number and phone_country pair, nicely formatted for display. | [optional] 
**ObfuscatedChannelName** | Pointer to **string** | Obfuscated slack channel name | [optional] 
**WorkspaceName** | Pointer to **string** | The workspace name of the slack subscriber. | [optional] 
**QuarantinedAt** | Pointer to **time.Time** | The timestamp when the subscriber was quarantined due to an issue reaching them. | [optional] 
**PurgeAt** | Pointer to **time.Time** | The timestamp when a quarantined subscriber will be purged (unsubscribed). | [optional] 
**Components** | Pointer to **string** | The components for which the subscriber has elected to receive updates. | [optional] 
**PageAccessUserId** | Pointer to **string** | The Page Access user this subscriber belongs to (only for audience-specific pages). | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSubscriber

`func NewSubscriber() *Subscriber`

NewSubscriber instantiates a new Subscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberWithDefaults

`func NewSubscriberWithDefaults() *Subscriber`

NewSubscriberWithDefaults instantiates a new Subscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscriber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscriber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscriber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscriber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSkipConfirmationNotification

`func (o *Subscriber) GetSkipConfirmationNotification() bool`

GetSkipConfirmationNotification returns the SkipConfirmationNotification field if non-nil, zero value otherwise.

### GetSkipConfirmationNotificationOk

`func (o *Subscriber) GetSkipConfirmationNotificationOk() (*bool, bool)`

GetSkipConfirmationNotificationOk returns a tuple with the SkipConfirmationNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirmationNotification

`func (o *Subscriber) SetSkipConfirmationNotification(v bool)`

SetSkipConfirmationNotification sets SkipConfirmationNotification field to given value.

### HasSkipConfirmationNotification

`func (o *Subscriber) HasSkipConfirmationNotification() bool`

HasSkipConfirmationNotification returns a boolean if a field has been set.

### GetMode

`func (o *Subscriber) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Subscriber) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Subscriber) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Subscriber) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetEmail

`func (o *Subscriber) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Subscriber) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Subscriber) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Subscriber) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndpoint

`func (o *Subscriber) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Subscriber) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Subscriber) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Subscriber) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Subscriber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Subscriber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Subscriber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Subscriber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneCountry

`func (o *Subscriber) GetPhoneCountry() string`

GetPhoneCountry returns the PhoneCountry field if non-nil, zero value otherwise.

### GetPhoneCountryOk

`func (o *Subscriber) GetPhoneCountryOk() (*string, bool)`

GetPhoneCountryOk returns a tuple with the PhoneCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountry

`func (o *Subscriber) SetPhoneCountry(v string)`

SetPhoneCountry sets PhoneCountry field to given value.

### HasPhoneCountry

`func (o *Subscriber) HasPhoneCountry() bool`

HasPhoneCountry returns a boolean if a field has been set.

### GetDisplayPhoneNumber

`func (o *Subscriber) GetDisplayPhoneNumber() string`

GetDisplayPhoneNumber returns the DisplayPhoneNumber field if non-nil, zero value otherwise.

### GetDisplayPhoneNumberOk

`func (o *Subscriber) GetDisplayPhoneNumberOk() (*string, bool)`

GetDisplayPhoneNumberOk returns a tuple with the DisplayPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPhoneNumber

`func (o *Subscriber) SetDisplayPhoneNumber(v string)`

SetDisplayPhoneNumber sets DisplayPhoneNumber field to given value.

### HasDisplayPhoneNumber

`func (o *Subscriber) HasDisplayPhoneNumber() bool`

HasDisplayPhoneNumber returns a boolean if a field has been set.

### GetObfuscatedChannelName

`func (o *Subscriber) GetObfuscatedChannelName() string`

GetObfuscatedChannelName returns the ObfuscatedChannelName field if non-nil, zero value otherwise.

### GetObfuscatedChannelNameOk

`func (o *Subscriber) GetObfuscatedChannelNameOk() (*string, bool)`

GetObfuscatedChannelNameOk returns a tuple with the ObfuscatedChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedChannelName

`func (o *Subscriber) SetObfuscatedChannelName(v string)`

SetObfuscatedChannelName sets ObfuscatedChannelName field to given value.

### HasObfuscatedChannelName

`func (o *Subscriber) HasObfuscatedChannelName() bool`

HasObfuscatedChannelName returns a boolean if a field has been set.

### GetWorkspaceName

`func (o *Subscriber) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *Subscriber) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *Subscriber) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.

### HasWorkspaceName

`func (o *Subscriber) HasWorkspaceName() bool`

HasWorkspaceName returns a boolean if a field has been set.

### GetQuarantinedAt

`func (o *Subscriber) GetQuarantinedAt() time.Time`

GetQuarantinedAt returns the QuarantinedAt field if non-nil, zero value otherwise.

### GetQuarantinedAtOk

`func (o *Subscriber) GetQuarantinedAtOk() (*time.Time, bool)`

GetQuarantinedAtOk returns a tuple with the QuarantinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantinedAt

`func (o *Subscriber) SetQuarantinedAt(v time.Time)`

SetQuarantinedAt sets QuarantinedAt field to given value.

### HasQuarantinedAt

`func (o *Subscriber) HasQuarantinedAt() bool`

HasQuarantinedAt returns a boolean if a field has been set.

### GetPurgeAt

`func (o *Subscriber) GetPurgeAt() time.Time`

GetPurgeAt returns the PurgeAt field if non-nil, zero value otherwise.

### GetPurgeAtOk

`func (o *Subscriber) GetPurgeAtOk() (*time.Time, bool)`

GetPurgeAtOk returns a tuple with the PurgeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeAt

`func (o *Subscriber) SetPurgeAt(v time.Time)`

SetPurgeAt sets PurgeAt field to given value.

### HasPurgeAt

`func (o *Subscriber) HasPurgeAt() bool`

HasPurgeAt returns a boolean if a field has been set.

### GetComponents

`func (o *Subscriber) GetComponents() string`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Subscriber) GetComponentsOk() (*string, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *Subscriber) SetComponents(v string)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *Subscriber) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetPageAccessUserId

`func (o *Subscriber) GetPageAccessUserId() string`

GetPageAccessUserId returns the PageAccessUserId field if non-nil, zero value otherwise.

### GetPageAccessUserIdOk

`func (o *Subscriber) GetPageAccessUserIdOk() (*string, bool)`

GetPageAccessUserIdOk returns a tuple with the PageAccessUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageAccessUserId

`func (o *Subscriber) SetPageAccessUserId(v string)`

SetPageAccessUserId sets PageAccessUserId field to given value.

### HasPageAccessUserId

`func (o *Subscriber) HasPageAccessUserId() bool`

HasPageAccessUserId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscriber) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscriber) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscriber) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscriber) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


