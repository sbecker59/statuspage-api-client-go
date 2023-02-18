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

// PostPagesPageIdMetricsProvidersMetricsProvider struct for PostPagesPageIdMetricsProvidersMetricsProvider
type PostPagesPageIdMetricsProvidersMetricsProvider struct {
	// Required by the Librato metrics provider.
	Email *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	// Required by the Datadog and NewRelic type metrics providers.
	ApiKey *string `json:"api_key,omitempty"`
	// Required by the Librato, Datadog and Pingdom type metrics providers.
	ApiToken *string `json:"api_token,omitempty"`
	// Required by the Pingdom-type metrics provider.
	ApplicationKey *string `json:"application_key,omitempty"`
	// One of \"Pingdom\", \"NewRelic\", \"Librato\", \"Datadog\", or \"Self\"
	Type *string `json:"type,omitempty"`
	// Required by the Datadog and NewRelic type metrics providers.
	MetricBaseUri *string `json:"metric_base_uri,omitempty"`
}

// NewPostPagesPageIdMetricsProvidersMetricsProvider instantiates a new PostPagesPageIdMetricsProvidersMetricsProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPagesPageIdMetricsProvidersMetricsProvider() *PostPagesPageIdMetricsProvidersMetricsProvider {
	this := PostPagesPageIdMetricsProvidersMetricsProvider{}
	return &this
}

// NewPostPagesPageIdMetricsProvidersMetricsProviderWithDefaults instantiates a new PostPagesPageIdMetricsProvidersMetricsProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPagesPageIdMetricsProvidersMetricsProviderWithDefaults() *PostPagesPageIdMetricsProvidersMetricsProvider {
	this := PostPagesPageIdMetricsProvidersMetricsProvider{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetPassword(v string) {
	o.Password = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApiToken() string {
	if o == nil || o.ApiToken == nil {
		var ret string
		return ret
	}
	return *o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApiTokenOk() (*string, bool) {
	if o == nil || o.ApiToken == nil {
		return nil, false
	}
	return o.ApiToken, true
}

// HasApiToken returns a boolean if a field has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasApiToken() bool {
	if o != nil && o.ApiToken != nil {
		return true
	}

	return false
}

// SetApiToken gets a reference to the given string and assigns it to the ApiToken field.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetApiToken(v string) {
	o.ApiToken = &v
}

// GetApplicationKey returns the ApplicationKey field value if set, zero value otherwise.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApplicationKey() string {
	if o == nil || o.ApplicationKey == nil {
		var ret string
		return ret
	}
	return *o.ApplicationKey
}

// GetApplicationKeyOk returns a tuple with the ApplicationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApplicationKeyOk() (*string, bool) {
	if o == nil || o.ApplicationKey == nil {
		return nil, false
	}
	return o.ApplicationKey, true
}

// HasApplicationKey returns a boolean if a field has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasApplicationKey() bool {
	if o != nil && o.ApplicationKey != nil {
		return true
	}

	return false
}

// SetApplicationKey gets a reference to the given string and assigns it to the ApplicationKey field.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetApplicationKey(v string) {
	o.ApplicationKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetType(v string) {
	o.Type = &v
}

// GetMetricBaseUri returns the MetricBaseUri field value if set, zero value otherwise.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetMetricBaseUri() string {
	if o == nil || o.MetricBaseUri == nil {
		var ret string
		return ret
	}
	return *o.MetricBaseUri
}

// GetMetricBaseUriOk returns a tuple with the MetricBaseUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetMetricBaseUriOk() (*string, bool) {
	if o == nil || o.MetricBaseUri == nil {
		return nil, false
	}
	return o.MetricBaseUri, true
}

// HasMetricBaseUri returns a boolean if a field has been set.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasMetricBaseUri() bool {
	if o != nil && o.MetricBaseUri != nil {
		return true
	}

	return false
}

// SetMetricBaseUri gets a reference to the given string and assigns it to the MetricBaseUri field.
func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetMetricBaseUri(v string) {
	o.MetricBaseUri = &v
}

func (o PostPagesPageIdMetricsProvidersMetricsProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.ApiToken != nil {
		toSerialize["api_token"] = o.ApiToken
	}
	if o.ApplicationKey != nil {
		toSerialize["application_key"] = o.ApplicationKey
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.MetricBaseUri != nil {
		toSerialize["metric_base_uri"] = o.MetricBaseUri
	}
	return json.Marshal(toSerialize)
}

type NullablePostPagesPageIdMetricsProvidersMetricsProvider struct {
	value *PostPagesPageIdMetricsProvidersMetricsProvider
	isSet bool
}

func (v NullablePostPagesPageIdMetricsProvidersMetricsProvider) Get() *PostPagesPageIdMetricsProvidersMetricsProvider {
	return v.value
}

func (v *NullablePostPagesPageIdMetricsProvidersMetricsProvider) Set(val *PostPagesPageIdMetricsProvidersMetricsProvider) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPagesPageIdMetricsProvidersMetricsProvider) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPagesPageIdMetricsProvidersMetricsProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPagesPageIdMetricsProvidersMetricsProvider(val *PostPagesPageIdMetricsProvidersMetricsProvider) *NullablePostPagesPageIdMetricsProvidersMetricsProvider {
	return &NullablePostPagesPageIdMetricsProvidersMetricsProvider{value: val, isSet: true}
}

func (v NullablePostPagesPageIdMetricsProvidersMetricsProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPagesPageIdMetricsProvidersMetricsProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


