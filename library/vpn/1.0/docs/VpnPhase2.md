# VpnPhase2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifeTime** | **int32** | VPN Tunnel IPSec Lifetime (sec) | 
**PerfectForwardSecrecy** | **string** | VPN Tunnel IPSec Perfect Forward Secrecy(PFS) | 
**RemoteSubnet** | **string** | VPN Tunnel IPSec Remote Subnet | 

## Methods

### NewVpnPhase2

`func NewVpnPhase2(lifeTime int32, perfectForwardSecrecy string, remoteSubnet string, ) *VpnPhase2`

NewVpnPhase2 instantiates a new VpnPhase2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPhase2WithDefaults

`func NewVpnPhase2WithDefaults() *VpnPhase2`

NewVpnPhase2WithDefaults instantiates a new VpnPhase2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifeTime

`func (o *VpnPhase2) GetLifeTime() int32`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *VpnPhase2) GetLifeTimeOk() (*int32, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *VpnPhase2) SetLifeTime(v int32)`

SetLifeTime sets LifeTime field to given value.


### GetPerfectForwardSecrecy

`func (o *VpnPhase2) GetPerfectForwardSecrecy() string`

GetPerfectForwardSecrecy returns the PerfectForwardSecrecy field if non-nil, zero value otherwise.

### GetPerfectForwardSecrecyOk

`func (o *VpnPhase2) GetPerfectForwardSecrecyOk() (*string, bool)`

GetPerfectForwardSecrecyOk returns a tuple with the PerfectForwardSecrecy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectForwardSecrecy

`func (o *VpnPhase2) SetPerfectForwardSecrecy(v string)`

SetPerfectForwardSecrecy sets PerfectForwardSecrecy field to given value.


### GetRemoteSubnet

`func (o *VpnPhase2) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VpnPhase2) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VpnPhase2) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


