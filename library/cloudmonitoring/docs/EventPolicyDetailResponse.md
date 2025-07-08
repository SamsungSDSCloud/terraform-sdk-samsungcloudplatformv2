# EventPolicyDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsgYn** | Pointer to **string** |  | [optional] 
**AttrListStr** | Pointer to **string** |  | [optional] 
**CheckAsg** | Pointer to **bool** |  | [optional] 
**CreateBy** | Pointer to **map[string]interface{}** |  | [optional] 
**CreateById** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByName** | Pointer to **string** |  | [optional] 
**CreatedDt** | Pointer to **time.Time** |  | [optional] 
**DisableObject** | Pointer to **string** | swagger.event.eventPolicyResponse.disableObject.value | [optional] 
**DisableYn** | Pointer to **string** |  | [optional] 
**DisplayEventRule** | Pointer to **string** | swagger.event.eventResponse.displayEventRule.value | [optional] 
**EventLevel** | **string** | 이벤트 레벨 | 
**EventMessagePrefix** | Pointer to **string** | 이벤트 메시지 접두사 | [optional] 
**EventOccurTimeZone** | Pointer to **string** | swagger.event.eventPolicyResponse.eventOccurTimeZone.value | [optional] 
**EventPolicyId** | **int64** | 이벤트 정책 아이디 | 
**EventPolicyStatistics** | Pointer to [**EventPolicyStatistics**](EventPolicyStatistics.md) |  | [optional] 
**EventThreshold** | [**EventThreshold**](EventThreshold.md) |  | 
**FtCount** | **int64** | 결함허용 개수 - 설정한 조건에 맞는 값이 몇 번 반복해서 발생하면 이벤트를 발생시킬지 설정하는 값 | 
**GroupSummary** | Pointer to [**GroupSummary**](GroupSummary.md) |  | [optional] 
**IsLogMetric** | **string** | 로그 메트릭 여부 | 
**MetricDescription** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricDescriptionEn** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricKey** | **string** | 메트릭 키 | 
**MetricName** | Pointer to **string** | 메트릭 이름 | [optional] 
**MetricSummary** | [**MetricSummary**](MetricSummary.md) |  | 
**MetricType** | Pointer to **string** | 메트릭 유형 | [optional] 
**MetricUnit** | Pointer to **string** | 메트릭 단위 | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**ModifiedByName** | Pointer to **string** |  | [optional] 
**ModifiedDt** | Pointer to **time.Time** |  | [optional] 
**ObjectDisplayName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectType** | Pointer to **string** | 개별항목 유형 | [optional] 
**ObjectTypeName** | Pointer to **string** | 개별항목 유형 이름 | [optional] 
**ProductInfoAttrs** | Pointer to [**[]ProductInfoAttr**](ProductInfoAttr.md) | 상품 정보 속성 | [optional] 
**ProductName** | Pointer to **string** | 상품 이름 | [optional] 
**ProductResourceId** | Pointer to **string** | 상품 리소스 아이디 | [optional] 
**ProductSq** | Pointer to **int64** | 상품 시퀀스 | [optional] 
**ProductSummary** | [**ProductSummary**](ProductSummary.md) |  | 
**ProductTargetType** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**ProductTargetTypeEn** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**StartDt** | Pointer to **time.Time** | 시작일시 | [optional] 
**UpdateBy** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdateById** | Pointer to **string** |  | [optional] 
**UserNameList** | Pointer to **[]string** | 사용자 이름 | [optional] 
**UserNameStr** | Pointer to **string** |  | [optional] 
**UserNames** | Pointer to **string** | swagger.event.eventPolicyResponse.disableObject.value | [optional] 

## Methods

### NewEventPolicyDetailResponse

`func NewEventPolicyDetailResponse(eventLevel string, eventPolicyId int64, eventThreshold EventThreshold, ftCount int64, isLogMetric string, metricKey string, metricSummary MetricSummary, productSummary ProductSummary, ) *EventPolicyDetailResponse`

NewEventPolicyDetailResponse instantiates a new EventPolicyDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPolicyDetailResponseWithDefaults

`func NewEventPolicyDetailResponseWithDefaults() *EventPolicyDetailResponse`

