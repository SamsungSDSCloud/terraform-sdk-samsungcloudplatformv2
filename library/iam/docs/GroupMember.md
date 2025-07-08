# GroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | 생성 일시 | 
**CreatedBy** | **string** | 생성자 | 
**CreatorCreatedAt** | Pointer to **time.Time** | 생성 일시 | [optional] 
**CreatorEmail** | **string** | 생성자 Email | 
**CreatorLastLoginAt** | Pointer to **NullableTime** |  | [optional] 
**CreatorName** | **string** | 생성자 성, 이름 | 
**GroupNames** | Pointer to **[]string** | Group names | [optional] [default to []]
**UserCreatedAt** | Pointer to **time.Time** | 생성 일시 | [optional] 
**UserEmail** | **string** | User Email | 
**UserId** | **string** | User ID | 
**UserLastLoginAt** | Pointer to **NullableTime** |  | [optional] 
**UserName** | **string** | User 성, 이름 | 

## Methods

### NewGroupMember

`func NewGroupMember(createdAt time.Time, createdBy string, creatorEmail string, creatorName string, userEmail string, userId string, userName string, ) *GroupMember`

NewGroupMember instantiates a new GroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberWithDefaults

`func NewGroupMemberWithDefaults() *GroupMember`

NewGroupMemberWithDefaults instantiates a new GroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *GroupMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *GroupMember) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GroupMember) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GroupMember) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorCreatedAt

`func (o *GroupMember) GetCreatorCreatedAt() time.Time`

GetCreatorCreatedAt returns the CreatorCreatedAt field if non-nil, zero value otherwise.

### GetCreatorCreatedAtOk

`func (o *GroupMember) GetCreatorCreatedAtOk() (*time.Time, bool)`

GetCreatorCreatedAtOk returns a tuple with the CreatorCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCreatedAt

`func (o *GroupMember) SetCreatorCreatedAt(v time.Time)`

SetCreatorCreatedAt sets CreatorCreatedAt field to given value.

### HasCreatorCreatedAt

`func (o *GroupMember) HasCreatorCreatedAt() bool`

HasCreatorCreatedAt returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *GroupMember) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *GroupMember) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *GroupMember) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.


### GetCreatorLastLoginAt

`func (o *GroupMember) GetCreatorLastLoginAt() time.Time`

GetCreatorLastLoginAt returns the CreatorLastLoginAt field if non-nil, zero value otherwise.

### GetCreatorLastLoginAtOk

`func (o *GroupMember) GetCreatorLastLoginAtOk() (*time.Time, bool)`

GetCreatorLastLoginAtOk returns a tuple with the CreatorLastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorLastLoginAt

`func (o *GroupMember) SetCreatorLastLoginAt(v time.Time)`

SetCreatorLastLoginAt sets CreatorLastLoginAt field to given value.

### HasCreatorLastLoginAt

`func (o *GroupMember) HasCreatorLastLoginAt() bool`

HasCreatorLastLoginAt returns a boolean if a field has been set.

### SetCreatorLastLoginAtNil

`func (o *GroupMember) SetCreatorLastLoginAtNil(b bool)`

 SetCreatorLastLoginAtNil sets the value for CreatorLastLoginAt to be an explicit nil

### UnsetCreatorLastLoginAt
`func (o *GroupMember) UnsetCreatorLastLoginAt()`

UnsetCreatorLastLoginAt ensures that no value is present for CreatorLastLoginAt, not even an explicit nil
### GetCreatorName

`func (o *GroupMember) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *GroupMember) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *GroupMember) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.


### GetGroupNames

`func (o *GroupMember) GetGroupNames() []string`

GetGroupNames returns the GroupNames field if non-nil, zero value otherwise.

### GetGroupNamesOk

`func (o *GroupMember) GetGroupNamesOk() (*[]string, bool)`

GetGroupNamesOk returns a tuple with the GroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNames

`func (o *GroupMember) SetGroupNames(v []string)`

SetGroupNames sets GroupNames field to given value.

### HasGroupNames

`func (o *GroupMember) HasGroupNames() bool`

HasGroupNames returns a boolean if a field has been set.

### GetUserCreatedAt

`func (o *GroupMember) GetUserCreatedAt() time.Time`

GetUserCreatedAt returns the UserCreatedAt field if non-nil, zero value otherwise.

### GetUserCreatedAtOk

`func (o *GroupMember) GetUserCreatedAtOk() (*time.Time, bool)`

GetUserCreatedAtOk returns a tuple with the UserCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCreatedAt

`func (o *GroupMember) SetUserCreatedAt(v time.Time)`

SetUserCreatedAt sets UserCreatedAt field to given value.

### HasUserCreatedAt

`func (o *GroupMember) HasUserCreatedAt() bool`

HasUserCreatedAt returns a boolean if a field has been set.

### GetUserEmail

`func (o *GroupMember) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *GroupMember) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *GroupMember) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserId

`func (o *GroupMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupMember) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserLastLoginAt

`func (o *GroupMember) GetUserLastLoginAt() time.Time`

GetUserLastLoginAt returns the UserLastLoginAt field if non-nil, zero value otherwise.

### GetUserLastLoginAtOk

`func (o *GroupMember) GetUserLastLoginAtOk() (*time.Time, bool)`

GetUserLastLoginAtOk returns a tuple with the UserLastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLastLoginAt

`func (o *GroupMember) SetUserLastLoginAt(v time.Time)`

SetUserLastLoginAt sets UserLastLoginAt field to given value.

### HasUserLastLoginAt

`func (o *GroupMember) HasUserLastLoginAt() bool`

HasUserLastLoginAt returns a boolean if a field has been set.

### SetUserLastLoginAtNil

`func (o *GroupMember) SetUserLastLoginAtNil(b bool)`

 SetUserLastLoginAtNil sets the value for UserLastLoginAt to be an explicit nil

### UnsetUserLastLoginAt
`func (o *GroupMember) UnsetUserLastLoginAt()`

UnsetUserLastLoginAt ensures that no value is present for UserLastLoginAt, not even an explicit nil
### GetUserName

`func (o *GroupMember) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GroupMember) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GroupMember) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


