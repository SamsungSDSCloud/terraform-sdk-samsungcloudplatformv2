# SearchEngineClusterCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**DbaasEngineVersionId** | **string** | DBaaS engine version ID | 
**InitConfigOption** | [**SearchEngineInitConfigOptionRequest**](SearchEngineInitConfigOptionRequest.md) | DB initial configuration option | 
**InstanceGroups** | [**[]InstanceGroupRequest**](InstanceGroupRequest.md) | Instance groups list | 
**InstanceNamePrefix** | **string** | Instance name prefix | 
**IsCombined** | Pointer to **bool** | Master Data server separation state | [optional] [default to false]
**License** | Pointer to **NullableString** |  | [optional] 
**MaintenanceOption** | Pointer to [**NullableMaintenanceOption**](MaintenanceOption.md) |  | [optional] 
**Name** | **string** | Cluster name | 
**NatEnabled** | Pointer to **bool** | NAT availability | [optional] [default to false]
**SubnetId** | **string** | Subnet ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Timezone** | **string** | Timezone | 

## Methods

### NewSearchEngineClusterCreateRequest

`func NewSearchEngineClusterCreateRequest(dbaasEngineVersionId string, initConfigOption SearchEngineInitConfigOptionRequest, instanceGroups []InstanceGroupRequest, instanceNamePrefix string, name string, subnetId string, timezone string, ) *SearchEngineClusterCreateRequest`

NewSearchEngineClusterCreateRequest instantiates a new SearchEngineClusterCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEngineClusterCreateRequestWithDefaults

`func NewSearchEngineClusterCreateRequestWithDefaults() *SearchEngineClusterCreateRequest`

NewSearchEngineClusterCreateRequestWithDefaults instantiates a new SearchEngineClusterCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowableIpAddresses

`func (o *SearchEngineClusterCreateRequest) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *SearchEngineClusterCreateRequest) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *SearchEngineClusterCreateRequest) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *SearchEngineClusterCreateRequest) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetDbaasEngineVersionId

`func (o *SearchEngineClusterCreateRequest) GetDbaasEngineVersionId() string`

GetDbaasEngineVersionId returns the DbaasEngineVersionId field if non-nil, zero value otherwise.

### GetDbaasEngineVersionIdOk

`func (o *SearchEngineClusterCreateRequest) GetDbaasEngineVersionIdOk() (*string, bool)`

GetDbaasEngineVersionIdOk returns a tuple with the DbaasEngineVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngineVersionId

`func (o *SearchEngineClusterCreateRequest) SetDbaasEngineVersionId(v string)`

SetDbaasEngineVersionId sets DbaasEngineVersionId field to given value.


### GetInitConfigOption

`func (o *SearchEngineClusterCreateRequest) GetInitConfigOption() SearchEngineInitConfigOptionRequest`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *SearchEngineClusterCreateRequest) GetInitConfigOptionOk() (*SearchEngineInitConfigOptionRequest, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *SearchEngineClusterCreateRequest) SetInitConfigOption(v SearchEngineInitConfigOptionRequest)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceGroups

`func (o *SearchEngineClusterCreateRequest) GetInstanceGroups() []InstanceGroupRequest`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *SearchEngineClusterCreateRequest) GetInstanceGroupsOk() (*[]InstanceGroupRequest, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *SearchEngineClusterCreateRequest) SetInstanceGroups(v []InstanceGroupRequest)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetInstanceNamePrefix

`func (o *SearchEngineClusterCreateRequest) GetInstanceNamePrefix() string`

GetInstanceNamePrefix returns the InstanceNamePrefix field if non-nil, zero value otherwise.

### GetInstanceNamePrefixOk

`func (o *SearchEngineClusterCreateRequest) GetInstanceNamePrefixOk() (*string, bool)`

GetInstanceNamePrefixOk returns a tuple with the InstanceNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNamePrefix

`func (o *SearchEngineClusterCreateRequest) SetInstanceNamePrefix(v string)`

SetInstanceNamePrefix sets InstanceNamePrefix field to given value.


### GetIsCombined

`func (o *SearchEngineClusterCreateRequest) GetIsCombined() bool`

GetIsCombined returns the IsCombined field if non-nil, zero value otherwise.

### GetIsCombinedOk

`func (o *SearchEngineClusterCreateRequest) GetIsCombinedOk() (*bool, bool)`

GetIsCombinedOk returns a tuple with the IsCombined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCombined

`func (o *SearchEngineClusterCreateRequest) SetIsCombined(v bool)`

SetIsCombined sets IsCombined field to given value.

### HasIsCombined

`func (o *SearchEngineClusterCreateRequest) HasIsCombined() bool`

HasIsCombined returns a boolean if a field has been set.

### GetLicense

`func (o *SearchEngineClusterCreateRequest) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SearchEngineClusterCreateRequest) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SearchEngineClusterCreateRequest) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *SearchEngineClusterCreateRequest) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### SetLicenseNil

`func (o *SearchEngineClusterCreateRequest) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *SearchEngineClusterCreateRequest) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetMaintenanceOption

`func (o *SearchEngineClusterCreateRequest) GetMaintenanceOption() MaintenanceOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *SearchEngineClusterCreateRequest) GetMaintenanceOptionOk() (*MaintenanceOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *SearchEngineClusterCreateRequest) SetMaintenanceOption(v MaintenanceOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *SearchEngineClusterCreateRequest) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *SearchEngineClusterCreateRequest) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *SearchEngineClusterCreateRequest) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetName

`func (o *SearchEngineClusterCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchEngineClusterCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchEngineClusterCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *SearchEngineClusterCreateRequest) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *SearchEngineClusterCreateRequest) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *SearchEngineClusterCreateRequest) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *SearchEngineClusterCreateRequest) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### GetSubnetId

`func (o *SearchEngineClusterCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *SearchEngineClusterCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *SearchEngineClusterCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *SearchEngineClusterCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SearchEngineClusterCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SearchEngineClusterCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SearchEngineClusterCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SearchEngineClusterCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SearchEngineClusterCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTimezone

`func (o *SearchEngineClusterCreateRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SearchEngineClusterCreateRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SearchEngineClusterCreateRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


