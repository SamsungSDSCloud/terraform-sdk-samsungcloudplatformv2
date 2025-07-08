# ServerInterfaceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixedIpAddress** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServerInterfaceUpdateRequest

`func NewServerInterfaceUpdateRequest() *ServerInterfaceUpdateRequest`

NewServerInterfaceUpdateRequest instantiates a new ServerInterfaceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInterfaceUpdateRequestWithDefaults

`func NewServerInterfaceUpdateRequestWithDefaults() *ServerInterfaceUpdateRequest`

NewServerInterfaceUpdateRequestWithDefaults instantiates a new ServerInterfaceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixedIpAddress

`func (o *ServerInterfaceUpdateRequest) GetFixedIpAddress() string`

GetFixedIpAddress returns the FixedIpAddress field if non-nil, zero value otherwise.

### GetFixedIpAddressOk

`func (o *ServerInterfaceUpdateRequest) GetFixedIpAddressOk() (*string, bool)`

GetFixedIpAddressOk returns a tuple with the FixedIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAddress

`func (o *ServerInterfaceUpdateRequest) SetFixedIpAddress(v string)`

SetFixedIpAddress sets FixedIpAddress field to given value.

### HasFixedIpAddress

`func (o *ServerInterfaceUpdateRequest) HasFixedIpAddress() bool`

HasFixedIpAddress returns a boolean if a field has been set.

### SetFixedIpAddressNil

`func (o *ServerInterfaceUpdateRequest) SetFixedIpAddressNil(b bool)`

 SetFixedIpAddressNil sets the value for FixedIpAddress to be an explicit nil

### UnsetFixedIpAddress
`func (o *ServerInterfaceUpdateRequest) UnsetFixedIpAddress()`

UnsetFixedIpAddress ensures that no value is present for FixedIpAddress, not even an explicit nil
### GetSubnetId

`func (o *ServerInterfaceUpdateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *ServerInterfaceUpdateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *ServerInterfaceUpdateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *ServerInterfaceUpdateRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *ServerInterfaceUpdateRequest) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *ServerInterfaceUpdateRequest) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


