# NatGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | NAT Gateway ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | NAT Gateway Name | 
**NatGatewayIpAddress** | **string** | NAT Gateway IP Address | 
**PublicipId** | Pointer to **NullableString** |  | [optional] 
**State** | **string** | NAT Gateway State | 
**SubnetCidr** | **string** | Subnet Cidr | 
**SubnetId** | **string** | Subnet Id | 
**SubnetName** | **string** | Subnet Name | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 

## Methods

### NewNatGateway

`func NewNatGateway(accountId string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, name string, natGatewayIpAddress string, state string, subnetCidr string, subnetId string, subnetName string, vpcId string, vpcName string, ) *NatGateway`

NewNatGateway instantiates a new NatGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatGatewayWithDefaults

`func NewNatGatewayWithDefaults() *NatGateway`

NewNatGatewayWithDefaults instantiates a new NatGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NatGateway) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NatGateway) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NatGateway) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *NatGateway) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NatGateway) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NatGateway) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *NatGateway) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NatGateway) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *NatGateway) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *NatGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NatGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NatGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NatGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NatGateway) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NatGateway) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *NatGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NatGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NatGateway) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *NatGateway) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *NatGateway) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *NatGateway) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *NatGateway) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *NatGateway) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *NatGateway) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *NatGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatGateway) SetName(v string)`

SetName sets Name field to given value.


### GetNatGatewayIpAddress

`func (o *NatGateway) GetNatGatewayIpAddress() string`

GetNatGatewayIpAddress returns the NatGatewayIpAddress field if non-nil, zero value otherwise.

### GetNatGatewayIpAddressOk

`func (o *NatGateway) GetNatGatewayIpAddressOk() (*string, bool)`

GetNatGatewayIpAddressOk returns a tuple with the NatGatewayIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatGatewayIpAddress

`func (o *NatGateway) SetNatGatewayIpAddress(v string)`

SetNatGatewayIpAddress sets NatGatewayIpAddress field to given value.


### GetPublicipId

`func (o *NatGateway) GetPublicipId() string`

GetPublicipId returns the PublicipId field if non-nil, zero value otherwise.

### GetPublicipIdOk

`func (o *NatGateway) GetPublicipIdOk() (*string, bool)`

GetPublicipIdOk returns a tuple with the PublicipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicipId

`func (o *NatGateway) SetPublicipId(v string)`

SetPublicipId sets PublicipId field to given value.

### HasPublicipId

`func (o *NatGateway) HasPublicipId() bool`

HasPublicipId returns a boolean if a field has been set.

### SetPublicipIdNil

`func (o *NatGateway) SetPublicipIdNil(b bool)`

 SetPublicipIdNil sets the value for PublicipId to be an explicit nil

### UnsetPublicipId
`func (o *NatGateway) UnsetPublicipId()`

UnsetPublicipId ensures that no value is present for PublicipId, not even an explicit nil
### GetState

`func (o *NatGateway) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NatGateway) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NatGateway) SetState(v string)`

SetState sets State field to given value.


### GetSubnetCidr

`func (o *NatGateway) GetSubnetCidr() string`

GetSubnetCidr returns the SubnetCidr field if non-nil, zero value otherwise.

### GetSubnetCidrOk

`func (o *NatGateway) GetSubnetCidrOk() (*string, bool)`

GetSubnetCidrOk returns a tuple with the SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidr

`func (o *NatGateway) SetSubnetCidr(v string)`

SetSubnetCidr sets SubnetCidr field to given value.


### GetSubnetId

`func (o *NatGateway) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *NatGateway) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *NatGateway) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSubnetName

`func (o *NatGateway) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *NatGateway) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *NatGateway) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.


### GetVpcId

`func (o *NatGateway) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *NatGateway) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *NatGateway) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *NatGateway) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *NatGateway) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *NatGateway) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


