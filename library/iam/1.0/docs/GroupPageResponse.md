# GroupPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Groups** | [**[]Group**](Group.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroupPageResponse

`func NewGroupPageResponse(count int32, groups []Group, page int32, size int32, ) *GroupPageResponse`

NewGroupPageResponse instantiates a new GroupPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPageResponseWithDefaults

`func NewGroupPageResponseWithDefaults() *GroupPageResponse`

NewGroupPageResponseWithDefaults instantiates a new GroupPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GroupPageResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupPageResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupPageResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetGroups

`func (o *GroupPageResponse) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupPageResponse) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupPageResponse) SetGroups(v []Group)`

SetGroups sets Groups field to given value.


### GetPage

`func (o *GroupPageResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GroupPageResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GroupPageResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *GroupPageResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GroupPageResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GroupPageResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *GroupPageResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GroupPageResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GroupPageResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GroupPageResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *GroupPageResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *GroupPageResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


