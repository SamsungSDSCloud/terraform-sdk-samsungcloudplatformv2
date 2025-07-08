# EventDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationSecond** | **int64** | 이벤트 지속 시간 - 단위: 초 | 
**EndDt** | Pointer to **time.Time** | 종료일시 | [optional] 
**EventId** | **string** | 이벤트 아이디 | 
**EventLevel** | **string** | 이벤트 레벨 | 
**EventMessage** | **string** | 이벤트 메시지 | 
**EventPolicyId** | Pointer to **int64** | 이벤트 정책 아이디 | [optional] 
**EventPolicySummary** | [**EventPolicyResponse**](EventPolicyResponse.md) |  | 
**EventState** | **string** | Event State | 
**MetricKey** | **string** | 메트릭 키 | 
**MetricName** | Pointer to **string** | 메트릭 이름 | [optional] 
**MetricSummary** | [**MetricSummary**](MetricSummary.md) |  | 
**ObjectDisplayName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectName** | Pointer to **string** | 오브젝트 이름 | [optional] 
**ObjectType** | Pointer to **string** | 개별항목 유형 | [optional] 
**ObjectTypeName** | Pointer to **string** | 개별항목 유형 이름 | [optional] 
**ProductIpAddress** | Pointer to **string** | 상품 IP주소 | [optional] 
**ProductName** | Pointer to **string** | 상품 이름 | [optional] 
**ProductResourceId** | **string** | 상품 리소스 아이디 | 
**ProductSummary** | [**ProductSummary**](ProductSummary.md) |  | 
**ProductTypeCode** | Pointer to **string** | 상품 유형 코드 | [optional] 
**StartDt** | **time.Time** | 시작일시 | 

## Methods

### NewEventDetailResponse

`func NewEventDetailResponse(durationSecond int64, eventId string, eventLevel string, eventMessage string, eventPolicySummary EventPolicyResponse, eventState string, metricKey string, metricSummary MetricSummary, productResourceId string, productSummary ProductSummary, startDt time.Time, ) *EventDetailResponse`

NewEventDetailResponse instantiates a new EventDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDetailResponseWithDefaults

`func NewEventDetailResponseWithDefaults() *EventDetailResponse`

NewEventDetailResponseWithDefaults instantiates a new EventDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationSecond

`func (o *EventDetailResponse) GetDurationSecond() int64`

GetDurationSecond returns the DurationSecond field if non-nil, zero value otherwise.

### GetDurationSecondOk

`func (o *EventDetailResponse) GetDurationSecondOk() (*int64, bool)`

GetDurationSecondOk returns a tuple with the DurationSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSecond

`func (o *EventDetailResponse) SetDurationSecond(v int64)`

SetDurationSecond sets DurationSecond field to given value.


### GetEndDt

`func (o *EventDetailResponse) GetEndDt() time.Time`

GetEndDt returns the EndDt field if non-nil, zero value otherwise.

### GetEndDtOk

`func (o *EventDetailResponse) GetEndDtOk() (*time.Time, bool)`

GetEndDtOk returns a tuple with the EndDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDt

`func (o *EventDetailResponse) SetEndDt(v time.Time)`

SetEndDt sets EndDt field to given value.

### HasEndDt

`func (o *EventDetailResponse) HasEndDt() bool`

HasEndDt returns a boolean if a field has been set.

### GetEventId

`func (o *EventDetailResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventDetailResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventDetailResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventLevel

`func (o *EventDetailResponse) GetEventLevel() string`

GetEventLevel returns the EventLevel field if non-nil, zero value otherwise.

### GetEventLevelOk

`func (o *EventDetailResponse) GetEventLevelOk() (*string, bool)`

GetEventLevelOk returns a tuple with the EventLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLevel

`func (o *EventDetailResponse) SetEventLevel(v string)`

SetEventLevel sets EventLevel field to given value.


### GetEventMessage

`func (o *EventDetailResponse) GetEventMessage() string`

GetEventMessage returns the EventMessage field if non-nil, zero value otherwise.

### GetEventMessageOk

`func (o *EventDetailResponse) GetEventMessageOk() (*string, bool)`

GetEventMessageOk returns a tuple with the EventMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMessage

`func (o *EventDetailResponse) SetEventMessage(v string)`

SetEventMessage sets EventMessage field to given value.


### GetEventPolicyId

`func (o *EventDetailResponse) GetEventPolicyId() int64`

GetEventPolicyId returns the EventPolicyId field if non-nil, zero value otherwise.

### GetEventPolicyIdOk

`func (o *EventDetailResponse) GetEventPolicyIdOk() (*int64, bool)`

GetEventPolicyIdOk returns a tuple with the EventPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicyId

`func (o *EventDetailResponse) SetEventPolicyId(v int64)`

SetEventPolicyId sets EventPolicyId field to given value.

### HasEventPolicyId

`func (o *EventDetailResponse) HasEventPolicyId() bool`

HasEventPolicyId returns a boolean if a field has been set.

### GetEventPolicySummary

`func (o *EventDetailResponse) GetEventPolicySummary() EventPolicyResponse`

GetEventPolicySummary returns the EventPolicySummary field if non-nil, zero value otherwise.

### GetEventPolicySummaryOk

`func (o *EventDetailResponse) GetEventPolicySummaryOk() (*EventPolicyResponse, bool)`

GetEventPolicySummaryOk returns a tuple with the EventPolicySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPolicySummary

`func (o *EventDetailResponse) SetEventPolicySummary(v EventPolicyResponse)`

SetEventPolicySummary sets EventPolicySummary field to given value.


### GetEventState

`func (o *EventDetailResponse) GetEventState() string`

GetEventState returns the EventState field if non-nil, zero value otherwise.

### GetEventStateOk

`func (o *EventDetailResponse) GetEventStateOk() (*string, bool)`

GetEventStateOk returns a tuple with the EventState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventState

`func (o *EventDetailResponse) SetEventState(v string)`

SetEventState sets EventState field to given value.


### GetMetricKey

`func (o *EventDetailResponse) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *EventDetailResponse) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *EventDetailResponse) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *EventDetailResponse) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *EventDetailResponse) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *EventDetailResponse) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *EventDetailResponse) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMetricSummary

