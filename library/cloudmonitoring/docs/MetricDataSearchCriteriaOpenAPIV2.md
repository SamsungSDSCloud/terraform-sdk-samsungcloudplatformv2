# MetricDataSearchCriteriaOpenAPIV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreInvalid** | Pointer to **string** | 유효하지 않은 값 제외 여부 | [optional] 
**MetricDataConditions** | [**[]MetricDataConditionOpenAPIV2**](MetricDataConditionOpenAPIV2.md) | 메트릭 데이터 조건 | 
**QueryEndDt** | **string** | 조회 종료 일시 | 
**QueryStartDt** | **string** | 조회 시작 일시 | 

## Methods

### NewMetricDataSearchCriteriaOpenAPIV2

`func NewMetricDataSearchCriteriaOpenAPIV2(metricDataConditions []MetricDataConditionOpenAPIV2, queryEndDt string, queryStartDt string, ) *MetricDataSearchCriteriaOpenAPIV2`

NewMetricDataSearchCriteriaOpenAPIV2 instantiates a new MetricDataSearchCriteriaOpenAPIV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDataSearchCriteriaOpenAPIV2WithDefaults

`func NewMetricDataSearchCriteriaOpenAPIV2WithDefaults() *MetricDataSearchCriteriaOpenAPIV2`

NewMetricDataSearchCriteriaOpenAPIV2WithDefaults instantiates a new MetricDataSearchCriteriaOpenAPIV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreInvalid

`func (o *MetricDataSearchCriteriaOpenAPIV2) GetIgnoreInvalid() string`

GetIgnoreInvalid returns the IgnoreInvalid field if non-nil, zero value otherwise.

### GetIgnoreInvalidOk

`func (o *MetricDataSearchCriteriaOpenAPIV2) GetIgnoreInvalidOk() (*string, bool)`

GetIgnoreInvalidOk returns a tuple with the IgnoreInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreInvalid

`func (o *MetricDataSearchCriteriaOpenAPIV2) SetIgnoreInvalid(v string)`

SetIgnoreInvalid sets IgnoreInvalid field to given value.

### HasIgnoreInvalid

`func (o *MetricDataSearchCriteriaOpenAPIV2) HasIgnoreInvalid() bool`

HasIgnoreInvalid returns a boolean if a field has been set.

### GetMetricDataConditions

`func (o *MetricDataSearchCriteriaOpenAPIV2) GetMetricDataConditions() []MetricDataConditionOpenAPIV2`

GetMetricDataConditions returns the MetricDataConditions field if non-nil, zero value otherwise.

### GetMetricDataConditionsOk

`func (o *MetricDataSearchCriteriaOpenAPIV2) GetMetricDataConditionsOk() (*[]MetricDataConditionOpenAPIV2, bool)`

GetMetricDataConditionsOk returns a tuple with the MetricDataConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDataConditions

`func (o *MetricDataSearchCriteriaOpenAPIV2) SetMetricDataConditions(v []MetricDataConditionOpenAPIV2)`

SetMetricDataConditions sets MetricDataConditions field to given value.


### GetQueryEndDt

`func (o *MetricDataSearchCriteriaOpenAPIV2) GetQueryEndDt() string`

GetQueryEndDt returns the QueryEndDt field if non-nil, zero value otherwise.

### GetQueryEndDtOk

`func (o *MetricDataSearchCriteriaOpenAPIV2) GetQueryEndDtOk() (*string, bool)`

GetQueryEndDtOk returns a tuple with the QueryEndDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEndDt

`func (o *MetricDataSearchCriteriaOpenAPIV2) SetQueryEndDt(v string)`

SetQueryEndDt sets QueryEndDt field to given value.


### GetQueryStartDt

`func (o *MetricDataSearchCriteriaOpenAPIV2) GetQueryStartDt() string`

GetQueryStartDt returns the QueryStartDt field if non-nil, zero value otherwise.

### GetQueryStartDtOk

`func (o *MetricDataSearchCriteriaOpenAPIV2) GetQueryStartDtOk() (*string, bool)`

GetQueryStartDtOk returns a tuple with the QueryStartDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStartDt

`func (o *MetricDataSearchCriteriaOpenAPIV2) SetQueryStartDt(v string)`

SetQueryStartDt sets QueryStartDt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


