# PrivateNatIpListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**PrivateNatIps** | [**[]PrivateNatIp**](PrivateNatIp.md) |  | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPrivateNatIpListResponse

`func NewPrivateNatIpListResponse(count int32, page int32, privateNatIps []PrivateNatIp, size int32, ) *PrivateNatIpListResponse`

NewPrivateNatIpListResponse instantiates a new PrivateNatIpListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNatIpListResponseWithDefaults

`func NewPrivateNatIpListResponseWithDefaults() *PrivateNatIpListResponse`

NewPrivateNatIpListResponseWithDefaults instantiates a new PrivateNatIpListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PrivateNatIpListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PrivateNatIpListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PrivateNatIpListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *PrivateNatIpListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PrivateNatIpListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PrivateNatIpListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPrivateNatIps

`func (o *PrivateNatIpListResponse) GetPrivateNatIps() []PrivateNatIp`

GetPrivateNatIps returns the PrivateNatIps field if non-nil, zero value otherwise.

### GetPrivateNatIpsOk

`func (o *PrivateNatIpListResponse) GetPrivateNatIpsOk() (*[]PrivateNatIp, bool)`

GetPrivateNatIpsOk returns a tuple with the PrivateNatIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNatIps

`func (o *PrivateNatIpListResponse) SetPrivateNatIps(v []PrivateNatIp)`

SetPrivateNatIps sets PrivateNatIps field to given value.


### GetSize

`func (o *PrivateNatIpListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PrivateNatIpListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PrivateNatIpListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *PrivateNatIpListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PrivateNatIpListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PrivateNatIpListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PrivateNatIpListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *PrivateNatIpListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *PrivateNatIpListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