NewEventPolicyDetailResponseWithDefaults instantiates a new EventPolicyDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsgYn

`func (o *EventPolicyDetailResponse) GetAsgYn() string`

GetAsgYn returns the AsgYn field if non-nil, zero value otherwise.

### GetAsgYnOk

`func (o *EventPolicyDetailResponse) GetAsgYnOk() (*string, bool)`

GetAsgYnOk returns a tuple with the AsgYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsgYn

`func (o *EventPolicyDetailResponse) SetAsgYn(v string)`

SetAsgYn sets AsgYn field to given value.

### HasAsgYn

`func (o *EventPolicyDetailResponse) HasAsgYn() bool`

HasAsgYn returns a boolean if a field has been set.

### GetAttrListStr

`func (o *EventPolicyDetailResponse) GetAttrListStr() string`

GetAttrListStr returns the AttrListStr field if non-nil, zero value otherwise.

### GetAttrListStrOk

`func (o *EventPolicyDetailResponse) GetAttrListStrOk() (*string, bool)`

GetAttrListStrOk returns a tuple with the AttrListStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttrListStr

`func (o *EventPolicyDetailResponse) SetAttrListStr(v string)`

SetAttrListStr sets AttrListStr field to given value.

### HasAttrListStr

`func (o *EventPolicyDetailResponse) HasAttrListStr() bool`

HasAttrListStr returns a boolean if a field has been set.

### GetCheckAsg

`func (o *EventPolicyDetailResponse) GetCheckAsg() bool`

GetCheckAsg returns the CheckAsg field if non-nil, zero value otherwise.

### GetCheckAsgOk

`func (o *EventPolicyDetailResponse) GetCheckAsgOk() (*bool, bool)`

GetCheckAsgOk returns a tuple with the CheckAsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAsg

`func (o *EventPolicyDetailResponse) SetCheckAsg(v bool)`

SetCheckAsg sets CheckAsg field to given value.

### HasCheckAsg

`func (o *EventPolicyDetailResponse) HasCheckAsg() bool`

HasCheckAsg returns a boolean if a field has been set.

### GetCreateBy

