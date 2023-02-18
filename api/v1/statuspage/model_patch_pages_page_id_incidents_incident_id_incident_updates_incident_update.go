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

// PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate struct for PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate
type PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate struct {
	// Controls whether to create twitter update.
	WantsTwitterUpdate *bool `json:"wants_twitter_update,omitempty"`
	// Incident update body.
	Body *string `json:"body,omitempty"`
	// Timestamp when incident update is happened.
	DisplayAt *time.Time `json:"display_at,omitempty"`
	// Controls whether to delivery notifications.
	DeliverNotifications *bool `json:"deliver_notifications,omitempty"`
}

// NewPatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate instantiates a new PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate() *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate {
	this := PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate{}
	return &this
}

// NewPatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateWithDefaults instantiates a new PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateWithDefaults() *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate {
	this := PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate{}
	return &this
}

// GetWantsTwitterUpdate returns the WantsTwitterUpdate field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) GetWantsTwitterUpdate() bool {
	if o == nil || o.WantsTwitterUpdate == nil {
		var ret bool
		return ret
	}
	return *o.WantsTwitterUpdate
}

// GetWantsTwitterUpdateOk returns a tuple with the WantsTwitterUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) GetWantsTwitterUpdateOk() (*bool, bool) {
	if o == nil || o.WantsTwitterUpdate == nil {
		return nil, false
	}
	return o.WantsTwitterUpdate, true
}

// HasWantsTwitterUpdate returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) HasWantsTwitterUpdate() bool {
	if o != nil && o.WantsTwitterUpdate != nil {
		return true
	}

	return false
}

// SetWantsTwitterUpdate gets a reference to the given bool and assigns it to the WantsTwitterUpdate field.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) SetWantsTwitterUpdate(v bool) {
	o.WantsTwitterUpdate = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) SetBody(v string) {
	o.Body = &v
}

// GetDisplayAt returns the DisplayAt field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) GetDisplayAt() time.Time {
	if o == nil || o.DisplayAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DisplayAt
}

// GetDisplayAtOk returns a tuple with the DisplayAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) GetDisplayAtOk() (*time.Time, bool) {
	if o == nil || o.DisplayAt == nil {
		return nil, false
	}
	return o.DisplayAt, true
}

// HasDisplayAt returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) HasDisplayAt() bool {
	if o != nil && o.DisplayAt != nil {
		return true
	}

	return false
}

// SetDisplayAt gets a reference to the given time.Time and assigns it to the DisplayAt field.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) SetDisplayAt(v time.Time) {
	o.DisplayAt = &v
}

// GetDeliverNotifications returns the DeliverNotifications field value if set, zero value otherwise.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) GetDeliverNotifications() bool {
	if o == nil || o.DeliverNotifications == nil {
		var ret bool
		return ret
	}
	return *o.DeliverNotifications
}

// GetDeliverNotificationsOk returns a tuple with the DeliverNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) GetDeliverNotificationsOk() (*bool, bool) {
	if o == nil || o.DeliverNotifications == nil {
		return nil, false
	}
	return o.DeliverNotifications, true
}

// HasDeliverNotifications returns a boolean if a field has been set.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) HasDeliverNotifications() bool {
	if o != nil && o.DeliverNotifications != nil {
		return true
	}

	return false
}

// SetDeliverNotifications gets a reference to the given bool and assigns it to the DeliverNotifications field.
func (o *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) SetDeliverNotifications(v bool) {
	o.DeliverNotifications = &v
}

func (o PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WantsTwitterUpdate != nil {
		toSerialize["wants_twitter_update"] = o.WantsTwitterUpdate
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.DisplayAt != nil {
		toSerialize["display_at"] = o.DisplayAt
	}
	if o.DeliverNotifications != nil {
		toSerialize["deliver_notifications"] = o.DeliverNotifications
	}
	return json.Marshal(toSerialize)
}

type NullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate struct {
	value *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate
	isSet bool
}

func (v NullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) Get() *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate {
	return v.value
}

func (v *NullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) Set(val *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate(val *PatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) *NullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate {
	return &NullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate{value: val, isSet: true}
}

func (v NullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


