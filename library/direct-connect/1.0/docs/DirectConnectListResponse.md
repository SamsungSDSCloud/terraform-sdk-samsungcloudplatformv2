# DirectConnectListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**DirectConnects** | [**[]DirectConnect**](DirectConnect.md) |  | 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewDirectConnectListResponse

`func NewDirectConnectListResponse(directConnects []DirectConnect, ) *DirectConnectListResponse`

NewDirectConnectListResponse instantiates a new DirectConnectListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectConnectListResponseWithDefaults

`func NewDirectConnectListResponseWithDefaults() *DirectConnectListResponse`

NewDirectConnectListResponseWithDefaults instantiates a new DirectConnectListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DirectConnectListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DirectConnectListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DirectConnectListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DirectConnectListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *DirectConnectListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *DirectConnectListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetDirectConnects

`func (o *DirectConnectListResponse) GetDirectConnects() []DirectConnect`

GetDirectConnects returns the DirectConnects field if non-nil, zero value otherwise.

### GetDirectConnectsOk

`func (o *DirectConnectListResponse) GetDirectConnectsOk() (*[]DirectConnect, bool)`

GetDirectConnectsOk returns a tuple with the DirectConnects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectConnects

`func (o *DirectConnectListResponse) SetDirectConnects(v []DirectConnect)`

SetDirectConnects sets DirectConnects field to given value.


### GetLinks

`func (o *DirectConnectListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DirectConnectListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DirectConnectListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DirectConnectListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *DirectConnectListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *DirectConnectListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


