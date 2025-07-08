# MetricInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableObject** | Pointer to **string** | swagger.product.MetricInfoDto.disableObject.value | [optional] 
**DisplayUnit** | Pointer to **string** |  | [optional] 
**FixedUnit** | Pointer to **string** | 메트릭 단위 (고정) | [optional] 
**IsLogMetric** | **string** | 로그 메트릭 여부 | 
**IsObjectExist** | **string** | 개별항목 성능 존재 여부 | 
**MetricDescription** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricDescriptionEn** | Pointer to **string** | 메트릭 설명 | [optional] 
**MetricKey** | **string** | 메트릭 키 | 
**MetricName** | **string** | 메트릭 이름 | 
**MetricOrder** | Pointer to **int32** | 메트릭 순서 | [optional] 
**MetricSetKey** | **string** | swagger.product.MetricInfoDto.metricSetKey.value | 
**MetricSetName** | **string** | 메트릭 세트 이름 | 
**MetricType** | **string** | 메트릭 유형 | 
**MetricUnit** | Pointer to **string** | 메트릭 단위 | [optional] 
**ObjectKeyName** | **string** | 오브젝트 키 이름 | 
**ObjectType** | Pointer to **string** | swagger.product.MetricInfoDto.disableObject.value | [optional] 
**ObjectTypeNameEng** | Pointer to **string** | swagger.product.MetricInfoDto.disableObject.value | [optional] 
**ObjectTypeNameLoc** | Pointer to **string** | swagger.product.MetricInfoDto.disableObject.value | [optional] 
**PerfTitle** | **string** | swagger.product.MetricInfoDto.perfTitle.value | 
**ProductTargetType** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**ProductTargetTypeEn** | Pointer to **string** | 메트릭 대상 유형 | [optional] 
**ProductTypeCode** | **string** | 상품 유형 코드 | 
**ProductTypeName** | **string** | 상품 유형 이름 | 

## Methods

### NewMetricInfoDto

`func NewMetricInfoDto(isLogMetric string, isObjectExist string, metricKey string, metricName string, metricSetKey string, metricSetName string, metricType string, objectKeyName string, perfTitle string, productTypeCode string, productTypeName string, ) *MetricInfoDto`

NewMetricInfoDto instantiates a new MetricInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricInfoDtoWithDefaults

`func NewMetricInfoDtoWithDefaults() *MetricInfoDto`

NewMetricInfoDtoWithDefaults instantiates a new MetricInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableObject

`func (o *MetricInfoDto) GetDisableObject() string`

GetDisableObject returns the DisableObject field if non-nil, zero value otherwise.

### GetDisableObjectOk

`func (o *MetricInfoDto) GetDisableObjectOk() (*string, bool)`

GetDisableObjectOk returns a tuple with the DisableObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableObject

`func (o *MetricInfoDto) SetDisableObject(v string)`

SetDisableObject sets DisableObject field to given value.

### HasDisableObject

`func (o *MetricInfoDto) HasDisableObject() bool`

HasDisableObject returns a boolean if a field has been set.

### GetDisplayUnit

`func (o *MetricInfoDto) GetDisplayUnit() string`

GetDisplayUnit returns the DisplayUnit field if non-nil, zero value otherwise.

### GetDisplayUnitOk

`func (o *MetricInfoDto) GetDisplayUnitOk() (*string, bool)`

GetDisplayUnitOk returns a tuple with the DisplayUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUnit

`func (o *MetricInfoDto) SetDisplayUnit(v string)`

SetDisplayUnit sets DisplayUnit field to given value.

### HasDisplayUnit

`func (o *MetricInfoDto) HasDisplayUnit() bool`

HasDisplayUnit returns a boolean if a field has been set.

### GetFixedUnit

`func (o *MetricInfoDto) GetFixedUnit() string`

GetFixedUnit returns the FixedUnit field if non-nil, zero value otherwise.

### GetFixedUnitOk

`func (o *MetricInfoDto) GetFixedUnitOk() (*string, bool)`

GetFixedUnitOk returns a tuple with the FixedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUnit

`func (o *MetricInfoDto) SetFixedUnit(v string)`

SetFixedUnit sets FixedUnit field to given value.

### HasFixedUnit

`func (o *MetricInfoDto) HasFixedUnit() bool`

HasFixedUnit returns a boolean if a field has been set.

### GetIsLogMetric

`func (o *MetricInfoDto) GetIsLogMetric() string`

GetIsLogMetric returns the IsLogMetric field if non-nil, zero value otherwise.

### GetIsLogMetricOk

`func (o *MetricInfoDto) GetIsLogMetricOk() (*string, bool)`

