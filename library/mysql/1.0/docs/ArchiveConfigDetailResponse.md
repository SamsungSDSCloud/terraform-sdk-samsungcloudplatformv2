# ArchiveConfigDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveEnabled** | **bool** | Archive enabled | 
**RetentionPeriodDay** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewArchiveConfigDetailResponse

`func NewArchiveConfigDetailResponse(archiveEnabled bool, ) *ArchiveConfigDetailResponse`

NewArchiveConfigDetailResponse instantiates a new ArchiveConfigDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveConfigDetailResponseWithDefaults

`func NewArchiveConfigDetailResponseWithDefaults() *ArchiveConfigDetailResponse`

NewArchiveConfigDetailResponseWithDefaults instantiates a new ArchiveConfigDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveEnabled

`func (o *ArchiveConfigDetailResponse) GetArchiveEnabled() bool`

GetArchiveEnabled returns the ArchiveEnabled field if non-nil, zero value otherwise.

### GetArchiveEnabledOk

`func (o *ArchiveConfigDetailResponse) GetArchiveEnabledOk() (*bool, bool)`

GetArchiveEnabledOk returns a tuple with the ArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveEnabled

`func (o *ArchiveConfigDetailResponse) SetArchiveEnabled(v bool)`

SetArchiveEnabled sets ArchiveEnabled field to given value.


### GetRetentionPeriodDay

`func (o *ArchiveConfigDetailResponse) GetRetentionPeriodDay() int32`

GetRetentionPeriodDay returns the RetentionPeriodDay field if non-nil, zero value otherwise.

### GetRetentionPeriodDayOk

`func (o *ArchiveConfigDetailResponse) GetRetentionPeriodDayOk() (*int32, bool)`

GetRetentionPeriodDayOk returns a tuple with the RetentionPeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDay

`func (o *ArchiveConfigDetailResponse) SetRetentionPeriodDay(v int32)`

SetRetentionPeriodDay sets RetentionPeriodDay field to given value.

### HasRetentionPeriodDay

`func (o *ArchiveConfigDetailResponse) HasRetentionPeriodDay() bool`

HasRetentionPeriodDay returns a boolean if a field has been set.

### SetRetentionPeriodDayNil

`func (o *ArchiveConfigDetailResponse) SetRetentionPeriodDayNil(b bool)`

 SetRetentionPeriodDayNil sets the value for RetentionPeriodDay to be an explicit nil

### UnsetRetentionPeriodDay
`func (o *ArchiveConfigDetailResponse) UnsetRetentionPeriodDay()`

UnsetRetentionPeriodDay ensures that no value is present for RetentionPeriodDay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


