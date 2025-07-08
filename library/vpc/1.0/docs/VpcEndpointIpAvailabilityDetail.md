# VpcEndpointIpAvailabilityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Availability of Endpoint IP Address | 
**EndpointIpAddress** | **string** | VPC Endpoint IP Address | 
**Message** | **string** | Vpc Endpoint Message | 
**SubnetId** | **string** | Subnet Id | 

## Methods

### NewVpcEndpointIpAvailabilityDetail

`func NewVpcEndpointIpAvailabilityDetail(available bool, endpointIpAddress string, message string, subnetId string, ) *VpcEndpointIpAvailabilityDetail`

NewVpcEndpointIpAvailabilityDetail instantiates a new VpcEndpointIpAvailabilityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcEndpointIpAvailabilityDetailWithDefaults

`func NewVpcEndpointIpAvailabilityDetailWithDefaults() *VpcEndpointIpAvailabilityDetail`

NewVpcEndpointIpAvailabilityDetailWithDefaults instantiates a new VpcEndpointIpAvailabilityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *VpcEndpointIpAvailabilityDetail) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *VpcEndpointIpAvailabilityDetail) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *VpcEndpointIpAvailabilityDetail) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetEndpointIpAddress

`func (o *VpcEndpointIpAvailabilityDetail) GetEndpointIpAddress() string`

GetEndpointIpAddress returns the EndpointIpAddress field if non-nil, zero value otherwise.

### GetEndpointIpAddressOk

`func (o *VpcEndpointIpAvailabilityDetail) GetEndpointIpAddressOk() (*string, bool)`

GetEndpointIpAddressOk returns a tuple with the EndpointIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointIpAddress

`func (o *VpcEndpointIpAvailabilityDetail) SetEndpointIpAddress(v string)`

SetEndpointIpAddress sets EndpointIpAddress field to given value.


### GetMessage

`func (o *VpcEndpointIpAvailabilityDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VpcEndpointIpAvailabilityDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VpcEndpointIpAvailabilityDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSubnetId

`func (o *VpcEndpointIpAvailabilityDetail) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *VpcEndpointIpAvailabilityDetail) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *VpcEndpointIpAvailabilityDetail) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


