# MemberListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**Members** | [**[]Member**](Member.md) |  | 

## Methods

### NewMemberListResponse

`func NewMemberListResponse(members []Member, ) *MemberListResponse`

NewMemberListResponse instantiates a new MemberListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberListResponseWithDefaults

`func NewMemberListResponseWithDefaults() *MemberListResponse`

NewMemberListResponseWithDefaults instantiates a new MemberListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *MemberListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MemberListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MemberListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MemberListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *MemberListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *MemberListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *MemberListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MemberListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MemberListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MemberListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *MemberListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *MemberListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMembers

`func (o *MemberListResponse) GetMembers() []Member`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MemberListResponse) GetMembersOk() (*[]Member, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MemberListResponse) SetMembers(v []Member)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


