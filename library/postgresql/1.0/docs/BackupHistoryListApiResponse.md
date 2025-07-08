# BackupHistoryListApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]BackupHistoryApiItem**](BackupHistoryApiItem.md) | backup list information | 
**TotalCount** | **int32** | total number of backup history | 

## Methods

### NewBackupHistoryListApiResponse

`func NewBackupHistoryListApiResponse(contents []BackupHistoryApiItem, totalCount int32, ) *BackupHistoryListApiResponse`

NewBackupHistoryListApiResponse instantiates a new BackupHistoryListApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupHistoryListApiResponseWithDefaults

`func NewBackupHistoryListApiResponseWithDefaults() *BackupHistoryListApiResponse`

NewBackupHistoryListApiResponseWithDefaults instantiates a new BackupHistoryListApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *BackupHistoryListApiResponse) GetContents() []BackupHistoryApiItem`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BackupHistoryListApiResponse) GetContentsOk() (*[]BackupHistoryApiItem, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BackupHistoryListApiResponse) SetContents(v []BackupHistoryApiItem)`

SetContents sets Contents field to given value.


### GetTotalCount

`func (o *BackupHistoryListApiResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BackupHistoryListApiResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BackupHistoryListApiResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


