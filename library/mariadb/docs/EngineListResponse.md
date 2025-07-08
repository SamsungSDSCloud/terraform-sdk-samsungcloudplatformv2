# EngineListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]Engine**](Engine.md) | DBaaS engine list | 

## Methods

### NewEngineListResponse

`func NewEngineListResponse(contents []Engine, ) *EngineListResponse`

NewEngineListResponse instantiates a new EngineListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineListResponseWithDefaults

`func NewEngineListResponseWithDefaults() *EngineListResponse`

NewEngineListResponseWithDefaults instantiates a new EngineListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *EngineListResponse) GetContents() []Engine`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *EngineListResponse) GetContentsOk() (*[]Engine, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *EngineListResponse) SetContents(v []Engine)`

SetContents sets Contents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


