# PolicyQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedPolicyIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyQueryRequest

`func NewPolicyQueryRequest() *PolicyQueryRequest`

NewPolicyQueryRequest instantiates a new PolicyQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyQueryRequestWithDefaults

`func NewPolicyQueryRequestWithDefaults() *PolicyQueryRequest`

NewPolicyQueryRequestWithDefaults instantiates a new PolicyQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedPolicyIds

`func (o *PolicyQueryRequest) GetExcludedPolicyIds() []string`

GetExcludedPolicyIds returns the ExcludedPolicyIds field if non-nil, zero value otherwise.

### GetExcludedPolicyIdsOk

`func (o *PolicyQueryRequest) GetExcludedPolicyIdsOk() (*[]string, bool)`

GetExcludedPolicyIdsOk returns a tuple with the ExcludedPolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPolicyIds

`func (o *PolicyQueryRequest) SetExcludedPolicyIds(v []string)`

SetExcludedPolicyIds sets ExcludedPolicyIds field to given value.

### HasExcludedPolicyIds

`func (o *PolicyQueryRequest) HasExcludedPolicyIds() bool`

HasExcludedPolicyIds returns a boolean if a field has been set.

### SetExcludedPolicyIdsNil

`func (o *PolicyQueryRequest) SetExcludedPolicyIdsNil(b bool)`

 SetExcludedPolicyIdsNil sets the value for ExcludedPolicyIds to be an explicit nil

### UnsetExcludedPolicyIds
`func (o *PolicyQueryRequest) UnsetExcludedPolicyIds()`

UnsetExcludedPolicyIds ensures that no value is present for ExcludedPolicyIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


