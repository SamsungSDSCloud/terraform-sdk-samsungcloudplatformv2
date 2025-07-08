# VolumeSnapshotScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) | Day of week(If this field is excluded, the schedule frequency is set to DAILY.) | [optional] 
**Hour** | **int32** | Hour | 

## Methods

### NewVolumeSnapshotScheduleRequest

`func NewVolumeSnapshotScheduleRequest(hour int32, ) *VolumeSnapshotScheduleRequest`

NewVolumeSnapshotScheduleRequest instantiates a new VolumeSnapshotScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSnapshotScheduleRequestWithDefaults

`func NewVolumeSnapshotScheduleRequestWithDefaults() *VolumeSnapshotScheduleRequest`

NewVolumeSnapshotScheduleRequestWithDefaults instantiates a new VolumeSnapshotScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *VolumeSnapshotScheduleRequest) GetDayOfWeek() DayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *VolumeSnapshotScheduleRequest) GetDayOfWeekOk() (*DayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *VolumeSnapshotScheduleRequest) SetDayOfWeek(v DayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *VolumeSnapshotScheduleRequest) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *VolumeSnapshotScheduleRequest) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *VolumeSnapshotScheduleRequest) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *VolumeSnapshotScheduleRequest) SetHour(v int32)`

SetHour sets Hour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


