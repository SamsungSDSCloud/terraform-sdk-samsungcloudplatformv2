# CommandListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]CommandItem**](CommandItem.md) | Command list information | 
**LastUpdatedAt** | **string** | Modification time | 
**TotalCount** | **int32** | Total number of commands | 

## Methods

### NewCommandListResponse

`func NewCommandListResponse(contents []CommandItem, lastUpdatedAt string, totalCount int32, ) *CommandListResponse`

NewCommandListResponse instantiates a new CommandListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandListResponseWithDefaults

`func NewCommandListResponseWithDefaults() *CommandListResponse`

NewCommandListResponseWithDefaults instantiates a new CommandListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *CommandListResponse) GetContents() []CommandItem`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *CommandListResponse) GetContentsOk() (*[]CommandItem, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *CommandListResponse) SetContents(v []CommandItem)`

SetContents sets Contents field to given value.


### GetLastUpdatedAt

`func (o *CommandListResponse) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *CommandListResponse) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *CommandListResponse) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.


### GetTotalCount

`func (o *CommandListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CommandListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CommandListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


