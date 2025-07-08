# ServerInterfaceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixedIps** | Pointer to [**[]InterfaceAttachmentFixedIp**](InterfaceAttachmentFixedIp.md) |  | [optional] 
**PortId** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServerInterfaceCreateRequest

`func NewServerInterfaceCreateRequest() *ServerInterfaceCreateRequest`

NewServerInterfaceCreateRequest instantiates a new ServerInterfaceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInterfaceCreateRequestWithDefaults

`func NewServerInterfaceCreateRequestWithDefaults() *ServerInterfaceCreateRequest`

NewServerInterfaceCreateRequestWithDefaults instantiates a new ServerInterfaceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixedIps

`func (o *ServerInterfaceCreateRequest) GetFixedIps() []InterfaceAttachmentFixedIp`

GetFixedIps returns the FixedIps field if non-nil, zero value otherwise.

### GetFixedIpsOk

`func (o *ServerInterfaceCreateRequest) GetFixedIpsOk() (*[]InterfaceAttachmentFixedIp, bool)`

GetFixedIpsOk returns a tuple with the FixedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIps

`func (o *ServerInterfaceCreateRequest) SetFixedIps(v []InterfaceAttachmentFixedIp)`

SetFixedIps sets FixedIps field to given value.

### HasFixedIps

`func (o *ServerInterfaceCreateRequest) HasFixedIps() bool`

HasFixedIps returns a boolean if a field has been set.

### SetFixedIpsNil

`func (o *ServerInterfaceCreateRequest) SetFixedIpsNil(b bool)`

 SetFixedIpsNil sets the value for FixedIps to be an explicit nil

### UnsetFixedIps
`func (o *ServerInterfaceCreateRequest) UnsetFixedIps()`

UnsetFixedIps ensures that no value is present for FixedIps, not even an explicit nil
### GetPortId

`func (o *ServerInterfaceCreateRequest) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *ServerInterfaceCreateRequest) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *ServerInterfaceCreateRequest) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *ServerInterfaceCreateRequest) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### SetPortIdNil

`func (o *ServerInterfaceCreateRequest) SetPortIdNil(b bool)`

 SetPortIdNil sets the value for PortId to be an explicit nil

### UnsetPortId
`func (o *ServerInterfaceCreateRequest) UnsetPortId()`

UnsetPortId ensures that no value is present for PortId, not even an explicit nil
### GetSubnetId

`func (o *ServerInterfaceCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *ServerInterfaceCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *ServerInterfaceCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *ServerInterfaceCreateRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *ServerInterfaceCreateRequest) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *ServerInterfaceCreateRequest) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


