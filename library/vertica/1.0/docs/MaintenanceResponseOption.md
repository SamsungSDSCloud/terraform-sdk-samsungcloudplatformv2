# MaintenanceResponseOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodHour** | Pointer to **NullableString** |  | [optional] 
**StartingDayOfWeek** | Pointer to [**NullableDayOfWeek**](DayOfWeek.md) |  | [optional] 
**StartingTime** | Pointer to **NullableString** |  | [optional] 
**UseMaintenanceOption** | Pointer to **bool** | Maintenance period usage state | [optional] [default to false]

## Methods

### NewMaintenanceResponseOption

`func NewMaintenanceResponseOption() *MaintenanceResponseOption`

NewMaintenanceResponseOption instantiates a new MaintenanceResponseOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceResponseOptionWithDefaults

`func NewMaintenanceResponseOptionWithDefaults() *MaintenanceResponseOption`

NewMaintenanceResponseOptionWithDefaults instantiates a new MaintenanceResponseOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodHour

`func (o *MaintenanceResponseOption) GetPeriodHour() string`

GetPeriodHour returns the PeriodHour field if non-nil, zero value otherwise.

### GetPeriodHourOk

`func (o *MaintenanceResponseOption) GetPeriodHourOk() (*string, bool)`

GetPeriodHourOk returns a tuple with the PeriodHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodHour

`func (o *MaintenanceResponseOption) SetPeriodHour(v string)`

SetPeriodHour sets PeriodHour field to given value.

### HasPeriodHour

`func (o *MaintenanceResponseOption) HasPeriodHour() bool`

HasPeriodHour returns a boolean if a field has been set.

### SetPeriodHourNil

`func (o *MaintenanceResponseOption) SetPeriodHourNil(b bool)`

 SetPeriodHourNil sets the value for PeriodHour to be an explicit nil

### UnsetPeriodHour
`func (o *MaintenanceResponseOption) UnsetPeriodHour()`

UnsetPeriodHour ensures that no value is present for PeriodHour, not even an explicit nil
### GetStartingDayOfWeek

`func (o *MaintenanceResponseOption) GetStartingDayOfWeek() DayOfWeek`

GetStartingDayOfWeek returns the StartingDayOfWeek field if non-nil, zero value otherwise.

### GetStartingDayOfWeekOk

`func (o *MaintenanceResponseOption) GetStartingDayOfWeekOk() (*DayOfWeek, bool)`

GetStartingDayOfWeekOk returns a tuple with the StartingDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingDayOfWeek

`func (o *MaintenanceResponseOption) SetStartingDayOfWeek(v DayOfWeek)`

SetStartingDayOfWeek sets StartingDayOfWeek field to given value.

### HasStartingDayOfWeek

`func (o *MaintenanceResponseOption) HasStartingDayOfWeek() bool`

HasStartingDayOfWeek returns a boolean if a field has been set.

### SetStartingDayOfWeekNil

`func (o *MaintenanceResponseOption) SetStartingDayOfWeekNil(b bool)`

 SetStartingDayOfWeekNil sets the value for StartingDayOfWeek to be an explicit nil

### UnsetStartingDayOfWeek
`func (o *MaintenanceResponseOption) UnsetStartingDayOfWeek()`

UnsetStartingDayOfWeek ensures that no value is present for StartingDayOfWeek, not even an explicit nil
### GetStartingTime

`func (o *MaintenanceResponseOption) GetStartingTime() string`

GetStartingTime returns the StartingTime field if non-nil, zero value otherwise.

### GetStartingTimeOk

`func (o *MaintenanceResponseOption) GetStartingTimeOk() (*string, bool)`

GetStartingTimeOk returns a tuple with the StartingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingTime

`func (o *MaintenanceResponseOption) SetStartingTime(v string)`

SetStartingTime sets StartingTime field to given value.

### HasStartingTime

`func (o *MaintenanceResponseOption) HasStartingTime() bool`

HasStartingTime returns a boolean if a field has been set.

### SetStartingTimeNil

`func (o *MaintenanceResponseOption) SetStartingTimeNil(b bool)`

 SetStartingTimeNil sets the value for StartingTime to be an explicit nil

### UnsetStartingTime
`func (o *MaintenanceResponseOption) UnsetStartingTime()`

UnsetStartingTime ensures that no value is present for StartingTime, not even an explicit nil
### GetUseMaintenanceOption

`func (o *MaintenanceResponseOption) GetUseMaintenanceOption() bool`

GetUseMaintenanceOption returns the UseMaintenanceOption field if non-nil, zero value otherwise.

### GetUseMaintenanceOptionOk

`func (o *MaintenanceResponseOption) GetUseMaintenanceOptionOk() (*bool, bool)`

GetUseMaintenanceOptionOk returns a tuple with the UseMaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMaintenanceOption

`func (o *MaintenanceResponseOption) SetUseMaintenanceOption(v bool)`

SetUseMaintenanceOption sets UseMaintenanceOption field to given value.

### HasUseMaintenanceOption

`func (o *MaintenanceResponseOption) HasUseMaintenanceOption() bool`

HasUseMaintenanceOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