`func (o *EventPolicyDetailResponse) GetCreateBy() map[string]interface{}`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *EventPolicyDetailResponse) GetCreateByOk() (*map[string]interface{}, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *EventPolicyDetailResponse) SetCreateBy(v map[string]interface{})`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *EventPolicyDetailResponse) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetCreateById

`func (o *EventPolicyDetailResponse) GetCreateById() string`

GetCreateById returns the CreateById field if non-nil, zero value otherwise.

### GetCreateByIdOk

`func (o *EventPolicyDetailResponse) GetCreateByIdOk() (*string, bool)`

GetCreateByIdOk returns a tuple with the CreateById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateById

`func (o *EventPolicyDetailResponse) SetCreateById(v string)`

SetCreateById sets CreateById field to given value.

### HasCreateById

`func (o *EventPolicyDetailResponse) HasCreateById() bool`

HasCreateById returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EventPolicyDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EventPolicyDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EventPolicyDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EventPolicyDetailResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByName

`func (o *EventPolicyDetailResponse) GetCreatedByName() string`

GetCreatedByName returns the CreatedByName field if non-nil, zero value otherwise.

### GetCreatedByNameOk

`func (o *EventPolicyDetailResponse) GetCreatedByNameOk() (*string, bool)`

GetCreatedByNameOk returns a tuple with the CreatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByName

`func (o *EventPolicyDetailResponse) SetCreatedByName(v string)`

SetCreatedByName sets CreatedByName field to given value.

### HasCreatedByName

`func (o *EventPolicyDetailResponse) HasCreatedByName() bool`

HasCreatedByName returns a boolean if a field has been set.

### GetCreatedDt

`func (o *EventPolicyDetailResponse) GetCreatedDt() time.Time`

GetCreatedDt returns the CreatedDt field if non-nil, zero value otherwise.

### GetCreatedDtOk

`func (o *EventPolicyDetailResponse) GetCreatedDtOk() (*time.Time, bool)`

GetCreatedDtOk returns a tuple with the CreatedDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDt

`func (o *EventPolicyDetailResponse) SetCreatedDt(v time.Time)`

SetCreatedDt sets CreatedDt field to given value.

### HasCreatedDt

`func (o *EventPolicyDetailResponse) HasCreatedDt() bool`

HasCreatedDt returns a boolean if a field has been set.

### GetDisableObject

`func (o *EventPolicyDetailResponse) GetDisableObject() string`

GetDisableObject returns the DisableObject field if non-nil, zero value otherwise.

### GetDisableObjectOk

`func (o *EventPolicyDetailResponse) GetDisableObjectOk() (*string, bool)`

GetDisableObjectOk returns a tuple with the DisableObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableObject

`func (o *EventPolicyDetailResponse) SetDisableObject(v string)`

SetDisableObject sets DisableObject field to given value.

### HasDisableObject

`func (o *EventPolicyDetailResponse) HasDisableObject() bool`

HasDisableObject returns a boolean if a field has been set.

### GetDisableYn

`func (o *EventPolicyDetailResponse) GetDisableYn() string`

GetDisableYn returns the DisableYn field if non-nil, zero value otherwise.

### GetDisableYnOk

`func (o *EventPolicyDetailResponse) GetDisableYnOk() (*string, bool)`

GetDisableYnOk returns a tuple with the DisableYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableYn

`func (o *EventPolicyDetailResponse) SetDisableYn(v string)`

SetDisableYn sets DisableYn field to given value.

### HasDisableYn

`func (o *EventPolicyDetailResponse) HasDisableYn() bool`

HasDisableYn returns a boolean if a field has been set.

### GetDisplayEventRule

`func (o *EventPolicyDetailResponse) GetDisplayEventRule() string`

GetDisplayEventRule returns the DisplayEventRule field if non-nil, zero value otherwise.

### GetDisplayEventRuleOk

`func (o *EventPolicyDetailResponse) GetDisplayEventRuleOk() (*string, bool)`

GetDisplayEventRuleOk returns a tuple with the DisplayEventRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayEventRule

`func (o *EventPolicyDetailResponse) SetDisplayEventRule(v string)`

SetDisplayEventRule sets DisplayEventRule field to given value.

### HasDisplayEventRule

`func (o *EventPolicyDetailResponse) HasDisplayEventRule() bool`

HasDisplayEventRule returns a boolean if a field has been set.

### GetEventLevel

`func (o *EventPolicyDetailResponse) GetEventLevel() string`

GetEventLevel returns the EventLevel field if non-nil, zero value otherwise.

### GetEventLevelOk

`func (o *EventPolicyDetailResponse) GetEventLevelOk() (*string, bool)`

GetEventLevelOk returns a tuple with the EventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLevel

`func (o *EventPolicyDetailResponse) SetEventLevel(v string)`

SetEventLevel sets EventLevel field to given value.


### GetEventMessagePrefix

`func (o *EventPolicyDetailResponse) GetEventMessagePrefix() string`

GetEventMessagePrefix returns the EventMessagePrefix field if non-nil, zero value otherwise.

### GetEventMessagePrefixOk

`func (o *EventPolicyDetailResponse) GetEventMessagePrefixOk() (*string, bool)`

GetEventMessagePrefixOk returns a tuple with the EventMessagePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessagePrefix

`func (o *EventPolicyDetailResponse) SetEventMessagePrefix(v string)`

SetEventMessagePrefix sets EventMessagePrefix field to given value.

### HasEventMessagePrefix

`func (o *EventPolicyDetailResponse) HasEventMessagePrefix() bool`

HasEventMessagePrefix returns a boolean if a field has been set.

### GetEventOccurTimeZone

`func (o *EventPolicyDetailResponse) GetEventOccurTimeZone() string`

GetEventOccurTimeZone returns the EventOccurTimeZone field if non-nil, zero value otherwise.

### GetEventOccurTimeZoneOk

`func (o *EventPolicyDetailResponse) GetEventOccurTimeZoneOk() (*string, bool)`

GetEventOccurTimeZoneOk returns a tuple with the EventOccurTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOccurTimeZone

`func (o *EventPolicyDetailResponse) SetEventOccurTimeZone(v string)`

SetEventOccurTimeZone sets EventOccurTimeZone field to given value.

### HasEventOccurTimeZone

`func (o *EventPolicyDetailResponse) HasEventOccurTimeZone() bool`

HasEventOccurTimeZone returns a boolean if a field has been set.

### GetEventPolicyId

`func (o *EventPolicyDetailResponse) GetEventPolicyId() int64`

GetEventPolicyId returns the EventPolicyId field if non-nil, zero value otherwise.

### GetEventPolicyIdOk

`func (o *EventPolicyDetailResponse) GetEventPolicyIdOk() (*int64, bool)`

GetEventPolicyIdOk returns a tuple with the EventPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyId

`func (o *EventPolicyDetailResponse) SetEventPolicyId(v int64)`

SetEventPolicyId sets EventPolicyId field to given value.


### GetEventPolicyStatistics

`func (o *EventPolicyDetailResponse) GetEventPolicyStatistics() EventPolicyStatistics`

GetEventPolicyStatistics returns the EventPolicyStatistics field if non-nil, zero value otherwise.

### GetEventPolicyStatisticsOk

`func (o *EventPolicyDetailResponse) GetEventPolicyStatisticsOk() (*EventPolicyStatistics, bool)`

GetEventPolicyStatisticsOk returns a tuple with the EventPolicyStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyStatistics

`func (o *EventPolicyDetailResponse) SetEventPolicyStatistics(v EventPolicyStatistics)`

SetEventPolicyStatistics sets EventPolicyStatistics field to given value.

### HasEventPolicyStatistics

`func (o *EventPolicyDetailResponse) HasEventPolicyStatistics() bool`

HasEventPolicyStatistics returns a boolean if a field has been set.

### GetEventThreshold

`func (o *EventPolicyDetailResponse) GetEventThreshold() EventThreshold`

GetEventThreshold returns the EventThreshold field if non-nil, zero value otherwise.

### GetEventThresholdOk

`func (o *EventPolicyDetailResponse) GetEventThresholdOk() (*EventThreshold, bool)`

GetEventThresholdOk returns a tuple with the EventThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThreshold

`func (o *EventPolicyDetailResponse) SetEventThreshold(v EventThreshold)`

SetEventThreshold sets EventThreshold field to given value.


### GetFtCount

`func (o *EventPolicyDetailResponse) GetFtCount() int64`

GetFtCount returns the FtCount field if non-nil, zero value otherwise.

### GetFtCountOk

`func (o *EventPolicyDetailResponse) GetFtCountOk() (*int64, bool)`

GetFtCountOk returns a tuple with the FtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtCount

`func (o *EventPolicyDetailResponse) SetFtCount(v int64)`

SetFtCount sets FtCount field to given value.


### GetGroupSummary

`func (o *EventPolicyDetailResponse) GetGroupSummary() GroupSummary`

GetGroupSummary returns the GroupSummary field if non-nil, zero value otherwise.

### GetGroupSummaryOk

`func (o *EventPolicyDetailResponse) GetGroupSummaryOk() (*GroupSummary, bool)`

GetGroupSummaryOk returns a tuple with the GroupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSummary

`func (o *EventPolicyDetailResponse) SetGroupSummary(v GroupSummary)`

SetGroupSummary sets GroupSummary field to given value.

### HasGroupSummary

`func (o *EventPolicyDetailResponse) HasGroupSummary() bool`

HasGroupSummary returns a boolean if a field has been set.

### GetIsLogMetric

`func (o *EventPolicyDetailResponse) GetIsLogMetric() string`

GetIsLogMetric returns the IsLogMetric field if non-nil, zero value otherwise.

### GetIsLogMetricOk

`func (o *EventPolicyDetailResponse) GetIsLogMetricOk() (*string, bool)`

GetIsLogMetricOk returns a tuple with the IsLogMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLogMetric

`func (o *EventPolicyDetailResponse) SetIsLogMetric(v string)`

SetIsLogMetric sets IsLogMetric field to given value.


### GetMetricDescription

`func (o *EventPolicyDetailResponse) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *EventPolicyDetailResponse) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *EventPolicyDetailResponse) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *EventPolicyDetailResponse) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricDescriptionEn

`func (o *EventPolicyDetailResponse) GetMetricDescriptionEn() string`

GetMetricDescriptionEn returns the MetricDescriptionEn field if non-nil, zero value otherwise.

### GetMetricDescriptionEnOk

`func (o *EventPolicyDetailResponse) GetMetricDescriptionEnOk() (*string, bool)`

GetMetricDescriptionEnOk returns a tuple with the MetricDescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescriptionEn

`func (o *EventPolicyDetailResponse) SetMetricDescriptionEn(v string)`

SetMetricDescriptionEn sets MetricDescriptionEn field to given value.

### HasMetricDescriptionEn

`func (o *EventPolicyDetailResponse) HasMetricDescriptionEn() bool`

HasMetricDescriptionEn returns a boolean if a field has been set.

### GetMetricKey

`func (o *EventPolicyDetailResponse) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *EventPolicyDetailResponse) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *EventPolicyDetailResponse) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *EventPolicyDetailResponse) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *EventPolicyDetailResponse) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *EventPolicyDetailResponse) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *EventPolicyDetailResponse) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMetricSummary

`func (o *EventPolicyDetailResponse) GetMetricSummary() MetricSummary`

GetMetricSummary returns the MetricSummary field if non-nil, zero value otherwise.

### GetMetricSummaryOk

`func (o *EventPolicyDetailResponse) GetMetricSummaryOk() (*MetricSummary, bool)`

GetMetricSummaryOk returns a tuple with the MetricSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSummary

`func (o *EventPolicyDetailResponse) SetMetricSummary(v MetricSummary)`

SetMetricSummary sets MetricSummary field to given value.


### GetMetricType

`func (o *EventPolicyDetailResponse) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *EventPolicyDetailResponse) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *EventPolicyDetailResponse) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *EventPolicyDetailResponse) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetMetricUnit

