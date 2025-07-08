# Firewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**FlavorName** | Pointer to **string** | Firewall Size Name | [optional] 
**FlavorRuleQuota** | Pointer to **int32** | Firewall Rule Quota | [optional] [default to 0]
**FwResourceId** | **string** | Fw Resource ID | 
**Id** | **string** | Firewall ID | 
**Loggable** | **bool** | Logging Use | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Firewall Name | 
**PreProductId** | Pointer to **string** | Pre Product ID | [optional] 
**ProductType** | [**FirewallProductType**](FirewallProductType.md) | Firewall Product Type | 
**State** | [**FirewallState**](FirewallState.md) | Firewall State | 
**Status** | [**FirewallStatusType**](FirewallStatusType.md) | Firewall Status | 
**TotalRuleCount** | Pointer to **int32** | Total Rule Count | [optional] [default to 0]
**VpcId** | **NullableString** |  | 
**VpcName** | **NullableString** |  | 

## Methods

### NewFirewall

`func NewFirewall(accountId string, createdAt time.Time, createdBy string, fwResourceId string, id string, loggable bool, modifiedAt time.Time, modifiedBy string, name string, productType FirewallProductType, state FirewallState, status FirewallStatusType, vpcId NullableString, vpcName NullableString, ) *Firewall`

NewFirewall instantiates a new Firewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallWithDefaults

`func NewFirewallWithDefaults() *Firewall`

NewFirewallWithDefaults instantiates a new Firewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Firewall) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Firewall) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Firewall) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *Firewall) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Firewall) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Firewall) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Firewall) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Firewall) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Firewall) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetFlavorName

`func (o *Firewall) GetFlavorName() string`

GetFlavorName returns the FlavorName field if non-nil, zero value otherwise.

### GetFlavorNameOk

`func (o *Firewall) GetFlavorNameOk() (*string, bool)`

GetFlavorNameOk returns a tuple with the FlavorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorName

`func (o *Firewall) SetFlavorName(v string)`

SetFlavorName sets FlavorName field to given value.

### HasFlavorName

`func (o *Firewall) HasFlavorName() bool`

HasFlavorName returns a boolean if a field has been set.

### GetFlavorRuleQuota

`func (o *Firewall) GetFlavorRuleQuota() int32`

GetFlavorRuleQuota returns the FlavorRuleQuota field if non-nil, zero value otherwise.

### GetFlavorRuleQuotaOk

`func (o *Firewall) GetFlavorRuleQuotaOk() (*int32, bool)`

GetFlavorRuleQuotaOk returns a tuple with the FlavorRuleQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorRuleQuota

`func (o *Firewall) SetFlavorRuleQuota(v int32)`

SetFlavorRuleQuota sets FlavorRuleQuota field to given value.

### HasFlavorRuleQuota

`func (o *Firewall) HasFlavorRuleQuota() bool`

HasFlavorRuleQuota returns a boolean if a field has been set.

### GetFwResourceId

`func (o *Firewall) GetFwResourceId() string`

GetFwResourceId returns the FwResourceId field if non-nil, zero value otherwise.

### GetFwResourceIdOk

`func (o *Firewall) GetFwResourceIdOk() (*string, bool)`

GetFwResourceIdOk returns a tuple with the FwResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwResourceId

`func (o *Firewall) SetFwResourceId(v string)`

SetFwResourceId sets FwResourceId field to given value.


### GetId

`func (o *Firewall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Firewall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Firewall) SetId(v string)`

SetId sets Id field to given value.


### GetLoggable

`func (o *Firewall) GetLoggable() bool`

GetLoggable returns the Loggable field if non-nil, zero value otherwise.

### GetLoggableOk

`func (o *Firewall) GetLoggableOk() (*bool, bool)`

GetLoggableOk returns a tuple with the Loggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggable

`func (o *Firewall) SetLoggable(v bool)`

SetLoggable sets Loggable field to given value.


### GetModifiedAt

`func (o *Firewall) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Firewall) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Firewall) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Firewall) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Firewall) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Firewall) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *Firewall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Firewall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Firewall) SetName(v string)`

SetName sets Name field to given value.


### GetPreProductId

`func (o *Firewall) GetPreProductId() string`

GetPreProductId returns the PreProductId field if non-nil, zero value otherwise.

### GetPreProductIdOk

`func (o *Firewall) GetPreProductIdOk() (*string, bool)`

GetPreProductIdOk returns a tuple with the PreProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProductId

`func (o *Firewall) SetPreProductId(v string)`

SetPreProductId sets PreProductId field to given value.

### HasPreProductId

`func (o *Firewall) HasPreProductId() bool`

HasPreProductId returns a boolean if a field has been set.

### GetProductType

`func (o *Firewall) GetProductType() FirewallProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Firewall) GetProductTypeOk() (*FirewallProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Firewall) SetProductType(v FirewallProductType)`

SetProductType sets ProductType field to given value.


### GetState

`func (o *Firewall) GetState() FirewallState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Firewall) GetStateOk() (*FirewallState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Firewall) SetState(v FirewallState)`

SetState sets State field to given value.


### GetStatus

`func (o *Firewall) GetStatus() FirewallStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Firewall) GetStatusOk() (*FirewallStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Firewall) SetStatus(v FirewallStatusType)`

SetStatus sets Status field to given value.


### GetTotalRuleCount

`func (o *Firewall) GetTotalRuleCount() int32`

GetTotalRuleCount returns the TotalRuleCount field if non-nil, zero value otherwise.

### GetTotalRuleCountOk

`func (o *Firewall) GetTotalRuleCountOk() (*int32, bool)`

GetTotalRuleCountOk returns a tuple with the TotalRuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuleCount

`func (o *Firewall) SetTotalRuleCount(v int32)`

SetTotalRuleCount sets TotalRuleCount field to given value.

### HasTotalRuleCount

`func (o *Firewall) HasTotalRuleCount() bool`

HasTotalRuleCount returns a boolean if a field has been set.

### GetVpcId

`func (o *Firewall) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *Firewall) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *Firewall) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### SetVpcIdNil

`func (o *Firewall) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *Firewall) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *Firewall) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *Firewall) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *Firewall) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.


### SetVpcNameNil

`func (o *Firewall) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *Firewall) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


