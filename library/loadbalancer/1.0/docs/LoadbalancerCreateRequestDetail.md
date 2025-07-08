# LoadbalancerCreateRequestDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**FirewallEnabled** | Pointer to **NullableBool** |  | [optional] 
**FirewallLoggingEnabled** | Pointer to **NullableBool** |  | [optional] 
**LayerType** | **string** | Layer type | 
**Name** | **string** | name | 
**PublicipId** | Pointer to **NullableString** |  | [optional] 
**ServiceIp** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | **string** | Subnet ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**VpcId** | **string** | VPC ID | 

## Methods

### NewLoadbalancerCreateRequestDetail

`func NewLoadbalancerCreateRequestDetail(layerType string, name string, subnetId string, vpcId string, ) *LoadbalancerCreateRequestDetail`

NewLoadbalancerCreateRequestDetail instantiates a new LoadbalancerCreateRequestDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerCreateRequestDetailWithDefaults

`func NewLoadbalancerCreateRequestDetailWithDefaults() *LoadbalancerCreateRequestDetail`

NewLoadbalancerCreateRequestDetailWithDefaults instantiates a new LoadbalancerCreateRequestDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LoadbalancerCreateRequestDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadbalancerCreateRequestDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadbalancerCreateRequestDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadbalancerCreateRequestDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LoadbalancerCreateRequestDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LoadbalancerCreateRequestDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFirewallEnabled

`func (o *LoadbalancerCreateRequestDetail) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *LoadbalancerCreateRequestDetail) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *LoadbalancerCreateRequestDetail) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *LoadbalancerCreateRequestDetail) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### SetFirewallEnabledNil

`func (o *LoadbalancerCreateRequestDetail) SetFirewallEnabledNil(b bool)`

 SetFirewallEnabledNil sets the value for FirewallEnabled to be an explicit nil

### UnsetFirewallEnabled
`func (o *LoadbalancerCreateRequestDetail) UnsetFirewallEnabled()`

UnsetFirewallEnabled ensures that no value is present for FirewallEnabled, not even an explicit nil
### GetFirewallLoggingEnabled

`func (o *LoadbalancerCreateRequestDetail) GetFirewallLoggingEnabled() bool`

GetFirewallLoggingEnabled returns the FirewallLoggingEnabled field if non-nil, zero value otherwise.

### GetFirewallLoggingEnabledOk

`func (o *LoadbalancerCreateRequestDetail) GetFirewallLoggingEnabledOk() (*bool, bool)`

GetFirewallLoggingEnabledOk returns a tuple with the FirewallLoggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallLoggingEnabled

`func (o *LoadbalancerCreateRequestDetail) SetFirewallLoggingEnabled(v bool)`

SetFirewallLoggingEnabled sets FirewallLoggingEnabled field to given value.

### HasFirewallLoggingEnabled

`func (o *LoadbalancerCreateRequestDetail) HasFirewallLoggingEnabled() bool`

HasFirewallLoggingEnabled returns a boolean if a field has been set.

### SetFirewallLoggingEnabledNil

`func (o *LoadbalancerCreateRequestDetail) SetFirewallLoggingEnabledNil(b bool)`

 SetFirewallLoggingEnabledNil sets the value for FirewallLoggingEnabled to be an explicit nil

### UnsetFirewallLoggingEnabled
`func (o *LoadbalancerCreateRequestDetail) UnsetFirewallLoggingEnabled()`

UnsetFirewallLoggingEnabled ensures that no value is present for FirewallLoggingEnabled, not even an explicit nil
### GetLayerType

`func (o *LoadbalancerCreateRequestDetail) GetLayerType() string`

GetLayerType returns the LayerType field if non-nil, zero value otherwise.

### GetLayerTypeOk

`func (o *LoadbalancerCreateRequestDetail) GetLayerTypeOk() (*string, bool)`

GetLayerTypeOk returns a tuple with the LayerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerType

`func (o *LoadbalancerCreateRequestDetail) SetLayerType(v string)`

SetLayerType sets LayerType field to given value.


### GetName

`func (o *LoadbalancerCreateRequestDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadbalancerCreateRequestDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadbalancerCreateRequestDetail) SetName(v string)`

SetName sets Name field to given value.


### GetPublicipId

`func (o *LoadbalancerCreateRequestDetail) GetPublicipId() string`

GetPublicipId returns the PublicipId field if non-nil, zero value otherwise.

### GetPublicipIdOk

`func (o *LoadbalancerCreateRequestDetail) GetPublicipIdOk() (*string, bool)`

GetPublicipIdOk returns a tuple with the PublicipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicipId

`func (o *LoadbalancerCreateRequestDetail) SetPublicipId(v string)`

SetPublicipId sets PublicipId field to given value.

### HasPublicipId

`func (o *LoadbalancerCreateRequestDetail) HasPublicipId() bool`

HasPublicipId returns a boolean if a field has been set.

### SetPublicipIdNil

`func (o *LoadbalancerCreateRequestDetail) SetPublicipIdNil(b bool)`

 SetPublicipIdNil sets the value for PublicipId to be an explicit nil

### UnsetPublicipId
`func (o *LoadbalancerCreateRequestDetail) UnsetPublicipId()`

UnsetPublicipId ensures that no value is present for PublicipId, not even an explicit nil
### GetServiceIp

`func (o *LoadbalancerCreateRequestDetail) GetServiceIp() string`

GetServiceIp returns the ServiceIp field if non-nil, zero value otherwise.

### GetServiceIpOk

`func (o *LoadbalancerCreateRequestDetail) GetServiceIpOk() (*string, bool)`

GetServiceIpOk returns a tuple with the ServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIp

`func (o *LoadbalancerCreateRequestDetail) SetServiceIp(v string)`

SetServiceIp sets ServiceIp field to given value.

### HasServiceIp

`func (o *LoadbalancerCreateRequestDetail) HasServiceIp() bool`

HasServiceIp returns a boolean if a field has been set.

### SetServiceIpNil

`func (o *LoadbalancerCreateRequestDetail) SetServiceIpNil(b bool)`

 SetServiceIpNil sets the value for ServiceIp to be an explicit nil

### UnsetServiceIp
`func (o *LoadbalancerCreateRequestDetail) UnsetServiceIp()`

UnsetServiceIp ensures that no value is present for ServiceIp, not even an explicit nil
### GetSubnetId

`func (o *LoadbalancerCreateRequestDetail) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *LoadbalancerCreateRequestDetail) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *LoadbalancerCreateRequestDetail) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *LoadbalancerCreateRequestDetail) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoadbalancerCreateRequestDetail) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LoadbalancerCreateRequestDetail) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LoadbalancerCreateRequestDetail) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVpcId

`func (o *LoadbalancerCreateRequestDetail) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *LoadbalancerCreateRequestDetail) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *LoadbalancerCreateRequestDetail) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


