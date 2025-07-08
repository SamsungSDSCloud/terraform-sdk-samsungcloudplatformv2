# CreatedGroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | 생성 일시 | 
**CreatedBy** | **string** | 생성자 | 
**CreatorEmail** | Pointer to **string** | 생성자 Email | [optional] [default to "-"]
**CreatorName** | Pointer to **string** | 생성자 성, 이름 | [optional] [default to "-"]
**UserEmail** | Pointer to **string** | User Email | [optional] [default to "-"]
**UserId** | **string** | User ID | 
**UserName** | Pointer to **string** | User 성, 이름 | [optional] [default to "-"]

## Methods

### NewCreatedGroupMember

`func NewCreatedGroupMember(createdAt time.Time, createdBy string, userId string, ) *CreatedGroupMember`

NewCreatedGroupMember instantiates a new CreatedGroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedGroupMemberWithDefaults

`func NewCreatedGroupMemberWithDefaults() *CreatedGroupMember`

NewCreatedGroupMemberWithDefaults instantiates a new CreatedGroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CreatedGroupMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreatedGroupMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreatedGroupMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *CreatedGroupMember) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreatedGroupMember) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreatedGroupMember) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorEmail

`func (o *CreatedGroupMember) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *CreatedGroupMember) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *CreatedGroupMember) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *CreatedGroupMember) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetCreatorName

`func (o *CreatedGroupMember) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *CreatedGroupMember) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *CreatedGroupMember) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *CreatedGroupMember) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetUserEmail

`func (o *CreatedGroupMember) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *CreatedGroupMember) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *CreatedGroupMember) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *CreatedGroupMember) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserId

`func (o *CreatedGroupMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreatedGroupMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreatedGroupMember) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserName

`func (o *CreatedGroupMember) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreatedGroupMember) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreatedGroupMember) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *CreatedGroupMember) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


