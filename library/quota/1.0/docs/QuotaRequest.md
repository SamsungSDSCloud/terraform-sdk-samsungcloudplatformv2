# QuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedAt** | Pointer to **NullableTime** |  | [optional] 
**AppliedValue** | **NullableInt32** |  | 
**ApprovalMessage** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**FreeRate** | Pointer to **NullableInt32** |  | [optional] 
**Id** | **string** | ID | 
**InitialValue** | **int32** | Initial value | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**QuotaItem** | **string** | Quota Item Name | 
**RequestClass** | **string** | Request Class | 
**RequestMessage** | **string** | Request message | 
**RequestValue** | **int32** | User Request Value | 
**Requester** | **string** | Requester | 
**Service** | **string** | Service Name | 
**State** | [**RequestStateEnum**](RequestStateEnum.md) | Request State | 
**Unit** | **string** | Unit in which the quota value is measured (e.g., EA, GB) | 

## Methods

### NewQuotaRequest

`func NewQuotaRequest(appliedValue NullableInt32, createdAt time.Time, createdBy string, id string, initialValue int32, modifiedAt time.Time, modifiedBy string, quotaItem string, requestClass string, requestMessage string, requestValue int32, requester string, service string, state RequestStateEnum, unit string, ) *QuotaRequest`

NewQuotaRequest instantiates a new QuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaRequestWithDefaults

`func NewQuotaRequestWithDefaults() *QuotaRequest`

NewQuotaRequestWithDefaults instantiates a new QuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedAt

`func (o *QuotaRequest) GetAppliedAt() time.Time`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *QuotaRequest) GetAppliedAtOk() (*time.Time, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *QuotaRequest) SetAppliedAt(v time.Time)`

SetAppliedAt sets AppliedAt field to given value.

### HasAppliedAt

`func (o *QuotaRequest) HasAppliedAt() bool`

HasAppliedAt returns a boolean if a field has been set.

### SetAppliedAtNil

`func (o *QuotaRequest) SetAppliedAtNil(b bool)`

 SetAppliedAtNil sets the value for AppliedAt to be an explicit nil

### UnsetAppliedAt
`func (o *QuotaRequest) UnsetAppliedAt()`

UnsetAppliedAt ensures that no value is present for AppliedAt, not even an explicit nil
### GetAppliedValue

`func (o *QuotaRequest) GetAppliedValue() int32`

GetAppliedValue returns the AppliedValue field if non-nil, zero value otherwise.

### GetAppliedValueOk

`func (o *QuotaRequest) GetAppliedValueOk() (*int32, bool)`

GetAppliedValueOk returns a tuple with the AppliedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedValue

`func (o *QuotaRequest) SetAppliedValue(v int32)`

SetAppliedValue sets AppliedValue field to given value.


### SetAppliedValueNil

`func (o *QuotaRequest) SetAppliedValueNil(b bool)`

 SetAppliedValueNil sets the value for AppliedValue to be an explicit nil

### UnsetAppliedValue
`func (o *QuotaRequest) UnsetAppliedValue()`

UnsetAppliedValue ensures that no value is present for AppliedValue, not even an explicit nil
### GetApprovalMessage

`func (o *QuotaRequest) GetApprovalMessage() string`

GetApprovalMessage returns the ApprovalMessage field if non-nil, zero value otherwise.

### GetApprovalMessageOk

`func (o *QuotaRequest) GetApprovalMessageOk() (*string, bool)`

GetApprovalMessageOk returns a tuple with the ApprovalMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalMessage

`func (o *QuotaRequest) SetApprovalMessage(v string)`

SetApprovalMessage sets ApprovalMessage field to given value.

### HasApprovalMessage

`func (o *QuotaRequest) HasApprovalMessage() bool`

HasApprovalMessage returns a boolean if a field has been set.

### SetApprovalMessageNil

`func (o *QuotaRequest) SetApprovalMessageNil(b bool)`

 SetApprovalMessageNil sets the value for ApprovalMessage to be an explicit nil

### UnsetApprovalMessage
`func (o *QuotaRequest) UnsetApprovalMessage()`

UnsetApprovalMessage ensures that no value is present for ApprovalMessage, not even an explicit nil
### GetCreatedAt

`func (o *QuotaRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuotaRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuotaRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *QuotaRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *QuotaRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *QuotaRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetFreeRate

`func (o *QuotaRequest) GetFreeRate() int32`

GetFreeRate returns the FreeRate field if non-nil, zero value otherwise.

### GetFreeRateOk

`func (o *QuotaRequest) GetFreeRateOk() (*int32, bool)`

GetFreeRateOk returns a tuple with the FreeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeRate

`func (o *QuotaRequest) SetFreeRate(v int32)`

SetFreeRate sets FreeRate field to given value.

### HasFreeRate

`func (o *QuotaRequest) HasFreeRate() bool`

HasFreeRate returns a boolean if a field has been set.

### SetFreeRateNil

`func (o *QuotaRequest) SetFreeRateNil(b bool)`

 SetFreeRateNil sets the value for FreeRate to be an explicit nil

### UnsetFreeRate
`func (o *QuotaRequest) UnsetFreeRate()`

UnsetFreeRate ensures that no value is present for FreeRate, not even an explicit nil
### GetId

`func (o *QuotaRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotaRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotaRequest) SetId(v string)`

SetId sets Id field to given value.


### GetInitialValue

`func (o *QuotaRequest) GetInitialValue() int32`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *QuotaRequest) GetInitialValueOk() (*int32, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *QuotaRequest) SetInitialValue(v int32)`

SetInitialValue sets InitialValue field to given value.


### GetModifiedAt

`func (o *QuotaRequest) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *QuotaRequest) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *QuotaRequest) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *QuotaRequest) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *QuotaRequest) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *QuotaRequest) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetQuotaItem

