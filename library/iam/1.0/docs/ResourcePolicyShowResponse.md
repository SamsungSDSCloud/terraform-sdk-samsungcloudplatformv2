# ResourcePolicyShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**CreatorEmail** | Pointer to **NullableString** |  | [optional] 
**CreatorName** | Pointer to **NullableString** |  | [optional] 
**DefaultVersionId** | Pointer to **string** | Policy Default Version Id | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**ModifierEmail** | Pointer to **NullableString** |  | [optional] 
**ModifierName** | Pointer to **NullableString** |  | [optional] 
**PolicyCategory** | Pointer to [**PolicyCategoryEnum**](PolicyCategoryEnum.md) | Policy Category | [optional] 
**PolicyType** | Pointer to [**PolicyTypeEnum**](PolicyTypeEnum.md) | Policy Type | [optional] 
**PolicyVersions** | Pointer to [**[]ResourcePolicyVersion**](ResourcePolicyVersion.md) | Policy Versions | [optional] 
**ServiceName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResourcePolicyShowResponse

`func NewResourcePolicyShowResponse(accountId NullableString, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, ) *ResourcePolicyShowResponse`

NewResourcePolicyShowResponse instantiates a new ResourcePolicyShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePolicyShowResponseWithDefaults

`func NewResourcePolicyShowResponseWithDefaults() *ResourcePolicyShowResponse`

NewResourcePolicyShowResponseWithDefaults instantiates a new ResourcePolicyShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ResourcePolicyShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ResourcePolicyShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ResourcePolicyShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *ResourcePolicyShowResponse) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *ResourcePolicyShowResponse) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetCreatedAt

`func (o *ResourcePolicyShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourcePolicyShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourcePolicyShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ResourcePolicyShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ResourcePolicyShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ResourcePolicyShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorEmail

`func (o *ResourcePolicyShowResponse) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *ResourcePolicyShowResponse) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *ResourcePolicyShowResponse) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *ResourcePolicyShowResponse) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### SetCreatorEmailNil

`func (o *ResourcePolicyShowResponse) SetCreatorEmailNil(b bool)`

 SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil

### UnsetCreatorEmail
`func (o *ResourcePolicyShowResponse) UnsetCreatorEmail()`

UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
### GetCreatorName

`func (o *ResourcePolicyShowResponse) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ResourcePolicyShowResponse) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ResourcePolicyShowResponse) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *ResourcePolicyShowResponse) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *ResourcePolicyShowResponse) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *ResourcePolicyShowResponse) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetDefaultVersionId

`func (o *ResourcePolicyShowResponse) GetDefaultVersionId() string`

GetDefaultVersionId returns the DefaultVersionId field if non-nil, zero value otherwise.

### GetDefaultVersionIdOk

`func (o *ResourcePolicyShowResponse) GetDefaultVersionIdOk() (*string, bool)`

GetDefaultVersionIdOk returns a tuple with the DefaultVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionId

`func (o *ResourcePolicyShowResponse) SetDefaultVersionId(v string)`

SetDefaultVersionId sets DefaultVersionId field to given value.

### HasDefaultVersionId

`func (o *ResourcePolicyShowResponse) HasDefaultVersionId() bool`

HasDefaultVersionId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ResourcePolicyShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ResourcePolicyShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ResourcePolicyShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *ResourcePolicyShowResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ResourcePolicyShowResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ResourcePolicyShowResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifierEmail

`func (o *ResourcePolicyShowResponse) GetModifierEmail() string`

GetModifierEmail returns the ModifierEmail field if non-nil, zero value otherwise.

### GetModifierEmailOk

`func (o *ResourcePolicyShowResponse) GetModifierEmailOk() (*string, bool)`

GetModifierEmailOk returns a tuple with the ModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierEmail

`func (o *ResourcePolicyShowResponse) SetModifierEmail(v string)`

SetModifierEmail sets ModifierEmail field to given value.

### HasModifierEmail

`func (o *ResourcePolicyShowResponse) HasModifierEmail() bool`

HasModifierEmail returns a boolean if a field has been set.

### SetModifierEmailNil

`func (o *ResourcePolicyShowResponse) SetModifierEmailNil(b bool)`

 SetModifierEmailNil sets the value for ModifierEmail to be an explicit nil

### UnsetModifierEmail
`func (o *ResourcePolicyShowResponse) UnsetModifierEmail()`

UnsetModifierEmail ensures that no value is present for ModifierEmail, not even an explicit nil
### GetModifierName

`func (o *ResourcePolicyShowResponse) GetModifierName() string`

GetModifierName returns the ModifierName field if non-nil, zero value otherwise.

### GetModifierNameOk

`func (o *ResourcePolicyShowResponse) GetModifierNameOk() (*string, bool)`

GetModifierNameOk returns a tuple with the ModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierName

`func (o *ResourcePolicyShowResponse) SetModifierName(v string)`

SetModifierName sets ModifierName field to given value.

### HasModifierName

`func (o *ResourcePolicyShowResponse) HasModifierName() bool`

HasModifierName returns a boolean if a field has been set.

### SetModifierNameNil

`func (o *ResourcePolicyShowResponse) SetModifierNameNil(b bool)`

 SetModifierNameNil sets the value for ModifierName to be an explicit nil

### UnsetModifierName
`func (o *ResourcePolicyShowResponse) UnsetModifierName()`

UnsetModifierName ensures that no value is present for ModifierName, not even an explicit nil
### GetPolicyCategory

`func (o *ResourcePolicyShowResponse) GetPolicyCategory() PolicyCategoryEnum`

GetPolicyCategory returns the PolicyCategory field if non-nil, zero value otherwise.

### GetPolicyCategoryOk

`func (o *ResourcePolicyShowResponse) GetPolicyCategoryOk() (*PolicyCategoryEnum, bool)`

GetPolicyCategoryOk returns a tuple with the PolicyCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCategory

`func (o *ResourcePolicyShowResponse) SetPolicyCategory(v PolicyCategoryEnum)`

SetPolicyCategory sets PolicyCategory field to given value.

### HasPolicyCategory

`func (o *ResourcePolicyShowResponse) HasPolicyCategory() bool`

HasPolicyCategory returns a boolean if a field has been set.

### GetPolicyType

`func (o *ResourcePolicyShowResponse) GetPolicyType() PolicyTypeEnum`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *ResourcePolicyShowResponse) GetPolicyTypeOk() (*PolicyTypeEnum, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *ResourcePolicyShowResponse) SetPolicyType(v PolicyTypeEnum)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *ResourcePolicyShowResponse) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetPolicyVersions

`func (o *ResourcePolicyShowResponse) GetPolicyVersions() []ResourcePolicyVersion`

GetPolicyVersions returns the PolicyVersions field if non-nil, zero value otherwise.

### GetPolicyVersionsOk

`func (o *ResourcePolicyShowResponse) GetPolicyVersionsOk() (*[]ResourcePolicyVersion, bool)`

GetPolicyVersionsOk returns a tuple with the PolicyVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersions

`func (o *ResourcePolicyShowResponse) SetPolicyVersions(v []ResourcePolicyVersion)`

SetPolicyVersions sets PolicyVersions field to given value.

### HasPolicyVersions

`func (o *ResourcePolicyShowResponse) HasPolicyVersions() bool`

HasPolicyVersions returns a boolean if a field has been set.

### GetServiceName

`func (o *ResourcePolicyShowResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourcePolicyShowResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourcePolicyShowResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ResourcePolicyShowResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *ResourcePolicyShowResponse) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *ResourcePolicyShowResponse) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


