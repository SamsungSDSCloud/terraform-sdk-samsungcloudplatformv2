# AutoScalingGroupScheduleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfMonth** | Pointer to **NullableString** |  | [optional] 
**DayOfWeek** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DesiredServerCount** | Pointer to **NullableInt32** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 
**Frequency** | Pointer to [**NullableJobSchedulingFrequencyEnum**](JobSchedulingFrequencyEnum.md) |  | [optional] 
**Hour** | Pointer to **NullableInt32** |  | [optional] 
**MaxServerCount** | Pointer to **NullableInt32** |  | [optional] 
**MinServerCount** | Pointer to **NullableInt32** |  | [optional] 
**Minute** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to [**JobSchedulingState**](JobSchedulingState.md) | Schedule state | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAutoScalingGroupScheduleUpdateRequest

`func NewAutoScalingGroupScheduleUpdateRequest() *AutoScalingGroupScheduleUpdateRequest`

NewAutoScalingGroupScheduleUpdateRequest instantiates a new AutoScalingGroupScheduleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupScheduleUpdateRequestWithDefaults

`func NewAutoScalingGroupScheduleUpdateRequestWithDefaults() *AutoScalingGroupScheduleUpdateRequest`

NewAutoScalingGroupScheduleUpdateRequestWithDefaults instantiates a new AutoScalingGroupScheduleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfMonth

`func (o *AutoScalingGroupScheduleUpdateRequest) GetDayOfMonth() string`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetDayOfMonthOk() (*string, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *AutoScalingGroupScheduleUpdateRequest) SetDayOfMonth(v string)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *AutoScalingGroupScheduleUpdateRequest) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### SetDayOfMonthNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetDayOfMonthNil(b bool)`

 SetDayOfMonthNil sets the value for DayOfMonth to be an explicit nil

### UnsetDayOfMonth
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetDayOfMonth()`

UnsetDayOfMonth ensures that no value is present for DayOfMonth, not even an explicit nil
### GetDayOfWeek

`func (o *AutoScalingGroupScheduleUpdateRequest) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *AutoScalingGroupScheduleUpdateRequest) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *AutoScalingGroupScheduleUpdateRequest) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetDescription

`func (o *AutoScalingGroupScheduleUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoScalingGroupScheduleUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutoScalingGroupScheduleUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesiredServerCount

`func (o *AutoScalingGroupScheduleUpdateRequest) GetDesiredServerCount() int32`

GetDesiredServerCount returns the DesiredServerCount field if non-nil, zero value otherwise.

### GetDesiredServerCountOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetDesiredServerCountOk() (*int32, bool)`

GetDesiredServerCountOk returns a tuple with the DesiredServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredServerCount

`func (o *AutoScalingGroupScheduleUpdateRequest) SetDesiredServerCount(v int32)`

SetDesiredServerCount sets DesiredServerCount field to given value.

### HasDesiredServerCount

`func (o *AutoScalingGroupScheduleUpdateRequest) HasDesiredServerCount() bool`

HasDesiredServerCount returns a boolean if a field has been set.

### SetDesiredServerCountNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetDesiredServerCountNil(b bool)`

 SetDesiredServerCountNil sets the value for DesiredServerCount to be an explicit nil

