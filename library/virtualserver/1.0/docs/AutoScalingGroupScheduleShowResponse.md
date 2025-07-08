# AutoScalingGroupScheduleShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**DayOfMonth** | **NullableString** |  | 
**DayOfWeek** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**DesiredServerCount** | **NullableInt32** |  | 
**EndDate** | **NullableString** |  | 
**Frequency** | [**JobSchedulingFrequencyEnum**](JobSchedulingFrequencyEnum.md) | Frequency | 
**Hour** | **int32** | Hour | 
**Id** | **string** | ID | 
**MaxServerCount** | **NullableInt32** |  | 
**MinServerCount** | **NullableInt32** |  | 
**Minute** | **int32** | Minute | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Schedule name | 
**StartDate** | **string** | Start date | 
**State** | **string** | Schedule state | 
**Timezone** | **string** | Timezone | 

## Methods

### NewAutoScalingGroupScheduleShowResponse

`func NewAutoScalingGroupScheduleShowResponse(accountId string, createdAt time.Time, createdBy string, dayOfMonth NullableString, dayOfWeek NullableString, description NullableString, desiredServerCount NullableInt32, endDate NullableString, frequency JobSchedulingFrequencyEnum, hour int32, id string, maxServerCount NullableInt32, minServerCount NullableInt32, minute int32, modifiedAt time.Time, modifiedBy string, name string, startDate string, state string, timezone string, ) *AutoScalingGroupScheduleShowResponse`

NewAutoScalingGroupScheduleShowResponse instantiates a new AutoScalingGroupScheduleShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupScheduleShowResponseWithDefaults

`func NewAutoScalingGroupScheduleShowResponseWithDefaults() *AutoScalingGroupScheduleShowResponse`

NewAutoScalingGroupScheduleShowResponseWithDefaults instantiates a new AutoScalingGroupScheduleShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AutoScalingGroupScheduleShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AutoScalingGroupScheduleShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AutoScalingGroupScheduleShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *AutoScalingGroupScheduleShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoScalingGroupScheduleShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoScalingGroupScheduleShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AutoScalingGroupScheduleShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AutoScalingGroupScheduleShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AutoScalingGroupScheduleShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDayOfMonth

`func (o *AutoScalingGroupScheduleShowResponse) GetDayOfMonth() string`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *AutoScalingGroupScheduleShowResponse) GetDayOfMonthOk() (*string, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *AutoScalingGroupScheduleShowResponse) SetDayOfMonth(v string)`

SetDayOfMonth sets DayOfMonth field to given value.


### SetDayOfMonthNil

`func (o *AutoScalingGroupScheduleShowResponse) SetDayOfMonthNil(b bool)`

 SetDayOfMonthNil sets the value for DayOfMonth to be an explicit nil

### UnsetDayOfMonth
`func (o *AutoScalingGroupScheduleShowResponse) UnsetDayOfMonth()`

UnsetDayOfMonth ensures that no value is present for DayOfMonth, not even an explicit nil
### GetDayOfWeek

`func (o *AutoScalingGroupScheduleShowResponse) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *AutoScalingGroupScheduleShowResponse) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *AutoScalingGroupScheduleShowResponse) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.


### SetDayOfWeekNil

