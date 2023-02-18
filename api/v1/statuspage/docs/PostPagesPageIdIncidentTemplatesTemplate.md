# PostPagesPageIdIncidentTemplatesTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the template, as shown in the list on the \&quot;Templates\&quot; tab of the \&quot;Incidents\&quot; page | 
**Title** | **string** | Title to be applied to the incident or maintenance when selecting this template | 
**Body** | **string** | The initial message, created as the first incident or maintenance update. | 
**GroupId** | Pointer to **string** | Identifier of Template Group this template belongs to | [optional] 
**UpdateStatus** | Pointer to **string** | The status the incident or maintenance should transition to when selecting this template | [optional] 
**ShouldTweet** | Pointer to **bool** | Whether the \&quot;tweet update\&quot; checkbox should be selected when selecting this template | [optional] 
**ShouldSendNotifications** | Pointer to **bool** | Whether the \&quot;deliver notifications\&quot; checkbox should be selected when selecting this template | [optional] 
**ComponentIds** | Pointer to **[]string** | List of component_ids affected by this incident | [optional] 

## Methods

### NewPostPagesPageIdIncidentTemplatesTemplate

`func NewPostPagesPageIdIncidentTemplatesTemplate(name string, title string, body string, ) *PostPagesPageIdIncidentTemplatesTemplate`

NewPostPagesPageIdIncidentTemplatesTemplate instantiates a new PostPagesPageIdIncidentTemplatesTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPagesPageIdIncidentTemplatesTemplateWithDefaults

`func NewPostPagesPageIdIncidentTemplatesTemplateWithDefaults() *PostPagesPageIdIncidentTemplatesTemplate`

NewPostPagesPageIdIncidentTemplatesTemplateWithDefaults instantiates a new PostPagesPageIdIncidentTemplatesTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostPagesPageIdIncidentTemplatesTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostPagesPageIdIncidentTemplatesTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PostPagesPageIdIncidentTemplatesTemplate) SetBody(v string)`

SetBody sets Body field to given value.


### GetGroupId

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PostPagesPageIdIncidentTemplatesTemplate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PostPagesPageIdIncidentTemplatesTemplate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUpdateStatus

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetUpdateStatus() string`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetUpdateStatusOk() (*string, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *PostPagesPageIdIncidentTemplatesTemplate) SetUpdateStatus(v string)`

SetUpdateStatus sets UpdateStatus field to given value.

### HasUpdateStatus

`func (o *PostPagesPageIdIncidentTemplatesTemplate) HasUpdateStatus() bool`

HasUpdateStatus returns a boolean if a field has been set.

### GetShouldTweet

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetShouldTweet() bool`

GetShouldTweet returns the ShouldTweet field if non-nil, zero value otherwise.

### GetShouldTweetOk

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetShouldTweetOk() (*bool, bool)`

GetShouldTweetOk returns a tuple with the ShouldTweet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldTweet

`func (o *PostPagesPageIdIncidentTemplatesTemplate) SetShouldTweet(v bool)`

SetShouldTweet sets ShouldTweet field to given value.

### HasShouldTweet

`func (o *PostPagesPageIdIncidentTemplatesTemplate) HasShouldTweet() bool`

HasShouldTweet returns a boolean if a field has been set.

### GetShouldSendNotifications

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetShouldSendNotifications() bool`

GetShouldSendNotifications returns the ShouldSendNotifications field if non-nil, zero value otherwise.

### GetShouldSendNotificationsOk

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetShouldSendNotificationsOk() (*bool, bool)`

GetShouldSendNotificationsOk returns a tuple with the ShouldSendNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldSendNotifications

`func (o *PostPagesPageIdIncidentTemplatesTemplate) SetShouldSendNotifications(v bool)`

SetShouldSendNotifications sets ShouldSendNotifications field to given value.

### HasShouldSendNotifications

`func (o *PostPagesPageIdIncidentTemplatesTemplate) HasShouldSendNotifications() bool`

HasShouldSendNotifications returns a boolean if a field has been set.

### GetComponentIds

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetComponentIds() []string`

GetComponentIds returns the ComponentIds field if non-nil, zero value otherwise.

### GetComponentIdsOk

`func (o *PostPagesPageIdIncidentTemplatesTemplate) GetComponentIdsOk() (*[]string, bool)`

GetComponentIdsOk returns a tuple with the ComponentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIds

`func (o *PostPagesPageIdIncidentTemplatesTemplate) SetComponentIds(v []string)`

SetComponentIds sets ComponentIds field to given value.

### HasComponentIds

`func (o *PostPagesPageIdIncidentTemplatesTemplate) HasComponentIds() bool`

HasComponentIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


