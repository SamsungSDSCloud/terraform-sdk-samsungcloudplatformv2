# MemberWithHealthStateListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Members** | [**[]MemberWithHealthState**](MemberWithHealthState.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMemberWithHealthStateListResponse

`func NewMemberWithHealthStateListResponse(count int32, members []MemberWithHealthState, page int32, size int32, ) *MemberWithHealthStateListResponse`

NewMemberWithHealthStateListResponse instantiates a new MemberWithHealthStateListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithHealthStateListResponseWithDefaults

`func NewMemberWithHealthStateListResponseWithDefaults() *MemberWithHealthStateListResponse`

NewMemberWithHealthStateListResponseWithDefaults instantiates a new MemberWithHealthStateListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *MemberWithHealthStateListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MemberWithHealthStateListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MemberWithHealthStateListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetMembers

`func (o *MemberWithHealthStateListResponse) GetMembers() []MemberWithHealthState`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *MemberWithHealthStateListResponse) GetMembersOk() (*[]MemberWithHealthState, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *MemberWithHealthStateListResponse) SetMembers(v []MemberWithHealthState)`

SetMembers sets Members field to given value.


### GetPage

`func (o *MemberWithHealthStateListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *MemberWithHealthStateListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *MemberWithHealthStateListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *MemberWithHealthStateListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MemberWithHealthStateListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MemberWithHealthStateListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *MemberWithHealthStateListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *MemberWithHealthStateListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *MemberWithHealthStateListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *MemberWithHealthStateListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *MemberWithHealthStateListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *MemberWithHealthStateListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


