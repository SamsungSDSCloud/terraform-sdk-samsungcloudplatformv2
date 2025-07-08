# MetricDataConditionOpenAPIV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricKey** | **string** | 메트릭 | 
**ObjectType** | Pointer to **string** | 개별항목 유형 | [optional] 
**ProductResourceInfos** | [**[]ProductResourceInfo**](ProductResourceInfo.md) | 상품 리소스 아이디 목록 - 상품 리소스 아이디는 50개까지 지정할 수 있습니다. | 
**StatisticsPeriod** | Pointer to **int32** | 통계 집계 주기 (단위: 분) | [optional] 
**StatisticsTypeList** | Pointer to **[]string** | 통계 유형 - 유효한 값 : MIN, MAX, AVG, SUM | [optional] 

## Methods

### NewMetricDataConditionOpenAPIV2

`func NewMetricDataConditionOpenAPIV2(metricKey string, productResourceInfos []ProductResourceInfo, ) *MetricDataConditionOpenAPIV2`

NewMetricDataConditionOpenAPIV2 instantiates a new MetricDataConditionOpenAPIV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDataConditionOpenAPIV2WithDefaults

`func NewMetricDataConditionOpenAPIV2WithDefaults() *MetricDataConditionOpenAPIV2`

NewMetricDataConditionOpenAPIV2WithDefaults instantiates a new MetricDataConditionOpenAPIV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricKey

`func (o *MetricDataConditionOpenAPIV2) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *MetricDataConditionOpenAPIV2) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *MetricDataConditionOpenAPIV2) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetObjectType

`func (o *MetricDataConditionOpenAPIV2) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricDataConditionOpenAPIV2) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricDataConditionOpenAPIV2) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *MetricDataConditionOpenAPIV2) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetProductResourceInfos

`func (o *MetricDataConditionOpenAPIV2) GetProductResourceInfos() []ProductResourceInfo`

GetProductResourceInfos returns the ProductResourceInfos field if non-nil, zero value otherwise.

### GetProductResourceInfosOk

`func (o *MetricDataConditionOpenAPIV2) GetProductResourceInfosOk() (*[]ProductResourceInfo, bool)`

GetProductResourceInfosOk returns a tuple with the ProductResourceInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceInfos

`func (o *MetricDataConditionOpenAPIV2) SetProductResourceInfos(v []ProductResourceInfo)`

SetProductResourceInfos sets ProductResourceInfos field to given value.


### GetStatisticsPeriod

`func (o *MetricDataConditionOpenAPIV2) GetStatisticsPeriod() int32`

GetStatisticsPeriod returns the StatisticsPeriod field if non-nil, zero value otherwise.

### GetStatisticsPeriodOk

`func (o *MetricDataConditionOpenAPIV2) GetStatisticsPeriodOk() (*int32, bool)`

GetStatisticsPeriodOk returns a tuple with the StatisticsPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsPeriod

`func (o *MetricDataConditionOpenAPIV2) SetStatisticsPeriod(v int32)`

SetStatisticsPeriod sets StatisticsPeriod field to given value.

### HasStatisticsPeriod

`func (o *MetricDataConditionOpenAPIV2) HasStatisticsPeriod() bool`

HasStatisticsPeriod returns a boolean if a field has been set.

### GetStatisticsTypeList

`func (o *MetricDataConditionOpenAPIV2) GetStatisticsTypeList() []string`

GetStatisticsTypeList returns the StatisticsTypeList field if non-nil, zero value otherwise.

### GetStatisticsTypeListOk

`func (o *MetricDataConditionOpenAPIV2) GetStatisticsTypeListOk() (*[]string, bool)`

GetStatisticsTypeListOk returns a tuple with the StatisticsTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsTypeList

`func (o *MetricDataConditionOpenAPIV2) SetStatisticsTypeList(v []string)`

SetStatisticsTypeList sets StatisticsTypeList field to given value.

### HasStatisticsTypeList

`func (o *MetricDataConditionOpenAPIV2) HasStatisticsTypeList() bool`

HasStatisticsTypeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


