# EventStreamsClusterCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AkhqEnabled** | Pointer to **bool** | AKHQ availability | [optional] [default to false]
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**DbaasEngineVersionId** | **string** | DBaaS engine version ID | 
**InitConfigOption** | [**EventStreamsInitConfigOptionRequest**](EventStreamsInitConfigOptionRequest.md) | DB initial configuration option | 
**InstanceGroups** | [**[]InstanceGroupRequest**](InstanceGroupRequest.md) | Instance groups list | 
**InstanceNamePrefix** | **string** | Instance name prefix | 
**IsCombined** | Pointer to **bool** | HA availability | [optional] [default to false]
**MaintenanceOption** | Pointer to [**NullableMaintenanceOption**](MaintenanceOption.md) |  | [optional] 
**Name** | **string** | Cluster name | 
**NatEnabled** | Pointer to **bool** | NAT availability | [optional] [default to false]
**SubnetId** | **string** | Subnet ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Timezone** | **string** | Timezone | 

## Methods

### NewEventStreamsClusterCreateRequest

`func NewEventStreamsClusterCreateRequest(dbaasEngineVersionId string, initConfigOption EventStreamsInitConfigOptionRequest, instanceGroups []InstanceGroupRequest, instanceNamePrefix string, name string, subnetId string, timezone string, ) *EventStreamsClusterCreateRequest`

NewEventStreamsClusterCreateRequest instantiates a new EventStreamsClusterCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStreamsClusterCreateRequestWithDefaults

`func NewEventStreamsClusterCreateRequestWithDefaults() *EventStreamsClusterCreateRequest`

NewEventStreamsClusterCreateRequestWithDefaults instantiates a new EventStreamsClusterCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAkhqEnabled

`func (o *EventStreamsClusterCreateRequest) GetAkhqEnabled() bool`

GetAkhqEnabled returns the AkhqEnabled field if non-nil, zero value otherwise.

### GetAkhqEnabledOk

`func (o *EventStreamsClusterCreateRequest) GetAkhqEnabledOk() (*bool, bool)`

GetAkhqEnabledOk returns a tuple with the AkhqEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkhqEnabled

`func (o *EventStreamsClusterCreateRequest) SetAkhqEnabled(v bool)`

SetAkhqEnabled sets AkhqEnabled field to given value.

### HasAkhqEnabled

`func (o *EventStreamsClusterCreateRequest) HasAkhqEnabled() bool`

HasAkhqEnabled returns a boolean if a field has been set.

### GetAllowableIpAddresses

`func (o *EventStreamsClusterCreateRequest) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *EventStreamsClusterCreateRequest) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *EventStreamsClusterCreateRequest) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *EventStreamsClusterCreateRequest) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetDbaasEngineVersionId

`func (o *EventStreamsClusterCreateRequest) GetDbaasEngineVersionId() string`

GetDbaasEngineVersionId returns the DbaasEngineVersionId field if non-nil, zero value otherwise.

### GetDbaasEngineVersionIdOk

`func (o *EventStreamsClusterCreateRequest) GetDbaasEngineVersionIdOk() (*string, bool)`

GetDbaasEngineVersionIdOk returns a tuple with the DbaasEngineVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngineVersionId

`func (o *EventStreamsClusterCreateRequest) SetDbaasEngineVersionId(v string)`

SetDbaasEngineVersionId sets DbaasEngineVersionId field to given value.


### GetInitConfigOption

`func (o *EventStreamsClusterCreateRequest) GetInitConfigOption() EventStreamsInitConfigOptionRequest`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *EventStreamsClusterCreateRequest) GetInitConfigOptionOk() (*EventStreamsInitConfigOptionRequest, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *EventStreamsClusterCreateRequest) SetInitConfigOption(v EventStreamsInitConfigOptionRequest)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceGroups

`func (o *EventStreamsClusterCreateRequest) GetInstanceGroups() []InstanceGroupRequest`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *EventStreamsClusterCreateRequest) GetInstanceGroupsOk() (*[]InstanceGroupRequest, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *EventStreamsClusterCreateRequest) SetInstanceGroups(v []InstanceGroupRequest)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetInstanceNamePrefix

`func (o *EventStreamsClusterCreateRequest) GetInstanceNamePrefix() string`

GetInstanceNamePrefix returns the InstanceNamePrefix field if non-nil, zero value otherwise.

### GetInstanceNamePrefixOk

`func (o *EventStreamsClusterCreateRequest) GetInstanceNamePrefixOk() (*string, bool)`

GetInstanceNamePrefixOk returns a tuple with the InstanceNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNamePrefix

`func (o *EventStreamsClusterCreateRequest) SetInstanceNamePrefix(v string)`

SetInstanceNamePrefix sets InstanceNamePrefix field to given value.


### GetIsCombined

`func (o *EventStreamsClusterCreateRequest) GetIsCombined() bool`

GetIsCombined returns the IsCombined field if non-nil, zero value otherwise.

### GetIsCombinedOk

`func (o *EventStreamsClusterCreateRequest) GetIsCombinedOk() (*bool, bool)`

GetIsCombinedOk returns a tuple with the IsCombined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCombined

`func (o *EventStreamsClusterCreateRequest) SetIsCombined(v bool)`

SetIsCombined sets IsCombined field to given value.

### HasIsCombined

`func (o *EventStreamsClusterCreateRequest) HasIsCombined() bool`

HasIsCombined returns a boolean if a field has been set.

### GetMaintenanceOption

`func (o *EventStreamsClusterCreateRequest) GetMaintenanceOption() MaintenanceOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *EventStreamsClusterCreateRequest) GetMaintenanceOptionOk() (*MaintenanceOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *EventStreamsClusterCreateRequest) SetMaintenanceOption(v MaintenanceOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *EventStreamsClusterCreateRequest) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *EventStreamsClusterCreateRequest) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *EventStreamsClusterCreateRequest) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetName

`func (o *EventStreamsClusterCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventStreamsClusterCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventStreamsClusterCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *EventStreamsClusterCreateRequest) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *EventStreamsClusterCreateRequest) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *EventStreamsClusterCreateRequest) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *EventStreamsClusterCreateRequest) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### GetSubnetId

`func (o *EventStreamsClusterCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *EventStreamsClusterCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *EventStreamsClusterCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *EventStreamsClusterCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EventStreamsClusterCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EventStreamsClusterCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EventStreamsClusterCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *EventStreamsClusterCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *EventStreamsClusterCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTimezone

`func (o *EventStreamsClusterCreateRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *EventStreamsClusterCreateRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *EventStreamsClusterCreateRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


