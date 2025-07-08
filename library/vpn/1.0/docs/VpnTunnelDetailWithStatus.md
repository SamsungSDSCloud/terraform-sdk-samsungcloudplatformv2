# VpnTunnelDetailWithStatus

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
**Phase1** | [**VpnPhase1Detail**](VpnPhase1Detail.md) | VPN Tunnel Phase1 | 
**Phase2** | [**VpnPhase2Detail**](VpnPhase2Detail.md) | VPN Tunnel Phase2 | 
**State** | [**VpnTunnelState**](VpnTunnelState.md) | State | 
**Status** | [**VpnTunnelStatus**](VpnTunnelStatus.md) | VPN Tunnel Status | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 
**VpnGatewayId** | **string** | VPN Gateway ID | 
**VpnGatewayIpAddress** | **string** | VPN Gateway IP Address | 
**VpnGatewayName** | **string** | VPN Gateway Name | 

## Methods

### NewVpnTunnelDetailWithStatus

`func NewVpnTunnelDetailWithStatus(accountId string, createdAt time.Time, createdBy string, description NullableString, id string, modifiedAt time.Time, modifiedBy string, name string, phase1 VpnPhase1Detail, phase2 VpnPhase2Detail, state VpnTunnelState, status VpnTunnelStatus, vpcId string, vpcName string, vpnGatewayId string, vpnGatewayIpAddress string, vpnGatewayName string, ) *VpnTunnelDetailWithStatus`

NewVpnTunnelDetailWithStatus instantiates a new VpnTunnelDetailWithStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelDetailWithStatusWithDefaults

`func NewVpnTunnelDetailWithStatusWithDefaults() *VpnTunnelDetailWithStatus`

NewVpnTunnelDetailWithStatusWithDefaults instantiates a new VpnTunnelDetailWithStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VpnTunnelDetailWithStatus) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VpnTunnelDetailWithStatus) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VpnTunnelDetailWithStatus) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *VpnTunnelDetailWithStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpnTunnelDetailWithStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpnTunnelDetailWithStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VpnTunnelDetailWithStatus) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VpnTunnelDetailWithStatus) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VpnTunnelDetailWithStatus) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *VpnTunnelDetailWithStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnTunnelDetailWithStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnTunnelDetailWithStatus) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *VpnTunnelDetailWithStatus) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpnTunnelDetailWithStatus) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *VpnTunnelDetailWithStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnTunnelDetailWithStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnTunnelDetailWithStatus) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *VpnTunnelDetailWithStatus) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VpnTunnelDetailWithStatus) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VpnTunnelDetailWithStatus) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *VpnTunnelDetailWithStatus) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VpnTunnelDetailWithStatus) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VpnTunnelDetailWithStatus) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *VpnTunnelDetailWithStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnTunnelDetailWithStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnTunnelDetailWithStatus) SetName(v string)`

SetName sets Name field to given value.


### GetPhase1

`func (o *VpnTunnelDetailWithStatus) GetPhase1() VpnPhase1Detail`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *VpnTunnelDetailWithStatus) GetPhase1Ok() (*VpnPhase1Detail, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *VpnTunnelDetailWithStatus) SetPhase1(v VpnPhase1Detail)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *VpnTunnelDetailWithStatus) GetPhase2() VpnPhase2Detail`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *VpnTunnelDetailWithStatus) GetPhase2Ok() (*VpnPhase2Detail, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *VpnTunnelDetailWithStatus) SetPhase2(v VpnPhase2Detail)`

SetPhase2 sets Phase2 field to given value.


### GetState

`func (o *VpnTunnelDetailWithStatus) GetState() VpnTunnelState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpnTunnelDetailWithStatus) GetStateOk() (*VpnTunnelState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpnTunnelDetailWithStatus) SetState(v VpnTunnelState)`

SetState sets State field to given value.


### GetStatus

`func (o *VpnTunnelDetailWithStatus) GetStatus() VpnTunnelStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnTunnelDetailWithStatus) GetStatusOk() (*VpnTunnelStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnTunnelDetailWithStatus) SetStatus(v VpnTunnelStatus)`

SetStatus sets Status field to given value.


### GetVpcId

`func (o *VpnTunnelDetailWithStatus) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *VpnTunnelDetailWithStatus) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *VpnTunnelDetailWithStatus) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *VpnTunnelDetailWithStatus) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *VpnTunnelDetailWithStatus) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *VpnTunnelDetailWithStatus) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.


### GetVpnGatewayId

`func (o *VpnTunnelDetailWithStatus) GetVpnGatewayId() string`

GetVpnGatewayId returns the VpnGatewayId field if non-nil, zero value otherwise.

### GetVpnGatewayIdOk

`func (o *VpnTunnelDetailWithStatus) GetVpnGatewayIdOk() (*string, bool)`

GetVpnGatewayIdOk returns a tuple with the VpnGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayId

`func (o *VpnTunnelDetailWithStatus) SetVpnGatewayId(v string)`

SetVpnGatewayId sets VpnGatewayId field to given value.


### GetVpnGatewayIpAddress

`func (o *VpnTunnelDetailWithStatus) GetVpnGatewayIpAddress() string`

GetVpnGatewayIpAddress returns the VpnGatewayIpAddress field if non-nil, zero value otherwise.

### GetVpnGatewayIpAddressOk

`func (o *VpnTunnelDetailWithStatus) GetVpnGatewayIpAddressOk() (*string, bool)`

GetVpnGatewayIpAddressOk returns a tuple with the VpnGatewayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayIpAddress

`func (o *VpnTunnelDetailWithStatus) SetVpnGatewayIpAddress(v string)`

SetVpnGatewayIpAddress sets VpnGatewayIpAddress field to given value.


### GetVpnGatewayName

`func (o *VpnTunnelDetailWithStatus) GetVpnGatewayName() string`

GetVpnGatewayName returns the VpnGatewayName field if non-nil, zero value otherwise.

### GetVpnGatewayNameOk

`func (o *VpnTunnelDetailWithStatus) GetVpnGatewayNameOk() (*string, bool)`

GetVpnGatewayNameOk returns a tuple with the VpnGatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayName

`func (o *VpnTunnelDetailWithStatus) SetVpnGatewayName(v string)`

SetVpnGatewayName sets VpnGatewayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