`func (o *EventPolicyDetailResponse) GetMetricUnit() string`

GetMetricUnit returns the MetricUnit field if non-nil, zero value otherwise.

### GetMetricUnitOk

`func (o *EventPolicyDetailResponse) GetMetricUnitOk() (*string, bool)`

GetMetricUnitOk returns a tuple with the MetricUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricUnit

`func (o *EventPolicyDetailResponse) SetMetricUnit(v string)`

SetMetricUnit sets MetricUnit field to given value.

### HasMetricUnit

`func (o *EventPolicyDetailResponse) HasMetricUnit() bool`

HasMetricUnit returns a boolean if a field has been set.

### GetModifiedBy

`func (o *EventPolicyDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *EventPolicyDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *EventPolicyDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *EventPolicyDetailResponse) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedByName

`func (o *EventPolicyDetailResponse) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *EventPolicyDetailResponse) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *EventPolicyDetailResponse) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *EventPolicyDetailResponse) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedDt

`func (o *EventPolicyDetailResponse) GetModifiedDt() time.Time`

GetModifiedDt returns the ModifiedDt field if non-nil, zero value otherwise.

### GetModifiedDtOk

`func (o *EventPolicyDetailResponse) GetModifiedDtOk() (*time.Time, bool)`

GetModifiedDtOk returns a tuple with the ModifiedDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDt

`func (o *EventPolicyDetailResponse) SetModifiedDt(v time.Time)`

SetModifiedDt sets ModifiedDt field to given value.

### HasModifiedDt

`func (o *EventPolicyDetailResponse) HasModifiedDt() bool`

HasModifiedDt returns a boolean if a field has been set.

### GetObjectDisplayName

`func (o *EventPolicyDetailResponse) GetObjectDisplayName() string`

GetObjectDisplayName returns the ObjectDisplayName field if non-nil, zero value otherwise.

### GetObjectDisplayNameOk

`func (o *EventPolicyDetailResponse) GetObjectDisplayNameOk() (*string, bool)`

GetObjectDisplayNameOk returns a tuple with the ObjectDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDisplayName

`func (o *EventPolicyDetailResponse) SetObjectDisplayName(v string)`

SetObjectDisplayName sets ObjectDisplayName field to given value.

### HasObjectDisplayName

`func (o *EventPolicyDetailResponse) HasObjectDisplayName() bool`

HasObjectDisplayName returns a boolean if a field has been set.

