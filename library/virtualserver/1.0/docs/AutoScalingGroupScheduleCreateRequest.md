# AutoScalingGroupScheduleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **NullableString** |  | [optional] 
**DayOfWeek** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DesiredServerCount** | Pointer to **NullableInt32** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**Frequency** | [**JobSchedulingFrequencyEnum**](JobSchedulingFrequencyEnum.md) | Frequency | 
**Hour** | **int32** | Hour | 
**MaxServerCount** | Pointer to **NullableInt32** |  | [optional] 
**MinServerCount** | Pointer to **NullableInt32** |  | [optional] 
**Minute** | **int32** | Minute | 
**Name** | **string** | Schedule name | 
**StartDate** | **string** | Start date | 
**State** | Pointer to [**JobSchedulingState**](JobSchedulingState.md) | Schedule state | [optional] 
**Timezone** | **string** | Timezone | 

## Methods

### NewAutoScalingGroupScheduleCreateRequest

`func NewAutoScalingGroupScheduleCreateRequest(frequency JobSchedulingFrequencyEnum, hour int32, minute int32, name string, startDate string, timezone string, ) *AutoScalingGroupScheduleCreateRequest`

NewAutoScalingGroupScheduleCreateRequest instantiates a new AutoScalingGroupScheduleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupScheduleCreateRequestWithDefaults

`func NewAutoScalingGroupScheduleCreateRequestWithDefaults() *AutoScalingGroupScheduleCreateRequest`

NewAutoScalingGroupScheduleCreateRequestWithDefaults instantiates a new AutoScalingGroupScheduleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *AutoScalingGroupScheduleCreateRequest) GetDayOfMonth() string`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetDayOfMonthOk() (*string, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *AutoScalingGroupScheduleCreateRequest) SetDayOfMonth(v string)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *AutoScalingGroupScheduleCreateRequest) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### SetDayOfMonthNil

`func (o *AutoScalingGroupScheduleCreateRequest) SetDayOfMonthNil(b bool)`

 SetDayOfMonthNil sets the value for DayOfMonth to be an explicit nil

### UnsetDayOfMonth
`func (o *AutoScalingGroupScheduleCreateRequest) UnsetDayOfMonth()`

UnsetDayOfMonth ensures that no value is present for DayOfMonth, not even an explicit nil
### GetDayOfWeek

`func (o *AutoScalingGroupScheduleCreateRequest) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *AutoScalingGroupScheduleCreateRequest) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *AutoScalingGroupScheduleCreateRequest) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *AutoScalingGroupScheduleCreateRequest) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *AutoScalingGroupScheduleCreateRequest) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetDescription

`func (o *AutoScalingGroupScheduleCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoScalingGroupScheduleCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoScalingGroupScheduleCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoScalingGroupScheduleCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoScalingGroupScheduleCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesiredServerCount

`func (o *AutoScalingGroupScheduleCreateRequest) GetDesiredServerCount() int32`

GetDesiredServerCount returns the DesiredServerCount field if non-nil, zero value otherwise.

### GetDesiredServerCountOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetDesiredServerCountOk() (*int32, bool)`

GetDesiredServerCountOk returns a tuple with the DesiredServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredServerCount

`func (o *AutoScalingGroupScheduleCreateRequest) SetDesiredServerCount(v int32)`

SetDesiredServerCount sets DesiredServerCount field to given value.

### HasDesiredServerCount

`func (o *AutoScalingGroupScheduleCreateRequest) HasDesiredServerCount() bool`

HasDesiredServerCount returns a boolean if a field has been set.

### SetDesiredServerCountNil

`func (o *AutoScalingGroupScheduleCreateRequest) SetDesiredServerCountNil(b bool)`

 SetDesiredServerCountNil sets the value for DesiredServerCount to be an explicit nil

### UnsetDesiredServerCount
`func (o *AutoScalingGroupScheduleCreateRequest) UnsetDesiredServerCount()`

UnsetDesiredServerCount ensures that no value is present for DesiredServerCount, not even an explicit nil
### GetEndDate

