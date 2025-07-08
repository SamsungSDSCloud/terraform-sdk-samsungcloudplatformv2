# EventPolicyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableYn** | **string** | swagger.event.eventPolicyResponse.disableYn.value | 
**EventLevel** | **string** | 이벤트 레벨 | 
**EventMessagePrefix** | Pointer to **string** | 이벤트 메시지 접두사 | [optional] 
**EventOccurTimeZone** | Pointer to **string** | swagger.event.eventPolicyResponse.eventOccurTimeZone.value | [optional] 
**EventPolicyStatistics** | Pointer to [**EventPolicyStatistics**](EventPolicyStatistics.md) |  | [optional] 
**EventThreshold** | [**EventThreshold**](EventThreshold.md) |  | 
**FtCount** | **int64** | 결함허용 개수 - 설정한 조건에 맞는 값이 몇 번 반복해서 발생하면 이벤트를 발생시킬지 설정하는 값 | 
**IsLogMetric** | **string** | 로그 메트릭 여부 | 
**MetricKey** | **string** | 메트릭 키 | 
**MetricName** | Pointer to **string** |  | [optional] 
**ObjectDisplayName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectType** | Pointer to **string** | 개별항목 유형 | [optional] 
**PodObjectDisplayName** | Pointer to **string** | swagger.event.eventPolicyResponse.podObjectDisplayName.value | [optional] 
**PodObjectName** | Pointer to **string** | swagger.event.eventPolicyResponse.podObjectName.value | [optional] 

## Methods

### NewEventPolicyInfo

`func NewEventPolicyInfo(disableYn string, eventLevel string, eventThreshold EventThreshold, ftCount int64, isLogMetric string, metricKey string, ) *EventPolicyInfo`

NewEventPolicyInfo instantiates a new EventPolicyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPolicyInfoWithDefaults

`func NewEventPolicyInfoWithDefaults() *EventPolicyInfo`

NewEventPolicyInfoWithDefaults instantiates a new EventPolicyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableYn

`func (o *EventPolicyInfo) GetDisableYn() string`

GetDisableYn returns the DisableYn field if non-nil, zero value otherwise.

### GetDisableYnOk

`func (o *EventPolicyInfo) GetDisableYnOk() (*string, bool)`

GetDisableYnOk returns a tuple with the DisableYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableYn

`func (o *EventPolicyInfo) SetDisableYn(v string)`

SetDisableYn sets DisableYn field to given value.


### GetEventLevel

`func (o *EventPolicyInfo) GetEventLevel() string`

GetEventLevel returns the EventLevel field if non-nil, zero value otherwise.

### GetEventLevelOk

`func (o *EventPolicyInfo) GetEventLevelOk() (*string, bool)`

GetEventLevelOk returns a tuple with the EventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLevel

`func (o *EventPolicyInfo) SetEventLevel(v string)`

SetEventLevel sets EventLevel field to given value.


### GetEventMessagePrefix

`func (o *EventPolicyInfo) GetEventMessagePrefix() string`

GetEventMessagePrefix returns the EventMessagePrefix field if non-nil, zero value otherwise.

### GetEventMessagePrefixOk

`func (o *EventPolicyInfo) GetEventMessagePrefixOk() (*string, bool)`

GetEventMessagePrefixOk returns a tuple with the EventMessagePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessagePrefix

`func (o *EventPolicyInfo) SetEventMessagePrefix(v string)`

SetEventMessagePrefix sets EventMessagePrefix field to given value.

### HasEventMessagePrefix

`func (o *EventPolicyInfo) HasEventMessagePrefix() bool`

HasEventMessagePrefix returns a boolean if a field has been set.

### GetEventOccurTimeZone

`func (o *EventPolicyInfo) GetEventOccurTimeZone() string`

GetEventOccurTimeZone returns the EventOccurTimeZone field if non-nil, zero value otherwise.

### GetEventOccurTimeZoneOk

`func (o *EventPolicyInfo) GetEventOccurTimeZoneOk() (*string, bool)`

GetEventOccurTimeZoneOk returns a tuple with the EventOccurTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOccurTimeZone

`func (o *EventPolicyInfo) SetEventOccurTimeZone(v string)`

SetEventOccurTimeZone sets EventOccurTimeZone field to given value.

### HasEventOccurTimeZone

`func (o *EventPolicyInfo) HasEventOccurTimeZone() bool`

HasEventOccurTimeZone returns a boolean if a field has been set.

### GetEventPolicyStatistics

`func (o *EventPolicyInfo) GetEventPolicyStatistics() EventPolicyStatistics`

GetEventPolicyStatistics returns the EventPolicyStatistics field if non-nil, zero value otherwise.

### GetEventPolicyStatisticsOk

