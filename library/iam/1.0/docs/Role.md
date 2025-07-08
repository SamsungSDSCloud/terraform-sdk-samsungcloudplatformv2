# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** |  | [optional] 
**AssumeRolePolicyDocument** | [**PolicyDocument**](PolicyDocument.md) | 역할 신뢰 정책 | 
**CreatedAt** | **time.Time** | 생성 일시 | 
**CreatedBy** | **string** | 생성자 | 
**CreatorEmail** | Pointer to **string** | 생성자 Email | [optional] [default to "-"]
**CreatorName** | Pointer to **string** | 생성자 성, 이름 | [optional] [default to "-"]
**Description** | **NullableString** |  | 
**Id** | **string** | ID | 
**MaxSessionDuration** | **int32** | 역할 세션 최대 지속 시간(초) | 
**ModifiedAt** | **time.Time** | 수정 일시 | 
**ModifiedBy** | **string** | 수정자 | 
**ModifierEmail** | Pointer to **string** | 수정자 Email | [optional] [default to "-"]
**ModifierName** | Pointer to **string** | 수정자 성, 이름 | [optional] [default to "-"]
**Name** | **string** | 역할명 | 
**Policies** | [**[]Policy**](Policy.md) |  | 
**Type** | Pointer to [**RoleTypeEnum**](RoleTypeEnum.md) | 역할 유형 | [optional] [default to ROLETYPEENUM_DEFAULT]

## Methods

### NewRole

`func NewRole(assumeRolePolicyDocument PolicyDocument, createdAt time.Time, createdBy string, description NullableString, id string, maxSessionDuration int32, modifiedAt time.Time, modifiedBy string, name string, policies []Policy, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Role) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Role) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Role) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Role) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *Role) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Role) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAssumeRolePolicyDocument

`func (o *Role) GetAssumeRolePolicyDocument() PolicyDocument`

GetAssumeRolePolicyDocument returns the AssumeRolePolicyDocument field if non-nil, zero value otherwise.

### GetAssumeRolePolicyDocumentOk

`func (o *Role) GetAssumeRolePolicyDocumentOk() (*PolicyDocument, bool)`

GetAssumeRolePolicyDocumentOk returns a tuple with the AssumeRolePolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumeRolePolicyDocument

`func (o *Role) SetAssumeRolePolicyDocument(v PolicyDocument)`

SetAssumeRolePolicyDocument sets AssumeRolePolicyDocument field to given value.


### GetCreatedAt

`func (o *Role) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Role) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Role) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Role) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Role) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Role) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorEmail

`func (o *Role) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *Role) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *Role) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *Role) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetCreatorName

`func (o *Role) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *Role) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *Role) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *Role) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Role) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Role) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Role) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v string)`

SetId sets Id field to given value.


### GetMaxSessionDuration

`func (o *Role) GetMaxSessionDuration() int32`

GetMaxSessionDuration returns the MaxSessionDuration field if non-nil, zero value otherwise.

### GetMaxSessionDurationOk

`func (o *Role) GetMaxSessionDurationOk() (*int32, bool)`

GetMaxSessionDurationOk returns a tuple with the MaxSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionDuration

`func (o *Role) SetMaxSessionDuration(v int32)`

SetMaxSessionDuration sets MaxSessionDuration field to given value.


### GetModifiedAt

`func (o *Role) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Role) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Role) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Role) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Role) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Role) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifierEmail

`func (o *Role) GetModifierEmail() string`

GetModifierEmail returns the ModifierEmail field if non-nil, zero value otherwise.

### GetModifierEmailOk

`func (o *Role) GetModifierEmailOk() (*string, bool)`

GetModifierEmailOk returns a tuple with the ModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierEmail

`func (o *Role) SetModifierEmail(v string)`

SetModifierEmail sets ModifierEmail field to given value.

### HasModifierEmail

`func (o *Role) HasModifierEmail() bool`

HasModifierEmail returns a boolean if a field has been set.

### GetModifierName

`func (o *Role) GetModifierName() string`

GetModifierName returns the ModifierName field if non-nil, zero value otherwise.

### GetModifierNameOk

`func (o *Role) GetModifierNameOk() (*string, bool)`

GetModifierNameOk returns a tuple with the ModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierName

`func (o *Role) SetModifierName(v string)`

SetModifierName sets ModifierName field to given value.

### HasModifierName

`func (o *Role) HasModifierName() bool`

HasModifierName returns a boolean if a field has been set.

### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.


### GetPolicies

`func (o *Role) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *Role) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *Role) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.


### SetPoliciesNil

`func (o *Role) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *Role) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil
### GetType

`func (o *Role) GetType() RoleTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Role) GetTypeOk() (*RoleTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Role) SetType(v RoleTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *Role) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


