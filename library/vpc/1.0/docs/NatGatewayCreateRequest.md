# NatGatewayCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | NAT Gateway Description | [optional] [default to ""]
**PublicipId** | **string** | PublicIP ID | 
**SubnetId** | **string** | Subnet CIDR | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]

## Methods

### NewNatGatewayCreateRequest

`func NewNatGatewayCreateRequest(publicipId string, subnetId string, ) *NatGatewayCreateRequest`

NewNatGatewayCreateRequest instantiates a new NatGatewayCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatGatewayCreateRequestWithDefaults

`func NewNatGatewayCreateRequestWithDefaults() *NatGatewayCreateRequest`

NewNatGatewayCreateRequestWithDefaults instantiates a new NatGatewayCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *NatGatewayCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NatGatewayCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NatGatewayCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NatGatewayCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublicipId

`func (o *NatGatewayCreateRequest) GetPublicipId() string`

GetPublicipId returns the PublicipId field if non-nil, zero value otherwise.

### GetPublicipIdOk

`func (o *NatGatewayCreateRequest) GetPublicipIdOk() (*string, bool)`

GetPublicipIdOk returns a tuple with the PublicipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicipId

`func (o *NatGatewayCreateRequest) SetPublicipId(v string)`

SetPublicipId sets PublicipId field to given value.


### GetSubnetId

`func (o *NatGatewayCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *NatGatewayCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *NatGatewayCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *NatGatewayCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NatGatewayCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NatGatewayCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NatGatewayCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


