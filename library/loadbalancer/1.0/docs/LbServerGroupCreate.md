# LbServerGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**LbHealthCheckId** | Pointer to **NullableString** |  | [optional] 
**LbMethod** | [**LbServerGroupLbMethod**](LbServerGroupLbMethod.md) | LB Method | 
**Name** | **string** | LB Health Check name | 
**Protocol** | [**LbServerGroupProtocol**](LbServerGroupProtocol.md) | Protocol | 
**SubnetId** | **string** | Service Subnet ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**VpcId** | **string** | VPC ID | 

## Methods

### NewLbServerGroupCreate

`func NewLbServerGroupCreate(lbMethod LbServerGroupLbMethod, name string, protocol LbServerGroupProtocol, subnetId string, vpcId string, ) *LbServerGroupCreate`

NewLbServerGroupCreate instantiates a new LbServerGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbServerGroupCreateWithDefaults

`func NewLbServerGroupCreateWithDefaults() *LbServerGroupCreate`

NewLbServerGroupCreateWithDefaults instantiates a new LbServerGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LbServerGroupCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LbServerGroupCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LbServerGroupCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LbServerGroupCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LbServerGroupCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LbServerGroupCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLbHealthCheckId

`func (o *LbServerGroupCreate) GetLbHealthCheckId() string`

GetLbHealthCheckId returns the LbHealthCheckId field if non-nil, zero value otherwise.

### GetLbHealthCheckIdOk

`func (o *LbServerGroupCreate) GetLbHealthCheckIdOk() (*string, bool)`

GetLbHealthCheckIdOk returns a tuple with the LbHealthCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbHealthCheckId

`func (o *LbServerGroupCreate) SetLbHealthCheckId(v string)`

SetLbHealthCheckId sets LbHealthCheckId field to given value.

### HasLbHealthCheckId

`func (o *LbServerGroupCreate) HasLbHealthCheckId() bool`

HasLbHealthCheckId returns a boolean if a field has been set.

### SetLbHealthCheckIdNil

`func (o *LbServerGroupCreate) SetLbHealthCheckIdNil(b bool)`

 SetLbHealthCheckIdNil sets the value for LbHealthCheckId to be an explicit nil

### UnsetLbHealthCheckId
`func (o *LbServerGroupCreate) UnsetLbHealthCheckId()`

UnsetLbHealthCheckId ensures that no value is present for LbHealthCheckId, not even an explicit nil
### GetLbMethod

`func (o *LbServerGroupCreate) GetLbMethod() LbServerGroupLbMethod`

GetLbMethod returns the LbMethod field if non-nil, zero value otherwise.

### GetLbMethodOk

`func (o *LbServerGroupCreate) GetLbMethodOk() (*LbServerGroupLbMethod, bool)`

GetLbMethodOk returns a tuple with the LbMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbMethod

`func (o *LbServerGroupCreate) SetLbMethod(v LbServerGroupLbMethod)`

SetLbMethod sets LbMethod field to given value.


### GetName

`func (o *LbServerGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LbServerGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LbServerGroupCreate) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *LbServerGroupCreate) GetProtocol() LbServerGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LbServerGroupCreate) GetProtocolOk() (*LbServerGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LbServerGroupCreate) SetProtocol(v LbServerGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetSubnetId

`func (o *LbServerGroupCreate) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *LbServerGroupCreate) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *LbServerGroupCreate) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *LbServerGroupCreate) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LbServerGroupCreate) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LbServerGroupCreate) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LbServerGroupCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVpcId

`func (o *LbServerGroupCreate) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *LbServerGroupCreate) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *LbServerGroupCreate) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


