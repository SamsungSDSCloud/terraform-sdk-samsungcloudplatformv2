# VpnTunnelDetail

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
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 
**VpnGatewayId** | **string** | VPN Gateway ID | 
**VpnGatewayIpAddress** | **string** | VPN Gateway IP Address | 
**VpnGatewayName** | **string** | VPN Gateway Name | 

## Methods

### NewVpnTunnelDetail

`func NewVpnTunnelDetail(accountId string, createdAt time.Time, createdBy string, description NullableString, id string, modifiedAt time.Time, modifiedBy string, name string, phase1 VpnPhase1Detail, phase2 VpnPhase2Detail, state VpnTunnelState, vpcId string, vpcName string, vpnGatewayId string, vpnGatewayIpAddress string, vpnGatewayName string, ) *VpnTunnelDetail`

NewVpnTunnelDetail instantiates a new VpnTunnelDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelDetailWithDefaults

`func NewVpnTunnelDetailWithDefaults() *VpnTunnelDetail`

NewVpnTunnelDetailWithDefaults instantiates a new VpnTunnelDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VpnTunnelDetail) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VpnTunnelDetail) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VpnTunnelDetail) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *VpnTunnelDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VpnTunnelDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VpnTunnelDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VpnTunnelDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VpnTunnelDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VpnTunnelDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *VpnTunnelDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnTunnelDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnTunnelDetail) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *VpnTunnelDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpnTunnelDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *VpnTunnelDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnTunnelDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnTunnelDetail) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *VpnTunnelDetail) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VpnTunnelDetail) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VpnTunnelDetail) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *VpnTunnelDetail) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VpnTunnelDetail) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VpnTunnelDetail) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *VpnTunnelDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnTunnelDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnTunnelDetail) SetName(v string)`

SetName sets Name field to given value.


### GetPhase1

`func (o *VpnTunnelDetail) GetPhase1() VpnPhase1Detail`

GetPhase1 returns the Phase1 field if non-nil, zero value otherwise.

### GetPhase1Ok

`func (o *VpnTunnelDetail) GetPhase1Ok() (*VpnPhase1Detail, bool)`

GetPhase1Ok returns a tuple with the Phase1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1

`func (o *VpnTunnelDetail) SetPhase1(v VpnPhase1Detail)`

SetPhase1 sets Phase1 field to given value.


### GetPhase2

`func (o *VpnTunnelDetail) GetPhase2() VpnPhase2Detail`

GetPhase2 returns the Phase2 field if non-nil, zero value otherwise.

### GetPhase2Ok

`func (o *VpnTunnelDetail) GetPhase2Ok() (*VpnPhase2Detail, bool)`

GetPhase2Ok returns a tuple with the Phase2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2

`func (o *VpnTunnelDetail) SetPhase2(v VpnPhase2Detail)`

SetPhase2 sets Phase2 field to given value.


### GetState

`func (o *VpnTunnelDetail) GetState() VpnTunnelState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpnTunnelDetail) GetStateOk() (*VpnTunnelState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VpnTunnelDetail) SetState(v VpnTunnelState)`

SetState sets State field to given value.


### GetVpcId

`func (o *VpnTunnelDetail) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *VpnTunnelDetail) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *VpnTunnelDetail) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *VpnTunnelDetail) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *VpnTunnelDetail) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *VpnTunnelDetail) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.


### GetVpnGatewayId

`func (o *VpnTunnelDetail) GetVpnGatewayId() string`

GetVpnGatewayId returns the VpnGatewayId field if non-nil, zero value otherwise.

### GetVpnGatewayIdOk

`func (o *VpnTunnelDetail) GetVpnGatewayIdOk() (*string, bool)`

GetVpnGatewayIdOk returns a tuple with the VpnGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayId

`func (o *VpnTunnelDetail) SetVpnGatewayId(v string)`

SetVpnGatewayId sets VpnGatewayId field to given value.


### GetVpnGatewayIpAddress

`func (o *VpnTunnelDetail) GetVpnGatewayIpAddress() string`

GetVpnGatewayIpAddress returns the VpnGatewayIpAddress field if non-nil, zero value otherwise.

### GetVpnGatewayIpAddressOk

`func (o *VpnTunnelDetail) GetVpnGatewayIpAddressOk() (*string, bool)`

GetVpnGatewayIpAddressOk returns a tuple with the VpnGatewayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayIpAddress

`func (o *VpnTunnelDetail) SetVpnGatewayIpAddress(v string)`

SetVpnGatewayIpAddress sets VpnGatewayIpAddress field to given value.


### GetVpnGatewayName

`func (o *VpnTunnelDetail) GetVpnGatewayName() string`

GetVpnGatewayName returns the VpnGatewayName field if non-nil, zero value otherwise.

### GetVpnGatewayNameOk

`func (o *VpnTunnelDetail) GetVpnGatewayNameOk() (*string, bool)`

GetVpnGatewayNameOk returns a tuple with the VpnGatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayName

`func (o *VpnTunnelDetail) SetVpnGatewayName(v string)`

SetVpnGatewayName sets VpnGatewayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


