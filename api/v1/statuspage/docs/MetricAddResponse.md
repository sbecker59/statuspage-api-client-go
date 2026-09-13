# MetricAddResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricId** | Pointer to [**[]MetricAddResponseMetricIdInner**](MetricAddResponseMetricIdInner.md) | Metric identifier to add data to | [optional] 

## Methods

### NewMetricAddResponse

`func NewMetricAddResponse() *MetricAddResponse`

NewMetricAddResponse instantiates a new MetricAddResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricAddResponseWithDefaults

`func NewMetricAddResponseWithDefaults() *MetricAddResponse`

NewMetricAddResponseWithDefaults instantiates a new MetricAddResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricId

`func (o *MetricAddResponse) GetMetricId() []MetricAddResponseMetricIdInner`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricAddResponse) GetMetricIdOk() (*[]MetricAddResponseMetricIdInner, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricAddResponse) SetMetricId(v []MetricAddResponseMetricIdInner)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *MetricAddResponse) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


