# RdbClusterPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]RdbClusterResponse**](RdbClusterResponse.md) |  | 
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRdbClusterPageResponse

`func NewRdbClusterPageResponse(contents []RdbClusterResponse, count int32, page int32, size int32, ) *RdbClusterPageResponse`

NewRdbClusterPageResponse instantiates a new RdbClusterPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdbClusterPageResponseWithDefaults

`func NewRdbClusterPageResponseWithDefaults() *RdbClusterPageResponse`

NewRdbClusterPageResponseWithDefaults instantiates a new RdbClusterPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *RdbClusterPageResponse) GetContents() []RdbClusterResponse`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *RdbClusterPageResponse) GetContentsOk() (*[]RdbClusterResponse, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *RdbClusterPageResponse) SetContents(v []RdbClusterResponse)`

SetContents sets Contents field to given value.


### SetContentsNil

`func (o *RdbClusterPageResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *RdbClusterPageResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCount

`func (o *RdbClusterPageResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RdbClusterPageResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RdbClusterPageResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *RdbClusterPageResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *RdbClusterPageResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *RdbClusterPageResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *RdbClusterPageResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RdbClusterPageResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RdbClusterPageResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *RdbClusterPageResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RdbClusterPageResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RdbClusterPageResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RdbClusterPageResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *RdbClusterPageResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *RdbClusterPageResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


