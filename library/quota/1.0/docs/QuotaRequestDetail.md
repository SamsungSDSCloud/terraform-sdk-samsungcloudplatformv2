# QuotaRequestDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountQuotaId** | **string** | Account Quota ID | 
**AppliedAt** | Pointer to **NullableTime** |  | [optional] 
**AppliedValue** | **NullableInt32** |  | 
**ApprovalMessage** | Pointer to **NullableString** |  | [optional] 
**ClassValue** | **string** | Value associated with the request class | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | **string** | Detailed description of the quota item | 
**FreeRate** | Pointer to **NullableInt32** |  | [optional] 
**Id** | **string** | ID | 
**InitialValue** | **int32** | Initial value | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**QuotaItem** | **string** | Quota Item Name | 
**Reduction** | **bool** | Reduction | 
**RequestMessage** | **string** | Request message | 
**RequestValue** | **int32** | User Request Value | 
**Requester** | **string** | Requester | 
**State** | [**RequestStateEnum**](RequestStateEnum.md) | Request State | 
**Unit** | **string** | Unit in which the quota value is measured (e.g., EA, GB) | 

## Methods

### NewQuotaRequestDetail

`func NewQuotaRequestDetail(accountQuotaId string, appliedValue NullableInt32, classValue string, createdAt time.Time, createdBy string, description string, id string, initialValue int32, modifiedAt time.Time, modifiedBy string, quotaItem string, reduction bool, requestMessage string, requestValue int32, requester string, state RequestStateEnum, unit string, ) *QuotaRequestDetail`

NewQuotaRequestDetail instantiates a new QuotaRequestDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaRequestDetailWithDefaults

`func NewQuotaRequestDetailWithDefaults() *QuotaRequestDetail`

NewQuotaRequestDetailWithDefaults instantiates a new QuotaRequestDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountQuotaId

`func (o *QuotaRequestDetail) GetAccountQuotaId() string`

GetAccountQuotaId returns the AccountQuotaId field if non-nil, zero value otherwise.

### GetAccountQuotaIdOk

`func (o *QuotaRequestDetail) GetAccountQuotaIdOk() (*string, bool)`

GetAccountQuotaIdOk returns a tuple with the AccountQuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountQuotaId

`func (o *QuotaRequestDetail) SetAccountQuotaId(v string)`

SetAccountQuotaId sets AccountQuotaId field to given value.


### GetAppliedAt

`func (o *QuotaRequestDetail) GetAppliedAt() time.Time`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *QuotaRequestDetail) GetAppliedAtOk() (*time.Time, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *QuotaRequestDetail) SetAppliedAt(v time.Time)`

SetAppliedAt sets AppliedAt field to given value.

### HasAppliedAt

`func (o *QuotaRequestDetail) HasAppliedAt() bool`

HasAppliedAt returns a boolean if a field has been set.

### SetAppliedAtNil

`func (o *QuotaRequestDetail) SetAppliedAtNil(b bool)`

 SetAppliedAtNil sets the value for AppliedAt to be an explicit nil

### UnsetAppliedAt
`func (o *QuotaRequestDetail) UnsetAppliedAt()`

UnsetAppliedAt ensures that no value is present for AppliedAt, not even an explicit nil
### GetAppliedValue

`func (o *QuotaRequestDetail) GetAppliedValue() int32`

GetAppliedValue returns the AppliedValue field if non-nil, zero value otherwise.

### GetAppliedValueOk

`func (o *QuotaRequestDetail) GetAppliedValueOk() (*int32, bool)`

GetAppliedValueOk returns a tuple with the AppliedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedValue

`func (o *QuotaRequestDetail) SetAppliedValue(v int32)`

SetAppliedValue sets AppliedValue field to given value.


### SetAppliedValueNil

`func (o *QuotaRequestDetail) SetAppliedValueNil(b bool)`

 SetAppliedValueNil sets the value for AppliedValue to be an explicit nil

### UnsetAppliedValue
`func (o *QuotaRequestDetail) UnsetAppliedValue()`

UnsetAppliedValue ensures that no value is present for AppliedValue, not even an explicit nil
### GetApprovalMessage

`func (o *QuotaRequestDetail) GetApprovalMessage() string`

GetApprovalMessage returns the ApprovalMessage field if non-nil, zero value otherwise.

### GetApprovalMessageOk

`func (o *QuotaRequestDetail) GetApprovalMessageOk() (*string, bool)`

GetApprovalMessageOk returns a tuple with the ApprovalMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalMessage

`func (o *QuotaRequestDetail) SetApprovalMessage(v string)`

SetApprovalMessage sets ApprovalMessage field to given value.

### HasApprovalMessage

`func (o *QuotaRequestDetail) HasApprovalMessage() bool`

HasApprovalMessage returns a boolean if a field has been set.

### SetApprovalMessageNil

`func (o *QuotaRequestDetail) SetApprovalMessageNil(b bool)`

 SetApprovalMessageNil sets the value for ApprovalMessage to be an explicit nil

### UnsetApprovalMessage
`func (o *QuotaRequestDetail) UnsetApprovalMessage()`

UnsetApprovalMessage ensures that no value is present for ApprovalMessage, not even an explicit nil
### GetClassValue

`func (o *QuotaRequestDetail) GetClassValue() string`

GetClassValue returns the ClassValue field if non-nil, zero value otherwise.

### GetClassValueOk

`func (o *QuotaRequestDetail) GetClassValueOk() (*string, bool)`

GetClassValueOk returns a tuple with the ClassValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassValue

`func (o *QuotaRequestDetail) SetClassValue(v string)`

SetClassValue sets ClassValue field to given value.


### GetCreatedAt

`func (o *QuotaRequestDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuotaRequestDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuotaRequestDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *QuotaRequestDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *QuotaRequestDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *QuotaRequestDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *QuotaRequestDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotaRequestDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotaRequestDetail) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFreeRate

`func (o *QuotaRequestDetail) GetFreeRate() int32`

GetFreeRate returns the FreeRate field if non-nil, zero value otherwise.

### GetFreeRateOk

`func (o *QuotaRequestDetail) GetFreeRateOk() (*int32, bool)`

GetFreeRateOk returns a tuple with the FreeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeRate

`func (o *QuotaRequestDetail) SetFreeRate(v int32)`

SetFreeRate sets FreeRate field to given value.

### HasFreeRate

`func (o *QuotaRequestDetail) HasFreeRate() bool`

HasFreeRate returns a boolean if a field has been set.

### SetFreeRateNil

`func (o *QuotaRequestDetail) SetFreeRateNil(b bool)`

 SetFreeRateNil sets the value for FreeRate to be an explicit nil

### UnsetFreeRate
`func (o *QuotaRequestDetail) UnsetFreeRate()`

UnsetFreeRate ensures that no value is present for FreeRate, not even an explicit nil
### GetId

`func (o *QuotaRequestDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotaRequestDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotaRequestDetail) SetId(v string)`

SetId sets Id field to given value.


### GetInitialValue

`func (o *QuotaRequestDetail) GetInitialValue() int32`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *QuotaRequestDetail) GetInitialValueOk() (*int32, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *QuotaRequestDetail) SetInitialValue(v int32)`

SetInitialValue sets InitialValue field to given value.


### GetModifiedAt

`func (o *QuotaRequestDetail) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *QuotaRequestDetail) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *QuotaRequestDetail) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *QuotaRequestDetail) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *QuotaRequestDetail) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *QuotaRequestDetail) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetQuotaItem

`func (o *QuotaRequestDetail) GetQuotaItem() string`

GetQuotaItem returns the QuotaItem field if non-nil, zero value otherwise.

### GetQuotaItemOk

`func (o *QuotaRequestDetail) GetQuotaItemOk() (*string, bool)`

GetQuotaItemOk returns a tuple with the QuotaItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaItem

`func (o *QuotaRequestDetail) SetQuotaItem(v string)`

SetQuotaItem sets QuotaItem field to given value.


### GetReduction

`func (o *QuotaRequestDetail) GetReduction() bool`

GetReduction returns the Reduction field if non-nil, zero value otherwise.

### GetReductionOk

`func (o *QuotaRequestDetail) GetReductionOk() (*bool, bool)`

GetReductionOk returns a tuple with the Reduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduction

`func (o *QuotaRequestDetail) SetReduction(v bool)`

SetReduction sets Reduction field to given value.


### GetRequestMessage

`func (o *QuotaRequestDetail) GetRequestMessage() string`

GetRequestMessage returns the RequestMessage field if non-nil, zero value otherwise.

### GetRequestMessageOk

`func (o *QuotaRequestDetail) GetRequestMessageOk() (*string, bool)`

GetRequestMessageOk returns a tuple with the RequestMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessage

`func (o *QuotaRequestDetail) SetRequestMessage(v string)`

SetRequestMessage sets RequestMessage field to given value.


### GetRequestValue

`func (o *QuotaRequestDetail) GetRequestValue() int32`

GetRequestValue returns the RequestValue field if non-nil, zero value otherwise.

### GetRequestValueOk

`func (o *QuotaRequestDetail) GetRequestValueOk() (*int32, bool)`

GetRequestValueOk returns a tuple with the RequestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestValue

`func (o *QuotaRequestDetail) SetRequestValue(v int32)`

SetRequestValue sets RequestValue field to given value.


### GetRequester

`func (o *QuotaRequestDetail) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *QuotaRequestDetail) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *QuotaRequestDetail) SetRequester(v string)`

SetRequester sets Requester field to given value.


### GetState

`func (o *QuotaRequestDetail) GetState() RequestStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuotaRequestDetail) GetStateOk() (*RequestStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuotaRequestDetail) SetState(v RequestStateEnum)`

SetState sets State field to given value.


### GetUnit

`func (o *QuotaRequestDetail) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *QuotaRequestDetail) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *QuotaRequestDetail) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


