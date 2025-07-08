# BackupAgentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]BackupAgentResponse**](BackupAgentResponse.md) |  | 
**Count** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBackupAgentListResponse

`func NewBackupAgentListResponse(contents []BackupAgentResponse, ) *BackupAgentListResponse`

NewBackupAgentListResponse instantiates a new BackupAgentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupAgentListResponseWithDefaults

`func NewBackupAgentListResponseWithDefaults() *BackupAgentListResponse`

NewBackupAgentListResponseWithDefaults instantiates a new BackupAgentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *BackupAgentListResponse) GetContents() []BackupAgentResponse`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BackupAgentListResponse) GetContentsOk() (*[]BackupAgentResponse, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BackupAgentListResponse) SetContents(v []BackupAgentResponse)`

SetContents sets Contents field to given value.


### SetContentsNil

`func (o *BackupAgentListResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *BackupAgentListResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCount

`func (o *BackupAgentListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BackupAgentListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BackupAgentListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BackupAgentListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *BackupAgentListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *BackupAgentListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


