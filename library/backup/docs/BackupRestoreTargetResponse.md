# BackupRestoreTargetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupStartTime** | **NullableTime** |  | 
**IsRestoreAvailable** | **bool** | Whether restore available | 
**RestoreTargetId** | **string** | Restore target ID | 
**RetentionPeriod** | **string** | Backup retention period | 
**ScheduleId** | **string** | Schedule ID | 
**ScheduleName** | **string** | Schedule name | 
**UsageGb** | **float32** | Backup size(GB) | 

## Methods

### NewBackupRestoreTargetResponse

`func NewBackupRestoreTargetResponse(backupStartTime NullableTime, isRestoreAvailable bool, restoreTargetId string, retentionPeriod string, scheduleId string, scheduleName string, usageGb float32, ) *BackupRestoreTargetResponse`

NewBackupRestoreTargetResponse instantiates a new BackupRestoreTargetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreTargetResponseWithDefaults

`func NewBackupRestoreTargetResponseWithDefaults() *BackupRestoreTargetResponse`

NewBackupRestoreTargetResponseWithDefaults instantiates a new BackupRestoreTargetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupStartTime

`func (o *BackupRestoreTargetResponse) GetBackupStartTime() time.Time`

GetBackupStartTime returns the BackupStartTime field if non-nil, zero value otherwise.

### GetBackupStartTimeOk

`func (o *BackupRestoreTargetResponse) GetBackupStartTimeOk() (*time.Time, bool)`

GetBackupStartTimeOk returns a tuple with the BackupStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupStartTime

`func (o *BackupRestoreTargetResponse) SetBackupStartTime(v time.Time)`

SetBackupStartTime sets BackupStartTime field to given value.


### SetBackupStartTimeNil

`func (o *BackupRestoreTargetResponse) SetBackupStartTimeNil(b bool)`

 SetBackupStartTimeNil sets the value for BackupStartTime to be an explicit nil

### UnsetBackupStartTime
`func (o *BackupRestoreTargetResponse) UnsetBackupStartTime()`

UnsetBackupStartTime ensures that no value is present for BackupStartTime, not even an explicit nil
### GetIsRestoreAvailable

`func (o *BackupRestoreTargetResponse) GetIsRestoreAvailable() bool`

GetIsRestoreAvailable returns the IsRestoreAvailable field if non-nil, zero value otherwise.

### GetIsRestoreAvailableOk

`func (o *BackupRestoreTargetResponse) GetIsRestoreAvailableOk() (*bool, bool)`

GetIsRestoreAvailableOk returns a tuple with the IsRestoreAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestoreAvailable

`func (o *BackupRestoreTargetResponse) SetIsRestoreAvailable(v bool)`

SetIsRestoreAvailable sets IsRestoreAvailable field to given value.


### GetRestoreTargetId

`func (o *BackupRestoreTargetResponse) GetRestoreTargetId() string`

GetRestoreTargetId returns the RestoreTargetId field if non-nil, zero value otherwise.

### GetRestoreTargetIdOk

`func (o *BackupRestoreTargetResponse) GetRestoreTargetIdOk() (*string, bool)`

GetRestoreTargetIdOk returns a tuple with the RestoreTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTargetId

`func (o *BackupRestoreTargetResponse) SetRestoreTargetId(v string)`

SetRestoreTargetId sets RestoreTargetId field to given value.


### GetRetentionPeriod

`func (o *BackupRestoreTargetResponse) GetRetentionPeriod() string`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *BackupRestoreTargetResponse) GetRetentionPeriodOk() (*string, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *BackupRestoreTargetResponse) SetRetentionPeriod(v string)`

SetRetentionPeriod sets RetentionPeriod field to given value.


### GetScheduleId

`func (o *BackupRestoreTargetResponse) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BackupRestoreTargetResponse) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BackupRestoreTargetResponse) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.


### GetScheduleName

`func (o *BackupRestoreTargetResponse) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *BackupRestoreTargetResponse) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *BackupRestoreTargetResponse) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.


### GetUsageGb

`func (o *BackupRestoreTargetResponse) GetUsageGb() float32`

GetUsageGb returns the UsageGb field if non-nil, zero value otherwise.

### GetUsageGbOk

`func (o *BackupRestoreTargetResponse) GetUsageGbOk() (*float32, bool)`

GetUsageGbOk returns a tuple with the UsageGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageGb

`func (o *BackupRestoreTargetResponse) SetUsageGb(v float32)`

SetUsageGb sets UsageGb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


