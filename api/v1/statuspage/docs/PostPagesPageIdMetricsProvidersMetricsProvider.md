# PostPagesPageIdMetricsProvidersMetricsProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Required by the Librato metrics provider. | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **string** | Required by the Datadog and NewRelic type metrics providers. | [optional] 
**ApiToken** | Pointer to **string** | Required by the Librato, Datadog and Pingdom type metrics providers. | [optional] 
**ApplicationKey** | Pointer to **string** | Required by the Pingdom-type metrics provider. | [optional] 
**Type** | Pointer to **string** | One of \&quot;Pingdom\&quot;, \&quot;NewRelic\&quot;, \&quot;Librato\&quot;, \&quot;Datadog\&quot;, or \&quot;Self\&quot; | [optional] 
**MetricBaseUri** | Pointer to **string** | Required by the Datadog and NewRelic type metrics providers. | [optional] 

## Methods

### NewPostPagesPageIdMetricsProvidersMetricsProvider

`func NewPostPagesPageIdMetricsProvidersMetricsProvider() *PostPagesPageIdMetricsProvidersMetricsProvider`

NewPostPagesPageIdMetricsProvidersMetricsProvider instantiates a new PostPagesPageIdMetricsProvidersMetricsProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPagesPageIdMetricsProvidersMetricsProviderWithDefaults

`func NewPostPagesPageIdMetricsProvidersMetricsProviderWithDefaults() *PostPagesPageIdMetricsProvidersMetricsProvider`

NewPostPagesPageIdMetricsProvidersMetricsProviderWithDefaults instantiates a new PostPagesPageIdMetricsProvidersMetricsProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetApiKey

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiToken

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetApplicationKey

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApplicationKey() string`

GetApplicationKey returns the ApplicationKey field if non-nil, zero value otherwise.

### GetApplicationKeyOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetApplicationKeyOk() (*string, bool)`

GetApplicationKeyOk returns a tuple with the ApplicationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKey

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetApplicationKey(v string)`

SetApplicationKey sets ApplicationKey field to given value.

### HasApplicationKey

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasApplicationKey() bool`

HasApplicationKey returns a boolean if a field has been set.

### GetType

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetricBaseUri

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetMetricBaseUri() string`

GetMetricBaseUri returns the MetricBaseUri field if non-nil, zero value otherwise.

### GetMetricBaseUriOk

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) GetMetricBaseUriOk() (*string, bool)`

GetMetricBaseUriOk returns a tuple with the MetricBaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricBaseUri

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) SetMetricBaseUri(v string)`

SetMetricBaseUri sets MetricBaseUri field to given value.

### HasMetricBaseUri

`func (o *PostPagesPageIdMetricsProvidersMetricsProvider) HasMetricBaseUri() bool`

HasMetricBaseUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


