# PageAccessGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Page Access Group Identifier | [optional] 
**PageId** | Pointer to **string** | Page Identifier. | [optional] 
**Name** | Pointer to **string** | Name for this Group. | [optional] 
**PageAccessUserIds** | Pointer to **[]string** |  | [optional] 
**ExternalIdentifier** | Pointer to **string** | Associates group with external group. | [optional] 
**MetricIds** | Pointer to **[]string** |  | [optional] 
**ComponentIds** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPageAccessGroup

`func NewPageAccessGroup() *PageAccessGroup`

NewPageAccessGroup instantiates a new PageAccessGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageAccessGroupWithDefaults

`func NewPageAccessGroupWithDefaults() *PageAccessGroup`

NewPageAccessGroupWithDefaults instantiates a new PageAccessGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PageAccessGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageAccessGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageAccessGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageAccessGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPageId

`func (o *PageAccessGroup) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PageAccessGroup) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PageAccessGroup) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *PageAccessGroup) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetName

`func (o *PageAccessGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PageAccessGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PageAccessGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PageAccessGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPageAccessUserIds

`func (o *PageAccessGroup) GetPageAccessUserIds() []string`

GetPageAccessUserIds returns the PageAccessUserIds field if non-nil, zero value otherwise.

### GetPageAccessUserIdsOk

`func (o *PageAccessGroup) GetPageAccessUserIdsOk() (*[]string, bool)`

GetPageAccessUserIdsOk returns a tuple with the PageAccessUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageAccessUserIds

`func (o *PageAccessGroup) SetPageAccessUserIds(v []string)`

SetPageAccessUserIds sets PageAccessUserIds field to given value.

### HasPageAccessUserIds

`func (o *PageAccessGroup) HasPageAccessUserIds() bool`

HasPageAccessUserIds returns a boolean if a field has been set.

### GetExternalIdentifier

`func (o *PageAccessGroup) GetExternalIdentifier() string`

GetExternalIdentifier returns the ExternalIdentifier field if non-nil, zero value otherwise.

### GetExternalIdentifierOk

`func (o *PageAccessGroup) GetExternalIdentifierOk() (*string, bool)`

GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentifier

`func (o *PageAccessGroup) SetExternalIdentifier(v string)`

SetExternalIdentifier sets ExternalIdentifier field to given value.

### HasExternalIdentifier

`func (o *PageAccessGroup) HasExternalIdentifier() bool`

HasExternalIdentifier returns a boolean if a field has been set.

### GetMetricIds

`func (o *PageAccessGroup) GetMetricIds() []string`

GetMetricIds returns the MetricIds field if non-nil, zero value otherwise.

### GetMetricIdsOk

`func (o *PageAccessGroup) GetMetricIdsOk() (*[]string, bool)`

GetMetricIdsOk returns a tuple with the MetricIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricIds

`func (o *PageAccessGroup) SetMetricIds(v []string)`

SetMetricIds sets MetricIds field to given value.

### HasMetricIds

`func (o *PageAccessGroup) HasMetricIds() bool`

HasMetricIds returns a boolean if a field has been set.

### GetComponentIds

`func (o *PageAccessGroup) GetComponentIds() []string`

GetComponentIds returns the ComponentIds field if non-nil, zero value otherwise.

### GetComponentIdsOk

`func (o *PageAccessGroup) GetComponentIdsOk() (*[]string, bool)`

GetComponentIdsOk returns a tuple with the ComponentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIds

`func (o *PageAccessGroup) SetComponentIds(v []string)`

SetComponentIds sets ComponentIds field to given value.

### HasComponentIds

`func (o *PageAccessGroup) HasComponentIds() bool`

HasComponentIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PageAccessGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageAccessGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageAccessGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PageAccessGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PageAccessGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PageAccessGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PageAccessGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PageAccessGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


