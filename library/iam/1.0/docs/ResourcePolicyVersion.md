# ResourcePolicyVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | Pointer to **string** | Policy Version Id | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**PolicyDocument** | [**IamPolicyDocument**](IamPolicyDocument.md) | Policy Document | 

## Methods

### NewResourcePolicyVersion

`func NewResourcePolicyVersion(createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, policyDocument IamPolicyDocument, ) *ResourcePolicyVersion`

NewResourcePolicyVersion instantiates a new ResourcePolicyVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePolicyVersionWithDefaults

`func NewResourcePolicyVersionWithDefaults() *ResourcePolicyVersion`

NewResourcePolicyVersionWithDefaults instantiates a new ResourcePolicyVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ResourcePolicyVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourcePolicyVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourcePolicyVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ResourcePolicyVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ResourcePolicyVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ResourcePolicyVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *ResourcePolicyVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcePolicyVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcePolicyVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourcePolicyVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ResourcePolicyVersion) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ResourcePolicyVersion) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ResourcePolicyVersion) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *ResourcePolicyVersion) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ResourcePolicyVersion) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ResourcePolicyVersion) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetPolicyDocument

`func (o *ResourcePolicyVersion) GetPolicyDocument() IamPolicyDocument`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *ResourcePolicyVersion) GetPolicyDocumentOk() (*IamPolicyDocument, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *ResourcePolicyVersion) SetPolicyDocument(v IamPolicyDocument)`

SetPolicyDocument sets PolicyDocument field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


