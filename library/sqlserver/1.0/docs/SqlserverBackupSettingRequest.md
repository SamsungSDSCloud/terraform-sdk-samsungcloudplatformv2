# SqlserverBackupSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveFrequencyMinute** | **string** | Backup starting time (minute) | 
**FullBackupDayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) | Full backup day of week | [optional] [default to DAYOFWEEK_SUN]
**RetentionPeriodDay** | **string** | Backup retention period (day) | 
**StartingTimeHour** | **string** | Backup starting time (hour) | 

## Methods

### NewSqlserverBackupSettingRequest

`func NewSqlserverBackupSettingRequest(archiveFrequencyMinute string, retentionPeriodDay string, startingTimeHour string, ) *SqlserverBackupSettingRequest`

NewSqlserverBackupSettingRequest instantiates a new SqlserverBackupSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverBackupSettingRequestWithDefaults

`func NewSqlserverBackupSettingRequestWithDefaults() *SqlserverBackupSettingRequest`

NewSqlserverBackupSettingRequestWithDefaults instantiates a new SqlserverBackupSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveFrequencyMinute

`func (o *SqlserverBackupSettingRequest) GetArchiveFrequencyMinute() string`

GetArchiveFrequencyMinute returns the ArchiveFrequencyMinute field if non-nil, zero value otherwise.

### GetArchiveFrequencyMinuteOk

`func (o *SqlserverBackupSettingRequest) GetArchiveFrequencyMinuteOk() (*string, bool)`

GetArchiveFrequencyMinuteOk returns a tuple with the ArchiveFrequencyMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFrequencyMinute

`func (o *SqlserverBackupSettingRequest) SetArchiveFrequencyMinute(v string)`

SetArchiveFrequencyMinute sets ArchiveFrequencyMinute field to given value.


### GetFullBackupDayOfWeek

`func (o *SqlserverBackupSettingRequest) GetFullBackupDayOfWeek() DayOfWeek`

GetFullBackupDayOfWeek returns the FullBackupDayOfWeek field if non-nil, zero value otherwise.

### GetFullBackupDayOfWeekOk

`func (o *SqlserverBackupSettingRequest) GetFullBackupDayOfWeekOk() (*DayOfWeek, bool)`

GetFullBackupDayOfWeekOk returns a tuple with the FullBackupDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupDayOfWeek

`func (o *SqlserverBackupSettingRequest) SetFullBackupDayOfWeek(v DayOfWeek)`

SetFullBackupDayOfWeek sets FullBackupDayOfWeek field to given value.

### HasFullBackupDayOfWeek

`func (o *SqlserverBackupSettingRequest) HasFullBackupDayOfWeek() bool`

HasFullBackupDayOfWeek returns a boolean if a field has been set.

### GetRetentionPeriodDay

`func (o *SqlserverBackupSettingRequest) GetRetentionPeriodDay() string`

GetRetentionPeriodDay returns the RetentionPeriodDay field if non-nil, zero value otherwise.

### GetRetentionPeriodDayOk

`func (o *SqlserverBackupSettingRequest) GetRetentionPeriodDayOk() (*string, bool)`

GetRetentionPeriodDayOk returns a tuple with the RetentionPeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDay

`func (o *SqlserverBackupSettingRequest) SetRetentionPeriodDay(v string)`

SetRetentionPeriodDay sets RetentionPeriodDay field to given value.


### GetStartingTimeHour

`func (o *SqlserverBackupSettingRequest) GetStartingTimeHour() string`

GetStartingTimeHour returns the StartingTimeHour field if non-nil, zero value otherwise.

### GetStartingTimeHourOk

`func (o *SqlserverBackupSettingRequest) GetStartingTimeHourOk() (*string, bool)`

GetStartingTimeHourOk returns a tuple with the StartingTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingTimeHour

`func (o *SqlserverBackupSettingRequest) SetStartingTimeHour(v string)`

SetStartingTimeHour sets StartingTimeHour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


