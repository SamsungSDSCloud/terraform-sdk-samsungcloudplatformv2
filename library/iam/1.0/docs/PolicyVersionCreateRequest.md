# PolicyVersionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyDocument** | [**IamPolicyDocument**](IamPolicyDocument.md) | Policy Document | 
**PolicyId** | Pointer to **string** | Policy ID | [optional] 
**PolicyVersionId** | Pointer to **string** | Policy Version Id | [optional] 
**PolicyVersionName** | Pointer to **string** | Policy Version Name | [optional] [default to "POLICY_VERSION_1"]

## Methods

### NewPolicyVersionCreateRequest

`func NewPolicyVersionCreateRequest(policyDocument IamPolicyDocument, ) *PolicyVersionCreateRequest`

NewPolicyVersionCreateRequest instantiates a new PolicyVersionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyVersionCreateRequestWithDefaults

`func NewPolicyVersionCreateRequestWithDefaults() *PolicyVersionCreateRequest`

NewPolicyVersionCreateRequestWithDefaults instantiates a new PolicyVersionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyDocument

`func (o *PolicyVersionCreateRequest) GetPolicyDocument() IamPolicyDocument`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *PolicyVersionCreateRequest) GetPolicyDocumentOk() (*IamPolicyDocument, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *PolicyVersionCreateRequest) SetPolicyDocument(v IamPolicyDocument)`

SetPolicyDocument sets PolicyDocument field to given value.


### GetPolicyId

`func (o *PolicyVersionCreateRequest) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyVersionCreateRequest) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyVersionCreateRequest) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyVersionCreateRequest) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyVersionId

`func (o *PolicyVersionCreateRequest) GetPolicyVersionId() string`

GetPolicyVersionId returns the PolicyVersionId field if non-nil, zero value otherwise.

### GetPolicyVersionIdOk

`func (o *PolicyVersionCreateRequest) GetPolicyVersionIdOk() (*string, bool)`

GetPolicyVersionIdOk returns a tuple with the PolicyVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersionId

`func (o *PolicyVersionCreateRequest) SetPolicyVersionId(v string)`

SetPolicyVersionId sets PolicyVersionId field to given value.

### HasPolicyVersionId

`func (o *PolicyVersionCreateRequest) HasPolicyVersionId() bool`

HasPolicyVersionId returns a boolean if a field has been set.

### GetPolicyVersionName

`func (o *PolicyVersionCreateRequest) GetPolicyVersionName() string`

GetPolicyVersionName returns the PolicyVersionName field if non-nil, zero value otherwise.

### GetPolicyVersionNameOk

`func (o *PolicyVersionCreateRequest) GetPolicyVersionNameOk() (*string, bool)`

GetPolicyVersionNameOk returns a tuple with the PolicyVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersionName

`func (o *PolicyVersionCreateRequest) SetPolicyVersionName(v string)`

SetPolicyVersionName sets PolicyVersionName field to given value.

### HasPolicyVersionName

`func (o *PolicyVersionCreateRequest) HasPolicyVersionName() bool`

HasPolicyVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


