# RoutingRuleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description | [optional] [default to ""]
**DestinationCidr** | **string** | Destination CIDR | 
**DestinationResourceId** | Pointer to **NullableString** |  | [optional] 
**DestinationType** | [**DirectConnectRoutingRuleDestinationType**](DirectConnectRoutingRuleDestinationType.md) | Destination Type | 

## Methods

### NewRoutingRuleCreateRequest

`func NewRoutingRuleCreateRequest(destinationCidr string, destinationType DirectConnectRoutingRuleDestinationType, ) *RoutingRuleCreateRequest`

NewRoutingRuleCreateRequest instantiates a new RoutingRuleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRuleCreateRequestWithDefaults

`func NewRoutingRuleCreateRequestWithDefaults() *RoutingRuleCreateRequest`

NewRoutingRuleCreateRequestWithDefaults instantiates a new RoutingRuleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoutingRuleCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingRuleCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingRuleCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingRuleCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinationCidr

`func (o *RoutingRuleCreateRequest) GetDestinationCidr() string`

GetDestinationCidr returns the DestinationCidr field if non-nil, zero value otherwise.

### GetDestinationCidrOk

`func (o *RoutingRuleCreateRequest) GetDestinationCidrOk() (*string, bool)`

GetDestinationCidrOk returns a tuple with the DestinationCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidr

`func (o *RoutingRuleCreateRequest) SetDestinationCidr(v string)`

SetDestinationCidr sets DestinationCidr field to given value.


### GetDestinationResourceId

`func (o *RoutingRuleCreateRequest) GetDestinationResourceId() string`

GetDestinationResourceId returns the DestinationResourceId field if non-nil, zero value otherwise.

### GetDestinationResourceIdOk

`func (o *RoutingRuleCreateRequest) GetDestinationResourceIdOk() (*string, bool)`

GetDestinationResourceIdOk returns a tuple with the DestinationResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationResourceId

`func (o *RoutingRuleCreateRequest) SetDestinationResourceId(v string)`

SetDestinationResourceId sets DestinationResourceId field to given value.

### HasDestinationResourceId

`func (o *RoutingRuleCreateRequest) HasDestinationResourceId() bool`

HasDestinationResourceId returns a boolean if a field has been set.

### SetDestinationResourceIdNil

`func (o *RoutingRuleCreateRequest) SetDestinationResourceIdNil(b bool)`

 SetDestinationResourceIdNil sets the value for DestinationResourceId to be an explicit nil

### UnsetDestinationResourceId
`func (o *RoutingRuleCreateRequest) UnsetDestinationResourceId()`

UnsetDestinationResourceId ensures that no value is present for DestinationResourceId, not even an explicit nil
### GetDestinationType

`func (o *RoutingRuleCreateRequest) GetDestinationType() DirectConnectRoutingRuleDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *RoutingRuleCreateRequest) GetDestinationTypeOk() (*DirectConnectRoutingRuleDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *RoutingRuleCreateRequest) SetDestinationType(v DirectConnectRoutingRuleDestinationType)`

SetDestinationType sets DestinationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


