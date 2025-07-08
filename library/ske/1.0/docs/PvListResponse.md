# PvListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**Pvs** | [**[]PvSummary**](PvSummary.md) |  | 

## Methods

### NewPvListResponse

`func NewPvListResponse(pvs []PvSummary, ) *PvListResponse`

NewPvListResponse instantiates a new PvListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvListResponseWithDefaults

`func NewPvListResponseWithDefaults() *PvListResponse`

NewPvListResponseWithDefaults instantiates a new PvListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PvListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PvListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PvListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PvListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PvListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PvListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *PvListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PvListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PvListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PvListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *PvListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *PvListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPvs

`func (o *PvListResponse) GetPvs() []PvSummary`

GetPvs returns the Pvs field if non-nil, zero value otherwise.

### GetPvsOk

`func (o *PvListResponse) GetPvsOk() (*[]PvSummary, bool)`

GetPvsOk returns a tuple with the Pvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvs

`func (o *PvListResponse) SetPvs(v []PvSummary)`

SetPvs sets Pvs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


