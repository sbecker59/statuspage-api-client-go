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

// PatchPagesPageIdMetricsProviders Update a metric provider
type PatchPagesPageIdMetricsProviders struct {
	MetricsProvider *PatchPagesPageIdMetricsProvidersMetricsProvider `json:"metrics_provider,omitempty"`
}

// NewPatchPagesPageIdMetricsProviders instantiates a new PatchPagesPageIdMetricsProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPagesPageIdMetricsProviders() *PatchPagesPageIdMetricsProviders {
	this := PatchPagesPageIdMetricsProviders{}
	return &this
}

// NewPatchPagesPageIdMetricsProvidersWithDefaults instantiates a new PatchPagesPageIdMetricsProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPagesPageIdMetricsProvidersWithDefaults() *PatchPagesPageIdMetricsProviders {
	this := PatchPagesPageIdMetricsProviders{}
	return &this
}

// GetMetricsProvider returns the MetricsProvider field value if set, zero value otherwise.
func (o *PatchPagesPageIdMetricsProviders) GetMetricsProvider() PatchPagesPageIdMetricsProvidersMetricsProvider {
	if o == nil || o.MetricsProvider == nil {
		var ret PatchPagesPageIdMetricsProvidersMetricsProvider
		return ret
	}
	return *o.MetricsProvider
}

// GetMetricsProviderOk returns a tuple with the MetricsProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdMetricsProviders) GetMetricsProviderOk() (*PatchPagesPageIdMetricsProvidersMetricsProvider, bool) {
	if o == nil || o.MetricsProvider == nil {
		return nil, false
	}
	return o.MetricsProvider, true
}

// HasMetricsProvider returns a boolean if a field has been set.
func (o *PatchPagesPageIdMetricsProviders) HasMetricsProvider() bool {
	if o != nil && o.MetricsProvider != nil {
		return true
	}

	return false
}

// SetMetricsProvider gets a reference to the given PatchPagesPageIdMetricsProvidersMetricsProvider and assigns it to the MetricsProvider field.
func (o *PatchPagesPageIdMetricsProviders) SetMetricsProvider(v PatchPagesPageIdMetricsProvidersMetricsProvider) {
	o.MetricsProvider = &v
}

func (o PatchPagesPageIdMetricsProviders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetricsProvider != nil {
		toSerialize["metrics_provider"] = o.MetricsProvider
	}
	return json.Marshal(toSerialize)
}

type NullablePatchPagesPageIdMetricsProviders struct {
	value *PatchPagesPageIdMetricsProviders
	isSet bool
}

func (v NullablePatchPagesPageIdMetricsProviders) Get() *PatchPagesPageIdMetricsProviders {
	return v.value
}

func (v *NullablePatchPagesPageIdMetricsProviders) Set(val *PatchPagesPageIdMetricsProviders) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPagesPageIdMetricsProviders) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPagesPageIdMetricsProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPagesPageIdMetricsProviders(val *PatchPagesPageIdMetricsProviders) *NullablePatchPagesPageIdMetricsProviders {
	return &NullablePatchPagesPageIdMetricsProviders{value: val, isSet: true}
}

func (v NullablePatchPagesPageIdMetricsProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPagesPageIdMetricsProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


