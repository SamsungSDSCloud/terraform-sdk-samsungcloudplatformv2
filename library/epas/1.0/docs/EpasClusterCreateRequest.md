# EpasClusterCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**DbaasEngineVersionId** | **string** | DBaaS engine version ID | 
**HaEnabled** | Pointer to **bool** | HA availability | [optional] [default to false]
**InitConfigOption** | [**EpasInitConfigOptionRequest**](EpasInitConfigOptionRequest.md) | DB initial configuration option | 
**InstanceGroups** | [**[]InstanceGroupRequest**](InstanceGroupRequest.md) | Instance groups list | 
**InstanceNamePrefix** | **string** | Instance name prefix | 
**MaintenanceOption** | Pointer to [**NullableMaintenanceOption**](MaintenanceOption.md) |  | [optional] 
**Name** | **string** | Cluster name | 
**NatEnabled** | Pointer to **bool** | NAT availability | [optional] [default to false]
**OriginClusterId** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | **string** | Subnet ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Timezone** | **string** | Timezone | 
**VipPublicIpId** | Pointer to **NullableString** |  | [optional] 
**VirtualIpAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEpasClusterCreateRequest

`func NewEpasClusterCreateRequest(dbaasEngineVersionId string, initConfigOption EpasInitConfigOptionRequest, instanceGroups []InstanceGroupRequest, instanceNamePrefix string, name string, subnetId string, timezone string, ) *EpasClusterCreateRequest`

NewEpasClusterCreateRequest instantiates a new EpasClusterCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpasClusterCreateRequestWithDefaults

`func NewEpasClusterCreateRequestWithDefaults() *EpasClusterCreateRequest`

NewEpasClusterCreateRequestWithDefaults instantiates a new EpasClusterCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowableIpAddresses

`func (o *EpasClusterCreateRequest) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *EpasClusterCreateRequest) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *EpasClusterCreateRequest) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *EpasClusterCreateRequest) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetDbaasEngineVersionId

`func (o *EpasClusterCreateRequest) GetDbaasEngineVersionId() string`

GetDbaasEngineVersionId returns the DbaasEngineVersionId field if non-nil, zero value otherwise.

### GetDbaasEngineVersionIdOk

`func (o *EpasClusterCreateRequest) GetDbaasEngineVersionIdOk() (*string, bool)`

GetDbaasEngineVersionIdOk returns a tuple with the DbaasEngineVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngineVersionId

`func (o *EpasClusterCreateRequest) SetDbaasEngineVersionId(v string)`

SetDbaasEngineVersionId sets DbaasEngineVersionId field to given value.


### GetHaEnabled

`func (o *EpasClusterCreateRequest) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *EpasClusterCreateRequest) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *EpasClusterCreateRequest) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *EpasClusterCreateRequest) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetInitConfigOption

`func (o *EpasClusterCreateRequest) GetInitConfigOption() EpasInitConfigOptionRequest`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *EpasClusterCreateRequest) GetInitConfigOptionOk() (*EpasInitConfigOptionRequest, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *EpasClusterCreateRequest) SetInitConfigOption(v EpasInitConfigOptionRequest)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceGroups

`func (o *EpasClusterCreateRequest) GetInstanceGroups() []InstanceGroupRequest`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *EpasClusterCreateRequest) GetInstanceGroupsOk() (*[]InstanceGroupRequest, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *EpasClusterCreateRequest) SetInstanceGroups(v []InstanceGroupRequest)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetInstanceNamePrefix

`func (o *EpasClusterCreateRequest) GetInstanceNamePrefix() string`

GetInstanceNamePrefix returns the InstanceNamePrefix field if non-nil, zero value otherwise.

### GetInstanceNamePrefixOk

`func (o *EpasClusterCreateRequest) GetInstanceNamePrefixOk() (*string, bool)`

GetInstanceNamePrefixOk returns a tuple with the InstanceNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNamePrefix

`func (o *EpasClusterCreateRequest) SetInstanceNamePrefix(v string)`

SetInstanceNamePrefix sets InstanceNamePrefix field to given value.


### GetMaintenanceOption

