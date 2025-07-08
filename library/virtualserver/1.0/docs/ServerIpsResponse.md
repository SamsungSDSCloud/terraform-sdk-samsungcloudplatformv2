# ServerIpsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]ServerShowResponseAddress**](ServerShowResponseAddress.md) | IP addresses | 

## Methods

### NewServerIpsResponse

`func NewServerIpsResponse(addresses []ServerShowResponseAddress, ) *ServerIpsResponse`

NewServerIpsResponse instantiates a new ServerIpsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerIpsResponseWithDefaults

`func NewServerIpsResponseWithDefaults() *ServerIpsResponse`

NewServerIpsResponseWithDefaults instantiates a new ServerIpsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ServerIpsResponse) GetAddresses() []ServerShowResponseAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ServerIpsResponse) GetAddressesOk() (*[]ServerShowResponseAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ServerIpsResponse) SetAddresses(v []ServerShowResponseAddress)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


