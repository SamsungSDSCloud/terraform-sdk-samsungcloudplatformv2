# VolumeGroupSnapshotScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) | Day of week(If this field is excluded, the schedule frequency is set to DAILY.) | [optional] 
**Hour** | **int32** | Hour | 

## Methods

### NewVolumeGroupSnapshotScheduleRequest

`func NewVolumeGroupSnapshotScheduleRequest(hour int32, ) *VolumeGroupSnapshotScheduleRequest`

NewVolumeGroupSnapshotScheduleRequest instantiates a new VolumeGroupSnapshotScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupSnapshotScheduleRequestWithDefaults

`func NewVolumeGroupSnapshotScheduleRequestWithDefaults() *VolumeGroupSnapshotScheduleRequest`

NewVolumeGroupSnapshotScheduleRequestWithDefaults instantiates a new VolumeGroupSnapshotScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *VolumeGroupSnapshotScheduleRequest) GetDayOfWeek() DayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *VolumeGroupSnapshotScheduleRequest) GetDayOfWeekOk() (*DayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *VolumeGroupSnapshotScheduleRequest) SetDayOfWeek(v DayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *VolumeGroupSnapshotScheduleRequest) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *VolumeGroupSnapshotScheduleRequest) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *VolumeGroupSnapshotScheduleRequest) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *VolumeGroupSnapshotScheduleRequest) SetHour(v int32)`

SetHour sets Hour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


