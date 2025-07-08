# MariadbBackupOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveFrequencyMinute** | Pointer to **NullableString** |  | [optional] 
**RetentionPeriodDay** | Pointer to **string** | Backup retention period (day) | [optional] 
**StartingTimeHour** | Pointer to **string** | Backup starting time (hour) | [optional] 

## Methods

### NewMariadbBackupOption

`func NewMariadbBackupOption() *MariadbBackupOption`

NewMariadbBackupOption instantiates a new MariadbBackupOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariadbBackupOptionWithDefaults

`func NewMariadbBackupOptionWithDefaults() *MariadbBackupOption`

NewMariadbBackupOptionWithDefaults instantiates a new MariadbBackupOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveFrequencyMinute

`func (o *MariadbBackupOption) GetArchiveFrequencyMinute() string`

GetArchiveFrequencyMinute returns the ArchiveFrequencyMinute field if non-nil, zero value otherwise.

### GetArchiveFrequencyMinuteOk

`func (o *MariadbBackupOption) GetArchiveFrequencyMinuteOk() (*string, bool)`

GetArchiveFrequencyMinuteOk returns a tuple with the ArchiveFrequencyMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFrequencyMinute

`func (o *MariadbBackupOption) SetArchiveFrequencyMinute(v string)`

SetArchiveFrequencyMinute sets ArchiveFrequencyMinute field to given value.

### HasArchiveFrequencyMinute

`func (o *MariadbBackupOption) HasArchiveFrequencyMinute() bool`

HasArchiveFrequencyMinute returns a boolean if a field has been set.

### SetArchiveFrequencyMinuteNil

`func (o *MariadbBackupOption) SetArchiveFrequencyMinuteNil(b bool)`

 SetArchiveFrequencyMinuteNil sets the value for ArchiveFrequencyMinute to be an explicit nil

### UnsetArchiveFrequencyMinute
`func (o *MariadbBackupOption) UnsetArchiveFrequencyMinute()`

UnsetArchiveFrequencyMinute ensures that no value is present for ArchiveFrequencyMinute, not even an explicit nil
### GetRetentionPeriodDay

`func (o *MariadbBackupOption) GetRetentionPeriodDay() string`

GetRetentionPeriodDay returns the RetentionPeriodDay field if non-nil, zero value otherwise.

### GetRetentionPeriodDayOk

`func (o *MariadbBackupOption) GetRetentionPeriodDayOk() (*string, bool)`

GetRetentionPeriodDayOk returns a tuple with the RetentionPeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDay

`func (o *MariadbBackupOption) SetRetentionPeriodDay(v string)`

SetRetentionPeriodDay sets RetentionPeriodDay field to given value.

### HasRetentionPeriodDay

`func (o *MariadbBackupOption) HasRetentionPeriodDay() bool`

HasRetentionPeriodDay returns a boolean if a field has been set.

### GetStartingTimeHour

`func (o *MariadbBackupOption) GetStartingTimeHour() string`

GetStartingTimeHour returns the StartingTimeHour field if non-nil, zero value otherwise.

### GetStartingTimeHourOk

`func (o *MariadbBackupOption) GetStartingTimeHourOk() (*string, bool)`

GetStartingTimeHourOk returns a tuple with the StartingTimeHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingTimeHour

`func (o *MariadbBackupOption) SetStartingTimeHour(v string)`

SetStartingTimeHour sets StartingTimeHour field to given value.

### HasStartingTimeHour

`func (o *MariadbBackupOption) HasStartingTimeHour() bool`

HasStartingTimeHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


