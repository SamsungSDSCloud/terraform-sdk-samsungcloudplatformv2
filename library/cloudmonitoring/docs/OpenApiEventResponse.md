# OpenApiEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationSecond** | **int64** | 이벤트 지속 시간 - 단위: 초 | 
**EndDt** | Pointer to **time.Time** | 종료일시 | [optional] 
**EventId** | **string** | 이벤트 아이디 | 
**EventLevel** | **string** | 이벤트 레벨 | 
**EventMessage** | **string** | 이벤트 메시지 | 
**EventPolicyId** | **int64** | 이벤트 정책 아이디 | 
**EventState** | **string** | Event State | 
**MetricKey** | **string** | 메트릭 키 | 
**MetricName** | **string** | 메트릭 이름 | 
**ObjectDisplayName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectType** | Pointer to **string** | 개별항목 유형 | [optional] 
**ObjectTypeName** | Pointer to **string** | 개별항목 유형 이름 | [optional] 
**ProductResourceId** | **string** | 상품 리소스 아이디 | 
**ProductTargetType** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**ProductTargetTypeEn** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**ProductTypeName** | **string** | 상품 유형 이름 | 
**StartDt** | **time.Time** | 시작일시 | 

## Methods

### NewOpenApiEventResponse

`func NewOpenApiEventResponse(durationSecond int64, eventId string, eventLevel string, eventMessage string, eventPolicyId int64, eventState string, metricKey string, metricName string, productResourceId string, productTypeName string, startDt time.Time, ) *OpenApiEventResponse`

NewOpenApiEventResponse instantiates a new OpenApiEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiEventResponseWithDefaults

`func NewOpenApiEventResponseWithDefaults() *OpenApiEventResponse`

NewOpenApiEventResponseWithDefaults instantiates a new OpenApiEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationSecond

`func (o *OpenApiEventResponse) GetDurationSecond() int64`

GetDurationSecond returns the DurationSecond field if non-nil, zero value otherwise.

### GetDurationSecondOk

`func (o *OpenApiEventResponse) GetDurationSecondOk() (*int64, bool)`

GetDurationSecondOk returns a tuple with the DurationSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSecond

`func (o *OpenApiEventResponse) SetDurationSecond(v int64)`

SetDurationSecond sets DurationSecond field to given value.


### GetEndDt

`func (o *OpenApiEventResponse) GetEndDt() time.Time`

GetEndDt returns the EndDt field if non-nil, zero value otherwise.

### GetEndDtOk

`func (o *OpenApiEventResponse) GetEndDtOk() (*time.Time, bool)`

GetEndDtOk returns a tuple with the EndDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDt

`func (o *OpenApiEventResponse) SetEndDt(v time.Time)`

SetEndDt sets EndDt field to given value.

### HasEndDt

`func (o *OpenApiEventResponse) HasEndDt() bool`

HasEndDt returns a boolean if a field has been set.

### GetEventId

`func (o *OpenApiEventResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *OpenApiEventResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *OpenApiEventResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventLevel

`func (o *OpenApiEventResponse) GetEventLevel() string`

GetEventLevel returns the EventLevel field if non-nil, zero value otherwise.

### GetEventLevelOk

`func (o *OpenApiEventResponse) GetEventLevelOk() (*string, bool)`

GetEventLevelOk returns a tuple with the EventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLevel

`func (o *OpenApiEventResponse) SetEventLevel(v string)`

SetEventLevel sets EventLevel field to given value.


### GetEventMessage

`func (o *OpenApiEventResponse) GetEventMessage() string`

GetEventMessage returns the EventMessage field if non-nil, zero value otherwise.

### GetEventMessageOk

`func (o *OpenApiEventResponse) GetEventMessageOk() (*string, bool)`

GetEventMessageOk returns a tuple with the EventMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessage

`func (o *OpenApiEventResponse) SetEventMessage(v string)`

SetEventMessage sets EventMessage field to given value.


### GetEventPolicyId

`func (o *OpenApiEventResponse) GetEventPolicyId() int64`

GetEventPolicyId returns the EventPolicyId field if non-nil, zero value otherwise.

### GetEventPolicyIdOk

`func (o *OpenApiEventResponse) GetEventPolicyIdOk() (*int64, bool)`

GetEventPolicyIdOk returns a tuple with the EventPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyId

`func (o *OpenApiEventResponse) SetEventPolicyId(v int64)`

SetEventPolicyId sets EventPolicyId field to given value.


### GetEventState

`func (o *OpenApiEventResponse) GetEventState() string`

GetEventState returns the EventState field if non-nil, zero value otherwise.

### GetEventStateOk

`func (o *OpenApiEventResponse) GetEventStateOk() (*string, bool)`

GetEventStateOk returns a tuple with the EventState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventState

`func (o *OpenApiEventResponse) SetEventState(v string)`

SetEventState sets EventState field to given value.


### GetMetricKey

`func (o *OpenApiEventResponse) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *OpenApiEventResponse) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *OpenApiEventResponse) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *OpenApiEventResponse) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *OpenApiEventResponse) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *OpenApiEventResponse) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetObjectDisplayName

