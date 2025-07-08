# EventPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckAsg** | Pointer to **bool** |  | [optional] 
**CreateBy** | Pointer to **map[string]interface{}** |  | [optional] 
**CreateById** | Pointer to **string** |  | [optional] 
**EventLevel** | **string** | 이벤트 레벨 | 
**EventMessagePrefix** | Pointer to **string** | 이벤트 메시지 접두사 | [optional] 
**EventOccurTimeZone** | Pointer to **string** | swagger.event.eventPolicyResponse.eventOccurTimeZone.value | [optional] 
**EventPolicyId** | **int64** | 이벤트 정책 아이디 | 
**EventPolicyStatistics** | Pointer to [**EventPolicyStatistics**](EventPolicyStatistics.md) |  | [optional] 
**EventThreshold** | [**EventThreshold**](EventThreshold.md) |  | 
**FtCount** | **int64** | 결함허용 개수 - 설정한 조건에 맞는 값이 몇 번 반복해서 발생하면 이벤트를 발생시킬지 설정하는 값 | 
**IsLogMetric** | **string** | 로그 메트릭 여부 | 
**MetricDescription** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricDescriptionEn** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricKey** | **string** | 메트릭 키 | 
**MetricName** | **string** | 메트릭 이름 | 
**ObjectName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ProductName** | **string** | 상품 이름 | 
**ProductResourceId** | **string** | 상품 리소스 아이디 | 
**ProductSq** | **int64** | 상품 시퀀스 | 
**ProductTargetType** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**ProductTargetTypeEn** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**UpdateBy** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdateById** | Pointer to **string** |  | [optional] 

## Methods

### NewEventPolicyResponse

`func NewEventPolicyResponse(eventLevel string, eventPolicyId int64, eventThreshold EventThreshold, ftCount int64, isLogMetric string, metricKey string, metricName string, productName string, productResourceId string, productSq int64, ) *EventPolicyResponse`

NewEventPolicyResponse instantiates a new EventPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPolicyResponseWithDefaults

`func NewEventPolicyResponseWithDefaults() *EventPolicyResponse`

NewEventPolicyResponseWithDefaults instantiates a new EventPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckAsg

`func (o *EventPolicyResponse) GetCheckAsg() bool`

GetCheckAsg returns the CheckAsg field if non-nil, zero value otherwise.

### GetCheckAsgOk

`func (o *EventPolicyResponse) GetCheckAsgOk() (*bool, bool)`

GetCheckAsgOk returns a tuple with the CheckAsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAsg

`func (o *EventPolicyResponse) SetCheckAsg(v bool)`

SetCheckAsg sets CheckAsg field to given value.

### HasCheckAsg

`func (o *EventPolicyResponse) HasCheckAsg() bool`

HasCheckAsg returns a boolean if a field has been set.

### GetCreateBy

