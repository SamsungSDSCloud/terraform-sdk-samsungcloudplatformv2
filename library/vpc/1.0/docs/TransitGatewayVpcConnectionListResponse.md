# TransitGatewayVpcConnectionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 
**TransitGatewayVpcConnections** | [**[]TransitGatewayVpcConnection**](TransitGatewayVpcConnection.md) |  | 

## Methods

### NewTransitGatewayVpcConnectionListResponse

`func NewTransitGatewayVpcConnectionListResponse(count int32, page int32, size int32, transitGatewayVpcConnections []TransitGatewayVpcConnection, ) *TransitGatewayVpcConnectionListResponse`

NewTransitGatewayVpcConnectionListResponse instantiates a new TransitGatewayVpcConnectionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayVpcConnectionListResponseWithDefaults

`func NewTransitGatewayVpcConnectionListResponseWithDefaults() *TransitGatewayVpcConnectionListResponse`

NewTransitGatewayVpcConnectionListResponseWithDefaults instantiates a new TransitGatewayVpcConnectionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TransitGatewayVpcConnectionListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransitGatewayVpcConnectionListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransitGatewayVpcConnectionListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *TransitGatewayVpcConnectionListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TransitGatewayVpcConnectionListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TransitGatewayVpcConnectionListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *TransitGatewayVpcConnectionListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TransitGatewayVpcConnectionListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TransitGatewayVpcConnectionListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *TransitGatewayVpcConnectionListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TransitGatewayVpcConnectionListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TransitGatewayVpcConnectionListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TransitGatewayVpcConnectionListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *TransitGatewayVpcConnectionListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *TransitGatewayVpcConnectionListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetTransitGatewayVpcConnections

`func (o *TransitGatewayVpcConnectionListResponse) GetTransitGatewayVpcConnections() []TransitGatewayVpcConnection`

GetTransitGatewayVpcConnections returns the TransitGatewayVpcConnections field if non-nil, zero value otherwise.

### GetTransitGatewayVpcConnectionsOk

`func (o *TransitGatewayVpcConnectionListResponse) GetTransitGatewayVpcConnectionsOk() (*[]TransitGatewayVpcConnection, bool)`

GetTransitGatewayVpcConnectionsOk returns a tuple with the TransitGatewayVpcConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewayVpcConnections

`func (o *TransitGatewayVpcConnectionListResponse) SetTransitGatewayVpcConnections(v []TransitGatewayVpcConnection)`

SetTransitGatewayVpcConnections sets TransitGatewayVpcConnections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


