# MemberSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberPort** | Pointer to **NullableInt32** |  | [optional] 
**MemberState** | Pointer to [**NullableStatusType**](StatusType.md) |  | [optional] 
**MemberWeight** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewMemberSet

`func NewMemberSet() *MemberSet`

NewMemberSet instantiates a new MemberSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSetWithDefaults

`func NewMemberSetWithDefaults() *MemberSet`

NewMemberSetWithDefaults instantiates a new MemberSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberPort

`func (o *MemberSet) GetMemberPort() int32`

GetMemberPort returns the MemberPort field if non-nil, zero value otherwise.

### GetMemberPortOk

`func (o *MemberSet) GetMemberPortOk() (*int32, bool)`

GetMemberPortOk returns a tuple with the MemberPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPort

`func (o *MemberSet) SetMemberPort(v int32)`

SetMemberPort sets MemberPort field to given value.

### HasMemberPort

`func (o *MemberSet) HasMemberPort() bool`

HasMemberPort returns a boolean if a field has been set.

### SetMemberPortNil

`func (o *MemberSet) SetMemberPortNil(b bool)`

 SetMemberPortNil sets the value for MemberPort to be an explicit nil

### UnsetMemberPort
`func (o *MemberSet) UnsetMemberPort()`

UnsetMemberPort ensures that no value is present for MemberPort, not even an explicit nil
### GetMemberState

`func (o *MemberSet) GetMemberState() StatusType`

GetMemberState returns the MemberState field if non-nil, zero value otherwise.

### GetMemberStateOk

`func (o *MemberSet) GetMemberStateOk() (*StatusType, bool)`

GetMemberStateOk returns a tuple with the MemberState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberState

`func (o *MemberSet) SetMemberState(v StatusType)`

SetMemberState sets MemberState field to given value.

### HasMemberState

`func (o *MemberSet) HasMemberState() bool`

HasMemberState returns a boolean if a field has been set.

### SetMemberStateNil

`func (o *MemberSet) SetMemberStateNil(b bool)`

 SetMemberStateNil sets the value for MemberState to be an explicit nil

### UnsetMemberState
`func (o *MemberSet) UnsetMemberState()`

UnsetMemberState ensures that no value is present for MemberState, not even an explicit nil
### GetMemberWeight

`func (o *MemberSet) GetMemberWeight() int32`

GetMemberWeight returns the MemberWeight field if non-nil, zero value otherwise.

### GetMemberWeightOk

`func (o *MemberSet) GetMemberWeightOk() (*int32, bool)`

GetMemberWeightOk returns a tuple with the MemberWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberWeight

`func (o *MemberSet) SetMemberWeight(v int32)`

SetMemberWeight sets MemberWeight field to given value.

### HasMemberWeight

`func (o *MemberSet) HasMemberWeight() bool`

HasMemberWeight returns a boolean if a field has been set.

### SetMemberWeightNil

`func (o *MemberSet) SetMemberWeightNil(b bool)`

 SetMemberWeightNil sets the value for MemberWeight to be an explicit nil

### UnsetMemberWeight
`func (o *MemberSet) UnsetMemberWeight()`

UnsetMemberWeight ensures that no value is present for MemberWeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


