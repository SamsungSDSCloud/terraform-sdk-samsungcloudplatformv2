# SnapshotScheduleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotPolicyEnabled** | Pointer to **NullableBool** |  | [optional] 
**SnapshotRetentionCount** | Pointer to **NullableString** |  | [optional] 
**SnapshotSchedule** | [**[]SnapshotScheduleShow**](SnapshotScheduleShow.md) |  | 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewSnapshotScheduleListResponse

`func NewSnapshotScheduleListResponse(snapshotSchedule []SnapshotScheduleShow, volumeId string, ) *SnapshotScheduleListResponse`

NewSnapshotScheduleListResponse instantiates a new SnapshotScheduleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotScheduleListResponseWithDefaults

`func NewSnapshotScheduleListResponseWithDefaults() *SnapshotScheduleListResponse`

NewSnapshotScheduleListResponseWithDefaults instantiates a new SnapshotScheduleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotPolicyEnabled

`func (o *SnapshotScheduleListResponse) GetSnapshotPolicyEnabled() bool`

GetSnapshotPolicyEnabled returns the SnapshotPolicyEnabled field if non-nil, zero value otherwise.

### GetSnapshotPolicyEnabledOk

`func (o *SnapshotScheduleListResponse) GetSnapshotPolicyEnabledOk() (*bool, bool)`

GetSnapshotPolicyEnabledOk returns a tuple with the SnapshotPolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotPolicyEnabled

`func (o *SnapshotScheduleListResponse) SetSnapshotPolicyEnabled(v bool)`

SetSnapshotPolicyEnabled sets SnapshotPolicyEnabled field to given value.

### HasSnapshotPolicyEnabled

`func (o *SnapshotScheduleListResponse) HasSnapshotPolicyEnabled() bool`

HasSnapshotPolicyEnabled returns a boolean if a field has been set.

### SetSnapshotPolicyEnabledNil

`func (o *SnapshotScheduleListResponse) SetSnapshotPolicyEnabledNil(b bool)`

 SetSnapshotPolicyEnabledNil sets the value for SnapshotPolicyEnabled to be an explicit nil

### UnsetSnapshotPolicyEnabled
`func (o *SnapshotScheduleListResponse) UnsetSnapshotPolicyEnabled()`

UnsetSnapshotPolicyEnabled ensures that no value is present for SnapshotPolicyEnabled, not even an explicit nil
### GetSnapshotRetentionCount

`func (o *SnapshotScheduleListResponse) GetSnapshotRetentionCount() string`

GetSnapshotRetentionCount returns the SnapshotRetentionCount field if non-nil, zero value otherwise.

### GetSnapshotRetentionCountOk

`func (o *SnapshotScheduleListResponse) GetSnapshotRetentionCountOk() (*string, bool)`

GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionCount

`func (o *SnapshotScheduleListResponse) SetSnapshotRetentionCount(v string)`

SetSnapshotRetentionCount sets SnapshotRetentionCount field to given value.

### HasSnapshotRetentionCount

`func (o *SnapshotScheduleListResponse) HasSnapshotRetentionCount() bool`

HasSnapshotRetentionCount returns a boolean if a field has been set.

### SetSnapshotRetentionCountNil

`func (o *SnapshotScheduleListResponse) SetSnapshotRetentionCountNil(b bool)`

 SetSnapshotRetentionCountNil sets the value for SnapshotRetentionCount to be an explicit nil

### UnsetSnapshotRetentionCount
`func (o *SnapshotScheduleListResponse) UnsetSnapshotRetentionCount()`

UnsetSnapshotRetentionCount ensures that no value is present for SnapshotRetentionCount, not even an explicit nil
### GetSnapshotSchedule

`func (o *SnapshotScheduleListResponse) GetSnapshotSchedule() []SnapshotScheduleShow`

GetSnapshotSchedule returns the SnapshotSchedule field if non-nil, zero value otherwise.

### GetSnapshotScheduleOk

`func (o *SnapshotScheduleListResponse) GetSnapshotScheduleOk() (*[]SnapshotScheduleShow, bool)`

GetSnapshotScheduleOk returns a tuple with the SnapshotSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSchedule

`func (o *SnapshotScheduleListResponse) SetSnapshotSchedule(v []SnapshotScheduleShow)`

SetSnapshotSchedule sets SnapshotSchedule field to given value.


### SetSnapshotScheduleNil

`func (o *SnapshotScheduleListResponse) SetSnapshotScheduleNil(b bool)`

 SetSnapshotScheduleNil sets the value for SnapshotSchedule to be an explicit nil

### UnsetSnapshotSchedule
`func (o *SnapshotScheduleListResponse) UnsetSnapshotSchedule()`

UnsetSnapshotSchedule ensures that no value is present for SnapshotSchedule, not even an explicit nil
### GetVolumeId

`func (o *SnapshotScheduleListResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *SnapshotScheduleListResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *SnapshotScheduleListResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