`func (o *AutoScalingGroupScheduleCreateRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AutoScalingGroupScheduleCreateRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AutoScalingGroupScheduleCreateRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *AutoScalingGroupScheduleCreateRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *AutoScalingGroupScheduleCreateRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetFrequency

`func (o *AutoScalingGroupScheduleCreateRequest) GetFrequency() JobSchedulingFrequencyEnum`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetFrequencyOk() (*JobSchedulingFrequencyEnum, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AutoScalingGroupScheduleCreateRequest) SetFrequency(v JobSchedulingFrequencyEnum)`

SetFrequency sets Frequency field to given value.


### GetHour

`func (o *AutoScalingGroupScheduleCreateRequest) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *AutoScalingGroupScheduleCreateRequest) SetHour(v int32)`

SetHour sets Hour field to given value.


### GetMaxServerCount

`func (o *AutoScalingGroupScheduleCreateRequest) GetMaxServerCount() int32`

GetMaxServerCount returns the MaxServerCount field if non-nil, zero value otherwise.

### GetMaxServerCountOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetMaxServerCountOk() (*int32, bool)`

GetMaxServerCountOk returns a tuple with the MaxServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerCount

`func (o *AutoScalingGroupScheduleCreateRequest) SetMaxServerCount(v int32)`

SetMaxServerCount sets MaxServerCount field to given value.

### HasMaxServerCount

`func (o *AutoScalingGroupScheduleCreateRequest) HasMaxServerCount() bool`

HasMaxServerCount returns a boolean if a field has been set.

### SetMaxServerCountNil

`func (o *AutoScalingGroupScheduleCreateRequest) SetMaxServerCountNil(b bool)`

 SetMaxServerCountNil sets the value for MaxServerCount to be an explicit nil

### UnsetMaxServerCount
`func (o *AutoScalingGroupScheduleCreateRequest) UnsetMaxServerCount()`

UnsetMaxServerCount ensures that no value is present for MaxServerCount, not even an explicit nil
### GetMinServerCount

`func (o *AutoScalingGroupScheduleCreateRequest) GetMinServerCount() int32`

GetMinServerCount returns the MinServerCount field if non-nil, zero value otherwise.

### GetMinServerCountOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetMinServerCountOk() (*int32, bool)`

GetMinServerCountOk returns a tuple with the MinServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinServerCount

`func (o *AutoScalingGroupScheduleCreateRequest) SetMinServerCount(v int32)`

SetMinServerCount sets MinServerCount field to given value.

### HasMinServerCount

`func (o *AutoScalingGroupScheduleCreateRequest) HasMinServerCount() bool`

HasMinServerCount returns a boolean if a field has been set.

### SetMinServerCountNil

`func (o *AutoScalingGroupScheduleCreateRequest) SetMinServerCountNil(b bool)`

 SetMinServerCountNil sets the value for MinServerCount to be an explicit nil

### UnsetMinServerCount
`func (o *AutoScalingGroupScheduleCreateRequest) UnsetMinServerCount()`

UnsetMinServerCount ensures that no value is present for MinServerCount, not even an explicit nil
### GetMinute

`func (o *AutoScalingGroupScheduleCreateRequest) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *AutoScalingGroupScheduleCreateRequest) SetMinute(v int32)`

SetMinute sets Minute field to given value.


### GetName

`func (o *AutoScalingGroupScheduleCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoScalingGroupScheduleCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *AutoScalingGroupScheduleCreateRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AutoScalingGroupScheduleCreateRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetState

`func (o *AutoScalingGroupScheduleCreateRequest) GetState() JobSchedulingState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetStateOk() (*JobSchedulingState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroupScheduleCreateRequest) SetState(v JobSchedulingState)`

SetState sets State field to given value.

### HasState

`func (o *AutoScalingGroupScheduleCreateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimezone

`func (o *AutoScalingGroupScheduleCreateRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AutoScalingGroupScheduleCreateRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AutoScalingGroupScheduleCreateRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


