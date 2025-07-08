# VpnTunnel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | **NullableString** |  | 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | VPN Tunnel Name | 
**Phase1** | [**VpnPhase1**](VpnPhase1.md) | VPN Tunnel Phase1 | 
**Phase2** | [**VpnPhase2**](VpnPhase2.md) | VPN Tunnel Phase2 | 
**State** | [**VpnTunnelState**](VpnTunnelState.md) | State | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 
**VpnGatewayId** | **string** | VPN Gateway ID | 
**VpnGatewayIpAddress** | **string** | VPN Gateway IP Address | 
**VpnGatewayName** | **string** | VPN Gateway Name | 

## Methods

### NewVpnTunnel

`func NewVpnTunnel(accountId string, createdAt time.Time, createdBy string, description NullableString, id string, modifiedAt time.Time, modifiedBy string, name string, phase1 VpnPhase1, phase2 VpnPhase2, state VpnTunnelState, vpcId string, vpcName string, vpnGatewayId string, vpnGatewayIpAddress string, vpnGatewayName string, ) *VpnTunnel`

NewVpnTunnel instantiates a new VpnTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelWithDefaults

`func NewVpnTunnelWithDefaults() *VpnTunnel`

NewVpnTunnelWithDefaults instantiates a new VpnTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VpnTunnel) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VpnTunnel) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VpnTunnel) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *VpnTunnel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpnTunnel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpnTunnel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VpnTunnel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VpnTunnel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VpnTunnel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *VpnTunnel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnTunnel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnTunnel) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *VpnTunnel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpnTunnel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *VpnTunnel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnTunnel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnTunnel) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *VpnTunnel) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VpnTunnel) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VpnTunnel) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *VpnTunnel) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VpnTunnel) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VpnTunnel) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *VpnTunnel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnTunnel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnTunnel) SetName(v string)`

SetName sets Name field to given value.


### GetPhase1

`func (o *VpnTunnel) GetPhase1() VpnPhase1`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *VpnTunnel) GetPhase1Ok() (*VpnPhase1, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *VpnTunnel) SetPhase1(v VpnPhase1)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *VpnTunnel) GetPhase2() VpnPhase2`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *VpnTunnel) GetPhase2Ok() (*VpnPhase2, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *VpnTunnel) SetPhase2(v VpnPhase2)`

SetPhase2 sets Phase2 field to given value.


### GetState

`func (o *VpnTunnel) GetState() VpnTunnelState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpnTunnel) GetStateOk() (*VpnTunnelState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpnTunnel) SetState(v VpnTunnelState)`

SetState sets State field to given value.


### GetVpcId

`func (o *VpnTunnel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *VpnTunnel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *VpnTunnel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *VpnTunnel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *VpnTunnel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *VpnTunnel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.


### GetVpnGatewayId

`func (o *VpnTunnel) GetVpnGatewayId() string`

GetVpnGatewayId returns the VpnGatewayId field if non-nil, zero value otherwise.

### GetVpnGatewayIdOk

`func (o *VpnTunnel) GetVpnGatewayIdOk() (*string, bool)`

GetVpnGatewayIdOk returns a tuple with the VpnGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayId

`func (o *VpnTunnel) SetVpnGatewayId(v string)`

SetVpnGatewayId sets VpnGatewayId field to given value.


### GetVpnGatewayIpAddress

`func (o *VpnTunnel) GetVpnGatewayIpAddress() string`

GetVpnGatewayIpAddress returns the VpnGatewayIpAddress field if non-nil, zero value otherwise.

### GetVpnGatewayIpAddressOk

`func (o *VpnTunnel) GetVpnGatewayIpAddressOk() (*string, bool)`

GetVpnGatewayIpAddressOk returns a tuple with the VpnGatewayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayIpAddress

`func (o *VpnTunnel) SetVpnGatewayIpAddress(v string)`

SetVpnGatewayIpAddress sets VpnGatewayIpAddress field to given value.


### GetVpnGatewayName

`func (o *VpnTunnel) GetVpnGatewayName() string`

GetVpnGatewayName returns the VpnGatewayName field if non-nil, zero value otherwise.

### GetVpnGatewayNameOk

`func (o *VpnTunnel) GetVpnGatewayNameOk() (*string, bool)`

GetVpnGatewayNameOk returns a tuple with the VpnGatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayName

`func (o *VpnTunnel) SetVpnGatewayName(v string)`

SetVpnGatewayName sets VpnGatewayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


