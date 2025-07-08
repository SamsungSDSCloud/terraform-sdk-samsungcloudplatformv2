# PortIpAvailabilityListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**PortIpAvailabilities** | [**[]PortIpAvailabilityResponse**](PortIpAvailabilityResponse.md) |  | 

## Methods

### NewPortIpAvailabilityListResponse

`func NewPortIpAvailabilityListResponse(portIpAvailabilities []PortIpAvailabilityResponse, ) *PortIpAvailabilityListResponse`

NewPortIpAvailabilityListResponse instantiates a new PortIpAvailabilityListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortIpAvailabilityListResponseWithDefaults

`func NewPortIpAvailabilityListResponseWithDefaults() *PortIpAvailabilityListResponse`

NewPortIpAvailabilityListResponseWithDefaults instantiates a new PortIpAvailabilityListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PortIpAvailabilityListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PortIpAvailabilityListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PortIpAvailabilityListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PortIpAvailabilityListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PortIpAvailabilityListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PortIpAvailabilityListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *PortIpAvailabilityListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PortIpAvailabilityListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PortIpAvailabilityListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PortIpAvailabilityListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *PortIpAvailabilityListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *PortIpAvailabilityListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPortIpAvailabilities

`func (o *PortIpAvailabilityListResponse) GetPortIpAvailabilities() []PortIpAvailabilityResponse`

GetPortIpAvailabilities returns the PortIpAvailabilities field if non-nil, zero value otherwise.

### GetPortIpAvailabilitiesOk

`func (o *PortIpAvailabilityListResponse) GetPortIpAvailabilitiesOk() (*[]PortIpAvailabilityResponse, bool)`

GetPortIpAvailabilitiesOk returns a tuple with the PortIpAvailabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIpAvailabilities

`func (o *PortIpAvailabilityListResponse) SetPortIpAvailabilities(v []PortIpAvailabilityResponse)`

SetPortIpAvailabilities sets PortIpAvailabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


