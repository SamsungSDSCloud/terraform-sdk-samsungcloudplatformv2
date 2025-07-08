# NatGatewayListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**NatGateways** | [**[]NatGateway**](NatGateway.md) |  | 

## Methods

### NewNatGatewayListResponse

`func NewNatGatewayListResponse(natGateways []NatGateway, ) *NatGatewayListResponse`

NewNatGatewayListResponse instantiates a new NatGatewayListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatGatewayListResponseWithDefaults

`func NewNatGatewayListResponseWithDefaults() *NatGatewayListResponse`

NewNatGatewayListResponseWithDefaults instantiates a new NatGatewayListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *NatGatewayListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NatGatewayListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NatGatewayListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NatGatewayListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *NatGatewayListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *NatGatewayListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *NatGatewayListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NatGatewayListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NatGatewayListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NatGatewayListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *NatGatewayListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *NatGatewayListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetNatGateways

`func (o *NatGatewayListResponse) GetNatGateways() []NatGateway`

GetNatGateways returns the NatGateways field if non-nil, zero value otherwise.

### GetNatGatewaysOk

`func (o *NatGatewayListResponse) GetNatGatewaysOk() (*[]NatGateway, bool)`

GetNatGatewaysOk returns a tuple with the NatGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatGateways

`func (o *NatGatewayListResponse) SetNatGateways(v []NatGateway)`

SetNatGateways sets NatGateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


