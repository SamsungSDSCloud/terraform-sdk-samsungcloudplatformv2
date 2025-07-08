# VpnTunnelCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | VPN Tunnel Name | 
**Phase1** | [**VpnPhase1CreateRequest**](VpnPhase1CreateRequest.md) | VPN Tunnel Phase1 | 
**Phase2** | [**VpnPhase2CreateRequest**](VpnPhase2CreateRequest.md) | VPN Tunnel Phase2 | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**VpnGatewayId** | **string** | VPN Gateway ID | 

## Methods

### NewVpnTunnelCreateRequest

`func NewVpnTunnelCreateRequest(name string, phase1 VpnPhase1CreateRequest, phase2 VpnPhase2CreateRequest, vpnGatewayId string, ) *VpnTunnelCreateRequest`

NewVpnTunnelCreateRequest instantiates a new VpnTunnelCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelCreateRequestWithDefaults

`func NewVpnTunnelCreateRequestWithDefaults() *VpnTunnelCreateRequest`

NewVpnTunnelCreateRequestWithDefaults instantiates a new VpnTunnelCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VpnTunnelCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnTunnelCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnTunnelCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnTunnelCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VpnTunnelCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpnTunnelCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *VpnTunnelCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnTunnelCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnTunnelCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPhase1

`func (o *VpnTunnelCreateRequest) GetPhase1() VpnPhase1CreateRequest`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *VpnTunnelCreateRequest) GetPhase1Ok() (*VpnPhase1CreateRequest, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *VpnTunnelCreateRequest) SetPhase1(v VpnPhase1CreateRequest)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *VpnTunnelCreateRequest) GetPhase2() VpnPhase2CreateRequest`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *VpnTunnelCreateRequest) GetPhase2Ok() (*VpnPhase2CreateRequest, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *VpnTunnelCreateRequest) SetPhase2(v VpnPhase2CreateRequest)`

SetPhase2 sets Phase2 field to given value.


### GetTags

`func (o *VpnTunnelCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnTunnelCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnTunnelCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnTunnelCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVpnGatewayId

`func (o *VpnTunnelCreateRequest) GetVpnGatewayId() string`

GetVpnGatewayId returns the VpnGatewayId field if non-nil, zero value otherwise.

### GetVpnGatewayIdOk

`func (o *VpnTunnelCreateRequest) GetVpnGatewayIdOk() (*string, bool)`

GetVpnGatewayIdOk returns a tuple with the VpnGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayId

`func (o *VpnTunnelCreateRequest) SetVpnGatewayId(v string)`

SetVpnGatewayId sets VpnGatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


