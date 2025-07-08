# LoadbalancerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Loadbalancers** | [**[]LoadbalancerListResponseDetail**](LoadbalancerListResponseDetail.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLoadbalancerListResponse

`func NewLoadbalancerListResponse(count int32, loadbalancers []LoadbalancerListResponseDetail, page int32, size int32, ) *LoadbalancerListResponse`

NewLoadbalancerListResponse instantiates a new LoadbalancerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerListResponseWithDefaults

`func NewLoadbalancerListResponseWithDefaults() *LoadbalancerListResponse`

NewLoadbalancerListResponseWithDefaults instantiates a new LoadbalancerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *LoadbalancerListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LoadbalancerListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LoadbalancerListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLoadbalancers

`func (o *LoadbalancerListResponse) GetLoadbalancers() []LoadbalancerListResponseDetail`

GetLoadbalancers returns the Loadbalancers field if non-nil, zero value otherwise.

### GetLoadbalancersOk

`func (o *LoadbalancerListResponse) GetLoadbalancersOk() (*[]LoadbalancerListResponseDetail, bool)`

GetLoadbalancersOk returns a tuple with the Loadbalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancers

`func (o *LoadbalancerListResponse) SetLoadbalancers(v []LoadbalancerListResponseDetail)`

SetLoadbalancers sets Loadbalancers field to given value.


### GetPage

`func (o *LoadbalancerListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LoadbalancerListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LoadbalancerListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *LoadbalancerListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LoadbalancerListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LoadbalancerListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *LoadbalancerListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LoadbalancerListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LoadbalancerListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LoadbalancerListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *LoadbalancerListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *LoadbalancerListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


