# LbListenerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Listeners** | [**[]ListenerForList**](ListenerForList.md) | A list of listeners. | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLbListenerListResponse

`func NewLbListenerListResponse(count int32, listeners []ListenerForList, page int32, size int32, ) *LbListenerListResponse`

NewLbListenerListResponse instantiates a new LbListenerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbListenerListResponseWithDefaults

`func NewLbListenerListResponseWithDefaults() *LbListenerListResponse`

NewLbListenerListResponseWithDefaults instantiates a new LbListenerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *LbListenerListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LbListenerListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LbListenerListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetListeners

`func (o *LbListenerListResponse) GetListeners() []ListenerForList`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *LbListenerListResponse) GetListenersOk() (*[]ListenerForList, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *LbListenerListResponse) SetListeners(v []ListenerForList)`

SetListeners sets Listeners field to given value.


### GetPage

`func (o *LbListenerListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LbListenerListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LbListenerListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *LbListenerListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LbListenerListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LbListenerListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *LbListenerListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LbListenerListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LbListenerListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LbListenerListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *LbListenerListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *LbListenerListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


