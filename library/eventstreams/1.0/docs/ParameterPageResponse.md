# ParameterPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]Parameters**](Parameters.md) | Parameter list | 
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewParameterPageResponse

`func NewParameterPageResponse(contents []Parameters, count int32, page int32, size int32, ) *ParameterPageResponse`

NewParameterPageResponse instantiates a new ParameterPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterPageResponseWithDefaults

`func NewParameterPageResponseWithDefaults() *ParameterPageResponse`

NewParameterPageResponseWithDefaults instantiates a new ParameterPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *ParameterPageResponse) GetContents() []Parameters`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ParameterPageResponse) GetContentsOk() (*[]Parameters, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ParameterPageResponse) SetContents(v []Parameters)`

SetContents sets Contents field to given value.


### GetCount

`func (o *ParameterPageResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ParameterPageResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ParameterPageResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *ParameterPageResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ParameterPageResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ParameterPageResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *ParameterPageResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ParameterPageResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ParameterPageResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *ParameterPageResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ParameterPageResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ParameterPageResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ParameterPageResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *ParameterPageResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *ParameterPageResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


