# TransitGatewayRuleAvailabilityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Transit Gateway Available | 
**ConnectedVpcId** | **string** | VPC ID Connected to Transit Gateway. | 
**DestinationCidr** | **string** | Destination CIDR | 
**DestinationType** | [**TransitGatewayRuleDestinationType**](TransitGatewayRuleDestinationType.md) | Destination Type | 
**Message** | **string** | Transit Gateway Message | 

## Methods

### NewTransitGatewayRuleAvailabilityDetail

`func NewTransitGatewayRuleAvailabilityDetail(available bool, connectedVpcId string, destinationCidr string, destinationType TransitGatewayRuleDestinationType, message string, ) *TransitGatewayRuleAvailabilityDetail`

NewTransitGatewayRuleAvailabilityDetail instantiates a new TransitGatewayRuleAvailabilityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayRuleAvailabilityDetailWithDefaults

`func NewTransitGatewayRuleAvailabilityDetailWithDefaults() *TransitGatewayRuleAvailabilityDetail`

NewTransitGatewayRuleAvailabilityDetailWithDefaults instantiates a new TransitGatewayRuleAvailabilityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *TransitGatewayRuleAvailabilityDetail) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TransitGatewayRuleAvailabilityDetail) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TransitGatewayRuleAvailabilityDetail) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetConnectedVpcId

`func (o *TransitGatewayRuleAvailabilityDetail) GetConnectedVpcId() string`

GetConnectedVpcId returns the ConnectedVpcId field if non-nil, zero value otherwise.

### GetConnectedVpcIdOk

`func (o *TransitGatewayRuleAvailabilityDetail) GetConnectedVpcIdOk() (*string, bool)`

GetConnectedVpcIdOk returns a tuple with the ConnectedVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedVpcId

`func (o *TransitGatewayRuleAvailabilityDetail) SetConnectedVpcId(v string)`

SetConnectedVpcId sets ConnectedVpcId field to given value.


### GetDestinationCidr

`func (o *TransitGatewayRuleAvailabilityDetail) GetDestinationCidr() string`

GetDestinationCidr returns the DestinationCidr field if non-nil, zero value otherwise.

### GetDestinationCidrOk

`func (o *TransitGatewayRuleAvailabilityDetail) GetDestinationCidrOk() (*string, bool)`

GetDestinationCidrOk returns a tuple with the DestinationCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidr

`func (o *TransitGatewayRuleAvailabilityDetail) SetDestinationCidr(v string)`

SetDestinationCidr sets DestinationCidr field to given value.


### GetDestinationType

`func (o *TransitGatewayRuleAvailabilityDetail) GetDestinationType() TransitGatewayRuleDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransitGatewayRuleAvailabilityDetail) GetDestinationTypeOk() (*TransitGatewayRuleDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransitGatewayRuleAvailabilityDetail) SetDestinationType(v TransitGatewayRuleDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMessage

`func (o *TransitGatewayRuleAvailabilityDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TransitGatewayRuleAvailabilityDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TransitGatewayRuleAvailabilityDetail) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


