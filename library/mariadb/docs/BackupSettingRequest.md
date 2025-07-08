# BackupSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveFrequencyMinute** | **string** | Backup starting time (minute) | 
**RetentionPeriodDay** | **string** | Backup retention period (day) | 
**StartingTimeHour** | **string** | Backup starting time (hour) | 

## Methods

### NewBackupSettingRequest

`func NewBackupSettingRequest(archiveFrequencyMinute string, retentionPeriodDay string, startingTimeHour string, ) *BackupSettingRequest`

NewBackupSettingRequest instantiates a new BackupSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSettingRequestWithDefaults

`func NewBackupSettingRequestWithDefaults() *BackupSettingRequest`

NewBackupSettingRequestWithDefaults instantiates a new BackupSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveFrequencyMinute

`func (o *BackupSettingRequest) GetArchiveFrequencyMinute() string`

GetArchiveFrequencyMinute returns the ArchiveFrequencyMinute field if non-nil, zero value otherwise.

### GetArchiveFrequencyMinuteOk

`func (o *BackupSettingRequest) GetArchiveFrequencyMinuteOk() (*string, bool)`

GetArchiveFrequencyMinuteOk returns a tuple with the ArchiveFrequencyMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFrequencyMinute

`func (o *BackupSettingRequest) SetArchiveFrequencyMinute(v string)`

SetArchiveFrequencyMinute sets ArchiveFrequencyMinute field to given value.


### GetRetentionPeriodDay

`func (o *BackupSettingRequest) GetRetentionPeriodDay() string`

GetRetentionPeriodDay returns the RetentionPeriodDay field if non-nil, zero value otherwise.

### GetRetentionPeriodDayOk

`func (o *BackupSettingRequest) GetRetentionPeriodDayOk() (*string, bool)`

GetRetentionPeriodDayOk returns a tuple with the RetentionPeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDay

`func (o *BackupSettingRequest) SetRetentionPeriodDay(v string)`

SetRetentionPeriodDay sets RetentionPeriodDay field to given value.


### GetStartingTimeHour

`func (o *BackupSettingRequest) GetStartingTimeHour() string`

GetStartingTimeHour returns the StartingTimeHour field if non-nil, zero value otherwise.

### GetStartingTimeHourOk

`func (o *BackupSettingRequest) GetStartingTimeHourOk() (*string, bool)`

GetStartingTimeHourOk returns a tuple with the StartingTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingTimeHour

`func (o *BackupSettingRequest) SetStartingTimeHour(v string)`

SetStartingTimeHour sets StartingTimeHour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


