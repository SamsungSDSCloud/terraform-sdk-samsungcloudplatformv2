# VpnPhase2Detail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiffieHellmanGroups** | Pointer to **[]int32** |  | [optional] 
**Encryptions** | **[]string** | VPN Tunnel IPSec Proposal List | 
**LifeTime** | **int32** | VPN Tunnel IPSec Lifetime (sec) | 
**PerfectForwardSecrecy** | **string** | VPN Tunnel IPSec Perfect Forward Secrecy(PFS) | 
**RemoteSubnet** | **string** | VPN Tunnel IPSec Remote Subnet | 

## Methods

### NewVpnPhase2Detail

`func NewVpnPhase2Detail(encryptions []string, lifeTime int32, perfectForwardSecrecy string, remoteSubnet string, ) *VpnPhase2Detail`

NewVpnPhase2Detail instantiates a new VpnPhase2Detail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPhase2DetailWithDefaults

`func NewVpnPhase2DetailWithDefaults() *VpnPhase2Detail`

NewVpnPhase2DetailWithDefaults instantiates a new VpnPhase2Detail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffieHellmanGroups

`func (o *VpnPhase2Detail) GetDiffieHellmanGroups() []int32`

GetDiffieHellmanGroups returns the DiffieHellmanGroups field if non-nil, zero value otherwise.

### GetDiffieHellmanGroupsOk

`func (o *VpnPhase2Detail) GetDiffieHellmanGroupsOk() (*[]int32, bool)`

GetDiffieHellmanGroupsOk returns a tuple with the DiffieHellmanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffieHellmanGroups

`func (o *VpnPhase2Detail) SetDiffieHellmanGroups(v []int32)`

SetDiffieHellmanGroups sets DiffieHellmanGroups field to given value.

### HasDiffieHellmanGroups

`func (o *VpnPhase2Detail) HasDiffieHellmanGroups() bool`

HasDiffieHellmanGroups returns a boolean if a field has been set.

### SetDiffieHellmanGroupsNil

`func (o *VpnPhase2Detail) SetDiffieHellmanGroupsNil(b bool)`

 SetDiffieHellmanGroupsNil sets the value for DiffieHellmanGroups to be an explicit nil

### UnsetDiffieHellmanGroups
`func (o *VpnPhase2Detail) UnsetDiffieHellmanGroups()`

UnsetDiffieHellmanGroups ensures that no value is present for DiffieHellmanGroups, not even an explicit nil
### GetEncryptions

`func (o *VpnPhase2Detail) GetEncryptions() []string`

GetEncryptions returns the Encryptions field if non-nil, zero value otherwise.

### GetEncryptionsOk

`func (o *VpnPhase2Detail) GetEncryptionsOk() (*[]string, bool)`

GetEncryptionsOk returns a tuple with the Encryptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptions

`func (o *VpnPhase2Detail) SetEncryptions(v []string)`

SetEncryptions sets Encryptions field to given value.


### GetLifeTime

`func (o *VpnPhase2Detail) GetLifeTime() int32`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *VpnPhase2Detail) GetLifeTimeOk() (*int32, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *VpnPhase2Detail) SetLifeTime(v int32)`

SetLifeTime sets LifeTime field to given value.


### GetPerfectForwardSecrecy

`func (o *VpnPhase2Detail) GetPerfectForwardSecrecy() string`

GetPerfectForwardSecrecy returns the PerfectForwardSecrecy field if non-nil, zero value otherwise.

### GetPerfectForwardSecrecyOk

`func (o *VpnPhase2Detail) GetPerfectForwardSecrecyOk() (*string, bool)`

GetPerfectForwardSecrecyOk returns a tuple with the PerfectForwardSecrecy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectForwardSecrecy

`func (o *VpnPhase2Detail) SetPerfectForwardSecrecy(v string)`

SetPerfectForwardSecrecy sets PerfectForwardSecrecy field to given value.


### GetRemoteSubnet

`func (o *VpnPhase2Detail) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VpnPhase2Detail) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VpnPhase2Detail) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


