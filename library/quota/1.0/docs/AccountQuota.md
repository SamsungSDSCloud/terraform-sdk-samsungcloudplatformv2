# AccountQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Unique identifier for the account | 
**AccountName** | **string** | Name of the account | 
**Adjustable** | **bool** | Flag indicating if additional quota is being requested | 
**AppliedValue** | Pointer to **NullableInt32** |  | [optional] 
**Approval** | **bool** | Approval | 
**ClassValue** | **string** | Value associated with the request class | 
**CreatedAt** | **time.Time** | Created At | 
**Description** | **string** | Detailed description of the quota item | 
**FreeRate** | **int32** | Free Rate | 
**Id** | **string** | Account Quota ID | 
**InitialValue** | **int32** | Initial quota value allocated | 
**ModifiedAt** | **time.Time** | Modified At | 
**QuotaItem** | **string** | Specific quota item within the resource | 
**Reduction** | Pointer to **bool** | Reduction | [optional] [default to false]
**Request** | Pointer to **bool** | Reqeust  | [optional] [default to false]
**RequestClass** | **string** | Classification of the quota request (e.g., Account, Region) | 
**ResourceType** | **string** | Type of the resource (e.g., Virtual Server, Storage) | 
**Service** | **string** | Name of the service to which quota applies | 
**Srn** | **string** | Service Resource Name for the quota item | 
**Unit** | **string** | Unit in which the quota value is measured (e.g., EA, GB) | 

## Methods

### NewAccountQuota

`func NewAccountQuota(accountId string, accountName string, adjustable bool, approval bool, classValue string, createdAt time.Time, description string, freeRate int32, id string, initialValue int32, modifiedAt time.Time, quotaItem string, requestClass string, resourceType string, service string, srn string, unit string, ) *AccountQuota`

NewAccountQuota instantiates a new AccountQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountQuotaWithDefaults

`func NewAccountQuotaWithDefaults() *AccountQuota`

NewAccountQuotaWithDefaults instantiates a new AccountQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountQuota) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountQuota) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountQuota) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccountName

`func (o *AccountQuota) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AccountQuota) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AccountQuota) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetAdjustable

`func (o *AccountQuota) GetAdjustable() bool`

GetAdjustable returns the Adjustable field if non-nil, zero value otherwise.

### GetAdjustableOk

`func (o *AccountQuota) GetAdjustableOk() (*bool, bool)`

GetAdjustableOk returns a tuple with the Adjustable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustable

`func (o *AccountQuota) SetAdjustable(v bool)`

SetAdjustable sets Adjustable field to given value.


### GetAppliedValue

`func (o *AccountQuota) GetAppliedValue() int32`

GetAppliedValue returns the AppliedValue field if non-nil, zero value otherwise.

### GetAppliedValueOk

`func (o *AccountQuota) GetAppliedValueOk() (*int32, bool)`

GetAppliedValueOk returns a tuple with the AppliedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedValue

`func (o *AccountQuota) SetAppliedValue(v int32)`

SetAppliedValue sets AppliedValue field to given value.

### HasAppliedValue

`func (o *AccountQuota) HasAppliedValue() bool`

HasAppliedValue returns a boolean if a field has been set.

### SetAppliedValueNil

`func (o *AccountQuota) SetAppliedValueNil(b bool)`

 SetAppliedValueNil sets the value for AppliedValue to be an explicit nil

### UnsetAppliedValue
`func (o *AccountQuota) UnsetAppliedValue()`

UnsetAppliedValue ensures that no value is present for AppliedValue, not even an explicit nil
### GetApproval

`func (o *AccountQuota) GetApproval() bool`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *AccountQuota) GetApprovalOk() (*bool, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *AccountQuota) SetApproval(v bool)`

SetApproval sets Approval field to given value.


### GetClassValue

`func (o *AccountQuota) GetClassValue() string`

GetClassValue returns the ClassValue field if non-nil, zero value otherwise.

### GetClassValueOk

`func (o *AccountQuota) GetClassValueOk() (*string, bool)`

GetClassValueOk returns a tuple with the ClassValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassValue

`func (o *AccountQuota) SetClassValue(v string)`

SetClassValue sets ClassValue field to given value.


### GetCreatedAt

`func (o *AccountQuota) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountQuota) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountQuota) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *AccountQuota) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountQuota) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountQuota) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFreeRate

`func (o *AccountQuota) GetFreeRate() int32`

GetFreeRate returns the FreeRate field if non-nil, zero value otherwise.

### GetFreeRateOk

`func (o *AccountQuota) GetFreeRateOk() (*int32, bool)`

GetFreeRateOk returns a tuple with the FreeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeRate

`func (o *AccountQuota) SetFreeRate(v int32)`

SetFreeRate sets FreeRate field to given value.


### GetId

`func (o *AccountQuota) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountQuota) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountQuota) SetId(v string)`

SetId sets Id field to given value.


### GetInitialValue

`func (o *AccountQuota) GetInitialValue() int32`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *AccountQuota) GetInitialValueOk() (*int32, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *AccountQuota) SetInitialValue(v int32)`

SetInitialValue sets InitialValue field to given value.


### GetModifiedAt

`func (o *AccountQuota) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AccountQuota) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AccountQuota) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetQuotaItem

`func (o *AccountQuota) GetQuotaItem() string`

GetQuotaItem returns the QuotaItem field if non-nil, zero value otherwise.

### GetQuotaItemOk

`func (o *AccountQuota) GetQuotaItemOk() (*string, bool)`

GetQuotaItemOk returns a tuple with the QuotaItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaItem

`func (o *AccountQuota) SetQuotaItem(v string)`

SetQuotaItem sets QuotaItem field to given value.


### GetReduction

`func (o *AccountQuota) GetReduction() bool`

GetReduction returns the Reduction field if non-nil, zero value otherwise.

### GetReductionOk

`func (o *AccountQuota) GetReductionOk() (*bool, bool)`

GetReductionOk returns a tuple with the Reduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduction

`func (o *AccountQuota) SetReduction(v bool)`

SetReduction sets Reduction field to given value.

### HasReduction

`func (o *AccountQuota) HasReduction() bool`

HasReduction returns a boolean if a field has been set.

### GetRequest

`func (o *AccountQuota) GetRequest() bool`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AccountQuota) GetRequestOk() (*bool, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AccountQuota) SetRequest(v bool)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *AccountQuota) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetRequestClass

`func (o *AccountQuota) GetRequestClass() string`

GetRequestClass returns the RequestClass field if non-nil, zero value otherwise.

### GetRequestClassOk

`func (o *AccountQuota) GetRequestClassOk() (*string, bool)`

GetRequestClassOk returns a tuple with the RequestClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestClass

`func (o *AccountQuota) SetRequestClass(v string)`

SetRequestClass sets RequestClass field to given value.


### GetResourceType

`func (o *AccountQuota) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AccountQuota) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AccountQuota) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetService

`func (o *AccountQuota) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AccountQuota) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AccountQuota) SetService(v string)`

SetService sets Service field to given value.


### GetSrn

`func (o *AccountQuota) GetSrn() string`

GetSrn returns the Srn field if non-nil, zero value otherwise.

### GetSrnOk

`func (o *AccountQuota) GetSrnOk() (*string, bool)`

GetSrnOk returns a tuple with the Srn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrn

`func (o *AccountQuota) SetSrn(v string)`

SetSrn sets Srn field to given value.


### GetUnit

`func (o *AccountQuota) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AccountQuota) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AccountQuota) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


