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

// PageAccessUser Delete metric for page access user
type PageAccessUser struct {
	// Page Access User Identifier
	Id *string `json:"id,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	Email *string `json:"email,omitempty"`
	// IDP login user id. Key is typically \"uid\".
	ExternalLogin *string `json:"external_login,omitempty"`
	PageAccessGroupId *string `json:"page_access_group_id,omitempty"`
	PageAccessGroupIds *[]string `json:"page_access_group_ids,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewPageAccessUser instantiates a new PageAccessUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageAccessUser() *PageAccessUser {
	this := PageAccessUser{}
	return &this
}

// NewPageAccessUserWithDefaults instantiates a new PageAccessUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageAccessUserWithDefaults() *PageAccessUser {
	this := PageAccessUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PageAccessUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAccessUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PageAccessUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PageAccessUser) SetId(v string) {
	o.Id = &v
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *PageAccessUser) GetPageId() string {
	if o == nil || o.PageId == nil {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAccessUser) GetPageIdOk() (*string, bool) {
	if o == nil || o.PageId == nil {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *PageAccessUser) HasPageId() bool {
	if o != nil && o.PageId != nil {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *PageAccessUser) SetPageId(v string) {
	o.PageId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PageAccessUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAccessUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PageAccessUser) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PageAccessUser) SetEmail(v string) {
	o.Email = &v
}

// GetExternalLogin returns the ExternalLogin field value if set, zero value otherwise.
func (o *PageAccessUser) GetExternalLogin() string {
	if o == nil || o.ExternalLogin == nil {
		var ret string
		return ret
	}
	return *o.ExternalLogin
}

// GetExternalLoginOk returns a tuple with the ExternalLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAccessUser) GetExternalLoginOk() (*string, bool) {
	if o == nil || o.ExternalLogin == nil {
		return nil, false
	}
	return o.ExternalLogin, true
}

// HasExternalLogin returns a boolean if a field has been set.
func (o *PageAccessUser) HasExternalLogin() bool {
	if o != nil && o.ExternalLogin != nil {
		return true
	}

	return false
}

// SetExternalLogin gets a reference to the given string and assigns it to the ExternalLogin field.
func (o *PageAccessUser) SetExternalLogin(v string) {
	o.ExternalLogin = &v
}

// GetPageAccessGroupId returns the PageAccessGroupId field value if set, zero value otherwise.
func (o *PageAccessUser) GetPageAccessGroupId() string {
	if o == nil || o.PageAccessGroupId == nil {
		var ret string
		return ret
	}
	return *o.PageAccessGroupId
}

// GetPageAccessGroupIdOk returns a tuple with the PageAccessGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAccessUser) GetPageAccessGroupIdOk() (*string, bool) {
	if o == nil || o.PageAccessGroupId == nil {
		return nil, false
	}
	return o.PageAccessGroupId, true
}

// HasPageAccessGroupId returns a boolean if a field has been set.
func (o *PageAccessUser) HasPageAccessGroupId() bool {
	if o != nil && o.PageAccessGroupId != nil {
		return true
	}

	return false
}

// SetPageAccessGroupId gets a reference to the given string and assigns it to the PageAccessGroupId field.
func (o *PageAccessUser) SetPageAccessGroupId(v string) {
	o.PageAccessGroupId = &v
}

// GetPageAccessGroupIds returns the PageAccessGroupIds field value if set, zero value otherwise.
func (o *PageAccessUser) GetPageAccessGroupIds() []string {
	if o == nil || o.PageAccessGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.PageAccessGroupIds
}

// GetPageAccessGroupIdsOk returns a tuple with the PageAccessGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAccessUser) GetPageAccessGroupIdsOk() (*[]string, bool) {
	if o == nil || o.PageAccessGroupIds == nil {
		return nil, false
	}
	return o.PageAccessGroupIds, true
}

// HasPageAccessGroupIds returns a boolean if a field has been set.
func (o *PageAccessUser) HasPageAccessGroupIds() bool {
	if o != nil && o.PageAccessGroupIds != nil {
		return true
	}

	return false
}

// SetPageAccessGroupIds gets a reference to the given string and assigns it to the PageAccessGroupIds field.
func (o *PageAccessUser) SetPageAccessGroupIds(v []string) {
	o.PageAccessGroupIds = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PageAccessUser) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAccessUser) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PageAccessUser) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PageAccessUser) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PageAccessUser) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageAccessUser) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PageAccessUser) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PageAccessUser) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PageAccessUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PageId != nil {
		toSerialize["page_id"] = o.PageId
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.ExternalLogin != nil {
		toSerialize["external_login"] = o.ExternalLogin
	}
	if o.PageAccessGroupId != nil {
		toSerialize["page_access_group_id"] = o.PageAccessGroupId
	}
	if o.PageAccessGroupIds != nil {
		toSerialize["page_access_group_ids"] = o.PageAccessGroupIds
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePageAccessUser struct {
	value *PageAccessUser
	isSet bool
}

func (v NullablePageAccessUser) Get() *PageAccessUser {
	return v.value
}

func (v *NullablePageAccessUser) Set(val *PageAccessUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePageAccessUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePageAccessUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageAccessUser(val *PageAccessUser) *NullablePageAccessUser {
	return &NullablePageAccessUser{value: val, isSet: true}
}

func (v NullablePageAccessUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageAccessUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


