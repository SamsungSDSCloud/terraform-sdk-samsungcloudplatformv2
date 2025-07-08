# AutoScalingGroupPolicyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonOperator** | Pointer to **NullableString** |  | [optional] 
**CooldownSeconds** | Pointer to **NullableInt32** |  | [optional] 
**EvaluationMinutes** | Pointer to **NullableInt32** |  | [optional] 
**MetricMethod** | Pointer to **NullableString** |  | [optional] 
**MetricType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ScaleMethod** | Pointer to **NullableString** |  | [optional] 
**ScaleType** | Pointer to **NullableString** |  | [optional] 
**ScaleValue** | Pointer to **NullableInt32** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**Threshold** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAutoScalingGroupPolicyUpdateRequest

`func NewAutoScalingGroupPolicyUpdateRequest() *AutoScalingGroupPolicyUpdateRequest`

NewAutoScalingGroupPolicyUpdateRequest instantiates a new AutoScalingGroupPolicyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupPolicyUpdateRequestWithDefaults

`func NewAutoScalingGroupPolicyUpdateRequestWithDefaults() *AutoScalingGroupPolicyUpdateRequest`

NewAutoScalingGroupPolicyUpdateRequestWithDefaults instantiates a new AutoScalingGroupPolicyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonOperator

`func (o *AutoScalingGroupPolicyUpdateRequest) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *AutoScalingGroupPolicyUpdateRequest) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.

### HasComparisonOperator

`func (o *AutoScalingGroupPolicyUpdateRequest) HasComparisonOperator() bool`

HasComparisonOperator returns a boolean if a field has been set.

### SetComparisonOperatorNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetComparisonOperatorNil(b bool)`

 SetComparisonOperatorNil sets the value for ComparisonOperator to be an explicit nil

### UnsetComparisonOperator
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetComparisonOperator()`

UnsetComparisonOperator ensures that no value is present for ComparisonOperator, not even an explicit nil
### GetCooldownSeconds

`func (o *AutoScalingGroupPolicyUpdateRequest) GetCooldownSeconds() int32`

GetCooldownSeconds returns the CooldownSeconds field if non-nil, zero value otherwise.

### GetCooldownSecondsOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetCooldownSecondsOk() (*int32, bool)`

GetCooldownSecondsOk returns a tuple with the CooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownSeconds

`func (o *AutoScalingGroupPolicyUpdateRequest) SetCooldownSeconds(v int32)`

SetCooldownSeconds sets CooldownSeconds field to given value.

### HasCooldownSeconds

`func (o *AutoScalingGroupPolicyUpdateRequest) HasCooldownSeconds() bool`

HasCooldownSeconds returns a boolean if a field has been set.

### SetCooldownSecondsNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetCooldownSecondsNil(b bool)`

 SetCooldownSecondsNil sets the value for CooldownSeconds to be an explicit nil

### UnsetCooldownSeconds
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetCooldownSeconds()`

UnsetCooldownSeconds ensures that no value is present for CooldownSeconds, not even an explicit nil
### GetEvaluationMinutes

`func (o *AutoScalingGroupPolicyUpdateRequest) GetEvaluationMinutes() int32`

GetEvaluationMinutes returns the EvaluationMinutes field if non-nil, zero value otherwise.

### GetEvaluationMinutesOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetEvaluationMinutesOk() (*int32, bool)`

GetEvaluationMinutesOk returns a tuple with the EvaluationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMinutes

`func (o *AutoScalingGroupPolicyUpdateRequest) SetEvaluationMinutes(v int32)`

SetEvaluationMinutes sets EvaluationMinutes field to given value.

### HasEvaluationMinutes

`func (o *AutoScalingGroupPolicyUpdateRequest) HasEvaluationMinutes() bool`

HasEvaluationMinutes returns a boolean if a field has been set.

### SetEvaluationMinutesNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetEvaluationMinutesNil(b bool)`

 SetEvaluationMinutesNil sets the value for EvaluationMinutes to be an explicit nil

### UnsetEvaluationMinutes
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetEvaluationMinutes()`

UnsetEvaluationMinutes ensures that no value is present for EvaluationMinutes, not even an explicit nil
### GetMetricMethod

`func (o *AutoScalingGroupPolicyUpdateRequest) GetMetricMethod() string`

GetMetricMethod returns the MetricMethod field if non-nil, zero value otherwise.

### GetMetricMethodOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetMetricMethodOk() (*string, bool)`

GetMetricMethodOk returns a tuple with the MetricMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricMethod

`func (o *AutoScalingGroupPolicyUpdateRequest) SetMetricMethod(v string)`

SetMetricMethod sets MetricMethod field to given value.

