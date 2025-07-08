# PolicyShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**CreatorEmail** | Pointer to **NullableString** |  | [optional] 
**CreatorName** | Pointer to **NullableString** |  | [optional] 
**DefaultVersionId** | Pointer to **string** | Policy Default Version Id | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DomainName** | **string** | 도메인 이름 | 
**Id** | Pointer to **string** | Policy ID | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**ModifierEmail** | Pointer to **NullableString** |  | [optional] 
**ModifierName** | Pointer to **NullableString** |  | [optional] 
**PolicyCategory** | Pointer to [**PolicyCategoryEnum**](PolicyCategoryEnum.md) | Policy Category | [optional] 
**PolicyName** | Pointer to **string** | Policy Name | [optional] 
**PolicyType** | Pointer to [**PolicyTypeEnum**](PolicyTypeEnum.md) | Policy Type | [optional] 
**PolicyVersions** | Pointer to [**[]PolicyVersion**](PolicyVersion.md) | Policy Versions | [optional] 
**ResourceType** | Pointer to **NullableString** |  | [optional] 
**ServiceName** | Pointer to **NullableString** |  | [optional] 
**ServiceType** | Pointer to **NullableString** |  | [optional] 
**Srn** | **string** | 삼성 리소스 자원명 | 
**State** | Pointer to [**PolicyStateEnum**](PolicyStateEnum.md) | Policy State | [optional] [default to POLICYSTATEENUM_ACTIVE]

## Methods

### NewPolicyShowResponse

`func NewPolicyShowResponse(accountId NullableString, createdAt time.Time, createdBy string, domainName string, modifiedAt time.Time, modifiedBy string, srn string, ) *PolicyShowResponse`

NewPolicyShowResponse instantiates a new PolicyShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyShowResponseWithDefaults

`func NewPolicyShowResponseWithDefaults() *PolicyShowResponse`

NewPolicyShowResponseWithDefaults instantiates a new PolicyShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PolicyShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PolicyShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PolicyShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *PolicyShowResponse) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *PolicyShowResponse) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetCreatedAt

`func (o *PolicyShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PolicyShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PolicyShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PolicyShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorEmail

`func (o *PolicyShowResponse) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *PolicyShowResponse) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *PolicyShowResponse) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *PolicyShowResponse) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### SetCreatorEmailNil

`func (o *PolicyShowResponse) SetCreatorEmailNil(b bool)`

 SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil

### UnsetCreatorEmail
`func (o *PolicyShowResponse) UnsetCreatorEmail()`

UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
### GetCreatorName

