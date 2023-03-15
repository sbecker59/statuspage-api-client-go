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

// PutOrganizationsOrganizationIdPermissionsPagesPageId struct for PutOrganizationsOrganizationIdPermissionsPagesPageId
type PutOrganizationsOrganizationIdPermissionsPagesPageId struct {
	// Whether or not user should have page configuration role. This field will only be present for pages with Role Based Access Control.
	PageConfiguration *bool `json:"page_configuration,omitempty"`
	// Whether or not user should have incident manager role. This field will only be present for pages with Role Based Access Control.
	IncidentManager *bool `json:"incident_manager,omitempty"`
	// Whether or not user should have maintenance manager role. This field will only be present for pages with Role Based Access Control.
	MaintenanceManager *bool `json:"maintenance_manager,omitempty"`
}

// NewPutOrganizationsOrganizationIdPermissionsPagesPageId instantiates a new PutOrganizationsOrganizationIdPermissionsPagesPageId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutOrganizationsOrganizationIdPermissionsPagesPageId() *PutOrganizationsOrganizationIdPermissionsPagesPageId {
	this := PutOrganizationsOrganizationIdPermissionsPagesPageId{}
	return &this
}

// NewPutOrganizationsOrganizationIdPermissionsPagesPageIdWithDefaults instantiates a new PutOrganizationsOrganizationIdPermissionsPagesPageId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutOrganizationsOrganizationIdPermissionsPagesPageIdWithDefaults() *PutOrganizationsOrganizationIdPermissionsPagesPageId {
	this := PutOrganizationsOrganizationIdPermissionsPagesPageId{}
	return &this
}

// GetPageConfiguration returns the PageConfiguration field value if set, zero value otherwise.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) GetPageConfiguration() bool {
	if o == nil || o.PageConfiguration == nil {
		var ret bool
		return ret
	}
	return *o.PageConfiguration
}

// GetPageConfigurationOk returns a tuple with the PageConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) GetPageConfigurationOk() (*bool, bool) {
	if o == nil || o.PageConfiguration == nil {
		return nil, false
	}
	return o.PageConfiguration, true
}

// HasPageConfiguration returns a boolean if a field has been set.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) HasPageConfiguration() bool {
	if o != nil && o.PageConfiguration != nil {
		return true
	}

	return false
}

// SetPageConfiguration gets a reference to the given bool and assigns it to the PageConfiguration field.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) SetPageConfiguration(v bool) {
	o.PageConfiguration = &v
}

// GetIncidentManager returns the IncidentManager field value if set, zero value otherwise.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) GetIncidentManager() bool {
	if o == nil || o.IncidentManager == nil {
		var ret bool
		return ret
	}
	return *o.IncidentManager
}

// GetIncidentManagerOk returns a tuple with the IncidentManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) GetIncidentManagerOk() (*bool, bool) {
	if o == nil || o.IncidentManager == nil {
		return nil, false
	}
	return o.IncidentManager, true
}

// HasIncidentManager returns a boolean if a field has been set.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) HasIncidentManager() bool {
	if o != nil && o.IncidentManager != nil {
		return true
	}

	return false
}

// SetIncidentManager gets a reference to the given bool and assigns it to the IncidentManager field.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) SetIncidentManager(v bool) {
	o.IncidentManager = &v
}

// GetMaintenanceManager returns the MaintenanceManager field value if set, zero value otherwise.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) GetMaintenanceManager() bool {
	if o == nil || o.MaintenanceManager == nil {
		var ret bool
		return ret
	}
	return *o.MaintenanceManager
}

// GetMaintenanceManagerOk returns a tuple with the MaintenanceManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) GetMaintenanceManagerOk() (*bool, bool) {
	if o == nil || o.MaintenanceManager == nil {
		return nil, false
	}
	return o.MaintenanceManager, true
}

// HasMaintenanceManager returns a boolean if a field has been set.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) HasMaintenanceManager() bool {
	if o != nil && o.MaintenanceManager != nil {
		return true
	}

	return false
}

// SetMaintenanceManager gets a reference to the given bool and assigns it to the MaintenanceManager field.
func (o *PutOrganizationsOrganizationIdPermissionsPagesPageId) SetMaintenanceManager(v bool) {
	o.MaintenanceManager = &v
}

func (o PutOrganizationsOrganizationIdPermissionsPagesPageId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageConfiguration != nil {
		toSerialize["page_configuration"] = o.PageConfiguration
	}
	if o.IncidentManager != nil {
		toSerialize["incident_manager"] = o.IncidentManager
	}
	if o.MaintenanceManager != nil {
		toSerialize["maintenance_manager"] = o.MaintenanceManager
	}
	return json.Marshal(toSerialize)
}

type NullablePutOrganizationsOrganizationIdPermissionsPagesPageId struct {
	value *PutOrganizationsOrganizationIdPermissionsPagesPageId
	isSet bool
}

func (v NullablePutOrganizationsOrganizationIdPermissionsPagesPageId) Get() *PutOrganizationsOrganizationIdPermissionsPagesPageId {
	return v.value
}

func (v *NullablePutOrganizationsOrganizationIdPermissionsPagesPageId) Set(val *PutOrganizationsOrganizationIdPermissionsPagesPageId) {
	v.value = val
	v.isSet = true
}

func (v NullablePutOrganizationsOrganizationIdPermissionsPagesPageId) IsSet() bool {
	return v.isSet
}

func (v *NullablePutOrganizationsOrganizationIdPermissionsPagesPageId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutOrganizationsOrganizationIdPermissionsPagesPageId(val *PutOrganizationsOrganizationIdPermissionsPagesPageId) *NullablePutOrganizationsOrganizationIdPermissionsPagesPageId {
	return &NullablePutOrganizationsOrganizationIdPermissionsPagesPageId{value: val, isSet: true}
}

func (v NullablePutOrganizationsOrganizationIdPermissionsPagesPageId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutOrganizationsOrganizationIdPermissionsPagesPageId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

