# PaymentContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessCategoryId** | **string** | 회사 업무 구분 ID | 
**BusinessCategoryName** | **NullableString** |  | 
**BusinessGroupCode** | **NullableString** |  | 
**CompanyId** | **string** | Company ID | 
**CompanyName** | **string** | Company 이름 | 
**Contract** | [**AccountContract**](AccountContract.md) | Account 계약 정보 | 
**CreatedAt** | **time.Time** | 생성 일시 | 
**CreatedBy** | **string** | 생성자 | 
**CreatorEmail** | **string** | 생성자 Email | 
**CreatorName** | **string** | 생성자 성, 이름 | 
**Description** | **NullableString** |  | 
**HasProjectYn** | [**YnEnum**](YnEnum.md) | Project 생성 여부 | 
**Id** | **string** | ID | 
**ManagedServiceType** | [**ManagedServiceTypeEnum**](ManagedServiceTypeEnum.md) | Service Type | 
**ModifiedAt** | **time.Time** | 수정 일시 | 
**ModifiedBy** | **string** | 수정자 | 
**ModifierEmail** | **string** | 수정자 Email | 
**ModifierName** | **string** | 수정자 성, 이름 | 
**Name** | **string** | Account 이름 | 
**PartnerCompanyCode** | **NullableString** |  | 

## Methods

### NewPaymentContract

`func NewPaymentContract(businessCategoryId string, businessCategoryName NullableString, businessGroupCode NullableString, companyId string, companyName string, contract AccountContract, createdAt time.Time, createdBy string, creatorEmail string, creatorName string, description NullableString, hasProjectYn YnEnum, id string, managedServiceType ManagedServiceTypeEnum, modifiedAt time.Time, modifiedBy string, modifierEmail string, modifierName string, name string, partnerCompanyCode NullableString, ) *PaymentContract`

NewPaymentContract instantiates a new PaymentContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentContractWithDefaults

`func NewPaymentContractWithDefaults() *PaymentContract`

NewPaymentContractWithDefaults instantiates a new PaymentContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessCategoryId

`func (o *PaymentContract) GetBusinessCategoryId() string`

GetBusinessCategoryId returns the BusinessCategoryId field if non-nil, zero value otherwise.

### GetBusinessCategoryIdOk

`func (o *PaymentContract) GetBusinessCategoryIdOk() (*string, bool)`

GetBusinessCategoryIdOk returns a tuple with the BusinessCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCategoryId

`func (o *PaymentContract) SetBusinessCategoryId(v string)`

SetBusinessCategoryId sets BusinessCategoryId field to given value.


### GetBusinessCategoryName

`func (o *PaymentContract) GetBusinessCategoryName() string`

GetBusinessCategoryName returns the BusinessCategoryName field if non-nil, zero value otherwise.

### GetBusinessCategoryNameOk

`func (o *PaymentContract) GetBusinessCategoryNameOk() (*string, bool)`

GetBusinessCategoryNameOk returns a tuple with the BusinessCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCategoryName

`func (o *PaymentContract) SetBusinessCategoryName(v string)`

SetBusinessCategoryName sets BusinessCategoryName field to given value.


### SetBusinessCategoryNameNil

`func (o *PaymentContract) SetBusinessCategoryNameNil(b bool)`

 SetBusinessCategoryNameNil sets the value for BusinessCategoryName to be an explicit nil

### UnsetBusinessCategoryName
`func (o *PaymentContract) UnsetBusinessCategoryName()`

UnsetBusinessCategoryName ensures that no value is present for BusinessCategoryName, not even an explicit nil
### GetBusinessGroupCode

`func (o *PaymentContract) GetBusinessGroupCode() string`

GetBusinessGroupCode returns the BusinessGroupCode field if non-nil, zero value otherwise.

### GetBusinessGroupCodeOk

`func (o *PaymentContract) GetBusinessGroupCodeOk() (*string, bool)`

GetBusinessGroupCodeOk returns a tuple with the BusinessGroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessGroupCode

`func (o *PaymentContract) SetBusinessGroupCode(v string)`

SetBusinessGroupCode sets BusinessGroupCode field to given value.


### SetBusinessGroupCodeNil

`func (o *PaymentContract) SetBusinessGroupCodeNil(b bool)`

 SetBusinessGroupCodeNil sets the value for BusinessGroupCode to be an explicit nil

### UnsetBusinessGroupCode
`func (o *PaymentContract) UnsetBusinessGroupCode()`

