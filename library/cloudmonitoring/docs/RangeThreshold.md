# RangeThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxComparisonOperator** | **string** | 최대 비교 연산자 | 
**MaxValue** | **float64** | 최대 값 | 
**MinComparisonOperator** | **string** | 최소 비교 연산자 | 
**MinValue** | **float64** | 최소 값 | 

## Methods

### NewRangeThreshold

`func NewRangeThreshold(maxComparisonOperator string, maxValue float64, minComparisonOperator string, minValue float64, ) *RangeThreshold`

NewRangeThreshold instantiates a new RangeThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeThresholdWithDefaults

`func NewRangeThresholdWithDefaults() *RangeThreshold`

NewRangeThresholdWithDefaults instantiates a new RangeThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxComparisonOperator

`func (o *RangeThreshold) GetMaxComparisonOperator() string`

GetMaxComparisonOperator returns the MaxComparisonOperator field if non-nil, zero value otherwise.

### GetMaxComparisonOperatorOk

`func (o *RangeThreshold) GetMaxComparisonOperatorOk() (*string, bool)`

GetMaxComparisonOperatorOk returns a tuple with the MaxComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxComparisonOperator

`func (o *RangeThreshold) SetMaxComparisonOperator(v string)`

SetMaxComparisonOperator sets MaxComparisonOperator field to given value.


### GetMaxValue

`func (o *RangeThreshold) GetMaxValue() float64`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *RangeThreshold) GetMaxValueOk() (*float64, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *RangeThreshold) SetMaxValue(v float64)`

SetMaxValue sets MaxValue field to given value.


### GetMinComparisonOperator

`func (o *RangeThreshold) GetMinComparisonOperator() string`

GetMinComparisonOperator returns the MinComparisonOperator field if non-nil, zero value otherwise.

### GetMinComparisonOperatorOk

`func (o *RangeThreshold) GetMinComparisonOperatorOk() (*string, bool)`

GetMinComparisonOperatorOk returns a tuple with the MinComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinComparisonOperator

`func (o *RangeThreshold) SetMinComparisonOperator(v string)`

SetMinComparisonOperator sets MinComparisonOperator field to given value.


### GetMinValue

`func (o *RangeThreshold) GetMinValue() float64`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *RangeThreshold) GetMinValueOk() (*float64, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *RangeThreshold) SetMinValue(v float64)`

SetMinValue sets MinValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


