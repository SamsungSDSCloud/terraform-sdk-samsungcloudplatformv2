# ArchiveConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveEnabled** | Pointer to **bool** | Archive enabled | [optional] [default to true]
**RetentionPeriodDay** | Pointer to **int32** | Backup retention period (day) | [optional] [default to 3]

## Methods

### NewArchiveConfigResponse

`func NewArchiveConfigResponse() *ArchiveConfigResponse`

NewArchiveConfigResponse instantiates a new ArchiveConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveConfigResponseWithDefaults

`func NewArchiveConfigResponseWithDefaults() *ArchiveConfigResponse`

NewArchiveConfigResponseWithDefaults instantiates a new ArchiveConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveEnabled

`func (o *ArchiveConfigResponse) GetArchiveEnabled() bool`

GetArchiveEnabled returns the ArchiveEnabled field if non-nil, zero value otherwise.

### GetArchiveEnabledOk

`func (o *ArchiveConfigResponse) GetArchiveEnabledOk() (*bool, bool)`

GetArchiveEnabledOk returns a tuple with the ArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveEnabled

`func (o *ArchiveConfigResponse) SetArchiveEnabled(v bool)`

SetArchiveEnabled sets ArchiveEnabled field to given value.

### HasArchiveEnabled

`func (o *ArchiveConfigResponse) HasArchiveEnabled() bool`

HasArchiveEnabled returns a boolean if a field has been set.

### GetRetentionPeriodDay

`func (o *ArchiveConfigResponse) GetRetentionPeriodDay() int32`

GetRetentionPeriodDay returns the RetentionPeriodDay field if non-nil, zero value otherwise.

### GetRetentionPeriodDayOk

`func (o *ArchiveConfigResponse) GetRetentionPeriodDayOk() (*int32, bool)`

GetRetentionPeriodDayOk returns a tuple with the RetentionPeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriodDay

`func (o *ArchiveConfigResponse) SetRetentionPeriodDay(v int32)`

SetRetentionPeriodDay sets RetentionPeriodDay field to given value.

### HasRetentionPeriodDay

`func (o *ArchiveConfigResponse) HasRetentionPeriodDay() bool`

HasRetentionPeriodDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


