# PortCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**FixedIpAddress** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Port Name | 
**SecurityGroups** | Pointer to **[]string** | Security Group List | [optional] [default to []]
**SubnetId** | **string** | Subnet Id | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]

## Methods

### NewPortCreateRequest

`func NewPortCreateRequest(name string, subnetId string, ) *PortCreateRequest`

NewPortCreateRequest instantiates a new PortCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortCreateRequestWithDefaults

`func NewPortCreateRequestWithDefaults() *PortCreateRequest`

NewPortCreateRequestWithDefaults instantiates a new PortCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PortCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PortCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PortCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PortCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PortCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PortCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFixedIpAddress

`func (o *PortCreateRequest) GetFixedIpAddress() string`

GetFixedIpAddress returns the FixedIpAddress field if non-nil, zero value otherwise.

### GetFixedIpAddressOk

`func (o *PortCreateRequest) GetFixedIpAddressOk() (*string, bool)`

GetFixedIpAddressOk returns a tuple with the FixedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAddress

`func (o *PortCreateRequest) SetFixedIpAddress(v string)`

SetFixedIpAddress sets FixedIpAddress field to given value.

### HasFixedIpAddress

`func (o *PortCreateRequest) HasFixedIpAddress() bool`

HasFixedIpAddress returns a boolean if a field has been set.

### SetFixedIpAddressNil

`func (o *PortCreateRequest) SetFixedIpAddressNil(b bool)`

 SetFixedIpAddressNil sets the value for FixedIpAddress to be an explicit nil

### UnsetFixedIpAddress
`func (o *PortCreateRequest) UnsetFixedIpAddress()`

UnsetFixedIpAddress ensures that no value is present for FixedIpAddress, not even an explicit nil
### GetName

`func (o *PortCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSecurityGroups

`func (o *PortCreateRequest) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *PortCreateRequest) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *PortCreateRequest) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *PortCreateRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSubnetId

`func (o *PortCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *PortCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *PortCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *PortCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PortCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PortCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PortCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


