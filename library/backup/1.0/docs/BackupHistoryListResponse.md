# BackupHistoryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]BackupHistoryResponse**](BackupHistoryResponse.md) |  | 
**Count** | **int32** | Count | 

## Methods

### NewBackupHistoryListResponse

`func NewBackupHistoryListResponse(contents []BackupHistoryResponse, count int32, ) *BackupHistoryListResponse`

NewBackupHistoryListResponse instantiates a new BackupHistoryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupHistoryListResponseWithDefaults

`func NewBackupHistoryListResponseWithDefaults() *BackupHistoryListResponse`

NewBackupHistoryListResponseWithDefaults instantiates a new BackupHistoryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *BackupHistoryListResponse) GetContents() []BackupHistoryResponse`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BackupHistoryListResponse) GetContentsOk() (*[]BackupHistoryResponse, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BackupHistoryListResponse) SetContents(v []BackupHistoryResponse)`

SetContents sets Contents field to given value.


### SetContentsNil

`func (o *BackupHistoryListResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *BackupHistoryListResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCount

`func (o *BackupHistoryListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BackupHistoryListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BackupHistoryListResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


