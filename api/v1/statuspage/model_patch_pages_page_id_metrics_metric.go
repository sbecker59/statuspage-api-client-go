/*
Statuspage API

# Code of Conduct Please don't abuse the API, and please report all feature requests and issues to https://support.atlassian.com/contact  # Rate Limiting Each API token is limited to 1 request / second as measured on a 60 second rolling window. To get this limit increased, please contact us at https://support.atlassian.com/contact  Error codes 420 or 429 indicate that you have exceeded the rate limit and the request has been rejected.  # Basics  ## HTTPS It's required  ## URL Prefix In order to maintain version integrity into the future, the API is versioned. All calls currently begin with the following prefix:    https://api.statuspage.io/v1/  ## RESTful Interface Wherever possible, the API seeks to implement repeatable patterns with logical, representative URLs and descriptive HTTP verbs. Below are some examples and conventions you will see throughout the documentation.  * Collections are buckets: https://api.statuspage.io/v1/pages/asdf123/incidents.json * Elements have unique IDs: https://api.statuspage.io/v1/pages/asdf123/incidents/jklm456.json * GET will retrieve information about a collection/element * POST will create an element in a collection * PATCH will update a single element * PUT will replace a single element in a collection (rarely used) * DELETE will destroy a single element  ## Sending Data Information can be sent in the body as form urlencoded or JSON, but make sure the Content-Type header matches the body structure or the server gremlins will be angry.  All examples are provided in JSON format, however they can easily be converted to form encoding if required.  Some examples of how to convert things are below:      // JSON     {       \"incident\": {         \"name\": \"test incident\",         \"components\": [\"8kbf7d35c070\", \"vtnh60py4yd7\"]       }     }      // Form Encoded (using curl as an example):     curl -X POST https://api.statuspage.io/v1/example \\       -d \"incident[name]=test incident\" \\       -d \"incident[components][]=8kbf7d35c070\" \\       -d \"incident[components][]=vtnh60py4yd7\"  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PatchPagesPageIdMetricsMetric struct for PatchPagesPageIdMetricsMetric
type PatchPagesPageIdMetricsMetric struct {
	// Name of metric
	Name *string `json:"name,omitempty"`
	// Metric Display identifier used to look up the metric data from the provider
	MetricIdentifier *string `json:"metric_identifier,omitempty"`
}

// NewPatchPagesPageIdMetricsMetric instantiates a new PatchPagesPageIdMetricsMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPagesPageIdMetricsMetric() *PatchPagesPageIdMetricsMetric {
	this := PatchPagesPageIdMetricsMetric{}
	return &this
}

// NewPatchPagesPageIdMetricsMetricWithDefaults instantiates a new PatchPagesPageIdMetricsMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPagesPageIdMetricsMetricWithDefaults() *PatchPagesPageIdMetricsMetric {
	this := PatchPagesPageIdMetricsMetric{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchPagesPageIdMetricsMetric) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdMetricsMetric) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchPagesPageIdMetricsMetric) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchPagesPageIdMetricsMetric) SetName(v string) {
	o.Name = &v
}

// GetMetricIdentifier returns the MetricIdentifier field value if set, zero value otherwise.
func (o *PatchPagesPageIdMetricsMetric) GetMetricIdentifier() string {
	if o == nil || o.MetricIdentifier == nil {
		var ret string
		return ret
	}
	return *o.MetricIdentifier
}

// GetMetricIdentifierOk returns a tuple with the MetricIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdMetricsMetric) GetMetricIdentifierOk() (*string, bool) {
	if o == nil || o.MetricIdentifier == nil {
		return nil, false
	}
	return o.MetricIdentifier, true
}

// HasMetricIdentifier returns a boolean if a field has been set.
func (o *PatchPagesPageIdMetricsMetric) HasMetricIdentifier() bool {
	if o != nil && o.MetricIdentifier != nil {
		return true
	}

	return false
}

// SetMetricIdentifier gets a reference to the given string and assigns it to the MetricIdentifier field.
func (o *PatchPagesPageIdMetricsMetric) SetMetricIdentifier(v string) {
	o.MetricIdentifier = &v
}

func (o PatchPagesPageIdMetricsMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.MetricIdentifier != nil {
		toSerialize["metric_identifier"] = o.MetricIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullablePatchPagesPageIdMetricsMetric struct {
	value *PatchPagesPageIdMetricsMetric
	isSet bool
}

func (v NullablePatchPagesPageIdMetricsMetric) Get() *PatchPagesPageIdMetricsMetric {
	return v.value
}

func (v *NullablePatchPagesPageIdMetricsMetric) Set(val *PatchPagesPageIdMetricsMetric) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPagesPageIdMetricsMetric) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPagesPageIdMetricsMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPagesPageIdMetricsMetric(val *PatchPagesPageIdMetricsMetric) *NullablePatchPagesPageIdMetricsMetric {
	return &NullablePatchPagesPageIdMetricsMetric{value: val, isSet: true}
}

func (v NullablePatchPagesPageIdMetricsMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPagesPageIdMetricsMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


