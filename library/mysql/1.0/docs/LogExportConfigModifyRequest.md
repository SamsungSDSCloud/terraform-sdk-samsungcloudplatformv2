# LogExportConfigModifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **NullableString** |  | 
**DeleteOnExport** | **NullableBool** |  | 
**ScheduleDayOfMonth** | **NullableString** |  | 
**ScheduleDayOfWeek** | [**NullableDayOfWeek**](DayOfWeek.md) |  | 
**ScheduleFrequencyType** | **NullableString** |  | 
**ScheduleHour** | **NullableString** |  | 
**SecretKey** | **NullableString** |  | 

## Methods

### NewLogExportConfigModifyRequest

`func NewLogExportConfigModifyRequest(accessKey NullableString, deleteOnExport NullableBool, scheduleDayOfMonth NullableString, scheduleDayOfWeek NullableDayOfWeek, scheduleFrequencyType NullableString, scheduleHour NullableString, secretKey NullableString, ) *LogExportConfigModifyRequest`

NewLogExportConfigModifyRequest instantiates a new LogExportConfigModifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogExportConfigModifyRequestWithDefaults

`func NewLogExportConfigModifyRequestWithDefaults() *LogExportConfigModifyRequest`

NewLogExportConfigModifyRequestWithDefaults instantiates a new LogExportConfigModifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *LogExportConfigModifyRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LogExportConfigModifyRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LogExportConfigModifyRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### SetAccessKeyNil

`func (o *LogExportConfigModifyRequest) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *LogExportConfigModifyRequest) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetDeleteOnExport

`func (o *LogExportConfigModifyRequest) GetDeleteOnExport() bool`

GetDeleteOnExport returns the DeleteOnExport field if non-nil, zero value otherwise.

### GetDeleteOnExportOk

`func (o *LogExportConfigModifyRequest) GetDeleteOnExportOk() (*bool, bool)`

GetDeleteOnExportOk returns a tuple with the DeleteOnExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExport

`func (o *LogExportConfigModifyRequest) SetDeleteOnExport(v bool)`

SetDeleteOnExport sets DeleteOnExport field to given value.


### SetDeleteOnExportNil

`func (o *LogExportConfigModifyRequest) SetDeleteOnExportNil(b bool)`

 SetDeleteOnExportNil sets the value for DeleteOnExport to be an explicit nil

### UnsetDeleteOnExport
`func (o *LogExportConfigModifyRequest) UnsetDeleteOnExport()`

UnsetDeleteOnExport ensures that no value is present for DeleteOnExport, not even an explicit nil
### GetScheduleDayOfMonth

`func (o *LogExportConfigModifyRequest) GetScheduleDayOfMonth() string`

GetScheduleDayOfMonth returns the ScheduleDayOfMonth field if non-nil, zero value otherwise.

### GetScheduleDayOfMonthOk

`func (o *LogExportConfigModifyRequest) GetScheduleDayOfMonthOk() (*string, bool)`

GetScheduleDayOfMonthOk returns a tuple with the ScheduleDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDayOfMonth

`func (o *LogExportConfigModifyRequest) SetScheduleDayOfMonth(v string)`

SetScheduleDayOfMonth sets ScheduleDayOfMonth field to given value.


### SetScheduleDayOfMonthNil

`func (o *LogExportConfigModifyRequest) SetScheduleDayOfMonthNil(b bool)`

 SetScheduleDayOfMonthNil sets the value for ScheduleDayOfMonth to be an explicit nil

### UnsetScheduleDayOfMonth
`func (o *LogExportConfigModifyRequest) UnsetScheduleDayOfMonth()`

UnsetScheduleDayOfMonth ensures that no value is present for ScheduleDayOfMonth, not even an explicit nil
### GetScheduleDayOfWeek

`func (o *LogExportConfigModifyRequest) GetScheduleDayOfWeek() DayOfWeek`

GetScheduleDayOfWeek returns the ScheduleDayOfWeek field if non-nil, zero value otherwise.

### GetScheduleDayOfWeekOk

`func (o *LogExportConfigModifyRequest) GetScheduleDayOfWeekOk() (*DayOfWeek, bool)`

GetScheduleDayOfWeekOk returns a tuple with the ScheduleDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDayOfWeek

`func (o *LogExportConfigModifyRequest) SetScheduleDayOfWeek(v DayOfWeek)`

SetScheduleDayOfWeek sets ScheduleDayOfWeek field to given value.


### SetScheduleDayOfWeekNil

`func (o *LogExportConfigModifyRequest) SetScheduleDayOfWeekNil(b bool)`

 SetScheduleDayOfWeekNil sets the value for ScheduleDayOfWeek to be an explicit nil

### UnsetScheduleDayOfWeek
`func (o *LogExportConfigModifyRequest) UnsetScheduleDayOfWeek()`

UnsetScheduleDayOfWeek ensures that no value is present for ScheduleDayOfWeek, not even an explicit nil
### GetScheduleFrequencyType

`func (o *LogExportConfigModifyRequest) GetScheduleFrequencyType() string`

GetScheduleFrequencyType returns the ScheduleFrequencyType field if non-nil, zero value otherwise.

### GetScheduleFrequencyTypeOk

`func (o *LogExportConfigModifyRequest) GetScheduleFrequencyTypeOk() (*string, bool)`

GetScheduleFrequencyTypeOk returns a tuple with the ScheduleFrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleFrequencyType

`func (o *LogExportConfigModifyRequest) SetScheduleFrequencyType(v string)`

SetScheduleFrequencyType sets ScheduleFrequencyType field to given value.


### SetScheduleFrequencyTypeNil

`func (o *LogExportConfigModifyRequest) SetScheduleFrequencyTypeNil(b bool)`

 SetScheduleFrequencyTypeNil sets the value for ScheduleFrequencyType to be an explicit nil

### UnsetScheduleFrequencyType
`func (o *LogExportConfigModifyRequest) UnsetScheduleFrequencyType()`

UnsetScheduleFrequencyType ensures that no value is present for ScheduleFrequencyType, not even an explicit nil
### GetScheduleHour

`func (o *LogExportConfigModifyRequest) GetScheduleHour() string`

GetScheduleHour returns the ScheduleHour field if non-nil, zero value otherwise.

### GetScheduleHourOk

`func (o *LogExportConfigModifyRequest) GetScheduleHourOk() (*string, bool)`

GetScheduleHourOk returns a tuple with the ScheduleHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleHour

`func (o *LogExportConfigModifyRequest) SetScheduleHour(v string)`

SetScheduleHour sets ScheduleHour field to given value.


### SetScheduleHourNil

`func (o *LogExportConfigModifyRequest) SetScheduleHourNil(b bool)`

 SetScheduleHourNil sets the value for ScheduleHour to be an explicit nil

### UnsetScheduleHour
`func (o *LogExportConfigModifyRequest) UnsetScheduleHour()`

UnsetScheduleHour ensures that no value is present for ScheduleHour, not even an explicit nil
### GetSecretKey

`func (o *LogExportConfigModifyRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LogExportConfigModifyRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LogExportConfigModifyRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### SetSecretKeyNil

`func (o *LogExportConfigModifyRequest) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *LogExportConfigModifyRequest) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


