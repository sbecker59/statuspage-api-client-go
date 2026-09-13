# IncidentImpact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Incident Impact Identifier | [optional] 
**TenantId** | Pointer to **string** | The tenant ID associated with the impact. | [optional] 
**AtlassianOrganizationId** | Pointer to **string** | The Atlassian organization ID associated with the impact. | [optional] 
**ProductName** | Pointer to **string** | The product name associated with the impact. | [optional] 
**Experiences** | Pointer to **[]string** | The list of experiences impacted. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the impact was created. | [optional] 

## Methods

### NewIncidentImpact

`func NewIncidentImpact() *IncidentImpact`

NewIncidentImpact instantiates a new IncidentImpact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentImpactWithDefaults

`func NewIncidentImpactWithDefaults() *IncidentImpact`

NewIncidentImpactWithDefaults instantiates a new IncidentImpact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IncidentImpact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncidentImpact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncidentImpact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncidentImpact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantId

`func (o *IncidentImpact) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IncidentImpact) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IncidentImpact) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IncidentImpact) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetAtlassianOrganizationId

`func (o *IncidentImpact) GetAtlassianOrganizationId() string`

GetAtlassianOrganizationId returns the AtlassianOrganizationId field if non-nil, zero value otherwise.

### GetAtlassianOrganizationIdOk

`func (o *IncidentImpact) GetAtlassianOrganizationIdOk() (*string, bool)`

GetAtlassianOrganizationIdOk returns a tuple with the AtlassianOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtlassianOrganizationId

`func (o *IncidentImpact) SetAtlassianOrganizationId(v string)`

SetAtlassianOrganizationId sets AtlassianOrganizationId field to given value.

### HasAtlassianOrganizationId

`func (o *IncidentImpact) HasAtlassianOrganizationId() bool`

HasAtlassianOrganizationId returns a boolean if a field has been set.

### GetProductName

`func (o *IncidentImpact) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *IncidentImpact) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *IncidentImpact) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *IncidentImpact) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetExperiences

`func (o *IncidentImpact) GetExperiences() []string`

GetExperiences returns the Experiences field if non-nil, zero value otherwise.

### GetExperiencesOk

`func (o *IncidentImpact) GetExperiencesOk() (*[]string, bool)`

GetExperiencesOk returns a tuple with the Experiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiences

`func (o *IncidentImpact) SetExperiences(v []string)`

SetExperiences sets Experiences field to given value.

### HasExperiences

`func (o *IncidentImpact) HasExperiences() bool`

HasExperiences returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IncidentImpact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IncidentImpact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IncidentImpact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IncidentImpact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