`func (o *EventPolicyResponse) GetCreateBy() map[string]interface{}`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *EventPolicyResponse) GetCreateByOk() (*map[string]interface{}, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *EventPolicyResponse) SetCreateBy(v map[string]interface{})`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *EventPolicyResponse) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetCreateById

`func (o *EventPolicyResponse) GetCreateById() string`

GetCreateById returns the CreateById field if non-nil, zero value otherwise.

### GetCreateByIdOk

`func (o *EventPolicyResponse) GetCreateByIdOk() (*string, bool)`

GetCreateByIdOk returns a tuple with the CreateById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateById

`func (o *EventPolicyResponse) SetCreateById(v string)`

SetCreateById sets CreateById field to given value.

### HasCreateById

`func (o *EventPolicyResponse) HasCreateById() bool`

HasCreateById returns a boolean if a field has been set.

### GetEventLevel

`func (o *EventPolicyResponse) GetEventLevel() string`

GetEventLevel returns the EventLevel field if non-nil, zero value otherwise.

### GetEventLevelOk

`func (o *EventPolicyResponse) GetEventLevelOk() (*string, bool)`

GetEventLevelOk returns a tuple with the EventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLevel

`func (o *EventPolicyResponse) SetEventLevel(v string)`

SetEventLevel sets EventLevel field to given value.


### GetEventMessagePrefix

`func (o *EventPolicyResponse) GetEventMessagePrefix() string`

GetEventMessagePrefix returns the EventMessagePrefix field if non-nil, zero value otherwise.

### GetEventMessagePrefixOk

`func (o *EventPolicyResponse) GetEventMessagePrefixOk() (*string, bool)`

GetEventMessagePrefixOk returns a tuple with the EventMessagePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessagePrefix

`func (o *EventPolicyResponse) SetEventMessagePrefix(v string)`

SetEventMessagePrefix sets EventMessagePrefix field to given value.

### HasEventMessagePrefix

`func (o *EventPolicyResponse) HasEventMessagePrefix() bool`

HasEventMessagePrefix returns a boolean if a field has been set.

### GetEventOccurTimeZone

`func (o *EventPolicyResponse) GetEventOccurTimeZone() string`

GetEventOccurTimeZone returns the EventOccurTimeZone field if non-nil, zero value otherwise.

### GetEventOccurTimeZoneOk

`func (o *EventPolicyResponse) GetEventOccurTimeZoneOk() (*string, bool)`

GetEventOccurTimeZoneOk returns a tuple with the EventOccurTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOccurTimeZone

`func (o *EventPolicyResponse) SetEventOccurTimeZone(v string)`

SetEventOccurTimeZone sets EventOccurTimeZone field to given value.

### HasEventOccurTimeZone

`func (o *EventPolicyResponse) HasEventOccurTimeZone() bool`

HasEventOccurTimeZone returns a boolean if a field has been set.

### GetEventPolicyId

`func (o *EventPolicyResponse) GetEventPolicyId() int64`

GetEventPolicyId returns the EventPolicyId field if non-nil, zero value otherwise.

### GetEventPolicyIdOk

`func (o *EventPolicyResponse) GetEventPolicyIdOk() (*int64, bool)`

GetEventPolicyIdOk returns a tuple with the EventPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyId

`func (o *EventPolicyResponse) SetEventPolicyId(v int64)`

SetEventPolicyId sets EventPolicyId field to given value.


### GetEventPolicyStatistics

`func (o *EventPolicyResponse) GetEventPolicyStatistics() EventPolicyStatistics`

GetEventPolicyStatistics returns the EventPolicyStatistics field if non-nil, zero value otherwise.

### GetEventPolicyStatisticsOk

`func (o *EventPolicyResponse) GetEventPolicyStatisticsOk() (*EventPolicyStatistics, bool)`

GetEventPolicyStatisticsOk returns a tuple with the EventPolicyStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyStatistics

`func (o *EventPolicyResponse) SetEventPolicyStatistics(v EventPolicyStatistics)`

SetEventPolicyStatistics sets EventPolicyStatistics field to given value.

### HasEventPolicyStatistics

`func (o *EventPolicyResponse) HasEventPolicyStatistics() bool`

HasEventPolicyStatistics returns a boolean if a field has been set.

### GetEventThreshold

`func (o *EventPolicyResponse) GetEventThreshold() EventThreshold`

GetEventThreshold returns the EventThreshold field if non-nil, zero value otherwise.

### GetEventThresholdOk

`func (o *EventPolicyResponse) GetEventThresholdOk() (*EventThreshold, bool)`

GetEventThresholdOk returns a tuple with the EventThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThreshold

`func (o *EventPolicyResponse) SetEventThreshold(v EventThreshold)`

SetEventThreshold sets EventThreshold field to given value.


### GetFtCount

`func (o *EventPolicyResponse) GetFtCount() int64`

GetFtCount returns the FtCount field if non-nil, zero value otherwise.

### GetFtCountOk

`func (o *EventPolicyResponse) GetFtCountOk() (*int64, bool)`

GetFtCountOk returns a tuple with the FtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtCount

`func (o *EventPolicyResponse) SetFtCount(v int64)`

SetFtCount sets FtCount field to given value.


### GetIsLogMetric

`func (o *EventPolicyResponse) GetIsLogMetric() string`

GetIsLogMetric returns the IsLogMetric field if non-nil, zero value otherwise.

### GetIsLogMetricOk

`func (o *EventPolicyResponse) GetIsLogMetricOk() (*string, bool)`

GetIsLogMetricOk returns a tuple with the IsLogMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLogMetric

`func (o *EventPolicyResponse) SetIsLogMetric(v string)`

SetIsLogMetric sets IsLogMetric field to given value.


### GetMetricDescription

`func (o *EventPolicyResponse) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *EventPolicyResponse) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *EventPolicyResponse) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *EventPolicyResponse) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricDescriptionEn

`func (o *EventPolicyResponse) GetMetricDescriptionEn() string`

