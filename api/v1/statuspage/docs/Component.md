# Component

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier for component | [optional] 
**PageId** | Pointer to **string** | Page identifier | [optional] 
**GroupId** | Pointer to **string** | Component Group identifier | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Group** | Pointer to **bool** | Is this component a group | [optional] 
**Name** | Pointer to **string** | Display name for component | [optional] 
**Description** | Pointer to **string** | More detailed description for component | [optional] 
**Position** | Pointer to **int32** | Order the component will appear on the page | [optional] 
**Status** | Pointer to **string** | Status of component | [optional] 
**Showcase** | Pointer to **bool** | Should this component be showcased | [optional] 
**OnlyShowIfDegraded** | Pointer to **bool** | Requires a special feature flag to be enabled | [optional] 
**AutomationEmail** | Pointer to **string** | Requires a special feature flag to be enabled | [optional] 
**StartDate** | Pointer to **string** | The date this component started being used | [optional] 

## Methods

### NewComponent

`func NewComponent() *Component`

NewComponent instantiates a new Component object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentWithDefaults

`func NewComponentWithDefaults() *Component`

NewComponentWithDefaults instantiates a new Component object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Component) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Component) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Component) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Component) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPageId

`func (o *Component) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *Component) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *Component) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *Component) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetGroupId

`func (o *Component) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Component) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Component) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Component) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Component) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Component) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Component) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Component) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Component) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Component) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Component) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Component) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetGroup

`func (o *Component) GetGroup() bool`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Component) GetGroupOk() (*bool, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Component) SetGroup(v bool)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Component) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetName

`func (o *Component) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Component) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Component) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Component) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Component) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Component) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Component) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Component) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPosition

`func (o *Component) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Component) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Component) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Component) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStatus

`func (o *Component) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Component) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Component) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Component) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetShowcase

`func (o *Component) GetShowcase() bool`

GetShowcase returns the Showcase field if non-nil, zero value otherwise.

### GetShowcaseOk

`func (o *Component) GetShowcaseOk() (*bool, bool)`

GetShowcaseOk returns a tuple with the Showcase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowcase

`func (o *Component) SetShowcase(v bool)`

SetShowcase sets Showcase field to given value.

### HasShowcase

`func (o *Component) HasShowcase() bool`

HasShowcase returns a boolean if a field has been set.

### GetOnlyShowIfDegraded

`func (o *Component) GetOnlyShowIfDegraded() bool`

GetOnlyShowIfDegraded returns the OnlyShowIfDegraded field if non-nil, zero value otherwise.

### GetOnlyShowIfDegradedOk

`func (o *Component) GetOnlyShowIfDegradedOk() (*bool, bool)`

GetOnlyShowIfDegradedOk returns a tuple with the OnlyShowIfDegraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyShowIfDegraded

`func (o *Component) SetOnlyShowIfDegraded(v bool)`

SetOnlyShowIfDegraded sets OnlyShowIfDegraded field to given value.

### HasOnlyShowIfDegraded

`func (o *Component) HasOnlyShowIfDegraded() bool`

HasOnlyShowIfDegraded returns a boolean if a field has been set.

### GetAutomationEmail

`func (o *Component) GetAutomationEmail() string`

GetAutomationEmail returns the AutomationEmail field if non-nil, zero value otherwise.

### GetAutomationEmailOk

`func (o *Component) GetAutomationEmailOk() (*string, bool)`

GetAutomationEmailOk returns a tuple with the AutomationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationEmail

`func (o *Component) SetAutomationEmail(v string)`

SetAutomationEmail sets AutomationEmail field to given value.

### HasAutomationEmail

`func (o *Component) HasAutomationEmail() bool`

HasAutomationEmail returns a boolean if a field has been set.

### GetStartDate

`func (o *Component) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Component) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Component) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Component) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


