# PolicySetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVersionId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyName** | Pointer to **NullableString** |  | [optional] 
**PolicyVersion** | Pointer to [**NullablePolicyVersionCreateRequest**](PolicyVersionCreateRequest.md) |  | [optional] 

## Methods

### NewPolicySetRequest

`func NewPolicySetRequest() *PolicySetRequest`

NewPolicySetRequest instantiates a new PolicySetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicySetRequestWithDefaults

`func NewPolicySetRequestWithDefaults() *PolicySetRequest`

NewPolicySetRequestWithDefaults instantiates a new PolicySetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVersionId

`func (o *PolicySetRequest) GetDefaultVersionId() string`

GetDefaultVersionId returns the DefaultVersionId field if non-nil, zero value otherwise.

### GetDefaultVersionIdOk

`func (o *PolicySetRequest) GetDefaultVersionIdOk() (*string, bool)`

GetDefaultVersionIdOk returns a tuple with the DefaultVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersionId

`func (o *PolicySetRequest) SetDefaultVersionId(v string)`

SetDefaultVersionId sets DefaultVersionId field to given value.

### HasDefaultVersionId

`func (o *PolicySetRequest) HasDefaultVersionId() bool`

HasDefaultVersionId returns a boolean if a field has been set.

### SetDefaultVersionIdNil

`func (o *PolicySetRequest) SetDefaultVersionIdNil(b bool)`

 SetDefaultVersionIdNil sets the value for DefaultVersionId to be an explicit nil

### UnsetDefaultVersionId
`func (o *PolicySetRequest) UnsetDefaultVersionId()`

UnsetDefaultVersionId ensures that no value is present for DefaultVersionId, not even an explicit nil
### GetDescription

`func (o *PolicySetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicySetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicySetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicySetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicySetRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicySetRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyName

`func (o *PolicySetRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicySetRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicySetRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicySetRequest) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *PolicySetRequest) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *PolicySetRequest) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPolicyVersion

`func (o *PolicySetRequest) GetPolicyVersion() PolicyVersionCreateRequest`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *PolicySetRequest) GetPolicyVersionOk() (*PolicyVersionCreateRequest, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *PolicySetRequest) SetPolicyVersion(v PolicyVersionCreateRequest)`

SetPolicyVersion sets PolicyVersion field to given value.

### HasPolicyVersion

`func (o *PolicySetRequest) HasPolicyVersion() bool`

HasPolicyVersion returns a boolean if a field has been set.

### SetPolicyVersionNil

`func (o *PolicySetRequest) SetPolicyVersionNil(b bool)`

 SetPolicyVersionNil sets the value for PolicyVersion to be an explicit nil

### UnsetPolicyVersion
`func (o *PolicySetRequest) UnsetPolicyVersion()`

UnsetPolicyVersion ensures that no value is present for PolicyVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


