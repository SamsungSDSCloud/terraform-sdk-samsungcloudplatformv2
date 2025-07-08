# SnapshotSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **string** | Day of week | [optional] [default to "MON"]
**Frequency** | Pointer to [**SnapshotFrequency**](SnapshotFrequency.md) | Frequency | [optional] 
**Hour** | Pointer to **int32** | Hour | [optional] [default to 0]

## Methods

### NewSnapshotSchedule

`func NewSnapshotSchedule() *SnapshotSchedule`

NewSnapshotSchedule instantiates a new SnapshotSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotScheduleWithDefaults

`func NewSnapshotScheduleWithDefaults() *SnapshotSchedule`

NewSnapshotScheduleWithDefaults instantiates a new SnapshotSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *SnapshotSchedule) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SnapshotSchedule) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SnapshotSchedule) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SnapshotSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetFrequency

`func (o *SnapshotSchedule) GetFrequency() SnapshotFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SnapshotSchedule) GetFrequencyOk() (*SnapshotFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SnapshotSchedule) SetFrequency(v SnapshotFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *SnapshotSchedule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetHour

`func (o *SnapshotSchedule) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *SnapshotSchedule) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *SnapshotSchedule) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *SnapshotSchedule) HasHour() bool`

HasHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


