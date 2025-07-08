# IngressListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Ingresses** | [**[]IngressSummary**](IngressSummary.md) |  | 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewIngressListResponse

`func NewIngressListResponse(ingresses []IngressSummary, ) *IngressListResponse`

NewIngressListResponse instantiates a new IngressListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressListResponseWithDefaults

`func NewIngressListResponseWithDefaults() *IngressListResponse`

NewIngressListResponseWithDefaults instantiates a new IngressListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IngressListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IngressListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IngressListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IngressListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *IngressListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *IngressListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetIngresses

`func (o *IngressListResponse) GetIngresses() []IngressSummary`

GetIngresses returns the Ingresses field if non-nil, zero value otherwise.

### GetIngressesOk

`func (o *IngressListResponse) GetIngressesOk() (*[]IngressSummary, bool)`

GetIngressesOk returns a tuple with the Ingresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresses

`func (o *IngressListResponse) SetIngresses(v []IngressSummary)`

SetIngresses sets Ingresses field to given value.


### GetLinks

`func (o *IngressListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IngressListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IngressListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IngressListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *IngressListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IngressListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


