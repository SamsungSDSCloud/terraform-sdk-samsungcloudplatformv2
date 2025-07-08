# PolicyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVersionId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyCategory** | Pointer to [**PolicyCategoryEnum**](PolicyCategoryEnum.md) | Policy Category | [optional] [default to POLICYCATEGORYENUM_IDENTITY_BASED]
**PolicyName** | **string** | Policy Name | 
**PolicyType** | Pointer to [**PolicyTypeEnum**](PolicyTypeEnum.md) | Policy Type | [optional] [default to POLICYTYPEENUM_USER_DEFINED]
**PolicyVersion** | [**PolicyVersionCreateRequest**](PolicyVersionCreateRequest.md) | Policy Version | 
**ServiceType** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]map[string]string** |  | [optional] 

## Methods

### NewPolicyCreateRequest

`func NewPolicyCreateRequest(policyName string, policyVersion PolicyVersionCreateRequest, ) *PolicyCreateRequest`

NewPolicyCreateRequest instantiates a new PolicyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCreateRequestWithDefaults

`func NewPolicyCreateRequestWithDefaults() *PolicyCreateRequest`

NewPolicyCreateRequestWithDefaults instantiates a new PolicyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVersionId

`func (o *PolicyCreateRequest) GetDefaultVersionId() string`

GetDefaultVersionId returns the DefaultVersionId field if non-nil, zero value otherwise.

### GetDefaultVersionIdOk

`func (o *PolicyCreateRequest) GetDefaultVersionIdOk() (*string, bool)`

GetDefaultVersionIdOk returns a tuple with the DefaultVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionId

`func (o *PolicyCreateRequest) SetDefaultVersionId(v string)`

SetDefaultVersionId sets DefaultVersionId field to given value.

### HasDefaultVersionId

`func (o *PolicyCreateRequest) HasDefaultVersionId() bool`

HasDefaultVersionId returns a boolean if a field has been set.

### SetDefaultVersionIdNil

`func (o *PolicyCreateRequest) SetDefaultVersionIdNil(b bool)`

 SetDefaultVersionIdNil sets the value for DefaultVersionId to be an explicit nil

### UnsetDefaultVersionId
`func (o *PolicyCreateRequest) UnsetDefaultVersionId()`

UnsetDefaultVersionId ensures that no value is present for DefaultVersionId, not even an explicit nil
### GetDescription

`func (o *PolicyCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyCategory

`func (o *PolicyCreateRequest) GetPolicyCategory() PolicyCategoryEnum`

GetPolicyCategory returns the PolicyCategory field if non-nil, zero value otherwise.

### GetPolicyCategoryOk

`func (o *PolicyCreateRequest) GetPolicyCategoryOk() (*PolicyCategoryEnum, bool)`

GetPolicyCategoryOk returns a tuple with the PolicyCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCategory

`func (o *PolicyCreateRequest) SetPolicyCategory(v PolicyCategoryEnum)`

SetPolicyCategory sets PolicyCategory field to given value.

### HasPolicyCategory

`func (o *PolicyCreateRequest) HasPolicyCategory() bool`

HasPolicyCategory returns a boolean if a field has been set.

### GetPolicyName

`func (o *PolicyCreateRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyCreateRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyCreateRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetPolicyType

`func (o *PolicyCreateRequest) GetPolicyType() PolicyTypeEnum`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *PolicyCreateRequest) GetPolicyTypeOk() (*PolicyTypeEnum, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *PolicyCreateRequest) SetPolicyType(v PolicyTypeEnum)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *PolicyCreateRequest) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetPolicyVersion

`func (o *PolicyCreateRequest) GetPolicyVersion() PolicyVersionCreateRequest`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *PolicyCreateRequest) GetPolicyVersionOk() (*PolicyVersionCreateRequest, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *PolicyCreateRequest) SetPolicyVersion(v PolicyVersionCreateRequest)`

SetPolicyVersion sets PolicyVersion field to given value.


### GetServiceType

`func (o *PolicyCreateRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PolicyCreateRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PolicyCreateRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *PolicyCreateRequest) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### SetServiceTypeNil

`func (o *PolicyCreateRequest) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *PolicyCreateRequest) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetTags

`func (o *PolicyCreateRequest) GetTags() []map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PolicyCreateRequest) GetTagsOk() (*[]map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PolicyCreateRequest) SetTags(v []map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PolicyCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *PolicyCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *PolicyCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


