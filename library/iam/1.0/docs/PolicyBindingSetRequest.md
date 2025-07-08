# PolicyBindingSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddIds** | Pointer to **[]string** |  | [optional] 
**IdentityType** | **string** | Identity Type | 
**RemoveIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyBindingSetRequest

`func NewPolicyBindingSetRequest(identityType string, ) *PolicyBindingSetRequest`

NewPolicyBindingSetRequest instantiates a new PolicyBindingSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBindingSetRequestWithDefaults

`func NewPolicyBindingSetRequestWithDefaults() *PolicyBindingSetRequest`

NewPolicyBindingSetRequestWithDefaults instantiates a new PolicyBindingSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddIds

`func (o *PolicyBindingSetRequest) GetAddIds() []string`

GetAddIds returns the AddIds field if non-nil, zero value otherwise.

### GetAddIdsOk

`func (o *PolicyBindingSetRequest) GetAddIdsOk() (*[]string, bool)`

GetAddIdsOk returns a tuple with the AddIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddIds

`func (o *PolicyBindingSetRequest) SetAddIds(v []string)`

SetAddIds sets AddIds field to given value.

### HasAddIds

`func (o *PolicyBindingSetRequest) HasAddIds() bool`

HasAddIds returns a boolean if a field has been set.

### SetAddIdsNil

`func (o *PolicyBindingSetRequest) SetAddIdsNil(b bool)`

 SetAddIdsNil sets the value for AddIds to be an explicit nil

### UnsetAddIds
`func (o *PolicyBindingSetRequest) UnsetAddIds()`

UnsetAddIds ensures that no value is present for AddIds, not even an explicit nil
### GetIdentityType

`func (o *PolicyBindingSetRequest) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *PolicyBindingSetRequest) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *PolicyBindingSetRequest) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.


### GetRemoveIds

`func (o *PolicyBindingSetRequest) GetRemoveIds() []string`

GetRemoveIds returns the RemoveIds field if non-nil, zero value otherwise.

### GetRemoveIdsOk

`func (o *PolicyBindingSetRequest) GetRemoveIdsOk() (*[]string, bool)`

GetRemoveIdsOk returns a tuple with the RemoveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveIds

`func (o *PolicyBindingSetRequest) SetRemoveIds(v []string)`

SetRemoveIds sets RemoveIds field to given value.

### HasRemoveIds

`func (o *PolicyBindingSetRequest) HasRemoveIds() bool`

HasRemoveIds returns a boolean if a field has been set.

### SetRemoveIdsNil

`func (o *PolicyBindingSetRequest) SetRemoveIdsNil(b bool)`

 SetRemoveIdsNil sets the value for RemoveIds to be an explicit nil

### UnsetRemoveIds
`func (o *PolicyBindingSetRequest) UnsetRemoveIds()`

UnsetRemoveIds ensures that no value is present for RemoveIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


