# LogExportConfigListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]LogExportConfig**](LogExportConfig.md) | Log export config contents | 
**Count** | **int32** | Log export config count | 

## Methods

### NewLogExportConfigListResponse

`func NewLogExportConfigListResponse(contents []LogExportConfig, count int32, ) *LogExportConfigListResponse`

NewLogExportConfigListResponse instantiates a new LogExportConfigListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogExportConfigListResponseWithDefaults

`func NewLogExportConfigListResponseWithDefaults() *LogExportConfigListResponse`

NewLogExportConfigListResponseWithDefaults instantiates a new LogExportConfigListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *LogExportConfigListResponse) GetContents() []LogExportConfig`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *LogExportConfigListResponse) GetContentsOk() (*[]LogExportConfig, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *LogExportConfigListResponse) SetContents(v []LogExportConfig)`

SetContents sets Contents field to given value.


### GetCount

`func (o *LogExportConfigListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LogExportConfigListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LogExportConfigListResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


