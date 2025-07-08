# VpcPeeringRuleAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**DestinationCidrAvailabilities** | [**[]VpcPeeringRuleAvailabilityDetail**](VpcPeeringRuleAvailabilityDetail.md) |  | 

## Methods

### NewVpcPeeringRuleAvailabilityResponse

`func NewVpcPeeringRuleAvailabilityResponse(count int32, destinationCidrAvailabilities []VpcPeeringRuleAvailabilityDetail, ) *VpcPeeringRuleAvailabilityResponse`

NewVpcPeeringRuleAvailabilityResponse instantiates a new VpcPeeringRuleAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPeeringRuleAvailabilityResponseWithDefaults

`func NewVpcPeeringRuleAvailabilityResponseWithDefaults() *VpcPeeringRuleAvailabilityResponse`

NewVpcPeeringRuleAvailabilityResponseWithDefaults instantiates a new VpcPeeringRuleAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VpcPeeringRuleAvailabilityResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VpcPeeringRuleAvailabilityResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VpcPeeringRuleAvailabilityResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetDestinationCidrAvailabilities

`func (o *VpcPeeringRuleAvailabilityResponse) GetDestinationCidrAvailabilities() []VpcPeeringRuleAvailabilityDetail`

GetDestinationCidrAvailabilities returns the DestinationCidrAvailabilities field if non-nil, zero value otherwise.

### GetDestinationCidrAvailabilitiesOk

`func (o *VpcPeeringRuleAvailabilityResponse) GetDestinationCidrAvailabilitiesOk() (*[]VpcPeeringRuleAvailabilityDetail, bool)`

GetDestinationCidrAvailabilitiesOk returns a tuple with the DestinationCidrAvailabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidrAvailabilities

`func (o *VpcPeeringRuleAvailabilityResponse) SetDestinationCidrAvailabilities(v []VpcPeeringRuleAvailabilityDetail)`

SetDestinationCidrAvailabilities sets DestinationCidrAvailabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


