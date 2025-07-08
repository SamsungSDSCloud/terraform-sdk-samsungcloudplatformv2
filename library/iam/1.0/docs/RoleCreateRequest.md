# RoleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** |  | [optional] 
**AssumeRolePolicyDocument** | [**PolicyDocument**](PolicyDocument.md) | 역할 신뢰 정책 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**MaxSessionDuration** | Pointer to **int32** | 역할 세션 최대 지속 시간(초) | [optional] [default to 3600]
**Name** | **string** | 역할명 | 
**PolicyIds** | **[]string** | 역할에 연결할 정책 ID 목록 | 
**Tags** | Pointer to **[]map[string]string** |  | [optional] 
**Types** | Pointer to **string** | 역할 유형 | [optional] [default to "USER_DEFINED"]

## Methods

### NewRoleCreateRequest

`func NewRoleCreateRequest(assumeRolePolicyDocument PolicyDocument, name string, policyIds []string, ) *RoleCreateRequest`

NewRoleCreateRequest instantiates a new RoleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleCreateRequestWithDefaults

`func NewRoleCreateRequestWithDefaults() *RoleCreateRequest`

NewRoleCreateRequestWithDefaults instantiates a new RoleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RoleCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RoleCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RoleCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *RoleCreateRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *RoleCreateRequest) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *RoleCreateRequest) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAssumeRolePolicyDocument

`func (o *RoleCreateRequest) GetAssumeRolePolicyDocument() PolicyDocument`

GetAssumeRolePolicyDocument returns the AssumeRolePolicyDocument field if non-nil, zero value otherwise.

### GetAssumeRolePolicyDocumentOk

`func (o *RoleCreateRequest) GetAssumeRolePolicyDocumentOk() (*PolicyDocument, bool)`

GetAssumeRolePolicyDocumentOk returns a tuple with the AssumeRolePolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumeRolePolicyDocument

`func (o *RoleCreateRequest) SetAssumeRolePolicyDocument(v PolicyDocument)`

SetAssumeRolePolicyDocument sets AssumeRolePolicyDocument field to given value.


### GetDescription

`func (o *RoleCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RoleCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RoleCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMaxSessionDuration

`func (o *RoleCreateRequest) GetMaxSessionDuration() int32`

GetMaxSessionDuration returns the MaxSessionDuration field if non-nil, zero value otherwise.

### GetMaxSessionDurationOk

`func (o *RoleCreateRequest) GetMaxSessionDurationOk() (*int32, bool)`

GetMaxSessionDurationOk returns a tuple with the MaxSessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionDuration

`func (o *RoleCreateRequest) SetMaxSessionDuration(v int32)`

SetMaxSessionDuration sets MaxSessionDuration field to given value.

### HasMaxSessionDuration

`func (o *RoleCreateRequest) HasMaxSessionDuration() bool`

HasMaxSessionDuration returns a boolean if a field has been set.

### GetName

`func (o *RoleCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPolicyIds

`func (o *RoleCreateRequest) GetPolicyIds() []string`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *RoleCreateRequest) GetPolicyIdsOk() (*[]string, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *RoleCreateRequest) SetPolicyIds(v []string)`

SetPolicyIds sets PolicyIds field to given value.


### GetTags

`func (o *RoleCreateRequest) GetTags() []map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoleCreateRequest) GetTagsOk() (*[]map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoleCreateRequest) SetTags(v []map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoleCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *RoleCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *RoleCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTypes

`func (o *RoleCreateRequest) GetTypes() string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *RoleCreateRequest) GetTypesOk() (*string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *RoleCreateRequest) SetTypes(v string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *RoleCreateRequest) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


