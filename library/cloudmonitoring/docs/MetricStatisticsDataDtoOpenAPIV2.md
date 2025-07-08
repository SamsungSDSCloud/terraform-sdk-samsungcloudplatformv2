# MetricStatisticsDataDtoOpenAPIV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | Pointer to **string** | 메트릭 키 | [optional] 
**MetricName** | Pointer to **string** | 메트릭 이름 | [optional] 
**MetricType** | Pointer to **string** | 메트릭 유형 | [optional] 
**MetricUnit** | Pointer to **string** | 메트릭 단위 | [optional] 
**ObjectDisplayName** | Pointer to **string** | 개별항목 출력 이름 | [optional] 
**ObjectName** | Pointer to **string** | 개별항목 이름 | [optional] 
**ObjectType** | Pointer to **string** | 개별항목 유형 | [optional] 
**PerfData** | Pointer to **[]map[string]interface{}** | 성능 데이터 | [optional] 
**ProductName** | Pointer to **string** | 상품의 이름 | [optional] 
**ProductResourceId** | Pointer to **string** | 상품 리소스 아이디 | [optional] 
**StatisticsPeriod** | **int32** | 통계 집계 주기 (단위: 분) | 
**StatisticsType** | **string** | 통계 유형 | 

## Methods

### NewMetricStatisticsDataDtoOpenAPIV2

`func NewMetricStatisticsDataDtoOpenAPIV2(statisticsPeriod int32, statisticsType string, ) *MetricStatisticsDataDtoOpenAPIV2`

NewMetricStatisticsDataDtoOpenAPIV2 instantiates a new MetricStatisticsDataDtoOpenAPIV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricStatisticsDataDtoOpenAPIV2WithDefaults

`func NewMetricStatisticsDataDtoOpenAPIV2WithDefaults() *MetricStatisticsDataDtoOpenAPIV2`

NewMetricStatisticsDataDtoOpenAPIV2WithDefaults instantiates a new MetricStatisticsDataDtoOpenAPIV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.

### HasMetricKey

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasMetricKey() bool`

HasMetricKey returns a boolean if a field has been set.

### GetMetricName

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMetricType

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetMetricUnit

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetMetricUnit() string`

GetMetricUnit returns the MetricUnit field if non-nil, zero value otherwise.

### GetMetricUnitOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetMetricUnitOk() (*string, bool)`

GetMetricUnitOk returns a tuple with the MetricUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricUnit

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetMetricUnit(v string)`

SetMetricUnit sets MetricUnit field to given value.

### HasMetricUnit

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasMetricUnit() bool`

HasMetricUnit returns a boolean if a field has been set.

### GetObjectDisplayName

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetObjectDisplayName() string`

GetObjectDisplayName returns the ObjectDisplayName field if non-nil, zero value otherwise.

### GetObjectDisplayNameOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetObjectDisplayNameOk() (*string, bool)`

GetObjectDisplayNameOk returns a tuple with the ObjectDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDisplayName

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetObjectDisplayName(v string)`

SetObjectDisplayName sets ObjectDisplayName field to given value.

### HasObjectDisplayName

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasObjectDisplayName() bool`

HasObjectDisplayName returns a boolean if a field has been set.

### GetObjectName

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPerfData

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetPerfData() []map[string]interface{}`

GetPerfData returns the PerfData field if non-nil, zero value otherwise.

### GetPerfDataOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetPerfDataOk() (*[]map[string]interface{}, bool)`

GetPerfDataOk returns a tuple with the PerfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfData

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetPerfData(v []map[string]interface{})`

SetPerfData sets PerfData field to given value.

### HasPerfData

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasPerfData() bool`

HasPerfData returns a boolean if a field has been set.

### GetProductName

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductResourceId

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.

### HasProductResourceId

`func (o *MetricStatisticsDataDtoOpenAPIV2) HasProductResourceId() bool`

HasProductResourceId returns a boolean if a field has been set.

### GetStatisticsPeriod

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetStatisticsPeriod() int32`

GetStatisticsPeriod returns the StatisticsPeriod field if non-nil, zero value otherwise.

### GetStatisticsPeriodOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetStatisticsPeriodOk() (*int32, bool)`

GetStatisticsPeriodOk returns a tuple with the StatisticsPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsPeriod

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetStatisticsPeriod(v int32)`

SetStatisticsPeriod sets StatisticsPeriod field to given value.


### GetStatisticsType

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetStatisticsType() string`

GetStatisticsType returns the StatisticsType field if non-nil, zero value otherwise.

### GetStatisticsTypeOk

`func (o *MetricStatisticsDataDtoOpenAPIV2) GetStatisticsTypeOk() (*string, bool)`

GetStatisticsTypeOk returns a tuple with the StatisticsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsType

`func (o *MetricStatisticsDataDtoOpenAPIV2) SetStatisticsType(v string)`

SetStatisticsType sets StatisticsType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


