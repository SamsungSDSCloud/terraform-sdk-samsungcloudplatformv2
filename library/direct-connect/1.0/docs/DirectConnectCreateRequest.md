# DirectConnectCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | **int32** | Port Bandwidth | 
**Description** | Pointer to **string** | Direct Connect Description | [optional] [default to ""]
**FirewallEnabled** | Pointer to **bool** | Firewall Use | [optional] [default to false]
**FirewallLoggable** | Pointer to **bool** | Firewall Loggable | [optional] [default to false]
**Name** | **string** | Direct Connect Name | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**VpcId** | **string** | VPC Id | 

## Methods

### NewDirectConnectCreateRequest

`func NewDirectConnectCreateRequest(bandwidth int32, name string, vpcId string, ) *DirectConnectCreateRequest`

NewDirectConnectCreateRequest instantiates a new DirectConnectCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectCreateRequestWithDefaults

`func NewDirectConnectCreateRequestWithDefaults() *DirectConnectCreateRequest`

NewDirectConnectCreateRequestWithDefaults instantiates a new DirectConnectCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *DirectConnectCreateRequest) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *DirectConnectCreateRequest) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *DirectConnectCreateRequest) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetDescription

`func (o *DirectConnectCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DirectConnectCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DirectConnectCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DirectConnectCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *DirectConnectCreateRequest) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *DirectConnectCreateRequest) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *DirectConnectCreateRequest) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *DirectConnectCreateRequest) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetFirewallLoggable

`func (o *DirectConnectCreateRequest) GetFirewallLoggable() bool`

GetFirewallLoggable returns the FirewallLoggable field if non-nil, zero value otherwise.

### GetFirewallLoggableOk

`func (o *DirectConnectCreateRequest) GetFirewallLoggableOk() (*bool, bool)`

GetFirewallLoggableOk returns a tuple with the FirewallLoggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallLoggable

`func (o *DirectConnectCreateRequest) SetFirewallLoggable(v bool)`

SetFirewallLoggable sets FirewallLoggable field to given value.

### HasFirewallLoggable

`func (o *DirectConnectCreateRequest) HasFirewallLoggable() bool`

HasFirewallLoggable returns a boolean if a field has been set.

### GetName

`func (o *DirectConnectCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectConnectCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectConnectCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *DirectConnectCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DirectConnectCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DirectConnectCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DirectConnectCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVpcId

`func (o *DirectConnectCreateRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *DirectConnectCreateRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *DirectConnectCreateRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


