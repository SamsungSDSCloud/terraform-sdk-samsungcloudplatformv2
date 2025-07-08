# VpnPhase2SetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerfectForwardSecrecy** | Pointer to [**NullableVpnPerfectForwardSecrecyType**](VpnPerfectForwardSecrecyType.md) |  | [optional] 
**Phase2DiffieHellmanGroups** | Pointer to **[]int32** |  | [optional] 
**Phase2Encryptions** | Pointer to **[]string** |  | [optional] 
**Phase2LifeTime** | Pointer to **NullableInt32** |  | [optional] 
**RemoteSubnet** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVpnPhase2SetRequest

`func NewVpnPhase2SetRequest() *VpnPhase2SetRequest`

NewVpnPhase2SetRequest instantiates a new VpnPhase2SetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnPhase2SetRequestWithDefaults

`func NewVpnPhase2SetRequestWithDefaults() *VpnPhase2SetRequest`

NewVpnPhase2SetRequestWithDefaults instantiates a new VpnPhase2SetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerfectForwardSecrecy

`func (o *VpnPhase2SetRequest) GetPerfectForwardSecrecy() VpnPerfectForwardSecrecyType`

GetPerfectForwardSecrecy returns the PerfectForwardSecrecy field if non-nil, zero value otherwise.

### GetPerfectForwardSecrecyOk

`func (o *VpnPhase2SetRequest) GetPerfectForwardSecrecyOk() (*VpnPerfectForwardSecrecyType, bool)`

GetPerfectForwardSecrecyOk returns a tuple with the PerfectForwardSecrecy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfectForwardSecrecy

`func (o *VpnPhase2SetRequest) SetPerfectForwardSecrecy(v VpnPerfectForwardSecrecyType)`

SetPerfectForwardSecrecy sets PerfectForwardSecrecy field to given value.

### HasPerfectForwardSecrecy

`func (o *VpnPhase2SetRequest) HasPerfectForwardSecrecy() bool`

HasPerfectForwardSecrecy returns a boolean if a field has been set.

### SetPerfectForwardSecrecyNil

`func (o *VpnPhase2SetRequest) SetPerfectForwardSecrecyNil(b bool)`

 SetPerfectForwardSecrecyNil sets the value for PerfectForwardSecrecy to be an explicit nil

### UnsetPerfectForwardSecrecy
`func (o *VpnPhase2SetRequest) UnsetPerfectForwardSecrecy()`

UnsetPerfectForwardSecrecy ensures that no value is present for PerfectForwardSecrecy, not even an explicit nil
### GetPhase2DiffieHellmanGroups

`func (o *VpnPhase2SetRequest) GetPhase2DiffieHellmanGroups() []int32`

GetPhase2DiffieHellmanGroups returns the Phase2DiffieHellmanGroups field if non-nil, zero value otherwise.

### GetPhase2DiffieHellmanGroupsOk

`func (o *VpnPhase2SetRequest) GetPhase2DiffieHellmanGroupsOk() (*[]int32, bool)`

GetPhase2DiffieHellmanGroupsOk returns a tuple with the Phase2DiffieHellmanGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2DiffieHellmanGroups

`func (o *VpnPhase2SetRequest) SetPhase2DiffieHellmanGroups(v []int32)`

SetPhase2DiffieHellmanGroups sets Phase2DiffieHellmanGroups field to given value.

### HasPhase2DiffieHellmanGroups

`func (o *VpnPhase2SetRequest) HasPhase2DiffieHellmanGroups() bool`

HasPhase2DiffieHellmanGroups returns a boolean if a field has been set.

### SetPhase2DiffieHellmanGroupsNil

`func (o *VpnPhase2SetRequest) SetPhase2DiffieHellmanGroupsNil(b bool)`

 SetPhase2DiffieHellmanGroupsNil sets the value for Phase2DiffieHellmanGroups to be an explicit nil

### UnsetPhase2DiffieHellmanGroups
`func (o *VpnPhase2SetRequest) UnsetPhase2DiffieHellmanGroups()`

UnsetPhase2DiffieHellmanGroups ensures that no value is present for Phase2DiffieHellmanGroups, not even an explicit nil
### GetPhase2Encryptions

`func (o *VpnPhase2SetRequest) GetPhase2Encryptions() []string`

GetPhase2Encryptions returns the Phase2Encryptions field if non-nil, zero value otherwise.

### GetPhase2EncryptionsOk

`func (o *VpnPhase2SetRequest) GetPhase2EncryptionsOk() (*[]string, bool)`

GetPhase2EncryptionsOk returns a tuple with the Phase2Encryptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Encryptions

`func (o *VpnPhase2SetRequest) SetPhase2Encryptions(v []string)`

SetPhase2Encryptions sets Phase2Encryptions field to given value.

### HasPhase2Encryptions

`func (o *VpnPhase2SetRequest) HasPhase2Encryptions() bool`

HasPhase2Encryptions returns a boolean if a field has been set.

### SetPhase2EncryptionsNil

`func (o *VpnPhase2SetRequest) SetPhase2EncryptionsNil(b bool)`

 SetPhase2EncryptionsNil sets the value for Phase2Encryptions to be an explicit nil

### UnsetPhase2Encryptions
`func (o *VpnPhase2SetRequest) UnsetPhase2Encryptions()`

UnsetPhase2Encryptions ensures that no value is present for Phase2Encryptions, not even an explicit nil
### GetPhase2LifeTime

`func (o *VpnPhase2SetRequest) GetPhase2LifeTime() int32`

GetPhase2LifeTime returns the Phase2LifeTime field if non-nil, zero value otherwise.

### GetPhase2LifeTimeOk

`func (o *VpnPhase2SetRequest) GetPhase2LifeTimeOk() (*int32, bool)`

GetPhase2LifeTimeOk returns a tuple with the Phase2LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2LifeTime

`func (o *VpnPhase2SetRequest) SetPhase2LifeTime(v int32)`

SetPhase2LifeTime sets Phase2LifeTime field to given value.

### HasPhase2LifeTime

`func (o *VpnPhase2SetRequest) HasPhase2LifeTime() bool`

HasPhase2LifeTime returns a boolean if a field has been set.

### SetPhase2LifeTimeNil

`func (o *VpnPhase2SetRequest) SetPhase2LifeTimeNil(b bool)`

 SetPhase2LifeTimeNil sets the value for Phase2LifeTime to be an explicit nil

### UnsetPhase2LifeTime
`func (o *VpnPhase2SetRequest) UnsetPhase2LifeTime()`

UnsetPhase2LifeTime ensures that no value is present for Phase2LifeTime, not even an explicit nil
### GetRemoteSubnet

`func (o *VpnPhase2SetRequest) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VpnPhase2SetRequest) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VpnPhase2SetRequest) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *VpnPhase2SetRequest) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### SetRemoteSubnetNil

`func (o *VpnPhase2SetRequest) SetRemoteSubnetNil(b bool)`

 SetRemoteSubnetNil sets the value for RemoteSubnet to be an explicit nil

### UnsetRemoteSubnet
`func (o *VpnPhase2SetRequest) UnsetRemoteSubnet()`

UnsetRemoteSubnet ensures that no value is present for RemoteSubnet, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


