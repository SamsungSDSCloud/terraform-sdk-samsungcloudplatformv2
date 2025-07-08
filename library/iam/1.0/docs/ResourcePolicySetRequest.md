# ResourcePolicySetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVersionId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyCategory** | Pointer to [**NullablePolicyCategoryEnum**](PolicyCategoryEnum.md) |  | [optional] 
**PolicyName** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**PolicyTypeEnum**](PolicyTypeEnum.md) | Policy Type | [optional] [default to POLICYTYPEENUM_USER_DEFINED]
**PolicyVersion** | [**PolicyVersionCreateRequest**](PolicyVersionCreateRequest.md) | Description | 
**ServiceType** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]map[string]string** |  | [optional] 

## Methods

### NewResourcePolicySetRequest

`func NewResourcePolicySetRequest(policyVersion PolicyVersionCreateRequest, ) *ResourcePolicySetRequest`

NewResourcePolicySetRequest instantiates a new ResourcePolicySetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePolicySetRequestWithDefaults

`func NewResourcePolicySetRequestWithDefaults() *ResourcePolicySetRequest`

NewResourcePolicySetRequestWithDefaults instantiates a new ResourcePolicySetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVersionId

`func (o *ResourcePolicySetRequest) GetDefaultVersionId() string`

GetDefaultVersionId returns the DefaultVersionId field if non-nil, zero value otherwise.

### GetDefaultVersionIdOk

`func (o *ResourcePolicySetRequest) GetDefaultVersionIdOk() (*string, bool)`

GetDefaultVersionIdOk returns a tuple with the DefaultVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionId

`func (o *ResourcePolicySetRequest) SetDefaultVersionId(v string)`

SetDefaultVersionId sets DefaultVersionId field to given value.

### HasDefaultVersionId

`func (o *ResourcePolicySetRequest) HasDefaultVersionId() bool`

HasDefaultVersionId returns a boolean if a field has been set.

### SetDefaultVersionIdNil

`func (o *ResourcePolicySetRequest) SetDefaultVersionIdNil(b bool)`

 SetDefaultVersionIdNil sets the value for DefaultVersionId to be an explicit nil

### UnsetDefaultVersionId
`func (o *ResourcePolicySetRequest) UnsetDefaultVersionId()`

UnsetDefaultVersionId ensures that no value is present for DefaultVersionId, not even an explicit nil
### GetDescription

`func (o *ResourcePolicySetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourcePolicySetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourcePolicySetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourcePolicySetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ResourcePolicySetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResourcePolicySetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyCategory

`func (o *ResourcePolicySetRequest) GetPolicyCategory() PolicyCategoryEnum`

GetPolicyCategory returns the PolicyCategory field if non-nil, zero value otherwise.

### GetPolicyCategoryOk

`func (o *ResourcePolicySetRequest) GetPolicyCategoryOk() (*PolicyCategoryEnum, bool)`

GetPolicyCategoryOk returns a tuple with the PolicyCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCategory

`func (o *ResourcePolicySetRequest) SetPolicyCategory(v PolicyCategoryEnum)`

SetPolicyCategory sets PolicyCategory field to given value.

### HasPolicyCategory

`func (o *ResourcePolicySetRequest) HasPolicyCategory() bool`

HasPolicyCategory returns a boolean if a field has been set.

### SetPolicyCategoryNil

`func (o *ResourcePolicySetRequest) SetPolicyCategoryNil(b bool)`

 SetPolicyCategoryNil sets the value for PolicyCategory to be an explicit nil

### UnsetPolicyCategory
`func (o *ResourcePolicySetRequest) UnsetPolicyCategory()`

UnsetPolicyCategory ensures that no value is present for PolicyCategory, not even an explicit nil
### GetPolicyName

`func (o *ResourcePolicySetRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ResourcePolicySetRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ResourcePolicySetRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ResourcePolicySetRequest) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *ResourcePolicySetRequest) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *ResourcePolicySetRequest) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPolicyType

`func (o *ResourcePolicySetRequest) GetPolicyType() PolicyTypeEnum`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *ResourcePolicySetRequest) GetPolicyTypeOk() (*PolicyTypeEnum, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *ResourcePolicySetRequest) SetPolicyType(v PolicyTypeEnum)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *ResourcePolicySetRequest) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetPolicyVersion

`func (o *ResourcePolicySetRequest) GetPolicyVersion() PolicyVersionCreateRequest`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *ResourcePolicySetRequest) GetPolicyVersionOk() (*PolicyVersionCreateRequest, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *ResourcePolicySetRequest) SetPolicyVersion(v PolicyVersionCreateRequest)`

SetPolicyVersion sets PolicyVersion field to given value.


### GetServiceType

`func (o *ResourcePolicySetRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ResourcePolicySetRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ResourcePolicySetRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ResourcePolicySetRequest) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### SetServiceTypeNil

`func (o *ResourcePolicySetRequest) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *ResourcePolicySetRequest) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetTags

`func (o *ResourcePolicySetRequest) GetTags() []map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResourcePolicySetRequest) GetTagsOk() (*[]map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResourcePolicySetRequest) SetTags(v []map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResourcePolicySetRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ResourcePolicySetRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ResourcePolicySetRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


