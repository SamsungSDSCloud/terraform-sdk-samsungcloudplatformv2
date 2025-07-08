# SnapshotScheduleSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotRetentionCount** | Pointer to **int32** | Retention count | [optional] [default to 0]
**SnapshotSchedule** | [**SnapshotSchedule**](SnapshotSchedule.md) |  | 
**VolumeId** | **NullableString** |  | 

## Methods

### NewSnapshotScheduleSetResponse

`func NewSnapshotScheduleSetResponse(snapshotSchedule SnapshotSchedule, volumeId NullableString, ) *SnapshotScheduleSetResponse`

NewSnapshotScheduleSetResponse instantiates a new SnapshotScheduleSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotScheduleSetResponseWithDefaults

`func NewSnapshotScheduleSetResponseWithDefaults() *SnapshotScheduleSetResponse`

NewSnapshotScheduleSetResponseWithDefaults instantiates a new SnapshotScheduleSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotRetentionCount

`func (o *SnapshotScheduleSetResponse) GetSnapshotRetentionCount() int32`

GetSnapshotRetentionCount returns the SnapshotRetentionCount field if non-nil, zero value otherwise.

### GetSnapshotRetentionCountOk

`func (o *SnapshotScheduleSetResponse) GetSnapshotRetentionCountOk() (*int32, bool)`

GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionCount

`func (o *SnapshotScheduleSetResponse) SetSnapshotRetentionCount(v int32)`

SetSnapshotRetentionCount sets SnapshotRetentionCount field to given value.

### HasSnapshotRetentionCount

`func (o *SnapshotScheduleSetResponse) HasSnapshotRetentionCount() bool`

HasSnapshotRetentionCount returns a boolean if a field has been set.

### GetSnapshotSchedule

`func (o *SnapshotScheduleSetResponse) GetSnapshotSchedule() SnapshotSchedule`

GetSnapshotSchedule returns the SnapshotSchedule field if non-nil, zero value otherwise.

### GetSnapshotScheduleOk

`func (o *SnapshotScheduleSetResponse) GetSnapshotScheduleOk() (*SnapshotSchedule, bool)`

GetSnapshotScheduleOk returns a tuple with the SnapshotSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSchedule

`func (o *SnapshotScheduleSetResponse) SetSnapshotSchedule(v SnapshotSchedule)`

SetSnapshotSchedule sets SnapshotSchedule field to given value.


### GetVolumeId

`func (o *SnapshotScheduleSetResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *SnapshotScheduleSetResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *SnapshotScheduleSetResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.


### SetVolumeIdNil

`func (o *SnapshotScheduleSetResponse) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *SnapshotScheduleSetResponse) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


