# LogExportConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **NullableString** |  | 
**DeleteOnExport** | **NullableBool** |  | 
**IsBucketDeleted** | **NullableBool** |  | 
**IsRegistered** | **bool** | Log type registration state | 
**LogLabel** | **string** | Log label | 
**LogType** | **string** | Log type | 
**ScheduleDayOfMonth** | **NullableString** |  | 
**ScheduleDayOfWeek** | [**NullableDayOfWeek**](DayOfWeek.md) |  | 
**ScheduleFrequencyType** | **NullableString** |  | 
**ScheduleHour** | **NullableString** |  | 

## Methods

### NewLogExportConfigDTO

`func NewLogExportConfigDTO(bucketName NullableString, deleteOnExport NullableBool, isBucketDeleted NullableBool, isRegistered bool, logLabel string, logType string, scheduleDayOfMonth NullableString, scheduleDayOfWeek NullableDayOfWeek, scheduleFrequencyType NullableString, scheduleHour NullableString, ) *LogExportConfigDTO`

NewLogExportConfigDTO instantiates a new LogExportConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogExportConfigDTOWithDefaults

`func NewLogExportConfigDTOWithDefaults() *LogExportConfigDTO`

NewLogExportConfigDTOWithDefaults instantiates a new LogExportConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *LogExportConfigDTO) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LogExportConfigDTO) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LogExportConfigDTO) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### SetBucketNameNil

`func (o *LogExportConfigDTO) SetBucketNameNil(b bool)`

 SetBucketNameNil sets the value for BucketName to be an explicit nil

### UnsetBucketName
`func (o *LogExportConfigDTO) UnsetBucketName()`

UnsetBucketName ensures that no value is present for BucketName, not even an explicit nil
### GetDeleteOnExport

`func (o *LogExportConfigDTO) GetDeleteOnExport() bool`

GetDeleteOnExport returns the DeleteOnExport field if non-nil, zero value otherwise.

### GetDeleteOnExportOk

`func (o *LogExportConfigDTO) GetDeleteOnExportOk() (*bool, bool)`

GetDeleteOnExportOk returns a tuple with the DeleteOnExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExport

`func (o *LogExportConfigDTO) SetDeleteOnExport(v bool)`

SetDeleteOnExport sets DeleteOnExport field to given value.


### SetDeleteOnExportNil

`func (o *LogExportConfigDTO) SetDeleteOnExportNil(b bool)`

 SetDeleteOnExportNil sets the value for DeleteOnExport to be an explicit nil

### UnsetDeleteOnExport
`func (o *LogExportConfigDTO) UnsetDeleteOnExport()`

UnsetDeleteOnExport ensures that no value is present for DeleteOnExport, not even an explicit nil
### GetIsBucketDeleted

`func (o *LogExportConfigDTO) GetIsBucketDeleted() bool`

GetIsBucketDeleted returns the IsBucketDeleted field if non-nil, zero value otherwise.

### GetIsBucketDeletedOk

`func (o *LogExportConfigDTO) GetIsBucketDeletedOk() (*bool, bool)`

GetIsBucketDeletedOk returns a tuple with the IsBucketDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBucketDeleted

`func (o *LogExportConfigDTO) SetIsBucketDeleted(v bool)`

SetIsBucketDeleted sets IsBucketDeleted field to given value.


### SetIsBucketDeletedNil

`func (o *LogExportConfigDTO) SetIsBucketDeletedNil(b bool)`

 SetIsBucketDeletedNil sets the value for IsBucketDeleted to be an explicit nil

### UnsetIsBucketDeleted
`func (o *LogExportConfigDTO) UnsetIsBucketDeleted()`

UnsetIsBucketDeleted ensures that no value is present for IsBucketDeleted, not even an explicit nil
### GetIsRegistered

`func (o *LogExportConfigDTO) GetIsRegistered() bool`

GetIsRegistered returns the IsRegistered field if non-nil, zero value otherwise.

### GetIsRegisteredOk

`func (o *LogExportConfigDTO) GetIsRegisteredOk() (*bool, bool)`

GetIsRegisteredOk returns a tuple with the IsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegistered

`func (o *LogExportConfigDTO) SetIsRegistered(v bool)`

SetIsRegistered sets IsRegistered field to given value.


### GetLogLabel

`func (o *LogExportConfigDTO) GetLogLabel() string`

GetLogLabel returns the LogLabel field if non-nil, zero value otherwise.

### GetLogLabelOk

`func (o *LogExportConfigDTO) GetLogLabelOk() (*string, bool)`

GetLogLabelOk returns a tuple with the LogLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLabel

