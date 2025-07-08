# BackupRestoreHistoryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]BackupRestoreHistoryResponse**](BackupRestoreHistoryResponse.md) |  | 
**Count** | **int32** | Count | 

## Methods

### NewBackupRestoreHistoryListResponse

`func NewBackupRestoreHistoryListResponse(contents []BackupRestoreHistoryResponse, count int32, ) *BackupRestoreHistoryListResponse`

NewBackupRestoreHistoryListResponse instantiates a new BackupRestoreHistoryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreHistoryListResponseWithDefaults

`func NewBackupRestoreHistoryListResponseWithDefaults() *BackupRestoreHistoryListResponse`

NewBackupRestoreHistoryListResponseWithDefaults instantiates a new BackupRestoreHistoryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *BackupRestoreHistoryListResponse) GetContents() []BackupRestoreHistoryResponse`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BackupRestoreHistoryListResponse) GetContentsOk() (*[]BackupRestoreHistoryResponse, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BackupRestoreHistoryListResponse) SetContents(v []BackupRestoreHistoryResponse)`

SetContents sets Contents field to given value.


### SetContentsNil

`func (o *BackupRestoreHistoryListResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *BackupRestoreHistoryListResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCount

`func (o *BackupRestoreHistoryListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BackupRestoreHistoryListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BackupRestoreHistoryListResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


