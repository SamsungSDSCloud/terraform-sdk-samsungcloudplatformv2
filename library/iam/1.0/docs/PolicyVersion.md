# PolicyVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | Pointer to **string** | Policy Version Id | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**PolicyDocument** | [**IamPolicyDocument**](IamPolicyDocument.md) | Policy Document | 
**PolicyId** | Pointer to **string** | Policy ID | [optional] 
**PolicyVersionName** | Pointer to **string** | Policy Version Name | [optional] 

## Methods

### NewPolicyVersion

`func NewPolicyVersion(createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, policyDocument IamPolicyDocument, ) *PolicyVersion`

NewPolicyVersion instantiates a new PolicyVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyVersionWithDefaults

`func NewPolicyVersionWithDefaults() *PolicyVersion`

NewPolicyVersionWithDefaults instantiates a new PolicyVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PolicyVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PolicyVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PolicyVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PolicyVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *PolicyVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *PolicyVersion) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PolicyVersion) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PolicyVersion) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *PolicyVersion) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PolicyVersion) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PolicyVersion) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetPolicyDocument

`func (o *PolicyVersion) GetPolicyDocument() IamPolicyDocument`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *PolicyVersion) GetPolicyDocumentOk() (*IamPolicyDocument, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *PolicyVersion) SetPolicyDocument(v IamPolicyDocument)`

SetPolicyDocument sets PolicyDocument field to given value.


### GetPolicyId

`func (o *PolicyVersion) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyVersion) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyVersion) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyVersion) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyVersionName

`func (o *PolicyVersion) GetPolicyVersionName() string`

GetPolicyVersionName returns the PolicyVersionName field if non-nil, zero value otherwise.

### GetPolicyVersionNameOk

`func (o *PolicyVersion) GetPolicyVersionNameOk() (*string, bool)`

GetPolicyVersionNameOk returns a tuple with the PolicyVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersionName

`func (o *PolicyVersion) SetPolicyVersionName(v string)`

SetPolicyVersionName sets PolicyVersionName field to given value.

### HasPolicyVersionName

`func (o *PolicyVersion) HasPolicyVersionName() bool`

HasPolicyVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