`func (o *LogExportConfigDTO) SetLogLabel(v string)`

SetLogLabel sets LogLabel field to given value.


### GetLogType

`func (o *LogExportConfigDTO) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *LogExportConfigDTO) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *LogExportConfigDTO) SetLogType(v string)`

SetLogType sets LogType field to given value.


### GetScheduleDayOfMonth

`func (o *LogExportConfigDTO) GetScheduleDayOfMonth() string`

GetScheduleDayOfMonth returns the ScheduleDayOfMonth field if non-nil, zero value otherwise.

### GetScheduleDayOfMonthOk

`func (o *LogExportConfigDTO) GetScheduleDayOfMonthOk() (*string, bool)`

GetScheduleDayOfMonthOk returns a tuple with the ScheduleDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDayOfMonth

`func (o *LogExportConfigDTO) SetScheduleDayOfMonth(v string)`

SetScheduleDayOfMonth sets ScheduleDayOfMonth field to given value.


### SetScheduleDayOfMonthNil

`func (o *LogExportConfigDTO) SetScheduleDayOfMonthNil(b bool)`

 SetScheduleDayOfMonthNil sets the value for ScheduleDayOfMonth to be an explicit nil

### UnsetScheduleDayOfMonth
`func (o *LogExportConfigDTO) UnsetScheduleDayOfMonth()`

UnsetScheduleDayOfMonth ensures that no value is present for ScheduleDayOfMonth, not even an explicit nil
### GetScheduleDayOfWeek

`func (o *LogExportConfigDTO) GetScheduleDayOfWeek() DayOfWeek`

GetScheduleDayOfWeek returns the ScheduleDayOfWeek field if non-nil, zero value otherwise.

### GetScheduleDayOfWeekOk

`func (o *LogExportConfigDTO) GetScheduleDayOfWeekOk() (*DayOfWeek, bool)`

GetScheduleDayOfWeekOk returns a tuple with the ScheduleDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDayOfWeek

`func (o *LogExportConfigDTO) SetScheduleDayOfWeek(v DayOfWeek)`

SetScheduleDayOfWeek sets ScheduleDayOfWeek field to given value.


### SetScheduleDayOfWeekNil

`func (o *LogExportConfigDTO) SetScheduleDayOfWeekNil(b bool)`

 SetScheduleDayOfWeekNil sets the value for ScheduleDayOfWeek to be an explicit nil

### UnsetScheduleDayOfWeek
`func (o *LogExportConfigDTO) UnsetScheduleDayOfWeek()`

UnsetScheduleDayOfWeek ensures that no value is present for ScheduleDayOfWeek, not even an explicit nil
### GetScheduleFrequencyType

`func (o *LogExportConfigDTO) GetScheduleFrequencyType() string`

GetScheduleFrequencyType returns the ScheduleFrequencyType field if non-nil, zero value otherwise.

### GetScheduleFrequencyTypeOk

`func (o *LogExportConfigDTO) GetScheduleFrequencyTypeOk() (*string, bool)`

GetScheduleFrequencyTypeOk returns a tuple with the ScheduleFrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleFrequencyType

`func (o *LogExportConfigDTO) SetScheduleFrequencyType(v string)`

SetScheduleFrequencyType sets ScheduleFrequencyType field to given value.


### SetScheduleFrequencyTypeNil

`func (o *LogExportConfigDTO) SetScheduleFrequencyTypeNil(b bool)`

 SetScheduleFrequencyTypeNil sets the value for ScheduleFrequencyType to be an explicit nil

### UnsetScheduleFrequencyType
`func (o *LogExportConfigDTO) UnsetScheduleFrequencyType()`

UnsetScheduleFrequencyType ensures that no value is present for ScheduleFrequencyType, not even an explicit nil
### GetScheduleHour

`func (o *LogExportConfigDTO) GetScheduleHour() string`

GetScheduleHour returns the ScheduleHour field if non-nil, zero value otherwise.

### GetScheduleHourOk

`func (o *LogExportConfigDTO) GetScheduleHourOk() (*string, bool)`

GetScheduleHourOk returns a tuple with the ScheduleHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleHour

`func (o *LogExportConfigDTO) SetScheduleHour(v string)`

SetScheduleHour sets ScheduleHour field to given value.


### SetScheduleHourNil

`func (o *LogExportConfigDTO) SetScheduleHourNil(b bool)`

 SetScheduleHourNil sets the value for ScheduleHour to be an explicit nil

### UnsetScheduleHour
`func (o *LogExportConfigDTO) UnsetScheduleHour()`

UnsetScheduleHour ensures that no value is present for ScheduleHour, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


