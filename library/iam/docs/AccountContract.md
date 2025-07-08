# AccountContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**BillingCorporationCode** | **string** | 과금 법인 Code | 
**BillingCorporationName** | **NullableString** |  | 
**BoCode** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | 생성 일시 | 
**CreatedBy** | **string** | 생성자 | 
**CreatorEmail** | **string** | 생성자 Email | 
**CreatorName** | **string** | 생성자 성, 이름 | 
**CurrencyCode** | [**CurrencyCodeEnum**](CurrencyCodeEnum.md) | 화폐 Code | 
**CustomerCode** | **string** | 고객사 Code | 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | 수정 일시 | 
**ModifiedBy** | **string** | 수정자 | 
**ModifierEmail** | **string** | 수정자 Email | 
**ModifierName** | **string** | 수정자 성, 이름 | 
**PaymentAccountId** | **string** | 납부자 Account ID | 
**ProjectCode** | Pointer to **NullableString** |  | [optional] 
**Wbses** | [**[]AccountContractWbs**](AccountContractWbs.md) | Account 계약 WBS 정보 | 

## Methods

### NewAccountContract

`func NewAccountContract(accountId string, billingCorporationCode string, billingCorporationName NullableString, createdAt time.Time, createdBy string, creatorEmail string, creatorName string, currencyCode CurrencyCodeEnum, customerCode string, id string, modifiedAt time.Time, modifiedBy string, modifierEmail string, modifierName string, paymentAccountId string, wbses []AccountContractWbs, ) *AccountContract`

NewAccountContract instantiates a new AccountContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountContractWithDefaults

`func NewAccountContractWithDefaults() *AccountContract`

NewAccountContractWithDefaults instantiates a new AccountContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountContract) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountContract) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountContract) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBillingCorporationCode

`func (o *AccountContract) GetBillingCorporationCode() string`

GetBillingCorporationCode returns the BillingCorporationCode field if non-nil, zero value otherwise.

### GetBillingCorporationCodeOk

`func (o *AccountContract) GetBillingCorporationCodeOk() (*string, bool)`

GetBillingCorporationCodeOk returns a tuple with the BillingCorporationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCorporationCode

`func (o *AccountContract) SetBillingCorporationCode(v string)`

SetBillingCorporationCode sets BillingCorporationCode field to given value.


### GetBillingCorporationName

`func (o *AccountContract) GetBillingCorporationName() string`

GetBillingCorporationName returns the BillingCorporationName field if non-nil, zero value otherwise.

### GetBillingCorporationNameOk

`func (o *AccountContract) GetBillingCorporationNameOk() (*string, bool)`

GetBillingCorporationNameOk returns a tuple with the BillingCorporationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCorporationName

`func (o *AccountContract) SetBillingCorporationName(v string)`

SetBillingCorporationName sets BillingCorporationName field to given value.


### SetBillingCorporationNameNil

`func (o *AccountContract) SetBillingCorporationNameNil(b bool)`

 SetBillingCorporationNameNil sets the value for BillingCorporationName to be an explicit nil

### UnsetBillingCorporationName
`func (o *AccountContract) UnsetBillingCorporationName()`

UnsetBillingCorporationName ensures that no value is present for BillingCorporationName, not even an explicit nil
### GetBoCode

`func (o *AccountContract) GetBoCode() string`

GetBoCode returns the BoCode field if non-nil, zero value otherwise.

### GetBoCodeOk

`func (o *AccountContract) GetBoCodeOk() (*string, bool)`

GetBoCodeOk returns a tuple with the BoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoCode

`func (o *AccountContract) SetBoCode(v string)`

SetBoCode sets BoCode field to given value.

### HasBoCode

`func (o *AccountContract) HasBoCode() bool`

HasBoCode returns a boolean if a field has been set.

### SetBoCodeNil

`func (o *AccountContract) SetBoCodeNil(b bool)`

 SetBoCodeNil sets the value for BoCode to be an explicit nil

### UnsetBoCode
`func (o *AccountContract) UnsetBoCode()`

