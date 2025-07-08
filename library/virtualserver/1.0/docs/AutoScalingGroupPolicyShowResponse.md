# AutoScalingGroupPolicyShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AutoScalingGroupId** | **string** | Auto-Scaling Group ID | 
**ComparisonOperator** | **string** | Comparison operator | 
**CooldownSeconds** | **int32** | Cooldown seconds | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**EvaluationMinutes** | **int32** | Evaluation minutes | 
**Id** | **string** | ID | 
**MetricMethod** | **string** | Metric method | 
**MetricType** | **string** | Metric type | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Policy name | 
**ScaleMethod** | **string** | Scale method | 
**ScaleType** | **string** | Scale type | 
**ScaleValue** | **int32** | Scale value | 
**State** | **string** | Policy state | 
**Threshold** | **int32** | Threshold | 
**ThresholdUnit** | **string** | Threshold unit | 

## Methods

### NewAutoScalingGroupPolicyShowResponse

`func NewAutoScalingGroupPolicyShowResponse(accountId string, autoScalingGroupId string, comparisonOperator string, cooldownSeconds int32, createdAt time.Time, createdBy string, evaluationMinutes int32, id string, metricMethod string, metricType string, modifiedAt time.Time, modifiedBy string, name string, scaleMethod string, scaleType string, scaleValue int32, state string, threshold int32, thresholdUnit string, ) *AutoScalingGroupPolicyShowResponse`

NewAutoScalingGroupPolicyShowResponse instantiates a new AutoScalingGroupPolicyShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupPolicyShowResponseWithDefaults

`func NewAutoScalingGroupPolicyShowResponseWithDefaults() *AutoScalingGroupPolicyShowResponse`

NewAutoScalingGroupPolicyShowResponseWithDefaults instantiates a new AutoScalingGroupPolicyShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AutoScalingGroupPolicyShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AutoScalingGroupPolicyShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AutoScalingGroupPolicyShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAutoScalingGroupId

`func (o *AutoScalingGroupPolicyShowResponse) GetAutoScalingGroupId() string`

GetAutoScalingGroupId returns the AutoScalingGroupId field if non-nil, zero value otherwise.

### GetAutoScalingGroupIdOk

`func (o *AutoScalingGroupPolicyShowResponse) GetAutoScalingGroupIdOk() (*string, bool)`

GetAutoScalingGroupIdOk returns a tuple with the AutoScalingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroupId

`func (o *AutoScalingGroupPolicyShowResponse) SetAutoScalingGroupId(v string)`

SetAutoScalingGroupId sets AutoScalingGroupId field to given value.


### GetComparisonOperator

`func (o *AutoScalingGroupPolicyShowResponse) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *AutoScalingGroupPolicyShowResponse) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *AutoScalingGroupPolicyShowResponse) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.


### GetCooldownSeconds

`func (o *AutoScalingGroupPolicyShowResponse) GetCooldownSeconds() int32`

GetCooldownSeconds returns the CooldownSeconds field if non-nil, zero value otherwise.

### GetCooldownSecondsOk

`func (o *AutoScalingGroupPolicyShowResponse) GetCooldownSecondsOk() (*int32, bool)`

GetCooldownSecondsOk returns a tuple with the CooldownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownSeconds

`func (o *AutoScalingGroupPolicyShowResponse) SetCooldownSeconds(v int32)`

SetCooldownSeconds sets CooldownSeconds field to given value.


### GetCreatedAt

`func (o *AutoScalingGroupPolicyShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoScalingGroupPolicyShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoScalingGroupPolicyShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AutoScalingGroupPolicyShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AutoScalingGroupPolicyShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AutoScalingGroupPolicyShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetEvaluationMinutes

`func (o *AutoScalingGroupPolicyShowResponse) GetEvaluationMinutes() int32`

GetEvaluationMinutes returns the EvaluationMinutes field if non-nil, zero value otherwise.

### GetEvaluationMinutesOk

`func (o *AutoScalingGroupPolicyShowResponse) GetEvaluationMinutesOk() (*int32, bool)`

GetEvaluationMinutesOk returns a tuple with the EvaluationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMinutes

`func (o *AutoScalingGroupPolicyShowResponse) SetEvaluationMinutes(v int32)`

SetEvaluationMinutes sets EvaluationMinutes field to given value.


### GetId

`func (o *AutoScalingGroupPolicyShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingGroupPolicyShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingGroupPolicyShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMetricMethod

`func (o *AutoScalingGroupPolicyShowResponse) GetMetricMethod() string`

GetMetricMethod returns the MetricMethod field if non-nil, zero value otherwise.

### GetMetricMethodOk

`func (o *AutoScalingGroupPolicyShowResponse) GetMetricMethodOk() (*string, bool)`

GetMetricMethodOk returns a tuple with the MetricMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricMethod

`func (o *AutoScalingGroupPolicyShowResponse) SetMetricMethod(v string)`

SetMetricMethod sets MetricMethod field to given value.


### GetMetricType

`func (o *AutoScalingGroupPolicyShowResponse) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *AutoScalingGroupPolicyShowResponse) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *AutoScalingGroupPolicyShowResponse) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetModifiedAt

`func (o *AutoScalingGroupPolicyShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AutoScalingGroupPolicyShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AutoScalingGroupPolicyShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *AutoScalingGroupPolicyShowResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *AutoScalingGroupPolicyShowResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *AutoScalingGroupPolicyShowResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *AutoScalingGroupPolicyShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoScalingGroupPolicyShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoScalingGroupPolicyShowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetScaleMethod

`func (o *AutoScalingGroupPolicyShowResponse) GetScaleMethod() string`

GetScaleMethod returns the ScaleMethod field if non-nil, zero value otherwise.

### GetScaleMethodOk

`func (o *AutoScalingGroupPolicyShowResponse) GetScaleMethodOk() (*string, bool)`

GetScaleMethodOk returns a tuple with the ScaleMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleMethod

`func (o *AutoScalingGroupPolicyShowResponse) SetScaleMethod(v string)`

SetScaleMethod sets ScaleMethod field to given value.


### GetScaleType

`func (o *AutoScalingGroupPolicyShowResponse) GetScaleType() string`

GetScaleType returns the ScaleType field if non-nil, zero value otherwise.

### GetScaleTypeOk

`func (o *AutoScalingGroupPolicyShowResponse) GetScaleTypeOk() (*string, bool)`

GetScaleTypeOk returns a tuple with the ScaleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleType

`func (o *AutoScalingGroupPolicyShowResponse) SetScaleType(v string)`

SetScaleType sets ScaleType field to given value.


### GetScaleValue

`func (o *AutoScalingGroupPolicyShowResponse) GetScaleValue() int32`

GetScaleValue returns the ScaleValue field if non-nil, zero value otherwise.

### GetScaleValueOk

`func (o *AutoScalingGroupPolicyShowResponse) GetScaleValueOk() (*int32, bool)`

GetScaleValueOk returns a tuple with the ScaleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleValue

`func (o *AutoScalingGroupPolicyShowResponse) SetScaleValue(v int32)`

SetScaleValue sets ScaleValue field to given value.


### GetState

`func (o *AutoScalingGroupPolicyShowResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroupPolicyShowResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroupPolicyShowResponse) SetState(v string)`

SetState sets State field to given value.


### GetThreshold

`func (o *AutoScalingGroupPolicyShowResponse) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AutoScalingGroupPolicyShowResponse) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AutoScalingGroupPolicyShowResponse) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetThresholdUnit

`func (o *AutoScalingGroupPolicyShowResponse) GetThresholdUnit() string`

GetThresholdUnit returns the ThresholdUnit field if non-nil, zero value otherwise.

### GetThresholdUnitOk

`func (o *AutoScalingGroupPolicyShowResponse) GetThresholdUnitOk() (*string, bool)`

GetThresholdUnitOk returns a tuple with the ThresholdUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdUnit

`func (o *AutoScalingGroupPolicyShowResponse) SetThresholdUnit(v string)`

SetThresholdUnit sets ThresholdUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


