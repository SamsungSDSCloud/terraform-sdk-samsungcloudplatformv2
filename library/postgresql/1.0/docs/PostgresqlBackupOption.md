# PostgresqlBackupOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveFrequencyMinute** | Pointer to **NullableString** |  | [optional] 
**RetentionPeriodDay** | Pointer to **NullableString** |  | [optional] 
**StartingTimeHour** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPostgresqlBackupOption

`func NewPostgresqlBackupOption() *PostgresqlBackupOption`

NewPostgresqlBackupOption instantiates a new PostgresqlBackupOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlBackupOptionWithDefaults

`func NewPostgresqlBackupOptionWithDefaults() *PostgresqlBackupOption`

NewPostgresqlBackupOptionWithDefaults instantiates a new PostgresqlBackupOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveFrequencyMinute

`func (o *PostgresqlBackupOption) GetArchiveFrequencyMinute() string`

GetArchiveFrequencyMinute returns the ArchiveFrequencyMinute field if non-nil, zero value otherwise.

### GetArchiveFrequencyMinuteOk

`func (o *PostgresqlBackupOption) GetArchiveFrequencyMinuteOk() (*string, bool)`

GetArchiveFrequencyMinuteOk returns a tuple with the ArchiveFrequencyMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFrequencyMinute

`func (o *PostgresqlBackupOption) SetArchiveFrequencyMinute(v string)`

SetArchiveFrequencyMinute sets ArchiveFrequencyMinute field to given value.

### HasArchiveFrequencyMinute

`func (o *PostgresqlBackupOption) HasArchiveFrequencyMinute() bool`

HasArchiveFrequencyMinute returns a boolean if a field has been set.

### SetArchiveFrequencyMinuteNil

`func (o *PostgresqlBackupOption) SetArchiveFrequencyMinuteNil(b bool)`

 SetArchiveFrequencyMinuteNil sets the value for ArchiveFrequencyMinute to be an explicit nil

### UnsetArchiveFrequencyMinute
`func (o *PostgresqlBackupOption) UnsetArchiveFrequencyMinute()`

UnsetArchiveFrequencyMinute ensures that no value is present for ArchiveFrequencyMinute, not even an explicit nil
### GetRetentionPeriodDay

`func (o *PostgresqlBackupOption) GetRetentionPeriodDay() string`

GetRetentionPeriodDay returns the RetentionPeriodDay field if non-nil, zero value otherwise.

### GetRetentionPeriodDayOk

`func (o *PostgresqlBackupOption) GetRetentionPeriodDayOk() (*string, bool)`

GetRetentionPeriodDayOk returns a tuple with the RetentionPeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDay

`func (o *PostgresqlBackupOption) SetRetentionPeriodDay(v string)`

SetRetentionPeriodDay sets RetentionPeriodDay field to given value.

### HasRetentionPeriodDay

`func (o *PostgresqlBackupOption) HasRetentionPeriodDay() bool`

HasRetentionPeriodDay returns a boolean if a field has been set.

### SetRetentionPeriodDayNil

`func (o *PostgresqlBackupOption) SetRetentionPeriodDayNil(b bool)`

 SetRetentionPeriodDayNil sets the value for RetentionPeriodDay to be an explicit nil

### UnsetRetentionPeriodDay
`func (o *PostgresqlBackupOption) UnsetRetentionPeriodDay()`

UnsetRetentionPeriodDay ensures that no value is present for RetentionPeriodDay, not even an explicit nil
### GetStartingTimeHour

`func (o *PostgresqlBackupOption) GetStartingTimeHour() string`

GetStartingTimeHour returns the StartingTimeHour field if non-nil, zero value otherwise.

### GetStartingTimeHourOk

`func (o *PostgresqlBackupOption) GetStartingTimeHourOk() (*string, bool)`

GetStartingTimeHourOk returns a tuple with the StartingTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingTimeHour

`func (o *PostgresqlBackupOption) SetStartingTimeHour(v string)`

SetStartingTimeHour sets StartingTimeHour field to given value.

### HasStartingTimeHour

`func (o *PostgresqlBackupOption) HasStartingTimeHour() bool`

HasStartingTimeHour returns a boolean if a field has been set.

### SetStartingTimeHourNil

`func (o *PostgresqlBackupOption) SetStartingTimeHourNil(b bool)`

 SetStartingTimeHourNil sets the value for StartingTimeHour to be an explicit nil

### UnsetStartingTimeHour
`func (o *PostgresqlBackupOption) UnsetStartingTimeHour()`

UnsetStartingTimeHour ensures that no value is present for StartingTimeHour, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


