# MaintenanceOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodHour** | Pointer to **string** | Period in hours | [optional] 
**StartingDayOfWeek** | Pointer to [**DayOfWeek**](DayOfWeek.md) | Starting day of week | [optional] 
**StartingTime** | Pointer to **string** | Starting time | [optional] 

## Methods

### NewMaintenanceOption

`func NewMaintenanceOption() *MaintenanceOption`

NewMaintenanceOption instantiates a new MaintenanceOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceOptionWithDefaults

`func NewMaintenanceOptionWithDefaults() *MaintenanceOption`

NewMaintenanceOptionWithDefaults instantiates a new MaintenanceOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodHour

`func (o *MaintenanceOption) GetPeriodHour() string`

GetPeriodHour returns the PeriodHour field if non-nil, zero value otherwise.

### GetPeriodHourOk

`func (o *MaintenanceOption) GetPeriodHourOk() (*string, bool)`

GetPeriodHourOk returns a tuple with the PeriodHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodHour

`func (o *MaintenanceOption) SetPeriodHour(v string)`

SetPeriodHour sets PeriodHour field to given value.

### HasPeriodHour

`func (o *MaintenanceOption) HasPeriodHour() bool`

HasPeriodHour returns a boolean if a field has been set.

### GetStartingDayOfWeek

`func (o *MaintenanceOption) GetStartingDayOfWeek() DayOfWeek`

GetStartingDayOfWeek returns the StartingDayOfWeek field if non-nil, zero value otherwise.

### GetStartingDayOfWeekOk

`func (o *MaintenanceOption) GetStartingDayOfWeekOk() (*DayOfWeek, bool)`

GetStartingDayOfWeekOk returns a tuple with the StartingDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingDayOfWeek

`func (o *MaintenanceOption) SetStartingDayOfWeek(v DayOfWeek)`

SetStartingDayOfWeek sets StartingDayOfWeek field to given value.

### HasStartingDayOfWeek

`func (o *MaintenanceOption) HasStartingDayOfWeek() bool`

HasStartingDayOfWeek returns a boolean if a field has been set.

### GetStartingTime

`func (o *MaintenanceOption) GetStartingTime() string`

GetStartingTime returns the StartingTime field if non-nil, zero value otherwise.

### GetStartingTimeOk

`func (o *MaintenanceOption) GetStartingTimeOk() (*string, bool)`

GetStartingTimeOk returns a tuple with the StartingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingTime

`func (o *MaintenanceOption) SetStartingTime(v string)`

SetStartingTime sets StartingTime field to given value.

### HasStartingTime

`func (o *MaintenanceOption) HasStartingTime() bool`

HasStartingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


