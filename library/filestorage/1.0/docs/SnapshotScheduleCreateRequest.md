# SnapshotScheduleCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotRetentionCount** | Pointer to **NullableInt32** |  | [optional] 
**SnapshotSchedule** | [**SnapshotSchedule**](SnapshotSchedule.md) |  | 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewSnapshotScheduleCreateRequest

`func NewSnapshotScheduleCreateRequest(snapshotSchedule SnapshotSchedule, volumeId string, ) *SnapshotScheduleCreateRequest`

NewSnapshotScheduleCreateRequest instantiates a new SnapshotScheduleCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotScheduleCreateRequestWithDefaults

`func NewSnapshotScheduleCreateRequestWithDefaults() *SnapshotScheduleCreateRequest`

NewSnapshotScheduleCreateRequestWithDefaults instantiates a new SnapshotScheduleCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotRetentionCount

`func (o *SnapshotScheduleCreateRequest) GetSnapshotRetentionCount() int32`

GetSnapshotRetentionCount returns the SnapshotRetentionCount field if non-nil, zero value otherwise.

### GetSnapshotRetentionCountOk

`func (o *SnapshotScheduleCreateRequest) GetSnapshotRetentionCountOk() (*int32, bool)`

GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionCount

`func (o *SnapshotScheduleCreateRequest) SetSnapshotRetentionCount(v int32)`

SetSnapshotRetentionCount sets SnapshotRetentionCount field to given value.

### HasSnapshotRetentionCount

`func (o *SnapshotScheduleCreateRequest) HasSnapshotRetentionCount() bool`

HasSnapshotRetentionCount returns a boolean if a field has been set.

### SetSnapshotRetentionCountNil

`func (o *SnapshotScheduleCreateRequest) SetSnapshotRetentionCountNil(b bool)`

 SetSnapshotRetentionCountNil sets the value for SnapshotRetentionCount to be an explicit nil

### UnsetSnapshotRetentionCount
`func (o *SnapshotScheduleCreateRequest) UnsetSnapshotRetentionCount()`

UnsetSnapshotRetentionCount ensures that no value is present for SnapshotRetentionCount, not even an explicit nil
### GetSnapshotSchedule

`func (o *SnapshotScheduleCreateRequest) GetSnapshotSchedule() SnapshotSchedule`

GetSnapshotSchedule returns the SnapshotSchedule field if non-nil, zero value otherwise.

### GetSnapshotScheduleOk

`func (o *SnapshotScheduleCreateRequest) GetSnapshotScheduleOk() (*SnapshotSchedule, bool)`

GetSnapshotScheduleOk returns a tuple with the SnapshotSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSchedule

`func (o *SnapshotScheduleCreateRequest) SetSnapshotSchedule(v SnapshotSchedule)`

SetSnapshotSchedule sets SnapshotSchedule field to given value.


### GetVolumeId

`func (o *SnapshotScheduleCreateRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *SnapshotScheduleCreateRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *SnapshotScheduleCreateRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


