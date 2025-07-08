# VpcPeeringListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 
**VpcPeerings** | [**[]VpcPeering**](VpcPeering.md) |  | 

## Methods

### NewVpcPeeringListResponse

`func NewVpcPeeringListResponse(count int32, page int32, size int32, vpcPeerings []VpcPeering, ) *VpcPeeringListResponse`

NewVpcPeeringListResponse instantiates a new VpcPeeringListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcPeeringListResponseWithDefaults

`func NewVpcPeeringListResponseWithDefaults() *VpcPeeringListResponse`

NewVpcPeeringListResponseWithDefaults instantiates a new VpcPeeringListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VpcPeeringListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VpcPeeringListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VpcPeeringListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *VpcPeeringListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *VpcPeeringListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *VpcPeeringListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *VpcPeeringListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VpcPeeringListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VpcPeeringListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *VpcPeeringListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *VpcPeeringListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *VpcPeeringListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *VpcPeeringListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *VpcPeeringListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *VpcPeeringListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetVpcPeerings

`func (o *VpcPeeringListResponse) GetVpcPeerings() []VpcPeering`

GetVpcPeerings returns the VpcPeerings field if non-nil, zero value otherwise.

### GetVpcPeeringsOk

`func (o *VpcPeeringListResponse) GetVpcPeeringsOk() (*[]VpcPeering, bool)`

GetVpcPeeringsOk returns a tuple with the VpcPeerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeerings

`func (o *VpcPeeringListResponse) SetVpcPeerings(v []VpcPeering)`

SetVpcPeerings sets VpcPeerings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