### GetObjectName

`func (o *EventPolicyDetailResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *EventPolicyDetailResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *EventPolicyDetailResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *EventPolicyDetailResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *EventPolicyDetailResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EventPolicyDetailResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EventPolicyDetailResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *EventPolicyDetailResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *EventPolicyDetailResponse) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *EventPolicyDetailResponse) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *EventPolicyDetailResponse) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *EventPolicyDetailResponse) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetProductInfoAttrs

`func (o *EventPolicyDetailResponse) GetProductInfoAttrs() []ProductInfoAttr`

GetProductInfoAttrs returns the ProductInfoAttrs field if non-nil, zero value otherwise.

### GetProductInfoAttrsOk

`func (o *EventPolicyDetailResponse) GetProductInfoAttrsOk() (*[]ProductInfoAttr, bool)`

GetProductInfoAttrsOk returns a tuple with the ProductInfoAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfoAttrs

`func (o *EventPolicyDetailResponse) SetProductInfoAttrs(v []ProductInfoAttr)`

SetProductInfoAttrs sets ProductInfoAttrs field to given value.

### HasProductInfoAttrs

`func (o *EventPolicyDetailResponse) HasProductInfoAttrs() bool`

HasProductInfoAttrs returns a boolean if a field has been set.

### GetProductName

`func (o *EventPolicyDetailResponse) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *EventPolicyDetailResponse) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *EventPolicyDetailResponse) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *EventPolicyDetailResponse) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductResourceId

`func (o *EventPolicyDetailResponse) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *EventPolicyDetailResponse) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *EventPolicyDetailResponse) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.

### HasProductResourceId

`func (o *EventPolicyDetailResponse) HasProductResourceId() bool`

HasProductResourceId returns a boolean if a field has been set.

### GetProductSq

`func (o *EventPolicyDetailResponse) GetProductSq() int64`

GetProductSq returns the ProductSq field if non-nil, zero value otherwise.

### GetProductSqOk

`func (o *EventPolicyDetailResponse) GetProductSqOk() (*int64, bool)`

GetProductSqOk returns a tuple with the ProductSq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSq

`func (o *EventPolicyDetailResponse) SetProductSq(v int64)`

SetProductSq sets ProductSq field to given value.

### HasProductSq

`func (o *EventPolicyDetailResponse) HasProductSq() bool`

HasProductSq returns a boolean if a field has been set.

### GetProductSummary

`func (o *EventPolicyDetailResponse) GetProductSummary() ProductSummary`

GetProductSummary returns the ProductSummary field if non-nil, zero value otherwise.

### GetProductSummaryOk

`func (o *EventPolicyDetailResponse) GetProductSummaryOk() (*ProductSummary, bool)`

GetProductSummaryOk returns a tuple with the ProductSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSummary

`func (o *EventPolicyDetailResponse) SetProductSummary(v ProductSummary)`

SetProductSummary sets ProductSummary field to given value.


### GetProductTargetType

`func (o *EventPolicyDetailResponse) GetProductTargetType() string`

GetProductTargetType returns the ProductTargetType field if non-nil, zero value otherwise.

### GetProductTargetTypeOk

`func (o *EventPolicyDetailResponse) GetProductTargetTypeOk() (*string, bool)`

GetProductTargetTypeOk returns a tuple with the ProductTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetType

`func (o *EventPolicyDetailResponse) SetProductTargetType(v string)`

SetProductTargetType sets ProductTargetType field to given value.

### HasProductTargetType

`func (o *EventPolicyDetailResponse) HasProductTargetType() bool`

HasProductTargetType returns a boolean if a field has been set.

### GetProductTargetTypeEn

`func (o *EventPolicyDetailResponse) GetProductTargetTypeEn() string`

GetProductTargetTypeEn returns the ProductTargetTypeEn field if non-nil, zero value otherwise.

### GetProductTargetTypeEnOk

`func (o *EventPolicyDetailResponse) GetProductTargetTypeEnOk() (*string, bool)`

GetProductTargetTypeEnOk returns a tuple with the ProductTargetTypeEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetTypeEn

`func (o *EventPolicyDetailResponse) SetProductTargetTypeEn(v string)`

SetProductTargetTypeEn sets ProductTargetTypeEn field to given value.

### HasProductTargetTypeEn

`func (o *EventPolicyDetailResponse) HasProductTargetTypeEn() bool`

HasProductTargetTypeEn returns a boolean if a field has been set.

### GetStartDt

`func (o *EventPolicyDetailResponse) GetStartDt() time.Time`

GetStartDt returns the StartDt field if non-nil, zero value otherwise.

### GetStartDtOk

`func (o *EventPolicyDetailResponse) GetStartDtOk() (*time.Time, bool)`

GetStartDtOk returns a tuple with the StartDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDt

`func (o *EventPolicyDetailResponse) SetStartDt(v time.Time)`

SetStartDt sets StartDt field to given value.

### HasStartDt

`func (o *EventPolicyDetailResponse) HasStartDt() bool`

HasStartDt returns a boolean if a field has been set.

### GetUpdateBy

`func (o *EventPolicyDetailResponse) GetUpdateBy() map[string]interface{}`

GetUpdateBy returns the UpdateBy field if non-nil, zero value otherwise.

### GetUpdateByOk

`func (o *EventPolicyDetailResponse) GetUpdateByOk() (*map[string]interface{}, bool)`

GetUpdateByOk returns a tuple with the UpdateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateBy

`func (o *EventPolicyDetailResponse) SetUpdateBy(v map[string]interface{})`

SetUpdateBy sets UpdateBy field to given value.

### HasUpdateBy

`func (o *EventPolicyDetailResponse) HasUpdateBy() bool`

HasUpdateBy returns a boolean if a field has been set.

### GetUpdateById

`func (o *EventPolicyDetailResponse) GetUpdateById() string`

GetUpdateById returns the UpdateById field if non-nil, zero value otherwise.

### GetUpdateByIdOk

`func (o *EventPolicyDetailResponse) GetUpdateByIdOk() (*string, bool)`

GetUpdateByIdOk returns a tuple with the UpdateById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateById

`func (o *EventPolicyDetailResponse) SetUpdateById(v string)`

SetUpdateById sets UpdateById field to given value.

### HasUpdateById

`func (o *EventPolicyDetailResponse) HasUpdateById() bool`

HasUpdateById returns a boolean if a field has been set.

### GetUserNameList

`func (o *EventPolicyDetailResponse) GetUserNameList() []string`

GetUserNameList returns the UserNameList field if non-nil, zero value otherwise.

### GetUserNameListOk

`func (o *EventPolicyDetailResponse) GetUserNameListOk() (*[]string, bool)`

GetUserNameListOk returns a tuple with the UserNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameList

`func (o *EventPolicyDetailResponse) SetUserNameList(v []string)`

SetUserNameList sets UserNameList field to given value.

### HasUserNameList

`func (o *EventPolicyDetailResponse) HasUserNameList() bool`

HasUserNameList returns a boolean if a field has been set.

### GetUserNameStr

`func (o *EventPolicyDetailResponse) GetUserNameStr() string`

GetUserNameStr returns the UserNameStr field if non-nil, zero value otherwise.

### GetUserNameStrOk

`func (o *EventPolicyDetailResponse) GetUserNameStrOk() (*string, bool)`

GetUserNameStrOk returns a tuple with the UserNameStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameStr

`func (o *EventPolicyDetailResponse) SetUserNameStr(v string)`

SetUserNameStr sets UserNameStr field to given value.

### HasUserNameStr

`func (o *EventPolicyDetailResponse) HasUserNameStr() bool`

HasUserNameStr returns a boolean if a field has been set.

### GetUserNames

`func (o *EventPolicyDetailResponse) GetUserNames() string`

GetUserNames returns the UserNames field if non-nil, zero value otherwise.

### GetUserNamesOk

`func (o *EventPolicyDetailResponse) GetUserNamesOk() (*string, bool)`

GetUserNamesOk returns a tuple with the UserNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNames

`func (o *EventPolicyDetailResponse) SetUserNames(v string)`

SetUserNames sets UserNames field to given value.

### HasUserNames

`func (o *EventPolicyDetailResponse) HasUserNames() bool`

HasUserNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