### UnsetDesiredServerCount
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetDesiredServerCount()`

UnsetDesiredServerCount ensures that no value is present for DesiredServerCount, not even an explicit nil
### GetEndDate

`func (o *AutoScalingGroupScheduleUpdateRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AutoScalingGroupScheduleUpdateRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AutoScalingGroupScheduleUpdateRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetFrequency

`func (o *AutoScalingGroupScheduleUpdateRequest) GetFrequency() JobSchedulingFrequencyEnum`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetFrequencyOk() (*JobSchedulingFrequencyEnum, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AutoScalingGroupScheduleUpdateRequest) SetFrequency(v JobSchedulingFrequencyEnum)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *AutoScalingGroupScheduleUpdateRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetHour

`func (o *AutoScalingGroupScheduleUpdateRequest) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *AutoScalingGroupScheduleUpdateRequest) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *AutoScalingGroupScheduleUpdateRequest) HasHour() bool`

HasHour returns a boolean if a field has been set.

### SetHourNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetHourNil(b bool)`

 SetHourNil sets the value for Hour to be an explicit nil

### UnsetHour
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetHour()`

UnsetHour ensures that no value is present for Hour, not even an explicit nil
### GetMaxServerCount

`func (o *AutoScalingGroupScheduleUpdateRequest) GetMaxServerCount() int32`

GetMaxServerCount returns the MaxServerCount field if non-nil, zero value otherwise.

### GetMaxServerCountOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetMaxServerCountOk() (*int32, bool)`

GetMaxServerCountOk returns a tuple with the MaxServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerCount

`func (o *AutoScalingGroupScheduleUpdateRequest) SetMaxServerCount(v int32)`

SetMaxServerCount sets MaxServerCount field to given value.

### HasMaxServerCount

`func (o *AutoScalingGroupScheduleUpdateRequest) HasMaxServerCount() bool`

HasMaxServerCount returns a boolean if a field has been set.

### SetMaxServerCountNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetMaxServerCountNil(b bool)`

 SetMaxServerCountNil sets the value for MaxServerCount to be an explicit nil

### UnsetMaxServerCount
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetMaxServerCount()`

UnsetMaxServerCount ensures that no value is present for MaxServerCount, not even an explicit nil
### GetMinServerCount

`func (o *AutoScalingGroupScheduleUpdateRequest) GetMinServerCount() int32`

GetMinServerCount returns the MinServerCount field if non-nil, zero value otherwise.

### GetMinServerCountOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetMinServerCountOk() (*int32, bool)`

GetMinServerCountOk returns a tuple with the MinServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinServerCount

`func (o *AutoScalingGroupScheduleUpdateRequest) SetMinServerCount(v int32)`

SetMinServerCount sets MinServerCount field to given value.

### HasMinServerCount

`func (o *AutoScalingGroupScheduleUpdateRequest) HasMinServerCount() bool`

HasMinServerCount returns a boolean if a field has been set.

### SetMinServerCountNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetMinServerCountNil(b bool)`

 SetMinServerCountNil sets the value for MinServerCount to be an explicit nil

### UnsetMinServerCount
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetMinServerCount()`

UnsetMinServerCount ensures that no value is present for MinServerCount, not even an explicit nil
### GetMinute

`func (o *AutoScalingGroupScheduleUpdateRequest) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *AutoScalingGroupScheduleUpdateRequest) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *AutoScalingGroupScheduleUpdateRequest) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### SetMinuteNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetMinuteNil(b bool)`

 SetMinuteNil sets the value for Minute to be an explicit nil

### UnsetMinute
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetMinute()`

UnsetMinute ensures that no value is present for Minute, not even an explicit nil
### GetName

`func (o *AutoScalingGroupScheduleUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoScalingGroupScheduleUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoScalingGroupScheduleUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartDate

`func (o *AutoScalingGroupScheduleUpdateRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AutoScalingGroupScheduleUpdateRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AutoScalingGroupScheduleUpdateRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetState

`func (o *AutoScalingGroupScheduleUpdateRequest) GetState() JobSchedulingState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetStateOk() (*JobSchedulingState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroupScheduleUpdateRequest) SetState(v JobSchedulingState)`

SetState sets State field to given value.

### HasState

`func (o *AutoScalingGroupScheduleUpdateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimezone

`func (o *AutoScalingGroupScheduleUpdateRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AutoScalingGroupScheduleUpdateRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AutoScalingGroupScheduleUpdateRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AutoScalingGroupScheduleUpdateRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *AutoScalingGroupScheduleUpdateRequest) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *AutoScalingGroupScheduleUpdateRequest) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


