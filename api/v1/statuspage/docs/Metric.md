# Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Metric identifier | [optional] 
**MetricsProviderId** | Pointer to **string** | Metric Provider identifier | [optional] 
**MetricIdentifier** | Pointer to **string** | Metric Display identifier used to look up the metric data from the provider | [optional] 
**Name** | Pointer to **string** | Name of metric | [optional] 
**Display** | Pointer to **bool** | Should the metric be displayed | [optional] 
**TooltipDescription** | Pointer to **string** |  | [optional] 
**Backfilled** | Pointer to **bool** |  | [optional] 
**YAxisMin** | Pointer to **float32** |  | [optional] 
**YAxisMax** | Pointer to **float32** |  | [optional] 
**YAxisHidden** | Pointer to **bool** | Should the values on the y axis be hidden on render | [optional] 
**Suffix** | Pointer to **string** | Suffix to describe the units on the graph | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**MostRecentDataAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**LastFetchedAt** | Pointer to **time.Time** |  | [optional] 
**BackfillPercentage** | Pointer to **int32** |  | [optional] 
**ReferenceName** | Pointer to **string** |  | [optional] 

## Methods

### NewMetric

`func NewMetric() *Metric`

NewMetric instantiates a new Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricWithDefaults

`func NewMetricWithDefaults() *Metric`

NewMetricWithDefaults instantiates a new Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Metric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Metric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Metric) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Metric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetricsProviderId

`func (o *Metric) GetMetricsProviderId() string`

GetMetricsProviderId returns the MetricsProviderId field if non-nil, zero value otherwise.

### GetMetricsProviderIdOk

`func (o *Metric) GetMetricsProviderIdOk() (*string, bool)`

GetMetricsProviderIdOk returns a tuple with the MetricsProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsProviderId

`func (o *Metric) SetMetricsProviderId(v string)`

SetMetricsProviderId sets MetricsProviderId field to given value.

### HasMetricsProviderId

`func (o *Metric) HasMetricsProviderId() bool`

HasMetricsProviderId returns a boolean if a field has been set.

### GetMetricIdentifier

`func (o *Metric) GetMetricIdentifier() string`

GetMetricIdentifier returns the MetricIdentifier field if non-nil, zero value otherwise.

### GetMetricIdentifierOk

`func (o *Metric) GetMetricIdentifierOk() (*string, bool)`

GetMetricIdentifierOk returns a tuple with the MetricIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricIdentifier

`func (o *Metric) SetMetricIdentifier(v string)`

SetMetricIdentifier sets MetricIdentifier field to given value.

### HasMetricIdentifier

`func (o *Metric) HasMetricIdentifier() bool`

HasMetricIdentifier returns a boolean if a field has been set.

### GetName

