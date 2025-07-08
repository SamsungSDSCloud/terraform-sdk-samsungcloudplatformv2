# EventPolicyHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateBy** | Pointer to **map[string]interface{}** |  | [optional] 
**CreateById** | Pointer to **string** |  | [optional] 
**DisableObject** | Pointer to **string** | swagger.event.eventPolicyResponse.disableObject.value | [optional] 
**DisableYn** | Pointer to **string** | swagger.event.eventPolicyResponse.disableObject.value | [optional] 
**EventLevel** | **string** | 이벤트 레벨 | 
**EventMessagePrefix** | Pointer to **string** | 이벤트 메시지 접두사 | [optional] 
**EventOccurTimeZone** | Pointer to **string** | swagger.event.eventPolicyResponse.eventOccurTimeZone.value | [optional] 
**EventPolicyHistoryId** | **int64** | 이벤트 정책 이력 아이디 | 
**EventPolicyHistoryType** | **string** | 이벤트 정책 이력 유형 - C : 등록, U : 수정, D : 삭제 | 
**EventPolicyId** | Pointer to **int64** | 이벤트 정책 아이디 | [optional] 
**EventPolicyStatistics** | Pointer to [**EventPolicyStatistics**](EventPolicyStatistics.md) |  | [optional] 
**EventThreshold** | [**EventThreshold**](EventThreshold.md) |  | 
**FtCount** | **int64** | 결함허용 개수 | 
**MetricDescription** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricDescriptionEn** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricKey** | **string** | 메트릭 키 | 
**MetricName** | **string** | 메트릭 이름 | 
**MetricUnit** | Pointer to **string** | 메트릭 단위 | [optional] 
**ModifiedBy** | Pointer to **string** |  | [optional] 
**ModifiedByName** | Pointer to **string** |  | [optional] 
**ModifiedDt** | Pointer to **time.Time** |  | [optional] 
**NotificationRecipientHistory** | Pointer to [**NotificationRecipientHistory**](NotificationRecipientHistory.md) |  | [optional] 
**ObjectDisplayName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectType** | Pointer to **string** | 개별항목 유형 | [optional] 
**ProductName** | **string** | 상품 이름 | 
**ProductResourceId** | **string** | 상품 리소스 아이디 | 
**ProductTargetType** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**ProductTargetTypeEn** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**UpdateBy** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdateById** | Pointer to **string** |  | [optional] 

## Methods

### NewEventPolicyHistoryResponse

`func NewEventPolicyHistoryResponse(eventLevel string, eventPolicyHistoryId int64, eventPolicyHistoryType string, eventThreshold EventThreshold, ftCount int64, metricKey string, metricName string, productName string, productResourceId string, ) *EventPolicyHistoryResponse`

NewEventPolicyHistoryResponse instantiates a new EventPolicyHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPolicyHistoryResponseWithDefaults

`func NewEventPolicyHistoryResponseWithDefaults() *EventPolicyHistoryResponse`

NewEventPolicyHistoryResponseWithDefaults instantiates a new EventPolicyHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateBy

