# PrivateNatCidrAvailabilityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Availability of Private NAT IP range | 
**Cidr** | **string** | Private NAT IP range | 
**DirectConnectId** | **string** | Direct Connect ID | 
**Message** | **string** | Message on the Availability of Private NAT IP Address Range | 

## Methods

### NewPrivateNatCidrAvailabilityDetail

`func NewPrivateNatCidrAvailabilityDetail(available bool, cidr string, directConnectId string, message string, ) *PrivateNatCidrAvailabilityDetail`

NewPrivateNatCidrAvailabilityDetail instantiates a new PrivateNatCidrAvailabilityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNatCidrAvailabilityDetailWithDefaults

`func NewPrivateNatCidrAvailabilityDetailWithDefaults() *PrivateNatCidrAvailabilityDetail`

NewPrivateNatCidrAvailabilityDetailWithDefaults instantiates a new PrivateNatCidrAvailabilityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *PrivateNatCidrAvailabilityDetail) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *PrivateNatCidrAvailabilityDetail) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *PrivateNatCidrAvailabilityDetail) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCidr

`func (o *PrivateNatCidrAvailabilityDetail) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *PrivateNatCidrAvailabilityDetail) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *PrivateNatCidrAvailabilityDetail) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDirectConnectId

`func (o *PrivateNatCidrAvailabilityDetail) GetDirectConnectId() string`

GetDirectConnectId returns the DirectConnectId field if non-nil, zero value otherwise.

### GetDirectConnectIdOk

`func (o *PrivateNatCidrAvailabilityDetail) GetDirectConnectIdOk() (*string, bool)`

GetDirectConnectIdOk returns a tuple with the DirectConnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnectId

`func (o *PrivateNatCidrAvailabilityDetail) SetDirectConnectId(v string)`

SetDirectConnectId sets DirectConnectId field to given value.


### GetMessage

`func (o *PrivateNatCidrAvailabilityDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PrivateNatCidrAvailabilityDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PrivateNatCidrAvailabilityDetail) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


