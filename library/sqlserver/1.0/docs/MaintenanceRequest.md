# MaintenanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDayOfWeek** | [**DayOfWeek**](DayOfWeek.md) | start_day_of_week | 
**StartMinute** | **string** | start_minute | 
**StartTime** | **string** | start_time | 
**TermHour** | **string** | term_hour | 

## Methods

### NewMaintenanceRequest

`func NewMaintenanceRequest(startDayOfWeek DayOfWeek, startMinute string, startTime string, termHour string, ) *MaintenanceRequest`

NewMaintenanceRequest instantiates a new MaintenanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceRequestWithDefaults

`func NewMaintenanceRequestWithDefaults() *MaintenanceRequest`

NewMaintenanceRequestWithDefaults instantiates a new MaintenanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDayOfWeek

`func (o *MaintenanceRequest) GetStartDayOfWeek() DayOfWeek`

GetStartDayOfWeek returns the StartDayOfWeek field if non-nil, zero value otherwise.

### GetStartDayOfWeekOk

`func (o *MaintenanceRequest) GetStartDayOfWeekOk() (*DayOfWeek, bool)`

GetStartDayOfWeekOk returns a tuple with the StartDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDayOfWeek

`func (o *MaintenanceRequest) SetStartDayOfWeek(v DayOfWeek)`

SetStartDayOfWeek sets StartDayOfWeek field to given value.


### GetStartMinute

`func (o *MaintenanceRequest) GetStartMinute() string`

GetStartMinute returns the StartMinute field if non-nil, zero value otherwise.

### GetStartMinuteOk

`func (o *MaintenanceRequest) GetStartMinuteOk() (*string, bool)`

GetStartMinuteOk returns a tuple with the StartMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMinute

`func (o *MaintenanceRequest) SetStartMinute(v string)`

SetStartMinute sets StartMinute field to given value.


### GetStartTime

`func (o *MaintenanceRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MaintenanceRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MaintenanceRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetTermHour

`func (o *MaintenanceRequest) GetTermHour() string`

GetTermHour returns the TermHour field if non-nil, zero value otherwise.

### GetTermHourOk

`func (o *MaintenanceRequest) GetTermHourOk() (*string, bool)`

GetTermHourOk returns a tuple with the TermHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermHour

`func (o *MaintenanceRequest) SetTermHour(v string)`

SetTermHour sets TermHour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


