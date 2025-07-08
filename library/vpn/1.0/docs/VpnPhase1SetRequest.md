# VpnPhase1SetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DpdRetryInterval** | Pointer to **NullableInt32** |  | [optional] 
**IkeVersion** | Pointer to **NullableInt32** |  | [optional] 
**PeerGatewayIp** | Pointer to **NullableString** |  | [optional] 
**Phase1DiffieHellmanGroups** | Pointer to **[]int32** |  | [optional] 
**Phase1Encryptions** | Pointer to **[]string** |  | [optional] 
**Phase1LifeTime** | Pointer to **NullableInt32** |  | [optional] 
**PreSharedKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVpnPhase1SetRequest

`func NewVpnPhase1SetRequest() *VpnPhase1SetRequest`

NewVpnPhase1SetRequest instantiates a new VpnPhase1SetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPhase1SetRequestWithDefaults

`func NewVpnPhase1SetRequestWithDefaults() *VpnPhase1SetRequest`

NewVpnPhase1SetRequestWithDefaults instantiates a new VpnPhase1SetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpdRetryInterval

`func (o *VpnPhase1SetRequest) GetDpdRetryInterval() int32`

GetDpdRetryInterval returns the DpdRetryInterval field if non-nil, zero value otherwise.

### GetDpdRetryIntervalOk

`func (o *VpnPhase1SetRequest) GetDpdRetryIntervalOk() (*int32, bool)`

GetDpdRetryIntervalOk returns a tuple with the DpdRetryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdRetryInterval

`func (o *VpnPhase1SetRequest) SetDpdRetryInterval(v int32)`

SetDpdRetryInterval sets DpdRetryInterval field to given value.

### HasDpdRetryInterval

`func (o *VpnPhase1SetRequest) HasDpdRetryInterval() bool`

HasDpdRetryInterval returns a boolean if a field has been set.

### SetDpdRetryIntervalNil

`func (o *VpnPhase1SetRequest) SetDpdRetryIntervalNil(b bool)`

 SetDpdRetryIntervalNil sets the value for DpdRetryInterval to be an explicit nil

### UnsetDpdRetryInterval
`func (o *VpnPhase1SetRequest) UnsetDpdRetryInterval()`

UnsetDpdRetryInterval ensures that no value is present for DpdRetryInterval, not even an explicit nil
### GetIkeVersion

`func (o *VpnPhase1SetRequest) GetIkeVersion() int32`

GetIkeVersion returns the IkeVersion field if non-nil, zero value otherwise.

### GetIkeVersionOk

`func (o *VpnPhase1SetRequest) GetIkeVersionOk() (*int32, bool)`

GetIkeVersionOk returns a tuple with the IkeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeVersion

`func (o *VpnPhase1SetRequest) SetIkeVersion(v int32)`

SetIkeVersion sets IkeVersion field to given value.

### HasIkeVersion

`func (o *VpnPhase1SetRequest) HasIkeVersion() bool`

HasIkeVersion returns a boolean if a field has been set.

### SetIkeVersionNil

`func (o *VpnPhase1SetRequest) SetIkeVersionNil(b bool)`

 SetIkeVersionNil sets the value for IkeVersion to be an explicit nil

### UnsetIkeVersion
`func (o *VpnPhase1SetRequest) UnsetIkeVersion()`

UnsetIkeVersion ensures that no value is present for IkeVersion, not even an explicit nil
### GetPeerGatewayIp

`func (o *VpnPhase1SetRequest) GetPeerGatewayIp() string`

GetPeerGatewayIp returns the PeerGatewayIp field if non-nil, zero value otherwise.

### GetPeerGatewayIpOk

`func (o *VpnPhase1SetRequest) GetPeerGatewayIpOk() (*string, bool)`

GetPeerGatewayIpOk returns a tuple with the PeerGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerGatewayIp

`func (o *VpnPhase1SetRequest) SetPeerGatewayIp(v string)`

SetPeerGatewayIp sets PeerGatewayIp field to given value.

### HasPeerGatewayIp

`func (o *VpnPhase1SetRequest) HasPeerGatewayIp() bool`

HasPeerGatewayIp returns a boolean if a field has been set.

### SetPeerGatewayIpNil

`func (o *VpnPhase1SetRequest) SetPeerGatewayIpNil(b bool)`

 SetPeerGatewayIpNil sets the value for PeerGatewayIp to be an explicit nil

### UnsetPeerGatewayIp
`func (o *VpnPhase1SetRequest) UnsetPeerGatewayIp()`

UnsetPeerGatewayIp ensures that no value is present for PeerGatewayIp, not even an explicit nil
### GetPhase1DiffieHellmanGroups