### HasMetricMethod

`func (o *AutoScalingGroupPolicyUpdateRequest) HasMetricMethod() bool`

HasMetricMethod returns a boolean if a field has been set.

### SetMetricMethodNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetMetricMethodNil(b bool)`

 SetMetricMethodNil sets the value for MetricMethod to be an explicit nil

### UnsetMetricMethod
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetMetricMethod()`

UnsetMetricMethod ensures that no value is present for MetricMethod, not even an explicit nil
### GetMetricType

`func (o *AutoScalingGroupPolicyUpdateRequest) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *AutoScalingGroupPolicyUpdateRequest) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *AutoScalingGroupPolicyUpdateRequest) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### SetMetricTypeNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetMetricTypeNil(b bool)`

 SetMetricTypeNil sets the value for MetricType to be an explicit nil

### UnsetMetricType
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetMetricType()`

UnsetMetricType ensures that no value is present for MetricType, not even an explicit nil
### GetName

`func (o *AutoScalingGroupPolicyUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoScalingGroupPolicyUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutoScalingGroupPolicyUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetScaleMethod

`func (o *AutoScalingGroupPolicyUpdateRequest) GetScaleMethod() string`

GetScaleMethod returns the ScaleMethod field if non-nil, zero value otherwise.

### GetScaleMethodOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetScaleMethodOk() (*string, bool)`

GetScaleMethodOk returns a tuple with the ScaleMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleMethod

`func (o *AutoScalingGroupPolicyUpdateRequest) SetScaleMethod(v string)`

SetScaleMethod sets ScaleMethod field to given value.

### HasScaleMethod

`func (o *AutoScalingGroupPolicyUpdateRequest) HasScaleMethod() bool`

HasScaleMethod returns a boolean if a field has been set.

### SetScaleMethodNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetScaleMethodNil(b bool)`

 SetScaleMethodNil sets the value for ScaleMethod to be an explicit nil

### UnsetScaleMethod
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetScaleMethod()`

UnsetScaleMethod ensures that no value is present for ScaleMethod, not even an explicit nil
### GetScaleType

`func (o *AutoScalingGroupPolicyUpdateRequest) GetScaleType() string`

GetScaleType returns the ScaleType field if non-nil, zero value otherwise.

### GetScaleTypeOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetScaleTypeOk() (*string, bool)`

GetScaleTypeOk returns a tuple with the ScaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleType

`func (o *AutoScalingGroupPolicyUpdateRequest) SetScaleType(v string)`

SetScaleType sets ScaleType field to given value.

### HasScaleType

`func (o *AutoScalingGroupPolicyUpdateRequest) HasScaleType() bool`

HasScaleType returns a boolean if a field has been set.

### SetScaleTypeNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetScaleTypeNil(b bool)`

 SetScaleTypeNil sets the value for ScaleType to be an explicit nil

### UnsetScaleType
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetScaleType()`

UnsetScaleType ensures that no value is present for ScaleType, not even an explicit nil
### GetScaleValue

`func (o *AutoScalingGroupPolicyUpdateRequest) GetScaleValue() int32`

GetScaleValue returns the ScaleValue field if non-nil, zero value otherwise.

### GetScaleValueOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetScaleValueOk() (*int32, bool)`

GetScaleValueOk returns a tuple with the ScaleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleValue

`func (o *AutoScalingGroupPolicyUpdateRequest) SetScaleValue(v int32)`

SetScaleValue sets ScaleValue field to given value.

### HasScaleValue

`func (o *AutoScalingGroupPolicyUpdateRequest) HasScaleValue() bool`

HasScaleValue returns a boolean if a field has been set.

### SetScaleValueNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetScaleValueNil(b bool)`

 SetScaleValueNil sets the value for ScaleValue to be an explicit nil

### UnsetScaleValue
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetScaleValue()`

UnsetScaleValue ensures that no value is present for ScaleValue, not even an explicit nil
### GetState

`func (o *AutoScalingGroupPolicyUpdateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroupPolicyUpdateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AutoScalingGroupPolicyUpdateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetThreshold

`func (o *AutoScalingGroupPolicyUpdateRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AutoScalingGroupPolicyUpdateRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AutoScalingGroupPolicyUpdateRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *AutoScalingGroupPolicyUpdateRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### SetThresholdNil

`func (o *AutoScalingGroupPolicyUpdateRequest) SetThresholdNil(b bool)`

 SetThresholdNil sets the value for Threshold to be an explicit nil

### UnsetThreshold
`func (o *AutoScalingGroupPolicyUpdateRequest) UnsetThreshold()`

UnsetThreshold ensures that no value is present for Threshold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