GetMetricDescriptionEn returns the MetricDescriptionEn field if non-nil, zero value otherwise.

### GetMetricDescriptionEnOk

`func (o *EventPolicyResponse) GetMetricDescriptionEnOk() (*string, bool)`

GetMetricDescriptionEnOk returns a tuple with the MetricDescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescriptionEn

`func (o *EventPolicyResponse) SetMetricDescriptionEn(v string)`

SetMetricDescriptionEn sets MetricDescriptionEn field to given value.

### HasMetricDescriptionEn

`func (o *EventPolicyResponse) HasMetricDescriptionEn() bool`

HasMetricDescriptionEn returns a boolean if a field has been set.

### GetMetricKey

`func (o *EventPolicyResponse) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *EventPolicyResponse) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *EventPolicyResponse) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *EventPolicyResponse) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *EventPolicyResponse) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *EventPolicyResponse) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetObjectName

`func (o *EventPolicyResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *EventPolicyResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *EventPolicyResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *EventPolicyResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetProductName

`func (o *EventPolicyResponse) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *EventPolicyResponse) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *EventPolicyResponse) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetProductResourceId

`func (o *EventPolicyResponse) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *EventPolicyResponse) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *EventPolicyResponse) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.


### GetProductSq

`func (o *EventPolicyResponse) GetProductSq() int64`

GetProductSq returns the ProductSq field if non-nil, zero value otherwise.

### GetProductSqOk

`func (o *EventPolicyResponse) GetProductSqOk() (*int64, bool)`

GetProductSqOk returns a tuple with the ProductSq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSq

`func (o *EventPolicyResponse) SetProductSq(v int64)`

SetProductSq sets ProductSq field to given value.


### GetProductTargetType

`func (o *EventPolicyResponse) GetProductTargetType() string`

GetProductTargetType returns the ProductTargetType field if non-nil, zero value otherwise.

### GetProductTargetTypeOk

`func (o *EventPolicyResponse) GetProductTargetTypeOk() (*string, bool)`

GetProductTargetTypeOk returns a tuple with the ProductTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetType

`func (o *EventPolicyResponse) SetProductTargetType(v string)`

SetProductTargetType sets ProductTargetType field to given value.

### HasProductTargetType

`func (o *EventPolicyResponse) HasProductTargetType() bool`

HasProductTargetType returns a boolean if a field has been set.

### GetProductTargetTypeEn

`func (o *EventPolicyResponse) GetProductTargetTypeEn() string`

GetProductTargetTypeEn returns the ProductTargetTypeEn field if non-nil, zero value otherwise.

### GetProductTargetTypeEnOk

`func (o *EventPolicyResponse) GetProductTargetTypeEnOk() (*string, bool)`

GetProductTargetTypeEnOk returns a tuple with the ProductTargetTypeEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetTypeEn

`func (o *EventPolicyResponse) SetProductTargetTypeEn(v string)`

SetProductTargetTypeEn sets ProductTargetTypeEn field to given value.

### HasProductTargetTypeEn

`func (o *EventPolicyResponse) HasProductTargetTypeEn() bool`

HasProductTargetTypeEn returns a boolean if a field has been set.

### GetUpdateBy

`func (o *EventPolicyResponse) GetUpdateBy() map[string]interface{}`

GetUpdateBy returns the UpdateBy field if non-nil, zero value otherwise.

### GetUpdateByOk

`func (o *EventPolicyResponse) GetUpdateByOk() (*map[string]interface{}, bool)`

GetUpdateByOk returns a tuple with the UpdateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateBy

`func (o *EventPolicyResponse) SetUpdateBy(v map[string]interface{})`

SetUpdateBy sets UpdateBy field to given value.

### HasUpdateBy

`func (o *EventPolicyResponse) HasUpdateBy() bool`

HasUpdateBy returns a boolean if a field has been set.

### GetUpdateById

`func (o *EventPolicyResponse) GetUpdateById() string`

GetUpdateById returns the UpdateById field if non-nil, zero value otherwise.

### GetUpdateByIdOk

`func (o *EventPolicyResponse) GetUpdateByIdOk() (*string, bool)`

GetUpdateByIdOk returns a tuple with the UpdateById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateById

`func (o *EventPolicyResponse) SetUpdateById(v string)`

SetUpdateById sets UpdateById field to given value.

### HasUpdateById

`func (o *EventPolicyResponse) HasUpdateById() bool`

HasUpdateById returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


