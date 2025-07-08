# GslbResourceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**GslbResources** | [**[]GslbResourceResponse**](GslbResourceResponse.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGslbResourceListResponse

`func NewGslbResourceListResponse(count int32, gslbResources []GslbResourceResponse, page int32, size int32, ) *GslbResourceListResponse`

NewGslbResourceListResponse instantiates a new GslbResourceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGslbResourceListResponseWithDefaults

`func NewGslbResourceListResponseWithDefaults() *GslbResourceListResponse`

NewGslbResourceListResponseWithDefaults instantiates a new GslbResourceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GslbResourceListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GslbResourceListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GslbResourceListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetGslbResources

`func (o *GslbResourceListResponse) GetGslbResources() []GslbResourceResponse`

GetGslbResources returns the GslbResources field if non-nil, zero value otherwise.

### GetGslbResourcesOk

`func (o *GslbResourceListResponse) GetGslbResourcesOk() (*[]GslbResourceResponse, bool)`

GetGslbResourcesOk returns a tuple with the GslbResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbResources

`func (o *GslbResourceListResponse) SetGslbResources(v []GslbResourceResponse)`

SetGslbResources sets GslbResources field to given value.


### GetPage

`func (o *GslbResourceListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GslbResourceListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GslbResourceListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *GslbResourceListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GslbResourceListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GslbResourceListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *GslbResourceListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GslbResourceListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GslbResourceListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GslbResourceListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *GslbResourceListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *GslbResourceListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


