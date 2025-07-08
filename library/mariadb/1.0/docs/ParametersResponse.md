# ParametersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]ParameterDTO**](ParameterDTO.md) | Parameter contents | 
**Count** | **int32** | Parameter count | 
**ModifiedAt** | **string** | Modification time | 

## Methods

### NewParametersResponse

`func NewParametersResponse(contents []ParameterDTO, count int32, modifiedAt string, ) *ParametersResponse`

NewParametersResponse instantiates a new ParametersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParametersResponseWithDefaults

`func NewParametersResponseWithDefaults() *ParametersResponse`

NewParametersResponseWithDefaults instantiates a new ParametersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *ParametersResponse) GetContents() []ParameterDTO`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ParametersResponse) GetContentsOk() (*[]ParameterDTO, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ParametersResponse) SetContents(v []ParameterDTO)`

SetContents sets Contents field to given value.


### GetCount

`func (o *ParametersResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ParametersResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ParametersResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetModifiedAt

`func (o *ParametersResponse) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ParametersResponse) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ParametersResponse) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


