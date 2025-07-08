# IngressClassListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**IngressClasses** | [**[]IngressClassSummary**](IngressClassSummary.md) |  | 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewIngressClassListResponse

`func NewIngressClassListResponse(ingressClasses []IngressClassSummary, ) *IngressClassListResponse`

NewIngressClassListResponse instantiates a new IngressClassListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressClassListResponseWithDefaults

`func NewIngressClassListResponseWithDefaults() *IngressClassListResponse`

NewIngressClassListResponseWithDefaults instantiates a new IngressClassListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IngressClassListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IngressClassListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IngressClassListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IngressClassListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *IngressClassListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *IngressClassListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetIngressClasses

`func (o *IngressClassListResponse) GetIngressClasses() []IngressClassSummary`

GetIngressClasses returns the IngressClasses field if non-nil, zero value otherwise.

### GetIngressClassesOk

`func (o *IngressClassListResponse) GetIngressClassesOk() (*[]IngressClassSummary, bool)`

GetIngressClassesOk returns a tuple with the IngressClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressClasses

`func (o *IngressClassListResponse) SetIngressClasses(v []IngressClassSummary)`

SetIngressClasses sets IngressClasses field to given value.


### GetLinks

`func (o *IngressClassListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IngressClassListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IngressClassListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IngressClassListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *IngressClassListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IngressClassListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


