# AutoScalingGroupPolicyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonOperator** | **string** | Comparison operator | 
**CooldownSeconds** | **int32** | Cooldown seconds | 
**EvaluationMinutes** | **int32** | Evaluation minutes | 
**MetricMethod** | **string** | Metric method | 
**MetricType** | **string** | Metric type | 
**Name** | **string** | Policy name | 
**ScaleMethod** | **string** | Scale method | 
**ScaleType** | **string** | Scale type | 
**ScaleValue** | **int32** | Scale value | 
**Threshold** | **int32** | Threshold | 

## Methods

### NewAutoScalingGroupPolicyCreateRequest

`func NewAutoScalingGroupPolicyCreateRequest(comparisonOperator string, cooldownSeconds int32, evaluationMinutes int32, metricMethod string, metricType string, name string, scaleMethod string, scaleType string, scaleValue int32, threshold int32, ) *AutoScalingGroupPolicyCreateRequest`

NewAutoScalingGroupPolicyCreateRequest instantiates a new AutoScalingGroupPolicyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupPolicyCreateRequestWithDefaults

`func NewAutoScalingGroupPolicyCreateRequestWithDefaults() *AutoScalingGroupPolicyCreateRequest`

NewAutoScalingGroupPolicyCreateRequestWithDefaults instantiates a new AutoScalingGroupPolicyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonOperator

`func (o *AutoScalingGroupPolicyCreateRequest) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *AutoScalingGroupPolicyCreateRequest) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.


### GetCooldownSeconds

`func (o *AutoScalingGroupPolicyCreateRequest) GetCooldownSeconds() int32`

GetCooldownSeconds returns the CooldownSeconds field if non-nil, zero value otherwise.

### GetCooldownSecondsOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetCooldownSecondsOk() (*int32, bool)`

GetCooldownSecondsOk returns a tuple with the CooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownSeconds

`func (o *AutoScalingGroupPolicyCreateRequest) SetCooldownSeconds(v int32)`

SetCooldownSeconds sets CooldownSeconds field to given value.


### GetEvaluationMinutes

`func (o *AutoScalingGroupPolicyCreateRequest) GetEvaluationMinutes() int32`

GetEvaluationMinutes returns the EvaluationMinutes field if non-nil, zero value otherwise.

### GetEvaluationMinutesOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetEvaluationMinutesOk() (*int32, bool)`

GetEvaluationMinutesOk returns a tuple with the EvaluationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMinutes

`func (o *AutoScalingGroupPolicyCreateRequest) SetEvaluationMinutes(v int32)`

SetEvaluationMinutes sets EvaluationMinutes field to given value.


### GetMetricMethod

`func (o *AutoScalingGroupPolicyCreateRequest) GetMetricMethod() string`

GetMetricMethod returns the MetricMethod field if non-nil, zero value otherwise.

### GetMetricMethodOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetMetricMethodOk() (*string, bool)`

GetMetricMethodOk returns a tuple with the MetricMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricMethod

`func (o *AutoScalingGroupPolicyCreateRequest) SetMetricMethod(v string)`

SetMetricMethod sets MetricMethod field to given value.


### GetMetricType

`func (o *AutoScalingGroupPolicyCreateRequest) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *AutoScalingGroupPolicyCreateRequest) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetName

`func (o *AutoScalingGroupPolicyCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoScalingGroupPolicyCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScaleMethod

`func (o *AutoScalingGroupPolicyCreateRequest) GetScaleMethod() string`

GetScaleMethod returns the ScaleMethod field if non-nil, zero value otherwise.

### GetScaleMethodOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetScaleMethodOk() (*string, bool)`

GetScaleMethodOk returns a tuple with the ScaleMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleMethod

`func (o *AutoScalingGroupPolicyCreateRequest) SetScaleMethod(v string)`

SetScaleMethod sets ScaleMethod field to given value.


### GetScaleType

`func (o *AutoScalingGroupPolicyCreateRequest) GetScaleType() string`

GetScaleType returns the ScaleType field if non-nil, zero value otherwise.

### GetScaleTypeOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetScaleTypeOk() (*string, bool)`

GetScaleTypeOk returns a tuple with the ScaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleType

`func (o *AutoScalingGroupPolicyCreateRequest) SetScaleType(v string)`

SetScaleType sets ScaleType field to given value.


### GetScaleValue

`func (o *AutoScalingGroupPolicyCreateRequest) GetScaleValue() int32`

GetScaleValue returns the ScaleValue field if non-nil, zero value otherwise.

### GetScaleValueOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetScaleValueOk() (*int32, bool)`

GetScaleValueOk returns a tuple with the ScaleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleValue

`func (o *AutoScalingGroupPolicyCreateRequest) SetScaleValue(v int32)`

SetScaleValue sets ScaleValue field to given value.


### GetThreshold

`func (o *AutoScalingGroupPolicyCreateRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AutoScalingGroupPolicyCreateRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AutoScalingGroupPolicyCreateRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


