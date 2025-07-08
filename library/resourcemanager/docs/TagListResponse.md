# TagListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]SrnKeyValue**](SrnKeyValue.md) |  | 
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTagListResponse

`func NewTagListResponse(content []SrnKeyValue, count int32, page int32, size int32, ) *TagListResponse`

NewTagListResponse instantiates a new TagListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagListResponseWithDefaults

`func NewTagListResponseWithDefaults() *TagListResponse`

NewTagListResponseWithDefaults instantiates a new TagListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *TagListResponse) GetContent() []SrnKeyValue`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TagListResponse) GetContentOk() (*[]SrnKeyValue, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TagListResponse) SetContent(v []SrnKeyValue)`

SetContent sets Content field to given value.


### GetCount

`func (o *TagListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TagListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TagListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *TagListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TagListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TagListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *TagListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TagListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TagListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *TagListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TagListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TagListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TagListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *TagListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *TagListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


