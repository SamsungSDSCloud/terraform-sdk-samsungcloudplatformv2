# GroupMemberDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreValidateProjectUser** | Pointer to **bool** | Project User가 최소 1개 Group에 속해야 하는 규칙 무시 | [optional] [default to false]

## Methods

### NewGroupMemberDeleteRequest

`func NewGroupMemberDeleteRequest() *GroupMemberDeleteRequest`

NewGroupMemberDeleteRequest instantiates a new GroupMemberDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberDeleteRequestWithDefaults

`func NewGroupMemberDeleteRequestWithDefaults() *GroupMemberDeleteRequest`

NewGroupMemberDeleteRequestWithDefaults instantiates a new GroupMemberDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreValidateProjectUser

`func (o *GroupMemberDeleteRequest) GetIgnoreValidateProjectUser() bool`

GetIgnoreValidateProjectUser returns the IgnoreValidateProjectUser field if non-nil, zero value otherwise.

### GetIgnoreValidateProjectUserOk

`func (o *GroupMemberDeleteRequest) GetIgnoreValidateProjectUserOk() (*bool, bool)`

GetIgnoreValidateProjectUserOk returns a tuple with the IgnoreValidateProjectUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreValidateProjectUser

`func (o *GroupMemberDeleteRequest) SetIgnoreValidateProjectUser(v bool)`

SetIgnoreValidateProjectUser sets IgnoreValidateProjectUser field to given value.

### HasIgnoreValidateProjectUser

`func (o *GroupMemberDeleteRequest) HasIgnoreValidateProjectUser() bool`

HasIgnoreValidateProjectUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


