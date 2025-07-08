# RoleTrustPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssumeRolePolicyDocument** | [**PolicyDocument**](PolicyDocument.md) | 역할 신뢰 정책 | 

## Methods

### NewRoleTrustPolicyRequest

`func NewRoleTrustPolicyRequest(assumeRolePolicyDocument PolicyDocument, ) *RoleTrustPolicyRequest`

NewRoleTrustPolicyRequest instantiates a new RoleTrustPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleTrustPolicyRequestWithDefaults

`func NewRoleTrustPolicyRequestWithDefaults() *RoleTrustPolicyRequest`

NewRoleTrustPolicyRequestWithDefaults instantiates a new RoleTrustPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssumeRolePolicyDocument

`func (o *RoleTrustPolicyRequest) GetAssumeRolePolicyDocument() PolicyDocument`

GetAssumeRolePolicyDocument returns the AssumeRolePolicyDocument field if non-nil, zero value otherwise.

### GetAssumeRolePolicyDocumentOk

`func (o *RoleTrustPolicyRequest) GetAssumeRolePolicyDocumentOk() (*PolicyDocument, bool)`

GetAssumeRolePolicyDocumentOk returns a tuple with the AssumeRolePolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumeRolePolicyDocument

`func (o *RoleTrustPolicyRequest) SetAssumeRolePolicyDocument(v PolicyDocument)`

SetAssumeRolePolicyDocument sets AssumeRolePolicyDocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


