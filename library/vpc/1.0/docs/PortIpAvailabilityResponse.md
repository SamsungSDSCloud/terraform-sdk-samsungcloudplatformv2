# PortIpAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Port IP Available | 
**FixedIpAddress** | **string** | Fixed IP | 
**Message** | **string** | Port Message | 
**SubnetId** | **string** | Subnet Id | 

## Methods

### NewPortIpAvailabilityResponse

`func NewPortIpAvailabilityResponse(available bool, fixedIpAddress string, message string, subnetId string, ) *PortIpAvailabilityResponse`

NewPortIpAvailabilityResponse instantiates a new PortIpAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortIpAvailabilityResponseWithDefaults

`func NewPortIpAvailabilityResponseWithDefaults() *PortIpAvailabilityResponse`

NewPortIpAvailabilityResponseWithDefaults instantiates a new PortIpAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *PortIpAvailabilityResponse) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *PortIpAvailabilityResponse) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *PortIpAvailabilityResponse) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetFixedIpAddress

`func (o *PortIpAvailabilityResponse) GetFixedIpAddress() string`

GetFixedIpAddress returns the FixedIpAddress field if non-nil, zero value otherwise.

### GetFixedIpAddressOk

`func (o *PortIpAvailabilityResponse) GetFixedIpAddressOk() (*string, bool)`

GetFixedIpAddressOk returns a tuple with the FixedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAddress

`func (o *PortIpAvailabilityResponse) SetFixedIpAddress(v string)`

SetFixedIpAddress sets FixedIpAddress field to given value.


### GetMessage

`func (o *PortIpAvailabilityResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PortIpAvailabilityResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PortIpAvailabilityResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSubnetId

`func (o *PortIpAvailabilityResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *PortIpAvailabilityResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *PortIpAvailabilityResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


