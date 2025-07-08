# DirectConnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**Bandwidth** | **int32** | Port Bandwidth | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**FirewallId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Direct Connect ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Direct Connect Name | 
**State** | [**DirectConnectState**](DirectConnectState.md) | State | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 

## Methods

### NewDirectConnect

`func NewDirectConnect(accountId string, bandwidth int32, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, name string, state DirectConnectState, vpcId string, vpcName string, ) *DirectConnect`

NewDirectConnect instantiates a new DirectConnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectWithDefaults

`func NewDirectConnectWithDefaults() *DirectConnect`

NewDirectConnectWithDefaults instantiates a new DirectConnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *DirectConnect) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DirectConnect) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *DirectConnect) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBandwidth

`func (o *DirectConnect) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *DirectConnect) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *DirectConnect) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.


### GetCreatedAt

`func (o *DirectConnect) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DirectConnect) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DirectConnect) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DirectConnect) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DirectConnect) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DirectConnect) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *DirectConnect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DirectConnect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DirectConnect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DirectConnect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DirectConnect) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DirectConnect) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFirewallId

`func (o *DirectConnect) GetFirewallId() string`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *DirectConnect) GetFirewallIdOk() (*string, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *DirectConnect) SetFirewallId(v string)`

SetFirewallId sets FirewallId field to given value.

### HasFirewallId

`func (o *DirectConnect) HasFirewallId() bool`

HasFirewallId returns a boolean if a field has been set.

### SetFirewallIdNil

`func (o *DirectConnect) SetFirewallIdNil(b bool)`

 SetFirewallIdNil sets the value for FirewallId to be an explicit nil

### UnsetFirewallId
`func (o *DirectConnect) UnsetFirewallId()`

UnsetFirewallId ensures that no value is present for FirewallId, not even an explicit nil
### GetId

`func (o *DirectConnect) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DirectConnect) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DirectConnect) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *DirectConnect) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DirectConnect) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DirectConnect) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *DirectConnect) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *DirectConnect) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *DirectConnect) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *DirectConnect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectConnect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectConnect) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *DirectConnect) GetState() DirectConnectState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectConnect) GetStateOk() (*DirectConnectState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DirectConnect) SetState(v DirectConnectState)`

SetState sets State field to given value.


### GetVpcId

`func (o *DirectConnect) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *DirectConnect) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *DirectConnect) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *DirectConnect) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *DirectConnect) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *DirectConnect) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


