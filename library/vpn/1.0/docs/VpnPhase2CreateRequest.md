# VpnPhase2CreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerfectForwardSecrecy** | [**VpnPerfectForwardSecrecyType**](VpnPerfectForwardSecrecyType.md) | VPN Tunnel IPSec Perfect Forward Secrecy(PFS) | 
**Phase2DiffieHellmanGroups** | **[]int32** | VPN Tunnel ISAKMP Diffie-Hellman Group List | 
**Phase2Encryptions** | **[]string** | VPN Tunnel ISAKMP Proposal List | 
**Phase2LifeTime** | **int32** | VPN Tunnel IPSec Lifetime (sec) | 
**RemoteSubnet** | **string** | VPN Tunnel IPSec Remote Subnet | 

## Methods

### NewVpnPhase2CreateRequest

`func NewVpnPhase2CreateRequest(perfectForwardSecrecy VpnPerfectForwardSecrecyType, phase2DiffieHellmanGroups []int32, phase2Encryptions []string, phase2LifeTime int32, remoteSubnet string, ) *VpnPhase2CreateRequest`

NewVpnPhase2CreateRequest instantiates a new VpnPhase2CreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPhase2CreateRequestWithDefaults

`func NewVpnPhase2CreateRequestWithDefaults() *VpnPhase2CreateRequest`

NewVpnPhase2CreateRequestWithDefaults instantiates a new VpnPhase2CreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerfectForwardSecrecy

`func (o *VpnPhase2CreateRequest) GetPerfectForwardSecrecy() VpnPerfectForwardSecrecyType`

GetPerfectForwardSecrecy returns the PerfectForwardSecrecy field if non-nil, zero value otherwise.

### GetPerfectForwardSecrecyOk

`func (o *VpnPhase2CreateRequest) GetPerfectForwardSecrecyOk() (*VpnPerfectForwardSecrecyType, bool)`

GetPerfectForwardSecrecyOk returns a tuple with the PerfectForwardSecrecy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectForwardSecrecy

`func (o *VpnPhase2CreateRequest) SetPerfectForwardSecrecy(v VpnPerfectForwardSecrecyType)`

SetPerfectForwardSecrecy sets PerfectForwardSecrecy field to given value.


### GetPhase2DiffieHellmanGroups

`func (o *VpnPhase2CreateRequest) GetPhase2DiffieHellmanGroups() []int32`

GetPhase2DiffieHellmanGroups returns the Phase2DiffieHellmanGroups field if non-nil, zero value otherwise.

### GetPhase2DiffieHellmanGroupsOk

`func (o *VpnPhase2CreateRequest) GetPhase2DiffieHellmanGroupsOk() (*[]int32, bool)`

GetPhase2DiffieHellmanGroupsOk returns a tuple with the Phase2DiffieHellmanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2DiffieHellmanGroups

`func (o *VpnPhase2CreateRequest) SetPhase2DiffieHellmanGroups(v []int32)`

SetPhase2DiffieHellmanGroups sets Phase2DiffieHellmanGroups field to given value.


### GetPhase2Encryptions

`func (o *VpnPhase2CreateRequest) GetPhase2Encryptions() []string`

GetPhase2Encryptions returns the Phase2Encryptions field if non-nil, zero value otherwise.

### GetPhase2EncryptionsOk

`func (o *VpnPhase2CreateRequest) GetPhase2EncryptionsOk() (*[]string, bool)`

GetPhase2EncryptionsOk returns a tuple with the Phase2Encryptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Encryptions

`func (o *VpnPhase2CreateRequest) SetPhase2Encryptions(v []string)`

SetPhase2Encryptions sets Phase2Encryptions field to given value.


### GetPhase2LifeTime

`func (o *VpnPhase2CreateRequest) GetPhase2LifeTime() int32`

GetPhase2LifeTime returns the Phase2LifeTime field if non-nil, zero value otherwise.

### GetPhase2LifeTimeOk

`func (o *VpnPhase2CreateRequest) GetPhase2LifeTimeOk() (*int32, bool)`

GetPhase2LifeTimeOk returns a tuple with the Phase2LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2LifeTime

`func (o *VpnPhase2CreateRequest) SetPhase2LifeTime(v int32)`

SetPhase2LifeTime sets Phase2LifeTime field to given value.


### GetRemoteSubnet

`func (o *VpnPhase2CreateRequest) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VpnPhase2CreateRequest) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VpnPhase2CreateRequest) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


