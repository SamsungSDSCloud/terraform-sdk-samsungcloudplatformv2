# LogExportConfigCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **NullableString** |  | 
**BucketName** | **string** | Bucket name | 
**DeleteOnExport** | **NullableBool** |  | 
**LogType** | **string** | Log type | 
**ScheduleDayOfMonth** | **NullableString** |  | 
**ScheduleDayOfWeek** | [**NullableDayOfWeek**](DayOfWeek.md) |  | 
**ScheduleFrequencyType** | **NullableString** |  | 
**ScheduleHour** | **NullableString** |  | 
**SecretKey** | **NullableString** |  | 

## Methods

### NewLogExportConfigCreateRequest

`func NewLogExportConfigCreateRequest(accessKey NullableString, bucketName string, deleteOnExport NullableBool, logType string, scheduleDayOfMonth NullableString, scheduleDayOfWeek NullableDayOfWeek, scheduleFrequencyType NullableString, scheduleHour NullableString, secretKey NullableString, ) *LogExportConfigCreateRequest`

NewLogExportConfigCreateRequest instantiates a new LogExportConfigCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogExportConfigCreateRequestWithDefaults

`func NewLogExportConfigCreateRequestWithDefaults() *LogExportConfigCreateRequest`

NewLogExportConfigCreateRequestWithDefaults instantiates a new LogExportConfigCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *LogExportConfigCreateRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LogExportConfigCreateRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LogExportConfigCreateRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### SetAccessKeyNil

`func (o *LogExportConfigCreateRequest) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *LogExportConfigCreateRequest) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetBucketName

`func (o *LogExportConfigCreateRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LogExportConfigCreateRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LogExportConfigCreateRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetDeleteOnExport

`func (o *LogExportConfigCreateRequest) GetDeleteOnExport() bool`

GetDeleteOnExport returns the DeleteOnExport field if non-nil, zero value otherwise.

### GetDeleteOnExportOk

`func (o *LogExportConfigCreateRequest) GetDeleteOnExportOk() (*bool, bool)`

GetDeleteOnExportOk returns a tuple with the DeleteOnExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExport

`func (o *LogExportConfigCreateRequest) SetDeleteOnExport(v bool)`

SetDeleteOnExport sets DeleteOnExport field to given value.


### SetDeleteOnExportNil

`func (o *LogExportConfigCreateRequest) SetDeleteOnExportNil(b bool)`

 SetDeleteOnExportNil sets the value for DeleteOnExport to be an explicit nil

### UnsetDeleteOnExport
`func (o *LogExportConfigCreateRequest) UnsetDeleteOnExport()`

UnsetDeleteOnExport ensures that no value is present for DeleteOnExport, not even an explicit nil
### GetLogType

`func (o *LogExportConfigCreateRequest) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *LogExportConfigCreateRequest) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *LogExportConfigCreateRequest) SetLogType(v string)`

SetLogType sets LogType field to given value.


### GetScheduleDayOfMonth

`func (o *LogExportConfigCreateRequest) GetScheduleDayOfMonth() string`

GetScheduleDayOfMonth returns the ScheduleDayOfMonth field if non-nil, zero value otherwise.

### GetScheduleDayOfMonthOk

`func (o *LogExportConfigCreateRequest) GetScheduleDayOfMonthOk() (*string, bool)`

GetScheduleDayOfMonthOk returns a tuple with the ScheduleDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDayOfMonth

`func (o *LogExportConfigCreateRequest) SetScheduleDayOfMonth(v string)`

SetScheduleDayOfMonth sets ScheduleDayOfMonth field to given value.


### SetScheduleDayOfMonthNil

`func (o *LogExportConfigCreateRequest) SetScheduleDayOfMonthNil(b bool)`

 SetScheduleDayOfMonthNil sets the value for ScheduleDayOfMonth to be an explicit nil

### UnsetScheduleDayOfMonth
`func (o *LogExportConfigCreateRequest) UnsetScheduleDayOfMonth()`

UnsetScheduleDayOfMonth ensures that no value is present for ScheduleDayOfMonth, not even an explicit nil
### GetScheduleDayOfWeek

`func (o *LogExportConfigCreateRequest) GetScheduleDayOfWeek() DayOfWeek`

GetScheduleDayOfWeek returns the ScheduleDayOfWeek field if non-nil, zero value otherwise.

### GetScheduleDayOfWeekOk

`func (o *LogExportConfigCreateRequest) GetScheduleDayOfWeekOk() (*DayOfWeek, bool)`

GetScheduleDayOfWeekOk returns a tuple with the ScheduleDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDayOfWeek

`func (o *LogExportConfigCreateRequest) SetScheduleDayOfWeek(v DayOfWeek)`

SetScheduleDayOfWeek sets ScheduleDayOfWeek field to given value.


### SetScheduleDayOfWeekNil

`func (o *LogExportConfigCreateRequest) SetScheduleDayOfWeekNil(b bool)`

 SetScheduleDayOfWeekNil sets the value for ScheduleDayOfWeek to be an explicit nil

### UnsetScheduleDayOfWeek
`func (o *LogExportConfigCreateRequest) UnsetScheduleDayOfWeek()`

UnsetScheduleDayOfWeek ensures that no value is present for ScheduleDayOfWeek, not even an explicit nil
### GetScheduleFrequencyType

`func (o *LogExportConfigCreateRequest) GetScheduleFrequencyType() string`

GetScheduleFrequencyType returns the ScheduleFrequencyType field if non-nil, zero value otherwise.

### GetScheduleFrequencyTypeOk

`func (o *LogExportConfigCreateRequest) GetScheduleFrequencyTypeOk() (*string, bool)`

GetScheduleFrequencyTypeOk returns a tuple with the ScheduleFrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleFrequencyType

`func (o *LogExportConfigCreateRequest) SetScheduleFrequencyType(v string)`

SetScheduleFrequencyType sets ScheduleFrequencyType field to given value.


### SetScheduleFrequencyTypeNil

`func (o *LogExportConfigCreateRequest) SetScheduleFrequencyTypeNil(b bool)`

 SetScheduleFrequencyTypeNil sets the value for ScheduleFrequencyType to be an explicit nil

### UnsetScheduleFrequencyType
`func (o *LogExportConfigCreateRequest) UnsetScheduleFrequencyType()`

UnsetScheduleFrequencyType ensures that no value is present for ScheduleFrequencyType, not even an explicit nil
### GetScheduleHour

`func (o *LogExportConfigCreateRequest) GetScheduleHour() string`

GetScheduleHour returns the ScheduleHour field if non-nil, zero value otherwise.

### GetScheduleHourOk

`func (o *LogExportConfigCreateRequest) GetScheduleHourOk() (*string, bool)`

GetScheduleHourOk returns a tuple with the ScheduleHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleHour

`func (o *LogExportConfigCreateRequest) SetScheduleHour(v string)`

SetScheduleHour sets ScheduleHour field to given value.


### SetScheduleHourNil

`func (o *LogExportConfigCreateRequest) SetScheduleHourNil(b bool)`

 SetScheduleHourNil sets the value for ScheduleHour to be an explicit nil

### UnsetScheduleHour
`func (o *LogExportConfigCreateRequest) UnsetScheduleHour()`

UnsetScheduleHour ensures that no value is present for ScheduleHour, not even an explicit nil
### GetSecretKey

`func (o *LogExportConfigCreateRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LogExportConfigCreateRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LogExportConfigCreateRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### SetSecretKeyNil

`func (o *LogExportConfigCreateRequest) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *LogExportConfigCreateRequest) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


