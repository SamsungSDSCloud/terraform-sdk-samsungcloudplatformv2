# GslbListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Gslbs** | [**[]GslbResponseCommon**](GslbResponseCommon.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGslbListResponse

`func NewGslbListResponse(count int32, gslbs []GslbResponseCommon, page int32, size int32, ) *GslbListResponse`

NewGslbListResponse instantiates a new GslbListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGslbListResponseWithDefaults

`func NewGslbListResponseWithDefaults() *GslbListResponse`

NewGslbListResponseWithDefaults instantiates a new GslbListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GslbListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GslbListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GslbListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetGslbs

`func (o *GslbListResponse) GetGslbs() []GslbResponseCommon`

GetGslbs returns the Gslbs field if non-nil, zero value otherwise.

### GetGslbsOk

`func (o *GslbListResponse) GetGslbsOk() (*[]GslbResponseCommon, bool)`

GetGslbsOk returns a tuple with the Gslbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbs

`func (o *GslbListResponse) SetGslbs(v []GslbResponseCommon)`

SetGslbs sets Gslbs field to given value.


### GetPage

`func (o *GslbListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GslbListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GslbListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *GslbListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GslbListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GslbListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *GslbListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GslbListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GslbListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GslbListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *GslbListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *GslbListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