UnsetBusinessGroupCode ensures that no value is present for BusinessGroupCode, not even an explicit nil
### GetCompanyId

`func (o *PaymentContract) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PaymentContract) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PaymentContract) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetCompanyName

`func (o *PaymentContract) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *PaymentContract) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *PaymentContract) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetContract

`func (o *PaymentContract) GetContract() AccountContract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *PaymentContract) GetContractOk() (*AccountContract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *PaymentContract) SetContract(v AccountContract)`

SetContract sets Contract field to given value.


### GetCreatedAt

`func (o *PaymentContract) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentContract) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentContract) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PaymentContract) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PaymentContract) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PaymentContract) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorEmail

`func (o *PaymentContract) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *PaymentContract) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *PaymentContract) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.


### GetCreatorName

`func (o *PaymentContract) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *PaymentContract) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *PaymentContract) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.


### GetDescription

`func (o *PaymentContract) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentContract) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentContract) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PaymentContract) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentContract) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHasProjectYn

`func (o *PaymentContract) GetHasProjectYn() YnEnum`

GetHasProjectYn returns the HasProjectYn field if non-nil, zero value otherwise.

### GetHasProjectYnOk

`func (o *PaymentContract) GetHasProjectYnOk() (*YnEnum, bool)`

GetHasProjectYnOk returns a tuple with the HasProjectYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjectYn

`func (o *PaymentContract) SetHasProjectYn(v YnEnum)`

SetHasProjectYn sets HasProjectYn field to given value.


### GetId

`func (o *PaymentContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentContract) SetId(v string)`

SetId sets Id field to given value.


### GetManagedServiceType

`func (o *PaymentContract) GetManagedServiceType() ManagedServiceTypeEnum`

GetManagedServiceType returns the ManagedServiceType field if non-nil, zero value otherwise.

### GetManagedServiceTypeOk

`func (o *PaymentContract) GetManagedServiceTypeOk() (*ManagedServiceTypeEnum, bool)`

GetManagedServiceTypeOk returns a tuple with the ManagedServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedServiceType

`func (o *PaymentContract) SetManagedServiceType(v ManagedServiceTypeEnum)`

SetManagedServiceType sets ManagedServiceType field to given value.


### GetModifiedAt

`func (o *PaymentContract) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PaymentContract) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PaymentContract) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *PaymentContract) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PaymentContract) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PaymentContract) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifierEmail

`func (o *PaymentContract) GetModifierEmail() string`

GetModifierEmail returns the ModifierEmail field if non-nil, zero value otherwise.

### GetModifierEmailOk

`func (o *PaymentContract) GetModifierEmailOk() (*string, bool)`

GetModifierEmailOk returns a tuple with the ModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierEmail

`func (o *PaymentContract) SetModifierEmail(v string)`

SetModifierEmail sets ModifierEmail field to given value.


### GetModifierName

`func (o *PaymentContract) GetModifierName() string`

GetModifierName returns the ModifierName field if non-nil, zero value otherwise.

### GetModifierNameOk

`func (o *PaymentContract) GetModifierNameOk() (*string, bool)`

GetModifierNameOk returns a tuple with the ModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierName

`func (o *PaymentContract) SetModifierName(v string)`

SetModifierName sets ModifierName field to given value.


### GetName

`func (o *PaymentContract) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentContract) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentContract) SetName(v string)`

SetName sets Name field to given value.


### GetPartnerCompanyCode

`func (o *PaymentContract) GetPartnerCompanyCode() string`

GetPartnerCompanyCode returns the PartnerCompanyCode field if non-nil, zero value otherwise.

### GetPartnerCompanyCodeOk

`func (o *PaymentContract) GetPartnerCompanyCodeOk() (*string, bool)`

GetPartnerCompanyCodeOk returns a tuple with the PartnerCompanyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCompanyCode

`func (o *PaymentContract) SetPartnerCompanyCode(v string)`

SetPartnerCompanyCode sets PartnerCompanyCode field to given value.


### SetPartnerCompanyCodeNil

`func (o *PaymentContract) SetPartnerCompanyCodeNil(b bool)`

 SetPartnerCompanyCodeNil sets the value for PartnerCompanyCode to be an explicit nil

### UnsetPartnerCompanyCode
`func (o *PaymentContract) UnsetPartnerCompanyCode()`

UnsetPartnerCompanyCode ensures that no value is present for PartnerCompanyCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


