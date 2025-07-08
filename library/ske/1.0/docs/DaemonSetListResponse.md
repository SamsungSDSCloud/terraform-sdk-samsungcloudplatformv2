# DaemonSetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**DaemonSets** | [**[]DaemonSetSummary**](DaemonSetSummary.md) |  | 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewDaemonSetListResponse

`func NewDaemonSetListResponse(daemonSets []DaemonSetSummary, ) *DaemonSetListResponse`

NewDaemonSetListResponse instantiates a new DaemonSetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaemonSetListResponseWithDefaults

`func NewDaemonSetListResponseWithDefaults() *DaemonSetListResponse`

NewDaemonSetListResponseWithDefaults instantiates a new DaemonSetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DaemonSetListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DaemonSetListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DaemonSetListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DaemonSetListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *DaemonSetListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *DaemonSetListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetDaemonSets

`func (o *DaemonSetListResponse) GetDaemonSets() []DaemonSetSummary`

GetDaemonSets returns the DaemonSets field if non-nil, zero value otherwise.

### GetDaemonSetsOk

`func (o *DaemonSetListResponse) GetDaemonSetsOk() (*[]DaemonSetSummary, bool)`

GetDaemonSetsOk returns a tuple with the DaemonSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSets

`func (o *DaemonSetListResponse) SetDaemonSets(v []DaemonSetSummary)`

SetDaemonSets sets DaemonSets field to given value.


### GetLinks

`func (o *DaemonSetListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaemonSetListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaemonSetListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DaemonSetListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *DaemonSetListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *DaemonSetListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


