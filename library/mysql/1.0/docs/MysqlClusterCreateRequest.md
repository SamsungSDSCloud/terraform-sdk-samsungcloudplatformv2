# MysqlClusterCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**DbaasEngineVersionId** | **string** | DBaaS engine version ID | 
**HaEnabled** | Pointer to **bool** | HA availability | [optional] [default to false]
**InitConfigOption** | [**MysqlInitConfigOptionRequest**](MysqlInitConfigOptionRequest.md) | DB initial configuration option | 
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

### NewMysqlClusterCreateRequest

`func NewMysqlClusterCreateRequest(dbaasEngineVersionId string, initConfigOption MysqlInitConfigOptionRequest, instanceGroups []InstanceGroupRequest, instanceNamePrefix string, name string, subnetId string, timezone string, ) *MysqlClusterCreateRequest`

NewMysqlClusterCreateRequest instantiates a new MysqlClusterCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlClusterCreateRequestWithDefaults

`func NewMysqlClusterCreateRequestWithDefaults() *MysqlClusterCreateRequest`

NewMysqlClusterCreateRequestWithDefaults instantiates a new MysqlClusterCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowableIpAddresses

`func (o *MysqlClusterCreateRequest) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *MysqlClusterCreateRequest) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *MysqlClusterCreateRequest) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *MysqlClusterCreateRequest) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetDbaasEngineVersionId

`func (o *MysqlClusterCreateRequest) GetDbaasEngineVersionId() string`

GetDbaasEngineVersionId returns the DbaasEngineVersionId field if non-nil, zero value otherwise.

### GetDbaasEngineVersionIdOk

`func (o *MysqlClusterCreateRequest) GetDbaasEngineVersionIdOk() (*string, bool)`

GetDbaasEngineVersionIdOk returns a tuple with the DbaasEngineVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngineVersionId

`func (o *MysqlClusterCreateRequest) SetDbaasEngineVersionId(v string)`

SetDbaasEngineVersionId sets DbaasEngineVersionId field to given value.


### GetHaEnabled

`func (o *MysqlClusterCreateRequest) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *MysqlClusterCreateRequest) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *MysqlClusterCreateRequest) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *MysqlClusterCreateRequest) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetInitConfigOption

`func (o *MysqlClusterCreateRequest) GetInitConfigOption() MysqlInitConfigOptionRequest`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *MysqlClusterCreateRequest) GetInitConfigOptionOk() (*MysqlInitConfigOptionRequest, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *MysqlClusterCreateRequest) SetInitConfigOption(v MysqlInitConfigOptionRequest)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceGroups

`func (o *MysqlClusterCreateRequest) GetInstanceGroups() []InstanceGroupRequest`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *MysqlClusterCreateRequest) GetInstanceGroupsOk() (*[]InstanceGroupRequest, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *MysqlClusterCreateRequest) SetInstanceGroups(v []InstanceGroupRequest)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetInstanceNamePrefix

`func (o *MysqlClusterCreateRequest) GetInstanceNamePrefix() string`

GetInstanceNamePrefix returns the InstanceNamePrefix field if non-nil, zero value otherwise.

### GetInstanceNamePrefixOk

`func (o *MysqlClusterCreateRequest) GetInstanceNamePrefixOk() (*string, bool)`

GetInstanceNamePrefixOk returns a tuple with the InstanceNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNamePrefix

`func (o *MysqlClusterCreateRequest) SetInstanceNamePrefix(v string)`

SetInstanceNamePrefix sets InstanceNamePrefix field to given value.


### GetMaintenanceOption

`func (o *MysqlClusterCreateRequest) GetMaintenanceOption() MaintenanceOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *MysqlClusterCreateRequest) GetMaintenanceOptionOk() (*MaintenanceOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *MysqlClusterCreateRequest) SetMaintenanceOption(v MaintenanceOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *MysqlClusterCreateRequest) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *MysqlClusterCreateRequest) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *MysqlClusterCreateRequest) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetName

`func (o *MysqlClusterCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlClusterCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlClusterCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *MysqlClusterCreateRequest) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *MysqlClusterCreateRequest) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *MysqlClusterCreateRequest) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *MysqlClusterCreateRequest) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### GetOriginClusterId