UnsetBoCode ensures that no value is present for BoCode, not even an explicit nil
### GetCreatedAt

`func (o *AccountContract) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountContract) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountContract) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AccountContract) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccountContract) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccountContract) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorEmail

`func (o *AccountContract) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *AccountContract) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *AccountContract) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.


### GetCreatorName

`func (o *AccountContract) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *AccountContract) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *AccountContract) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.


### GetCurrencyCode

`func (o *AccountContract) GetCurrencyCode() CurrencyCodeEnum`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AccountContract) GetCurrencyCodeOk() (*CurrencyCodeEnum, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AccountContract) SetCurrencyCode(v CurrencyCodeEnum)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetCustomerCode

`func (o *AccountContract) GetCustomerCode() string`

GetCustomerCode returns the CustomerCode field if non-nil, zero value otherwise.

### GetCustomerCodeOk

`func (o *AccountContract) GetCustomerCodeOk() (*string, bool)`

GetCustomerCodeOk returns a tuple with the CustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCode

`func (o *AccountContract) SetCustomerCode(v string)`

SetCustomerCode sets CustomerCode field to given value.


### GetId

`func (o *AccountContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountContract) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *AccountContract) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AccountContract) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AccountContract) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *AccountContract) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *AccountContract) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *AccountContract) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifierEmail

`func (o *AccountContract) GetModifierEmail() string`

GetModifierEmail returns the ModifierEmail field if non-nil, zero value otherwise.

### GetModifierEmailOk

`func (o *AccountContract) GetModifierEmailOk() (*string, bool)`

GetModifierEmailOk returns a tuple with the ModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierEmail

`func (o *AccountContract) SetModifierEmail(v string)`

SetModifierEmail sets ModifierEmail field to given value.


### GetModifierName

`func (o *AccountContract) GetModifierName() string`

GetModifierName returns the ModifierName field if non-nil, zero value otherwise.

### GetModifierNameOk

`func (o *AccountContract) GetModifierNameOk() (*string, bool)`

GetModifierNameOk returns a tuple with the ModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierName

`func (o *AccountContract) SetModifierName(v string)`

SetModifierName sets ModifierName field to given value.


### GetPaymentAccountId

`func (o *AccountContract) GetPaymentAccountId() string`

GetPaymentAccountId returns the PaymentAccountId field if non-nil, zero value otherwise.

### GetPaymentAccountIdOk

`func (o *AccountContract) GetPaymentAccountIdOk() (*string, bool)`

GetPaymentAccountIdOk returns a tuple with the PaymentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAccountId

`func (o *AccountContract) SetPaymentAccountId(v string)`

SetPaymentAccountId sets PaymentAccountId field to given value.


### GetProjectCode

`func (o *AccountContract) GetProjectCode() string`

GetProjectCode returns the ProjectCode field if non-nil, zero value otherwise.

### GetProjectCodeOk

`func (o *AccountContract) GetProjectCodeOk() (*string, bool)`

GetProjectCodeOk returns a tuple with the ProjectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCode

`func (o *AccountContract) SetProjectCode(v string)`

SetProjectCode sets ProjectCode field to given value.

### HasProjectCode

`func (o *AccountContract) HasProjectCode() bool`

HasProjectCode returns a boolean if a field has been set.

### SetProjectCodeNil

`func (o *AccountContract) SetProjectCodeNil(b bool)`

 SetProjectCodeNil sets the value for ProjectCode to be an explicit nil

### UnsetProjectCode
`func (o *AccountContract) UnsetProjectCode()`

UnsetProjectCode ensures that no value is present for ProjectCode, not even an explicit nil
### GetWbses

`func (o *AccountContract) GetWbses() []AccountContractWbs`

GetWbses returns the Wbses field if non-nil, zero value otherwise.

### GetWbsesOk

`func (o *AccountContract) GetWbsesOk() (*[]AccountContractWbs, bool)`

GetWbsesOk returns a tuple with the Wbses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbses

`func (o *AccountContract) SetWbses(v []AccountContractWbs)`

SetWbses sets Wbses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


