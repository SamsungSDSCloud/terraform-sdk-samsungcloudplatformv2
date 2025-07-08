# BackupRestoreHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **string** | Backup ID | 
**BackupStartTime** | **time.Time** | Backup time | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**RestoreEndTime** | **NullableTime** |  | 
**RestoreServerId** | **NullableString** |  | 
**RestoreServerName** | **string** | Restore server name | 
**RestoreStartTime** | **time.Time** | Restore start time | 
**RestoreState** | [**BackupRestoreState**](BackupRestoreState.md) | Restore state | 
**ScheduleName** | **string** | Schedule name | 

## Methods

### NewBackupRestoreHistoryResponse

`func NewBackupRestoreHistoryResponse(backupId string, backupStartTime time.Time, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, restoreEndTime NullableTime, restoreServerId NullableString, restoreServerName string, restoreStartTime time.Time, restoreState BackupRestoreState, scheduleName string, ) *BackupRestoreHistoryResponse`

NewBackupRestoreHistoryResponse instantiates a new BackupRestoreHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreHistoryResponseWithDefaults

`func NewBackupRestoreHistoryResponseWithDefaults() *BackupRestoreHistoryResponse`

NewBackupRestoreHistoryResponseWithDefaults instantiates a new BackupRestoreHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *BackupRestoreHistoryResponse) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *BackupRestoreHistoryResponse) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *BackupRestoreHistoryResponse) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetBackupStartTime

`func (o *BackupRestoreHistoryResponse) GetBackupStartTime() time.Time`

GetBackupStartTime returns the BackupStartTime field if non-nil, zero value otherwise.

### GetBackupStartTimeOk

`func (o *BackupRestoreHistoryResponse) GetBackupStartTimeOk() (*time.Time, bool)`

GetBackupStartTimeOk returns a tuple with the BackupStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStartTime

`func (o *BackupRestoreHistoryResponse) SetBackupStartTime(v time.Time)`

SetBackupStartTime sets BackupStartTime field to given value.


### GetCreatedAt

`func (o *BackupRestoreHistoryResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupRestoreHistoryResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupRestoreHistoryResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *BackupRestoreHistoryResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BackupRestoreHistoryResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BackupRestoreHistoryResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *BackupRestoreHistoryResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BackupRestoreHistoryResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BackupRestoreHistoryResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *BackupRestoreHistoryResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BackupRestoreHistoryResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BackupRestoreHistoryResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetRestoreEndTime

`func (o *BackupRestoreHistoryResponse) GetRestoreEndTime() time.Time`

GetRestoreEndTime returns the RestoreEndTime field if non-nil, zero value otherwise.

### GetRestoreEndTimeOk

`func (o *BackupRestoreHistoryResponse) GetRestoreEndTimeOk() (*time.Time, bool)`

GetRestoreEndTimeOk returns a tuple with the RestoreEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEndTime

`func (o *BackupRestoreHistoryResponse) SetRestoreEndTime(v time.Time)`

SetRestoreEndTime sets RestoreEndTime field to given value.


### SetRestoreEndTimeNil

`func (o *BackupRestoreHistoryResponse) SetRestoreEndTimeNil(b bool)`

 SetRestoreEndTimeNil sets the value for RestoreEndTime to be an explicit nil

### UnsetRestoreEndTime
`func (o *BackupRestoreHistoryResponse) UnsetRestoreEndTime()`

UnsetRestoreEndTime ensures that no value is present for RestoreEndTime, not even an explicit nil
### GetRestoreServerId

`func (o *BackupRestoreHistoryResponse) GetRestoreServerId() string`

GetRestoreServerId returns the RestoreServerId field if non-nil, zero value otherwise.

### GetRestoreServerIdOk

`func (o *BackupRestoreHistoryResponse) GetRestoreServerIdOk() (*string, bool)`

GetRestoreServerIdOk returns a tuple with the RestoreServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreServerId

`func (o *BackupRestoreHistoryResponse) SetRestoreServerId(v string)`

SetRestoreServerId sets RestoreServerId field to given value.


### SetRestoreServerIdNil

`func (o *BackupRestoreHistoryResponse) SetRestoreServerIdNil(b bool)`

 SetRestoreServerIdNil sets the value for RestoreServerId to be an explicit nil

### UnsetRestoreServerId
`func (o *BackupRestoreHistoryResponse) UnsetRestoreServerId()`

UnsetRestoreServerId ensures that no value is present for RestoreServerId, not even an explicit nil
### GetRestoreServerName

`func (o *BackupRestoreHistoryResponse) GetRestoreServerName() string`

GetRestoreServerName returns the RestoreServerName field if non-nil, zero value otherwise.

### GetRestoreServerNameOk

`func (o *BackupRestoreHistoryResponse) GetRestoreServerNameOk() (*string, bool)`

GetRestoreServerNameOk returns a tuple with the RestoreServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreServerName

`func (o *BackupRestoreHistoryResponse) SetRestoreServerName(v string)`

SetRestoreServerName sets RestoreServerName field to given value.


### GetRestoreStartTime

`func (o *BackupRestoreHistoryResponse) GetRestoreStartTime() time.Time`

GetRestoreStartTime returns the RestoreStartTime field if non-nil, zero value otherwise.

### GetRestoreStartTimeOk

`func (o *BackupRestoreHistoryResponse) GetRestoreStartTimeOk() (*time.Time, bool)`

GetRestoreStartTimeOk returns a tuple with the RestoreStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreStartTime

`func (o *BackupRestoreHistoryResponse) SetRestoreStartTime(v time.Time)`

SetRestoreStartTime sets RestoreStartTime field to given value.


### GetRestoreState

`func (o *BackupRestoreHistoryResponse) GetRestoreState() BackupRestoreState`

GetRestoreState returns the RestoreState field if non-nil, zero value otherwise.

### GetRestoreStateOk

`func (o *BackupRestoreHistoryResponse) GetRestoreStateOk() (*BackupRestoreState, bool)`

GetRestoreStateOk returns a tuple with the RestoreState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreState

`func (o *BackupRestoreHistoryResponse) SetRestoreState(v BackupRestoreState)`

SetRestoreState sets RestoreState field to given value.


### GetScheduleName

`func (o *BackupRestoreHistoryResponse) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *BackupRestoreHistoryResponse) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *BackupRestoreHistoryResponse) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


