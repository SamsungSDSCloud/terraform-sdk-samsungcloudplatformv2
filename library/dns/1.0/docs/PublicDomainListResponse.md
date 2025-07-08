# PublicDomainListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**PublicDomainNames** | [**[]PublicDomainListItem**](PublicDomainListItem.md) |  | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPublicDomainListResponse

`func NewPublicDomainListResponse(count int32, page int32, publicDomainNames []PublicDomainListItem, size int32, ) *PublicDomainListResponse`

NewPublicDomainListResponse instantiates a new PublicDomainListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDomainListResponseWithDefaults

`func NewPublicDomainListResponseWithDefaults() *PublicDomainListResponse`

NewPublicDomainListResponseWithDefaults instantiates a new PublicDomainListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PublicDomainListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PublicDomainListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PublicDomainListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *PublicDomainListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PublicDomainListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PublicDomainListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPublicDomainNames

`func (o *PublicDomainListResponse) GetPublicDomainNames() []PublicDomainListItem`

GetPublicDomainNames returns the PublicDomainNames field if non-nil, zero value otherwise.

### GetPublicDomainNamesOk

`func (o *PublicDomainListResponse) GetPublicDomainNamesOk() (*[]PublicDomainListItem, bool)`

GetPublicDomainNamesOk returns a tuple with the PublicDomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDomainNames

`func (o *PublicDomainListResponse) SetPublicDomainNames(v []PublicDomainListItem)`

SetPublicDomainNames sets PublicDomainNames field to given value.


### GetSize

`func (o *PublicDomainListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PublicDomainListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PublicDomainListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *PublicDomainListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PublicDomainListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PublicDomainListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PublicDomainListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *PublicDomainListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *PublicDomainListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


