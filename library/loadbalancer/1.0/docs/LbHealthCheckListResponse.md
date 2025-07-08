# LbHealthCheckListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**LbHealthChecks** | [**[]LbHealthCheckList**](LbHealthCheckList.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLbHealthCheckListResponse

`func NewLbHealthCheckListResponse(count int32, lbHealthChecks []LbHealthCheckList, page int32, size int32, ) *LbHealthCheckListResponse`

NewLbHealthCheckListResponse instantiates a new LbHealthCheckListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbHealthCheckListResponseWithDefaults

`func NewLbHealthCheckListResponseWithDefaults() *LbHealthCheckListResponse`

NewLbHealthCheckListResponseWithDefaults instantiates a new LbHealthCheckListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *LbHealthCheckListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LbHealthCheckListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LbHealthCheckListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLbHealthChecks

`func (o *LbHealthCheckListResponse) GetLbHealthChecks() []LbHealthCheckList`

GetLbHealthChecks returns the LbHealthChecks field if non-nil, zero value otherwise.

### GetLbHealthChecksOk

`func (o *LbHealthCheckListResponse) GetLbHealthChecksOk() (*[]LbHealthCheckList, bool)`

GetLbHealthChecksOk returns a tuple with the LbHealthChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbHealthChecks

`func (o *LbHealthCheckListResponse) SetLbHealthChecks(v []LbHealthCheckList)`

SetLbHealthChecks sets LbHealthChecks field to given value.


### GetPage

`func (o *LbHealthCheckListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LbHealthCheckListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LbHealthCheckListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *LbHealthCheckListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LbHealthCheckListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LbHealthCheckListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *LbHealthCheckListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LbHealthCheckListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LbHealthCheckListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LbHealthCheckListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *LbHealthCheckListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *LbHealthCheckListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


