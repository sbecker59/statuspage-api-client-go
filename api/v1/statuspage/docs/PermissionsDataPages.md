# PermissionsDataPages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageId** | Pointer to **string** | Page identifier | [optional] 
**PageConfiguration** | Pointer to **bool** | User has page configuration role. This field will only be present if the organization has Role Based Access Control. | [optional] 
**IncidentManager** | Pointer to **bool** | User has incident manager role. This field will only be present if the organization has Role Based Access Control. | [optional] 
**MaintenanceManager** | Pointer to **bool** | User has maintenance manager role. This field will only be present if the organization has Role Based Access Control. | [optional] 

## Methods

### NewPermissionsDataPages

`func NewPermissionsDataPages() *PermissionsDataPages`

NewPermissionsDataPages instantiates a new PermissionsDataPages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsDataPagesWithDefaults

`func NewPermissionsDataPagesWithDefaults() *PermissionsDataPages`

NewPermissionsDataPagesWithDefaults instantiates a new PermissionsDataPages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageId

`func (o *PermissionsDataPages) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PermissionsDataPages) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PermissionsDataPages) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *PermissionsDataPages) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetPageConfiguration

`func (o *PermissionsDataPages) GetPageConfiguration() bool`

GetPageConfiguration returns the PageConfiguration field if non-nil, zero value otherwise.

### GetPageConfigurationOk

`func (o *PermissionsDataPages) GetPageConfigurationOk() (*bool, bool)`

GetPageConfigurationOk returns a tuple with the PageConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageConfiguration

`func (o *PermissionsDataPages) SetPageConfiguration(v bool)`

SetPageConfiguration sets PageConfiguration field to given value.

### HasPageConfiguration

`func (o *PermissionsDataPages) HasPageConfiguration() bool`

HasPageConfiguration returns a boolean if a field has been set.

### GetIncidentManager

`func (o *PermissionsDataPages) GetIncidentManager() bool`

GetIncidentManager returns the IncidentManager field if non-nil, zero value otherwise.

### GetIncidentManagerOk

`func (o *PermissionsDataPages) GetIncidentManagerOk() (*bool, bool)`

GetIncidentManagerOk returns a tuple with the IncidentManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentManager

`func (o *PermissionsDataPages) SetIncidentManager(v bool)`

SetIncidentManager sets IncidentManager field to given value.

### HasIncidentManager

`func (o *PermissionsDataPages) HasIncidentManager() bool`

HasIncidentManager returns a boolean if a field has been set.

### GetMaintenanceManager

`func (o *PermissionsDataPages) GetMaintenanceManager() bool`

GetMaintenanceManager returns the MaintenanceManager field if non-nil, zero value otherwise.

### GetMaintenanceManagerOk

`func (o *PermissionsDataPages) GetMaintenanceManagerOk() (*bool, bool)`

GetMaintenanceManagerOk returns a tuple with the MaintenanceManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceManager

`func (o *PermissionsDataPages) SetMaintenanceManager(v bool)`

SetMaintenanceManager sets MaintenanceManager field to given value.

### HasMaintenanceManager

`func (o *PermissionsDataPages) HasMaintenanceManager() bool`

HasMaintenanceManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


