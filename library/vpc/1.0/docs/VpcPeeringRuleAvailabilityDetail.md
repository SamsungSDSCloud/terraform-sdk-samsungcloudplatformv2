# VpcPeeringRuleAvailabilityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | VPC Peering Available | 
**DestinationCidr** | **string** | Destination CIDR | 
**DestinationVpcType** | [**VpcPeeringRuleDestinationVpcType**](VpcPeeringRuleDestinationVpcType.md) | Destination VPC Type | 
**Message** | **string** | VPC Peering Message | 

## Methods

### NewVpcPeeringRuleAvailabilityDetail

`func NewVpcPeeringRuleAvailabilityDetail(available bool, destinationCidr string, destinationVpcType VpcPeeringRuleDestinationVpcType, message string, ) *VpcPeeringRuleAvailabilityDetail`

NewVpcPeeringRuleAvailabilityDetail instantiates a new VpcPeeringRuleAvailabilityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPeeringRuleAvailabilityDetailWithDefaults

`func NewVpcPeeringRuleAvailabilityDetailWithDefaults() *VpcPeeringRuleAvailabilityDetail`

NewVpcPeeringRuleAvailabilityDetailWithDefaults instantiates a new VpcPeeringRuleAvailabilityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *VpcPeeringRuleAvailabilityDetail) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *VpcPeeringRuleAvailabilityDetail) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *VpcPeeringRuleAvailabilityDetail) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetDestinationCidr

`func (o *VpcPeeringRuleAvailabilityDetail) GetDestinationCidr() string`

GetDestinationCidr returns the DestinationCidr field if non-nil, zero value otherwise.

### GetDestinationCidrOk

`func (o *VpcPeeringRuleAvailabilityDetail) GetDestinationCidrOk() (*string, bool)`

GetDestinationCidrOk returns a tuple with the DestinationCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidr

`func (o *VpcPeeringRuleAvailabilityDetail) SetDestinationCidr(v string)`

SetDestinationCidr sets DestinationCidr field to given value.


### GetDestinationVpcType

`func (o *VpcPeeringRuleAvailabilityDetail) GetDestinationVpcType() VpcPeeringRuleDestinationVpcType`

GetDestinationVpcType returns the DestinationVpcType field if non-nil, zero value otherwise.

### GetDestinationVpcTypeOk

`func (o *VpcPeeringRuleAvailabilityDetail) GetDestinationVpcTypeOk() (*VpcPeeringRuleDestinationVpcType, bool)`

GetDestinationVpcTypeOk returns a tuple with the DestinationVpcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVpcType

`func (o *VpcPeeringRuleAvailabilityDetail) SetDestinationVpcType(v VpcPeeringRuleDestinationVpcType)`

SetDestinationVpcType sets DestinationVpcType field to given value.


### GetMessage

`func (o *VpcPeeringRuleAvailabilityDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VpcPeeringRuleAvailabilityDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VpcPeeringRuleAvailabilityDetail) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