`func (o *EventPolicyInfo) GetEventPolicyStatisticsOk() (*EventPolicyStatistics, bool)`

GetEventPolicyStatisticsOk returns a tuple with the EventPolicyStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyStatistics

`func (o *EventPolicyInfo) SetEventPolicyStatistics(v EventPolicyStatistics)`

SetEventPolicyStatistics sets EventPolicyStatistics field to given value.

### HasEventPolicyStatistics

`func (o *EventPolicyInfo) HasEventPolicyStatistics() bool`

HasEventPolicyStatistics returns a boolean if a field has been set.

### GetEventThreshold

`func (o *EventPolicyInfo) GetEventThreshold() EventThreshold`

GetEventThreshold returns the EventThreshold field if non-nil, zero value otherwise.

### GetEventThresholdOk

`func (o *EventPolicyInfo) GetEventThresholdOk() (*EventThreshold, bool)`

GetEventThresholdOk returns a tuple with the EventThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThreshold

`func (o *EventPolicyInfo) SetEventThreshold(v EventThreshold)`

SetEventThreshold sets EventThreshold field to given value.


### GetFtCount

`func (o *EventPolicyInfo) GetFtCount() int64`

GetFtCount returns the FtCount field if non-nil, zero value otherwise.

### GetFtCountOk

`func (o *EventPolicyInfo) GetFtCountOk() (*int64, bool)`

GetFtCountOk returns a tuple with the FtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtCount

`func (o *EventPolicyInfo) SetFtCount(v int64)`

SetFtCount sets FtCount field to given value.


### GetIsLogMetric

`func (o *EventPolicyInfo) GetIsLogMetric() string`

GetIsLogMetric returns the IsLogMetric field if non-nil, zero value otherwise.

### GetIsLogMetricOk

`func (o *EventPolicyInfo) GetIsLogMetricOk() (*string, bool)`

GetIsLogMetricOk returns a tuple with the IsLogMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLogMetric

`func (o *EventPolicyInfo) SetIsLogMetric(v string)`

SetIsLogMetric sets IsLogMetric field to given value.


### GetMetricKey

`func (o *EventPolicyInfo) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *EventPolicyInfo) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *EventPolicyInfo) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *EventPolicyInfo) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *EventPolicyInfo) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *EventPolicyInfo) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *EventPolicyInfo) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetObjectDisplayName

`func (o *EventPolicyInfo) GetObjectDisplayName() string`

GetObjectDisplayName returns the ObjectDisplayName field if non-nil, zero value otherwise.

### GetObjectDisplayNameOk

`func (o *EventPolicyInfo) GetObjectDisplayNameOk() (*string, bool)`

GetObjectDisplayNameOk returns a tuple with the ObjectDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDisplayName

`func (o *EventPolicyInfo) SetObjectDisplayName(v string)`

SetObjectDisplayName sets ObjectDisplayName field to given value.

### HasObjectDisplayName

`func (o *EventPolicyInfo) HasObjectDisplayName() bool`

HasObjectDisplayName returns a boolean if a field has been set.

### GetObjectName

`func (o *EventPolicyInfo) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *EventPolicyInfo) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *EventPolicyInfo) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *EventPolicyInfo) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *EventPolicyInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EventPolicyInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EventPolicyInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *EventPolicyInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPodObjectDisplayName

`func (o *EventPolicyInfo) GetPodObjectDisplayName() string`

GetPodObjectDisplayName returns the PodObjectDisplayName field if non-nil, zero value otherwise.

### GetPodObjectDisplayNameOk

`func (o *EventPolicyInfo) GetPodObjectDisplayNameOk() (*string, bool)`

GetPodObjectDisplayNameOk returns a tuple with the PodObjectDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodObjectDisplayName

`func (o *EventPolicyInfo) SetPodObjectDisplayName(v string)`

SetPodObjectDisplayName sets PodObjectDisplayName field to given value.

### HasPodObjectDisplayName

`func (o *EventPolicyInfo) HasPodObjectDisplayName() bool`

HasPodObjectDisplayName returns a boolean if a field has been set.

### GetPodObjectName

`func (o *EventPolicyInfo) GetPodObjectName() string`

GetPodObjectName returns the PodObjectName field if non-nil, zero value otherwise.

### GetPodObjectNameOk

`func (o *EventPolicyInfo) GetPodObjectNameOk() (*string, bool)`

GetPodObjectNameOk returns a tuple with the PodObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodObjectName

`func (o *EventPolicyInfo) SetPodObjectName(v string)`

SetPodObjectName sets PodObjectName field to given value.

### HasPodObjectName

`func (o *EventPolicyInfo) HasPodObjectName() bool`

HasPodObjectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


