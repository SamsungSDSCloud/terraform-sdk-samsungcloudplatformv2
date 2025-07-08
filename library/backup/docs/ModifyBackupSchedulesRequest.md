# ModifyBackupSchedulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | [**[]BackupScheduleCreateRequest**](BackupScheduleCreateRequest.md) | Backup schedule list | 

## Methods

### NewModifyBackupSchedulesRequest

`func NewModifyBackupSchedulesRequest(schedules []BackupScheduleCreateRequest, ) *ModifyBackupSchedulesRequest`

NewModifyBackupSchedulesRequest instantiates a new ModifyBackupSchedulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyBackupSchedulesRequestWithDefaults

`func NewModifyBackupSchedulesRequestWithDefaults() *ModifyBackupSchedulesRequest`

NewModifyBackupSchedulesRequestWithDefaults instantiates a new ModifyBackupSchedulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedules

`func (o *ModifyBackupSchedulesRequest) GetSchedules() []BackupScheduleCreateRequest`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ModifyBackupSchedulesRequest) GetSchedulesOk() (*[]BackupScheduleCreateRequest, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ModifyBackupSchedulesRequest) SetSchedules(v []BackupScheduleCreateRequest)`

SetSchedules sets Schedules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


