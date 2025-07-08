# ListAccessKeyRemoveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | Access key id list | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Marker** | Pointer to **NullableString** |  | [optional] 
**Sort** | Pointer to **NullableString** |  | [optional] 
**WithCount** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListAccessKeyRemoveRequest

`func NewListAccessKeyRemoveRequest() *ListAccessKeyRemoveRequest`

NewListAccessKeyRemoveRequest instantiates a new ListAccessKeyRemoveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccessKeyRemoveRequestWithDefaults

`func NewListAccessKeyRemoveRequestWithDefaults() *ListAccessKeyRemoveRequest`

NewListAccessKeyRemoveRequestWithDefaults instantiates a new ListAccessKeyRemoveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListAccessKeyRemoveRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListAccessKeyRemoveRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListAccessKeyRemoveRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListAccessKeyRemoveRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetLimit

`func (o *ListAccessKeyRemoveRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListAccessKeyRemoveRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListAccessKeyRemoveRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListAccessKeyRemoveRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ListAccessKeyRemoveRequest) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ListAccessKeyRemoveRequest) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetMarker

`func (o *ListAccessKeyRemoveRequest) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *ListAccessKeyRemoveRequest) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *ListAccessKeyRemoveRequest) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *ListAccessKeyRemoveRequest) HasMarker() bool`

HasMarker returns a boolean if a field has been set.

### SetMarkerNil

`func (o *ListAccessKeyRemoveRequest) SetMarkerNil(b bool)`

 SetMarkerNil sets the value for Marker to be an explicit nil

### UnsetMarker
`func (o *ListAccessKeyRemoveRequest) UnsetMarker()`

UnsetMarker ensures that no value is present for Marker, not even an explicit nil
### GetSort

`func (o *ListAccessKeyRemoveRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ListAccessKeyRemoveRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ListAccessKeyRemoveRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ListAccessKeyRemoveRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *ListAccessKeyRemoveRequest) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *ListAccessKeyRemoveRequest) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetWithCount

`func (o *ListAccessKeyRemoveRequest) GetWithCount() string`

GetWithCount returns the WithCount field if non-nil, zero value otherwise.

### GetWithCountOk

`func (o *ListAccessKeyRemoveRequest) GetWithCountOk() (*string, bool)`

GetWithCountOk returns a tuple with the WithCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCount

`func (o *ListAccessKeyRemoveRequest) SetWithCount(v string)`

SetWithCount sets WithCount field to given value.

### HasWithCount

`func (o *ListAccessKeyRemoveRequest) HasWithCount() bool`

HasWithCount returns a boolean if a field has been set.

### SetWithCountNil

`func (o *ListAccessKeyRemoveRequest) SetWithCountNil(b bool)`

 SetWithCountNil sets the value for WithCount to be an explicit nil

### UnsetWithCount
`func (o *ListAccessKeyRemoveRequest) UnsetWithCount()`

UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