GetIsLogMetricOk returns a tuple with the IsLogMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLogMetric

`func (o *MetricInfoDto) SetIsLogMetric(v string)`

SetIsLogMetric sets IsLogMetric field to given value.


### GetIsObjectExist

`func (o *MetricInfoDto) GetIsObjectExist() string`

GetIsObjectExist returns the IsObjectExist field if non-nil, zero value otherwise.

### GetIsObjectExistOk

`func (o *MetricInfoDto) GetIsObjectExistOk() (*string, bool)`

GetIsObjectExistOk returns a tuple with the IsObjectExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObjectExist

`func (o *MetricInfoDto) SetIsObjectExist(v string)`

SetIsObjectExist sets IsObjectExist field to given value.


### GetMetricDescription

`func (o *MetricInfoDto) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *MetricInfoDto) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *MetricInfoDto) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *MetricInfoDto) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricDescriptionEn

`func (o *MetricInfoDto) GetMetricDescriptionEn() string`

GetMetricDescriptionEn returns the MetricDescriptionEn field if non-nil, zero value otherwise.

### GetMetricDescriptionEnOk

`func (o *MetricInfoDto) GetMetricDescriptionEnOk() (*string, bool)`

GetMetricDescriptionEnOk returns a tuple with the MetricDescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescriptionEn

`func (o *MetricInfoDto) SetMetricDescriptionEn(v string)`

SetMetricDescriptionEn sets MetricDescriptionEn field to given value.

### HasMetricDescriptionEn

`func (o *MetricInfoDto) HasMetricDescriptionEn() bool`

HasMetricDescriptionEn returns a boolean if a field has been set.

### GetMetricKey

`func (o *MetricInfoDto) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *MetricInfoDto) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *MetricInfoDto) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricName

`func (o *MetricInfoDto) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *MetricInfoDto) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *MetricInfoDto) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetMetricOrder

`func (o *MetricInfoDto) GetMetricOrder() int32`

GetMetricOrder returns the MetricOrder field if non-nil, zero value otherwise.

### GetMetricOrderOk

`func (o *MetricInfoDto) GetMetricOrderOk() (*int32, bool)`

GetMetricOrderOk returns a tuple with the MetricOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricOrder

`func (o *MetricInfoDto) SetMetricOrder(v int32)`

SetMetricOrder sets MetricOrder field to given value.

### HasMetricOrder

`func (o *MetricInfoDto) HasMetricOrder() bool`

HasMetricOrder returns a boolean if a field has been set.

### GetMetricSetKey

`func (o *MetricInfoDto) GetMetricSetKey() string`

GetMetricSetKey returns the MetricSetKey field if non-nil, zero value otherwise.

### GetMetricSetKeyOk

`func (o *MetricInfoDto) GetMetricSetKeyOk() (*string, bool)`

GetMetricSetKeyOk returns a tuple with the MetricSetKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSetKey

`func (o *MetricInfoDto) SetMetricSetKey(v string)`

SetMetricSetKey sets MetricSetKey field to given value.


### GetMetricSetName

`func (o *MetricInfoDto) GetMetricSetName() string`

GetMetricSetName returns the MetricSetName field if non-nil, zero value otherwise.

### GetMetricSetNameOk

`func (o *MetricInfoDto) GetMetricSetNameOk() (*string, bool)`

GetMetricSetNameOk returns a tuple with the MetricSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSetName

`func (o *MetricInfoDto) SetMetricSetName(v string)`

SetMetricSetName sets MetricSetName field to given value.


### GetMetricType

`func (o *MetricInfoDto) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *MetricInfoDto) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *MetricInfoDto) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetMetricUnit

`func (o *MetricInfoDto) GetMetricUnit() string`

GetMetricUnit returns the MetricUnit field if non-nil, zero value otherwise.

### GetMetricUnitOk

`func (o *MetricInfoDto) GetMetricUnitOk() (*string, bool)`

GetMetricUnitOk returns a tuple with the MetricUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricUnit

`func (o *MetricInfoDto) SetMetricUnit(v string)`

SetMetricUnit sets MetricUnit field to given value.

### HasMetricUnit

`func (o *MetricInfoDto) HasMetricUnit() bool`

HasMetricUnit returns a boolean if a field has been set.

### GetObjectKeyName

`func (o *MetricInfoDto) GetObjectKeyName() string`

GetObjectKeyName returns the ObjectKeyName field if non-nil, zero value otherwise.

### GetObjectKeyNameOk

`func (o *MetricInfoDto) GetObjectKeyNameOk() (*string, bool)`

GetObjectKeyNameOk returns a tuple with the ObjectKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectKeyName

`func (o *MetricInfoDto) SetObjectKeyName(v string)`

