# VpnTunnelListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 
**VpnTunnels** | [**[]VpnTunnel**](VpnTunnel.md) |  | 

## Methods

### NewVpnTunnelListResponse

`func NewVpnTunnelListResponse(count int32, page int32, size int32, vpnTunnels []VpnTunnel, ) *VpnTunnelListResponse`

NewVpnTunnelListResponse instantiates a new VpnTunnelListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelListResponseWithDefaults

`func NewVpnTunnelListResponseWithDefaults() *VpnTunnelListResponse`

NewVpnTunnelListResponseWithDefaults instantiates a new VpnTunnelListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VpnTunnelListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VpnTunnelListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VpnTunnelListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *VpnTunnelListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *VpnTunnelListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *VpnTunnelListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *VpnTunnelListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VpnTunnelListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VpnTunnelListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *VpnTunnelListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *VpnTunnelListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *VpnTunnelListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *VpnTunnelListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *VpnTunnelListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *VpnTunnelListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetVpnTunnels

`func (o *VpnTunnelListResponse) GetVpnTunnels() []VpnTunnel`

GetVpnTunnels returns the VpnTunnels field if non-nil, zero value otherwise.

### GetVpnTunnelsOk

`func (o *VpnTunnelListResponse) GetVpnTunnelsOk() (*[]VpnTunnel, bool)`

GetVpnTunnelsOk returns a tuple with the VpnTunnels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnTunnels

`func (o *VpnTunnelListResponse) SetVpnTunnels(v []VpnTunnel)`

SetVpnTunnels sets VpnTunnels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


