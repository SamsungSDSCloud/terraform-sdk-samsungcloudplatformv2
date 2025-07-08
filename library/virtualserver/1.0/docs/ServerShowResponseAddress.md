# ServerShowResponseAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddresses** | [**[]ServerShowResponseAddressIpAddress**](ServerShowResponseAddressIpAddress.md) | IP addresses | 
**SubnetName** | **string** | Subnet name | 

## Methods

### NewServerShowResponseAddress

`func NewServerShowResponseAddress(ipAddresses []ServerShowResponseAddressIpAddress, subnetName string, ) *ServerShowResponseAddress`

NewServerShowResponseAddress instantiates a new ServerShowResponseAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerShowResponseAddressWithDefaults

`func NewServerShowResponseAddressWithDefaults() *ServerShowResponseAddress`

NewServerShowResponseAddressWithDefaults instantiates a new ServerShowResponseAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddresses

`func (o *ServerShowResponseAddress) GetIpAddresses() []ServerShowResponseAddressIpAddress`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *ServerShowResponseAddress) GetIpAddressesOk() (*[]ServerShowResponseAddressIpAddress, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *ServerShowResponseAddress) SetIpAddresses(v []ServerShowResponseAddressIpAddress)`

SetIpAddresses sets IpAddresses field to given value.


### GetSubnetName

`func (o *ServerShowResponseAddress) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *ServerShowResponseAddress) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *ServerShowResponseAddress) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


