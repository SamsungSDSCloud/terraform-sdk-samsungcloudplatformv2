# PrivateNat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**Cidr** | **string** | Private NAT IP range | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DirectConnectId** | **string** | Direct Connect ID | 
**DirectConnectName** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Private NAT ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Private NAT Name | 
**State** | [**PrivateNatState**](PrivateNatState.md) | Private NAT State | 
**VpcId** | **string** | VPC Id | 
**VpcName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrivateNat

`func NewPrivateNat(accountId string, cidr string, createdAt time.Time, createdBy string, directConnectId string, id string, modifiedAt time.Time, modifiedBy string, name string, state PrivateNatState, vpcId string, ) *PrivateNat`

NewPrivateNat instantiates a new PrivateNat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNatWithDefaults

`func NewPrivateNatWithDefaults() *PrivateNat`

NewPrivateNatWithDefaults instantiates a new PrivateNat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PrivateNat) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PrivateNat) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PrivateNat) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCidr

`func (o *PrivateNat) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *PrivateNat) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *PrivateNat) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetCreatedAt

`func (o *PrivateNat) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrivateNat) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrivateNat) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PrivateNat) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrivateNat) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrivateNat) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *PrivateNat) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateNat) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateNat) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateNat) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PrivateNat) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PrivateNat) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDirectConnectId

`func (o *PrivateNat) GetDirectConnectId() string`

GetDirectConnectId returns the DirectConnectId field if non-nil, zero value otherwise.

### GetDirectConnectIdOk

`func (o *PrivateNat) GetDirectConnectIdOk() (*string, bool)`

GetDirectConnectIdOk returns a tuple with the DirectConnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnectId

`func (o *PrivateNat) SetDirectConnectId(v string)`

SetDirectConnectId sets DirectConnectId field to given value.


### GetDirectConnectName

`func (o *PrivateNat) GetDirectConnectName() string`

GetDirectConnectName returns the DirectConnectName field if non-nil, zero value otherwise.

### GetDirectConnectNameOk

`func (o *PrivateNat) GetDirectConnectNameOk() (*string, bool)`

GetDirectConnectNameOk returns a tuple with the DirectConnectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnectName

`func (o *PrivateNat) SetDirectConnectName(v string)`

SetDirectConnectName sets DirectConnectName field to given value.

### HasDirectConnectName

`func (o *PrivateNat) HasDirectConnectName() bool`

HasDirectConnectName returns a boolean if a field has been set.

### SetDirectConnectNameNil

`func (o *PrivateNat) SetDirectConnectNameNil(b bool)`

 SetDirectConnectNameNil sets the value for DirectConnectName to be an explicit nil

### UnsetDirectConnectName
`func (o *PrivateNat) UnsetDirectConnectName()`

UnsetDirectConnectName ensures that no value is present for DirectConnectName, not even an explicit nil
### GetId

`func (o *PrivateNat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateNat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateNat) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *PrivateNat) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PrivateNat) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PrivateNat) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *PrivateNat) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PrivateNat) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PrivateNat) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *PrivateNat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateNat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateNat) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *PrivateNat) GetState() PrivateNatState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PrivateNat) GetStateOk() (*PrivateNatState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PrivateNat) SetState(v PrivateNatState)`

SetState sets State field to given value.


### GetVpcId

`func (o *PrivateNat) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *PrivateNat) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *PrivateNat) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *PrivateNat) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *PrivateNat) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *PrivateNat) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *PrivateNat) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *PrivateNat) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *PrivateNat) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


