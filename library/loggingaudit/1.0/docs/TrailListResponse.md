# TrailListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 
**Trails** | [**[]Trail**](Trail.md) |  | 

## Methods

### NewTrailListResponse

`func NewTrailListResponse(count int32, page int32, size int32, trails []Trail, ) *TrailListResponse`

NewTrailListResponse instantiates a new TrailListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailListResponseWithDefaults

`func NewTrailListResponseWithDefaults() *TrailListResponse`

NewTrailListResponseWithDefaults instantiates a new TrailListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TrailListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TrailListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TrailListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *TrailListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TrailListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TrailListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *TrailListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TrailListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TrailListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *TrailListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TrailListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TrailListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TrailListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *TrailListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *TrailListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetTrails

`func (o *TrailListResponse) GetTrails() []Trail`

GetTrails returns the Trails field if non-nil, zero value otherwise.

### GetTrailsOk

`func (o *TrailListResponse) GetTrailsOk() (*[]Trail, bool)`

GetTrailsOk returns a tuple with the Trails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrails

`func (o *TrailListResponse) SetTrails(v []Trail)`

SetTrails sets Trails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


