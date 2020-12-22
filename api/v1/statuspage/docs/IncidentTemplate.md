# IncidentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Incident Template Identifier | [optional] 
**Components** | Pointer to [**[]Component**](Component.md) | Affected components | [optional] 
**Name** | Pointer to **string** | Name of the template, as shown in the list on the \&quot;Templates\&quot; tab of the \&quot;Incidents\&quot; page | [optional] 
**Title** | Pointer to **string** | Title to be applied to the incident or maintenance when selecting this template | [optional] 
**Body** | Pointer to **string** | Body of the incident or maintenance update to be applied when selecting this template | [optional] 
**GroupId** | Pointer to **string** | Identifier of Template Group this template belongs to | [optional] 
**UpdateStatus** | Pointer to **string** | The status the incident or maintenance should transition to when selecting this template | [optional] 
**ShouldTweet** | Pointer to **bool** | Whether the \&quot;tweet update\&quot; checkbox should be selected when selecting this template | [optional] 
**ShouldSendNotifications** | Pointer to **bool** | Whether the \&quot;deliver notifications\&quot; checkbox should be selected when selecting this template | [optional] 

## Methods

### NewIncidentTemplate

`func NewIncidentTemplate() *IncidentTemplate`

NewIncidentTemplate instantiates a new IncidentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentTemplateWithDefaults

`func NewIncidentTemplateWithDefaults() *IncidentTemplate`

NewIncidentTemplateWithDefaults instantiates a new IncidentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IncidentTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetComponents

`func (o *IncidentTemplate) GetComponents() []Component`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *IncidentTemplate) GetComponentsOk() (*[]Component, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *IncidentTemplate) SetComponents(v []Component)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *IncidentTemplate) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetName

`func (o *IncidentTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *IncidentTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IncidentTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IncidentTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IncidentTemplate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *IncidentTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *IncidentTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *IncidentTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *IncidentTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetGroupId

`func (o *IncidentTemplate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *IncidentTemplate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *IncidentTemplate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *IncidentTemplate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUpdateStatus

`func (o *IncidentTemplate) GetUpdateStatus() string`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *IncidentTemplate) GetUpdateStatusOk() (*string, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *IncidentTemplate) SetUpdateStatus(v string)`

SetUpdateStatus sets UpdateStatus field to given value.

### HasUpdateStatus

`func (o *IncidentTemplate) HasUpdateStatus() bool`

HasUpdateStatus returns a boolean if a field has been set.

### GetShouldTweet

`func (o *IncidentTemplate) GetShouldTweet() bool`

GetShouldTweet returns the ShouldTweet field if non-nil, zero value otherwise.

### GetShouldTweetOk

`func (o *IncidentTemplate) GetShouldTweetOk() (*bool, bool)`

GetShouldTweetOk returns a tuple with the ShouldTweet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldTweet

`func (o *IncidentTemplate) SetShouldTweet(v bool)`

SetShouldTweet sets ShouldTweet field to given value.

### HasShouldTweet

`func (o *IncidentTemplate) HasShouldTweet() bool`

HasShouldTweet returns a boolean if a field has been set.

### GetShouldSendNotifications

`func (o *IncidentTemplate) GetShouldSendNotifications() bool`

GetShouldSendNotifications returns the ShouldSendNotifications field if non-nil, zero value otherwise.

### GetShouldSendNotificationsOk

`func (o *IncidentTemplate) GetShouldSendNotificationsOk() (*bool, bool)`

GetShouldSendNotificationsOk returns a tuple with the ShouldSendNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendNotifications

`func (o *IncidentTemplate) SetShouldSendNotifications(v bool)`

SetShouldSendNotifications sets ShouldSendNotifications field to given value.

### HasShouldSendNotifications

`func (o *IncidentTemplate) HasShouldSendNotifications() bool`

HasShouldSendNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


