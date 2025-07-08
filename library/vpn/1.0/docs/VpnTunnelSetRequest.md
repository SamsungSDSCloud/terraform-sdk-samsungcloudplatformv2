# VpnTunnelSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Phase1** | Pointer to [**NullableVpnPhase1SetRequest**](VpnPhase1SetRequest.md) |  | [optional] 
**Phase2** | Pointer to [**NullableVpnPhase2SetRequest**](VpnPhase2SetRequest.md) |  | [optional] 

## Methods

### NewVpnTunnelSetRequest

`func NewVpnTunnelSetRequest() *VpnTunnelSetRequest`

NewVpnTunnelSetRequest instantiates a new VpnTunnelSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelSetRequestWithDefaults

`func NewVpnTunnelSetRequestWithDefaults() *VpnTunnelSetRequest`

NewVpnTunnelSetRequestWithDefaults instantiates a new VpnTunnelSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VpnTunnelSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnTunnelSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnTunnelSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnTunnelSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VpnTunnelSetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpnTunnelSetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPhase1

`func (o *VpnTunnelSetRequest) GetPhase1() VpnPhase1SetRequest`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *VpnTunnelSetRequest) GetPhase1Ok() (*VpnPhase1SetRequest, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *VpnTunnelSetRequest) SetPhase1(v VpnPhase1SetRequest)`

SetPhase1 sets Phase1 field to given value.

### HasPhase1

`func (o *VpnTunnelSetRequest) HasPhase1() bool`

HasPhase1 returns a boolean if a field has been set.

### SetPhase1Nil

`func (o *VpnTunnelSetRequest) SetPhase1Nil(b bool)`

 SetPhase1Nil sets the value for Phase1 to be an explicit nil

### UnsetPhase1
`func (o *VpnTunnelSetRequest) UnsetPhase1()`

UnsetPhase1 ensures that no value is present for Phase1, not even an explicit nil
### GetPhase2

`func (o *VpnTunnelSetRequest) GetPhase2() VpnPhase2SetRequest`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *VpnTunnelSetRequest) GetPhase2Ok() (*VpnPhase2SetRequest, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *VpnTunnelSetRequest) SetPhase2(v VpnPhase2SetRequest)`

SetPhase2 sets Phase2 field to given value.

### HasPhase2

`func (o *VpnTunnelSetRequest) HasPhase2() bool`

HasPhase2 returns a boolean if a field has been set.

### SetPhase2Nil

`func (o *VpnTunnelSetRequest) SetPhase2Nil(b bool)`

 SetPhase2Nil sets the value for Phase2 to be an explicit nil

### UnsetPhase2
`func (o *VpnTunnelSetRequest) UnsetPhase2()`

UnsetPhase2 ensures that no value is present for Phase2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


