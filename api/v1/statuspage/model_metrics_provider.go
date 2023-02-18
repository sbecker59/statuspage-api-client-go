/*
Statuspage API

# Code of Conduct Please don't abuse the API, and please report all feature requests and issues to https://support.atlassian.com/contact  # Rate Limiting Each API token is limited to 1 request / second as measured on a 60 second rolling window. To get this limit increased, please contact us at https://support.atlassian.com/contact  Error codes 420 or 429 indicate that you have exceeded the rate limit and the request has been rejected.  # Basics  ## HTTPS It's required  ## URL Prefix In order to maintain version integrity into the future, the API is versioned. All calls currently begin with the following prefix:    https://api.statuspage.io/v1/  ## RESTful Interface Wherever possible, the API seeks to implement repeatable patterns with logical, representative URLs and descriptive HTTP verbs. Below are some examples and conventions you will see throughout the documentation.  * Collections are buckets: https://api.statuspage.io/v1/pages/asdf123/incidents.json * Elements have unique IDs: https://api.statuspage.io/v1/pages/asdf123/incidents/jklm456.json * GET will retrieve information about a collection/element * POST will create an element in a collection * PATCH will update a single element * PUT will replace a single element in a collection (rarely used) * DELETE will destroy a single element  ## Sending Data Information can be sent in the body as form urlencoded or JSON, but make sure the Content-Type header matches the body structure or the server gremlins will be angry.  All examples are provided in JSON format, however they can easily be converted to form encoding if required.  Some examples of how to convert things are below:      // JSON     {       \"incident\": {         \"name\": \"test incident\",         \"components\": [\"8kbf7d35c070\", \"vtnh60py4yd7\"]       }     }      // Form Encoded (using curl as an example):     curl -X POST https://api.statuspage.io/v1/example \\       -d \"incident[name]=test incident\" \\       -d \"incident[components][]=8kbf7d35c070\" \\       -d \"incident[components][]=vtnh60py4yd7\"  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MetricsProvider Delete a metric provider
type MetricsProvider struct {
	// Identifier for Metrics Provider
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	MetricBaseUri *string `json:"metric_base_uri,omitempty"`
	LastRevalidatedAt *time.Time `json:"last_revalidated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	PageId *int32 `json:"page_id,omitempty"`
}

// NewMetricsProvider instantiates a new MetricsProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsProvider() *MetricsProvider {
	this := MetricsProvider{}
	return &this
}

// NewMetricsProviderWithDefaults instantiates a new MetricsProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsProviderWithDefaults() *MetricsProvider {
	this := MetricsProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricsProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricsProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricsProvider) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetricsProvider) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsProvider) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetricsProvider) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MetricsProvider) SetType(v string) {
	o.Type = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *MetricsProvider) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsProvider) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *MetricsProvider) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *MetricsProvider) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetMetricBaseUri returns the MetricBaseUri field value if set, zero value otherwise.
func (o *MetricsProvider) GetMetricBaseUri() string {
	if o == nil || o.MetricBaseUri == nil {
		var ret string
		return ret
	}
	return *o.MetricBaseUri
}

// GetMetricBaseUriOk returns a tuple with the MetricBaseUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsProvider) GetMetricBaseUriOk() (*string, bool) {
	if o == nil || o.MetricBaseUri == nil {
		return nil, false
	}
	return o.MetricBaseUri, true
}

// HasMetricBaseUri returns a boolean if a field has been set.
func (o *MetricsProvider) HasMetricBaseUri() bool {
	if o != nil && o.MetricBaseUri != nil {
		return true
	}

	return false
}

// SetMetricBaseUri gets a reference to the given string and assigns it to the MetricBaseUri field.
func (o *MetricsProvider) SetMetricBaseUri(v string) {
	o.MetricBaseUri = &v
}

// GetLastRevalidatedAt returns the LastRevalidatedAt field value if set, zero value otherwise.
func (o *MetricsProvider) GetLastRevalidatedAt() time.Time {
	if o == nil || o.LastRevalidatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRevalidatedAt
}

// GetLastRevalidatedAtOk returns a tuple with the LastRevalidatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsProvider) GetLastRevalidatedAtOk() (*time.Time, bool) {
	if o == nil || o.LastRevalidatedAt == nil {
		return nil, false
	}
	return o.LastRevalidatedAt, true
}

// HasLastRevalidatedAt returns a boolean if a field has been set.
func (o *MetricsProvider) HasLastRevalidatedAt() bool {
	if o != nil && o.LastRevalidatedAt != nil {
		return true
	}

	return false
}

// SetLastRevalidatedAt gets a reference to the given time.Time and assigns it to the LastRevalidatedAt field.
func (o *MetricsProvider) SetLastRevalidatedAt(v time.Time) {
	o.LastRevalidatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MetricsProvider) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsProvider) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MetricsProvider) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MetricsProvider) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MetricsProvider) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsProvider) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MetricsProvider) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MetricsProvider) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *MetricsProvider) GetPageId() int32 {
	if o == nil || o.PageId == nil {
		var ret int32
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsProvider) GetPageIdOk() (*int32, bool) {
	if o == nil || o.PageId == nil {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *MetricsProvider) HasPageId() bool {
	if o != nil && o.PageId != nil {
		return true
	}

	return false
}

// SetPageId gets a reference to the given int32 and assigns it to the PageId field.
func (o *MetricsProvider) SetPageId(v int32) {
	o.PageId = &v
}

func (o MetricsProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.MetricBaseUri != nil {
		toSerialize["metric_base_uri"] = o.MetricBaseUri
	}
	if o.LastRevalidatedAt != nil {
		toSerialize["last_revalidated_at"] = o.LastRevalidatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.PageId != nil {
		toSerialize["page_id"] = o.PageId
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsProvider struct {
	value *MetricsProvider
	isSet bool
}

func (v NullableMetricsProvider) Get() *MetricsProvider {
	return v.value
}

func (v *NullableMetricsProvider) Set(val *MetricsProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsProvider(val *MetricsProvider) *NullableMetricsProvider {
	return &NullableMetricsProvider{value: val, isSet: true}
}

func (v NullableMetricsProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


