# BackupSettingExcludingArchiveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionPeriodDay** | **string** | Backup retention period (day) | 
**StartingTimeHour** | **string** | Backup starting time (hour) | 

## Methods

### NewBackupSettingExcludingArchiveRequest

`func NewBackupSettingExcludingArchiveRequest(retentionPeriodDay string, startingTimeHour string, ) *BackupSettingExcludingArchiveRequest`

NewBackupSettingExcludingArchiveRequest instantiates a new BackupSettingExcludingArchiveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSettingExcludingArchiveRequestWithDefaults

`func NewBackupSettingExcludingArchiveRequestWithDefaults() *BackupSettingExcludingArchiveRequest`

NewBackupSettingExcludingArchiveRequestWithDefaults instantiates a new BackupSettingExcludingArchiveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionPeriodDay

`func (o *BackupSettingExcludingArchiveRequest) GetRetentionPeriodDay() string`

GetRetentionPeriodDay returns the RetentionPeriodDay field if non-nil, zero value otherwise.

### GetRetentionPeriodDayOk

`func (o *BackupSettingExcludingArchiveRequest) GetRetentionPeriodDayOk() (*string, bool)`

GetRetentionPeriodDayOk returns a tuple with the RetentionPeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDay

`func (o *BackupSettingExcludingArchiveRequest) SetRetentionPeriodDay(v string)`

SetRetentionPeriodDay sets RetentionPeriodDay field to given value.


### GetStartingTimeHour

`func (o *BackupSettingExcludingArchiveRequest) GetStartingTimeHour() string`

GetStartingTimeHour returns the StartingTimeHour field if non-nil, zero value otherwise.

### GetStartingTimeHourOk

`func (o *BackupSettingExcludingArchiveRequest) GetStartingTimeHourOk() (*string, bool)`

GetStartingTimeHourOk returns a tuple with the StartingTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingTimeHour

`func (o *BackupSettingExcludingArchiveRequest) SetStartingTimeHour(v string)`

SetStartingTimeHour sets StartingTimeHour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


