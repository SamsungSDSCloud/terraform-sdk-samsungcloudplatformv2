# VpnPhase1Detail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiffieHellmanGroups** | **[]int32** | VPN Tunnel ISAKMP Diffie-Hellman Group List | 
**DpdRetryInterval** | **int32** | VPN Tunnel Dead Peer Detection(DPD) Retry Interval (sec) | 
**Encryptions** | **[]string** | VPN Tunnel ISAKMP Proposal List | 
**IkeVersion** | **int32** | VPN Tunnel IKE Version | 
**LifeTime** | **int32** | VPN Tunnel ISAKMP Lifetime (sec) | 
**PeerGatewayIp** | **string** | VPN Tunnel Peer Gateway IP Address | 

## Methods

### NewVpnPhase1Detail

`func NewVpnPhase1Detail(diffieHellmanGroups []int32, dpdRetryInterval int32, encryptions []string, ikeVersion int32, lifeTime int32, peerGatewayIp string, ) *VpnPhase1Detail`

NewVpnPhase1Detail instantiates a new VpnPhase1Detail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPhase1DetailWithDefaults

`func NewVpnPhase1DetailWithDefaults() *VpnPhase1Detail`

NewVpnPhase1DetailWithDefaults instantiates a new VpnPhase1Detail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffieHellmanGroups

`func (o *VpnPhase1Detail) GetDiffieHellmanGroups() []int32`

GetDiffieHellmanGroups returns the DiffieHellmanGroups field if non-nil, zero value otherwise.

### GetDiffieHellmanGroupsOk

`func (o *VpnPhase1Detail) GetDiffieHellmanGroupsOk() (*[]int32, bool)`

GetDiffieHellmanGroupsOk returns a tuple with the DiffieHellmanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffieHellmanGroups

`func (o *VpnPhase1Detail) SetDiffieHellmanGroups(v []int32)`

SetDiffieHellmanGroups sets DiffieHellmanGroups field to given value.


### GetDpdRetryInterval

`func (o *VpnPhase1Detail) GetDpdRetryInterval() int32`

GetDpdRetryInterval returns the DpdRetryInterval field if non-nil, zero value otherwise.

### GetDpdRetryIntervalOk

`func (o *VpnPhase1Detail) GetDpdRetryIntervalOk() (*int32, bool)`

GetDpdRetryIntervalOk returns a tuple with the DpdRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdRetryInterval

`func (o *VpnPhase1Detail) SetDpdRetryInterval(v int32)`

SetDpdRetryInterval sets DpdRetryInterval field to given value.


### GetEncryptions

`func (o *VpnPhase1Detail) GetEncryptions() []string`

GetEncryptions returns the Encryptions field if non-nil, zero value otherwise.

### GetEncryptionsOk

`func (o *VpnPhase1Detail) GetEncryptionsOk() (*[]string, bool)`

GetEncryptionsOk returns a tuple with the Encryptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptions

`func (o *VpnPhase1Detail) SetEncryptions(v []string)`

SetEncryptions sets Encryptions field to given value.


### GetIkeVersion

`func (o *VpnPhase1Detail) GetIkeVersion() int32`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *VpnPhase1Detail) GetIkeVersionOk() (*int32, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *VpnPhase1Detail) SetIkeVersion(v int32)`

SetIkeVersion sets IkeVersion field to given value.


### GetLifeTime

`func (o *VpnPhase1Detail) GetLifeTime() int32`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *VpnPhase1Detail) GetLifeTimeOk() (*int32, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *VpnPhase1Detail) SetLifeTime(v int32)`

SetLifeTime sets LifeTime field to given value.


### GetPeerGatewayIp

`func (o *VpnPhase1Detail) GetPeerGatewayIp() string`

GetPeerGatewayIp returns the PeerGatewayIp field if non-nil, zero value otherwise.

### GetPeerGatewayIpOk

`func (o *VpnPhase1Detail) GetPeerGatewayIpOk() (*string, bool)`

GetPeerGatewayIpOk returns a tuple with the PeerGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGatewayIp

`func (o *VpnPhase1Detail) SetPeerGatewayIp(v string)`

SetPeerGatewayIp sets PeerGatewayIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


