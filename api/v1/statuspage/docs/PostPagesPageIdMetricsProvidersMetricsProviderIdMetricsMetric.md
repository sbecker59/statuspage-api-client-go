# PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of metric | [optional] 
**MetricIdentifier** | Pointer to **string** | The identifier used to look up the metric data from the provider | [optional] 
**Transform** | Pointer to **string** | The transform to apply to metric before pulling into Statuspage. One of: \&quot;average\&quot;, \&quot;count\&quot;, \&quot;max\&quot;, \&quot;min\&quot;, or \&quot;sum\&quot; | [optional] 
**Suffix** | Pointer to **string** | Suffix to describe the units on the graph | [optional] 
**YAxisMin** | Pointer to **int32** | The lower bound of the y axis | [optional] 
**YAxisMax** | Pointer to **int32** | The upper bound of the y axis | [optional] 
**YAxisHidden** | Pointer to **bool** | Should the values on the y axis be hidden on render | [optional] 
**Display** | Pointer to **bool** | Should the metric be displayed | [optional] 
**DecimalPlaces** | Pointer to **int32** | How many decimal places to render on the graph | [optional] 
**TooltipDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewPostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric

`func NewPostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric() *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric`

NewPostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric instantiates a new PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetricWithDefaults

`func NewPostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetricWithDefaults() *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric`

NewPostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetricWithDefaults instantiates a new PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetricIdentifier

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetMetricIdentifier() string`

GetMetricIdentifier returns the MetricIdentifier field if non-nil, zero value otherwise.

### GetMetricIdentifierOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetMetricIdentifierOk() (*string, bool)`

GetMetricIdentifierOk returns a tuple with the MetricIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricIdentifier

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetMetricIdentifier(v string)`

SetMetricIdentifier sets MetricIdentifier field to given value.

### HasMetricIdentifier

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasMetricIdentifier() bool`

HasMetricIdentifier returns a boolean if a field has been set.

### GetTransform

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetTransform() string`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetTransformOk() (*string, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetTransform(v string)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasTransform() bool`

HasTransform returns a boolean if a field has been set.

### GetSuffix

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetYAxisMin

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetYAxisMin() int32`

GetYAxisMin returns the YAxisMin field if non-nil, zero value otherwise.

### GetYAxisMinOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetYAxisMinOk() (*int32, bool)`

GetYAxisMinOk returns a tuple with the YAxisMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisMin

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetYAxisMin(v int32)`

SetYAxisMin sets YAxisMin field to given value.

### HasYAxisMin

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasYAxisMin() bool`

HasYAxisMin returns a boolean if a field has been set.

### GetYAxisMax

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetYAxisMax() int32`

GetYAxisMax returns the YAxisMax field if non-nil, zero value otherwise.

### GetYAxisMaxOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetYAxisMaxOk() (*int32, bool)`

GetYAxisMaxOk returns a tuple with the YAxisMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisMax

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetYAxisMax(v int32)`

SetYAxisMax sets YAxisMax field to given value.

### HasYAxisMax

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasYAxisMax() bool`

HasYAxisMax returns a boolean if a field has been set.

### GetYAxisHidden

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetYAxisHidden() bool`

GetYAxisHidden returns the YAxisHidden field if non-nil, zero value otherwise.

### GetYAxisHiddenOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetYAxisHiddenOk() (*bool, bool)`

GetYAxisHiddenOk returns a tuple with the YAxisHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisHidden

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetYAxisHidden(v bool)`

SetYAxisHidden sets YAxisHidden field to given value.

### HasYAxisHidden

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasYAxisHidden() bool`

HasYAxisHidden returns a boolean if a field has been set.

### GetDisplay

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetDisplay() bool`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetDisplayOk() (*bool, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetDisplay(v bool)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetTooltipDescription

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetTooltipDescription() string`

GetTooltipDescription returns the TooltipDescription field if non-nil, zero value otherwise.

### GetTooltipDescriptionOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) GetTooltipDescriptionOk() (*string, bool)`

GetTooltipDescriptionOk returns a tuple with the TooltipDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltipDescription

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) SetTooltipDescription(v string)`

SetTooltipDescription sets TooltipDescription field to given value.

### HasTooltipDescription

`func (o *PostPagesPageIdMetricsProvidersMetricsProviderIdMetricsMetric) HasTooltipDescription() bool`

HasTooltipDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


