# GroupMemberPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**GroupMembers** | [**[]GroupMember**](GroupMember.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroupMemberPageResponse

`func NewGroupMemberPageResponse(count int32, groupMembers []GroupMember, page int32, size int32, ) *GroupMemberPageResponse`

NewGroupMemberPageResponse instantiates a new GroupMemberPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberPageResponseWithDefaults

`func NewGroupMemberPageResponseWithDefaults() *GroupMemberPageResponse`

NewGroupMemberPageResponseWithDefaults instantiates a new GroupMemberPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GroupMemberPageResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupMemberPageResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupMemberPageResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetGroupMembers

`func (o *GroupMemberPageResponse) GetGroupMembers() []GroupMember`

GetGroupMembers returns the GroupMembers field if non-nil, zero value otherwise.

### GetGroupMembersOk

`func (o *GroupMemberPageResponse) GetGroupMembersOk() (*[]GroupMember, bool)`

GetGroupMembersOk returns a tuple with the GroupMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembers

`func (o *GroupMemberPageResponse) SetGroupMembers(v []GroupMember)`

SetGroupMembers sets GroupMembers field to given value.


### GetPage

`func (o *GroupMemberPageResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GroupMemberPageResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GroupMemberPageResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *GroupMemberPageResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GroupMemberPageResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GroupMemberPageResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *GroupMemberPageResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GroupMemberPageResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GroupMemberPageResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GroupMemberPageResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *GroupMemberPageResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *GroupMemberPageResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


