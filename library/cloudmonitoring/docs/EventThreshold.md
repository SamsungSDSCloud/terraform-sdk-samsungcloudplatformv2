# EventThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricFunction** | Pointer to **string** | 메트릭 함수 - 유효한 값: [delta] - 이벤트 발생 여부를 판단하기 위해 성능항목의 현재 값을 사용할지, 이전 값과 현재 값의 차이를 사용할지를 지정합니다. - 이전 값과 현재 값의 차이를 사용하고자 할 경우에만 delta로 지정합니다. | [optional] 
**RangeThreshold** | Pointer to [**RangeThreshold**](RangeThreshold.md) |  | [optional] 
**SingleThreshold** | Pointer to [**SingleThreshold**](SingleThreshold.md) |  | [optional] 
**ThresholdType** | **string** | 임계치 유형 - 임계치를 범위로 지정할 경우에는 \&quot;RANGE_VALUE\&quot;를, 임계치를 단일값으로 지정할 경우에는 \&quot;SINGLE_VALUE\&quot; 지정해야 합니다. | 

## Methods

### NewEventThreshold

`func NewEventThreshold(thresholdType string, ) *EventThreshold`

NewEventThreshold instantiates a new EventThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventThresholdWithDefaults

`func NewEventThresholdWithDefaults() *EventThreshold`

NewEventThresholdWithDefaults instantiates a new EventThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricFunction

`func (o *EventThreshold) GetMetricFunction() string`

GetMetricFunction returns the MetricFunction field if non-nil, zero value otherwise.

### GetMetricFunctionOk

`func (o *EventThreshold) GetMetricFunctionOk() (*string, bool)`

GetMetricFunctionOk returns a tuple with the MetricFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricFunction

`func (o *EventThreshold) SetMetricFunction(v string)`

SetMetricFunction sets MetricFunction field to given value.

### HasMetricFunction

`func (o *EventThreshold) HasMetricFunction() bool`

HasMetricFunction returns a boolean if a field has been set.

### GetRangeThreshold

`func (o *EventThreshold) GetRangeThreshold() RangeThreshold`

GetRangeThreshold returns the RangeThreshold field if non-nil, zero value otherwise.

### GetRangeThresholdOk

`func (o *EventThreshold) GetRangeThresholdOk() (*RangeThreshold, bool)`

GetRangeThresholdOk returns a tuple with the RangeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeThreshold

`func (o *EventThreshold) SetRangeThreshold(v RangeThreshold)`

SetRangeThreshold sets RangeThreshold field to given value.

### HasRangeThreshold

`func (o *EventThreshold) HasRangeThreshold() bool`

HasRangeThreshold returns a boolean if a field has been set.

### GetSingleThreshold

`func (o *EventThreshold) GetSingleThreshold() SingleThreshold`

GetSingleThreshold returns the SingleThreshold field if non-nil, zero value otherwise.

### GetSingleThresholdOk

`func (o *EventThreshold) GetSingleThresholdOk() (*SingleThreshold, bool)`

GetSingleThresholdOk returns a tuple with the SingleThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleThreshold

`func (o *EventThreshold) SetSingleThreshold(v SingleThreshold)`

SetSingleThreshold sets SingleThreshold field to given value.

### HasSingleThreshold

`func (o *EventThreshold) HasSingleThreshold() bool`

HasSingleThreshold returns a boolean if a field has been set.

### GetThresholdType

`func (o *EventThreshold) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *EventThreshold) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *EventThreshold) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