`func (o *MysqlClusterCreateRequest) GetOriginClusterId() string`

GetOriginClusterId returns the OriginClusterId field if non-nil, zero value otherwise.

### GetOriginClusterIdOk

`func (o *MysqlClusterCreateRequest) GetOriginClusterIdOk() (*string, bool)`

GetOriginClusterIdOk returns a tuple with the OriginClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginClusterId

`func (o *MysqlClusterCreateRequest) SetOriginClusterId(v string)`

SetOriginClusterId sets OriginClusterId field to given value.

### HasOriginClusterId

`func (o *MysqlClusterCreateRequest) HasOriginClusterId() bool`

HasOriginClusterId returns a boolean if a field has been set.

### SetOriginClusterIdNil

`func (o *MysqlClusterCreateRequest) SetOriginClusterIdNil(b bool)`

 SetOriginClusterIdNil sets the value for OriginClusterId to be an explicit nil

### UnsetOriginClusterId
`func (o *MysqlClusterCreateRequest) UnsetOriginClusterId()`

UnsetOriginClusterId ensures that no value is present for OriginClusterId, not even an explicit nil
### GetSubnetId

`func (o *MysqlClusterCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *MysqlClusterCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *MysqlClusterCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *MysqlClusterCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MysqlClusterCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MysqlClusterCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MysqlClusterCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *MysqlClusterCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *MysqlClusterCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTimezone

`func (o *MysqlClusterCreateRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MysqlClusterCreateRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MysqlClusterCreateRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetVipPublicIpId

`func (o *MysqlClusterCreateRequest) GetVipPublicIpId() string`

GetVipPublicIpId returns the VipPublicIpId field if non-nil, zero value otherwise.

### GetVipPublicIpIdOk

`func (o *MysqlClusterCreateRequest) GetVipPublicIpIdOk() (*string, bool)`

GetVipPublicIpIdOk returns a tuple with the VipPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPublicIpId

`func (o *MysqlClusterCreateRequest) SetVipPublicIpId(v string)`

SetVipPublicIpId sets VipPublicIpId field to given value.

### HasVipPublicIpId

`func (o *MysqlClusterCreateRequest) HasVipPublicIpId() bool`

HasVipPublicIpId returns a boolean if a field has been set.

### SetVipPublicIpIdNil

`func (o *MysqlClusterCreateRequest) SetVipPublicIpIdNil(b bool)`

 SetVipPublicIpIdNil sets the value for VipPublicIpId to be an explicit nil

### UnsetVipPublicIpId
`func (o *MysqlClusterCreateRequest) UnsetVipPublicIpId()`

UnsetVipPublicIpId ensures that no value is present for VipPublicIpId, not even an explicit nil
### GetVirtualIpAddress

`func (o *MysqlClusterCreateRequest) GetVirtualIpAddress() string`

GetVirtualIpAddress returns the VirtualIpAddress field if non-nil, zero value otherwise.

### GetVirtualIpAddressOk

`func (o *MysqlClusterCreateRequest) GetVirtualIpAddressOk() (*string, bool)`

GetVirtualIpAddressOk returns a tuple with the VirtualIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIpAddress

`func (o *MysqlClusterCreateRequest) SetVirtualIpAddress(v string)`

SetVirtualIpAddress sets VirtualIpAddress field to given value.

### HasVirtualIpAddress

`func (o *MysqlClusterCreateRequest) HasVirtualIpAddress() bool`

HasVirtualIpAddress returns a boolean if a field has been set.

### SetVirtualIpAddressNil

`func (o *MysqlClusterCreateRequest) SetVirtualIpAddressNil(b bool)`

 SetVirtualIpAddressNil sets the value for VirtualIpAddress to be an explicit nil

### UnsetVirtualIpAddress
`func (o *MysqlClusterCreateRequest) UnsetVirtualIpAddress()`

UnsetVirtualIpAddress ensures that no value is present for VirtualIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


