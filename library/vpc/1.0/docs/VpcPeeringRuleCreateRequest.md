# VpcPeeringRuleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationCidr** | **string** | Destination CIDR | 
**DestinationVpcType** | [**VpcPeeringRuleDestinationVpcType**](VpcPeeringRuleDestinationVpcType.md) | Destination VPC Type | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]

## Methods

### NewVpcPeeringRuleCreateRequest

`func NewVpcPeeringRuleCreateRequest(destinationCidr string, destinationVpcType VpcPeeringRuleDestinationVpcType, ) *VpcPeeringRuleCreateRequest`

NewVpcPeeringRuleCreateRequest instantiates a new VpcPeeringRuleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPeeringRuleCreateRequestWithDefaults

`func NewVpcPeeringRuleCreateRequestWithDefaults() *VpcPeeringRuleCreateRequest`

NewVpcPeeringRuleCreateRequestWithDefaults instantiates a new VpcPeeringRuleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationCidr

`func (o *VpcPeeringRuleCreateRequest) GetDestinationCidr() string`

GetDestinationCidr returns the DestinationCidr field if non-nil, zero value otherwise.

### GetDestinationCidrOk

`func (o *VpcPeeringRuleCreateRequest) GetDestinationCidrOk() (*string, bool)`

GetDestinationCidrOk returns a tuple with the DestinationCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidr

`func (o *VpcPeeringRuleCreateRequest) SetDestinationCidr(v string)`

SetDestinationCidr sets DestinationCidr field to given value.


### GetDestinationVpcType

`func (o *VpcPeeringRuleCreateRequest) GetDestinationVpcType() VpcPeeringRuleDestinationVpcType`

GetDestinationVpcType returns the DestinationVpcType field if non-nil, zero value otherwise.

### GetDestinationVpcTypeOk

`func (o *VpcPeeringRuleCreateRequest) GetDestinationVpcTypeOk() (*VpcPeeringRuleDestinationVpcType, bool)`

GetDestinationVpcTypeOk returns a tuple with the DestinationVpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVpcType

`func (o *VpcPeeringRuleCreateRequest) SetDestinationVpcType(v VpcPeeringRuleDestinationVpcType)`

SetDestinationVpcType sets DestinationVpcType field to given value.


### GetTags

`func (o *VpcPeeringRuleCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpcPeeringRuleCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpcPeeringRuleCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpcPeeringRuleCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


