# EventPolicyInfoEditable

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
**ObjectDisplayName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectType** | Pointer to **string** | 개별항목 유형 | [optional] 
**PodObjectDisplayName** | Pointer to **string** | swagger.event.eventPolicyResponse.podObjectName.value | [optional] 
**PodObjectName** | Pointer to **string** | swagger.event.eventPolicyResponse.podObjectName.value | [optional] 

## Methods

### NewEventPolicyInfoEditable

`func NewEventPolicyInfoEditable(disableYn string, eventLevel string, eventThreshold EventThreshold, ftCount int64, ) *EventPolicyInfoEditable`

NewEventPolicyInfoEditable instantiates a new EventPolicyInfoEditable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPolicyInfoEditableWithDefaults

`func NewEventPolicyInfoEditableWithDefaults() *EventPolicyInfoEditable`

NewEventPolicyInfoEditableWithDefaults instantiates a new EventPolicyInfoEditable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableYn

`func (o *EventPolicyInfoEditable) GetDisableYn() string`

GetDisableYn returns the DisableYn field if non-nil, zero value otherwise.

### GetDisableYnOk

`func (o *EventPolicyInfoEditable) GetDisableYnOk() (*string, bool)`

GetDisableYnOk returns a tuple with the DisableYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableYn

`func (o *EventPolicyInfoEditable) SetDisableYn(v string)`

SetDisableYn sets DisableYn field to given value.


### GetEventLevel

`func (o *EventPolicyInfoEditable) GetEventLevel() string`

GetEventLevel returns the EventLevel field if non-nil, zero value otherwise.

### GetEventLevelOk

`func (o *EventPolicyInfoEditable) GetEventLevelOk() (*string, bool)`

GetEventLevelOk returns a tuple with the EventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLevel

`func (o *EventPolicyInfoEditable) SetEventLevel(v string)`

SetEventLevel sets EventLevel field to given value.


### GetEventMessagePrefix

`func (o *EventPolicyInfoEditable) GetEventMessagePrefix() string`

GetEventMessagePrefix returns the EventMessagePrefix field if non-nil, zero value otherwise.

### GetEventMessagePrefixOk

`func (o *EventPolicyInfoEditable) GetEventMessagePrefixOk() (*string, bool)`

GetEventMessagePrefixOk returns a tuple with the EventMessagePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessagePrefix

`func (o *EventPolicyInfoEditable) SetEventMessagePrefix(v string)`

SetEventMessagePrefix sets EventMessagePrefix field to given value.

### HasEventMessagePrefix

`func (o *EventPolicyInfoEditable) HasEventMessagePrefix() bool`

HasEventMessagePrefix returns a boolean if a field has been set.

### GetEventOccurTimeZone

`func (o *EventPolicyInfoEditable) GetEventOccurTimeZone() string`

GetEventOccurTimeZone returns the EventOccurTimeZone field if non-nil, zero value otherwise.

### GetEventOccurTimeZoneOk

`func (o *EventPolicyInfoEditable) GetEventOccurTimeZoneOk() (*string, bool)`

GetEventOccurTimeZoneOk returns a tuple with the EventOccurTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOccurTimeZone

`func (o *EventPolicyInfoEditable) SetEventOccurTimeZone(v string)`

SetEventOccurTimeZone sets EventOccurTimeZone field to given value.

### HasEventOccurTimeZone

`func (o *EventPolicyInfoEditable) HasEventOccurTimeZone() bool`

HasEventOccurTimeZone returns a boolean if a field has been set.

### GetEventPolicyStatistics

`func (o *EventPolicyInfoEditable) GetEventPolicyStatistics() EventPolicyStatistics`

GetEventPolicyStatistics returns the EventPolicyStatistics field if non-nil, zero value otherwise.

### GetEventPolicyStatisticsOk

`func (o *EventPolicyInfoEditable) GetEventPolicyStatisticsOk() (*EventPolicyStatistics, bool)`

GetEventPolicyStatisticsOk returns a tuple with the EventPolicyStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyStatistics

