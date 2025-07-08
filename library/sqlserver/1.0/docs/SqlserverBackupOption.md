# SqlserverBackupOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveFrequencyMinute** | **string** | Backup starting time (minute) | 
**FullBackupDayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) | Full backup day of week | [optional] [default to DAYOFWEEK_SUN]
**RetentionPeriodDay** | Pointer to **NullableString** |  | [optional] 
**StartingTimeHour** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSqlserverBackupOption

`func NewSqlserverBackupOption(archiveFrequencyMinute string, ) *SqlserverBackupOption`

NewSqlserverBackupOption instantiates a new SqlserverBackupOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverBackupOptionWithDefaults

`func NewSqlserverBackupOptionWithDefaults() *SqlserverBackupOption`

NewSqlserverBackupOptionWithDefaults instantiates a new SqlserverBackupOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveFrequencyMinute

`func (o *SqlserverBackupOption) GetArchiveFrequencyMinute() string`

GetArchiveFrequencyMinute returns the ArchiveFrequencyMinute field if non-nil, zero value otherwise.

### GetArchiveFrequencyMinuteOk

`func (o *SqlserverBackupOption) GetArchiveFrequencyMinuteOk() (*string, bool)`

GetArchiveFrequencyMinuteOk returns a tuple with the ArchiveFrequencyMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFrequencyMinute

`func (o *SqlserverBackupOption) SetArchiveFrequencyMinute(v string)`

SetArchiveFrequencyMinute sets ArchiveFrequencyMinute field to given value.


### GetFullBackupDayOfWeek

`func (o *SqlserverBackupOption) GetFullBackupDayOfWeek() DayOfWeek`

GetFullBackupDayOfWeek returns the FullBackupDayOfWeek field if non-nil, zero value otherwise.

### GetFullBackupDayOfWeekOk

`func (o *SqlserverBackupOption) GetFullBackupDayOfWeekOk() (*DayOfWeek, bool)`

GetFullBackupDayOfWeekOk returns a tuple with the FullBackupDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupDayOfWeek

`func (o *SqlserverBackupOption) SetFullBackupDayOfWeek(v DayOfWeek)`

SetFullBackupDayOfWeek sets FullBackupDayOfWeek field to given value.

### HasFullBackupDayOfWeek

`func (o *SqlserverBackupOption) HasFullBackupDayOfWeek() bool`

HasFullBackupDayOfWeek returns a boolean if a field has been set.

### GetRetentionPeriodDay

`func (o *SqlserverBackupOption) GetRetentionPeriodDay() string`

GetRetentionPeriodDay returns the RetentionPeriodDay field if non-nil, zero value otherwise.

### GetRetentionPeriodDayOk

`func (o *SqlserverBackupOption) GetRetentionPeriodDayOk() (*string, bool)`

GetRetentionPeriodDayOk returns a tuple with the RetentionPeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDay

`func (o *SqlserverBackupOption) SetRetentionPeriodDay(v string)`

SetRetentionPeriodDay sets RetentionPeriodDay field to given value.

### HasRetentionPeriodDay

`func (o *SqlserverBackupOption) HasRetentionPeriodDay() bool`

HasRetentionPeriodDay returns a boolean if a field has been set.

### SetRetentionPeriodDayNil

`func (o *SqlserverBackupOption) SetRetentionPeriodDayNil(b bool)`

 SetRetentionPeriodDayNil sets the value for RetentionPeriodDay to be an explicit nil

### UnsetRetentionPeriodDay
`func (o *SqlserverBackupOption) UnsetRetentionPeriodDay()`

UnsetRetentionPeriodDay ensures that no value is present for RetentionPeriodDay, not even an explicit nil
### GetStartingTimeHour

`func (o *SqlserverBackupOption) GetStartingTimeHour() string`

GetStartingTimeHour returns the StartingTimeHour field if non-nil, zero value otherwise.

### GetStartingTimeHourOk

`func (o *SqlserverBackupOption) GetStartingTimeHourOk() (*string, bool)`

GetStartingTimeHourOk returns a tuple with the StartingTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingTimeHour

`func (o *SqlserverBackupOption) SetStartingTimeHour(v string)`

SetStartingTimeHour sets StartingTimeHour field to given value.

### HasStartingTimeHour

`func (o *SqlserverBackupOption) HasStartingTimeHour() bool`

HasStartingTimeHour returns a boolean if a field has been set.

### SetStartingTimeHourNil

`func (o *SqlserverBackupOption) SetStartingTimeHourNil(b bool)`

 SetStartingTimeHourNil sets the value for StartingTimeHour to be an explicit nil

### UnsetStartingTimeHour
`func (o *SqlserverBackupOption) UnsetStartingTimeHour()`

UnsetStartingTimeHour ensures that no value is present for StartingTimeHour, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


