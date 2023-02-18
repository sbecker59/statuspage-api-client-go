# PermissionsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | User identifier | [optional] 
**Pages** | Pointer to [**PermissionsDataPages**](PermissionsDataPages.md) |  | [optional] 

## Methods

### NewPermissionsData

`func NewPermissionsData() *PermissionsData`

NewPermissionsData instantiates a new PermissionsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsDataWithDefaults

`func NewPermissionsDataWithDefaults() *PermissionsData`

NewPermissionsDataWithDefaults instantiates a new PermissionsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PermissionsData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PermissionsData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PermissionsData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PermissionsData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPages

`func (o *PermissionsData) GetPages() PermissionsDataPages`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PermissionsData) GetPagesOk() (*PermissionsDataPages, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PermissionsData) SetPages(v PermissionsDataPages)`

SetPages sets Pages field to given value.

### HasPages

`func (o *PermissionsData) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


