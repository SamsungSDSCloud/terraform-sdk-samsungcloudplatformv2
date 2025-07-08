# AddrbookAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberCount** | **int64** | 멤버수 | 
**UserAdditionalInfo** | Pointer to [**[]Member**](Member.md) |  | [optional] 

## Methods

### NewAddrbookAdditionalInfo

`func NewAddrbookAdditionalInfo(memberCount int64, ) *AddrbookAdditionalInfo`

NewAddrbookAdditionalInfo instantiates a new AddrbookAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddrbookAdditionalInfoWithDefaults

`func NewAddrbookAdditionalInfoWithDefaults() *AddrbookAdditionalInfo`

NewAddrbookAdditionalInfoWithDefaults instantiates a new AddrbookAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberCount

`func (o *AddrbookAdditionalInfo) GetMemberCount() int64`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *AddrbookAdditionalInfo) GetMemberCountOk() (*int64, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *AddrbookAdditionalInfo) SetMemberCount(v int64)`

SetMemberCount sets MemberCount field to given value.


### GetUserAdditionalInfo

`func (o *AddrbookAdditionalInfo) GetUserAdditionalInfo() []Member`

GetUserAdditionalInfo returns the UserAdditionalInfo field if non-nil, zero value otherwise.

### GetUserAdditionalInfoOk

`func (o *AddrbookAdditionalInfo) GetUserAdditionalInfoOk() (*[]Member, bool)`

GetUserAdditionalInfoOk returns a tuple with the UserAdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAdditionalInfo

`func (o *AddrbookAdditionalInfo) SetUserAdditionalInfo(v []Member)`

SetUserAdditionalInfo sets UserAdditionalInfo field to given value.

### HasUserAdditionalInfo

`func (o *AddrbookAdditionalInfo) HasUserAdditionalInfo() bool`

HasUserAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


