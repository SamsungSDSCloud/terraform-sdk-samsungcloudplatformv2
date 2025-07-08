# QuotaRequestListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**QuotaRequests** | [**[]QuotaRequest**](QuotaRequest.md) |  | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewQuotaRequestListResponse

`func NewQuotaRequestListResponse(count int32, page int32, quotaRequests []QuotaRequest, size int32, ) *QuotaRequestListResponse`

NewQuotaRequestListResponse instantiates a new QuotaRequestListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaRequestListResponseWithDefaults

`func NewQuotaRequestListResponseWithDefaults() *QuotaRequestListResponse`

NewQuotaRequestListResponseWithDefaults instantiates a new QuotaRequestListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *QuotaRequestListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *QuotaRequestListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *QuotaRequestListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *QuotaRequestListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *QuotaRequestListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *QuotaRequestListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetQuotaRequests

`func (o *QuotaRequestListResponse) GetQuotaRequests() []QuotaRequest`

GetQuotaRequests returns the QuotaRequests field if non-nil, zero value otherwise.

### GetQuotaRequestsOk

`func (o *QuotaRequestListResponse) GetQuotaRequestsOk() (*[]QuotaRequest, bool)`

GetQuotaRequestsOk returns a tuple with the QuotaRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaRequests

`func (o *QuotaRequestListResponse) SetQuotaRequests(v []QuotaRequest)`

SetQuotaRequests sets QuotaRequests field to given value.


### GetSize

`func (o *QuotaRequestListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *QuotaRequestListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *QuotaRequestListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *QuotaRequestListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *QuotaRequestListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *QuotaRequestListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *QuotaRequestListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *QuotaRequestListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *QuotaRequestListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