`func (o *OpenApiEventResponse) GetObjectDisplayName() string`

GetObjectDisplayName returns the ObjectDisplayName field if non-nil, zero value otherwise.

### GetObjectDisplayNameOk

`func (o *OpenApiEventResponse) GetObjectDisplayNameOk() (*string, bool)`

GetObjectDisplayNameOk returns a tuple with the ObjectDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDisplayName

`func (o *OpenApiEventResponse) SetObjectDisplayName(v string)`

SetObjectDisplayName sets ObjectDisplayName field to given value.

### HasObjectDisplayName

`func (o *OpenApiEventResponse) HasObjectDisplayName() bool`

HasObjectDisplayName returns a boolean if a field has been set.

### GetObjectName

`func (o *OpenApiEventResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *OpenApiEventResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *OpenApiEventResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *OpenApiEventResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *OpenApiEventResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OpenApiEventResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OpenApiEventResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *OpenApiEventResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *OpenApiEventResponse) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *OpenApiEventResponse) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *OpenApiEventResponse) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *OpenApiEventResponse) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetProductResourceId

`func (o *OpenApiEventResponse) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *OpenApiEventResponse) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *OpenApiEventResponse) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.


### GetProductTargetType

`func (o *OpenApiEventResponse) GetProductTargetType() string`

GetProductTargetType returns the ProductTargetType field if non-nil, zero value otherwise.

### GetProductTargetTypeOk

`func (o *OpenApiEventResponse) GetProductTargetTypeOk() (*string, bool)`

GetProductTargetTypeOk returns a tuple with the ProductTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetType

`func (o *OpenApiEventResponse) SetProductTargetType(v string)`

SetProductTargetType sets ProductTargetType field to given value.

### HasProductTargetType

`func (o *OpenApiEventResponse) HasProductTargetType() bool`

HasProductTargetType returns a boolean if a field has been set.

### GetProductTargetTypeEn

`func (o *OpenApiEventResponse) GetProductTargetTypeEn() string`

GetProductTargetTypeEn returns the ProductTargetTypeEn field if non-nil, zero value otherwise.

### GetProductTargetTypeEnOk

`func (o *OpenApiEventResponse) GetProductTargetTypeEnOk() (*string, bool)`

GetProductTargetTypeEnOk returns a tuple with the ProductTargetTypeEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetTypeEn

`func (o *OpenApiEventResponse) SetProductTargetTypeEn(v string)`

SetProductTargetTypeEn sets ProductTargetTypeEn field to given value.

### HasProductTargetTypeEn

`func (o *OpenApiEventResponse) HasProductTargetTypeEn() bool`

HasProductTargetTypeEn returns a boolean if a field has been set.

### GetProductTypeName

`func (o *OpenApiEventResponse) GetProductTypeName() string`

GetProductTypeName returns the ProductTypeName field if non-nil, zero value otherwise.

### GetProductTypeNameOk

`func (o *OpenApiEventResponse) GetProductTypeNameOk() (*string, bool)`

GetProductTypeNameOk returns a tuple with the ProductTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeName

`func (o *OpenApiEventResponse) SetProductTypeName(v string)`

SetProductTypeName sets ProductTypeName field to given value.


### GetStartDt

`func (o *OpenApiEventResponse) GetStartDt() time.Time`

GetStartDt returns the StartDt field if non-nil, zero value otherwise.

### GetStartDtOk

`func (o *OpenApiEventResponse) GetStartDtOk() (*time.Time, bool)`

GetStartDtOk returns a tuple with the StartDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDt

`func (o *OpenApiEventResponse) SetStartDt(v time.Time)`

SetStartDt sets StartDt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


