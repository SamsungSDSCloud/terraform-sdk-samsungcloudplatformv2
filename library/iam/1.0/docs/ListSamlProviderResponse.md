# ListSamlProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**SamlProviders** | [**[]SamlProviderResponse**](SamlProviderResponse.md) |  | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewListSamlProviderResponse

`func NewListSamlProviderResponse(count int32, page int32, samlProviders []SamlProviderResponse, size int32, ) *ListSamlProviderResponse`

NewListSamlProviderResponse instantiates a new ListSamlProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSamlProviderResponseWithDefaults

`func NewListSamlProviderResponseWithDefaults() *ListSamlProviderResponse`

NewListSamlProviderResponseWithDefaults instantiates a new ListSamlProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListSamlProviderResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListSamlProviderResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListSamlProviderResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *ListSamlProviderResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListSamlProviderResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListSamlProviderResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSamlProviders

`func (o *ListSamlProviderResponse) GetSamlProviders() []SamlProviderResponse`

GetSamlProviders returns the SamlProviders field if non-nil, zero value otherwise.

### GetSamlProvidersOk

`func (o *ListSamlProviderResponse) GetSamlProvidersOk() (*[]SamlProviderResponse, bool)`

GetSamlProvidersOk returns a tuple with the SamlProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlProviders

`func (o *ListSamlProviderResponse) SetSamlProviders(v []SamlProviderResponse)`

SetSamlProviders sets SamlProviders field to given value.


### GetSize

`func (o *ListSamlProviderResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListSamlProviderResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListSamlProviderResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *ListSamlProviderResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ListSamlProviderResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ListSamlProviderResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ListSamlProviderResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *ListSamlProviderResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *ListSamlProviderResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


