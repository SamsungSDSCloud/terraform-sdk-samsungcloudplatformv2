# VpnPhase1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpdRetryInterval** | **int32** | VPN Tunnel Dead Peer Detection(DPD) Retry Interval (sec) | 
**IkeVersion** | **int32** | VPN Tunnel IKE Version | 
**LifeTime** | **int32** | VPN Tunnel ISAKMP Lifetime (sec) | 
**PeerGatewayIp** | **string** | VPN Tunnel Peer Gateway IP Address | 

## Methods

### NewVpnPhase1

`func NewVpnPhase1(dpdRetryInterval int32, ikeVersion int32, lifeTime int32, peerGatewayIp string, ) *VpnPhase1`

NewVpnPhase1 instantiates a new VpnPhase1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPhase1WithDefaults

`func NewVpnPhase1WithDefaults() *VpnPhase1`

NewVpnPhase1WithDefaults instantiates a new VpnPhase1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpdRetryInterval

`func (o *VpnPhase1) GetDpdRetryInterval() int32`

GetDpdRetryInterval returns the DpdRetryInterval field if non-nil, zero value otherwise.

### GetDpdRetryIntervalOk

`func (o *VpnPhase1) GetDpdRetryIntervalOk() (*int32, bool)`

GetDpdRetryIntervalOk returns a tuple with the DpdRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdRetryInterval

`func (o *VpnPhase1) SetDpdRetryInterval(v int32)`

SetDpdRetryInterval sets DpdRetryInterval field to given value.


### GetIkeVersion

`func (o *VpnPhase1) GetIkeVersion() int32`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *VpnPhase1) GetIkeVersionOk() (*int32, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *VpnPhase1) SetIkeVersion(v int32)`

SetIkeVersion sets IkeVersion field to given value.


### GetLifeTime

`func (o *VpnPhase1) GetLifeTime() int32`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *VpnPhase1) GetLifeTimeOk() (*int32, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *VpnPhase1) SetLifeTime(v int32)`

SetLifeTime sets LifeTime field to given value.


### GetPeerGatewayIp

`func (o *VpnPhase1) GetPeerGatewayIp() string`

GetPeerGatewayIp returns the PeerGatewayIp field if non-nil, zero value otherwise.

### GetPeerGatewayIpOk

`func (o *VpnPhase1) GetPeerGatewayIpOk() (*string, bool)`

GetPeerGatewayIpOk returns a tuple with the PeerGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGatewayIp

`func (o *VpnPhase1) SetPeerGatewayIp(v string)`

SetPeerGatewayIp sets PeerGatewayIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