SetObjectKeyName sets ObjectKeyName field to given value.


### GetObjectType

`func (o *MetricInfoDto) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetricInfoDto) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetricInfoDto) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *MetricInfoDto) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeNameEng

`func (o *MetricInfoDto) GetObjectTypeNameEng() string`

GetObjectTypeNameEng returns the ObjectTypeNameEng field if non-nil, zero value otherwise.

### GetObjectTypeNameEngOk

`func (o *MetricInfoDto) GetObjectTypeNameEngOk() (*string, bool)`

GetObjectTypeNameEngOk returns a tuple with the ObjectTypeNameEng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeNameEng

`func (o *MetricInfoDto) SetObjectTypeNameEng(v string)`

SetObjectTypeNameEng sets ObjectTypeNameEng field to given value.

### HasObjectTypeNameEng

`func (o *MetricInfoDto) HasObjectTypeNameEng() bool`

HasObjectTypeNameEng returns a boolean if a field has been set.

### GetObjectTypeNameLoc

`func (o *MetricInfoDto) GetObjectTypeNameLoc() string`

GetObjectTypeNameLoc returns the ObjectTypeNameLoc field if non-nil, zero value otherwise.

### GetObjectTypeNameLocOk

`func (o *MetricInfoDto) GetObjectTypeNameLocOk() (*string, bool)`

GetObjectTypeNameLocOk returns a tuple with the ObjectTypeNameLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeNameLoc

`func (o *MetricInfoDto) SetObjectTypeNameLoc(v string)`

SetObjectTypeNameLoc sets ObjectTypeNameLoc field to given value.

### HasObjectTypeNameLoc

`func (o *MetricInfoDto) HasObjectTypeNameLoc() bool`

HasObjectTypeNameLoc returns a boolean if a field has been set.

### GetPerfTitle

`func (o *MetricInfoDto) GetPerfTitle() string`

GetPerfTitle returns the PerfTitle field if non-nil, zero value otherwise.

### GetPerfTitleOk

`func (o *MetricInfoDto) GetPerfTitleOk() (*string, bool)`

GetPerfTitleOk returns a tuple with the PerfTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfTitle

`func (o *MetricInfoDto) SetPerfTitle(v string)`

SetPerfTitle sets PerfTitle field to given value.


### GetProductTargetType

`func (o *MetricInfoDto) GetProductTargetType() string`

GetProductTargetType returns the ProductTargetType field if non-nil, zero value otherwise.

### GetProductTargetTypeOk

`func (o *MetricInfoDto) GetProductTargetTypeOk() (*string, bool)`

GetProductTargetTypeOk returns a tuple with the ProductTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetType

`func (o *MetricInfoDto) SetProductTargetType(v string)`

SetProductTargetType sets ProductTargetType field to given value.

### HasProductTargetType

`func (o *MetricInfoDto) HasProductTargetType() bool`

HasProductTargetType returns a boolean if a field has been set.

### GetProductTargetTypeEn

`func (o *MetricInfoDto) GetProductTargetTypeEn() string`

GetProductTargetTypeEn returns the ProductTargetTypeEn field if non-nil, zero value otherwise.

### GetProductTargetTypeEnOk

`func (o *MetricInfoDto) GetProductTargetTypeEnOk() (*string, bool)`

GetProductTargetTypeEnOk returns a tuple with the ProductTargetTypeEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTargetTypeEn

`func (o *MetricInfoDto) SetProductTargetTypeEn(v string)`

SetProductTargetTypeEn sets ProductTargetTypeEn field to given value.

### HasProductTargetTypeEn

`func (o *MetricInfoDto) HasProductTargetTypeEn() bool`

HasProductTargetTypeEn returns a boolean if a field has been set.

### GetProductTypeCode

`func (o *MetricInfoDto) GetProductTypeCode() string`

GetProductTypeCode returns the ProductTypeCode field if non-nil, zero value otherwise.

### GetProductTypeCodeOk

`func (o *MetricInfoDto) GetProductTypeCodeOk() (*string, bool)`

GetProductTypeCodeOk returns a tuple with the ProductTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeCode

`func (o *MetricInfoDto) SetProductTypeCode(v string)`

SetProductTypeCode sets ProductTypeCode field to given value.


### GetProductTypeName

`func (o *MetricInfoDto) GetProductTypeName() string`

GetProductTypeName returns the ProductTypeName field if non-nil, zero value otherwise.

### GetProductTypeNameOk

`func (o *MetricInfoDto) GetProductTypeNameOk() (*string, bool)`

GetProductTypeNameOk returns a tuple with the ProductTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeName

`func (o *MetricInfoDto) SetProductTypeName(v string)`

SetProductTypeName sets ProductTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


