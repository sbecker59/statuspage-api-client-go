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

// ComponentUptime Get uptime data for a component that has uptime showcase enabled
type ComponentUptime struct {
	// Start date used for uptime calculation (see the warnings field in the response if this value does not match the start parameter you provided).
	RangeStart *time.Time `json:"range_start,omitempty"`
	// End date used for uptime calculation (see the warnings field in the response if this value does not match the end parameter you provided).
	RangeEnd *time.Time `json:"range_end,omitempty"`
	// Uptime percentage for a component
	UptimePercentage *float32 `json:"uptime_percentage,omitempty"`
	// Seconds of major outage
	MajorOutage *int32 `json:"major_outage,omitempty"`
	// Seconds of partial outage
	PartialOutage *int32 `json:"partial_outage,omitempty"`
	// Warning messages related to the uptime query that may occur
	Warnings *string `json:"warnings,omitempty"`
	// Component identifier
	Id *string `json:"id,omitempty"`
	// Component display name
	Name *string `json:"name,omitempty"`
	RelatedEvents *ComponentUptimeRelatedEvents `json:"related_events,omitempty"`
}

// NewComponentUptime instantiates a new ComponentUptime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentUptime() *ComponentUptime {
	this := ComponentUptime{}
	return &this
}

// NewComponentUptimeWithDefaults instantiates a new ComponentUptime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentUptimeWithDefaults() *ComponentUptime {
	this := ComponentUptime{}
	return &this
}

// GetRangeStart returns the RangeStart field value if set, zero value otherwise.
func (o *ComponentUptime) GetRangeStart() time.Time {
	if o == nil || o.RangeStart == nil {
		var ret time.Time
		return ret
	}
	return *o.RangeStart
}

// GetRangeStartOk returns a tuple with the RangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptime) GetRangeStartOk() (*time.Time, bool) {
	if o == nil || o.RangeStart == nil {
		return nil, false
	}
	return o.RangeStart, true
}

// HasRangeStart returns a boolean if a field has been set.
func (o *ComponentUptime) HasRangeStart() bool {
	if o != nil && o.RangeStart != nil {
		return true
	}

	return false
}

// SetRangeStart gets a reference to the given time.Time and assigns it to the RangeStart field.
func (o *ComponentUptime) SetRangeStart(v time.Time) {
	o.RangeStart = &v
}

// GetRangeEnd returns the RangeEnd field value if set, zero value otherwise.
func (o *ComponentUptime) GetRangeEnd() time.Time {
	if o == nil || o.RangeEnd == nil {
		var ret time.Time
		return ret
	}
	return *o.RangeEnd
}

// GetRangeEndOk returns a tuple with the RangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptime) GetRangeEndOk() (*time.Time, bool) {
	if o == nil || o.RangeEnd == nil {
		return nil, false
	}
	return o.RangeEnd, true
}

// HasRangeEnd returns a boolean if a field has been set.
func (o *ComponentUptime) HasRangeEnd() bool {
	if o != nil && o.RangeEnd != nil {
		return true
	}

	return false
}

// SetRangeEnd gets a reference to the given time.Time and assigns it to the RangeEnd field.
func (o *ComponentUptime) SetRangeEnd(v time.Time) {
	o.RangeEnd = &v
}

// GetUptimePercentage returns the UptimePercentage field value if set, zero value otherwise.
func (o *ComponentUptime) GetUptimePercentage() float32 {
	if o == nil || o.UptimePercentage == nil {
		var ret float32
		return ret
	}
	return *o.UptimePercentage
}

// GetUptimePercentageOk returns a tuple with the UptimePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptime) GetUptimePercentageOk() (*float32, bool) {
	if o == nil || o.UptimePercentage == nil {
		return nil, false
	}
	return o.UptimePercentage, true
}

// HasUptimePercentage returns a boolean if a field has been set.
func (o *ComponentUptime) HasUptimePercentage() bool {
	if o != nil && o.UptimePercentage != nil {
		return true
	}

	return false
}

// SetUptimePercentage gets a reference to the given float32 and assigns it to the UptimePercentage field.
func (o *ComponentUptime) SetUptimePercentage(v float32) {
	o.UptimePercentage = &v
}

