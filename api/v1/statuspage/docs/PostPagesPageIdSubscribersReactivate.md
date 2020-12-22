# PostPagesPageIdSubscribersReactivate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscribers** | **string** | The array of quarantined subscriber codes to reactivate, or \&quot;all\&quot; to reactivate all quarantined subscribers. | 
**Type** | Pointer to **string** | If this is present, only reactivate subscribers of this type. | [optional] 

## Methods

### NewPostPagesPageIdSubscribersReactivate

`func NewPostPagesPageIdSubscribersReactivate(subscribers string, ) *PostPagesPageIdSubscribersReactivate`

NewPostPagesPageIdSubscribersReactivate instantiates a new PostPagesPageIdSubscribersReactivate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPagesPageIdSubscribersReactivateWithDefaults

`func NewPostPagesPageIdSubscribersReactivateWithDefaults() *PostPagesPageIdSubscribersReactivate`

NewPostPagesPageIdSubscribersReactivateWithDefaults instantiates a new PostPagesPageIdSubscribersReactivate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribers

`func (o *PostPagesPageIdSubscribersReactivate) GetSubscribers() string`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *PostPagesPageIdSubscribersReactivate) GetSubscribersOk() (*string, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *PostPagesPageIdSubscribersReactivate) SetSubscribers(v string)`

SetSubscribers sets Subscribers field to given value.


### GetType

`func (o *PostPagesPageIdSubscribersReactivate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostPagesPageIdSubscribersReactivate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostPagesPageIdSubscribersReactivate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostPagesPageIdSubscribersReactivate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


