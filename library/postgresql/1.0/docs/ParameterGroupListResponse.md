# ParameterGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]ParameterGroup**](ParameterGroup.md) | Parameter group list | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewParameterGroupListResponse

`func NewParameterGroupListResponse(contents []ParameterGroup, ) *ParameterGroupListResponse`

NewParameterGroupListResponse instantiates a new ParameterGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterGroupListResponseWithDefaults

`func NewParameterGroupListResponseWithDefaults() *ParameterGroupListResponse`

NewParameterGroupListResponseWithDefaults instantiates a new ParameterGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *ParameterGroupListResponse) GetContents() []ParameterGroup`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ParameterGroupListResponse) GetContentsOk() (*[]ParameterGroup, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ParameterGroupListResponse) SetContents(v []ParameterGroup)`

SetContents sets Contents field to given value.


### GetSort

`func (o *ParameterGroupListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ParameterGroupListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ParameterGroupListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ParameterGroupListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *ParameterGroupListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *ParameterGroupListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