`func (o *EventPolicyInfoEditable) SetEventPolicyStatistics(v EventPolicyStatistics)`

SetEventPolicyStatistics sets EventPolicyStatistics field to given value.

### HasEventPolicyStatistics

`func (o *EventPolicyInfoEditable) HasEventPolicyStatistics() bool`

HasEventPolicyStatistics returns a boolean if a field has been set.

### GetEventThreshold

`func (o *EventPolicyInfoEditable) GetEventThreshold() EventThreshold`

GetEventThreshold returns the EventThreshold field if non-nil, zero value otherwise.

### GetEventThresholdOk

`func (o *EventPolicyInfoEditable) GetEventThresholdOk() (*EventThreshold, bool)`

GetEventThresholdOk returns a tuple with the EventThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventThreshold

`func (o *EventPolicyInfoEditable) SetEventThreshold(v EventThreshold)`

SetEventThreshold sets EventThreshold field to given value.


### GetFtCount

`func (o *EventPolicyInfoEditable) GetFtCount() int64`

GetFtCount returns the FtCount field if non-nil, zero value otherwise.

### GetFtCountOk

`func (o *EventPolicyInfoEditable) GetFtCountOk() (*int64, bool)`

GetFtCountOk returns a tuple with the FtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtCount

`func (o *EventPolicyInfoEditable) SetFtCount(v int64)`

SetFtCount sets FtCount field to given value.


### GetObjectDisplayName

`func (o *EventPolicyInfoEditable) GetObjectDisplayName() string`

GetObjectDisplayName returns the ObjectDisplayName field if non-nil, zero value otherwise.

### GetObjectDisplayNameOk

`func (o *EventPolicyInfoEditable) GetObjectDisplayNameOk() (*string, bool)`

GetObjectDisplayNameOk returns a tuple with the ObjectDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDisplayName

`func (o *EventPolicyInfoEditable) SetObjectDisplayName(v string)`

SetObjectDisplayName sets ObjectDisplayName field to given value.

### HasObjectDisplayName

`func (o *EventPolicyInfoEditable) HasObjectDisplayName() bool`

HasObjectDisplayName returns a boolean if a field has been set.

### GetObjectName

`func (o *EventPolicyInfoEditable) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *EventPolicyInfoEditable) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *EventPolicyInfoEditable) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *EventPolicyInfoEditable) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *EventPolicyInfoEditable) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EventPolicyInfoEditable) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EventPolicyInfoEditable) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *EventPolicyInfoEditable) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPodObjectDisplayName

`func (o *EventPolicyInfoEditable) GetPodObjectDisplayName() string`

GetPodObjectDisplayName returns the PodObjectDisplayName field if non-nil, zero value otherwise.

### GetPodObjectDisplayNameOk

`func (o *EventPolicyInfoEditable) GetPodObjectDisplayNameOk() (*string, bool)`

GetPodObjectDisplayNameOk returns a tuple with the PodObjectDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodObjectDisplayName

`func (o *EventPolicyInfoEditable) SetPodObjectDisplayName(v string)`

SetPodObjectDisplayName sets PodObjectDisplayName field to given value.

### HasPodObjectDisplayName

`func (o *EventPolicyInfoEditable) HasPodObjectDisplayName() bool`

HasPodObjectDisplayName returns a boolean if a field has been set.

### GetPodObjectName

`func (o *EventPolicyInfoEditable) GetPodObjectName() string`

GetPodObjectName returns the PodObjectName field if non-nil, zero value otherwise.

### GetPodObjectNameOk

`func (o *EventPolicyInfoEditable) GetPodObjectNameOk() (*string, bool)`

GetPodObjectNameOk returns a tuple with the PodObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodObjectName

`func (o *EventPolicyInfoEditable) SetPodObjectName(v string)`

SetPodObjectName sets PodObjectName field to given value.

### HasPodObjectName

`func (o *EventPolicyInfoEditable) HasPodObjectName() bool`

HasPodObjectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