`func (o *EventDetailResponse) GetMetricSummary() MetricSummary`

GetMetricSummary returns the MetricSummary field if non-nil, zero value otherwise.

### GetMetricSummaryOk

`func (o *EventDetailResponse) GetMetricSummaryOk() (*MetricSummary, bool)`

GetMetricSummaryOk returns a tuple with the MetricSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSummary

`func (o *EventDetailResponse) SetMetricSummary(v MetricSummary)`

SetMetricSummary sets MetricSummary field to given value.


### GetObjectDisplayName

`func (o *EventDetailResponse) GetObjectDisplayName() string`

GetObjectDisplayName returns the ObjectDisplayName field if non-nil, zero value otherwise.

### GetObjectDisplayNameOk

`func (o *EventDetailResponse) GetObjectDisplayNameOk() (*string, bool)`

GetObjectDisplayNameOk returns a tuple with the ObjectDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectDisplayName

`func (o *EventDetailResponse) SetObjectDisplayName(v string)`

SetObjectDisplayName sets ObjectDisplayName field to given value.

### HasObjectDisplayName

`func (o *EventDetailResponse) HasObjectDisplayName() bool`

HasObjectDisplayName returns a boolean if a field has been set.

### GetObjectName

`func (o *EventDetailResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *EventDetailResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *EventDetailResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *EventDetailResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *EventDetailResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EventDetailResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EventDetailResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *EventDetailResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *EventDetailResponse) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *EventDetailResponse) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *EventDetailResponse) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *EventDetailResponse) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetProductIpAddress

`func (o *EventDetailResponse) GetProductIpAddress() string`

GetProductIpAddress returns the ProductIpAddress field if non-nil, zero value otherwise.

### GetProductIpAddressOk

`func (o *EventDetailResponse) GetProductIpAddressOk() (*string, bool)`

GetProductIpAddressOk returns a tuple with the ProductIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIpAddress

`func (o *EventDetailResponse) SetProductIpAddress(v string)`

SetProductIpAddress sets ProductIpAddress field to given value.

### HasProductIpAddress

`func (o *EventDetailResponse) HasProductIpAddress() bool`

HasProductIpAddress returns a boolean if a field has been set.

### GetProductName

`func (o *EventDetailResponse) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *EventDetailResponse) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *EventDetailResponse) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *EventDetailResponse) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductResourceId

`func (o *EventDetailResponse) GetProductResourceId() string`

GetProductResourceId returns the ProductResourceId field if non-nil, zero value otherwise.

### GetProductResourceIdOk

`func (o *EventDetailResponse) GetProductResourceIdOk() (*string, bool)`

GetProductResourceIdOk returns a tuple with the ProductResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductResourceId

`func (o *EventDetailResponse) SetProductResourceId(v string)`

SetProductResourceId sets ProductResourceId field to given value.


### GetProductSummary

`func (o *EventDetailResponse) GetProductSummary() ProductSummary`

GetProductSummary returns the ProductSummary field if non-nil, zero value otherwise.

### GetProductSummaryOk

`func (o *EventDetailResponse) GetProductSummaryOk() (*ProductSummary, bool)`

GetProductSummaryOk returns a tuple with the ProductSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSummary

`func (o *EventDetailResponse) SetProductSummary(v ProductSummary)`

SetProductSummary sets ProductSummary field to given value.


### GetProductTypeCode

`func (o *EventDetailResponse) GetProductTypeCode() string`

GetProductTypeCode returns the ProductTypeCode field if non-nil, zero value otherwise.

### GetProductTypeCodeOk

`func (o *EventDetailResponse) GetProductTypeCodeOk() (*string, bool)`

GetProductTypeCodeOk returns a tuple with the ProductTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeCode

`func (o *EventDetailResponse) SetProductTypeCode(v string)`

SetProductTypeCode sets ProductTypeCode field to given value.

### HasProductTypeCode

`func (o *EventDetailResponse) HasProductTypeCode() bool`

HasProductTypeCode returns a boolean if a field has been set.

### GetStartDt

`func (o *EventDetailResponse) GetStartDt() time.Time`

GetStartDt returns the StartDt field if non-nil, zero value otherwise.

### GetStartDtOk

`func (o *EventDetailResponse) GetStartDtOk() (*time.Time, bool)`

GetStartDtOk returns a tuple with the StartDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDt

`func (o *EventDetailResponse) SetStartDt(v time.Time)`

SetStartDt sets StartDt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