`func (o *VpnPhase1SetRequest) GetPhase1DiffieHellmanGroups() []int32`

GetPhase1DiffieHellmanGroups returns the Phase1DiffieHellmanGroups field if non-nil, zero value otherwise.

### GetPhase1DiffieHellmanGroupsOk

`func (o *VpnPhase1SetRequest) GetPhase1DiffieHellmanGroupsOk() (*[]int32, bool)`

GetPhase1DiffieHellmanGroupsOk returns a tuple with the Phase1DiffieHellmanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1DiffieHellmanGroups

`func (o *VpnPhase1SetRequest) SetPhase1DiffieHellmanGroups(v []int32)`

SetPhase1DiffieHellmanGroups sets Phase1DiffieHellmanGroups field to given value.

### HasPhase1DiffieHellmanGroups

`func (o *VpnPhase1SetRequest) HasPhase1DiffieHellmanGroups() bool`

HasPhase1DiffieHellmanGroups returns a boolean if a field has been set.

### SetPhase1DiffieHellmanGroupsNil

`func (o *VpnPhase1SetRequest) SetPhase1DiffieHellmanGroupsNil(b bool)`

 SetPhase1DiffieHellmanGroupsNil sets the value for Phase1DiffieHellmanGroups to be an explicit nil

### UnsetPhase1DiffieHellmanGroups
`func (o *VpnPhase1SetRequest) UnsetPhase1DiffieHellmanGroups()`

UnsetPhase1DiffieHellmanGroups ensures that no value is present for Phase1DiffieHellmanGroups, not even an explicit nil
### GetPhase1Encryptions

`func (o *VpnPhase1SetRequest) GetPhase1Encryptions() []string`

GetPhase1Encryptions returns the Phase1Encryptions field if non-nil, zero value otherwise.

### GetPhase1EncryptionsOk

`func (o *VpnPhase1SetRequest) GetPhase1EncryptionsOk() (*[]string, bool)`

GetPhase1EncryptionsOk returns a tuple with the Phase1Encryptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Encryptions

`func (o *VpnPhase1SetRequest) SetPhase1Encryptions(v []string)`

SetPhase1Encryptions sets Phase1Encryptions field to given value.

### HasPhase1Encryptions

`func (o *VpnPhase1SetRequest) HasPhase1Encryptions() bool`

HasPhase1Encryptions returns a boolean if a field has been set.

### SetPhase1EncryptionsNil

`func (o *VpnPhase1SetRequest) SetPhase1EncryptionsNil(b bool)`

 SetPhase1EncryptionsNil sets the value for Phase1Encryptions to be an explicit nil

### UnsetPhase1Encryptions
`func (o *VpnPhase1SetRequest) UnsetPhase1Encryptions()`

UnsetPhase1Encryptions ensures that no value is present for Phase1Encryptions, not even an explicit nil
### GetPhase1LifeTime

`func (o *VpnPhase1SetRequest) GetPhase1LifeTime() int32`

GetPhase1LifeTime returns the Phase1LifeTime field if non-nil, zero value otherwise.

### GetPhase1LifeTimeOk

`func (o *VpnPhase1SetRequest) GetPhase1LifeTimeOk() (*int32, bool)`

GetPhase1LifeTimeOk returns a tuple with the Phase1LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1LifeTime

`func (o *VpnPhase1SetRequest) SetPhase1LifeTime(v int32)`

SetPhase1LifeTime sets Phase1LifeTime field to given value.

### HasPhase1LifeTime

`func (o *VpnPhase1SetRequest) HasPhase1LifeTime() bool`

HasPhase1LifeTime returns a boolean if a field has been set.

### SetPhase1LifeTimeNil

`func (o *VpnPhase1SetRequest) SetPhase1LifeTimeNil(b bool)`

 SetPhase1LifeTimeNil sets the value for Phase1LifeTime to be an explicit nil

### UnsetPhase1LifeTime
`func (o *VpnPhase1SetRequest) UnsetPhase1LifeTime()`

UnsetPhase1LifeTime ensures that no value is present for Phase1LifeTime, not even an explicit nil
### GetPreSharedKey

`func (o *VpnPhase1SetRequest) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *VpnPhase1SetRequest) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *VpnPhase1SetRequest) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *VpnPhase1SetRequest) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### SetPreSharedKeyNil

`func (o *VpnPhase1SetRequest) SetPreSharedKeyNil(b bool)`

 SetPreSharedKeyNil sets the value for PreSharedKey to be an explicit nil

### UnsetPreSharedKey
`func (o *VpnPhase1SetRequest) UnsetPreSharedKey()`

UnsetPreSharedKey ensures that no value is present for PreSharedKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


