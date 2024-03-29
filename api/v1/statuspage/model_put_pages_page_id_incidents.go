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

// PutPagesPageIdIncidents Update an incident
type PutPagesPageIdIncidents struct {
	Incident *PatchPagesPageIdIncidentsIncident `json:"incident,omitempty"`
}

// NewPutPagesPageIdIncidents instantiates a new PutPagesPageIdIncidents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutPagesPageIdIncidents() *PutPagesPageIdIncidents {
	this := PutPagesPageIdIncidents{}
	return &this
}

// NewPutPagesPageIdIncidentsWithDefaults instantiates a new PutPagesPageIdIncidents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutPagesPageIdIncidentsWithDefaults() *PutPagesPageIdIncidents {
	this := PutPagesPageIdIncidents{}
	return &this
}

// GetIncident returns the Incident field value if set, zero value otherwise.
func (o *PutPagesPageIdIncidents) GetIncident() PatchPagesPageIdIncidentsIncident {
	if o == nil || o.Incident == nil {
		var ret PatchPagesPageIdIncidentsIncident
		return ret
	}
	return *o.Incident
}

// GetIncidentOk returns a tuple with the Incident field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutPagesPageIdIncidents) GetIncidentOk() (*PatchPagesPageIdIncidentsIncident, bool) {
	if o == nil || o.Incident == nil {
		return nil, false
	}
	return o.Incident, true
}

// HasIncident returns a boolean if a field has been set.
func (o *PutPagesPageIdIncidents) HasIncident() bool {
	if o != nil && o.Incident != nil {
		return true
	}

	return false
}

// SetIncident gets a reference to the given PatchPagesPageIdIncidentsIncident and assigns it to the Incident field.
func (o *PutPagesPageIdIncidents) SetIncident(v PatchPagesPageIdIncidentsIncident) {
	o.Incident = &v
}

func (o PutPagesPageIdIncidents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Incident != nil {
		toSerialize["incident"] = o.Incident
	}
	return json.Marshal(toSerialize)
}

type NullablePutPagesPageIdIncidents struct {
	value *PutPagesPageIdIncidents
	isSet bool
}

func (v NullablePutPagesPageIdIncidents) Get() *PutPagesPageIdIncidents {
	return v.value
}

func (v *NullablePutPagesPageIdIncidents) Set(val *PutPagesPageIdIncidents) {
	v.value = val
	v.isSet = true
}

func (v NullablePutPagesPageIdIncidents) IsSet() bool {
	return v.isSet
}

func (v *NullablePutPagesPageIdIncidents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutPagesPageIdIncidents(val *PutPagesPageIdIncidents) *NullablePutPagesPageIdIncidents {
	return &NullablePutPagesPageIdIncidents{value: val, isSet: true}
}

func (v NullablePutPagesPageIdIncidents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutPagesPageIdIncidents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


