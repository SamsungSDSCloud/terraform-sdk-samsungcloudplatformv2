# VolumeGroupSnapshotScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) | Day of week | [optional] 
**Frequency** | Pointer to [**SnapshotFrequency**](SnapshotFrequency.md) | Frequency DAILY|WEEKLY|NONE | [optional] 
**Hour** | Pointer to **int32** | Hour | [optional] [default to 0]
**VolumeGroupId** | Pointer to **string** | Volume group id | [optional] 

## Methods

### NewVolumeGroupSnapshotScheduleResponse

`func NewVolumeGroupSnapshotScheduleResponse() *VolumeGroupSnapshotScheduleResponse`

NewVolumeGroupSnapshotScheduleResponse instantiates a new VolumeGroupSnapshotScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupSnapshotScheduleResponseWithDefaults

`func NewVolumeGroupSnapshotScheduleResponseWithDefaults() *VolumeGroupSnapshotScheduleResponse`

NewVolumeGroupSnapshotScheduleResponseWithDefaults instantiates a new VolumeGroupSnapshotScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *VolumeGroupSnapshotScheduleResponse) GetDayOfWeek() DayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *VolumeGroupSnapshotScheduleResponse) GetDayOfWeekOk() (*DayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *VolumeGroupSnapshotScheduleResponse) SetDayOfWeek(v DayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *VolumeGroupSnapshotScheduleResponse) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetFrequency

`func (o *VolumeGroupSnapshotScheduleResponse) GetFrequency() SnapshotFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *VolumeGroupSnapshotScheduleResponse) GetFrequencyOk() (*SnapshotFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *VolumeGroupSnapshotScheduleResponse) SetFrequency(v SnapshotFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *VolumeGroupSnapshotScheduleResponse) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetHour

`func (o *VolumeGroupSnapshotScheduleResponse) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *VolumeGroupSnapshotScheduleResponse) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *VolumeGroupSnapshotScheduleResponse) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *VolumeGroupSnapshotScheduleResponse) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetVolumeGroupId

`func (o *VolumeGroupSnapshotScheduleResponse) GetVolumeGroupId() string`

GetVolumeGroupId returns the VolumeGroupId field if non-nil, zero value otherwise.

### GetVolumeGroupIdOk

`func (o *VolumeGroupSnapshotScheduleResponse) GetVolumeGroupIdOk() (*string, bool)`

GetVolumeGroupIdOk returns a tuple with the VolumeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupId

`func (o *VolumeGroupSnapshotScheduleResponse) SetVolumeGroupId(v string)`

SetVolumeGroupId sets VolumeGroupId field to given value.

### HasVolumeGroupId

`func (o *VolumeGroupSnapshotScheduleResponse) HasVolumeGroupId() bool`

HasVolumeGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