`func (o *EpasClusterCreateRequest) GetMaintenanceOption() MaintenanceOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *EpasClusterCreateRequest) GetMaintenanceOptionOk() (*MaintenanceOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *EpasClusterCreateRequest) SetMaintenanceOption(v MaintenanceOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *EpasClusterCreateRequest) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *EpasClusterCreateRequest) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *EpasClusterCreateRequest) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetName

`func (o *EpasClusterCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EpasClusterCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EpasClusterCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *EpasClusterCreateRequest) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *EpasClusterCreateRequest) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *EpasClusterCreateRequest) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *EpasClusterCreateRequest) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### GetOriginClusterId

`func (o *EpasClusterCreateRequest) GetOriginClusterId() string`

GetOriginClusterId returns the OriginClusterId field if non-nil, zero value otherwise.

### GetOriginClusterIdOk

`func (o *EpasClusterCreateRequest) GetOriginClusterIdOk() (*string, bool)`

GetOriginClusterIdOk returns a tuple with the OriginClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginClusterId

`func (o *EpasClusterCreateRequest) SetOriginClusterId(v string)`

SetOriginClusterId sets OriginClusterId field to given value.

### HasOriginClusterId

`func (o *EpasClusterCreateRequest) HasOriginClusterId() bool`

HasOriginClusterId returns a boolean if a field has been set.

### SetOriginClusterIdNil

`func (o *EpasClusterCreateRequest) SetOriginClusterIdNil(b bool)`

 SetOriginClusterIdNil sets the value for OriginClusterId to be an explicit nil

### UnsetOriginClusterId
`func (o *EpasClusterCreateRequest) UnsetOriginClusterId()`

UnsetOriginClusterId ensures that no value is present for OriginClusterId, not even an explicit nil
### GetSubnetId

`func (o *EpasClusterCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *EpasClusterCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *EpasClusterCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *EpasClusterCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EpasClusterCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EpasClusterCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EpasClusterCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *EpasClusterCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *EpasClusterCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTimezone

`func (o *EpasClusterCreateRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *EpasClusterCreateRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *EpasClusterCreateRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetVipPublicIpId

`func (o *EpasClusterCreateRequest) GetVipPublicIpId() string`

GetVipPublicIpId returns the VipPublicIpId field if non-nil, zero value otherwise.

### GetVipPublicIpIdOk

`func (o *EpasClusterCreateRequest) GetVipPublicIpIdOk() (*string, bool)`

GetVipPublicIpIdOk returns a tuple with the VipPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPublicIpId

`func (o *EpasClusterCreateRequest) SetVipPublicIpId(v string)`

SetVipPublicIpId sets VipPublicIpId field to given value.

### HasVipPublicIpId

`func (o *EpasClusterCreateRequest) HasVipPublicIpId() bool`

HasVipPublicIpId returns a boolean if a field has been set.

### SetVipPublicIpIdNil

`func (o *EpasClusterCreateRequest) SetVipPublicIpIdNil(b bool)`

 SetVipPublicIpIdNil sets the value for VipPublicIpId to be an explicit nil

### UnsetVipPublicIpId
`func (o *EpasClusterCreateRequest) UnsetVipPublicIpId()`

UnsetVipPublicIpId ensures that no value is present for VipPublicIpId, not even an explicit nil
### GetVirtualIpAddress

`func (o *EpasClusterCreateRequest) GetVirtualIpAddress() string`

GetVirtualIpAddress returns the VirtualIpAddress field if non-nil, zero value otherwise.

### GetVirtualIpAddressOk

`func (o *EpasClusterCreateRequest) GetVirtualIpAddressOk() (*string, bool)`

GetVirtualIpAddressOk returns a tuple with the VirtualIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIpAddress

`func (o *EpasClusterCreateRequest) SetVirtualIpAddress(v string)`

SetVirtualIpAddress sets VirtualIpAddress field to given value.

### HasVirtualIpAddress

`func (o *EpasClusterCreateRequest) HasVirtualIpAddress() bool`

HasVirtualIpAddress returns a boolean if a field has been set.

### SetVirtualIpAddressNil

`func (o *EpasClusterCreateRequest) SetVirtualIpAddressNil(b bool)`

 SetVirtualIpAddressNil sets the value for VirtualIpAddress to be an explicit nil

### UnsetVirtualIpAddress
`func (o *EpasClusterCreateRequest) UnsetVirtualIpAddress()`

UnsetVirtualIpAddress ensures that no value is present for VirtualIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