`func (o *PolicyShowResponse) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *PolicyShowResponse) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *PolicyShowResponse) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *PolicyShowResponse) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *PolicyShowResponse) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *PolicyShowResponse) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetDefaultVersionId

`func (o *PolicyShowResponse) GetDefaultVersionId() string`

GetDefaultVersionId returns the DefaultVersionId field if non-nil, zero value otherwise.

### GetDefaultVersionIdOk

`func (o *PolicyShowResponse) GetDefaultVersionIdOk() (*string, bool)`

GetDefaultVersionIdOk returns a tuple with the DefaultVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionId

`func (o *PolicyShowResponse) SetDefaultVersionId(v string)`

SetDefaultVersionId sets DefaultVersionId field to given value.

### HasDefaultVersionId

`func (o *PolicyShowResponse) HasDefaultVersionId() bool`

HasDefaultVersionId returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyShowResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyShowResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyShowResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyShowResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyShowResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyShowResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomainName

`func (o *PolicyShowResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *PolicyShowResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *PolicyShowResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetId

`func (o *PolicyShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyShowResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyShowResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PolicyShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PolicyShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PolicyShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *PolicyShowResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PolicyShowResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PolicyShowResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifierEmail

`func (o *PolicyShowResponse) GetModifierEmail() string`

GetModifierEmail returns the ModifierEmail field if non-nil, zero value otherwise.

### GetModifierEmailOk

`func (o *PolicyShowResponse) GetModifierEmailOk() (*string, bool)`

GetModifierEmailOk returns a tuple with the ModifierEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierEmail

`func (o *PolicyShowResponse) SetModifierEmail(v string)`

SetModifierEmail sets ModifierEmail field to given value.

### HasModifierEmail

`func (o *PolicyShowResponse) HasModifierEmail() bool`

HasModifierEmail returns a boolean if a field has been set.

### SetModifierEmailNil

`func (o *PolicyShowResponse) SetModifierEmailNil(b bool)`

 SetModifierEmailNil sets the value for ModifierEmail to be an explicit nil

### UnsetModifierEmail
`func (o *PolicyShowResponse) UnsetModifierEmail()`

UnsetModifierEmail ensures that no value is present for ModifierEmail, not even an explicit nil
### GetModifierName

`func (o *PolicyShowResponse) GetModifierName() string`

GetModifierName returns the ModifierName field if non-nil, zero value otherwise.

### GetModifierNameOk

`func (o *PolicyShowResponse) GetModifierNameOk() (*string, bool)`

GetModifierNameOk returns a tuple with the ModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierName

`func (o *PolicyShowResponse) SetModifierName(v string)`

SetModifierName sets ModifierName field to given value.

### HasModifierName

`func (o *PolicyShowResponse) HasModifierName() bool`

HasModifierName returns a boolean if a field has been set.

### SetModifierNameNil

`func (o *PolicyShowResponse) SetModifierNameNil(b bool)`

 SetModifierNameNil sets the value for ModifierName to be an explicit nil

### UnsetModifierName
`func (o *PolicyShowResponse) UnsetModifierName()`

UnsetModifierName ensures that no value is present for ModifierName, not even an explicit nil
### GetPolicyCategory

`func (o *PolicyShowResponse) GetPolicyCategory() PolicyCategoryEnum`

GetPolicyCategory returns the PolicyCategory field if non-nil, zero value otherwise.

### GetPolicyCategoryOk

`func (o *PolicyShowResponse) GetPolicyCategoryOk() (*PolicyCategoryEnum, bool)`

GetPolicyCategoryOk returns a tuple with the PolicyCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCategory

`func (o *PolicyShowResponse) SetPolicyCategory(v PolicyCategoryEnum)`

SetPolicyCategory sets PolicyCategory field to given value.

### HasPolicyCategory

`func (o *PolicyShowResponse) HasPolicyCategory() bool`

HasPolicyCategory returns a boolean if a field has been set.

### GetPolicyName

`func (o *PolicyShowResponse) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyShowResponse) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyShowResponse) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyShowResponse) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyType

`func (o *PolicyShowResponse) GetPolicyType() PolicyTypeEnum`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *PolicyShowResponse) GetPolicyTypeOk() (*PolicyTypeEnum, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *PolicyShowResponse) SetPolicyType(v PolicyTypeEnum)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *PolicyShowResponse) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetPolicyVersions

`func (o *PolicyShowResponse) GetPolicyVersions() []PolicyVersion`

GetPolicyVersions returns the PolicyVersions field if non-nil, zero value otherwise.

### GetPolicyVersionsOk

`func (o *PolicyShowResponse) GetPolicyVersionsOk() (*[]PolicyVersion, bool)`

GetPolicyVersionsOk returns a tuple with the PolicyVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersions

`func (o *PolicyShowResponse) SetPolicyVersions(v []PolicyVersion)`

SetPolicyVersions sets PolicyVersions field to given value.

### HasPolicyVersions

`func (o *PolicyShowResponse) HasPolicyVersions() bool`

HasPolicyVersions returns a boolean if a field has been set.

### GetResourceType

`func (o *PolicyShowResponse) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PolicyShowResponse) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PolicyShowResponse) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *PolicyShowResponse) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *PolicyShowResponse) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *PolicyShowResponse) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetServiceName

`func (o *PolicyShowResponse) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PolicyShowResponse) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PolicyShowResponse) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *PolicyShowResponse) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *PolicyShowResponse) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *PolicyShowResponse) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetServiceType

`func (o *PolicyShowResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PolicyShowResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PolicyShowResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *PolicyShowResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### SetServiceTypeNil

`func (o *PolicyShowResponse) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *PolicyShowResponse) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetSrn

`func (o *PolicyShowResponse) GetSrn() string`

GetSrn returns the Srn field if non-nil, zero value otherwise.

### GetSrnOk

`func (o *PolicyShowResponse) GetSrnOk() (*string, bool)`

GetSrnOk returns a tuple with the Srn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrn

`func (o *PolicyShowResponse) SetSrn(v string)`

SetSrn sets Srn field to given value.


### GetState

`func (o *PolicyShowResponse) GetState() PolicyStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PolicyShowResponse) GetStateOk() (*PolicyStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PolicyShowResponse) SetState(v PolicyStateEnum)`

SetState sets State field to given value.

### HasState

`func (o *PolicyShowResponse) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


