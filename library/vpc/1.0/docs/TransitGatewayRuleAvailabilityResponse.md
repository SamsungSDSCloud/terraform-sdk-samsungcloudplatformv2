# TransitGatewayRuleAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**DestinationCidrAvailabilities** | [**[]TransitGatewayRuleAvailabilityDetail**](TransitGatewayRuleAvailabilityDetail.md) |  | 

## Methods

### NewTransitGatewayRuleAvailabilityResponse

`func NewTransitGatewayRuleAvailabilityResponse(count int32, destinationCidrAvailabilities []TransitGatewayRuleAvailabilityDetail, ) *TransitGatewayRuleAvailabilityResponse`

NewTransitGatewayRuleAvailabilityResponse instantiates a new TransitGatewayRuleAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayRuleAvailabilityResponseWithDefaults

`func NewTransitGatewayRuleAvailabilityResponseWithDefaults() *TransitGatewayRuleAvailabilityResponse`

NewTransitGatewayRuleAvailabilityResponseWithDefaults instantiates a new TransitGatewayRuleAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TransitGatewayRuleAvailabilityResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransitGatewayRuleAvailabilityResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransitGatewayRuleAvailabilityResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetDestinationCidrAvailabilities

`func (o *TransitGatewayRuleAvailabilityResponse) GetDestinationCidrAvailabilities() []TransitGatewayRuleAvailabilityDetail`

GetDestinationCidrAvailabilities returns the DestinationCidrAvailabilities field if non-nil, zero value otherwise.

### GetDestinationCidrAvailabilitiesOk

`func (o *TransitGatewayRuleAvailabilityResponse) GetDestinationCidrAvailabilitiesOk() (*[]TransitGatewayRuleAvailabilityDetail, bool)`

GetDestinationCidrAvailabilitiesOk returns a tuple with the DestinationCidrAvailabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidrAvailabilities

`func (o *TransitGatewayRuleAvailabilityResponse) SetDestinationCidrAvailabilities(v []TransitGatewayRuleAvailabilityDetail)`

SetDestinationCidrAvailabilities sets DestinationCidrAvailabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


