# SnapshotScheduleShow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **NullableString** |  | [optional] 
**Frequency** | **string** | Frequency of the snapshot | 
**Hour** | **string** | Hour of the snapshot schedule | 
**Id** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSnapshotScheduleShow

`func NewSnapshotScheduleShow(frequency string, hour string, ) *SnapshotScheduleShow`

NewSnapshotScheduleShow instantiates a new SnapshotScheduleShow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotScheduleShowWithDefaults

`func NewSnapshotScheduleShowWithDefaults() *SnapshotScheduleShow`

NewSnapshotScheduleShowWithDefaults instantiates a new SnapshotScheduleShow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *SnapshotScheduleShow) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *SnapshotScheduleShow) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *SnapshotScheduleShow) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *SnapshotScheduleShow) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *SnapshotScheduleShow) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *SnapshotScheduleShow) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetFrequency

`func (o *SnapshotScheduleShow) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *SnapshotScheduleShow) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *SnapshotScheduleShow) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.


### GetHour

`func (o *SnapshotScheduleShow) GetHour() string`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *SnapshotScheduleShow) GetHourOk() (*string, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *SnapshotScheduleShow) SetHour(v string)`

SetHour sets Hour field to given value.


### GetId

`func (o *SnapshotScheduleShow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotScheduleShow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotScheduleShow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotScheduleShow) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SnapshotScheduleShow) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SnapshotScheduleShow) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


