# VpnPhase1CreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpdRetryInterval** | **int32** | VPN Tunnel Dead Peer Detection(DPD) Retry Interval (sec) | 
**IkeVersion** | **int32** | VPN Tunnel IKE Version | 
**PeerGatewayIp** | **string** | VPN Tunnel Peer Gateway IP Address | 
**Phase1DiffieHellmanGroups** | **[]int32** | VPN Tunnel ISAKMP Diffie-Hellman Group List | 
**Phase1Encryptions** | **[]string** | VPN Tunnel ISAKMP Proposal List | 
**Phase1LifeTime** | **int32** | VPN Tunnel ISAKMP Lifetime (sec) | 
**PreSharedKey** | **string** | VPN Tunnel ISAKMP Authentication : Pre-shared Key | 

## Methods

### NewVpnPhase1CreateRequest

`func NewVpnPhase1CreateRequest(dpdRetryInterval int32, ikeVersion int32, peerGatewayIp string, phase1DiffieHellmanGroups []int32, phase1Encryptions []string, phase1LifeTime int32, preSharedKey string, ) *VpnPhase1CreateRequest`

NewVpnPhase1CreateRequest instantiates a new VpnPhase1CreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPhase1CreateRequestWithDefaults

`func NewVpnPhase1CreateRequestWithDefaults() *VpnPhase1CreateRequest`

NewVpnPhase1CreateRequestWithDefaults instantiates a new VpnPhase1CreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpdRetryInterval

`func (o *VpnPhase1CreateRequest) GetDpdRetryInterval() int32`

GetDpdRetryInterval returns the DpdRetryInterval field if non-nil, zero value otherwise.

### GetDpdRetryIntervalOk

`func (o *VpnPhase1CreateRequest) GetDpdRetryIntervalOk() (*int32, bool)`

GetDpdRetryIntervalOk returns a tuple with the DpdRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdRetryInterval

`func (o *VpnPhase1CreateRequest) SetDpdRetryInterval(v int32)`

SetDpdRetryInterval sets DpdRetryInterval field to given value.


### GetIkeVersion

`func (o *VpnPhase1CreateRequest) GetIkeVersion() int32`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *VpnPhase1CreateRequest) GetIkeVersionOk() (*int32, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *VpnPhase1CreateRequest) SetIkeVersion(v int32)`

SetIkeVersion sets IkeVersion field to given value.


### GetPeerGatewayIp

`func (o *VpnPhase1CreateRequest) GetPeerGatewayIp() string`

GetPeerGatewayIp returns the PeerGatewayIp field if non-nil, zero value otherwise.

### GetPeerGatewayIpOk

`func (o *VpnPhase1CreateRequest) GetPeerGatewayIpOk() (*string, bool)`

GetPeerGatewayIpOk returns a tuple with the PeerGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGatewayIp

`func (o *VpnPhase1CreateRequest) SetPeerGatewayIp(v string)`

SetPeerGatewayIp sets PeerGatewayIp field to given value.


### GetPhase1DiffieHellmanGroups

`func (o *VpnPhase1CreateRequest) GetPhase1DiffieHellmanGroups() []int32`

GetPhase1DiffieHellmanGroups returns the Phase1DiffieHellmanGroups field if non-nil, zero value otherwise.

### GetPhase1DiffieHellmanGroupsOk

`func (o *VpnPhase1CreateRequest) GetPhase1DiffieHellmanGroupsOk() (*[]int32, bool)`

GetPhase1DiffieHellmanGroupsOk returns a tuple with the Phase1DiffieHellmanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1DiffieHellmanGroups

`func (o *VpnPhase1CreateRequest) SetPhase1DiffieHellmanGroups(v []int32)`

SetPhase1DiffieHellmanGroups sets Phase1DiffieHellmanGroups field to given value.


### GetPhase1Encryptions

`func (o *VpnPhase1CreateRequest) GetPhase1Encryptions() []string`

GetPhase1Encryptions returns the Phase1Encryptions field if non-nil, zero value otherwise.

### GetPhase1EncryptionsOk

`func (o *VpnPhase1CreateRequest) GetPhase1EncryptionsOk() (*[]string, bool)`

GetPhase1EncryptionsOk returns a tuple with the Phase1Encryptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Encryptions

`func (o *VpnPhase1CreateRequest) SetPhase1Encryptions(v []string)`

SetPhase1Encryptions sets Phase1Encryptions field to given value.


### GetPhase1LifeTime

`func (o *VpnPhase1CreateRequest) GetPhase1LifeTime() int32`

GetPhase1LifeTime returns the Phase1LifeTime field if non-nil, zero value otherwise.

### GetPhase1LifeTimeOk

`func (o *VpnPhase1CreateRequest) GetPhase1LifeTimeOk() (*int32, bool)`

GetPhase1LifeTimeOk returns a tuple with the Phase1LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1LifeTime

`func (o *VpnPhase1CreateRequest) SetPhase1LifeTime(v int32)`

SetPhase1LifeTime sets Phase1LifeTime field to given value.


### GetPreSharedKey

`func (o *VpnPhase1CreateRequest) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VpnPhase1CreateRequest) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VpnPhase1CreateRequest) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