// GetMajorOutage returns the MajorOutage field value if set, zero value otherwise.
func (o *ComponentUptime) GetMajorOutage() int32 {
	if o == nil || o.MajorOutage == nil {
		var ret int32
		return ret
	}
	return *o.MajorOutage
}

// GetMajorOutageOk returns a tuple with the MajorOutage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptime) GetMajorOutageOk() (*int32, bool) {
	if o == nil || o.MajorOutage == nil {
		return nil, false
	}
	return o.MajorOutage, true
}

// HasMajorOutage returns a boolean if a field has been set.
func (o *ComponentUptime) HasMajorOutage() bool {
	if o != nil && o.MajorOutage != nil {
		return true
	}

	return false
}

// SetMajorOutage gets a reference to the given int32 and assigns it to the MajorOutage field.
func (o *ComponentUptime) SetMajorOutage(v int32) {
	o.MajorOutage = &v
}

// GetPartialOutage returns the PartialOutage field value if set, zero value otherwise.
func (o *ComponentUptime) GetPartialOutage() int32 {
	if o == nil || o.PartialOutage == nil {
		var ret int32
		return ret
	}
	return *o.PartialOutage
}

// GetPartialOutageOk returns a tuple with the PartialOutage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptime) GetPartialOutageOk() (*int32, bool) {
	if o == nil || o.PartialOutage == nil {
		return nil, false
	}
	return o.PartialOutage, true
}

// HasPartialOutage returns a boolean if a field has been set.
func (o *ComponentUptime) HasPartialOutage() bool {
	if o != nil && o.PartialOutage != nil {
		return true
	}

	return false
}

// SetPartialOutage gets a reference to the given int32 and assigns it to the PartialOutage field.
func (o *ComponentUptime) SetPartialOutage(v int32) {
	o.PartialOutage = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ComponentUptime) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptime) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ComponentUptime) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *ComponentUptime) SetWarnings(v string) {
	o.Warnings = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentUptime) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptime) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentUptime) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComponentUptime) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentUptime) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptime) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentUptime) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentUptime) SetName(v string) {
	o.Name = &v
}

// GetRelatedEvents returns the RelatedEvents field value if set, zero value otherwise.
func (o *ComponentUptime) GetRelatedEvents() ComponentUptimeRelatedEvents {
	if o == nil || o.RelatedEvents == nil {
		var ret ComponentUptimeRelatedEvents
		return ret
	}
	return *o.RelatedEvents
}

// GetRelatedEventsOk returns a tuple with the RelatedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentUptime) GetRelatedEventsOk() (*ComponentUptimeRelatedEvents, bool) {
	if o == nil || o.RelatedEvents == nil {
		return nil, false
	}
	return o.RelatedEvents, true
}

// HasRelatedEvents returns a boolean if a field has been set.
func (o *ComponentUptime) HasRelatedEvents() bool {
	if o != nil && o.RelatedEvents != nil {
		return true
	}

	return false
}

// SetRelatedEvents gets a reference to the given ComponentUptimeRelatedEvents and assigns it to the RelatedEvents field.
func (o *ComponentUptime) SetRelatedEvents(v ComponentUptimeRelatedEvents) {
	o.RelatedEvents = &v
}

func (o ComponentUptime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RangeStart != nil {
		toSerialize["range_start"] = o.RangeStart
	}
	if o.RangeEnd != nil {
		toSerialize["range_end"] = o.RangeEnd
	}
	if o.UptimePercentage != nil {
		toSerialize["uptime_percentage"] = o.UptimePercentage
	}
	if o.MajorOutage != nil {
		toSerialize["major_outage"] = o.MajorOutage
	}
	if o.PartialOutage != nil {
		toSerialize["partial_outage"] = o.PartialOutage
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RelatedEvents != nil {
		toSerialize["related_events"] = o.RelatedEvents
	}
	return json.Marshal(toSerialize)
}

type NullableComponentUptime struct {
	value *ComponentUptime
	isSet bool
}

func (v NullableComponentUptime) Get() *ComponentUptime {
	return v.value
}

func (v *NullableComponentUptime) Set(val *ComponentUptime) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentUptime) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentUptime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentUptime(val *ComponentUptime) *NullableComponentUptime {
	return &NullableComponentUptime{value: val, isSet: true}
}

func (v NullableComponentUptime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentUptime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


