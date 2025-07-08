# VpcPeeringAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**VpcPeeringAvailability** | [**[]VpcPeeringAvailabilityDetail**](VpcPeeringAvailabilityDetail.md) |  | 

## Methods

### NewVpcPeeringAvailabilityResponse

`func NewVpcPeeringAvailabilityResponse(count int32, vpcPeeringAvailability []VpcPeeringAvailabilityDetail, ) *VpcPeeringAvailabilityResponse`

NewVpcPeeringAvailabilityResponse instantiates a new VpcPeeringAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPeeringAvailabilityResponseWithDefaults

`func NewVpcPeeringAvailabilityResponseWithDefaults() *VpcPeeringAvailabilityResponse`

NewVpcPeeringAvailabilityResponseWithDefaults instantiates a new VpcPeeringAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VpcPeeringAvailabilityResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VpcPeeringAvailabilityResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VpcPeeringAvailabilityResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetVpcPeeringAvailability

`func (o *VpcPeeringAvailabilityResponse) GetVpcPeeringAvailability() []VpcPeeringAvailabilityDetail`

GetVpcPeeringAvailability returns the VpcPeeringAvailability field if non-nil, zero value otherwise.

### GetVpcPeeringAvailabilityOk

`func (o *VpcPeeringAvailabilityResponse) GetVpcPeeringAvailabilityOk() (*[]VpcPeeringAvailabilityDetail, bool)`

GetVpcPeeringAvailabilityOk returns a tuple with the VpcPeeringAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeeringAvailability

`func (o *VpcPeeringAvailabilityResponse) SetVpcPeeringAvailability(v []VpcPeeringAvailabilityDetail)`

SetVpcPeeringAvailability sets VpcPeeringAvailability field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


