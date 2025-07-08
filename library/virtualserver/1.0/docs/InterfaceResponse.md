# InterfaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixedIps** | [**[]InterfaceAttachmentFixedIp**](InterfaceAttachmentFixedIp.md) | Fixed IP List | 
**MacAddr** | **string** | MAC Address | 
**PortId** | **string** | Port ID | 
**PortState** | **string** | Port State | 
**PrivateStaticNat** | Pointer to [**PrivateStaticNat**](PrivateStaticNat.md) | Private Static NAT | [optional] 
**StaticNat** | Pointer to [**PublicStaticNat**](PublicStaticNat.md) | Public Static NAT | [optional] 
**SubnetId** | **string** | Subnet ID | 

## Methods

### NewInterfaceResponse

`func NewInterfaceResponse(fixedIps []InterfaceAttachmentFixedIp, macAddr string, portId string, portState string, subnetId string, ) *InterfaceResponse`

NewInterfaceResponse instantiates a new InterfaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceResponseWithDefaults

`func NewInterfaceResponseWithDefaults() *InterfaceResponse`

NewInterfaceResponseWithDefaults instantiates a new InterfaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixedIps

`func (o *InterfaceResponse) GetFixedIps() []InterfaceAttachmentFixedIp`

GetFixedIps returns the FixedIps field if non-nil, zero value otherwise.

### GetFixedIpsOk

`func (o *InterfaceResponse) GetFixedIpsOk() (*[]InterfaceAttachmentFixedIp, bool)`

GetFixedIpsOk returns a tuple with the FixedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIps

`func (o *InterfaceResponse) SetFixedIps(v []InterfaceAttachmentFixedIp)`

SetFixedIps sets FixedIps field to given value.


### GetMacAddr

`func (o *InterfaceResponse) GetMacAddr() string`

GetMacAddr returns the MacAddr field if non-nil, zero value otherwise.

### GetMacAddrOk

`func (o *InterfaceResponse) GetMacAddrOk() (*string, bool)`

GetMacAddrOk returns a tuple with the MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddr

`func (o *InterfaceResponse) SetMacAddr(v string)`

SetMacAddr sets MacAddr field to given value.


### GetPortId

`func (o *InterfaceResponse) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *InterfaceResponse) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *InterfaceResponse) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetPortState

`func (o *InterfaceResponse) GetPortState() string`

GetPortState returns the PortState field if non-nil, zero value otherwise.

### GetPortStateOk

`func (o *InterfaceResponse) GetPortStateOk() (*string, bool)`

GetPortStateOk returns a tuple with the PortState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortState

`func (o *InterfaceResponse) SetPortState(v string)`

SetPortState sets PortState field to given value.


### GetPrivateStaticNat

`func (o *InterfaceResponse) GetPrivateStaticNat() PrivateStaticNat`

GetPrivateStaticNat returns the PrivateStaticNat field if non-nil, zero value otherwise.

### GetPrivateStaticNatOk

`func (o *InterfaceResponse) GetPrivateStaticNatOk() (*PrivateStaticNat, bool)`

GetPrivateStaticNatOk returns a tuple with the PrivateStaticNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateStaticNat

`func (o *InterfaceResponse) SetPrivateStaticNat(v PrivateStaticNat)`

SetPrivateStaticNat sets PrivateStaticNat field to given value.

### HasPrivateStaticNat

`func (o *InterfaceResponse) HasPrivateStaticNat() bool`

HasPrivateStaticNat returns a boolean if a field has been set.

### GetStaticNat

`func (o *InterfaceResponse) GetStaticNat() PublicStaticNat`

GetStaticNat returns the StaticNat field if non-nil, zero value otherwise.

### GetStaticNatOk

`func (o *InterfaceResponse) GetStaticNatOk() (*PublicStaticNat, bool)`

GetStaticNatOk returns a tuple with the StaticNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticNat

`func (o *InterfaceResponse) SetStaticNat(v PublicStaticNat)`

SetStaticNat sets StaticNat field to given value.

### HasStaticNat

`func (o *InterfaceResponse) HasStaticNat() bool`

HasStaticNat returns a boolean if a field has been set.

### GetSubnetId

`func (o *InterfaceResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *InterfaceResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *InterfaceResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