`func (o *EventPolicyHistoryResponse) GetCreateBy() map[string]interface{}`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *EventPolicyHistoryResponse) GetCreateByOk() (*map[string]interface{}, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *EventPolicyHistoryResponse) SetCreateBy(v map[string]interface{})`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *EventPolicyHistoryResponse) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetCreateById

`func (o *EventPolicyHistoryResponse) GetCreateById() string`

GetCreateById returns the CreateById field if non-nil, zero value otherwise.

### GetCreateByIdOk

`func (o *EventPolicyHistoryResponse) GetCreateByIdOk() (*string, bool)`

GetCreateByIdOk returns a tuple with the CreateById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateById

`func (o *EventPolicyHistoryResponse) SetCreateById(v string)`

SetCreateById sets CreateById field to given value.

### HasCreateById

`func (o *EventPolicyHistoryResponse) HasCreateById() bool`

HasCreateById returns a boolean if a field has been set.

### GetDisableObject

`func (o *EventPolicyHistoryResponse) GetDisableObject() string`

GetDisableObject returns the DisableObject field if non-nil, zero value otherwise.

### GetDisableObjectOk

`func (o *EventPolicyHistoryResponse) GetDisableObjectOk() (*string, bool)`

GetDisableObjectOk returns a tuple with the DisableObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableObject

`func (o *EventPolicyHistoryResponse) SetDisableObject(v string)`

SetDisableObject sets DisableObject field to given value.

### HasDisableObject

`func (o *EventPolicyHistoryResponse) HasDisableObject() bool`

HasDisableObject returns a boolean if a field has been set.

### GetDisableYn

`func (o *EventPolicyHistoryResponse) GetDisableYn() string`

GetDisableYn returns the DisableYn field if non-nil, zero value otherwise.

### GetDisableYnOk

`func (o *EventPolicyHistoryResponse) GetDisableYnOk() (*string, bool)`

GetDisableYnOk returns a tuple with the DisableYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableYn

`func (o *EventPolicyHistoryResponse) SetDisableYn(v string)`

SetDisableYn sets DisableYn field to given value.

### HasDisableYn

`func (o *EventPolicyHistoryResponse) HasDisableYn() bool`

HasDisableYn returns a boolean if a field has been set.

### GetEventLevel

`func (o *EventPolicyHistoryResponse) GetEventLevel() string`

GetEventLevel returns the EventLevel field if non-nil, zero value otherwise.

### GetEventLevelOk

`func (o *EventPolicyHistoryResponse) GetEventLevelOk() (*string, bool)`

GetEventLevelOk returns a tuple with the EventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLevel

`func (o *EventPolicyHistoryResponse) SetEventLevel(v string)`

SetEventLevel sets EventLevel field to given value.


### GetEventMessagePrefix

`func (o *EventPolicyHistoryResponse) GetEventMessagePrefix() string`

GetEventMessagePrefix returns the EventMessagePrefix field if non-nil, zero value otherwise.

### GetEventMessagePrefixOk

`func (o *EventPolicyHistoryResponse) GetEventMessagePrefixOk() (*string, bool)`

GetEventMessagePrefixOk returns a tuple with the EventMessagePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessagePrefix

`func (o *EventPolicyHistoryResponse) SetEventMessagePrefix(v string)`

SetEventMessagePrefix sets EventMessagePrefix field to given value.

### HasEventMessagePrefix

`func (o *EventPolicyHistoryResponse) HasEventMessagePrefix() bool`

HasEventMessagePrefix returns a boolean if a field has been set.

### GetEventOccurTimeZone

`func (o *EventPolicyHistoryResponse) GetEventOccurTimeZone() string`

GetEventOccurTimeZone returns the EventOccurTimeZone field if non-nil, zero value otherwise.

### GetEventOccurTimeZoneOk

`func (o *EventPolicyHistoryResponse) GetEventOccurTimeZoneOk() (*string, bool)`

GetEventOccurTimeZoneOk returns a tuple with the EventOccurTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOccurTimeZone

`func (o *EventPolicyHistoryResponse) SetEventOccurTimeZone(v string)`

SetEventOccurTimeZone sets EventOccurTimeZone field to given value.

### HasEventOccurTimeZone

`func (o *EventPolicyHistoryResponse) HasEventOccurTimeZone() bool`

HasEventOccurTimeZone returns a boolean if a field has been set.

### GetEventPolicyHistoryId

`func (o *EventPolicyHistoryResponse) GetEventPolicyHistoryId() int64`

GetEventPolicyHistoryId returns the EventPolicyHistoryId field if non-nil, zero value otherwise.

### GetEventPolicyHistoryIdOk

`func (o *EventPolicyHistoryResponse) GetEventPolicyHistoryIdOk() (*int64, bool)`

GetEventPolicyHistoryIdOk returns a tuple with the EventPolicyHistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyHistoryId

`func (o *EventPolicyHistoryResponse) SetEventPolicyHistoryId(v int64)`

SetEventPolicyHistoryId sets EventPolicyHistoryId field to given value.


### GetEventPolicyHistoryType

`func (o *EventPolicyHistoryResponse) GetEventPolicyHistoryType() string`

GetEventPolicyHistoryType returns the EventPolicyHistoryType field if non-nil, zero value otherwise.

### GetEventPolicyHistoryTypeOk

`func (o *EventPolicyHistoryResponse) GetEventPolicyHistoryTypeOk() (*string, bool)`

GetEventPolicyHistoryTypeOk returns a tuple with the EventPolicyHistoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyHistoryType

`func (o *EventPolicyHistoryResponse) SetEventPolicyHistoryType(v string)`

SetEventPolicyHistoryType sets EventPolicyHistoryType field to given value.


### GetEventPolicyId

`func (o *EventPolicyHistoryResponse) GetEventPolicyId() int64`

GetEventPolicyId returns the EventPolicyId field if non-nil, zero value otherwise.

### GetEventPolicyIdOk

`func (o *EventPolicyHistoryResponse) GetEventPolicyIdOk() (*int64, bool)`

GetEventPolicyIdOk returns a tuple with the EventPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyId

`func (o *EventPolicyHistoryResponse) SetEventPolicyId(v int64)`

SetEventPolicyId sets EventPolicyId field to given value.

### HasEventPolicyId

`func (o *EventPolicyHistoryResponse) HasEventPolicyId() bool`

HasEventPolicyId returns a boolean if a field has been set.

### GetEventPolicyStatistics

`func (o *EventPolicyHistoryResponse) GetEventPolicyStatistics() EventPolicyStatistics`

GetEventPolicyStatistics returns the EventPolicyStatistics field if non-nil, zero value otherwise.

### GetEventPolicyStatisticsOk

`func (o *EventPolicyHistoryResponse) GetEventPolicyStatisticsOk() (*EventPolicyStatistics, bool)`

GetEventPolicyStatisticsOk returns a tuple with the EventPolicyStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyStatistics

`func (o *EventPolicyHistoryResponse) SetEventPolicyStatistics(v EventPolicyStatistics)`

SetEventPolicyStatistics sets EventPolicyStatistics field to given value.

### HasEventPolicyStatistics

`func (o *EventPolicyHistoryResponse) HasEventPolicyStatistics() bool`

HasEventPolicyStatistics returns a boolean if a field has been set.

### GetEventThreshold

`func (o *EventPolicyHistoryResponse) GetEventThreshold() EventThreshold`

GetEventThreshold returns the EventThreshold field if non-nil, zero value otherwise.

### GetEventThresholdOk

`func (o *EventPolicyHistoryResponse) GetEventThresholdOk() (*EventThreshold, bool)`

GetEventThresholdOk returns a tuple with the EventThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThreshold

`func (o *EventPolicyHistoryResponse) SetEventThreshold(v EventThreshold)`

SetEventThreshold sets EventThreshold field to given value.


### GetFtCount

`func (o *EventPolicyHistoryResponse) GetFtCount() int64`

GetFtCount returns the FtCount field if non-nil, zero value otherwise.

### GetFtCountOk

`func (o *EventPolicyHistoryResponse) GetFtCountOk() (*int64, bool)`

GetFtCountOk returns a tuple with the FtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtCount

`func (o *EventPolicyHistoryResponse) SetFtCount(v int64)`

SetFtCount sets FtCount field to given value.


### GetMetricDescription

`func (o *EventPolicyHistoryResponse) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *EventPolicyHistoryResponse) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *EventPolicyHistoryResponse) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *EventPolicyHistoryResponse) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricDescriptionEn

`func (o *EventPolicyHistoryResponse) GetMetricDescriptionEn() string`

GetMetricDescriptionEn returns the MetricDescriptionEn field if non-nil, zero value otherwise.

### GetMetricDescriptionEnOk

`func (o *EventPolicyHistoryResponse) GetMetricDescriptionEnOk() (*string, bool)`

GetMetricDescriptionEnOk returns a tuple with the MetricDescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescriptionEn

`func (o *EventPolicyHistoryResponse) SetMetricDescriptionEn(v string)`

SetMetricDescriptionEn sets MetricDescriptionEn field to given value.

### HasMetricDescriptionEn

`func (o *EventPolicyHistoryResponse) HasMetricDescriptionEn() bool`

HasMetricDescriptionEn returns a boolean if a field has been set.

### GetMetricKey

`func (o *EventPolicyHistoryResponse) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *EventPolicyHistoryResponse) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *EventPolicyHistoryResponse) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *EventPolicyHistoryResponse) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *EventPolicyHistoryResponse) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *EventPolicyHistoryResponse) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMetricUnit

`func (o *EventPolicyHistoryResponse) GetMetricUnit() string`

GetMetricUnit returns the MetricUnit field if non-nil, zero value otherwise.

### GetMetricUnitOk

`func (o *EventPolicyHistoryResponse) GetMetricUnitOk() (*string, bool)`

GetMetricUnitOk returns a tuple with the MetricUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricUnit

`func (o *EventPolicyHistoryResponse) SetMetricUnit(v string)`

SetMetricUnit sets MetricUnit field to given value.

### HasMetricUnit

`func (o *EventPolicyHistoryResponse) HasMetricUnit() bool`

HasMetricUnit returns a boolean if a field has been set.

### GetModifiedBy

`func (o *EventPolicyHistoryResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *EventPolicyHistoryResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *EventPolicyHistoryResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *EventPolicyHistoryResponse) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedByName

