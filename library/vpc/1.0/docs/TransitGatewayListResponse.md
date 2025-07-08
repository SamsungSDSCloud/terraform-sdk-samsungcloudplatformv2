# TransitGatewayListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 
**TransitGateways** | [**[]TransitGateway**](TransitGateway.md) |  | 

## Methods

### NewTransitGatewayListResponse

`func NewTransitGatewayListResponse(count int32, page int32, size int32, transitGateways []TransitGateway, ) *TransitGatewayListResponse`

NewTransitGatewayListResponse instantiates a new TransitGatewayListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayListResponseWithDefaults

`func NewTransitGatewayListResponseWithDefaults() *TransitGatewayListResponse`

NewTransitGatewayListResponseWithDefaults instantiates a new TransitGatewayListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TransitGatewayListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransitGatewayListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransitGatewayListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *TransitGatewayListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TransitGatewayListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TransitGatewayListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *TransitGatewayListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TransitGatewayListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TransitGatewayListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *TransitGatewayListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TransitGatewayListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TransitGatewayListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TransitGatewayListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *TransitGatewayListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *TransitGatewayListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetTransitGateways

`func (o *TransitGatewayListResponse) GetTransitGateways() []TransitGateway`

GetTransitGateways returns the TransitGateways field if non-nil, zero value otherwise.

### GetTransitGatewaysOk

`func (o *TransitGatewayListResponse) GetTransitGatewaysOk() (*[]TransitGateway, bool)`

GetTransitGatewaysOk returns a tuple with the TransitGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGateways

`func (o *TransitGatewayListResponse) SetTransitGateways(v []TransitGateway)`

SetTransitGateways sets TransitGateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


