# MetricSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLogMetric** | **string** | 로그 메트릭 여부 | 
**IsObjectExist** | **string** | 오브젝트 존재여부 | 
**MetricDescription** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricDescriptionEn** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricKey** | **string** | 메트릭 키 | 
**MetricName** | **string** | 메트릭 이름 | 
**MetricSetKey** | **string** | 메트릭 셋 키 | 
**MetricSetName** | **string** | 메트릭 셋 이름 | 
**MetricType** | **string** | 메트릭 유형 | 
**MetricUnit** | **string** | 메트릭 단위 | 
**ProductTargetType** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**ProductTargetTypeEn** | Pointer to **string** | 메트릭 대상 유형 | [optional] 

## Methods

### NewMetricSummary

`func NewMetricSummary(isLogMetric string, isObjectExist string, metricKey string, metricName string, metricSetKey string, metricSetName string, metricType string, metricUnit string, ) *MetricSummary`

NewMetricSummary instantiates a new MetricSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricSummaryWithDefaults

`func NewMetricSummaryWithDefaults() *MetricSummary`

NewMetricSummaryWithDefaults instantiates a new MetricSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLogMetric

`func (o *MetricSummary) GetIsLogMetric() string`

GetIsLogMetric returns the IsLogMetric field if non-nil, zero value otherwise.

### GetIsLogMetricOk

`func (o *MetricSummary) GetIsLogMetricOk() (*string, bool)`

GetIsLogMetricOk returns a tuple with the IsLogMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLogMetric

`func (o *MetricSummary) SetIsLogMetric(v string)`

SetIsLogMetric sets IsLogMetric field to given value.


### GetIsObjectExist

`func (o *MetricSummary) GetIsObjectExist() string`

GetIsObjectExist returns the IsObjectExist field if non-nil, zero value otherwise.

### GetIsObjectExistOk

`func (o *MetricSummary) GetIsObjectExistOk() (*string, bool)`

GetIsObjectExistOk returns a tuple with the IsObjectExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObjectExist

`func (o *MetricSummary) SetIsObjectExist(v string)`

SetIsObjectExist sets IsObjectExist field to given value.


### GetMetricDescription

`func (o *MetricSummary) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *MetricSummary) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *MetricSummary) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *MetricSummary) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricDescriptionEn

`func (o *MetricSummary) GetMetricDescriptionEn() string`

GetMetricDescriptionEn returns the MetricDescriptionEn field if non-nil, zero value otherwise.

### GetMetricDescriptionEnOk

`func (o *MetricSummary) GetMetricDescriptionEnOk() (*string, bool)`

GetMetricDescriptionEnOk returns a tuple with the MetricDescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescriptionEn

`func (o *MetricSummary) SetMetricDescriptionEn(v string)`

SetMetricDescriptionEn sets MetricDescriptionEn field to given value.

### HasMetricDescriptionEn

`func (o *MetricSummary) HasMetricDescriptionEn() bool`

HasMetricDescriptionEn returns a boolean if a field has been set.

### GetMetricKey

`func (o *MetricSummary) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *MetricSummary) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *MetricSummary) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *MetricSummary) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricSummary) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricSummary) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMetricSetKey

`func (o *MetricSummary) GetMetricSetKey() string`

GetMetricSetKey returns the MetricSetKey field if non-nil, zero value otherwise.

### GetMetricSetKeyOk

`func (o *MetricSummary) GetMetricSetKeyOk() (*string, bool)`

GetMetricSetKeyOk returns a tuple with the MetricSetKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSetKey

`func (o *MetricSummary) SetMetricSetKey(v string)`

SetMetricSetKey sets MetricSetKey field to given value.


### GetMetricSetName

`func (o *MetricSummary) GetMetricSetName() string`

GetMetricSetName returns the MetricSetName field if non-nil, zero value otherwise.

### GetMetricSetNameOk

`func (o *MetricSummary) GetMetricSetNameOk() (*string, bool)`

GetMetricSetNameOk returns a tuple with the MetricSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSetName

`func (o *MetricSummary) SetMetricSetName(v string)`

SetMetricSetName sets MetricSetName field to given value.


### GetMetricType

`func (o *MetricSummary) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *MetricSummary) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *MetricSummary) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetMetricUnit

`func (o *MetricSummary) GetMetricUnit() string`

GetMetricUnit returns the MetricUnit field if non-nil, zero value otherwise.

### GetMetricUnitOk

`func (o *MetricSummary) GetMetricUnitOk() (*string, bool)`

GetMetricUnitOk returns a tuple with the MetricUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricUnit

`func (o *MetricSummary) SetMetricUnit(v string)`

SetMetricUnit sets MetricUnit field to given value.


### GetProductTargetType

`func (o *MetricSummary) GetProductTargetType() string`

GetProductTargetType returns the ProductTargetType field if non-nil, zero value otherwise.

### GetProductTargetTypeOk

`func (o *MetricSummary) GetProductTargetTypeOk() (*string, bool)`

GetProductTargetTypeOk returns a tuple with the ProductTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetType

`func (o *MetricSummary) SetProductTargetType(v string)`

SetProductTargetType sets ProductTargetType field to given value.

### HasProductTargetType

`func (o *MetricSummary) HasProductTargetType() bool`

HasProductTargetType returns a boolean if a field has been set.

### GetProductTargetTypeEn

`func (o *MetricSummary) GetProductTargetTypeEn() string`

GetProductTargetTypeEn returns the ProductTargetTypeEn field if non-nil, zero value otherwise.

### GetProductTargetTypeEnOk

`func (o *MetricSummary) GetProductTargetTypeEnOk() (*string, bool)`

GetProductTargetTypeEnOk returns a tuple with the ProductTargetTypeEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetTypeEn

`func (o *MetricSummary) SetProductTargetTypeEn(v string)`

SetProductTargetTypeEn sets ProductTargetTypeEn field to given value.

### HasProductTargetTypeEn

`func (o *MetricSummary) HasProductTargetTypeEn() bool`

HasProductTargetTypeEn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


