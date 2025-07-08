# BackupHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupEndTime** | **NullableTime** |  | 
**BackupJobGuid** | **NullableString** |  | 
**BackupJobState** | **string** | Backup history state | 
**BackupStartTime** | **NullableTime** |  | 
**ScheduleId** | **string** | Schedule ID | 
**ScheduleName** | **string** | Schedule name | 

## Methods

### NewBackupHistoryResponse

`func NewBackupHistoryResponse(backupEndTime NullableTime, backupJobGuid NullableString, backupJobState string, backupStartTime NullableTime, scheduleId string, scheduleName string, ) *BackupHistoryResponse`

NewBackupHistoryResponse instantiates a new BackupHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupHistoryResponseWithDefaults

`func NewBackupHistoryResponseWithDefaults() *BackupHistoryResponse`

NewBackupHistoryResponseWithDefaults instantiates a new BackupHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupEndTime

`func (o *BackupHistoryResponse) GetBackupEndTime() time.Time`

GetBackupEndTime returns the BackupEndTime field if non-nil, zero value otherwise.

### GetBackupEndTimeOk

`func (o *BackupHistoryResponse) GetBackupEndTimeOk() (*time.Time, bool)`

GetBackupEndTimeOk returns a tuple with the BackupEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEndTime

`func (o *BackupHistoryResponse) SetBackupEndTime(v time.Time)`

SetBackupEndTime sets BackupEndTime field to given value.


### SetBackupEndTimeNil

`func (o *BackupHistoryResponse) SetBackupEndTimeNil(b bool)`

 SetBackupEndTimeNil sets the value for BackupEndTime to be an explicit nil

### UnsetBackupEndTime
`func (o *BackupHistoryResponse) UnsetBackupEndTime()`

UnsetBackupEndTime ensures that no value is present for BackupEndTime, not even an explicit nil
### GetBackupJobGuid

`func (o *BackupHistoryResponse) GetBackupJobGuid() string`

GetBackupJobGuid returns the BackupJobGuid field if non-nil, zero value otherwise.

### GetBackupJobGuidOk

`func (o *BackupHistoryResponse) GetBackupJobGuidOk() (*string, bool)`

GetBackupJobGuidOk returns a tuple with the BackupJobGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJobGuid

`func (o *BackupHistoryResponse) SetBackupJobGuid(v string)`

SetBackupJobGuid sets BackupJobGuid field to given value.


### SetBackupJobGuidNil

`func (o *BackupHistoryResponse) SetBackupJobGuidNil(b bool)`

 SetBackupJobGuidNil sets the value for BackupJobGuid to be an explicit nil

### UnsetBackupJobGuid
`func (o *BackupHistoryResponse) UnsetBackupJobGuid()`

UnsetBackupJobGuid ensures that no value is present for BackupJobGuid, not even an explicit nil
### GetBackupJobState

`func (o *BackupHistoryResponse) GetBackupJobState() string`

GetBackupJobState returns the BackupJobState field if non-nil, zero value otherwise.

### GetBackupJobStateOk

`func (o *BackupHistoryResponse) GetBackupJobStateOk() (*string, bool)`

GetBackupJobStateOk returns a tuple with the BackupJobState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJobState

`func (o *BackupHistoryResponse) SetBackupJobState(v string)`

SetBackupJobState sets BackupJobState field to given value.


### GetBackupStartTime

`func (o *BackupHistoryResponse) GetBackupStartTime() time.Time`

GetBackupStartTime returns the BackupStartTime field if non-nil, zero value otherwise.

### GetBackupStartTimeOk

`func (o *BackupHistoryResponse) GetBackupStartTimeOk() (*time.Time, bool)`

GetBackupStartTimeOk returns a tuple with the BackupStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStartTime

`func (o *BackupHistoryResponse) SetBackupStartTime(v time.Time)`

SetBackupStartTime sets BackupStartTime field to given value.


### SetBackupStartTimeNil

`func (o *BackupHistoryResponse) SetBackupStartTimeNil(b bool)`

 SetBackupStartTimeNil sets the value for BackupStartTime to be an explicit nil

### UnsetBackupStartTime
`func (o *BackupHistoryResponse) UnsetBackupStartTime()`

UnsetBackupStartTime ensures that no value is present for BackupStartTime, not even an explicit nil
### GetScheduleId

`func (o *BackupHistoryResponse) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BackupHistoryResponse) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BackupHistoryResponse) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.


### GetScheduleName

`func (o *BackupHistoryResponse) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *BackupHistoryResponse) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *BackupHistoryResponse) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


