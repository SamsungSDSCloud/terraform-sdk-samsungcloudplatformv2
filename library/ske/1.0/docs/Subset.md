# Subset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** | Addresses | 
**NotReadyAddresses** | **[]string** | Not Ready Addresses | 
**Ports** | [**[]EndpointPort**](EndpointPort.md) | Ports | 

## Methods

### NewSubset

`func NewSubset(addresses []string, notReadyAddresses []string, ports []EndpointPort, ) *Subset`

NewSubset instantiates a new Subset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubsetWithDefaults

`func NewSubsetWithDefaults() *Subset`

NewSubsetWithDefaults instantiates a new Subset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *Subset) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Subset) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Subset) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetNotReadyAddresses

`func (o *Subset) GetNotReadyAddresses() []string`

GetNotReadyAddresses returns the NotReadyAddresses field if non-nil, zero value otherwise.

### GetNotReadyAddressesOk

`func (o *Subset) GetNotReadyAddressesOk() (*[]string, bool)`

GetNotReadyAddressesOk returns a tuple with the NotReadyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotReadyAddresses

`func (o *Subset) SetNotReadyAddresses(v []string)`

SetNotReadyAddresses sets NotReadyAddresses field to given value.


### GetPorts

`func (o *Subset) GetPorts() []EndpointPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Subset) GetPortsOk() (*[]EndpointPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Subset) SetPorts(v []EndpointPort)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


