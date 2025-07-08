# LbServerGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**LbServerGroups** | [**[]LbServerGroupList**](LbServerGroupList.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLbServerGroupListResponse

`func NewLbServerGroupListResponse(count int32, lbServerGroups []LbServerGroupList, page int32, size int32, ) *LbServerGroupListResponse`

NewLbServerGroupListResponse instantiates a new LbServerGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbServerGroupListResponseWithDefaults

`func NewLbServerGroupListResponseWithDefaults() *LbServerGroupListResponse`

NewLbServerGroupListResponseWithDefaults instantiates a new LbServerGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *LbServerGroupListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LbServerGroupListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LbServerGroupListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLbServerGroups

`func (o *LbServerGroupListResponse) GetLbServerGroups() []LbServerGroupList`

GetLbServerGroups returns the LbServerGroups field if non-nil, zero value otherwise.

### GetLbServerGroupsOk

`func (o *LbServerGroupListResponse) GetLbServerGroupsOk() (*[]LbServerGroupList, bool)`

GetLbServerGroupsOk returns a tuple with the LbServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbServerGroups

`func (o *LbServerGroupListResponse) SetLbServerGroups(v []LbServerGroupList)`

SetLbServerGroups sets LbServerGroups field to given value.


### GetPage

`func (o *LbServerGroupListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LbServerGroupListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LbServerGroupListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *LbServerGroupListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LbServerGroupListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LbServerGroupListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *LbServerGroupListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LbServerGroupListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LbServerGroupListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LbServerGroupListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *LbServerGroupListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *LbServerGroupListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


