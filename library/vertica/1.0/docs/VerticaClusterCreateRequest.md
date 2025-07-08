# VerticaClusterCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**DbaasEngineVersionId** | **string** | DBaaS engine version ID | 
**InitConfigOption** | [**VerticaInitConfigOptionRequest**](VerticaInitConfigOptionRequest.md) | DB initial configuration option | 
**InstanceGroups** | [**[]InstanceGroupRequest**](InstanceGroupRequest.md) | Instance groups list | 
**InstanceNamePrefix** | **string** | Instance name prefix | 
**License** | Pointer to **string** | license | [optional] 
**MaintenanceOption** | Pointer to [**NullableMaintenanceOption**](MaintenanceOption.md) |  | [optional] 
**Name** | **string** | Cluster name | 
**NatEnabled** | Pointer to **NullableBool** |  | [optional] 
**SubnetId** | **string** | Subnet ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Timezone** | **string** | Timezone | 

## Methods

### NewVerticaClusterCreateRequest

`func NewVerticaClusterCreateRequest(dbaasEngineVersionId string, initConfigOption VerticaInitConfigOptionRequest, instanceGroups []InstanceGroupRequest, instanceNamePrefix string, name string, subnetId string, timezone string, ) *VerticaClusterCreateRequest`

NewVerticaClusterCreateRequest instantiates a new VerticaClusterCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerticaClusterCreateRequestWithDefaults

`func NewVerticaClusterCreateRequestWithDefaults() *VerticaClusterCreateRequest`

NewVerticaClusterCreateRequestWithDefaults instantiates a new VerticaClusterCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowableIpAddresses

`func (o *VerticaClusterCreateRequest) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *VerticaClusterCreateRequest) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *VerticaClusterCreateRequest) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *VerticaClusterCreateRequest) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetDbaasEngineVersionId

`func (o *VerticaClusterCreateRequest) GetDbaasEngineVersionId() string`

GetDbaasEngineVersionId returns the DbaasEngineVersionId field if non-nil, zero value otherwise.

### GetDbaasEngineVersionIdOk

`func (o *VerticaClusterCreateRequest) GetDbaasEngineVersionIdOk() (*string, bool)`

GetDbaasEngineVersionIdOk returns a tuple with the DbaasEngineVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngineVersionId

`func (o *VerticaClusterCreateRequest) SetDbaasEngineVersionId(v string)`

SetDbaasEngineVersionId sets DbaasEngineVersionId field to given value.


### GetInitConfigOption

`func (o *VerticaClusterCreateRequest) GetInitConfigOption() VerticaInitConfigOptionRequest`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *VerticaClusterCreateRequest) GetInitConfigOptionOk() (*VerticaInitConfigOptionRequest, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *VerticaClusterCreateRequest) SetInitConfigOption(v VerticaInitConfigOptionRequest)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceGroups

`func (o *VerticaClusterCreateRequest) GetInstanceGroups() []InstanceGroupRequest`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *VerticaClusterCreateRequest) GetInstanceGroupsOk() (*[]InstanceGroupRequest, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *VerticaClusterCreateRequest) SetInstanceGroups(v []InstanceGroupRequest)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetInstanceNamePrefix

`func (o *VerticaClusterCreateRequest) GetInstanceNamePrefix() string`

GetInstanceNamePrefix returns the InstanceNamePrefix field if non-nil, zero value otherwise.

### GetInstanceNamePrefixOk

`func (o *VerticaClusterCreateRequest) GetInstanceNamePrefixOk() (*string, bool)`

GetInstanceNamePrefixOk returns a tuple with the InstanceNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNamePrefix

`func (o *VerticaClusterCreateRequest) SetInstanceNamePrefix(v string)`

SetInstanceNamePrefix sets InstanceNamePrefix field to given value.


### GetLicense

`func (o *VerticaClusterCreateRequest) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *VerticaClusterCreateRequest) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *VerticaClusterCreateRequest) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *VerticaClusterCreateRequest) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMaintenanceOption

`func (o *VerticaClusterCreateRequest) GetMaintenanceOption() MaintenanceOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *VerticaClusterCreateRequest) GetMaintenanceOptionOk() (*MaintenanceOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *VerticaClusterCreateRequest) SetMaintenanceOption(v MaintenanceOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *VerticaClusterCreateRequest) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *VerticaClusterCreateRequest) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *VerticaClusterCreateRequest) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetName

`func (o *VerticaClusterCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerticaClusterCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerticaClusterCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *VerticaClusterCreateRequest) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *VerticaClusterCreateRequest) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *VerticaClusterCreateRequest) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *VerticaClusterCreateRequest) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### SetNatEnabledNil

`func (o *VerticaClusterCreateRequest) SetNatEnabledNil(b bool)`

 SetNatEnabledNil sets the value for NatEnabled to be an explicit nil

### UnsetNatEnabled
`func (o *VerticaClusterCreateRequest) UnsetNatEnabled()`

UnsetNatEnabled ensures that no value is present for NatEnabled, not even an explicit nil
### GetSubnetId

`func (o *VerticaClusterCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *VerticaClusterCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *VerticaClusterCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *VerticaClusterCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VerticaClusterCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VerticaClusterCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VerticaClusterCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *VerticaClusterCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *VerticaClusterCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTimezone

`func (o *VerticaClusterCreateRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *VerticaClusterCreateRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *VerticaClusterCreateRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