`func (o *AutoScalingGroupScheduleShowResponse) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *AutoScalingGroupScheduleShowResponse) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetDescription

`func (o *AutoScalingGroupScheduleShowResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutoScalingGroupScheduleShowResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutoScalingGroupScheduleShowResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AutoScalingGroupScheduleShowResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AutoScalingGroupScheduleShowResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDesiredServerCount

`func (o *AutoScalingGroupScheduleShowResponse) GetDesiredServerCount() int32`

GetDesiredServerCount returns the DesiredServerCount field if non-nil, zero value otherwise.

### GetDesiredServerCountOk

`func (o *AutoScalingGroupScheduleShowResponse) GetDesiredServerCountOk() (*int32, bool)`

GetDesiredServerCountOk returns a tuple with the DesiredServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredServerCount

`func (o *AutoScalingGroupScheduleShowResponse) SetDesiredServerCount(v int32)`

SetDesiredServerCount sets DesiredServerCount field to given value.


### SetDesiredServerCountNil

`func (o *AutoScalingGroupScheduleShowResponse) SetDesiredServerCountNil(b bool)`

 SetDesiredServerCountNil sets the value for DesiredServerCount to be an explicit nil

### UnsetDesiredServerCount
`func (o *AutoScalingGroupScheduleShowResponse) UnsetDesiredServerCount()`

UnsetDesiredServerCount ensures that no value is present for DesiredServerCount, not even an explicit nil
### GetEndDate

`func (o *AutoScalingGroupScheduleShowResponse) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AutoScalingGroupScheduleShowResponse) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AutoScalingGroupScheduleShowResponse) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *AutoScalingGroupScheduleShowResponse) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *AutoScalingGroupScheduleShowResponse) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetFrequency

`func (o *AutoScalingGroupScheduleShowResponse) GetFrequency() JobSchedulingFrequencyEnum`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AutoScalingGroupScheduleShowResponse) GetFrequencyOk() (*JobSchedulingFrequencyEnum, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AutoScalingGroupScheduleShowResponse) SetFrequency(v JobSchedulingFrequencyEnum)`

SetFrequency sets Frequency field to given value.


### GetHour

`func (o *AutoScalingGroupScheduleShowResponse) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *AutoScalingGroupScheduleShowResponse) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *AutoScalingGroupScheduleShowResponse) SetHour(v int32)`

SetHour sets Hour field to given value.


### GetId

`func (o *AutoScalingGroupScheduleShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingGroupScheduleShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingGroupScheduleShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMaxServerCount

`func (o *AutoScalingGroupScheduleShowResponse) GetMaxServerCount() int32`

GetMaxServerCount returns the MaxServerCount field if non-nil, zero value otherwise.

### GetMaxServerCountOk

`func (o *AutoScalingGroupScheduleShowResponse) GetMaxServerCountOk() (*int32, bool)`

GetMaxServerCountOk returns a tuple with the MaxServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerCount

`func (o *AutoScalingGroupScheduleShowResponse) SetMaxServerCount(v int32)`

SetMaxServerCount sets MaxServerCount field to given value.


### SetMaxServerCountNil

`func (o *AutoScalingGroupScheduleShowResponse) SetMaxServerCountNil(b bool)`

 SetMaxServerCountNil sets the value for MaxServerCount to be an explicit nil

### UnsetMaxServerCount
`func (o *AutoScalingGroupScheduleShowResponse) UnsetMaxServerCount()`

UnsetMaxServerCount ensures that no value is present for MaxServerCount, not even an explicit nil
### GetMinServerCount

`func (o *AutoScalingGroupScheduleShowResponse) GetMinServerCount() int32`

GetMinServerCount returns the MinServerCount field if non-nil, zero value otherwise.

### GetMinServerCountOk

`func (o *AutoScalingGroupScheduleShowResponse) GetMinServerCountOk() (*int32, bool)`

GetMinServerCountOk returns a tuple with the MinServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinServerCount

`func (o *AutoScalingGroupScheduleShowResponse) SetMinServerCount(v int32)`

SetMinServerCount sets MinServerCount field to given value.


### SetMinServerCountNil

`func (o *AutoScalingGroupScheduleShowResponse) SetMinServerCountNil(b bool)`

 SetMinServerCountNil sets the value for MinServerCount to be an explicit nil

### UnsetMinServerCount
`func (o *AutoScalingGroupScheduleShowResponse) UnsetMinServerCount()`

UnsetMinServerCount ensures that no value is present for MinServerCount, not even an explicit nil
### GetMinute

`func (o *AutoScalingGroupScheduleShowResponse) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *AutoScalingGroupScheduleShowResponse) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *AutoScalingGroupScheduleShowResponse) SetMinute(v int32)`

SetMinute sets Minute field to given value.


### GetModifiedAt

`func (o *AutoScalingGroupScheduleShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AutoScalingGroupScheduleShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AutoScalingGroupScheduleShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *AutoScalingGroupScheduleShowResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *AutoScalingGroupScheduleShowResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *AutoScalingGroupScheduleShowResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *AutoScalingGroupScheduleShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoScalingGroupScheduleShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoScalingGroupScheduleShowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *AutoScalingGroupScheduleShowResponse) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AutoScalingGroupScheduleShowResponse) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AutoScalingGroupScheduleShowResponse) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetState

`func (o *AutoScalingGroupScheduleShowResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroupScheduleShowResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroupScheduleShowResponse) SetState(v string)`

SetState sets State field to given value.


### GetTimezone

`func (o *AutoScalingGroupScheduleShowResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AutoScalingGroupScheduleShowResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AutoScalingGroupScheduleShowResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


