# TransitGatewayRuleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description | [optional] [default to ""]
**DestinationCidr** | **string** | Destination CIDR | 
**DestinationType** | [**TransitGatewayRuleDestinationType**](TransitGatewayRuleDestinationType.md) | Destination Type | 
**TgwConnectionVpcId** | **string** | VPC ID Connected to Transit Gateway. | 

## Methods

### NewTransitGatewayRuleCreateRequest

`func NewTransitGatewayRuleCreateRequest(destinationCidr string, destinationType TransitGatewayRuleDestinationType, tgwConnectionVpcId string, ) *TransitGatewayRuleCreateRequest`

NewTransitGatewayRuleCreateRequest instantiates a new TransitGatewayRuleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayRuleCreateRequestWithDefaults

`func NewTransitGatewayRuleCreateRequestWithDefaults() *TransitGatewayRuleCreateRequest`

NewTransitGatewayRuleCreateRequestWithDefaults instantiates a new TransitGatewayRuleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TransitGatewayRuleCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransitGatewayRuleCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransitGatewayRuleCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransitGatewayRuleCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinationCidr

`func (o *TransitGatewayRuleCreateRequest) GetDestinationCidr() string`

GetDestinationCidr returns the DestinationCidr field if non-nil, zero value otherwise.

### GetDestinationCidrOk

`func (o *TransitGatewayRuleCreateRequest) GetDestinationCidrOk() (*string, bool)`

GetDestinationCidrOk returns a tuple with the DestinationCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidr

`func (o *TransitGatewayRuleCreateRequest) SetDestinationCidr(v string)`

SetDestinationCidr sets DestinationCidr field to given value.


### GetDestinationType

`func (o *TransitGatewayRuleCreateRequest) GetDestinationType() TransitGatewayRuleDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *TransitGatewayRuleCreateRequest) GetDestinationTypeOk() (*TransitGatewayRuleDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *TransitGatewayRuleCreateRequest) SetDestinationType(v TransitGatewayRuleDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetTgwConnectionVpcId

`func (o *TransitGatewayRuleCreateRequest) GetTgwConnectionVpcId() string`

GetTgwConnectionVpcId returns the TgwConnectionVpcId field if non-nil, zero value otherwise.

### GetTgwConnectionVpcIdOk

`func (o *TransitGatewayRuleCreateRequest) GetTgwConnectionVpcIdOk() (*string, bool)`

GetTgwConnectionVpcIdOk returns a tuple with the TgwConnectionVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwConnectionVpcId

`func (o *TransitGatewayRuleCreateRequest) SetTgwConnectionVpcId(v string)`

SetTgwConnectionVpcId sets TgwConnectionVpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


