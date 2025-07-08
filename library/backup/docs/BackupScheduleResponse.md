# BackupScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **string** | Backup ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Frequency** | [**BackupScheduleFrequency**](BackupScheduleFrequency.md) | Schedule frequency type | 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Schedule name | 
**StartDay** | **NullableString** |  | 
**StartTime** | **NullableString** |  | 
**StartWeek** | **NullableString** |  | 
**State** | **string** | Backup Schedule state | 
**Type** | **string** | Schedule type | 

## Methods

### NewBackupScheduleResponse

`func NewBackupScheduleResponse(backupId string, createdAt time.Time, createdBy string, frequency BackupScheduleFrequency, id string, modifiedAt time.Time, modifiedBy string, name string, startDay NullableString, startTime NullableString, startWeek NullableString, state string, type_ string, ) *BackupScheduleResponse`

NewBackupScheduleResponse instantiates a new BackupScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleResponseWithDefaults

`func NewBackupScheduleResponseWithDefaults() *BackupScheduleResponse`

NewBackupScheduleResponseWithDefaults instantiates a new BackupScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *BackupScheduleResponse) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *BackupScheduleResponse) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *BackupScheduleResponse) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetCreatedAt

`func (o *BackupScheduleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupScheduleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupScheduleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *BackupScheduleResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BackupScheduleResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BackupScheduleResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetFrequency

`func (o *BackupScheduleResponse) GetFrequency() BackupScheduleFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *BackupScheduleResponse) GetFrequencyOk() (*BackupScheduleFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *BackupScheduleResponse) SetFrequency(v BackupScheduleFrequency)`

SetFrequency sets Frequency field to given value.


### GetId

`func (o *BackupScheduleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupScheduleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupScheduleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *BackupScheduleResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BackupScheduleResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BackupScheduleResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *BackupScheduleResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BackupScheduleResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BackupScheduleResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *BackupScheduleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupScheduleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupScheduleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStartDay

`func (o *BackupScheduleResponse) GetStartDay() string`

GetStartDay returns the StartDay field if non-nil, zero value otherwise.

### GetStartDayOk

`func (o *BackupScheduleResponse) GetStartDayOk() (*string, bool)`

GetStartDayOk returns a tuple with the StartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDay

`func (o *BackupScheduleResponse) SetStartDay(v string)`

SetStartDay sets StartDay field to given value.


### SetStartDayNil

`func (o *BackupScheduleResponse) SetStartDayNil(b bool)`

 SetStartDayNil sets the value for StartDay to be an explicit nil

### UnsetStartDay
`func (o *BackupScheduleResponse) UnsetStartDay()`

UnsetStartDay ensures that no value is present for StartDay, not even an explicit nil
### GetStartTime

`func (o *BackupScheduleResponse) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupScheduleResponse) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupScheduleResponse) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### SetStartTimeNil

`func (o *BackupScheduleResponse) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BackupScheduleResponse) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetStartWeek

`func (o *BackupScheduleResponse) GetStartWeek() string`

GetStartWeek returns the StartWeek field if non-nil, zero value otherwise.

### GetStartWeekOk

`func (o *BackupScheduleResponse) GetStartWeekOk() (*string, bool)`

GetStartWeekOk returns a tuple with the StartWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartWeek

`func (o *BackupScheduleResponse) SetStartWeek(v string)`

SetStartWeek sets StartWeek field to given value.


### SetStartWeekNil

`func (o *BackupScheduleResponse) SetStartWeekNil(b bool)`

 SetStartWeekNil sets the value for StartWeek to be an explicit nil

### UnsetStartWeek
`func (o *BackupScheduleResponse) UnsetStartWeek()`

UnsetStartWeek ensures that no value is present for StartWeek, not even an explicit nil
### GetState

`func (o *BackupScheduleResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupScheduleResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupScheduleResponse) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *BackupScheduleResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupScheduleResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupScheduleResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