`func (o *QuotaRequest) GetQuotaItem() string`

GetQuotaItem returns the QuotaItem field if non-nil, zero value otherwise.

### GetQuotaItemOk

`func (o *QuotaRequest) GetQuotaItemOk() (*string, bool)`

GetQuotaItemOk returns a tuple with the QuotaItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaItem

`func (o *QuotaRequest) SetQuotaItem(v string)`

SetQuotaItem sets QuotaItem field to given value.


### GetRequestClass

`func (o *QuotaRequest) GetRequestClass() string`

GetRequestClass returns the RequestClass field if non-nil, zero value otherwise.

### GetRequestClassOk

`func (o *QuotaRequest) GetRequestClassOk() (*string, bool)`

GetRequestClassOk returns a tuple with the RequestClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestClass

`func (o *QuotaRequest) SetRequestClass(v string)`

SetRequestClass sets RequestClass field to given value.


### GetRequestMessage

`func (o *QuotaRequest) GetRequestMessage() string`

GetRequestMessage returns the RequestMessage field if non-nil, zero value otherwise.

### GetRequestMessageOk

`func (o *QuotaRequest) GetRequestMessageOk() (*string, bool)`

GetRequestMessageOk returns a tuple with the RequestMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessage

`func (o *QuotaRequest) SetRequestMessage(v string)`

SetRequestMessage sets RequestMessage field to given value.


### GetRequestValue

`func (o *QuotaRequest) GetRequestValue() int32`

GetRequestValue returns the RequestValue field if non-nil, zero value otherwise.

### GetRequestValueOk

`func (o *QuotaRequest) GetRequestValueOk() (*int32, bool)`

GetRequestValueOk returns a tuple with the RequestValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestValue

`func (o *QuotaRequest) SetRequestValue(v int32)`

SetRequestValue sets RequestValue field to given value.


### GetRequester

`func (o *QuotaRequest) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *QuotaRequest) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *QuotaRequest) SetRequester(v string)`

SetRequester sets Requester field to given value.


### GetService

`func (o *QuotaRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *QuotaRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *QuotaRequest) SetService(v string)`

SetService sets Service field to given value.


### GetState

`func (o *QuotaRequest) GetState() RequestStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QuotaRequest) GetStateOk() (*RequestStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QuotaRequest) SetState(v RequestStateEnum)`

SetState sets State field to given value.


### GetUnit

`func (o *QuotaRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *QuotaRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *QuotaRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


