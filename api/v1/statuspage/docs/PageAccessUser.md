# PageAccessUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Page Access User Identifier | [optional] 
**PageId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ExternalLogin** | Pointer to **string** | IDP login user id. Key is typically \&quot;uid\&quot;. | [optional] 
**PageAccessGroupId** | Pointer to **string** |  | [optional] 
**PageAccessGroupIds** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPageAccessUser

`func NewPageAccessUser() *PageAccessUser`

NewPageAccessUser instantiates a new PageAccessUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageAccessUserWithDefaults

`func NewPageAccessUserWithDefaults() *PageAccessUser`

NewPageAccessUserWithDefaults instantiates a new PageAccessUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageAccessUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageAccessUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageAccessUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageAccessUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPageId

`func (o *PageAccessUser) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PageAccessUser) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PageAccessUser) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *PageAccessUser) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetEmail

`func (o *PageAccessUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PageAccessUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PageAccessUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PageAccessUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalLogin

`func (o *PageAccessUser) GetExternalLogin() string`

GetExternalLogin returns the ExternalLogin field if non-nil, zero value otherwise.

### GetExternalLoginOk

`func (o *PageAccessUser) GetExternalLoginOk() (*string, bool)`

GetExternalLoginOk returns a tuple with the ExternalLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLogin

`func (o *PageAccessUser) SetExternalLogin(v string)`

SetExternalLogin sets ExternalLogin field to given value.

### HasExternalLogin

`func (o *PageAccessUser) HasExternalLogin() bool`

HasExternalLogin returns a boolean if a field has been set.

### GetPageAccessGroupId

`func (o *PageAccessUser) GetPageAccessGroupId() string`

GetPageAccessGroupId returns the PageAccessGroupId field if non-nil, zero value otherwise.

### GetPageAccessGroupIdOk

`func (o *PageAccessUser) GetPageAccessGroupIdOk() (*string, bool)`

GetPageAccessGroupIdOk returns a tuple with the PageAccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageAccessGroupId

`func (o *PageAccessUser) SetPageAccessGroupId(v string)`

SetPageAccessGroupId sets PageAccessGroupId field to given value.

### HasPageAccessGroupId

`func (o *PageAccessUser) HasPageAccessGroupId() bool`

HasPageAccessGroupId returns a boolean if a field has been set.

### GetPageAccessGroupIds

`func (o *PageAccessUser) GetPageAccessGroupIds() string`

GetPageAccessGroupIds returns the PageAccessGroupIds field if non-nil, zero value otherwise.

### GetPageAccessGroupIdsOk

`func (o *PageAccessUser) GetPageAccessGroupIdsOk() (*string, bool)`

GetPageAccessGroupIdsOk returns a tuple with the PageAccessGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageAccessGroupIds

`func (o *PageAccessUser) SetPageAccessGroupIds(v string)`

SetPageAccessGroupIds sets PageAccessGroupIds field to given value.

### HasPageAccessGroupIds

`func (o *PageAccessUser) HasPageAccessGroupIds() bool`

HasPageAccessGroupIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PageAccessUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageAccessUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageAccessUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PageAccessUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PageAccessUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PageAccessUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PageAccessUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PageAccessUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


