# SnapshotScheduleCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotRetentionCount** | Pointer to **NullableString** |  | [optional] 
**SnapshotSchedule** | [**SnapshotSchedule**](SnapshotSchedule.md) |  | 
**VolumeId** | **NullableString** |  | 

## Methods

### NewSnapshotScheduleCreateResponse

`func NewSnapshotScheduleCreateResponse(snapshotSchedule SnapshotSchedule, volumeId NullableString, ) *SnapshotScheduleCreateResponse`

NewSnapshotScheduleCreateResponse instantiates a new SnapshotScheduleCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotScheduleCreateResponseWithDefaults

`func NewSnapshotScheduleCreateResponseWithDefaults() *SnapshotScheduleCreateResponse`

NewSnapshotScheduleCreateResponseWithDefaults instantiates a new SnapshotScheduleCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotRetentionCount

`func (o *SnapshotScheduleCreateResponse) GetSnapshotRetentionCount() string`

GetSnapshotRetentionCount returns the SnapshotRetentionCount field if non-nil, zero value otherwise.

### GetSnapshotRetentionCountOk

`func (o *SnapshotScheduleCreateResponse) GetSnapshotRetentionCountOk() (*string, bool)`

GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionCount

`func (o *SnapshotScheduleCreateResponse) SetSnapshotRetentionCount(v string)`

SetSnapshotRetentionCount sets SnapshotRetentionCount field to given value.

### HasSnapshotRetentionCount

`func (o *SnapshotScheduleCreateResponse) HasSnapshotRetentionCount() bool`

HasSnapshotRetentionCount returns a boolean if a field has been set.

### SetSnapshotRetentionCountNil

`func (o *SnapshotScheduleCreateResponse) SetSnapshotRetentionCountNil(b bool)`

 SetSnapshotRetentionCountNil sets the value for SnapshotRetentionCount to be an explicit nil

### UnsetSnapshotRetentionCount
`func (o *SnapshotScheduleCreateResponse) UnsetSnapshotRetentionCount()`

UnsetSnapshotRetentionCount ensures that no value is present for SnapshotRetentionCount, not even an explicit nil
### GetSnapshotSchedule

`func (o *SnapshotScheduleCreateResponse) GetSnapshotSchedule() SnapshotSchedule`

GetSnapshotSchedule returns the SnapshotSchedule field if non-nil, zero value otherwise.

### GetSnapshotScheduleOk

`func (o *SnapshotScheduleCreateResponse) GetSnapshotScheduleOk() (*SnapshotSchedule, bool)`

GetSnapshotScheduleOk returns a tuple with the SnapshotSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSchedule

`func (o *SnapshotScheduleCreateResponse) SetSnapshotSchedule(v SnapshotSchedule)`

SetSnapshotSchedule sets SnapshotSchedule field to given value.


### GetVolumeId

`func (o *SnapshotScheduleCreateResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *SnapshotScheduleCreateResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *SnapshotScheduleCreateResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.


### SetVolumeIdNil

`func (o *SnapshotScheduleCreateResponse) SetVolumeIdNil(b bool)`

 SetVolumeIdNil sets the value for VolumeId to be an explicit nil

### UnsetVolumeId
`func (o *SnapshotScheduleCreateResponse) UnsetVolumeId()`

UnsetVolumeId ensures that no value is present for VolumeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