`func (o *EventPolicyHistoryResponse) GetModifiedByName() string`

GetModifiedByName returns the ModifiedByName field if non-nil, zero value otherwise.

### GetModifiedByNameOk

`func (o *EventPolicyHistoryResponse) GetModifiedByNameOk() (*string, bool)`

GetModifiedByNameOk returns a tuple with the ModifiedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByName

`func (o *EventPolicyHistoryResponse) SetModifiedByName(v string)`

SetModifiedByName sets ModifiedByName field to given value.

### HasModifiedByName

`func (o *EventPolicyHistoryResponse) HasModifiedByName() bool`

HasModifiedByName returns a boolean if a field has been set.

### GetModifiedDt

`func (o *EventPolicyHistoryResponse) GetModifiedDt() time.Time`

GetModifiedDt returns the ModifiedDt field if non-nil, zero value otherwise.

### GetModifiedDtOk

`func (o *EventPolicyHistoryResponse) GetModifiedDtOk() (*time.Time, bool)`

GetModifiedDtOk returns a tuple with the ModifiedDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDt

`func (o *EventPolicyHistoryResponse) SetModifiedDt(v time.Time)`

SetModifiedDt sets ModifiedDt field to given value.

### HasModifiedDt

`func (o *EventPolicyHistoryResponse) HasModifiedDt() bool`

HasModifiedDt returns a boolean if a field has been set.

### GetNotificationRecipientHistory

`func (o *EventPolicyHistoryResponse) GetNotificationRecipientHistory() NotificationRecipientHistory`

GetNotificationRecipientHistory returns the NotificationRecipientHistory field if non-nil, zero value otherwise.

### GetNotificationRecipientHistoryOk

`func (o *EventPolicyHistoryResponse) GetNotificationRecipientHistoryOk() (*NotificationRecipientHistory, bool)`

GetNotificationRecipientHistoryOk returns a tuple with the NotificationRecipientHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationRecipientHistory

`func (o *EventPolicyHistoryResponse) SetNotificationRecipientHistory(v NotificationRecipientHistory)`

SetNotificationRecipientHistory sets NotificationRecipientHistory field to given value.

### HasNotificationRecipientHistory

`func (o *EventPolicyHistoryResponse) HasNotificationRecipientHistory() bool`

HasNotificationRecipientHistory returns a boolean if a field has been set.

### GetObjectDisplayName

`func (o *EventPolicyHistoryResponse) GetObjectDisplayName() string`

GetObjectDisplayName returns the ObjectDisplayName field if non-nil, zero value otherwise.

### GetObjectDisplayNameOk

