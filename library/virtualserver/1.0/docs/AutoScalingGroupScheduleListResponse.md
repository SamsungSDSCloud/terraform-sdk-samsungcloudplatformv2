# AutoScalingGroupScheduleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Schedules** | [**[]AutoScalingGroupScheduleShowResponse**](AutoScalingGroupScheduleShowResponse.md) |  | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAutoScalingGroupScheduleListResponse

`func NewAutoScalingGroupScheduleListResponse(count int32, page int32, schedules []AutoScalingGroupScheduleShowResponse, size int32, ) *AutoScalingGroupScheduleListResponse`

NewAutoScalingGroupScheduleListResponse instantiates a new AutoScalingGroupScheduleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupScheduleListResponseWithDefaults

`func NewAutoScalingGroupScheduleListResponseWithDefaults() *AutoScalingGroupScheduleListResponse`

NewAutoScalingGroupScheduleListResponseWithDefaults instantiates a new AutoScalingGroupScheduleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AutoScalingGroupScheduleListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AutoScalingGroupScheduleListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AutoScalingGroupScheduleListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *AutoScalingGroupScheduleListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AutoScalingGroupScheduleListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AutoScalingGroupScheduleListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSchedules

`func (o *AutoScalingGroupScheduleListResponse) GetSchedules() []AutoScalingGroupScheduleShowResponse`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *AutoScalingGroupScheduleListResponse) GetSchedulesOk() (*[]AutoScalingGroupScheduleShowResponse, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *AutoScalingGroupScheduleListResponse) SetSchedules(v []AutoScalingGroupScheduleShowResponse)`

SetSchedules sets Schedules field to given value.


### GetSize

`func (o *AutoScalingGroupScheduleListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AutoScalingGroupScheduleListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AutoScalingGroupScheduleListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *AutoScalingGroupScheduleListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AutoScalingGroupScheduleListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AutoScalingGroupScheduleListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AutoScalingGroupScheduleListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *AutoScalingGroupScheduleListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *AutoScalingGroupScheduleListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


