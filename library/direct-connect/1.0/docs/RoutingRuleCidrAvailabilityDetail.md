# RoutingRuleCidrAvailabilityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Availability of destination CIDR | 
**DestinationCidr** | **string** | Destination CIDR | 
**DestinationType** | [**DirectConnectRoutingRuleDestinationType**](DirectConnectRoutingRuleDestinationType.md) | Destination Type | 

## Methods

### NewRoutingRuleCidrAvailabilityDetail

`func NewRoutingRuleCidrAvailabilityDetail(available bool, destinationCidr string, destinationType DirectConnectRoutingRuleDestinationType, ) *RoutingRuleCidrAvailabilityDetail`

NewRoutingRuleCidrAvailabilityDetail instantiates a new RoutingRuleCidrAvailabilityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRuleCidrAvailabilityDetailWithDefaults

`func NewRoutingRuleCidrAvailabilityDetailWithDefaults() *RoutingRuleCidrAvailabilityDetail`

NewRoutingRuleCidrAvailabilityDetailWithDefaults instantiates a new RoutingRuleCidrAvailabilityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *RoutingRuleCidrAvailabilityDetail) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *RoutingRuleCidrAvailabilityDetail) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *RoutingRuleCidrAvailabilityDetail) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetDestinationCidr

`func (o *RoutingRuleCidrAvailabilityDetail) GetDestinationCidr() string`

GetDestinationCidr returns the DestinationCidr field if non-nil, zero value otherwise.

### GetDestinationCidrOk

`func (o *RoutingRuleCidrAvailabilityDetail) GetDestinationCidrOk() (*string, bool)`

GetDestinationCidrOk returns a tuple with the DestinationCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidr

`func (o *RoutingRuleCidrAvailabilityDetail) SetDestinationCidr(v string)`

SetDestinationCidr sets DestinationCidr field to given value.


### GetDestinationType

`func (o *RoutingRuleCidrAvailabilityDetail) GetDestinationType() DirectConnectRoutingRuleDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *RoutingRuleCidrAvailabilityDetail) GetDestinationTypeOk() (*DirectConnectRoutingRuleDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *RoutingRuleCidrAvailabilityDetail) SetDestinationType(v DirectConnectRoutingRuleDestinationType)`

SetDestinationType sets DestinationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


