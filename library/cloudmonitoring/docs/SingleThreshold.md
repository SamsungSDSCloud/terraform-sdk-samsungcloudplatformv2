# SingleThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonOperator** | **string** | 비교 연산자 | 
**Value** | **float64** | 값 | 

## Methods

### NewSingleThreshold

`func NewSingleThreshold(comparisonOperator string, value float64, ) *SingleThreshold`

NewSingleThreshold instantiates a new SingleThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleThresholdWithDefaults

`func NewSingleThresholdWithDefaults() *SingleThreshold`

NewSingleThresholdWithDefaults instantiates a new SingleThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonOperator

`func (o *SingleThreshold) GetComparisonOperator() string`

GetComparisonOperator returns the ComparisonOperator field if non-nil, zero value otherwise.

### GetComparisonOperatorOk

`func (o *SingleThreshold) GetComparisonOperatorOk() (*string, bool)`

GetComparisonOperatorOk returns a tuple with the ComparisonOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonOperator

`func (o *SingleThreshold) SetComparisonOperator(v string)`

SetComparisonOperator sets ComparisonOperator field to given value.


### GetValue

`func (o *SingleThreshold) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SingleThreshold) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SingleThreshold) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


