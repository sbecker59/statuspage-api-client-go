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

// ComponentUptimeRelatedEvents Related incidents
type ComponentUptimeRelatedEvents struct {
	// Incident identifier
	Id *string `json:"id,omitempty"`
}

// NewComponentUptimeRelatedEvents instantiates a new ComponentUptimeRelatedEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentUptimeRelatedEvents() *ComponentUptimeRelatedEvents {
	this := ComponentUptimeRelatedEvents{}
	return &this
}

// NewComponentUptimeRelatedEventsWithDefaults instantiates a new ComponentUptimeRelatedEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentUptimeRelatedEventsWithDefaults() *ComponentUptimeRelatedEvents {
	this := ComponentUptimeRelatedEvents{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentUptimeRelatedEvents) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptimeRelatedEvents) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentUptimeRelatedEvents) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComponentUptimeRelatedEvents) SetId(v string) {
	o.Id = &v
}

func (o ComponentUptimeRelatedEvents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableComponentUptimeRelatedEvents struct {
	value *ComponentUptimeRelatedEvents
	isSet bool
}

func (v NullableComponentUptimeRelatedEvents) Get() *ComponentUptimeRelatedEvents {
	return v.value
}

func (v *NullableComponentUptimeRelatedEvents) Set(val *ComponentUptimeRelatedEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentUptimeRelatedEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentUptimeRelatedEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentUptimeRelatedEvents(val *ComponentUptimeRelatedEvents) *NullableComponentUptimeRelatedEvents {
	return &NullableComponentUptimeRelatedEvents{value: val, isSet: true}
}

func (v NullableComponentUptimeRelatedEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentUptimeRelatedEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


