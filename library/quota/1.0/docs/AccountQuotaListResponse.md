# AccountQuotaListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountQuotas** | [**[]AccountQuota**](AccountQuota.md) |  | 
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccountQuotaListResponse

`func NewAccountQuotaListResponse(accountQuotas []AccountQuota, count int32, page int32, size int32, ) *AccountQuotaListResponse`

NewAccountQuotaListResponse instantiates a new AccountQuotaListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountQuotaListResponseWithDefaults

`func NewAccountQuotaListResponseWithDefaults() *AccountQuotaListResponse`

NewAccountQuotaListResponseWithDefaults instantiates a new AccountQuotaListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountQuotas

`func (o *AccountQuotaListResponse) GetAccountQuotas() []AccountQuota`

GetAccountQuotas returns the AccountQuotas field if non-nil, zero value otherwise.

### GetAccountQuotasOk

`func (o *AccountQuotaListResponse) GetAccountQuotasOk() (*[]AccountQuota, bool)`

GetAccountQuotasOk returns a tuple with the AccountQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountQuotas

`func (o *AccountQuotaListResponse) SetAccountQuotas(v []AccountQuota)`

SetAccountQuotas sets AccountQuotas field to given value.


### GetCount

`func (o *AccountQuotaListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AccountQuotaListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AccountQuotaListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *AccountQuotaListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AccountQuotaListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AccountQuotaListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *AccountQuotaListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AccountQuotaListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AccountQuotaListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *AccountQuotaListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AccountQuotaListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AccountQuotaListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AccountQuotaListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *AccountQuotaListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *AccountQuotaListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