`func (o *Metric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Metric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplay

`func (o *Metric) GetDisplay() bool`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Metric) GetDisplayOk() (*bool, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Metric) SetDisplay(v bool)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *Metric) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetTooltipDescription

`func (o *Metric) GetTooltipDescription() string`

GetTooltipDescription returns the TooltipDescription field if non-nil, zero value otherwise.

### GetTooltipDescriptionOk

`func (o *Metric) GetTooltipDescriptionOk() (*string, bool)`

GetTooltipDescriptionOk returns a tuple with the TooltipDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltipDescription

`func (o *Metric) SetTooltipDescription(v string)`

SetTooltipDescription sets TooltipDescription field to given value.

### HasTooltipDescription

`func (o *Metric) HasTooltipDescription() bool`

HasTooltipDescription returns a boolean if a field has been set.

### GetBackfilled

`func (o *Metric) GetBackfilled() bool`

GetBackfilled returns the Backfilled field if non-nil, zero value otherwise.

### GetBackfilledOk

`func (o *Metric) GetBackfilledOk() (*bool, bool)`

GetBackfilledOk returns a tuple with the Backfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfilled

`func (o *Metric) SetBackfilled(v bool)`

SetBackfilled sets Backfilled field to given value.

### HasBackfilled

`func (o *Metric) HasBackfilled() bool`

HasBackfilled returns a boolean if a field has been set.

### GetYAxisMin

`func (o *Metric) GetYAxisMin() float32`

GetYAxisMin returns the YAxisMin field if non-nil, zero value otherwise.

### GetYAxisMinOk

`func (o *Metric) GetYAxisMinOk() (*float32, bool)`

GetYAxisMinOk returns a tuple with the YAxisMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisMin

`func (o *Metric) SetYAxisMin(v float32)`

SetYAxisMin sets YAxisMin field to given value.

### HasYAxisMin

`func (o *Metric) HasYAxisMin() bool`

HasYAxisMin returns a boolean if a field has been set.

### GetYAxisMax

`func (o *Metric) GetYAxisMax() float32`

GetYAxisMax returns the YAxisMax field if non-nil, zero value otherwise.

### GetYAxisMaxOk

`func (o *Metric) GetYAxisMaxOk() (*float32, bool)`

GetYAxisMaxOk returns a tuple with the YAxisMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisMax

`func (o *Metric) SetYAxisMax(v float32)`

SetYAxisMax sets YAxisMax field to given value.

### HasYAxisMax

`func (o *Metric) HasYAxisMax() bool`

HasYAxisMax returns a boolean if a field has been set.

### GetYAxisHidden

`func (o *Metric) GetYAxisHidden() bool`

GetYAxisHidden returns the YAxisHidden field if non-nil, zero value otherwise.

### GetYAxisHiddenOk

`func (o *Metric) GetYAxisHiddenOk() (*bool, bool)`

GetYAxisHiddenOk returns a tuple with the YAxisHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisHidden

`func (o *Metric) SetYAxisHidden(v bool)`

SetYAxisHidden sets YAxisHidden field to given value.

### HasYAxisHidden

`func (o *Metric) HasYAxisHidden() bool`

HasYAxisHidden returns a boolean if a field has been set.

### GetSuffix

`func (o *Metric) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *Metric) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *Metric) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *Metric) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *Metric) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *Metric) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *Metric) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *Metric) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetMostRecentDataAt

`func (o *Metric) GetMostRecentDataAt() time.Time`

GetMostRecentDataAt returns the MostRecentDataAt field if non-nil, zero value otherwise.

### GetMostRecentDataAtOk

`func (o *Metric) GetMostRecentDataAtOk() (*time.Time, bool)`

GetMostRecentDataAtOk returns a tuple with the MostRecentDataAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentDataAt

`func (o *Metric) SetMostRecentDataAt(v time.Time)`

SetMostRecentDataAt sets MostRecentDataAt field to given value.

### HasMostRecentDataAt

`func (o *Metric) HasMostRecentDataAt() bool`

HasMostRecentDataAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Metric) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Metric) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Metric) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Metric) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Metric) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Metric) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Metric) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Metric) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastFetchedAt

`func (o *Metric) GetLastFetchedAt() time.Time`

GetLastFetchedAt returns the LastFetchedAt field if non-nil, zero value otherwise.

### GetLastFetchedAtOk

`func (o *Metric) GetLastFetchedAtOk() (*time.Time, bool)`

GetLastFetchedAtOk returns a tuple with the LastFetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFetchedAt

`func (o *Metric) SetLastFetchedAt(v time.Time)`

SetLastFetchedAt sets LastFetchedAt field to given value.

### HasLastFetchedAt

`func (o *Metric) HasLastFetchedAt() bool`

HasLastFetchedAt returns a boolean if a field has been set.

### GetBackfillPercentage

`func (o *Metric) GetBackfillPercentage() int32`

GetBackfillPercentage returns the BackfillPercentage field if non-nil, zero value otherwise.

### GetBackfillPercentageOk

`func (o *Metric) GetBackfillPercentageOk() (*int32, bool)`

GetBackfillPercentageOk returns a tuple with the BackfillPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillPercentage

`func (o *Metric) SetBackfillPercentage(v int32)`

SetBackfillPercentage sets BackfillPercentage field to given value.

### HasBackfillPercentage

`func (o *Metric) HasBackfillPercentage() bool`

HasBackfillPercentage returns a boolean if a field has been set.

### GetReferenceName

`func (o *Metric) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *Metric) GetReferenceNameOk() (*string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceName

`func (o *Metric) SetReferenceName(v string)`

SetReferenceName sets ReferenceName field to given value.

### HasReferenceName

`func (o *Metric) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