`func (o *EventPolicyHistoryResponse) GetObjectDisplayNameOk() (*string, bool)`

GetObjectDisplayNameOk returns a tuple with the ObjectDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDisplayName

`func (o *EventPolicyHistoryResponse) SetObjectDisplayName(v string)`

SetObjectDisplayName sets ObjectDisplayName field to given value.

### HasObjectDisplayName

`func (o *EventPolicyHistoryResponse) HasObjectDisplayName() bool`

HasObjectDisplayName returns a boolean if a field has been set.

### GetObjectName

`func (o *EventPolicyHistoryResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *EventPolicyHistoryResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *EventPolicyHistoryResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *EventPolicyHistoryResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *EventPolicyHistoryResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EventPolicyHistoryResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EventPolicyHistoryResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *EventPolicyHistoryResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetProductName

`func (o *EventPolicyHistoryResponse) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *EventPolicyHistoryResponse) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *EventPolicyHistoryResponse) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetProductResourceId

`func (o *EventPolicyHistoryResponse) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *EventPolicyHistoryResponse) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *EventPolicyHistoryResponse) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.


### GetProductTargetType

`func (o *EventPolicyHistoryResponse) GetProductTargetType() string`

GetProductTargetType returns the ProductTargetType field if non-nil, zero value otherwise.

### GetProductTargetTypeOk

`func (o *EventPolicyHistoryResponse) GetProductTargetTypeOk() (*string, bool)`

GetProductTargetTypeOk returns a tuple with the ProductTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetType

`func (o *EventPolicyHistoryResponse) SetProductTargetType(v string)`

SetProductTargetType sets ProductTargetType field to given value.

### HasProductTargetType

`func (o *EventPolicyHistoryResponse) HasProductTargetType() bool`

HasProductTargetType returns a boolean if a field has been set.

### GetProductTargetTypeEn

`func (o *EventPolicyHistoryResponse) GetProductTargetTypeEn() string`

GetProductTargetTypeEn returns the ProductTargetTypeEn field if non-nil, zero value otherwise.

### GetProductTargetTypeEnOk

`func (o *EventPolicyHistoryResponse) GetProductTargetTypeEnOk() (*string, bool)`

GetProductTargetTypeEnOk returns a tuple with the ProductTargetTypeEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetTypeEn

`func (o *EventPolicyHistoryResponse) SetProductTargetTypeEn(v string)`

SetProductTargetTypeEn sets ProductTargetTypeEn field to given value.

### HasProductTargetTypeEn

`func (o *EventPolicyHistoryResponse) HasProductTargetTypeEn() bool`

HasProductTargetTypeEn returns a boolean if a field has been set.

### GetUpdateBy

`func (o *EventPolicyHistoryResponse) GetUpdateBy() map[string]interface{}`

GetUpdateBy returns the UpdateBy field if non-nil, zero value otherwise.

### GetUpdateByOk

`func (o *EventPolicyHistoryResponse) GetUpdateByOk() (*map[string]interface{}, bool)`

GetUpdateByOk returns a tuple with the UpdateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateBy

`func (o *EventPolicyHistoryResponse) SetUpdateBy(v map[string]interface{})`

SetUpdateBy sets UpdateBy field to given value.

### HasUpdateBy

`func (o *EventPolicyHistoryResponse) HasUpdateBy() bool`

HasUpdateBy returns a boolean if a field has been set.

### GetUpdateById

`func (o *EventPolicyHistoryResponse) GetUpdateById() string`

GetUpdateById returns the UpdateById field if non-nil, zero value otherwise.

### GetUpdateByIdOk

`func (o *EventPolicyHistoryResponse) GetUpdateByIdOk() (*string, bool)`

GetUpdateByIdOk returns a tuple with the UpdateById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateById

`func (o *EventPolicyHistoryResponse) SetUpdateById(v string)`

SetUpdateById sets UpdateById field to given value.

### HasUpdateById

`func (o *EventPolicyHistoryResponse) HasUpdateById() bool`

HasUpdateById returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


