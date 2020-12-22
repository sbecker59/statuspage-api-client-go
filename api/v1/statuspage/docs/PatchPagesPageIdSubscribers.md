# PatchPagesPageIdSubscribers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIds** | Pointer to **[]string** | A list of component ids for which the subscriber should recieve updates for. Components must be an array with at least one element if it is passed at all. Each component must belong to the page indicated in the path. To set the subscriber to be subscribed to all components on the page, exclude this parameter. | [optional] 

## Methods

### NewPatchPagesPageIdSubscribers

`func NewPatchPagesPageIdSubscribers() *PatchPagesPageIdSubscribers`

NewPatchPagesPageIdSubscribers instantiates a new PatchPagesPageIdSubscribers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPagesPageIdSubscribersWithDefaults

`func NewPatchPagesPageIdSubscribersWithDefaults() *PatchPagesPageIdSubscribers`

NewPatchPagesPageIdSubscribersWithDefaults instantiates a new PatchPagesPageIdSubscribers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIds

`func (o *PatchPagesPageIdSubscribers) GetComponentIds() []string`

GetComponentIds returns the ComponentIds field if non-nil, zero value otherwise.

### GetComponentIdsOk

`func (o *PatchPagesPageIdSubscribers) GetComponentIdsOk() (*[]string, bool)`

GetComponentIdsOk returns a tuple with the ComponentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIds

`func (o *PatchPagesPageIdSubscribers) SetComponentIds(v []string)`

SetComponentIds sets ComponentIds field to given value.

### HasComponentIds

`func (o *PatchPagesPageIdSubscribers) HasComponentIds() bool`

HasComponentIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


