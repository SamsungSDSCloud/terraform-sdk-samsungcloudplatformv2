# BackupScheduleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | [**BackupScheduleFrequency**](BackupScheduleFrequency.md) | Schedule frequency type | 
**StartDay** | Pointer to [**NullableBackupScheduleDay**](BackupScheduleDay.md) |  | [optional] 
**StartTime** | **string** | Backup schedule start time | 
**StartWeek** | Pointer to [**NullableBackupScheduleWeek**](BackupScheduleWeek.md) |  | [optional] 
**Type** | [**BackupScheduleType**](BackupScheduleType.md) | Schedule type | 

## Methods

### NewBackupScheduleCreateRequest

`func NewBackupScheduleCreateRequest(frequency BackupScheduleFrequency, startTime string, type_ BackupScheduleType, ) *BackupScheduleCreateRequest`

NewBackupScheduleCreateRequest instantiates a new BackupScheduleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleCreateRequestWithDefaults

`func NewBackupScheduleCreateRequestWithDefaults() *BackupScheduleCreateRequest`

NewBackupScheduleCreateRequestWithDefaults instantiates a new BackupScheduleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *BackupScheduleCreateRequest) GetFrequency() BackupScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *BackupScheduleCreateRequest) GetFrequencyOk() (*BackupScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *BackupScheduleCreateRequest) SetFrequency(v BackupScheduleFrequency)`

SetFrequency sets Frequency field to given value.


### GetStartDay

`func (o *BackupScheduleCreateRequest) GetStartDay() BackupScheduleDay`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *BackupScheduleCreateRequest) GetStartDayOk() (*BackupScheduleDay, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *BackupScheduleCreateRequest) SetStartDay(v BackupScheduleDay)`

SetStartDay sets StartDay field to given value.

### HasStartDay

`func (o *BackupScheduleCreateRequest) HasStartDay() bool`

HasStartDay returns a boolean if a field has been set.

### SetStartDayNil

`func (o *BackupScheduleCreateRequest) SetStartDayNil(b bool)`

 SetStartDayNil sets the value for StartDay to be an explicit nil

### UnsetStartDay
`func (o *BackupScheduleCreateRequest) UnsetStartDay()`

UnsetStartDay ensures that no value is present for StartDay, not even an explicit nil
### GetStartTime

`func (o *BackupScheduleCreateRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupScheduleCreateRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupScheduleCreateRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetStartWeek

`func (o *BackupScheduleCreateRequest) GetStartWeek() BackupScheduleWeek`

GetStartWeek returns the StartWeek field if non-nil, zero value otherwise.

### GetStartWeekOk

`func (o *BackupScheduleCreateRequest) GetStartWeekOk() (*BackupScheduleWeek, bool)`

GetStartWeekOk returns a tuple with the StartWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWeek

`func (o *BackupScheduleCreateRequest) SetStartWeek(v BackupScheduleWeek)`

SetStartWeek sets StartWeek field to given value.

### HasStartWeek

`func (o *BackupScheduleCreateRequest) HasStartWeek() bool`

HasStartWeek returns a boolean if a field has been set.

### SetStartWeekNil

`func (o *BackupScheduleCreateRequest) SetStartWeekNil(b bool)`

 SetStartWeekNil sets the value for StartWeek to be an explicit nil

### UnsetStartWeek
`func (o *BackupScheduleCreateRequest) UnsetStartWeek()`

UnsetStartWeek ensures that no value is present for StartWeek, not even an explicit nil
### GetType

`func (o *BackupScheduleCreateRequest) GetType() BackupScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupScheduleCreateRequest) GetTypeOk() (*BackupScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupScheduleCreateRequest) SetType(v BackupScheduleType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


